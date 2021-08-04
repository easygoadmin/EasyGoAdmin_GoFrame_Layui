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
 * 文件上传-控制器
 * @author 半城风雨
 * @since 2021/7/23
 * @File : upload
 */
package controller

import (
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
)

// 控制器管理对象
var Upload = new(uploadCtl)

type uploadCtl struct{}

func (c *uploadCtl) UploadImage(r *ghttp.Request) {
	// 上传文件
	upFile := r.GetUploadFile("file")
	if upFile == nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  "上传文件不能为空",
		})
	}
	// 调用上传方法
	result, err := service.Upload.UpdImg(upFile)
	if err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 拼接图片地址
	result.FileUrl = utils.GetImageUrl(result.FileUrl)

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "上传成功",
		Data: result,
	})
}
