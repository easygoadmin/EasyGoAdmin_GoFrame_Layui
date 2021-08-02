/**
 *
 * @author 摆渡人
 * @since 2021/7/26
 * @File : notice
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
var Notice = new(noticeService)

type noticeService struct{}

func (s *noticeService) GetList(req *model.NoticePageReq) ([]model.NoticeInfoVo, int, error) {
	// 创建查询实例
	query := dao.Notice.Where("mark=1")
	// 查询条件
	if req != nil {
		// 通知标题
		if req.Title != "" {
			query = query.Where("title like ?", "%"+req.Title+"%")
		}
		// 通知来源
		if req.Source > 0 {
			query = query.Where("source=?", req.Source)
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
	var list []model.Notice
	query.Structs(&list)

	// 数据处理
	var result []model.NoticeInfoVo
	for _, v := range list {
		item := model.NoticeInfoVo{}
		item.Notice = v
		item.SourceName = common.NOTICE_SOURCE_LIST[v.Source]
		result = append(result, item)
	}
	return result, count, nil
}

func (s *noticeService) Add(req *model.NoticeAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Notice
	entity.Title = req.Title
	entity.Content = req.Content
	entity.IsTop = req.IsTop
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.Notice.Insert(entity)
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

func (s *noticeService) Update(req *model.NoticeUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.Notice.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置参数
	info.Title = req.Title
	info.Content = req.Content
	info.IsTop = req.IsTop
	info.Source = req.Source
	info.Status = req.Status
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.Notice.Save(info)
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

func (s *noticeService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除数据
	result, err := dao.Notice.Delete("id in (?)", idsArr)
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
