package service

import (
	"context"
	"dora/internal/config"
	mailRes "dora/internal/datasource/mail"
	dataRedis "dora/internal/datasource/redis"
	"dora/pkg/utils"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type UserService interface {
	SendEmailCaptcha(captchaType string, toEmail string) error
	VerifyCaptcha(captchaType string, toEmail string, Captcha string) (verify bool, e error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{
	}
}

func (u *userService) SendEmailCaptcha(captchaType string, toEmail string) error {
	ctx := context.Background()

	key := fmt.Sprintf("%s_%s", captchaType, toEmail)

	result, err := dataRedis.Instance().Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "redisutil Get 失败 %s", key)
	}
	if result != "" {
		return errors.New("验证码发送太频繁")
	}
	// 设置新的
	code, err := utils.EncodeToString(6)
	if err != nil {
		return errors.Wrap(err, "验证码 EncodeToString 报错")
	}

	err = dataRedis.Instance().Set(ctx, key, code, 60*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, "redisutil Set 失败 %s", key)
	}

	// 发送邮箱验证码
	conf := config.GetMail()
	m := mailRes.BuilderEmail(toEmail, fmt.Sprintf("Dora Robot <%s>", conf.Username),
		"Dora 登录验证码", fmt.Sprintf("您的登录验证码是：<h1>%s</h1>", code))
	err = mailRes.GetPool().Send(m, 3*time.Second)

	if err != nil {
		return errors.Wrap(err, "发送失败")
	}
	return nil
}

func (u *userService) VerifyCaptcha(captchaType string, toEmail string, Captcha string) (verify bool, e error) {
	ctx := context.Background()

	key := fmt.Sprintf("%s_%s", captchaType, toEmail)
	result, err := dataRedis.Instance().Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, errors.Wrapf(err, "redisutil Get 失败 %s", key)
	}
	if result == "" {
		return false, errors.New("验证码已过期")
	}

	if result != Captcha {
		return false, errors.New("验证码错误")
	}

	return true, nil
}
