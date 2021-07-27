/**
 *
 * @author 摆渡人
 * @since 2021/7/21
 * @File : dict
 */
package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
)

// 控制器管理对象
var Dict = new(dictCtl)

type dictCtl struct{}

func (c *dictCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "dict/index.html").WriteTpl()
}

func (c *dictCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.DictQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list := service.Dict.GetList(req)

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *dictCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.DictAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.Dict.Add(req)
		if err != nil || id == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "添加成功",
		})
	}
}

func (c *dictCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.DictUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用修改方法
		rows, err := service.Dict.Update(req)
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
}

func (c *dictCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.DictDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.Dict.Delete(req.Ids)
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "删除成功",
		})
	}
}
