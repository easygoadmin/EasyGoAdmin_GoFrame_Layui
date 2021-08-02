/**
 *
 * @author 摆渡人
 * @since 2021/7/24
 * @File : item
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var Item = new(itemService)

type itemService struct{}

func (s *itemService) GetList(req *model.ItemPageReq) ([]model.ItemInfoVo, int, error) {
	// 创建查询实例
	query := dao.Item.Where("mark=1")
	// 查询条件
	if req != nil {
		// 站点名称
		if req.Name != "" {
			query = query.Where("name like %?%", req.Name)
		}
		// 站点类型
		if req.Type > 0 {
			query = query.Where("type=?", req.Type)
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
	var list []model.Item
	query.Structs(&list)

	// 数据处理
	var result = make([]model.ItemInfoVo, 0)
	for _, v := range list {
		item := model.ItemInfoVo{}
		item.Item = v
		// 站点类型
		typeName, ok := common.ITEM_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 站点图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		result = append(result, item)
	}
	return result, count, nil
}

func (s *itemService) Add(req *model.ItemAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Item
	entity.Name = req.Name
	entity.Type = req.Type
	entity.Url = req.Url
	entity.Status = req.Status
	entity.Note = req.Name
	entity.Sort = req.Status
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 图片处理
	image, err := utils.SaveImage(req.Image, "item")
	if err != nil {
		return 0, err
	}
	entity.Image = image

	// 插入数据
	result, err := dao.Item.Insert(entity)
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

func (s *itemService) Update(req *model.ItemUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.Item.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Name = req.Name
	info.Type = req.Type
	info.Url = req.Url
	info.Status = req.Status
	info.Note = req.Name
	info.Sort = req.Status

	// 图片处理
	image, err := utils.SaveImage(req.Image, "item")
	if err != nil {
		return 0, err
	}
	info.Image = image
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.Item.Save(info)
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

func (s *itemService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	result, err := dao.Item.Delete("id in (?)", idsArr)
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
