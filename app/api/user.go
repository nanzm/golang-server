package api

import (
	"dora/app/dao"
	"dora/app/datasource"
	"dora/app/dto"
	"dora/app/middleware"
	"dora/app/model"
	"dora/app/service"
	"dora/config"
	"dora/pkg/ginutil"
	"dora/pkg/jwtutil"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const userSessionKey = "user"

type UserResource struct {
	Conf *config.Conf
}

func NewUserResource() Resource {
	return &UserResource{
		Conf: config.GetConf(),
	}
}

func (u *UserResource) Register(router *gin.RouterGroup) {
	router.POST("/user/signup", u.SingUp)
	router.POST("/user/login", u.UserLogin)
	router.POST("/email/code/login", u.EmailCodeLogin)
	router.POST("/captcha", u.Captcha)
	router.GET("/user/loginOut", u.LoginOut)

	router.GET("/user/info", middleware.JWTAuthMiddleware(), u.Info)
	router.POST("/user/info/update", middleware.JWTAuthMiddleware(), u.UpdateInfo)
	router.POST("/user/status", middleware.JWTAuthMiddleware(), u.Status)
	router.GET("/user/list", middleware.JWTAuthMiddleware(), u.List)

	router.GET("/user/projects", middleware.JWTAuthMiddleware(), u.UseProjects)
	router.GET("/user/setting", middleware.JWTAuthMiddleware(), u.Setting)
	router.POST("/user/setting/update", middleware.JWTAuthMiddleware(), u.UpdateSetting)
}

func (u *UserResource) SingUp(c *gin.Context) {
	var p dto.SignUpParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	userDao := dao.NewUserDao()
	exist, err := userDao.GetByEmail(p.Email)
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	utils.PrettyPrint(exist)
	if exist.ID != 0 {
		ginutil.JSONFail(c, -1, "该邮箱已注册")
		return
	}

	// 保存入库
	newUser := &model.User{
		NickName: p.NickName,
		Email:    p.Email,
		Password: p.Password,
		Status:   1,
		RoleId:   3,
	}

	created, err := userDao.Create(newUser)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	ginutil.JSONOk(c, created)
}

func (u *UserResource) UserLogin(c *gin.Context) {
	var p dto.LoginParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	userDao := dao.NewUserDao()
	user, err := userDao.GetByEmail(p.Email)
	if err != nil {
		ginutil.JSONFail(c, -1, "账号或密码错误")
		return
	}
	if user.ID == 0 {
		ginutil.JSONFail(c, -1, "该账号未注册")
		return
	}
	if user.Password != p.Password {
		ginutil.JSONFail(c, -1, "用户名或密码错误")
		return
	}

	// token
	token, err := jwtutil.GenToken(user.ID, user.Email)
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	ginutil.JSONOk(c, dto.UserLoginVo{
		Token: token,
		User:  user,
	})
}

func (u *UserResource) EmailCodeLogin(c *gin.Context) {
	var p dto.EmailLoginParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	var user model.User

	err := datasource.GormInstance().First(&user, model.User{Email: p.Email}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, -1, "该账号未注册")
		return
	}
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	userService := service.NewUserService(u.Conf)
	_, err = userService.VerifyCaptcha("login", p.Email, p.Captcha)
	if err != nil {
		ginutil.JSONFail(c, -1, err.Error())
		return
	}

	// cookie
	session := sessions.Default(c)
	session.Set(userSessionKey, user)
	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 1, // 1hrs
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	err = session.Save()
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, user)
}

func (u *UserResource) Captcha(c *gin.Context) {
	var p dto.CaptchaParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	var user model.User

	db := datasource.GormInstance().First(&user, model.User{
		Email: p.Email,
	})

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, -1, "该账号未注册")
		return
	}

	if db.Error != nil {
		ginutil.JSONFail(c, -1, "账号或密码错误")
		return
	}

	// 发送验证码
	codeString, err := utils.EncodeToString(6)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	logger.Printf("%v \n", codeString)

	userService := service.NewUserService(u.Conf)
	err = userService.SendEmailCaptcha(p.Type, p.Email)
	if err != nil {
		ginutil.JSONFail(c, http.StatusBadRequest, err.Error())
		return
	}

	ginutil.JSONOk(c, nil, "验证码已发送")
}

func (u *UserResource) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userSessionKey)
	err := session.Save()
	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, nil)
}

