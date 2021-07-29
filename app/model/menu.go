// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// Menu is the golang structure for table sys_menu.
type Menu internal.Menu

// Fill with you ideas below.

// 菜单Vo
type TreeNode struct {
	Menu
	Children []*TreeNode `json:"children"` // 子菜单
}

// 列表查询条件
type MenuQueryReq struct {
	Name string `p:"name"` // 菜单名称
}

// 添加菜单
type MenuAddReq struct {
	Name       string `p:"name" v:"required#菜单名称不能为空"`  // 菜单名称
	Icon       string `p:"icon" v:"required#请选择菜单图标"`   // 图标
	Url        string `p:"url" v:"required#菜单URL不能为空"`  // URL地址
	Param      string `p:"param"`                       // 参数
	Pid        int    `p:"pid"`                         // 上级ID
	Type       int    `p:"type"v:"required#请选择菜单类型"`    // 类型：1模块 2导航 3菜单 4节点
	Permission string `p:"permission"`                  // 权限标识
	Status     int    `p:"status" v:"required#请选择菜单状态"` // 是否显示：1显示 2不显示
	IsPublic   int    `p:"is_public"`                   // 是否公共：1是 2否
	Note       string `p:"note"`                        // 菜单备注
	Sort       int    `p:"sort" v:"required请输入菜单排序号"`   // 显示顺序
	Func       string `p:"func"`                        // 权限节点
}

// 更新菜单
type MenuUpdateReq struct {
	Id         int    `p:id v:"required#主键ID不能为空"`
	Name       string `p:"name" v:"required#菜单名称不能为空"`  // 菜单名称
	Icon       string `p:"icon" v:"required#请选择菜单图标"`   // 图标
	Url        string `p:"url" v:"required#菜单URL不能为空"`  // URL地址
	Param      string `p:"param"`                       // 参数
	Pid        int    `p:"pid"`                         // 上级ID
	Type       int    `p:"type"v:"required#请选择菜单类型"`    // 类型：1模块 2导航 3菜单 4节点
	Permission string `p:"permission"`                  // 权限标识
	Status     int    `p:"status" v:"required#请选择菜单状态"` // 是否显示：1显示 2不显示
	IsPublic   int    `p:"is_public"`                   // 是否公共：1是 2否
	Note       string `p:"note"`                        // 菜单备注
	Sort       int    `p:"sort" v:"required请输入菜单排序号"`   // 显示顺序
	Func       string `p:"func"`                        // 权限节点
}

type MenuDeleteReq struct {
	Ids string `p:ids v:"require#请选择需要删除的数据记录"`
}