// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

package router

import (
	"easygoadmin/app/controller"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/**
 * 演示一-路由
 * @author 半城风雨
 * @since 2021/08/07
 * @File : example
 */
func init() {
	s := g.Server()

	/* 案例演示 */
	s.Group("example", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Example.Index)
		group.POST("/list", controller.Example.List)
		group.GET("/edit", controller.Example.Edit)
		group.POST("/add", controller.Example.Add)
		group.POST("/update", controller.Example.Update)
		group.POST("/delete", controller.Example.Delete)
		group.POST("/setStatus", controller.Example.Status)
		group.POST("/setIsVip", controller.Example.IsVip)
	})
}
