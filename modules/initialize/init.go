package initialize

import (
	"dora/config/constant"
	"dora/modules/datasource"
	"dora/modules/model/dao"
	"dora/modules/model/entity"
	"dora/modules/service"
	"dora/pkg/utils/logx"
)

func Run() {
	dbMigrate()
	createRoles()
	createUser()
	putMd5ListToCache()
}

// 表同步
func dbMigrate() {
	err := datasource.Migrate(datasource.GormInstance(), entity.Tables())
	if err != nil {
		panic(err)
	}
}

// 创建 roles
func createRoles() {
	var roles = make([]entity.Role, 0)
	err := datasource.GormInstance().Limit(10).Find(&roles).Error
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

	err = datasource.GormInstance().Model(&roles).Create(&data).Error
	if err != nil {
		panic(err)
	}
	logx.Println("默认角色已创建！")
}

// 创建 admin
func createUser() {
	var user entity.User
	err := datasource.GormInstance().Where("email = ?", "demo@dora.com").Find(&user).Error
	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		// 创建默认用户
		admin := entity.User{NickName: "live demo", Email: "demo@dora.com", Password: "123", RoleId: 1}
		logx.Println("-------------------------------------")
		logx.Printf("初始化用户：%v 密码：%v", admin.Email, admin.Password)
		logx.Println("-------------------------------------")
		err = datasource.GormInstance().Create(&admin).Error
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

// 将所有 md5 放入 redis
func putMd5ListToCache() {
	issues := service.NewIssuesService()
	md5s := issues.GetAllMd5()
	if len(md5s) > 0 {
		datasource.RedisSetAdd(constant.Md5ListHas, md5s)
	} else {
		logx.Infof("none issues md5 values")
	}
}
