/**
 * 常用方法封装【基础类】
 * @author 牧羊人
 * @date 2018/7/13
 */
layui.define(['form', 'layer', 'laydate', 'upload', 'element', 'base'], function (exports) {
    "use strict";

    // 变量声明
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        laydate = layui.laydate,
        upload = layui.upload,
        element = layui.element,
        base = layui.base,
        $ = layui.$;

    /**
     * 自定义模块接口对象
     */
    var active = {
        /**
         * 添加、编辑公共函数
         */
        edit: function (title, id = 0, width = 0, height = 0, param = [], callback = null, isClose = false) {

            // 窗口标题
            var titleStr = id > 0 ? "修改" : "新增";
            if (base.isEmpty(title)) {
                titleStr += '内容';
            } else {
                titleStr += title;
            }

            // URL逻辑处理
            var url = cUrl + "/edit?id=" + id;
            if (Array.isArray(param)) {
                for (var i in param) {
                    url += "&" + param[i];
                }
            }
            // 调用内部方法
            active.showWin(titleStr, url, width, height, param, 2, [], function (index, type) {
                if (callback) {
                    callback(index, type);
                }
            }, isClose);

        },
        /**
         * 数据详情函数
         */
        detail: function (title, id, width = 0, height = 0, isClose = false) {
            // 调用内部方法
            var url = cUrl + "/detail?id=" + id;
            active.showWin(title + "详情", url, width, height, [], 2, [], null, isClose);
        },
        /**
         * 重置缓存
         */
        cache: function (id) {
            var url = cUrl + "/cache";
            active.ajaxPost(url, {"id": id}, function (data, res) {
            });
        },
        /**
         * 一键复制
         */
        copy: function (title, id, width = 0, height = 0) {
            var url = cUrl + "/copy?id=" + id;
            active.showWin(title + "复制", url, width, height);
        },
        /**
         * 删除单条数据函数
         */
        delete: function (id, callback = null) {
            layer.confirm('您确定要删除吗？删除后将无法恢复！', {
                icon: 3,
                skin: 'layer-ext-moon',
                btn: ['确认', '取消'] //按钮
            }, function (index) {

                // 调用内部方法
                var url = cUrl + "/delete?ids=" + id;
                console.log(url);
                active.ajaxPost(url, {}, function (data, flag) {
                    if (callback) {
                        // 关闭弹窗
                        layer.close(index);
                        // 回调
                        callback(data, flag);
                    }
                }, '正在删除。。。');

            });

        },
        /**
         * 批量操作方法
         */
        batchFunc: function (option, callback = null) {
            // 基础参数
            var url = option.url,
                title = option.title,
                form = option.form || '',
                confirm = option.confirm || false,
                show_tips = option.show_tips || '处理中...',
                item = option.data || [],
                param = option.param || [],
                type = option.type || 'POST';

            if (title != "导出数据" && item.length == 0) {
                layer.msg("请选择数据", {icon: 5});
                return false;
            }

            // 选择数据ID
            var ids = [];
            for (var i in item) {
                ids.push(item[i].id);
            }
            // 选择数据ID字符串（逗号‘,’分隔）
            var ids_str = ids.join(",");

            var data = {};
            data['ids'] = ids_str;

            // 自定义参数解析
            if (Array.isArray(param)) {
                for (var i in param) {
                    var subItem = param[i].split('=');
                    data[subItem[0]] = subItem[1];
                }
            }
            console.log(data)

            if (confirm) {
                // 弹出确认
                layer.confirm('您确定要【' + title + '】选中的数据吗？', {icon: 3, title: '提示信息'}, function (index) {
                    if (type == "POST") {
                        active.ajaxPost(url, data, callback, show_tips);
                    } else {
                        active.ajaxGet(url + "/" + ids_str, {}, callback, show_tips);
                    }
                });
            } else {
                // 直接请求
                if (type == "POST") {
                    active.ajaxPost(url, data, callback, show_tips);
                } else {
                    active.ajaxGet(url + "/" + ids_str, {}, callback, show_tips);
                }
            }
        },
        /**
         * 表单验证函数
         */
        verify: function () {
            form.verify({
                number: [/^[0-9]*$/, '请输入数字']
                , username: function (value, item) {
                    // 特殊字符验证
                    if (!new RegExp("^[a-zA-Z0-9_\u4e00-\u9fa5\\s·]+$").test(value)) {
                        return title + '不能含有特殊字符';
                    }
                    // 下划线验证
                    if (/(^\_)|(\__)|(\_+$)/.test(value)) {
                        return title + '首尾不能出现下划线\'_\'';
                    }
                    // 数字验证
                    if (/^\d+\d+\d$/.test(value)) {
                        return title + '不能全为数字';
                    }
                }
                // 数组的两个值分别代表：[正则匹配、匹配不符时的提示文字]
                , pass: [/^[\S]{6,12}$/, '密码必须6到12位，且不能出现空格']
            });
        },
        /**
         * 表单提交函数
         */
        submitForm: function (data, url = null, callback = null, isClose = true) {
            // 带[]中括号下标的字段特殊处理
            var nameArr = [];
            var itemArr = [];
            var param = data;
            $.each(param, function (key, val) {
                // 正则验证字段是否存在中括号[]
                var regex = /\[|\]|【|】/g
                if (!regex.test(key)) {
                    return;
                }

                // 处理带括号[]的字段
                var regex1 = /\[(.+?)\]/g;   // [] 中括号及内容

                // 获取括号及括号内容
                var content = key.match(regex1);

                // 获取括号内容值
                var regex2 = "\\[(.+?)\\]";
                var item = key.match(regex2);
                val = item[1];

                // 获取字段名
                var name = key.replace(content, "");
                // 字段名临时存储
                if ($.inArray(name, nameArr) < 0) {
                    nameArr.push(name);
                }

                // 字段名数组初始化
                if (!itemArr[name]) {
                    itemArr[name] = [];
                }
                itemArr[name].push(val);
            });
            // 遍历数组
            $.each(nameArr, function (i, name) {
                var item = [];
                $.each(itemArr[name], function (key, val) {
                    item.push(val);
                    // 移除指定元素
                    delete param[name + "[" + val + "]"];
                });
                param[name] = item.join(",");
            });

            // 请求地址
            if (url == null) {
                var url = cUrl;
                var action = $("form").attr("action");
                if (!base.isEmpty(action)) {
                    // 自定义网络请求地址
                    url = action;
                } else if (data.id != null) {
                    // 根据常规表单提交判断是新增还是编辑
                    if (data.id == 0) {
                        url += "/add";
                    } else if (data.id > 0) {
                        url += "/update";
                    }
                }
            }
            active.ajaxPost(url, param, function (res, success) {
                if (success) {
                    // 延迟0.5秒
                    if (isClose) {
                        setTimeout(function () {
                            // 关闭窗体
                            // layer.closeAll("iframe");
                            // console.log("关闭所有弹窗")
                            // 获取窗口索引
                            var index = parent.layer.getFrameIndex(window.name)
                            // 关闭layer弹出层
                            parent.layer.close(index)
                            // 刷新父页面
                            //parent.location.reload();
                        }, 500);
                    }

                    // 数据回调
                    if (callback) {
                        callback(res, success);
                    }
                    return false;
                } else {
                    // 网络请求异常处理
                }
            });
        },
        /**
         * 搜索函数
         */
        searchForm: function (table, data, tableList = 'tableList') {
            // 执行重载
            table.reload(tableList, {
                page: {
                    curr: 1
                },
                where: data.field
            });
        },
        /**
         * 初始化日期组件函数
         */
        initDate: function (item, callback = null) {
            if (Array.isArray(item)) {
                for (var i in item) {
                    var subItem = item[i].split('|');
                    if (subItem[2]) {
                        var param = subItem[2].split(',');
                    }

                    // 日期组件数据重组
                    var options = {};
                    options.elem = "#" + subItem[0];
                    options.type = subItem[1];
                    options.theme = 'molv';// 主题颜色[molv,#393D49,grid]
                    options.range = subItem[3] === "true" ? true : subItem[3];// 开启左右面板
                    options.calendar = true;// 是否显示公历节日
                    options.show = false;// 默认显示
                    options.position = 'absolute';// [fixed,absolute,static]
                    options.trigger = 'click';// 定义鼠标悬停时弹出控件[click,mouseover]
                    options.btns = ['clear', 'now', 'confirm'];// 工具按钮 默认值['clear', 'now', 'confirm']
                    options.mark = {'0-06-25': "生日", '0-12-31': "跨年"};// 自定义标注重要日子
                    // 控件在打开时触发，回调返回一个参数
                    options.ready = function (date) {
                        // console.log("组件面板打开：" + date);
                    }
                    // 日期时间被切换后的回调
                    options.change = function (value, date, endDate) {
                        // console.log(value); // 得到日期生成的值，如：2017-08-18
                        // console.log(date); // 得到日期时间对象：{year: 2017, month: 8, date: 18, hours: 0, minutes: 0, seconds: 0}
                        // console.log(endDate); // 得结束的日期时间对象，开启范围选择（range: true）才会返回。对象成员同上。
                    }
                    // 控件选择完毕后的回调
                    options.done = function (value, date, endDate) {
                        if (callback) {
                            callback(value, date);
                        }
                    }
                    if (param) {
                        // 最小值
                        var minV = param[0];
                        if (minV) {
                            var isNum = !isNaN(minV);
                            if (isNum) {
                                // 数字
                                options.min = parseInt(minV);
                            } else {
                                // 非数字
                                options.min = minV;
                            }
                        }
                        // 最大值
                        var maxV = param[1];
                        if (maxV) {
                            var isNum2 = !isNaN(maxV);
                            if (isNum2) {
                                // 数字
                                options.max = parseInt(maxV);
                            } else {
                                // 非数字
                                options.max = maxV;
                            }
                        }
                    }

                    // 日期选择组件
                    laydate.render(options);
                }
            }
        },
        /**
         * 弹出窗体函数
         */
        showWin: function (title, url, width = 0, height = 0, param = [], type = 2, btn = [], callback = null, isClose = false) {
            var index = layui.layer.open({
                title: title,
                type: type,
                area: [width + "px", height + "px"],
                content: url,
                // closeBtn: false,
                shadeClose: isClose,// 点击遮罩关闭
                shade: 0.4,
                // maxmin: true, // 开启最大化最小化按钮
                // skin: 'layui-layer-rim', // 加上边框
                // skin: 'layui-layer-molv', // 加上边框
                skin: 'layui-layer-admin',
                // btn: btn,
                // btnAlign: 'c',
                success: function (layero, index) {

                    // 窗体传值【支持多值传递】
                    if (Array.isArray(param)) {
                        for (var i in param) {
                            var item = param[i].split('=');
                            // console.log("传值：" + item[0] + "," + item[1]);
                            var body = layui.layer.getChildFrame('body', index);
                            body.find("#" + item[0]).val(item[1]);
                        }
                    }

                    // 回调函数
                    if (callback) {
                        callback(index, 1);
                    }

                    // // 延迟0.5秒
                    // setTimeout(function () {
                    //     layer.tips('点击此处返回列表', '.layui-layer-setwin .layui-layer-close', {
                    //         tips: 3
                    //     });
                    // }, 500);

                },
                end: function () {
                    callback(index, 2);
                }
            });

            if (width == 0) {
                // 全屏设置
                layui.layer.full(index);
                $(window).on("resize", function () {
                    layui.layer.full(index);
                });
            }

        },
        /**
         * 网络请求函数(POST)
         */
        ajaxPost: function (url, data, callback = null, msg = '处理中,请稍后...') {
            var index = null;
            $.ajax({
                type: "POST",
                url: url,
                data: JSON.stringify(data),
                contentType: "application/json",
                dataType: "json",
                beforeSend: function () {
                    index = layer.msg(msg, {
                        icon: 16
                        , shade: 0.01
                        , time: 0
                    });
                },
                success: function (res) {
                    if (res.code == 0) {
                        //0.5秒后关闭
                        layer.msg(res.msg, {icon: 1, time: 500}, function () {
                            layer.close(index);
                            if (callback) {
                                callback(res, true);
                            }
                        });
                    } else {
                        layer.close(index);
                        layer.msg(res.msg, {icon: 5});
                        return false;
                    }
                },
                error: function () {
                    layer.close(index);
                    layer.msg("AJAX请求异常");
                    if (callback) {
                        callback(null, false);
                    }
                }
            });
        },
        /**
         * 网络请求函数(POST)
         */
        ajaxGet: function (url, data, callback = null, msg = '处理中,请稍后...') {
            var index = null;
            $.ajax({
                type: "GET",
                url: url,
                data: data,
                contentType: "application/json",
                dataType: "json",
                beforeSend: function () {
                    index = layer.msg(msg, {
                        icon: 16
                        , shade: 0.01
                        , time: 0
                    });
                },
                success: function (res) {
                    if (res.code == 0) {
                        //0.5秒后关闭
                        layer.msg(res.msg, {icon: 1, time: 500}, function () {
                            layer.close(index);
                            if (callback) {
                                callback(res, true);
                            }
                        });
                    } else {
                        layer.msg(res.msg, {icon: 5});
                        return false;
                    }
                },
                error: function () {
                    layer.msg("AJAX请求异常");
                    if (callback) {
                        callback(null, false);
                    }
                }
            });
        },
        /**
         * 开关
         */
        formSwitch: function (name, url = '', callback = null) {
            form.on('switch(' + name + ')', function (obj) {
                // 开关的值
                var value = this.checked ? '1' : '2';

                // URL自定义
                if (base.isEmpty(url)) {
                    url = cUrl + "/set" + name.substring(0, 1).toUpperCase() + name.substring(1);
                }

                // JSON数据
                var json_data = {};
                json_data['id'] = this.value;
                json_data[name] = value;
                // JSON字符串
                var json_str = JSON.stringify(json_data);
                // JSON数据
                var json = JSON.parse(json_str);
                // 发起POST请求
                active.ajaxPost(url, json_data, function (data, res) {
                    if (callback) {
                        callback(data, res);
                    }
                });

            });
        },
        /**
         * 上传文件
         */
        uploadFile: function (elem_id, callback = null, url = '', exts = 'xls|xlsx', size = 10240, data = {}) {
            if (base.isEmpty(url)) {
                url = cUrl + "/uploadFile";
            }
            upload.render({
                elem: '#' + elem_id
                , url: url
                , auto: false
                , exts: exts
                , accept: 'file' // 允许上传的文件类型
                , size: size // 最大允许上传的文件大小
                , method: 'post' // 可选项。HTTP类型，默认post
                , data: data // 可选项。额外的参数，如：{id: 123, abc: 'xxx'}
                , before: function (obj) {
                    // 预读本地文件
                    layer.msg('上传并处理中。。。', {
                        icon: 16
                        , shade: 0.01
                        , time: 0
                    });
                }
                , done: function (res) {
                    // 上传完毕回调

                    // 关闭所有弹窗
                    layer.closeAll();

                    // 上传成功
                    if (res.code == 0) {
                        layer.alert(res.msg, {
                            title: "上传反馈"
                            , skin: 'layui-layer-molv' //样式类名  自定义样式
                            , closeBtn: 1    // 是否显示关闭按钮
                            , anim: 0 //动画类型
                            , btn: ['确定', '取消'] //按钮
                            , icon: 6    // icon
                            , yes: function () {
                                // 回调
                                if (callback) {
                                    callback(res, true);
                                }
                            }
                            , btn2: function () {
                            }
                        });
                    } else {
                        layer.msg(res.msg, {icon: 5});
                    }
                    return false;
                }
                , error: function () {
                    // 请求异常回调
                    return layer.msg('数据请求异常');
                }
            });
        },
    };

    /**
     * 输入自定义模块(此模块接口是对象)
     */
    exports('common', active);
});