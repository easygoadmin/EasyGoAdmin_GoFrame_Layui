/**
 *
 * @author 摆渡人
 * @since 2021/7/21
 * @File : config
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
var Config = new(configService)

type configService struct{}

func (s *configService) GetList(req *model.ConfigQueryReq) []model.Config {
	// 创建查询对象
	query := dao.Config.Where("mark=1")
	// 查询条件
	if req != nil {
		// 字典名称
		if req.Name != "" {
			query = query.Where("name like ?", req.Name)
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.Config
	query.Structs(&list)
	return list
}

func (s *configService) Add(req *model.ConfigAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Config
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.Config.Insert(entity)
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

func (s *configService) Update(req *model.ConfigUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.Config.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Name = req.Name
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新数据
	result, err := dao.Config.Save(info)
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

func (s *configService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.Config.Delete("id in (?)", idsArr)
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
