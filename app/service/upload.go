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
 * 文件上传-服务类
 * @author 半城风雨
 * @since 2021/7/23
 * @File : upload
 */
package service

import (
	"easygoadmin/app/utils"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// 中间件管理服务
var Upload = new(uploadService)

type uploadService struct{}

// 上传路径
var uploadPath string

// 初始化
func init() {
	// 文件上传路径
	uploadPath = g.Cfg().GetString("server.ServerRoot") + "/uploads/"
	fmt.Println("初始化配置")
}

// 上传得文件信息
type FileInfo struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	FileUrl  string `json:"fileUrl"`
	FileType string `json:"fileType"`
}

func (s *uploadService) UpdImg(file *ghttp.UploadFile) (FileInfo, error) {
	// 允许上传文件后缀
	fileExt := "jpg,gif,png,bmp,jpeg,JPG"
	// 检查上传文件后缀
	if !checkFileExt(file.Filename, fileExt) {
		return FileInfo{}, gerror.New("上传文件格式不正确，文件后缀只允许为：" + fileExt + "的文件")
	}
	// 允许文件上传最大值
	fileSize := "1M"
	// 检查上传文件大小
	isvalid, err := checkFileSize(file.Size, fileSize)
	if err != nil {
		return FileInfo{}, err
	}
	if !isvalid {
		return FileInfo{}, gerror.New("上传文件大小不得超过：" + fileSize)
	}

	// 临时存储目录
	savePath := utils.TempPath() + "/" + gtime.Now().Format("Ymd")

	// 创建文件夹
	ok := utils.CreateDir(savePath)
	if !ok {
		return FileInfo{}, gerror.New("存储路径创建失败")
	}

	// 上传文件
	fileName, err := file.Save(savePath, true)
	if err != nil {
		return FileInfo{}, err
	}

	// 返回结果
	result := FileInfo{
		FileName: file.Filename,
		FileSize: file.Size,
		FileUrl:  gstr.Replace(savePath, utils.UploadPath(), "") + "/" + fileName,
	}
	return result, nil
}

// 检查文件格式是否合法
func checkFileExt(fileName string, typeString string) bool {
	// 上传文件后缀
	suffix := gstr.SubStrRune(fileName, gstr.PosRRune(fileName, ".")+1, gstr.LenRune(fileName)-1)
	// 允许上传文件后缀
	exts := gstr.Split(typeString, ",")
	// 是否验证通过
	isValid := false
	for _, v := range exts {
		// 对比文件后缀
		if gstr.Equal(suffix, v) {
			isValid = true
			break
		}
	}
	return isValid
}

// 检查上传文件大小
func checkFileSize(fileSize int64, maxSize string) (bool, error) {
	// 匹配上传文件最大值
	match, err := gregex.MatchString(`^([0-9]+)(?i:([a-z]*))$`, maxSize)
	if err != nil {
		return false, err
	}
	if len(match) == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB）")
		return false, err
	}
	var cfSize int64
	switch gstr.ToUpper(match[2]) {
	case "MB", "M":
		cfSize = gconv.Int64(match[1]) * 1024 * 1024
	case "KB", "K":
		cfSize = gconv.Int64(match[1]) * 1024
	case "":
		cfSize = gconv.Int64(match[1])
	}
	if cfSize == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB），最大单位为MB")
		return false, err
	}
	return cfSize >= fileSize, nil
}
