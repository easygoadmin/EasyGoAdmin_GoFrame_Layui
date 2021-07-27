/**
 * 菜单
 * @auth 鲲鹏
 * @date 2020-05-07
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
            , {
                field: 'type', width: 80, title: '类型', align: 'center', templet(d) {
                    if (d.type == 1) {
                        // 模块
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">模块</span>';
                    } else if (d.type == 2) {
                        // 导航
                        return '<span class="layui-btn layui-btn-danger layui-btn-xs">导航</span>';
                    } else if (d.type == 3) {
                        // 菜单
                        return '<span class="layui-btn layui-btn-warm layui-btn-xs">菜单</span>';
                    } else if (d.type == 4) {
                        // 节点
                        cls = "layui-btn-primary";
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">节点</span>';
                    }
                }
            }
            , {
                field: 'icon',
                width: 80,
                title: '图标',
                align: 'center',
                templet: '<p><i class="layui-icon {{d.icon}}"></i></p>'
            }
            , {field: 'url', width: 150, title: 'URL地址', align: 'center'}
            , {field: 'permission', width: 180, title: '权限标识', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center'}
            , {
                field: 'isPublic', width: 100, title: '是否公共', align: 'center', templet(d) {
                    var cls = "";
                    if (d.isPublic == 1) {
                        // 是
                        cls = "layui-btn-normal";
                    } else if (d.isPublic == 2) {
                        // 否
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.isPublicName + '</span>';
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
        var type = $("#type").val();
        if (type == 3) {
            $(".func").removeClass("layui-hide");
        } else {
            $(".func").addClass("layui-hide");
        }

        /**
         * 权限节点隐藏控制
         */
        form.on('select(type)', function (data) {
            var val = data.value;
            if (val == 3) {
                $(".func").removeClass("layui-hide");
            } else {
                $(".func").addClass("layui-hide");
            }
        });

        /**
         * 提交表单
         */
        form.on('submit(submitForm2)', function (data) {
            if (data.field['type'] == 3) {
                // 获取穿梭组件的选中值
                var funcList = transfer.getData('funcIds'); //获取右侧数据
                // 重组数据并赋值给字段
                var item = [];
                $.each(funcList, function (key, val) {
                    item.push(val['value']);
                });
                data.field['funcIds'] = item.join(",");
            }
            // 提交表单
            common.submitForm(data.field, null, function (res, success) {
                console.log("保存成功回调");
            });
            return false;
        });
    }
});
