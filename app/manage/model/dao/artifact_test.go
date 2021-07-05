package dao

import (
	"dora/app/manage/model/entity"
	"dora/pkg/utils"
	"testing"
)

func TestArtifact_Create(t *testing.T) {
	dao := NewArtifactDao()
	create, err := dao.Create(&entity.Artifact{
		AppId:       "dfb38446941dcf83aa5173d430f6",
		ProjectName: "dadada",
		FileName:    "dist.zip",
		FileType:    "full",
		FilePath:    "store/test/",
		GitName:     "nan",
		GitEmail:    "nan@nancode.cn",
		GitBranch:   "dev",
		Commit:      "feat: add rate middleware",
		CommitSha:   "dfb38446941dca6f2d67919b4f83aa5173d430f6",
		CommitTs:    "July 4, 2021 at 21:06:18 GMT+8",
	})

	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(create)
}

func TestArtifact_Delete(t *testing.T) {

}

func TestArtifact_List(t *testing.T) {

}

func TestNewArtifactDao(t *testing.T) {

}
