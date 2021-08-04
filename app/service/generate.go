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
 * 代码生成器-服务类
 * @author 半城风雨
 * @since 2021/8/2
 * @File : generate
 */
package service

import (
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"os"
	"regexp"
	"strings"
)

// 中间件管理服务
var Generate = new(generateService)

type generateService struct{}

func (s *generateService) GetList(req *model.GeneratePageReq) ([]model.GenerateInfo, error) {
	// 查询SQL
	sql := "SHOW TABLE STATUS"
	// 查询条件
	if req != nil {
		// 表名称
		if req.Name != "" {
			sql += " WHERE Name like \"%" + req.Name + "%\""
		}
		// 表描述
		if req.Comment != "" {
			sql += " WHERE Comment like \"%" + req.Comment + "%\""
		}
	}
	// 对象转换
	var list []model.GenerateInfo
	err := g.DB().GetScan(&list, sql)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return list, nil
}

func (s *generateService) Generate(r *ghttp.Request) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 参数验证
	var req *model.GenerateFileReq
	if err := r.Parse(&req); err != nil {
		return 0, err
	}
	// 数据表名
	tableName := req.Name
	// 数据表描述
	moduleTitle := req.Comment
	// 替换“表”
	if gstr.Contains(moduleTitle, "表") {
		moduleTitle = gstr.Replace(moduleTitle, "表", "")
	}
	// 替换“管理”
	if gstr.Contains(moduleTitle, "管理") {
		moduleTitle = gstr.Replace(moduleTitle, "管理", "")
	}
	// 模型名称
	moduleName := gstr.Replace(tableName, "sys_", "")
	//// 控制器名
	//controllerName := tableName
	// 作者名称
	authorName := "半城风雨"

	// 获取字段列表
	columnList, err := GetColumnList(tableName)
	if err != nil {
		return 0, err
	}

	// 生成控制器
	if err := GenerateController(r, tableName, authorName, moduleName, moduleTitle); err != nil {
		return 0, err
	}

	// 生成服务类
	if err := GenerateService(r, columnList, authorName, moduleName, moduleTitle); err != nil {
		return 0, err
	}

	// 生成模块index.html
	if err := GenerateIndex(r, columnList, moduleName, moduleTitle); err != nil {
		return 0, err
	}

	return 1, nil
}

