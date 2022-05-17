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
 * 系统路由
 * @author 半城风雨
 * @since 2021/7/26
 * @File : submit
 */
package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	// 登录验证中间件
	s.Use(middleware.CheckLogin)

	/* 文件上传 */
	s.Group("/upload", func(group *ghttp.RouterGroup) {
		// 上传图片
		group.POST("/uploadImage", controller.Upload.UploadImage)
	})

	/* 登录注册 */
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", controller.Index.Index)
		group.ALL("/login", controller.Login.Login)
		group.GET("/captcha", controller.Login.Captcha)
		group.GET("/index", controller.Index.Index)
		group.GET("/main", controller.Index.Main)
		group.ALL("/userInfo", controller.Index.UserInfo)
		group.ALL("/updatePwd", controller.Index.UpdatePwd)
		group.GET("/logout", controller.Index.Logout)
	})

	/* 用户管理 */
	s.Group("user", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.User.Index)
		group.POST("/list", controller.User.List)
		group.GET("/edit", controller.User.Edit)
		group.POST("/add", controller.User.Add)
		group.POST("/update", controller.User.Update)
		group.POST("/delete", controller.User.Delete)
		group.POST("/setStatus", controller.User.Status)
		group.POST("/resetPwd", controller.User.ResetPwd)
	})

	/* 职级管理 */
	s.Group("level", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Level.Index)
		group.POST("/list", controller.Level.List)
		group.GET("/edit", controller.Level.Edit)
		group.POST("/add", controller.Level.Add)
		group.POST("/update", controller.Level.Update)
		group.POST("/delete", controller.Level.Delete)
		group.POST("/setStatus", controller.Level.Status)
	})

	/* 岗位路由 */
	s.Group("position", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Position.Index)
		group.POST("/list", controller.Position.List)
		group.GET("/edit", controller.Position.Edit)
		group.POST("/add", controller.Position.Add)
		group.POST("/update", controller.Position.Update)
		group.POST("/delete", controller.Position.Delete)
	})

	/* 角色路由 */
	s.Group("role", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Role.Index)
		group.POST("/list", controller.Role.List)
		group.GET("/edit", controller.Role.Edit)
		group.POST("/add", controller.Role.Add)
		group.POST("/update", controller.Role.Update)
		group.POST("/delete", controller.Role.Delete)
		group.POST("/setStatus", controller.Role.Status)
	})

	/* 角色菜单权限 */
	s.Group("rolemenu", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.RoleMenu.Index)
		group.POST("/save", controller.RoleMenu.Save)
	})

	/* 部门管理 */
	s.Group("dept", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Dept.Index)
		group.POST("/list", controller.Dept.List)
		group.GET("/edit", controller.Dept.Edit)
		group.POST("/add", controller.Dept.Add)
		group.POST("/update", controller.Dept.Update)
		group.POST("/delete", controller.Dept.Delete)
	})

	/* 菜单管理 */
	s.Group("menu", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Menu.Index)
		group.POST("/list", controller.Menu.List)
		group.GET("/edit", controller.Menu.Edit)
		group.POST("/add", controller.Menu.Add)
		group.POST("/update", controller.Menu.Update)
		group.POST("/delete", controller.Menu.Delete)
	})

	/* 城市管理 */
	s.Group("city", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.City.Index)
		group.POST("/list", controller.City.List)
		group.GET("/edit", controller.City.Edit)
		group.POST("/add", controller.City.Add)
		group.POST("/update", controller.City.Update)
		group.POST("/delete", controller.City.Delete)
		group.POST("/getChilds", controller.City.GetChilds)
	})

	/* 字典管理 */
	s.Group("dict", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Dict.Index)
		group.POST("/list", controller.Dict.List)
		group.POST("/add", controller.Dict.Add)
		group.POST("/update", controller.Dict.Update)
		group.POST("/delete", controller.Dict.Delete)
	})

	/* 字典项管理 */
	s.Group("dictdata", func(group *ghttp.RouterGroup) {
		group.POST("/list", controller.DictData.List)
		group.POST("/add", controller.DictData.Add)
		group.POST("/update", controller.DictData.Update)
		group.POST("/delete", controller.DictData.Delete)
	})

	/* 配置管理 */
	s.Group("config", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Config.Index)
		group.POST("/list", controller.Config.List)
		group.POST("/add", controller.Config.Add)
		group.POST("/update", controller.Config.Update)
		group.POST("/delete", controller.Config.Delete)
	})

	/* 字典项管理 */
	s.Group("configdata", func(group *ghttp.RouterGroup) {
		group.POST("/list", controller.ConfigData.List)
		group.GET("/edit", controller.ConfigData.Edit)
		group.POST("/add", controller.ConfigData.Add)
		group.POST("/update", controller.ConfigData.Update)
		group.POST("/delete", controller.ConfigData.Delete)
		group.POST("/setStatus", controller.ConfigData.Status)
	})

	/* 友链管理 */
	s.Group("link", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Link.Index)
		group.POST("/list", controller.Link.List)
		group.GET("/edit", controller.Link.Edit)
		group.POST("/add", controller.Link.Add)
		group.POST("/update", controller.Link.Update)
		group.POST("/delete", controller.Link.Delete)
		group.POST("/setStatus", controller.Link.Status)
	})

	/* 站点管理 */
	s.Group("item", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Item.Index)
		group.POST("/list", controller.Item.List)
		group.GET("/edit", controller.Item.Edit)
		group.POST("/add", controller.Item.Add)
		group.POST("/update", controller.Item.Update)
		group.POST("/delete", controller.Item.Delete)
	})

	/* 栏目管理 */
	s.Group("itemcate", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.ItemCate.Index)
		group.POST("/list", controller.ItemCate.List)
		group.GET("/edit", controller.ItemCate.Edit)
		group.POST("/add", controller.ItemCate.Add)
		group.POST("/update", controller.ItemCate.Update)
		group.POST("/delete", controller.ItemCate.Delete)
		group.GET("/getCateTreeList", controller.ItemCate.GetCateTreeList)
	})

	/* 广告位管理 */
	s.Group("adsort", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.AdSort.Index)
		group.POST("/list", controller.AdSort.List)
		group.GET("/edit", controller.AdSort.Edit)
		group.POST("/add", controller.AdSort.Add)
		group.POST("/update", controller.AdSort.Update)
		group.POST("/delete", controller.AdSort.Delete)
	})

	/* 广告管理 */
	s.Group("ad", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Ad.Index)
		group.POST("/list", controller.Ad.List)
		group.GET("/edit", controller.Ad.Edit)
		group.POST("/add", controller.Ad.Add)
		group.POST("/update", controller.Ad.Update)
		group.POST("/delete", controller.Ad.Delete)
		group.POST("/setStatus", controller.Ad.Status)
	})

	/* 通知管理 */
	s.Group("notice", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Notice.Index)
		group.POST("/list", controller.Notice.List)
		group.GET("/edit", controller.Notice.Edit)
		group.POST("/add", controller.Notice.Add)
		group.POST("/update", controller.Notice.Update)
		group.POST("/delete", controller.Notice.Delete)
	})

	/* 网站设置 */
	s.Group("configweb", func(group *ghttp.RouterGroup) {
		group.ALL("/index", controller.ConfigWeb.Index)
	})

	/* 会员等级 */
	s.Group("memberlevel", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.MemberLevel.Index)
		group.POST("/list", controller.MemberLevel.List)
		group.GET("/edit", controller.MemberLevel.Edit)
		group.POST("/add", controller.MemberLevel.Add)
		group.POST("/update", controller.MemberLevel.Update)
		group.POST("/delete", controller.MemberLevel.Delete)
	})

	/* 会员管理 */
	s.Group("member", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Member.Index)
		group.POST("/list", controller.Member.List)
		group.GET("/edit", controller.Member.Edit)
		group.POST("/add", controller.Member.Add)
		group.POST("/update", controller.Member.Update)
		group.POST("/delete", controller.Member.Delete)
		group.POST("/setStatus", controller.Member.Status)
	})

	/* 统计分析 */
	s.Group("analysis", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Analysis.Index)
	})

	/* 代码生成器 */
	s.Group("generate", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Generate.Index)
		group.POST("/list", controller.Generate.List)
		group.POST("/generate", controller.Generate.Generate)
	})

}
