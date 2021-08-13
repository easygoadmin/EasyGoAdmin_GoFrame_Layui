/**
 *
 * @author 摆渡人
 * @since 2021/5/20
 * @File : api_response
 */
package response

import (
	"github.com/gogf/gf/net/ghttp"
	"easygoadmin/app/utils/common"
)

// 通用API响应
type ApiResp struct {
	c *common.JsonResult
	r *ghttp.Request
}

// 返回一个成功的消息体
func SucessResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  0,
		Btype: common.BOther,
		Msg:   "操作成功",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 返回一个错误的消息体
func ErrorResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  500,
		Btype: common.BOther,
		Msg:   "操作失败",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 返回一个拒绝访问的消息体
func ForbiddenResp(r *ghttp.Request) *ApiResp {
	msg := common.JsonResult{
		Code:  403,
		Btype: common.BOther,
		Msg:   "无操作权限",
	}
	var a = ApiResp{
		c: &msg,
		r: r,
	}
	return &a
}

// 设置消息体的内容
func (resp *ApiResp) SetMsg(msg string) *ApiResp {
	resp.c.Msg = msg
	return resp
}

// 设置消息体的编码
func (resp *ApiResp) SetCode(code int) *ApiResp {
	resp.c.Code = code
	return resp
}

// 设置消息体的数据
func (resp *ApiResp) SetData(data interface{}) *ApiResp {
	resp.c.Data = data
	return resp
}

// 设置消息体的业务类型
func (resp *ApiResp) SetBtype(btype common.BunissType) *ApiResp {
	resp.c.Btype = btype
	return resp
}

// 输出json到客户端
func (resp *ApiResp) WriteJsonExit() {
	resp.r.Response.WriteJsonExit(resp.c)
}
