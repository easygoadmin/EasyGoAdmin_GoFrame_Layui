/**
 *
 * @author 摆渡人
 * @since 2021/7/28
 * @File : config_set
 */
package controller

import (
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var ConfigSet = new(configSet)

type configSet struct{}

func (c *configSet) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
		"mainTpl": "config_set/index.html",
	})
}
