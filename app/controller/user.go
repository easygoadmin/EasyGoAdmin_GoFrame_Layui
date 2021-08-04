// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 用户管理-控制器
 * @author 半城风雨
 * @since 2021/7/27
 * @File : user
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
	"github.com/gogf/gf/util/gutil"
)

// 控制器管理对象
var User = new(userCtl)

type userCtl struct{}

func (c *userCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "user/index.html",
	})
}

func (c *userCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.UserPageReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list, count, err := service.User.GetList(req)
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

func (c *userCtl) Edit(r *ghttp.Request) {
	// 获取职级
	levelAll, _ := dao.Level.Where("status=1 and mark=1").All()
	levelList := make(map[int]string, 0)
	for _, v := range levelAll {
		levelList[v.Id] = v.Name
	}
	// 获取岗位
	positionAll, _ := dao.Position.Where("status=1 and mark=1").All()
	positionList := make(map[int]string, 0)
	for _, v := range positionAll {
		positionList[v.Id] = v.Name
	}
	// 获取部门列表
	deptData, _ := service.Dept.GetDeptTreeList()
	deptList := service.Dept.MakeList(deptData)
	// 获取角色
	roleData, _ := dao.Role.Where("status=1 and mark=1").All()
	roleList := make(map[int]string)
	for _, v := range roleData {
		roleList[v.Id] = v.Name
	}

	// 记录ID
	id := r.GetQueryInt("id")
	if id > 0 {
		// 编辑
		info, err := dao.User.FindOne("id=?", id)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		var userInfo = model.UserInfoVo{}
		userInfo.User = *info
		// 头像
		userInfo.Avatar = utils.GetImageUrl(info.Avatar)

		// 角色ID
		var userRoleList []model.UserRole
		dao.UserRole.Where("user_id=?", utils.Uid(r.Session)).Structs(&userRoleList)
		roleIds := gutil.ListItemValuesUnique(&userRoleList, "RoleId")
		userInfo.RoleIds = roleIds

		// 渲染模板
		response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
			"mainTpl":      "user/edit.html",
			"info":         userInfo,
			"genderList":   utils.GENDER_LIST,
			"levelList":    levelList,
			"positionList": positionList,
			"deptList":     deptList,
			"roleList":     roleList,
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
			"mainTpl":      "user/edit.html",
			"genderList":   utils.GENDER_LIST,
			"levelList":    levelList,
			"positionList": positionList,
			"deptList":     deptList,
			"roleList":     roleList,
		})
	}
}

func (c *userCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UserAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.User.Add(req, utils.Uid(r.Session))
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

func (c *userCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UserUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		rows, err := service.User.Update(req, utils.Uid(r.Session))
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

func (c *userCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UserDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.User.Delete(req.Ids)
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

func (c *userCtl) Status(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *model.UserStatusReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		result, err := service.User.Status(req, utils.Uid(r.Session))
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

func (c *userCtl) ResetPwd(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.UserResetPwdReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用重置密码方法
		rows, err := service.User.ResetPwd(req.Id, utils.Uid(r.Session))
		if err != nil || rows == 0 {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "重置密码成功",
		})
	}
}
