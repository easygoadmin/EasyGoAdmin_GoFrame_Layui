/**
 *
 * @author 摆渡人
 * @since 2021/7/16
 * @File : dept
 */
package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
)

// 中间件管理服务
var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req *model.DeptPageReq) []model.Dept {
	// 创建查询对象
	query := dao.Dept.Where("mark=1")
	// 查询条件
	if req != nil {
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.Dept
	query.Structs(&list)
	return list
}

func (s *deptService) Add(req *model.DeptAddReq) (int64, error) {
	// 实例化对象
	var data model.Dept
	data.Name = req.Name
	data.Code = req.Code
	data.Fullname = req.Fullname
	data.Type = req.Type
	data.Sort = req.Sort
	data.Note = req.Note
	data.CreateUser = 1
	data.CreateTime = gtime.Now()
	data.Mark = 1

	// 插入记录
	result, err := dao.Dept.Insert(data)
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
func (s *deptService) Update(req *model.DeptUpdateReq) (int64, error) {
	// 查询记录
	info, err := dao.Dept.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置参数
	info.Id = req.Id
	info.Name = req.Name
	info.Code = req.Code
	info.Fullname = req.Fullname
	info.Type = req.Type
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = 1
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.Dept.Save(info)
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

func (s *deptService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.Dept.Delete("id in (?)", idsArr)
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
