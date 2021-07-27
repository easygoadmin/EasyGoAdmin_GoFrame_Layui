/**
 *
 * @author 摆渡人
 * @since 2021/3/3
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var User = new(userService)

type userService struct{}

var SessionList = gmap.New(true)

// 系统登录
func (s *userService) SignIn(username, password string, session *ghttp.Session) error {
	// 获取用户信息
	user, err := dao.User.FindOne("username=? and mark=1", username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd := gmd5.MustEncryptString(gmd5.MustEncryptString(password + "IgtUdEQJyVevaCxQnY"))
	if user.Password != pwd {
		return errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return errors.New("您的账号已被禁用,请联系管理员")
	}
	// 设置SESSION信息
	session.Set("userId", user.Id)
	session.Set("userInfo", user)
	sessionId := session.Id()
	SessionList.Set(sessionId, session)
	return nil
}

// 获取个人信息
func (s *userService) GetProfile(session *ghttp.Session) (u *model.User) {
	_ = session.GetStruct("userInfo", &u)
	return
}
