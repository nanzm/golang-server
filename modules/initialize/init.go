package initialize

import (
	"dora/app/manage/model/dao"
	"dora/app/manage/model/entity"
	"dora/config"
	"dora/modules/datasource/elastic"
	"dora/modules/datasource/gorm"
	"dora/pkg/utils/logx"
	"fmt"
	"strings"
)

func Run() {
	dbMigrate()
	createRoles()
	createUser()
	createDocMapping()
}

// 表同步
func dbMigrate() {
	err := gorm.Instance().AutoMigrate(
		&entity.SysLog{},

		&entity.Project{},
		&entity.Role{},
		&entity.User{},
		&entity.UserSetting{},

		&entity.Issue{},
		&entity.IssueUserStatus{},
		&entity.SourceMap{},

		&entity.Artifact{},
		&entity.AlarmLog{},
	)
	if err != nil {
		panic(err)
	}
}

// 创建 roles
func createRoles() {
	var roles = make([]entity.Role, 0)
	err := gorm.Instance().Limit(10).Find(&roles).Error
	if err != nil {
		panic(err)
	}
	if len(roles) > 0 {
		return
	}

	data := []entity.Role{
		{
			Key:     "admin",
			Name:    "管理员",
			Remarks: "",
		}, {
			Key:     "developer",
			Name:    "开发者",
			Remarks: "",
		}, {
			Key:     "visitor",
			Name:    "游客",
			Remarks: "",
		}}

	err = gorm.Instance().Model(&roles).Create(&data).Error
	if err != nil {
		panic(err)
	}
	logx.Println("默认角色已创建！")
}

// 创建 admin
func createUser() {
	var user entity.User
	err := gorm.Instance().Where("email = ?", "demo@dora.com").Find(&user).Error
	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		// 创建默认用户
		admin := entity.User{NickName: "live demo", Email: "demo@dora.com", Password: "123", RoleId: 1}
		logx.Println("-------------------------------------")
		logx.Printf("初始化用户：%v 密码：%v", admin.Email, admin.Password)
		logx.Println("-------------------------------------")
		err = gorm.Instance().Create(&admin).Error
		if err != nil {
			panic(err)
		}

		// 创建项目
		project := dao.NewProjectDao()
		pro := entity.Project{
			AppId:             "44992867-5a85-4804-849a-d525be1fa77c",
			Name:              "demo",
			Alias:             "live demo",
			Type:              "browser",
			GitRepositoryUrl:  "",
			GitRepositoryName: "",
		}
		demoProject, err := project.Create(&pro, admin.ID)
		if err != nil {
			panic(err)
		}

		// 默认设置
		setting := dao.NewUserSettingDao()
		err = setting.UpdateOrCreate(admin.ID, demoProject.ID)
		if err != nil {
			panic(err)
		}

		return
	}
}

func createDocMapping() {
	enable := config.GetLogStore().Enable
	if enable == "elastic" {
		es := elastic.GetClient()

		conf := config.GetElastic()
		doc := conf.Index

		exists, _ := es.Indices.Exists([]string{doc})
		fmt.Printf("%v \n", exists)

		if exists != nil && exists.StatusCode == 200 {
			fmt.Printf("%v \n", "")
			logx.Infof("elastic docs %v has exists", doc)
			return
		}

		logx.Infof("elastic need create doc %s", doc)

		_, err := es.Indices.Create(doc,
			es.Indices.Create.WithBody(strings.NewReader(elasticMapping)))
		if err != nil {
			logx.Fatalf("elastic docs create error %s", err)
		}
		logx.Infof("elastic docs %v has created", doc)
	}
}
