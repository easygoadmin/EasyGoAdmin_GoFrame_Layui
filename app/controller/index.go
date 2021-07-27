/**
 *
 * @author 摆渡人
 * @since 2021/5/19
 * @File : index
 */
package controller

import (
	"easygoadmin/app/service"
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
		user := service.User.GetProfile(r.Session)
		// 获取菜单列表
		menuList := service.Menu.GetPermissionList(user.Id)
		//fmt.Println(g.Map{
		//	"realname": user.Realname,
		//	"avatar":   function.GetImageUrl(user.Avatar),
		//	"menuList": menuList,
		//})
		// 渲染模板并绑定数据
		response.BuildTpl(r, "index.html").WriteTpl(g.Map{
			"realname": user.Realname,
			"avatar":   function.GetImageUrl(user.Avatar),
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
