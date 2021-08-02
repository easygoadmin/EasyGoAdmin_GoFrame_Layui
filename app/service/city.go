/**
 *
 * @author 摆渡人
 * @since 2021/7/19
 * @File : city
 */
package service

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/utils/convert"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

// 中间件管理服务
var City = new(cityService)

type cityService struct{}

func (s *cityService) GetList(req *model.CityQueryReq) []model.CityInfoVo {
	// 创建查询对象
	query := dao.City.Where("mark=1")
	// 查询条件
	if req != nil {
		// 上级ID
		query = query.Where("pid=?", req.Pid)
		// 城市名称
		if req.Name != "" {
			query = query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	// 排序
	query = query.Order("sort asc")
	// 对象转换
	var list []model.City
	query.Structs(&list)

	var result = make([]model.CityInfoVo, 0)
	// 遍历数据
	for _, v := range list {
		item := model.CityInfoVo{}
		item.City = v
		if v.Level < 3 {
			item.HaveChild = true
		} else {
			item.HaveChild = false
		}
		result = append(result, item)
	}
	return result
}

func (s *cityService) Add(req *model.CityAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.City
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.Level = req.Level
	entity.Citycode = req.Citycode
	entity.PAdcode = req.PAdcode
	entity.Adcode = req.Adcode
	entity.Lng = req.Lng
	entity.Lat = req.Lat
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = gtime.Now()
	entity.Mark = 1

	// 插入记录
	result, err := dao.City.Insert(entity)
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

func (s *cityService) Update(req *model.CityUpdateReq, userId int) (int64, error) {
	// 查询记录
	info, err := dao.City.FindOne("id=?", req.Id)
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, gerror.New("记录不存在")
	}

	// 设置对象属性
	info.Name = req.Name
	info.Pid = req.Pid
	info.Level = req.Level
	info.Citycode = req.Citycode
	info.PAdcode = req.PAdcode
	info.Adcode = req.Adcode
	info.Lng = req.Lng
	info.Lat = req.Lat
	info.Sort = req.Sort
	info.UpdateUser = userId
	info.UpdateTime = gtime.Now()

	// 更新记录
	result, err := dao.City.Save(info)
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

func (s *cityService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := convert.ToInt64Array(ids, ",")
	// 删除记录
	result, err := dao.City.Delete("id in (?)", idsArr)
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

func (s *cityService) GetChilds(cityCode string) ([]*model.City, error) {
	info, err := dao.City.Where("citycode=?", cityCode).FindOne()
	if err != nil {
		return nil, gerror.New("城市不能存在")
	}
	list, err := dao.City.Where("pid=? and mark=1", info.Id).FindAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *cityService) GetCityName(cityCode string, delimiter string) string {
	info, err := dao.City.Where("citycode=?", cityCode).FindOne()
	if err != nil {
		return ""
	}
	// 城市ID
	cityId := info.Id
	// 声明数组
	list := make([]string, 0)
	for {
		if cityId <= 0 {
			// 退出
			break
		}
		// 业务处理
		info, err := dao.City.FindOne("id=?", cityId)
		if err != nil || info == nil {
			break
		}
		// 上级栏目ID
		cityId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组翻转
		return gstr.Implode(delimiter, list)
	}
	return ""
}
