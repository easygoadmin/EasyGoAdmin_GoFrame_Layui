// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 登录-服务类
 * @author 半城风雨
 * @since 2021/3/3
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// 中间件管理服务
var Login = new(loginService)

type loginService struct{}

var SessionList = gmap.New(true)

// 系统登录
func (s *loginService) UserLogin(username, password string, session *ghttp.Session) error {
	// 获取用户信息
	user, err := dao.User.FindOne("username=? and mark=1", username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username) //gmd5.MustEncryptString(gmd5.MustEncryptString(password + "IgtUdEQJyVevaCxQnY"))
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
func (s *loginService) GetProfile(session *ghttp.Session) (u *model.User) {
	_ = session.GetStruct("userInfo", &u)
	// 头像
	if u.Avatar != "" && !gstr.Contains(u.Avatar, utils.ImgUrl()) {
		u.Avatar = utils.GetImageUrl(u.Avatar)
	}
	return
}
