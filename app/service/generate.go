/**
 *
 * @author 摆渡人
 * @since 2021/8/2
 * @File : generate
 */
package service

import (
	"easygoadmin/app/model"
	"github.com/gogf/gf/frame/g"
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
