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
		if utils.AppDebug() {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  "演示环境，暂无权限操作",
			})
		}
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
