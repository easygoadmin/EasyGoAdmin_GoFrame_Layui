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
