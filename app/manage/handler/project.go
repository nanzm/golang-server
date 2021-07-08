package handler

import (
	"bytes"
	"dora/app/manage/model/dao"
	"dora/app/manage/model/dto"
	"dora/app/manage/model/entity"
	"dora/config"
	"dora/modules/middleware"
	"dora/pkg/utils"
	"dora/pkg/utils/fs"
	"dora/pkg/utils/ginutil"
	"dora/pkg/utils/unarchive"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
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
	router.POST("/project/upload/backup", pro.BackupUpload)
	router.GET("/project/backup", pro.BackupList)

	// sourcemap 上传
	router.POST("/project/upload/sourcemap", pro.SourcemapUpload)
	router.GET("/project/sourcemap", pro.SourcemapList)
	router.DELETE("/project/sourcemap", pro.SourcemapDelete)

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

func (pro *ProjectResource) BackupUpload(c *gin.Context) {
	var u dto.BackUpParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	// 文件存储目录
	destDir := config.BackupDir + "/" + u.AppId
	err := fs.EnsureDir(destDir)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	nowTimeStr := utils.CurrentTimePathFriendly()
	fileName := fmt.Sprintf("%s_%s_%s", u.ProjectName, nowTimeStr, u.FileName)
	fileDest := fmt.Sprintf("%v/%v", destDir, fileName)

	if err = c.SaveUploadedFile(u.File, fileDest); err != nil {
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
		FilePath: fileDest,

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

	ginutil.JSONOk(c, fileDest)
}

func (pro *ProjectResource) BackupList(c *gin.Context) {
	var u dto.BackUpListParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	artifactDao := dao.NewArtifactDao()
	list, count, err := artifactDao.List(u.Current, u.PageSize, u.AppId, u.FileType)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONListPages(c, list, u.Current, u.PageSize, count)
}

func (pro *ProjectResource) SourcemapUpload(c *gin.Context) {
	var u dto.UploadSourcemapParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	destDir := config.SourcemapDir + "/" + u.AppId
	err := fs.EnsureDir(destDir)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	nowTimeStr := utils.CurrentTimePathFriendly()
	fileName := fmt.Sprintf("%s_%s_%s", u.ProjectName, nowTimeStr, u.FileName)
	fileDest := fmt.Sprintf("%v/%v", destDir, fileName)

	// 保存
	if err := c.SaveUploadedFile(u.File, fileDest); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 解压
	err = unarchive.Save(fileDest, strings.TrimSuffix(fileDest, filepath.Ext(fileDest)))
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	// 记录到 db
	sourcemapDao := dao.NewSourcemapDao()
	_, err = sourcemapDao.Create(&entity.Sourcemap{
		AppId:      u.AppId,
		AppVersion: u.AppVersion,
		Path:       fileDest,
		Size:       fs.FileSize(fileDest),
	})
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, fileDest)
}

func (pro *ProjectResource) SourcemapList(c *gin.Context) {
	var u dto.SourcemapListParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	sourcemapDao := dao.NewSourcemapDao()
	list, total, err := sourcemapDao.List(u.Current, u.PageSize, u.AppId)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONListPages(c, list, u.Current, u.PageSize, total)
}

func (pro *ProjectResource) SourcemapDelete(c *gin.Context) {
	var u dto.SourcemapDeleteParam
	if err := c.ShouldBind(&u); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	sourcemapDao := dao.NewSourcemapDao()
	err := sourcemapDao.Delete(u.Id)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONOk(c, "ok")
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
