/**
 *
 * @author 摆渡人
 * @since 2021/7/19
 * @File : menu
 */
package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
)

var Menu = new(menuCtl)

type menuCtl struct{}

func (c *menuCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "menu/index.html",
	})
}

func (c *menuCtl) List(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MenuQueryReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用查询方法
		list := service.Menu.List(req)
		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Data: list,
			Msg:  "操作成功",
		})
	}
}

func (c *menuCtl) Edit(r *ghttp.Request) {
	id := r.GetQueryUint64("id")
	if id > 0 {
		// 编辑
		info, err := dao.Menu.FindOne("id=?", id)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":  "menu/edit.html",
			"info":     info,
			"typeList": common.MENU_TYPE_LIST,
			"funcList": []int{1, 5},
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl": "menu/edit.html",
		})
	}
}

func (c *menuCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MenuAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.Menu.Add(req)
		if err != nil || id == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回成功提示
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "添加成功",
		})
	}
}

func (c *menuCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MenuUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		rows, err := service.Menu.Update(req)
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回成功提示
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})

	}
}

func (c *menuCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MenuDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.Menu.Delete(req.Ids)
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
