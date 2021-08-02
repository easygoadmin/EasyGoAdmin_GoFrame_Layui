/**
 *
 * @author 摆渡人
 * @since 2021/5/20
 * @File : level
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
var Level = new(levelCtl)

type levelCtl struct{}

func (c *levelCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "level/index.html",
	})
}

func (c *levelCtl) List(r *ghttp.Request) {
	// 请求参数
	var req *model.LevelQueryReq
	// 请求验证
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表函数
	list, count, err := service.Level.GetList(req)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果集
	r.Response.WriteJsonExit(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "操作成功",
		Count: count,
	})
}

func (c *levelCtl) Edit(r *ghttp.Request) {
	// 查询记录
	id := r.GetQueryInt64("id")
	if id > 0 {
		info, err := dao.Level.FindOne("id=?", id)
		if err != nil || info == nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl": "level/edit.html",
			"info":    info,
		})
	} else {
		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl": "level/edit.html",
		})
	}

}

func (c *levelCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.LevelAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		id, err := service.Level.Add(req, utils.Uid(r.Session))
		if err != nil || id <= 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 保存成功
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "保存成功",
		})
	}
}

func (c *levelCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.LevelEditReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 调用更新方法
		result, err := service.Level.Update(req, utils.Uid(r.Session))
		if err != nil || result == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 更新成功提示
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
}

func (c *levelCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.LevelDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		result := service.Level.Delete(req.Ids)
		if result > 0 {
			// 删除成功
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 0,
				Msg:  "删除成功",
			})
		} else {
			// 删除失败
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  "删除失败",
			})
		}
	}
}

func (c *levelCtl) Status(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.LevelStatusReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		result, err := service.Level.Status(req, utils.Uid(r.Session))
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
