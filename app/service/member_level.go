/**
 *
 * @author 摆渡人
 * @since 2021/7/29
 * @File : member_level
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var MemberLevel = new(memberLevelService)

type memberLevelService struct{}

func (s *memberLevelService) GetList(req *model.MemberLevelPageReq) ([]model.MemberLevel, int, error) {
	// 创建查询实例
	query := dao.MemberLevel.Where("mark=1")
	// 查询条件
	if req != nil {
		// 等级名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("sort asc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.MemberLevel
	query.Structs(&list)
	return list, count, err
}

func (s *memberLevelService) Add(req *model.MemberLevelAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.MemberLevel
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.MemberLevel.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *memberLevelService) Update(req *model.MemberLevelUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.MemberLevel.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置参数
	info.Name = req.Name
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.MemberLevel.Save(info)
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *memberLevelService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.MemberLevel.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
