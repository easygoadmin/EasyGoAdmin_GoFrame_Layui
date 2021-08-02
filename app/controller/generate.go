/**
 *
 * @author 摆渡人
 * @since 2021/8/2
 * @File : generate
 */
package controller

import (
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Generate = new(generateCtl)

type generateCtl struct{}

func (c *generateCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "generate/index.html",
	})
}

func (c *generateCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.GeneratePageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list, err := service.Generate.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  list,
		Count: len(list),
	})
}

func (c *generateCtl) Generate(r *ghttp.Request) {
	r.Response.WriteJsonExit(common.JsonResult{
		Code: -1,
		Msg:  "功能开发中,请耐心等待...",
	})
}
