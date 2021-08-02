/**
 *
 * @author 摆渡人
 * @since 2021/7/29
 * @File : member
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

//
var Member = new(memberCtl)

type memberCtl struct{}

func (c *memberCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "member/index.html",
	})
}

func (c *memberCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.MemberPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list, count, err := service.Member.GetList(req)
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
		Count: count,
	})
}

func (c *memberCtl) Edit(r *ghttp.Request) {
	// 记录ID
	id := r.GetQueryInt("id")

	// 会员等级
	list, _ := dao.MemberLevel.Where("mark=1").All()
	memberLevelList := make(map[int]string, 0)
	for _, v := range list {
		memberLevelList[v.Id] = v.Name
	}

	if id > 0 {
		// 编辑
		info, err := dao.Member.FindOne("id=?", id)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 头像
		if info.Avatar != "" {
			info.Avatar = utils.GetImageUrl(info.Avatar)
		}
		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":         "member/edit.html",
			"info":            info,
			"memberLevelList": memberLevelList,
			"deviceList":      common.MEMBER_DEVICE_LIST,
			"sourceList":      common.MEMBER_SOURCE_LIST,
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":         "member/edit.html",
			"memberLevelList": memberLevelList,
			"deviceList":      common.MEMBER_DEVICE_LIST,
			"sourceList":      common.MEMBER_SOURCE_LIST,
		})
	}
}

func (c *memberCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MemberAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.Member.Add(req, utils.Uid(r.Session))
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

func (c *memberCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MemberUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		rows, err := service.Member.Update(req, utils.Uid(r.Session))
		if err != nil || rows == 0 {
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

func (c *memberCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.MemberDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.Member.Delete(req.Ids)
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

func (c *memberCtl) Status(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.MemberStatusReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		result, err := service.Member.Status(req, utils.Uid(r.Session))
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
