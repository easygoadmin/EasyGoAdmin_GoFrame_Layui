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
 * 栏目管理-服务类
 * @author 半城风雨
 * @since 2021/7/24
 * @File : item_cate
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/convert"
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"reflect"
)

// 中间件管理服务
var ItemCate = new(itemCateService)

type itemCateService struct{}

func (s *itemCateService) GetList(req *model.ItemCateQueryReq) []model.ItemCateInfoVo {
	// 创建查询对象
	query := dao.ItemCate.Where("mark=1")
	// 查询条件
	if req != nil {
		// 栏目名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
		// 站点ID
		if req.ItemId > 0 {
			query = query.Where("item_id=?", req.ItemId)
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.ItemCate
	query.Structs(&list)

	// 数据处理
	var result []model.ItemCateInfoVo
	for _, v := range list {
		item := model.ItemCateInfoVo{}
		item.ItemCate = v
		// 获取栏目
		if v.ItemId > 0 {
			itemInfo, _ := dao.Item.FindOne("id=?", item.ItemId)
			item.ItemName = itemInfo.Name
		}
		// 加入数组
		result = append(result, item)
	}
	return result
}

func (s *itemCateService) Add(req *model.ItemCateAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity model.ItemCate
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.ItemId = req.ItemId
	entity.Pinyin = req.Pinyin
	entity.Code = req.Code
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 封面
	entity.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		entity.Cover = cover
	} else {
		// 没封面
		entity.Cover = ""
	}
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入数据
	result, err := dao.ItemCate.Insert(entity)
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

func (s *itemCateService) Update(req *model.ItemCateUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info, err := dao.ItemCate.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象
	info.Name = req.Name
	info.Pid = req.Pid
	info.ItemId = req.ItemId
	info.Pinyin = req.Pinyin
	info.Code = req.Code
	info.Status = req.Status
	info.Note = req.Note
	info.Sort = req.Sort

	// 封面
	info.IsCover = req.IsCover
	if req.IsCover == 1 {
		// 有封面
		cover, err := utils.SaveImage(req.Cover, "item_cate")
		if err != nil {
			return 0, err
		}
		info.Cover = cover
	} else {
		// 没封面
		info.Cover = ""
	}
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.ItemCate.Save(info)
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

func (s *itemCateService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.ItemCate.Delete("id in (?)", idsArr)
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

// 获取子级菜单
func (s *itemCateService) GetCateTreeList(itemId int, pid int) ([]*model.CateTreeNode, error) {
	var cateNote model.CateTreeNode
	// 创建查询实例
	query := dao.ItemCate.Where("mark=1")
	// 站点ID
	if itemId > 0 {
		query = query.Where("item_id=?", itemId)
	}
	// 返回字段
	query.Fields("id,name,pid")
	// 排序
	query = query.Order("sort asc")
	// 查询所有
	data, err := query.FindAll()
	if err != nil {
		return nil, errors.New("系统错误")
	}
	makeCateTree(data, &cateNote)
	return cateNote.Children, nil
}

//递归生成分类列表
func makeCateTree(cate []*model.ItemCate, tn *model.CateTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &model.CateTreeNode{}
			child.ItemCate = *c
			tn.Children = append(tn.Children, child)
			makeCateTree(cate, child)
		}
	}
}

// 数据源转换
func (s *itemCateService) MakeList(data []*model.CateTreeNode) []map[string]string {
	cateList := make([]map[string]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			item := map[string]string{}
			item["id"] = gconv.String(val.Id)
			item["name"] = val.Name
			cateList = append(cateList, item)

			// 二级栏目
			for _, v := range val.Children {
				item2 := map[string]string{}
				item2["id"] = gconv.String(v.Id)
				item2["name"] = "|--" + v.Name
				cateList = append(cateList, item2)

				// 三级栏目
				for _, vt := range v.Children {
					item3 := map[string]string{}
					item3["id"] = gconv.String(vt.Id)
					item3["name"] = "|--|--" + vt.Name
					cateList = append(cateList, item3)
				}
			}
		}
	}
	return cateList
}

func (s *itemCateService) GetCateName(cateId int, delimiter string) string {
	// 声明数组
	list := make([]string, 0)
	for {
		if cateId <= 0 {
			// 退出
			break
		}
		// 业务处理
		info, err := dao.ItemCate.FindOne("id=?", cateId)
		if err != nil || info == nil {
			break
		}
		// 上级栏目ID
		cateId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		return gstr.Implode(delimiter, list)
	}
	return ""
}
