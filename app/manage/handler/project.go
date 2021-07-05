package handler

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"dora/app/manage/model/dao"
	"dora/app/manage/model/dto"
	"dora/app/manage/model/entity"
	"dora/config"
	"dora/modules/middleware"
	"dora/pkg/utils"
	"dora/pkg/utils/ginutil"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mholt/archiver/v3"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type ProjectResource struct {
}

func NewProjectResource() ginutil.Resource {
	return &ProjectResource{
	}
}

func (pro *ProjectResource) Register(router *gin.RouterGroup) {
	router.GET("/project", middleware.JWTAuthMiddleware(), pro.Get)
	router.POST("/project", middleware.JWTAuthMiddleware(), pro.Create)
	router.GET("/project/users", middleware.JWTAuthMiddleware(), pro.GetProjectUsers)

	// 备份项目
	// 1、 xx.all.zip
	// 2、 xx.prod.zip
	// 3、 xx.sourcemap.zip
	router.POST("/project/upload/backup", pro.UploadBackup)

	// sourcemap 上传
	router.POST("/project/upload/sourcemap", pro.UploadSourcemap)

	router.POST("/project/sourcemap/parse", pro.SourcemapParse)
}

func (pro *ProjectResource) Get(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.Get(u.ProjectId)

	// not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (pro *ProjectResource) Create(c *gin.Context) {
	uid, _ := c.Get("uid")
	var body dto.CreateProject
	if err := c.ShouldBind(&body); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	project := entity.Project{
		AppId:             uuid.New().String(),
		Name:              body.Name,
		Alias:             body.Alias,
		Type:              body.Type,
		GitRepositoryUrl:  body.GitRepositoryUrl,
		GitRepositoryName: body.GitRepositoryName,
	}

	d := dao.NewProjectDao()

	// 校验项目名 名字
	p, err := d.GetByName(body.Name)
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	if p.ID != 0 {
		ginutil.JSONFail(c, -1, "该项目名已存在")
		return
	}

	// 创建
	result, err := d.Create(&project, uid.(uint))
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	// 切换到当前项目
	settingDao := dao.NewUserSettingDao()
	err = settingDao.UpdateOrCreate(uid.(uint), result.ID)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, result)
}

func (pro *ProjectResource) UploadBackup(c *gin.Context) {
	var u dto.BackUpParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	// 文件存储目录
	destDir := config.BackupDir + "/" + u.AppId
	_, err := os.Stat(destDir)
	if err != nil {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
	}
	nowTimeStr := utils.CurrentTimePathFriendly()
	fileName := fmt.Sprintf("%s_%s_%s", u.ProjectName, nowTimeStr, u.FileName)
	fileURL := fmt.Sprintf("%v/%v", destDir, fileName)

	if err = c.SaveUploadedFile(u.File, fileURL); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 记录到 db
	artifactDao := dao.NewArtifactDao()
	_, err = artifactDao.Create(&entity.Artifact{
		AppId:       u.AppId,
		ProjectName: u.ProjectName,

		FileName: fileName,
		FileType: u.FileType,
		FilePath: fileURL,

		GitName:   u.GitName,
		GitEmail:  u.GitEmail,
		GitBranch: u.GitBranch,

		Commit:    u.Commit,
		CommitSha: u.CommitSha,
		CommitTs:  u.CommitTs,
	})
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, fileURL)
}

func (pro *ProjectResource) UploadSourcemap(c *gin.Context) {
	var u dto.UploadSourcemapParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	// 文件存储目录
	destDir := "tmp/sourcemap/" + u.AppId
	// 文件解压存储目录
	decompressDestDir := "tmp/sourcemap/" + u.AppId + "/decompress"

	_, err := os.Stat(destDir)
	if err != nil {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
	}

	fileURL := destDir + "/" + u.File.Filename
	if err = c.SaveUploadedFile(u.File, fileURL); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 解压
	if path.Ext(fileURL) == ".zip" {
		z := archiver.Zip{
			OverwriteExisting:    true,
			MkdirAll:             true,
			SelectiveCompression: true,
			CompressionLevel:     flate.DefaultCompression,
			FileMethod:           archiver.Deflate,
		}
		err := z.Unarchive(fileURL, decompressDestDir)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
	}

	if path.Ext(fileURL) == ".gz" {
		gz := archiver.TarGz{
			CompressionLevel: gzip.DefaultCompression,
			Tar: &archiver.Tar{
				OverwriteExisting: true,
				MkdirAll:          true,
			},
		}
		err := gz.Unarchive(fileURL, decompressDestDir)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
	}

	ginutil.JSONOk(c, fileURL)
}

func (pro *ProjectResource) SourcemapParse(c *gin.Context) {
	var u dto.SourcemapParseParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	sourcemap, err := utils.GetStackSourceMap("tmp/sourcemap/"+u.AppId+"/decompress", u.Stack)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	body := dto.ReqPostParseData{
		Stack:        u.Stack,
		RawSourcemap: string(sourcemap),
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8220", bytes.NewBuffer(marshal))
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 解析响应
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 返回
	var result map[string]interface{}
	err = json.Unmarshal(s, &result)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONOk(c, result)
}

func (pro *ProjectResource) GetProjectUsers(c *gin.Context) {
	var u dto.QueryDetail
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}

	d := dao.NewProjectDao()
	get, err := d.ProjectUsers(u.ProjectId)
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, get)
}
