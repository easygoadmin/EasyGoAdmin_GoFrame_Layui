// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 友链管理-服务类
 * @author 半城风雨
 * @since 2021/7/22
 * @File : link
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
var Link = new(linkService)

type linkService struct{}

func (s *linkService) GetList(req *model.LinkPageReq) ([]model.LinkInfoVo, int, error) {
	// 实例化对象
	query := dao.Link.Where("mark=1")
	// 查询条件
	if req != nil {
		// 友链名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
		// 友链类型
		if req.Type > 0 {
			query = query.Where("type=?", req.Type)
		}
		// 投放平台
		if req.Platform > 0 {
			query = query.Where("platform=?", req.Platform)
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
	var list []model.Link
	query.Structs(&list)

	var result []model.LinkInfoVo
	for _, v := range list {
		item := model.LinkInfoVo{}
		item.Link = v
		// 友链图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		// 友链类型
		typeName, ok := common.LINK_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 友链形式
		formName, ok := common.LINK_FORM_LIST[v.Form]
		if ok {
			item.FormName = formName
		}
		// 投放平台
		platformName, ok := common.LINK_PLATFORM_LIST[v.Platform]
		if ok {
			item.PlatformName = platformName
		}
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *linkService) Add(req *model.LinkAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.Link
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.Platform = req.Platform
	entity.Form = req.Form
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 插入数据
	result, err := dao.Link.Insert(entity)
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

func (s *linkService) Update(req *model.LinkUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.Link.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 保存图片
	image, err := utils.SaveImage(req.Image, "link")
	if err != nil {
		return 0, err
	}
	info.Image = image

	// 设置对象
	info.Name = req.Name
	info.Type = req.Type
	info.Url = req.Url
	info.ItemId = req.ItemId
	info.CateId = req.CateId
	info.Platform = req.Platform
	info.Form = req.Form
	info.Status = req.Status
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		info.Image = image
	}

	// 更新记录
	result, err := dao.Link.Save(info)
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

func (s *linkService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.Link.Delete("id in (?)", idsArr)
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

func (s *linkService) Status(req *model.LinkStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	info, err := dao.Link.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Link.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Level.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
