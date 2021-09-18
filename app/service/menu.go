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
 * 菜单管理-服务类
 * @author 半城风雨
 * @since 2021/5/19
 * @File : menu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"strings"
)

// 中间件管理服务
var Menu = new(menuService)

type menuService struct{}

// 获取菜单权限列表
func (s *menuService) GetPermissionMenuList(userId int) interface{} {
	if userId == 1 {
		// 管理员(拥有全部权限)
		menuList, _ := Menu.GetTreeList()
		return menuList
	} else {
		// 非管理员
		// 创建查询实例
		query := dao.Menu.As("m").Clone()
		// 内联查询
		query = query.InnerJoin("sys_role_menu as r", "m.id = r.menu_id")
		query = query.InnerJoin("sys_user_role ur", "ur.role_id=r.role_id")
		query = query.Where("ur.user_id=? AND m.type=0 AND m.`status`=1 AND m.mark=1", userId)
		// 获取字段
		query.Fields("m.*")
		// 排序
		query = query.Order("m.id asc")
		// 数据转换
		var list []*model.Menu
		query.Structs(&list)
		// 数据处理
		var menuNode model.TreeNode
		makeTree(list, &menuNode)
		return menuNode.Children
	}
}

// 获取子级菜单
func (s *menuService) GetTreeList() ([]*model.TreeNode, error) {
	var menuNode model.TreeNode
	data, err := dao.Menu.Where("type=0 and mark=1").Fields("id,name,pid,icon,url,target").Order("sort").FindAll()
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeTree(data, &menuNode)
	return menuNode.Children, nil
}

//递归生成分类列表
func makeTree(menu []*model.Menu, tn *model.TreeNode) {
	for _, c := range menu {
		if c.Pid == tn.Id {
			child := &model.TreeNode{}
			child.Menu = *c
			tn.Children = append(tn.Children, child)
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

func (s *menuService) Add(req *model.MenuAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
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
	entity.Target = req.Target
	entity.Note = req.Note
	entity.Sort = req.Sort
	entity.CreateUser = userId
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
	setPermission(req.Type, req.Func, req.Url, gconv.Int(id), userId)

	return id, nil
}

func (s *menuService) Update(req *model.MenuUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
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
	info.Target = req.Target
	info.Note = req.Note
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新数据
	result, err := dao.Menu.Save(info)
	if err != nil {
		return 0, err
	}

	// 添加节点
	setPermission(req.Type, req.Func, req.Url, req.Id, userId)

	// 获取数影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *menuService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")

	// 判断是否有子级
	child, err := dao.Menu.Where("pid in (?)", idsArr).Count()
	if err != nil {
		return 0, err
	}
	if child > 0 {
		return 0, gerror.New("有子级无法删除")
	}

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
func setPermission(menuType int, funcIds string, url string, pid int, userId int) {
	if menuType != 0 || funcIds == "" || url == "" {
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
			} else if value == 50 {
				entity.Name = "导出数据"
				entity.Url = "/" + moduleName + "/export"
				entity.Permission = "sys:" + moduleName + ":export"
			} else if value == 55 {
				entity.Name = "导入数据"
				entity.Url = "/" + moduleName + "/import"
				entity.Permission = "sys:" + moduleName + ":import"
			} else if value == 60 {
				entity.Name = "分配权限"
				entity.Url = "/" + moduleName + "/permission"
				entity.Permission = "sys:" + moduleName + ":permission"
			} else if value == 65 {
				entity.Name = "重置密码"
				entity.Url = "/" + moduleName + "/resetPwd"
				entity.Permission = "sys:" + moduleName + ":resetPwd"
			}
			entity.Pid = pid
			entity.Type = 1
			entity.Status = 1
			entity.Target = 1
			entity.Sort = value
			entity.CreateUser = userId
			entity.CreateTime = gtime.Now()
			entity.UpdateUser = userId
			entity.UpdateTime = gtime.Now()
			entity.Mark = 1

			// 插入节点
			dao.Menu.Insert(entity)
		}
	}
}

// 数据源转换
func (s *menuService) MakeList(data []*model.TreeNode) map[int]string {
	menuList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			menuList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				menuList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					menuList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return menuList
}
