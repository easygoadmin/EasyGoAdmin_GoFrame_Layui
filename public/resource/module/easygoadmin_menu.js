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
 * 菜单管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'common', 'form', 'transfer'], function () {

    //声明变量
    var func = layui.func
        , common = layui.common
        , form = layui.form
        , transfer = layui.transfer
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '菜单名称', align: 'left'}
            , {field: 'type', width: 80, title: '类型', align: 'center', templet(d) {
                    if (d.type == 0) {
                        // 菜单
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">菜单</span>';
                    } else if (d.type == 1) {
                        // 节点
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">节点</span>';
                    }
                }
            }
            , { field: 'icon', width: 80, title: '图标', align: 'center', templet: '<p><i class="layui-icon {{d.icon}}"></i></p>'}
            , {field: 'url', width: 150, title: 'URL地址', align: 'center'}
            , {field: 'permission', width: 180, title: '权限标识', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                    if (d.status == 1) {
                        // 在用
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">在用</span>';
                    } else {
                        // 停用
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">停用</span>';
                    }
                }
            }
            , {field: 'target', width: 100, title: '是否公共', align: 'center', templet(d) {
                    if (d.target == 1) {
                        // 内部打开
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">内部打开</span>';
                    } else if (d.target == 2) {
                        // 外部打开
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">外部打开</span>';
                    }
                }
            }
            , {field: 'sort', width: 90, title: '显示顺序', align: 'center'}
            , {fixed: 'right', width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("菜单");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });

    } else {

        // 初始化
        var type = $("#type").val()
        if (type == 0) {
            $(".func").removeClass("layui-hide");
        } else {
            $(".func").addClass("layui-hide");
        }

        // 菜单类型选择事件
        form.on('select(type)', function (data) {
            var val = data.value;
            if (val == 0) {
                $(".func").removeClass("layui-hide");
            } else {
                $(".func").addClass("layui-hide");
            }
        });

        /**
         * 提交表单
         */
        form.on('submit(submitForm2)', function (data) {
            if (data.field['type'] == 0) {
                // 获取穿梭组件的选中值
                var funcList = transfer.getData('func'); //获取右侧数据
                // 重组数据并赋值给字段
                var item = [];
                $.each(funcList, function (key, val) {
                    item.push(val['value']);
                });
                data.field['func'] = item.join(",");
            }
            // 提交表单
            common.submitForm(data.field, null, function (res, success) {
                console.log("保存成功回调");
            });
            return false;
        });
    }
});
