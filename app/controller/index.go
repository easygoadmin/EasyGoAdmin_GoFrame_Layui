// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 系统主页
 * @author 半城风雨
 * @since 2021/5/19
 * @File : index
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/function"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Index = new(indexCtl)

type indexCtl struct{}

func (c *indexCtl) Index(r *ghttp.Request) {
	if function.IsLogin(r.Session) {
		// 已登录
		// 获取用户信息
		userInfo := service.Login.GetProfile(r.Session)
		// 获取菜单列表
		menuList := service.Menu.GetPermissionMenuList(userInfo.Id)
		// 渲染模板并绑定数据
		response.BuildTpl(r, "index.html").WriteTpl(g.Map{
			"userInfo": userInfo,
			"menuList": menuList,
		})
	} else {
		// 未登录
		r.Response.RedirectTo("/login")
	}
}

func (c *indexCtl) Main(r *ghttp.Request) {
	response.BuildTpl(r, "main.html").WriteTpl()
}

// 个人中心
func (c *indexCtl) UserInfo(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UserInfoReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 更新信息
		_, err := service.User.UpdateUserInfo(req, r.Session)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
	// 获取用户信息
	userInfo := service.Login.GetProfile(r.Session)
	// 渲染模板
	response.BuildTpl(r, "user_info.html").WriteTpl(g.Map{
		"userInfo": userInfo,
	})
}

// 更新密码
func (c *indexCtl) UpdatePwd(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UpdatePwd
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新密码方法
		rows, err := service.User.UpdatePwd(req, utils.Uid(r.Session))
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新密码成功",
		})
	}
}

// 退出登录
func (c *indexCtl) Logout(r *ghttp.Request) {
	// 设置SESSION为空
	r.Session.Set("userId", "")
	r.Session.Set("userInfo", "")
	// 重定向至登录页面
	r.Response.RedirectTo("/login")
}
