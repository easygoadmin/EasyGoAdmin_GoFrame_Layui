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
 * 演示二管理-服务类
 * @author 半城风雨
 * @since 2021/08/07
 * @File : example2
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var Example2 = new(example2Service)

type example2Service struct{}

func (s *example2Service) GetList(req *model.Example2QueryReq) ([]model.Example2InfoVo, int, error) {
	// 创建查询对象
	query := dao.Example2.Where("mark=1")
	// 查询条件
	if req != nil {

		// 演示名称

		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}

		// 状态：1正常 2停用

		if req.Status > 0 {
			query = query.Where("status = ?", req.Status)
		}

	}
	// 获取记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("id asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.Example2
	query.Structs(&list)

	// 数据处理
	var result []model.Example2InfoVo
	for _, v := range list {
		item := model.Example2InfoVo{}
		item.Example2 = v

		result = append(result, item)
	}

	// 返回结果
	return result, count, nil
}

func (s *example2Service) Add(req *model.Example2AddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}

	// 实例化模型
	var entity model.Example2

	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.Example2.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

func (s *example2Service) Update(req *model.Example2UpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 获取记录信息
	info, err := dao.Example2.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 对象赋值

	info.Name = req.Name
	info.Status = req.Status
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()
	// 调用更新方法
	result, err := dao.Example2.Save(info)
	if err != nil {
		return 0, err
	}
	// 获取受影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *example2Service) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	idsArr := convert.ToInt64Array(ids, ",")
	result, err := dao.Example2.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响的行数
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *example2Service) Status(req *model.Example2StatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Example2.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Example2.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Example2.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
