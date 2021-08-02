/**
 *
 * @author 摆渡人
 * @since 2021/7/13
 * @File : level
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req *model.LevelQueryReq) ([]model.Level, int, error) {
	query := dao.Level.Clone()
	query = query.Where("mark=1")
	if req != nil {
		// 职级名称查询
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
	var list []model.Level
	query.Structs(&list)
	return list, count, nil
}

func (s *levelService) Add(req *model.LevelAddReq, userId int) (int64, error) {
	// 模型
	var entity model.Level
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1
	// 插入数据
	result, err := dao.Level.Insert(entity)
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

func (s *levelService) Update(req *model.LevelEditReq, userId int) (int64, error) {
	// 查询当前记录是否存在
	info, err := dao.Level.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	info.Name = req.Name
	info.Status = req.Status
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()
	result, err := dao.Level.Save(info)
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

// 删除
func (s *levelService) Delete(Ids string) int64 {
	idsArr := convert.ToInt64Array(Ids, ",")
	result, err := dao.Level.Delete("id in (?)", idsArr)
	if err != nil {
		return 0
	}
	res, err := result.RowsAffected()
	return res
}

func (s *levelService) Status(req *model.LevelStatusReq, userId int) (int64, error) {
	info, err := dao.Level.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Level.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Level.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
