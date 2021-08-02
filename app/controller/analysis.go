package controller

import (
	"easygoadmin/app/utils/response"
	"github.com/gogf/gf/net/ghttp"
)

var Analysis = new(analysisCtl)

type analysisCtl struct{}

func (c *analysisCtl) Index(r *ghttp.Request) {
	// 渲染模板
	response.BuildTpl(r, "analysis/index.html").WriteTpl()
}
