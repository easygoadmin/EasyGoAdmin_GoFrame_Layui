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
 * 工具类库
 * @author 半城风雨
 * @since 2021/7/23
 * @File : utils
 */
package utils

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"log"
	"os"
	"regexp"
	"strings"
)

// 附件目录
func UploadPath() string {
	// 附件存储路径
	upload_dir := g.Cfg().GetString("easygoadmin.upload_dir")
	if upload_dir != "" {
		return upload_dir
	} else {
		// 获取项目根目录
		curDir, _ := os.Getwd()
		return curDir + "/public/uploads"
	}
}

// 临时目录
func TempPath() string {
	return UploadPath() + "/temp"
}

// 图片存放目录
func ImagePath() string {
	return UploadPath() + "/images"
}

// 文件目录(非图片目录)
func FilePath() string {
	return UploadPath() + "/file"
}

// 创建文件夹并设置权限
func CreateDir(path string) bool {
	// 判断文件夹是否存在
	if IsExist(path) {
		return true
	}
	// 创建文件夹
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// 判断文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		if os.IsExist(err) {
			// 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

// 判断文件是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func ImgUrl() string {
	return g.Cfg().GetString("easygoadmin.image_url")
}

// 获取文件地址
func GetImageUrl(path string) string {
	return ImgUrl() + path
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func SaveImage(url string, dirname string) (string, error) {
	// 判断文件地址是否为空
	if gstr.Equal(url, "") {
		return "", gerror.New("文件地址不能为空")
	}

	// 判断是否本站图片
	if gstr.Contains(url, ImgUrl()) {
		// 本站图片

		// 是否临时图片
		if gstr.Contains(url, "temp") {
			// 临时图片

			// 创建目录
			dirPath := ImagePath() + "/" + dirname + "/" + gtime.Now().Format("Ymd")
			if !CreateDir(dirPath) {
				return "", gerror.New("文件目录创建失败")
			}
			// 原始图片地址
			oldPath := gstr.Replace(url, ImgUrl(), UploadPath())
			// 目标目录地址
			newPath := ImagePath() + "/" + dirname + gstr.Replace(url, ImgUrl()+"/temp", "")
			// 移动文件
			os.Rename(oldPath, newPath)
			return gstr.Replace(newPath, UploadPath(), ""), nil
		} else {
			// 非临时图片
			path := gstr.Replace(url, ImgUrl(), "")
			return path, nil
		}
	} else {
		// 远程图片
		// TODO...
	}
	return "", gerror.New("保存文件异常")
}

// 处理富文本
func SaveImageContent(content string, title string, dirname string) string {
	str := `<img src="(?s:(.*?))"`
	//解析、编译正则
	ret := regexp.MustCompile(str)
	// 提取图片信息
	alls := ret.FindAllStringSubmatch(content, -1)
	// 遍历图片数据
	for _, v := range alls {
		// 获取图片地址
		item := v[1]
		if item == "" {
			continue
		}
		// 保存图片至正式目录
		image, _ := SaveImage(item, dirname)
		if image != "" {
			content = strings.ReplaceAll(content, item, "[IMG_URL]"+image)
		}
	}
	// 设置ALT标题
	if strings.Contains(content, "alt=\"\"") && title != "" {
		content = strings.ReplaceAll(content, "alt=\"\"", "alt=\""+title+"\"")
	}
	return content
}

func Md5(password string) (string, error) {
	// 第一次MD5加密
	password, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	password2, err := gmd5.Encrypt(password)
	if err != nil {
		return "", err
	}
	return password2, nil
}

// 判断元素是否在数组中
func InArray(value string, array []interface{}) bool {
	for _, v := range array {
		if gconv.String(v) == value {
			return true
		}
	}
	return false
}

// 登录用户ID
func Uid(session *ghttp.Session) int {
	return session.GetInt("userId")
}

// 判断用户登录状态
func IsLogin(session *ghttp.Session) bool {
	// 获取用户ID
	userId := session.GetInt("userId")
	return userId > 0
}

// 获取数据库表
func GetDatabase() (string, error) {
	// 获取数据库连接
	link := g.Cfg().GetString("database.link")
	if link == "" {
		return "", gerror.New("数据库配置读取错误")
	}
	// 分裂字符串
	linkArr := strings.Split(link, "/")
	return linkArr[1], nil
}

// 调试模式
func AppDebug() bool {
	return g.Cfg().GetString("easygoadmin.app_debug") == "true"
}

// 系统版本号
func GetVersion() string {
	return g.Cfg().GetString("easygoadmin.version")
}

// 数组反转
func Reverse(arr *[]string) {
	length := len(*arr)
	var temp string
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}
