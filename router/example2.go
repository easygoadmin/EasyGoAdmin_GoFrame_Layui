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
 * 演示二-路由
 * @author 半城风雨
 * @since 2021/08/07
 * @File : example2
 */
func init() {
	s := g.Server()

	/* 案例演示 */
	s.Group("example2", func(group *ghttp.RouterGroup) {
		group.GET("/index", controller.Example2.Index)
		group.POST("/list", controller.Example2.List)
		group.GET("/edit", controller.Example2.Edit)
		group.POST("/add", controller.Example2.Add)
		group.POST("/update", controller.Example2.Update)
		group.POST("/delete", controller.Example2.Delete)
		group.POST("/setStatus", controller.Example2.Status)
	})
}
