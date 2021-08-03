/**
 *
 * @author 摆渡人
 * @since 2021/7/24
 * @File : ad_sort
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/common"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
)

// 中间件管理服务
var AdSort = new(adSortService)

type adSortService struct{}

func (s *adSortService) GetList(req *model.AdSortPageReq) ([]model.AdSortInfoVo, int, error) {
	// 创建查询实例
	query := dao.AdSort.Where("mark=1")
	// 查询条件
	if req != nil {
		// 广告位名称
		if req.Description != "" {
			query = query.Where("description like ?", "%"+req.Description+"%")
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
	var list []model.AdSort
	query.Structs(&list)

	// 数据处理
	var result = make([]model.AdSortInfoVo, 0)
	for _, v := range list {
		platformName, ok := common.ADSORT_PLATFORM_LIST[v.Platform]
		item := model.AdSortInfoVo{}
		item.AdSort = v
		if ok {
			item.PlatformName = platformName
		}
		// 站点名称
		if v.ItemId > 0 {
			info, err := dao.Item.FindOne("id=?", v.ItemId)
			if err == nil && info != nil {
				item.ItemName = info.Name
			}
		}

		// 栏目名称
		if v.CateId > 0 {
			cateName := ItemCate.GetCateName(v.CateId, ">>")
			item.CateName = cateName
		}

		// 加入数组
		result = append(result, item)
	}

	return result, count, nil
}

func (s *adSortService) Add(req *model.AdSortAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.AdSort
	entity.Description = req.Description
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.LocId = req.LocId
	entity.Platform = req.Platform
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.AdSort.Insert(entity)
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

func (s *adSortService) Update(req *model.AdSortUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.AdSort.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Description = req.Description
	info.ItemId = req.ItemId
	info.CateId = req.CateId
	info.LocId = req.LocId
	info.Platform = req.Platform
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.AdSort.Save(info)
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

func (s *adSortService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.AdSort.Delete("id in (?)", idsArr)
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
