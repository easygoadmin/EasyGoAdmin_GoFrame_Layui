// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// DictData is the golang structure for table sys_dict_data.
type DictData struct {
    Id         int         `orm:"id,primary"  json:"id"`         // 主键ID             
    Name       string      `orm:"name,unique" json:"name"`       // 字典项名称         
    Code       string      `orm:"code"        json:"code"`       // 字典项值           
    DictId     int         `orm:"dict_id"     json:"dictId"`     // 字典类型ID         
    Status     int         `orm:"status"      json:"status"`     // 状态：1在用 2停用  
    Note       string      `orm:"note"        json:"note"`       // 备注               
    Sort       int         `orm:"sort"        json:"sort"`       // 显示顺序           
    CreateUser int         `orm:"create_user" json:"createUser"` // 添加人             
    CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 添加时间           
    UpdateUser int         `orm:"update_user" json:"updateUser"` // 更新人             
    UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间           
    Mark       int         `orm:"mark"        json:"mark"`       // 有效标记           
}