/**
 *
 * @author 摆渡人
 * @since 2021/8/2
 * @File : generate
 */
package model

import "github.com/gogf/gf/os/gtime"

// 数据库信息
type GenerateInfo struct {
	Name          string      `json:"name"`           // 表名
	Engine        string      `json:"engine"`         // 引擎
	Version       string      `json:"version"`        // 版本
	Collation     string      `json:"collation"`      // 编码
	Rows          int         `json:"rows"`           // 记录数
	DataLength    int         `json:"data_length"`    // 大小
	AutoIncrement int         `json:"auto_increment"` // 自增索引
	Comment       string      `json:"comment"`        // 表备注
	CreateTime    *gtime.Time `json:"createTime"`     // 添加时间
	UpdateTime    *gtime.Time `json:"updateTime"`     // 更新时间
}

// 分页查询条件
type GeneratePageReq struct {
	Name    string `p:"name"`    // 表名称
	Comment string `p:"comment"` // 表描述
	Page    int    `p:page`      // 页码
	Limit   int    `p:limit`     // 每页数
}

// 生成文件
type GenerateFileReq struct {
	Name    string `p:"name"`    // 表名称
	Comment string `p:"comment"` // 表描述
}

//// 数据表字段信息
//type GenerateTableColumnInfo struct {
//	ColumnName    string      `json:"columnName"`     // 列名字
//	Engine        string      `json:"engine"`         // 引擎
//	Version       string      `json:"version"`        // 版本
//	Collation     string      `json:"collation"`      // 编码
//	Rows          int         `json:"rows"`           // 记录数
//	DataLength    int         `json:"data_length"`    // 大小
//	AutoIncrement int         `json:"auto_increment"` // 自增索引
//	Comment       string      `json:"comment"`        // 表备注
//	CreateTime    *gtime.Time `json:"createTime"`     // 添加时间
//	UpdateTime    *gtime.Time `json:"updateTime"`     // 更新时间
//}
