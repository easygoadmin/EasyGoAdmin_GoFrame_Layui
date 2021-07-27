/**
 *
 * @author 摆渡人
 * @since 2021/7/24
 * @File : item_cate
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
var ItemCate = new(itemCateCtl)

type itemCateCtl struct{}

func (c *itemCateCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl": "item_cate/index.html",
	})
}

func (c *itemCateCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.ItemCateQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询列表方法
	list := service.ItemCate.GetList(req)

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: list,
	})
}
func (c *itemCateCtl) Edit(r *ghttp.Request) {
	// 站点列表
	result, _ := dao.Item.All("mark=1")
	var itemList = map[int]string{}
	for _, v := range result {
		itemList[v.Id] = v.Name
	}

	// 记录ID
	id := r.GetQueryInt("id")
	if id > 0 {
		// 编辑
		info, err := dao.ItemCate.FindOne("id=?", id)
		if err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 渲染模板
		response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
			"mainTpl":  "item_cate/edit.html",
			"info":     info,
			"itemList": itemList,
		})
	} else {
		// 添加
		response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
			"mainTpl":  "item_cate/edit.html",
			"itemList": itemList,
		})
	}
}

func (c *itemCateCtl) Add(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.ItemCateAddReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用添加方法
		id, err := service.ItemCate.Add(req)
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

func (c *itemCateCtl) Update(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.ItemCateUpdateReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用更新方法
		rows, err := service.ItemCate.Update(req)
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

func (c *itemCateCtl) Delete(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// 参数验证
		var req *model.ItemCateDeleteReq
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 调用删除方法
		rows, err := service.ItemCate.Delete(req.Ids)
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

func (c *itemCateCtl) GetCateTreeList(r *ghttp.Request) {
	itemId := r.GetQueryInt("item_id")
	list, err := service.ItemCate.GetCateTreeList(itemId, 0)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 数据源转换
	result := service.ItemCate.MakeList(list)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: result,
	})
}
