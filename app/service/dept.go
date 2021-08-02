/**
 *
 * @author 摆渡人
 * @since 2021/7/16
 * @File : dept
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"reflect"
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

func (s *deptService) Add(req *model.DeptAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.Dept.Insert(entity)
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
func (s *deptService) Update(req *model.DeptUpdateReq, userId int) (int64, error) {
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
	info.Pid = req.Pid
	info.Sort = req.Sort
	info.Note = req.Note
	info.UpdateUser = userId
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

// 获取子级菜单
func (s *deptService) GetDeptTreeList() ([]*model.DeptTreeNode, error) {
	var deptNode model.DeptTreeNode
	// 查询列表
	list, _ := dao.Dept.Where("mark=1").Fields("id,name,pid").Order("sort asc").FindAll()
	makeDeptTree(list, &deptNode)
	return deptNode.Children, nil
}

//递归生成分类列表
func makeDeptTree(cate []*model.Dept, tn *model.DeptTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &model.DeptTreeNode{}
			child.Dept = *c
			tn.Children = append(tn.Children, child)
			makeDeptTree(cate, child)
		}
	}
}

// 数据源转换
func (s *deptService) MakeList(data []*model.DeptTreeNode) map[int]string {
	deptList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			deptList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				deptList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					deptList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return deptList
}
