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
 * 会员管理-服务类
 * @author 半城风雨
 * @since 2021/7/29
 * @File : member
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var Member = new(memberService)

type memberService struct{}

func (s *memberService) GetList(req *model.MemberPageReq) ([]model.MemberInfoVo, int, error) {
	// 创建查询实例
	query := dao.Member.Where("mark=1")
	// 查询条件
	if req != nil {
		// 用户名
		if req.Username != "" {
			query = query.Where("username=?", req.Username)
		}
	}
	// 查询记录总数
	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 排序
	query = query.Order("id desc")
	// 分页
	query = query.Page(req.Page, req.Limit)
	// 对象转换
	var list []model.Member
	query.Structs(&list)

	// 数据处理
	var result = make([]model.MemberInfoVo, 0)
	for _, v := range list {
		item := model.MemberInfoVo{}
		item.Member = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = utils.GENDER_LIST[v.Gender]
		}
		// 设备类型
		if v.Device > 0 {
			item.DeviceName = common.MEMBER_DEVICE_LIST[v.Device]
		}
		// 会员来源
		if v.Source > 0 {
			item.SourceName = common.MEMBER_SOURCE_LIST[v.Source]
		}
		// 所属城市
		if v.DistrictCode != "" {
			item.CityName = City.GetCityName(v.DistrictCode, ">>")
		}
		// 加入数组
		result = append(result, item)
	}
	return result, count, nil
}

func (s *memberService) Add(req *model.MemberAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Member
	entity.Username = req.Username
	entity.MemberLevel = req.MemberLevel
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Birthday = req.Birthday
	entity.ProvinceCode = req.ProvinceCode
	entity.CityCode = req.CityCode
	entity.DistrictCode = req.DistrictCode
	entity.Address = req.Address
	entity.Intro = req.Intro
	entity.Signature = req.Signature
	entity.Device = req.Device
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 头像处理
	avatar, err := utils.SaveImage(req.Avatar, "member")
	if err != nil {
		return 0, err
	}
	entity.Avatar = avatar

	// 插入数据
	result, err := dao.Member.Insert(entity)
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

func (s *memberService) Update(req *model.MemberUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Member.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, err
	}
	info.Username = req.Username
	info.MemberLevel = req.MemberLevel
	info.Realname = req.Realname
	info.Nickname = req.Nickname
	info.Gender = req.Gender
	info.Birthday = req.Birthday
	info.ProvinceCode = req.ProvinceCode
	info.CityCode = req.CityCode
	info.DistrictCode = req.DistrictCode
	info.Address = req.Address
	info.Intro = req.Intro
	info.Signature = req.Signature
	info.Device = req.Device
	info.Source = req.Source
	info.Status = req.Status
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 头像处理
	avatar, err := utils.SaveImage(req.Avatar, "member")
	if err != nil {
		return 0, err
	}
	info.Avatar = avatar

	// 调用更新方法
	result, err := dao.Member.Save(info)
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

func (s *memberService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.Member.Delete("id in (?)", idsArr)
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

func (s *memberService) Status(req *model.MemberStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Member.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Member.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Member.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
