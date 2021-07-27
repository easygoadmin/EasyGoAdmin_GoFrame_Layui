/**
 *
 * @author 摆渡人
 * @since 2021/7/22
 * @File : link
 */
package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
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

func (s *linkService) Add(req *model.LinkAddReq) (int64, error) {
	// 实例化对象
	var entity model.Link
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.Platform = req.Platform
	entity.Form = req.Form
	entity.Image = req.Image
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = 1
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

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

func (s *linkService) Update(req *model.LinkUpdateReq) (int64, error) {
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
	info.UpdateUser = 1
	info.UpdateTime = gtime.Now()

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
