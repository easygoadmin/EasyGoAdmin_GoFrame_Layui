// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
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
