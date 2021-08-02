/**
 *
 * @author 摆渡人
 * @since 2021/7/22
 * @File : link
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Link = new(linkCtl)

type linkCtl struct{}

func (c *linkCtl) Index(r *ghttp.Request) {
	// 模板渲染
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl":      "link/index.html",
		"typeList":     common.LINK_TYPE_LIST,
		"platformList": common.LINK_PLATFORM_LIST,
	})
}

func (c *linkCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.LinkPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list, count, err := service.Link.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *linkCtl) Edit(r *ghttp.Request) {
	id := r.GetQueryInt("id")
	if id > 0 {
		// 修改
		info, err := dao.Link.FindOne("id=?", id)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 设置图片链接
		info.Image = utils.GetImageUrl(info.Image)

		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":      "link/edit.html",
			"info":         info,
			"typeList":     common.LINK_TYPE_LIST,
			"formList":     common.LINK_FORM_LIST,
			"platformList": common.LINK_PLATFORM_LIST,
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":      "link/edit.html",
			"typeList":     common.LINK_TYPE_LIST,
			"formList":     common.LINK_FORM_LIST,
			"platformList": common.LINK_PLATFORM_LIST,
		})
	}
}

func (c *linkCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.LinkAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.Link.Add(req, utils.Uid(r.Session))
		if err != nil || id == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "添加成功",
		})
	}
}

func (c *linkCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.LinkUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		rows, err := service.Link.Update(req, utils.Uid(r.Session))
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "修改成功",
		})
	}
}

func (c *linkCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.LinkDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.Link.Delete(req.Ids)
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

func (c *linkCtl) Status(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.LinkStatusReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		result, err := service.Link.Status(req, utils.Uid(r.Session))
		if err != nil || result == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 保存成功
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "设置成功",
		})
	}
}
