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
 * 案例演示管理-服务类
 * @author 半城风雨
 * @since 2021/08/06
 * @File : example
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
var Example = new(exampleService)

type exampleService struct{}

func (s *exampleService) GetList(req *model.ExampleQueryReq) ([]model.ExampleInfoVo, int, error) {
	// 创建查询对象
	query := dao.Example.Where("mark=1")
	// 查询条件
	if req != nil {

		// 测试名称

		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}

		// 状态：1正常 2停用

		if req.Status > 0 {
			query = query.Where("status = ?", req.Status)
		}

		// 类型：1京东 2淘宝 3拼多多 4唯品会

		if req.Type > 0 {
			query = query.Where("type = ?", req.Type)
		}

		// 是否VIP：1是 2否

		if req.IsVip > 0 {
			query = query.Where("is_vip = ?", req.IsVip)
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
	var list []model.Example
	query.Structs(&list)

	// 数据处理
	var result []model.ExampleInfoVo
	for _, v := range list {
		item := model.ExampleInfoVo{}
		item.Example = v

		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}

		result = append(result, item)
	}

	// 返回结果
	return result, count, nil
}

func (s *exampleService) Add(req *model.ExampleAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}

	// 实例化模型
	var entity model.Example

	entity.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.Content = req.Content
	entity.Status = req.Status
	entity.Type = req.Type
	entity.IsVip = req.IsVip
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.Example.Insert(entity)
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

func (s *exampleService) Update(req *model.ExampleUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 获取记录信息
	info, err := dao.Example.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 对象赋值

	info.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		info.Avatar = avatar
	}
	info.Content = req.Content
	info.Status = req.Status
	info.Type = req.Type
	info.IsVip = req.IsVip
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()
	// 调用更新方法
	result, err := dao.Example.Save(info)
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

func (s *exampleService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	idsArr := convert.ToInt64Array(ids, ",")
	result, err := dao.Example.Delete("id in (?)", idsArr)
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

func (s *exampleService) Status(req *model.ExampleStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Example.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Example.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Example.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *exampleService) IsVip(req *model.ExampleIsVipReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Example.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置是否VIP
	result, err := dao.Example.Data(g.Map{
		"is_vip":      req.IsVip,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Example.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
