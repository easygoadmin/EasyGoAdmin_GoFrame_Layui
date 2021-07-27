/**
 *
 * @author 摆渡人
 * @since 2021/7/15
 * @File : rolemenu
 */
package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"easygoadmin/app/utils/function"
)

var RoleMenu = new(roleMenuService)

type roleMenuService struct{}

func (s *roleMenuService) GetRoleMenuList(roleId int64) ([]model.RoleMenuInfo, error) {
	// 获取全部菜单列表
	var menuList []model.Menu
	dao.Menu.Where("status=1").Where("mark=1").Order("sort asc").Structs(&menuList)
	if len(menuList) == 0 {
		return nil, gerror.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	var roleMenuList []model.RoleMenu
	dao.RoleMenu.Where("role_id=?", roleId).Structs(&roleMenuList)
	idList := gutil.ListItemValuesUnique(&roleMenuList, "MenuId")

	// 对象处理
	var list []model.RoleMenuInfo
	if len(menuList) > 0 {
		for _, m := range menuList {
			var info model.RoleMenuInfo
			info.Id = m.Id
			info.Name = m.Name
			info.Open = true
			info.Pid = m.Pid
			// 节点选中值
			if function.InArray(gconv.String(m.Id), idList) {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

func (s *roleMenuService) Save(req *model.RoleMenuSaveReq) error {
	itemArr := convert.ToInt64Array(req.MenuIds, ",")
	if len(itemArr) == 0 {
		return gerror.New("请选择权限节点")
	}
	// 删除现有的角色权限数据
	dao.RoleMenu.Delete("role_id=?", req.RoleId)
	// 遍历创建新角色权限数据
	for i := range itemArr {
		var entity model.RoleMenu
		entity.RoleId = req.RoleId
		entity.MenuId = gconv.Int(itemArr[i])
		dao.RoleMenu.Insert(entity)
	}
	// 批量插入
	//dao.RoleMenu.Data(list).Batch(2).Insert()
	return nil
}
