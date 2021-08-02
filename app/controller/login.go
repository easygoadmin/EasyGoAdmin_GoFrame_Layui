/**
 *
 * @author 摆渡人
 * @since 2021/5/18
 * @File : login
 */
package controller

import (
	"easygoadmin/app/service"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/mojocn/base64Captcha"
)

// 用户控制器管理对象
var Login = new(loginCtl)

type loginCtl struct{}

type LoginReq struct {
	UserName string `p:"username" v:required|length:5,30#请输入登录账号|账号长度为：min-max位`
	Password string `p:"password" v:"required|length:6,12#请输入密码|密码长度为：min-max位"`
	Captcha  string `p:"captcha" v:"required|length:4,6#请输入验证码|验证码长度不够"`
	IdKey    string `p:"IdKey" v:"required#验证码KEY不能为空"`
}

// 系统登录
func (c *loginCtl) Login(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		var req *LoginReq

		// 获取参数并验证
		if err := r.Parse(&req); err != nil {
			// 返回错误信息
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		//
		//// 校验验证码
		//verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		//if !verifyRes {
		//	r.Response.WriteJsonExit(common.JsonResult{
		//		Code: -1,
		//		Msg:  "验证码不正确",
		//	})
		//}

		// 系统登录
		if err := service.Login.UserLogin(req.UserName, req.Password, r.Session); err != nil {
			// 登录错误
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		} else {
			// 登录成功
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 0,
				Msg:  "登录成功",
			})
		}
	}
	// 渲染模板
	response.BuildTpl(r, "login.html").WriteTpl()
}

// 验证码
func (c *loginCtl) Captcha(r *ghttp.Request) {
	// 验证码参数配置：字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	///create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	// 返回结果集
	r.Response.WriteJsonExit(common.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}
