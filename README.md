<p align="center">
    <img src="https://images.gitee.com/uploads/images/2021/0803/150932_211843b7_478496.png" width="620" height="120" />
</p>

## 项目介绍
一款Golang语言基于GoFrame、Layui、MySQL等框架精心打造的一款模块化、高性能、企业级的敏捷开发框架，本着简化开发、提升开发效率的初衷触发，框架自研了一套个性化的组件，实现了可插拔的组件式开发方式：单图上传、多图上传、下拉选择、开关按钮、单选按钮、多选按钮、图片裁剪等等一系列个性化、轻量级的组件，是一款真正意义上实现组件化开发的敏捷开发框架。

## 🍻 项目特点

+ 模块化、松耦合
+ 模块丰富、开箱即用
+ 简洁易用、快速接入
+ 文档详尽、易于维护
+ 自顶向下、体系化设计
+ 统一框架、统一组件、降低选择成本
+ 开发规范、设计模式、代码分层模型
+ 强大便捷的开发工具链
+ 完善的本地中文化支持
+ 设计为团队及企业使用

## 🍪 内置模块
+ 用户管理：用于维护管理系统的用户，常规信息的维护与账号设置。
+ 角色管理：角色菜单管理与权限分配、设置角色所拥有的菜单权限。
+ 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
+ 职级管理：主要管理用户的职级。
+ 岗位管理：主要管理用户担任职务。
+ 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
+ 字典管理：对系统中常用的较为固定的数据进行统一维护。
+ 配置管理：对系统的常规配置信息进行维护，网站配置管理功能进行统一维护。
+ 通知公告：系统通知公告信息发布维护。
+ 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
+ 登录日志：系统登录日志记录查询包含登录异常。
+ 代码生成：一键生成模块CRUD的功能，包括后端Go和前端HTML、JS等相关代码。
+ 案例演示：常规代码生成器一键生成后的演示案例。

