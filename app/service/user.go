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
 * 用户管理-服务类
 * @author 半城风雨
 * @since 2021/7/27
 * @File : user
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// 中间件管理服务
var User = new(userService)

type userService struct{}

func (s *userService) GetList(req *model.UserPageReq) ([]model.UserInfoVo, int, error) {
	// 创建查询实例
	query := dao.User.Where("mark=1")
	// 查询条件
	if req != nil {
		// 用户姓名
		if req.Realname != "" {
			query = query.Where("realname like ?", "%"+req.Realname+"%")
		}
		// 性别
		if req.Gender > 0 {
			query = query.Where("gender=?", req.Gender)
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
	var list []model.User
	query.Structs(&list)

	// 获取职级列表
	levelList, _ := dao.Level.Where("mark=1").Fields("id,name").All()
	var levelMap = make(map[int]string)
	for _, v := range levelList {
		levelMap[v.Id] = v.Name
	}
	// 获取岗位列表
	positionList, _ := dao.Position.Where("mark=1").Fields("id,name").All()
	var positionMap = make(map[int]string)
	for _, v := range positionList {
		positionMap[v.Id] = v.Name
	}
	// 部门
	deptList, _ := dao.Dept.Where("mark=1").Fields("id,name").All()
	var deptMap = make(map[int]string)
	for _, v := range deptList {
		deptMap[v.Id] = v.Name
	}

	// 数据处理
	var result []model.UserInfoVo
	for _, v := range list {
		item := model.UserInfoVo{}
		item.User = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = utils.GENDER_LIST[v.Gender]
		}
		// 职级
		if v.LevelId > 0 {
			item.LevelName = levelMap[v.LevelId]
		}
		// 岗位
		if v.PositionId > 0 {
			item.PositionName = positionMap[v.PositionId]
		}
		// 部门
		if v.DeptId > 0 {
			item.DeptName = deptMap[v.DeptId]
		}
		// 角色信息
		queryRole := dao.Role.As("r").Clone()
		queryRole = queryRole.InnerJoin("sys_user_role as ur", "r.id = ur.role_id")
		queryRole = queryRole.Where("ur.user_id=? AND r.mark=1", v.Id)
		// 获取字段
		queryRole.Fields("r.*")
		// 排序
		queryRole = queryRole.Order("r.sort asc")
		// 数据转换
		var roles []*model.Role
		queryRole.Structs(&roles)
		// 数据重组
		roleList := make([]string, 0)
		if len(roles) > 0 {
			for _, val := range roles {
				roleList = append(roleList, val.Name)
			}
		}
		if len(roleList) > 0 {
			item.RoleName = strings.Join(roleList, "、")
		}

		result = append(result, item)
	}
	return result, count, nil
}

func (s *userService) Add(req *model.UserAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.User
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Avatar = req.Avatar
	entity.Mobile = req.Mobile
	entity.Email = req.Email
	entity.Birthday = req.Birthday
	entity.DeptId = req.DeptId
	entity.LevelId = req.LevelId
	entity.PositionId = req.PositionId
	entity.ProvinceCode = req.ProvinceCode
	entity.CityCode = req.CityCode
	entity.DistrictCode = req.DistrictCode
	entity.Address = req.Address
	entity.Username = req.Username
	entity.Intro = req.Intro
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入对象
	result, err := dao.User.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 删除用户角色关系
	dao.UserRole.Delete("user_id=?", entity.Id)
	// 创建人员角色关系
	roleIds := strings.Split(req.RoleIds, ",")
	for _, v := range roleIds {
		if gconv.Int(v) == 0 {
			continue
		}
		var userRole model.UserRole
		userRole.UserId = entity.Id
		userRole.RoleId = gconv.Int(v)
		dao.UserRole.Insert(userRole)
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *userService) Update(req *model.UserUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置对象
	info.Realname = req.Realname
	info.Nickname = req.Nickname
	info.Gender = req.Gender
	info.Avatar = req.Avatar
	info.Mobile = req.Mobile
	info.Email = req.Email
	info.Birthday = req.Birthday
	info.DeptId = req.DeptId
	info.LevelId = req.LevelId
	info.PositionId = req.PositionId
	info.ProvinceCode = req.ProvinceCode
	info.CityCode = req.CityCode
	info.DistrictCode = req.DistrictCode
	info.Address = req.Address
	info.Username = req.Username
	info.Intro = req.Intro
	info.Status = req.Status
	info.Note = req.Note
	info.Sort = req.Sort

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		info.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		info.Avatar = avatar
	}
	info.CreateUser = userId
	info.CreateTime = gtime.Now()

	// 更新记录
	result, err := dao.User.Save(info)
	if err != nil {
		return 0, err
	}

	// 删除用户角色关系
	dao.UserRole.Delete("user_id=?", info.Id)
	// 创建人员角色关系
	roleIds := strings.Split(req.RoleIds, ",")
	for _, v := range roleIds {
		if gconv.Int(v) == 0 {
			continue
		}
		var userRole model.UserRole
		userRole.UserId = info.Id
		userRole.RoleId = gconv.Int(v)
		dao.UserRole.Insert(userRole)
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.User.Delete("id in (?)", idsArr)
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

func (s *userService) Status(req *model.UserStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.User.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.User.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (s *userService) ResetPwd(id int, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.User.FindOne("id=?", id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置初始密码
	password, err := utils.Md5("123456" + info.Username)
	if err != nil {
		return 0, err
	}

	// 初始化密码
	result, err := dao.User.Data(g.Map{
		"password":    password,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.User.Columns.Id, info.Id).Update()
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

func (s *userService) UpdateUserInfo(req *model.UserInfoReq, session *ghttp.Session) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 用户ID
	userId := utils.Uid(session)
	// 更新用户信息
	result, err := dao.User.Data(g.Map{
		"realname": req.Realname,
		"email":    req.Email,
		"mobile":   req.Mobile,
		"address":  req.Address,
		"intro":    req.Intro,
	}).Where("id", userId).Update()
	if err != nil {
		return 0, err
	}

	// 获取受影响行数
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, gerror.New("更新失败")
	}

	// 获取信息
	userInfo, _ := dao.User.FindOne("id=?", userId)
	// 设置SESSON
	session.Set("userInfo", userInfo)
	return rows, nil
}

func (s *userService) UpdatePwd(req *model.UpdatePwd, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询信息
	info, err := dao.User.FindOne("id=?", userId)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 比对旧密码
	oldPwd, err := utils.Md5(req.OldPassword + info.Username)
	if err != nil {
		return 0, err
	}
	if oldPwd != info.Password {
		return 0, gerror.New("旧密码不正确")
	}

	// 设置新密码
	if req.NewPassword != req.RePassword {
		return 0, gerror.New("两次输入的新密码不一致")
	}
	newPwd, err := utils.Md5(req.NewPassword + info.Username)
	if err != nil {
		return 0, err
	}

	result, err := dao.User.Data(g.Map{
		"password": newPwd,
	}).Where(dao.User.Columns.Id, userId).Update()
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
