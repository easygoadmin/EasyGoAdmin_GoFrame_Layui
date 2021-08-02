/**
 *
 * @author 摆渡人
 * @since 2021/7/15
 * @File : role
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
var Role = new(roleCtl)

type roleCtl struct{}

func (c *roleCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "role/index.html",
	})
}

func (c *roleCtl) List(r *ghttp.Request) {
	var req *model.RolePageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表方法
	list, count, err := service.Role.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回查询结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *roleCtl) Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")
	if id > 0 {
		// 编辑
		// 编辑
		info, err := dao.Role.FindOne("id=?", id)
		if err != nil || info == nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl": "role/edit.html",
			"info":    info,
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl": "role/edit.html",
		})
	}
}

func (c *roleCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.RoleAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 调用添加方法
		id, err := service.Role.Add(req, utils.Uid(r.Session))
		if err != nil || id <= 0 {
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

func (c *roleCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.RoleUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		result, err := service.Role.Update(req, utils.Uid(r.Session))
		if err != nil || result == 0 {
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

func (c *roleCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.RoleDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		result, err := service.Role.Delete(req.Ids)
		if err != nil || result == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "删除成功",
		})
	}
}

func (c *roleCtl) Status(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.RoleStatusReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用设置方法
		result, err := service.Role.Staus(req, utils.Uid(r.Session))
		if err != nil || result == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "设置成功",
		})
	}
}