func (u *UserResource) Info(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		ginutil.JSONFail(c, -1, "请重新登录")
		return
	}

	userDao := dao.NewUserDao()
	userInfo, err := userDao.Get(uid.(uint))
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	ginutil.JSONOk(c, userInfo)
}

func (u *UserResource) UpdateSession(c *gin.Context) {
	get, exists := c.Get(userSessionKey)
	if !exists {
		ginutil.JSONFail(c, -1, "请重新登录")
		return
	}

	// 去数据库查最新的
	var user model.User
	first := datasource.GormInstance().First(&user, model.User{
		ID: get.(model.User).ID,
	})
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, -1, "不存在的用户id")
		return
	}

	// 更新 session
	session := sessions.Default(c)
	session.Set(userSessionKey, user)
	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 1, // 1hrs
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	err := session.Save()
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONOk(c, get)
}

func (u *UserResource) UpdateInfo(c *gin.Context) {
	var p dto.UpdateParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	var user model.User
	first := datasource.GormInstance().First(&user, model.User{
		ID: p.Id,
	})
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		ginutil.JSONFail(c, -1, "不存在的用户id")
		return
	}

	// 更新
	err := datasource.GormInstance().Model(&user).Updates(model.User{
		NickName: p.NickName,
		Email:    p.Email,
		Password: p.Password,
	}).Error

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err, "更新异常")
		return
	}

	session := sessions.Default(c)
	userCache := session.Get(userSessionKey)

	// 如果是本人在修改 更新session 及时刷新
	if userCache.(model.User).ID == p.Id {
		session.Set(userSessionKey, user)
		session.Options(sessions.Options{
			Path:     "/",
			MaxAge:   3600 * 1, // 1hrs
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		})
		session.Save()
	}

	ginutil.JSONOk(c, user)
}

func (u *UserResource) Status(c *gin.Context) {
	var p dto.StatusParam
	if err := c.ShouldBind(&p); err != nil {
		errorTrans(c, err)
		return
	}

	var user model.User
	first := datasource.GormInstance().First(&user, model.User{
		ID: p.Id,
	})
	if first.Error != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, first.Error, "不存在的用户id")
		return
	}

	// 更新 Status
	err := datasource.GormInstance().Model(&user).Updates(model.User{
		Status: p.Status,
	}).Error

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONOk(c, user)
}

func (u *UserResource) List(c *gin.Context) {
	var p dto.ListSearchParam
	if err := c.ShouldBind(&p); err != nil {
		logger.Printf("%v \n", err)
		return
	}

	var users []model.User
	var total int64

	current := 1
	if p.Current > 0 {
		current = p.Current
	}

	size := 10
	if p.PageSize > 0 {
		size = p.PageSize
	}

	db := datasource.GormInstance().Limit(size).Offset((current - 1) * size)
	var err error
	cond := "%" + p.SearchStr + "%"
	logger.Printf("%v\n", cond)

	if p.SearchStr != "" {
		err = db.Find(&users, "nick_name LIKE ? Or email LIKE ? Or phone LIKE ? Or username LIKE ?",
			cond, cond, cond, cond).Count(&total).Error
	} else {
		err = db.Find(&users).Count(&total).Error
	}

	if err != nil {
		ginutil.JSONError(c, http.StatusInternalServerError, err)
		return
	}

	ginutil.JSONList(c, users, int(total))
}

func (u *UserResource) UseProjects(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		ginutil.JSONFail(c, -1, "请重新登录")
		return
	}

	userDao := dao.NewUserDao()
	userInfo, err := userDao.UserProjects(uid.(uint))
	if err != nil {
		ginutil.JSONServerError(c, err)
	}
	ginutil.JSONOk(c, userInfo)
}

func (u *UserResource) Setting(c *gin.Context) {
	uid, _ := c.Get("uid")

	settingDao := dao.NewUserSettingDao()
	vo, err := settingDao.Get(uid.(uint))
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	if vo.ID == 0 {
		ginutil.JSONOk(c, nil)
		return
	}
	ginutil.JSONOk(c, vo)
}

func (u *UserResource) UpdateSetting(c *gin.Context) {
	var r dto.UpdateDefaultSettingReq
	if err := c.ShouldBind(&r); err != nil {
		errorTrans(c, err)
		return
	}

	get, _ := c.Get("uid")
	settingDao := dao.NewUserSettingDao()
	err := settingDao.UpdateOrCreate(get.(uint), r.ProjectId)
	if err != nil {
		ginutil.JSONError(c, http.StatusBadRequest, err)
		return
	}
	ginutil.JSONOk(c, nil)
}
