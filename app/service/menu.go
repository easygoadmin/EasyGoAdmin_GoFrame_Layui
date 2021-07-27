/**
 *
 * @author 摆渡人
 * @since 2021/5/19
 * @File : menu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// 中间件管理服务
var Menu = new(menuService)

type menuService struct{}

// 获取菜单权限列表
func (s *menuService) GetPermissionList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		menuList, _ := getTreeList()
		//fmt.Print(menuList)
		return menuList
	} else {
		// 非管理员
		return nil
	}
}

// 获取子级菜单
func getTreeList() (interface{}, error) {
	var catNet model.TreeNode
	data, err := dao.Menu.Where("type<=3 and status=1 and mark=1").Fields("id,name,pid,icon,url").Order("sort").FindAll()
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeTree(data, &catNet)
	return catNet.Children, nil
}

//递归生成分类列表
func makeTree(menu []*model.Menu, tn *model.TreeNode) {
	for _, c := range menu {
		if c.Pid == tn.Id {
			child := &model.TreeNode{}
			child.Menu = *c
			tn.Children = append(tn.Children, child)
			fmt.Println(child)
			makeTree(menu, child)
		}
	}
}

func (s *menuService) List(req *model.MenuQueryReq) []model.Menu {
	// 创建查询条件
	query := dao.Menu.Where("mark=1")
	// 查询条件
	if req != nil {
		// 菜单名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.Menu
	query.Structs(&list)
	return list
}

func (s *menuService) Add(req *model.MenuAddReq) (int64, error) {
	// 实例化对象
	var entity model.Menu
	entity.Name = req.Name
	entity.Icon = req.Icon
	entity.Url = req.Url
	entity.Param = req.Param
	entity.Pid = req.Pid
	entity.Type = req.Type
	entity.Permission = req.Permission
	entity.Status = req.Status
	entity.IsPublic = req.IsPublic
	entity.Note = req.Note
	entity.Sort = req.Sort
	entity.CreateUser = 1
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.Menu.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 添加节点
	setPermission(req.Type, req.Func, req.Url, gconv.Int(id))

	return id, nil
}

func (s *menuService) Update(req *model.MenuUpdateReq) (int64, error) {
	// 查询记录
	info, err := dao.Menu.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置参数值
	info.Name = req.Name
	info.Icon = req.Icon
	info.Url = req.Url
	info.Param = req.Param
	info.Pid = req.Pid
	info.Type = req.Type
	info.Permission = req.Permission
	info.Status = req.Status
	info.IsPublic = req.IsPublic
	info.Note = req.Note
	info.Sort = req.Sort
	info.UpdateUser = 1
	info.UpdateTime = gtime.Now()

	// 更新数据
	result, err := dao.Menu.Save(info)
	if err != nil {
		return 0, err
	}

	// 添加节点
	setPermission(req.Type, req.Func, req.Url, req.Pid)

	// 获取数影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *menuService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")

	// 删除记录
	result, err := dao.Menu.Delete("id in (?)", idsArr)
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

// 添加节点
func setPermission(menuType int, funcIds string, url string, pid int) {
	if menuType != 3 || funcIds == "" || url == "" {
		return
	}
	// 删除现有节点
	dao.Menu.Delete("pid=?", pid)

	// 创建权限节点
	urlArr := strings.Split(url, "/")
	if len(urlArr) == 3 {
		// 模块名
		moduleName := urlArr[1]
		// 节点处理
		funcArr := strings.Split(funcIds, ",")
		for _, v := range funcArr {
			// 实例化对象
			var entity model.Menu
			// 节点索引
			value := gconv.Int(v)
			if value == 1 {
				entity.Name = "列表"
				entity.Url = "/" + moduleName + "/list"
				entity.Permission = "sys:" + moduleName + ":list"
			} else if value == 5 {
				entity.Name = "添加"
				entity.Url = "/" + moduleName + "/add"
				entity.Permission = "sys:" + moduleName + ":add"
			} else if value == 10 {
				entity.Name = "修改"
				entity.Url = "/" + moduleName + "/update"
				entity.Permission = "sys:" + moduleName + ":update"
			} else if value == 15 {
				entity.Name = "删除"
				entity.Url = "/" + moduleName + "/delete"
				entity.Permission = "sys:" + moduleName + ":delete"
			} else if value == 20 {
				entity.Name = "详情"
				entity.Url = "/" + moduleName + "/detail"
				entity.Permission = "sys:" + moduleName + ":detail"
			} else if value == 25 {
				entity.Name = "状态"
				entity.Url = "/" + moduleName + "/status"
				entity.Permission = "sys:" + moduleName + ":status"
			} else if value == 30 {
				entity.Name = "批量删除"
				entity.Url = "/" + moduleName + "/dall"
				entity.Permission = "sys:" + moduleName + ":dall"
			} else if value == 35 {
				entity.Name = "添加子级"
				entity.Url = "/" + moduleName + "/addz"
				entity.Permission = "sys:" + moduleName + ":addz"
			} else if value == 40 {
				entity.Name = "全部展开"
				entity.Url = "/" + moduleName + "/expand"
				entity.Permission = "sys:" + moduleName + ":expand"
			} else if value == 45 {
				entity.Name = "全部折叠"
				entity.Url = "/" + moduleName + "/collapse"
				entity.Permission = "sys:" + moduleName + ":collapse"
			}
			entity.Pid = pid
			entity.Type = 4
			entity.Status = 1
			entity.IsPublic = 2
			entity.Sort = value
			// 插入节点
			dao.Menu.Insert(entity)
		}

	}
}