## 👷 开发者信息
* 系统名称：EasyGoAdmin敏捷开发框架GoFrame+Layui版本
* 作者：半城风雨
* 作者QQ：[1175401194](http://wpa.qq.com/msgrd?v=3&amp;uin=1175401194&amp;site=qq&amp;menu=yes)
* 官网网址：[http://www.easygoadmin.vip/](http://www.easygoadmin.vip/)
* 文档网址：[http://docs.goframe.layui.easygoadmin.vip/](http://docs.goframe.layui.easygoadmin.vip/)

## 🎨 系统演示

+ 演示地址：http://goframe.layui.rxthink.cn/login

账号 | 密码| 操作权限
---|---|---
admin | 123456| 演示环境无法进行修改删除操作

## 👷 技术支持

[技术支持QQ：1175401194](http://wpa.qq.com/msgrd?v=3&amp;uin=1175401194&amp;site=qq&amp;menu=yes)

## 📌 版本说明
版本名称 | 版本说明 | 版本地址
---|---|---
GoFrame+Layui混编版 | 采用GoFrame、Layui等框架研发 | https://gitee.com/easygoadmin/EasyGoAdmin_GoFrame_Layui
Beego+Layui混编版 | 采用Beego、Layui等框架研发 | https://gitee.com/easygoadmin/EasyGoAdmin_Beego_Layui
Gin+Layui混编版 | 采用Gin、Layui等框架研发 | https://gitee.com/easygoadmin/EasyGoAdmin_Gin_Layui
Iris+Layui混编版 | 采用Iris、Layui等框架研发 | https://gitee.com/easygoadmin/EasyGoAdmin_Iris_Layui
GoFrame+EleVue前后端分离版 | 采用GoFrame、Vue、ElementUI等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_GoFrame_EleVue
Beego+EleVue前后端分离版 | 采用Beego、Vue、ElementUI等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Beego_EleVue
Gin+EleVue前后端分离版 | 采用Gin、Vue、ElementUI等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Gin_EleVue
Iris+EleVue前后端分离版 | 采用Iris、Vue、ElementUI等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Iris_EleVue
GoFrame+AntdVue前后端分离版 | 采用GoFrame、Vue、AntDesign等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_GoFrame_AntdVue
Beego+AntdVue前后端分离版 | 采用Beego、Vue、AntDesign等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Beego_AntdVue
Gin+AntdVue前后端分离版 | 采用Gin、Vue、AntDesign等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Gin_AntdVue
Iris+AntdVue前后端分离版 | 采用Iris、Vue、AntDesign等框架研发前后端分离版本 | https://gitee.com/easygoadmin/EasyGoAdmin_Iris_AntdVue

## 🍪 项目结构

```
├── app             // 应用目录
│   ├── controller  // 控制器
│   ├── dao         // DAO层
│   ├── model       // 模型层
│   └── service     // 服务层
│   └── utils       // 系统工具
│   └── widget      // 核心组件
├── boot
├── config          // 系统配置
├── docker
├── document        // 文档目录
├── i18n            // 国际化
├── library         // 类库
├── packed
├── public          // 资源目录
├── router          // 路由
├── template        // 模板
├── Dockerfile
├── go.mod
└── main.go
```

## 📚 核心组件

+ 单图上传组件
```
{{upload_image "avatar|头像|90x90|建议上传尺寸450x450|450x450" .info.Avatar "" 0}}
```
+ 多图上传组件
```
{{album "avatar|图集|90x90|20|建议上传尺寸450x450" .info.Avatar "" 0}}
```
+ 下拉选择组件
```
{{select "gender|1|性别|name|id" "1=男,2=女,3=保密" .info.Gender}}
```
+ 单选按钮组件
```
{{radio "gender|name|id" "1=男,2=女,3=保密" .info.Gender}}
```
+ 复选框组件
```
{{checkbox "role_ids|name|id" .roleList .info.RoleIds}}
```
+ 城市选择组件
```
{{city .info.DistrictCode 3 1}}
```
+ 开关组件
```
{{switch "status" "在用|禁用" .info.Status}}
```
+ 日期组件
```
{{date "birthday|1|出生日期|date" .info.Birthday}}
```
+ 图标组件
```
{{icon "icon" .info.Icon}}
```
+ 穿梭组件
```
{{transfer "func|0|全部节点,已赋予节点|name|id|220x350" "1=列表,5=添加,10=修改,15=删除,20=详情,25=状态,30=批量删除,35=添加子级,40=全部展开,45=全部折叠" .funcList}}
```

## 🔧 模块展示

![效果图](./public/uploads/demo/1.png)

![效果图](./public/uploads/demo/2.png)

![效果图](./public/uploads/demo/3.png)

![效果图](./public/uploads/demo/4.png)

![效果图](./public/uploads/demo/5.png)

![效果图](./public/uploads/demo/6.png)

![效果图](./public/uploads/demo/7.png)

![效果图](./public/uploads/demo/8.png)

![效果图](./public/uploads/demo/9.png)

![效果图](./public/uploads/demo/10.png)

![效果图](./public/uploads/demo/11.png)

![效果图](./public/uploads/demo/12.png)

![效果图](./public/uploads/demo/13.png)

![效果图](./public/uploads/demo/14.png)

![效果图](./public/uploads/demo/15.png)

![效果图](./public/uploads/demo/16.png)

![效果图](./public/uploads/demo/17.png)

![效果图](./public/uploads/demo/18.png)

![效果图](./public/uploads/demo/19.png)

![效果图](./public/uploads/demo/20.png)

![效果图](./public/uploads/demo/21.png)

![效果图](./public/uploads/demo/22.png)

![效果图](./public/uploads/demo/23.png)

![效果图](./public/uploads/demo/24.png)

![效果图](./public/uploads/demo/25.png)


## 🍻 贡献者名单

@半城风雨

## 🍻  安全&缺陷
如果你发现了一个安全漏洞或缺陷，请发送邮件到 1175401194@qq.com,所有的安全漏洞都将及时得到解决。

## ✨  特别鸣谢
感谢[GoFrame](https://goframe.org/#all-updates)、[Layui](http://www.layui.com)等优秀开源项目。

## 📚 版权信息

商业版使用需授权，未授权禁止恶意传播和用于商业用途，否则将追究相关人的法律责任。

本项目包含的第三方源码和二进制文件之版权信息另行标注。

版权所有Copyright © 2021 easygoadmin.vip (http://www.easygoadmin.vip)

All rights reserved。

更多细节参阅 [LICENSE.txt](LICENSE.txt)