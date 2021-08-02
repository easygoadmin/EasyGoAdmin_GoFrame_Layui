/**
 *
 * @author 摆渡人
 * @since 2021/7/21
 * @File : config_data
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var ConfigData = new(configDataService)

type configDataService struct{}

func (s *configDataService) GetList(req *model.ConfigDataPageReq) ([]model.ConfigDataVo, int, error) {
	// 创建查询对象
	query := dao.ConfigData.Where("config_id=?", req.ConfigId).Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典项名称
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
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
	var list []model.ConfigData
	query.Structs(&list)

	var result = make([]model.ConfigDataVo, 0)
	for _, v := range list {
		typeName, ok := common.CONFIG_DATA_TYPE_LIST[v.Type]
		item := model.ConfigDataVo{}
		item.ConfigData = v
		if ok {
			item.TypeName = typeName
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *configDataService) Add(req *model.ConfigDataAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.ConfigData
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.ConfigData.Insert(entity)
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

func (s *configDataService) Update(req *model.ConfigDataUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.ConfigData.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Title = req.Title
	info.Code = req.Code
	info.Value = req.Value
	info.Options = req.Options
	info.ConfigId = req.ConfigId
	info.Type = req.Type
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.ConfigData.Save(info)
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

func (s *configDataService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.ConfigData.Delete("id in (?)", idsArr)
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

func (s *configDataService) Status(req *model.ConfigDataStatusReq, userId int) (int64, error) {
	info, err := dao.ConfigData.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.ConfigData.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.ConfigData.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
