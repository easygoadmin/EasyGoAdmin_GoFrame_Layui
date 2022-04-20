// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 菜单管理-控制器
 * @author 半城风雨
 * @since 2021/7/19
 * @File : menu
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
	// 获取菜单列表
	menuTreeList, _ := service.Menu.GetTreeList()
	// 数据源转换
	menuList := service.Menu.MakeList(menuTreeList)

	// 记录ID
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

		// 获取节点
		funcList, _ := dao.Menu.Where("pid=? and type=1 and mark=1", id).All()
		sortList := gutil.ListItemValuesUnique(&funcList, "Sort")

		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":  "menu/edit.html",
			"info":     info,
			"typeList": common.MENU_TYPE_LIST,
			"funcList": sortList,
			"menuList": menuList,
		})
	} else {
		// 添加

		pid := r.GetInt("pid")
		var info model.Menu
		info.Pid = pid
		info.Status = 1
		info.Target = 1

		// 渲染模板
		response.BuildTpl(r, "public/form.html").WriteTpl(g.Map{
			"mainTpl":  "menu/edit.html",
			"info":     info,
			"typeList": common.MENU_TYPE_LIST,
			"funcList": make([]interface{}, 0),
			"menuList": menuList,
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
		id, err := service.Menu.Add(req, utils.Uid(r.Session))
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
		rows, err := service.Menu.Update(req, utils.Uid(r.Session))
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
