/**
 *
 * @author 摆渡人
 * @since 2021/7/26
 * @File : ad
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
	"github.com/gogf/gf/util/gconv"
)

// 中间件管理服务
var Ad = new(adService)

type adService struct{}

func (s *adService) GetList(req *model.AdPageReq) ([]model.AdInfoVo, int, error) {
	// 创建查询实例
	query := dao.Ad.Where("mark=1")
	// 查询条件
	if req != nil {
		// 广告标题
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
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
	var list []model.Ad
	query.Structs(&list)

	// 数据处理
	result := make([]model.AdInfoVo, 0)
	for _, v := range list {
		item := model.AdInfoVo{}
		item.Ad = v
		// 广告封面
		if v.Cover != "" {
			item.Cover = utils.GetImageUrl(v.Cover)
		}
		// 广告类型
		if v.Type > 0 {
			item.TypeName = common.AD_TYPE_LIST[v.Type]
		}
		// 所属广告位
		if v.AdSortId > 0 {
			itemInfo, err := dao.AdSort.FindOne("id=?", v.AdSortId)
			if err == nil && itemInfo != nil {
				item.AdSortDesc = itemInfo.Description + " >> " + gconv.String(itemInfo.LocId)
			}
		}
		result = append(result, item)
	}

	// 返回结果
	return result, count, err
}

func (s *adService) Add(req *model.AdAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Ad
	entity.Title = req.Title
	entity.AdSortId = req.AdSortId
	entity.Type = req.Type
	entity.Description = req.Description
	entity.Content = req.Content
	entity.Url = req.Url
	entity.Width = req.Width
	entity.Height = req.Height
	entity.StartTime = req.StartTime
	entity.EndTime = req.EndTime
	entity.Status = req.Status
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 广告封面
	if req.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		entity.Cover = ""
	}

	// 插入数据
	result, err := dao.Ad.Insert(entity)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 返回结果
	return id, nil
}

func (s *adService) Update(req *model.AdUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.Ad.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}
	// 设置对象
	info.Title = req.Title
	info.AdSortId = req.AdSortId
	info.Type = req.Type
	info.Description = req.Description
	info.Content = req.Content
	info.Url = req.Url
	info.Width = req.Width
	info.Height = req.Height
	info.StartTime = req.StartTime
	info.EndTime = req.EndTime
	info.Status = req.Status
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 广告封面
	if req.Type == 1 {
		// 图片
		cover, err := utils.SaveImage(req.Cover, "ad")
		if err != nil {
			return 0, err
		}
		info.Cover = cover
	} else {
		info.Cover = ""
	}

	// 更新信息
	result, err := dao.Ad.Save(info)
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

func (s *adService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除
	result, err := dao.Ad.Delete("id in (?)", idsArr)
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

func (s *adService) Status(req *model.AdStatusReq, userId int) (int64, error) {
	info, err := dao.Ad.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置状态
	result, err := dao.Ad.Data(g.Map{
		"status":      req.Status,
		"update_user": userId,
		"update_time": gtime.Now(),
	}).Where(dao.Ad.Columns.Id, info.Id).Update()
	if err != nil {
		return 0, err
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return res, nil
}