// 生成控制器
func GenerateController(r *ghttp.Request, tableName string, authorName string, moduleName string, moduleTitle string) error {
	// 获取字段列表
	columnList, err := GetColumnList(tableName)
	if err != nil {
		return err
	}

	// 读取HTML模板并绑定数据
	if filePath, err := r.Response.ParseTpl("tpl/controller.html", g.Map{
		"author":      authorName,
		"since":       gtime.Now().Format("Y/m/d"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName := strings.Join([]string{curDir, "/app/controller/", moduleName, ".go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(filePath)
			}
			f.Close()
		}
	}
	return nil
}

// 生成服务类
func GenerateService(r *ghttp.Request, columnList *garray.Array, authorName string, moduleName string, moduleTitle string) error {
	// 读取HTML模板并绑定数据
	if filePath, err := r.Response.ParseTpl("tpl/service.html", g.Map{
		"author":      authorName,
		"since":       gtime.Now().Format("Y/m/d"),
		"moduleName":  moduleName,
		"entityName":  gstr.UcWords(moduleName),
		"moduleTitle": moduleTitle,
		"columnList":  columnList,
	}); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName := strings.Join([]string{curDir, "/app/service/", moduleName, ".go"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(filePath)
			}
			f.Close()
		}
	}
	return nil
}

func GenerateIndex(r *ghttp.Request, columnList *garray.Array, moduleName string, moduleTitle string) error {
	// 读取HTML模板并绑定数据
	if filePath, err := r.Response.ParseTpl("tpl/index.html", g.Map{
		"queryList": columnList,

		"funcList1": `{{query "查询"}}
                {{add "添加` + moduleTitle + `" "{}"}}
                {{dall "批量删除"}}`,
		"funcList2": `{{edit "编辑"}}
    {{delete "删除"}}`,
	}); err == nil {
		// 获取项目根目录
		curDir, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName := strings.Join([]string{curDir, "/template/", moduleName, "/index.html"}, "")
		if !gfile.Exists(fileName) {
			f, err := gfile.Create(fileName)
			if err == nil {
				f.WriteString(filePath)
			}
			f.Close()
		}
	}
	return nil
}

// 获取表字段列表
func GetColumnList(tableName string) (*garray.Array, error) {
	// 获取数据库名
	DbName, err := utils.GetDatabase()
	if err != nil {
		return nil, err
	}

	// 获取字段列表
	data, err := g.DB().GetAll("SELECT COLUMN_NAME,COLUMN_DEFAULT,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT FROM information_schema.`COLUMNS` where TABLE_SCHEMA = ? AND TABLE_NAME = ?", DbName, tableName)
	if err != nil {
		return nil, err
	}

	// 初始化数组
	result := garray.NewArray(true)
	for _, v := range data {
		// 初始化Map
		item := make(map[string]interface{})
		// 字段列名
		columnName := v["COLUMN_NAME"].String()
		// 系统常规字段直接跳过
		if columnName == "create_user" ||
			columnName == "create_time" ||
			columnName == "update_user" ||
			columnName == "update_time" ||
			columnName == "mark" {
			continue
		}
		item["columnName"] = columnName
		// 字段默认值
		item["columnDefault"] = v["COLUMN_DEFAULT"].String()
		// 字段数据类型
		item["dataType"] = v["DATA_TYPE"].String()
		// 字段描述
		columnComment := v["COLUMN_COMMENT"].String()
		item["columnComment"] = columnComment
		// 判断是否有规则描述
		if gstr.Contains(columnComment, ":") || gstr.Contains(columnComment, "：") {
			// 正则根据冒号分裂字符串
			re := regexp.MustCompile("[：；]")
			commentItem := gstr.Split(re.ReplaceAllString(columnComment, "|"), "|")
			// 字段描述
			item["columnTitle"] = commentItem[0]

			// 字段描述数据处理
			commentStr := gstr.Replace(commentItem[1], " ", "|")
			commentArr := gstr.Split(commentStr, "|")

			// 实例化字段描述参数数组
			columnValue := make([]string, 0)
			// 实例化字段描述文字数组
			columnSwitchValue := make([]string, 0)
			for _, v := range commentArr {
				fmt.Println(v)
				// 正则提取数字键
				regexp := regexp.MustCompile(`[0-9]+`)
				vItem := regexp.FindStringSubmatch(v)
				// 键
				key := vItem[0]
				// 值
				value := gstr.Replace(v, vItem[0], "")
				// 加入数组
				columnValue = append(columnValue, key+"="+value)
				columnSwitchValue = append(columnSwitchValue, value)
			}
			// 字符串逗号拼接
			item["columnValue"] = gstr.Join(columnValue, ",")
			// 开关参数处理
			if columnName == "status" || gstr.SubStr(columnName, 0, 3) == "is_" {
				item["columnSwitch"] = true
				item["columnSwitchValue"] = gstr.Join(columnSwitchValue, "|")
				// 方法名处理
				columnSwitchName := ""
				if gstr.Contains(columnName, "_") {
					switchArr := gstr.Split(columnName, "_")
					columnSwitchName = "set" + gstr.UcWords(switchArr[0]) + gstr.UcWords(switchArr[1])
				} else {
					columnSwitchName = "set" + gstr.UcWords(columnName)
				}
				item["columnSwitchName"] = columnSwitchName
			} else {
				item["columnSwitch"] = false
			}
		} else {
			// 字段描述
			item["columnTitle"] = columnComment
		}

		// 判断是否是图片
		if gstr.Contains(columnName, "cover") ||
			gstr.Contains(columnName, "avatar") ||
			gstr.Contains(columnName, "image") ||
			gstr.Contains(columnName, "logo") ||
			gstr.Contains(columnName, "pic") {
			item["columnImage"] = true
		} else {
			item["columnImage"] = false
		}

		// 判断是否多行文本或富文本
		if gstr.Contains(columnName, "note") ||
			gstr.Contains(columnName, "remark") ||
			gstr.Contains(columnName, "content") ||
			gstr.Contains(columnName, "description") ||
			gstr.Contains(columnName, "intro") {
			item["columnTextArea"] = true
		}

		// 加入数组
		result.Append(item)
	}
	return result, nil
}
