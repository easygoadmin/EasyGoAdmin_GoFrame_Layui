/**
 *
 * @author 摆渡人
 * @since 2021/7/16
 * @File : dept
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Dept = new(deptCtl)

type deptCtl struct{}

func (c *deptCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "dept/index.html",
	})
}

func (c *deptCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.DeptPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用获取部门列表方法
	list := service.Dept.GetList(req)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *deptCtl) Edit(r *ghttp.Request) {
	id := r.GetQueryInt64("id")
	if id > 0 {
		// 编辑
		info, err := dao.Dept.FindOne("id=?", id)
		if err != nil || info == nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":  "dept/edit.html",
			"info":     info,
			"typeList": common.DEPT_TYPE_LIST,
		})
	} else {
		// 添加
		//var list []model.Position
		//dao.Position.Structs(&list)
		//data := make([]map[int]string, 0)
		//if len(list) > 0 {
		//	for _, v := range list {
		//		info := make(map[int]string)
		//		info[v.Id] = v.Name
		//		data = append(data, info)
		//	}
		//}
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":  "dept/edit.html",
			"typeList": common.DEPT_TYPE_LIST,
		})
	}
}

func (c *deptCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.DeptAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.Dept.Add(req)
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

func (c *deptCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.DeptUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用编辑方法
		result, err := service.Dept.Update(req)
		if err != nil || result == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
}

func (c *deptCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.DeptDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法

	}
}
