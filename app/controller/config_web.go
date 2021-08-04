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
 * 网站配置-控制器
 * @author 半城风雨
 * @since 2021/7/28
 * @File : config_set
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/response"
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"regexp"
	"strings"
)

// 控制器管理对象
var ConfigWeb = new(configWeb)

type configWeb struct{}

func (c *configWeb) Index(r *ghttp.Request) {
	if r.IsAjaxRequest() {
		// key：string类型，value：interface{}  类型能存任何数据类型
		var jsonObj map[string]interface{}
		json.Unmarshal(r.GetBody(), &jsonObj)
		// 遍历处理数据源
		for key, val := range jsonObj {
			// 参数处理
			if key == "checkbox" {
				// 复选框

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
			} else if strings.Contains(key, "upimage") {
				// 单图上传

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				if strings.Contains(gconv.String(val), "temp") {
					image, _ := utils.SaveImage(gconv.String(val), "config")
					// 赋值给参数
					val = image
				} else {
					// 赋值给参数
					val = gstr.Replace(gconv.String(val), utils.ImgUrl(), "")
				}
			} else if strings.Contains(key, "upimgs") {
				// 多图上传
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				// 图片地址处理
				urlArr := gstr.Split(gconv.String(val), ",")
				list := make([]string, 0)
				for _, v := range urlArr {
					if strings.Contains(gconv.String(v), "temp") {
						image, _ := utils.SaveImage(v, "config")
						list = append(list, image)
					} else {
						image := gstr.Replace(v, utils.ImgUrl(), "")
						list = append(list, image)
					}
				}
				// 数组转字符串，逗号分隔
				val = strings.Join(list, ",")
			} else if strings.Contains(key, "ueditor") {
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]

				// 富文本处理(待完善)
				// TODO...
			}

			// 查询记录
			info, err := dao.ConfigData.FindOne("code=?", key)
			if err != nil || info == nil {
				continue
			}

			// 更新记录
			result, err := dao.ConfigData.Data(g.Map{
				"value": val,
			}).Where("code=?", key).Update()
			if err != nil {
				continue
			}

			// 获取受影响行数
			rows, _ := result.RowsAffected()
			if rows == 0 {
				continue
			}
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "保存成功",
		})
	}
	// 配置ID
	configId := r.GetQueryInt("configId", 1)

	// 获取配置列表
	configData, _ := dao.Config.Where("mark=1").All()
	configList := make(map[string]string)
	for _, v := range configData {
		configList[gconv.String(v.Id)] = v.Name
	}

	// 获取配置项列表
	itemData, _ := dao.ConfigData.Where("config_id=? and status=1 and mark=1", configId).Order("sort asc").All()
	itemList := make([]map[string]interface{}, 0)
	for _, v := range itemData {
		item := make(map[string]interface{})
		item["id"] = v.Id
		item["title"] = v.Title
		item["code"] = v.Code
		item["value"] = v.Value
		item["type"] = v.Type

		if v.Type == "checkbox" {
			// 复选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, val := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "radio" {
			// 单选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList

		} else if v.Type == "select" {
			// 下拉选择框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "image" {
			// 单图片
			item["value"] = utils.GetImageUrl(v.Value)
		} else if v.Type == "images" {
			// 多图片
			list := gstr.Split(v.Value, ",")
			itemList := make([]string, 0)
			for _, v := range list {
				// 图片地址
				item := utils.GetImageUrl(v)
				itemList = append(itemList, item)
			}
			item["value"] = itemList
		}
		itemList = append(itemList, item)
	}

	// 渲染模板
	response.BuildTpl(r, "public/layout.html").WriteTpl(g.Map{
		"mainTpl":    "config_web/index.html",
		"configId":   configId,
		"configList": configList,
		"itemList":   itemList,
	})
}
