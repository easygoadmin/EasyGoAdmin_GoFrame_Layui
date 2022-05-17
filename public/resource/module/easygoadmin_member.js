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
 * 会员管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'form'], function () {
    var func = layui.func;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'username', width: 120, title: '用户名', align: 'center'}
            , {field: 'nickname', width: 120, title: '用户昵称', align: 'center'}
            , {
                field: 'gender', width: 60, title: '性别', align: 'center', templet(d) {
                    var cls = "";
                    if (d.gender == 1) {
                        // 男
                        cls = "layui-btn-normal";
                    } else if (d.gender == 2) {
                        // 女
                        cls = "layui-btn-warm";
                    } else if (d.gender == 3) {
                        // 保密
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.genderName + '</span>';
                }
            }
            , {field: 'avatar', width: 90, title: '用户头像', align: 'center', templet: function (d) {
                    var avatarStr = "";
                    if (d.avatar) {
                        avatarStr = '<a href="' + d.avatar + '" target="_blank"><img src="' + d.avatar + '" height="26" /></a>';
                    }
                    return avatarStr;
                }
            }
            , {
                field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'cityName', width: 250, title: '所在地区', align: 'center'}
            , {field: 'device', width: 100, title: '设备类型', align: 'center', templet(d) {
                    var cls = "";
                    if (d.device == 1) {
                        // 苹果
                        cls = "layui-btn-normal";
                    } else if (d.device == 2) {
                        // 安卓
                        cls = "layui-btn-warm";
                    } else if (d.device == 3) {
                        // WAP站
                        cls = "layui-btn-danger";
                    } else if (d.device == 4) {
                        // PC站
                        cls = "layui-btn-primary";
                    } else if (d.device == 5) {
                        // 微信小程序
                        cls = "layui-btn-disabled";
                    }

                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.deviceName+'</span>';
                }}
            , {field: 'source', width: 100, title: '用户来源', align: 'center', templet(d) {
                    var cls = "";
                    if (d.source == 1) {
                        // 注册会员
                        cls = "layui-btn-normal";
                    } else if (d.source == 2) {
                        // 马甲会员
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.sourceName+'</span>';
                }}
            , {field: 'createTime', width: 180, title: '注册时间', align: 'center',templet:'<div>{{ layui.util.toDateString(d.create_time, "yyyy-MM-dd HH:mm:ss") }}</div>'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("会员用户");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
