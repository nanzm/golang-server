package boots

import (
	"dora/app/constant"
	"dora/app/datasource"
	"dora/app/model"
	"dora/app/service"
	"dora/pkg/logger"
)

func Run() {
	dbMigrate()
	createRoles()
	createUser()
	putMd5ListToCache()
}

// 表同步
func dbMigrate() {
	err := datasource.Migrate(datasource.GormInstance(), model.Tables())
	if err != nil {
		panic(err)
	}
}

// 创建 roles
func createRoles() {
	var roles = make([]model.Role, 0)
	err := datasource.GormInstance().Limit(10).Find(&roles).Error
	if err != nil {
		panic(err)
	}
	if len(roles) > 0 {
		return
	}

	data := []model.Role{
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
	logger.Println("默认角色已创建！")
}

// 创建 admin
func createUser() {
	var user model.User
	err := datasource.GormInstance().Where("email = ?", "demo@dora.com").Find(&user).Error
	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		admin := model.User{NickName: "demo_auto_generate", Email: "demo@dora.com", Password: "123", RoleId: 1}
		logger.Println("-------------------------------------")
		logger.Printf("初始化用户：%v 密码：%v", admin.Email, admin.Password)
		logger.Println("-------------------------------------")
		err = datasource.GormInstance().Create(&admin).Error
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
		logger.Infof("none issues md5 values")
	}
}
