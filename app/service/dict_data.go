/**
 *
 * @author 摆渡人
 * @since 2021/7/21
 * @File : dict_data
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
var DictData = new(dictDataService)

type dictDataService struct{}

func (s *dictDataService) GetList(req *model.DictDataPageReq) ([]model.DictData, int, error) {
	// 创建查询对象
	query := dao.DictData.Where("dict_id=?", req.DictId).Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典项名称
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
	var list []model.DictData
	query.Structs(&list)
	return list, count, nil
}

func (s *dictDataService) Add(req *model.DictDataAddReq) (int64, error) {
	// 实例化对象
	var entity model.DictData
	entity.DictId = req.DictId
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = 1
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.DictData.Insert(entity)
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

func (s *dictDataService) Update(req *model.DictDataUpdateReq) (int64, error) {
	// 查询记录
	info, err := dao.DictData.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.DictId = req.DictId
	info.Name = req.Name
	info.Code = req.Code
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = 1
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.DictData.Save(info)
	if err != nil {
		return 0, err
	}

	// 获取受影响函数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *dictDataService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.DictData.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}
	// 获取受影响函数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
