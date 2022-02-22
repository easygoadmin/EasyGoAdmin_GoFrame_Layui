// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------
// | 免责声明:
// | 本软件框架禁止任何单位和个人用于任何违法、侵害他人合法利益等恶意的行为，禁止用于任何违
// | 反我国法律法规的一切平台研发，任何单位和个人使用本软件框架用于产品研发而产生的任何意外
// | 、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、附带
// | 或衍生的损失等)，本团队不承担任何法律责任。本软件框架只能用于公司和个人内部的法律所允
// | 许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 角色菜单-服务类
 * @author 半城风雨
 * @since 2021/7/15
 * @File : rolemenu
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"easygoadmin/app/utils/function"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
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
	if utils.AppDebug() {
		return gerror.New("演示环境，暂无权限操作")
	}
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
