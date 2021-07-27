/**
 * 常用方法封装【方法调用类】
 * @author 鲲鹏
 * @date 2018/7/12
 */
layui.define(['form', 'layer', 'table', 'common', 'treeTable'], function (exports) {
    "use strict";

    var form = layui.form,
        table = layui.table,
        layer = layui.layer,
        common = layui.common,
        treeTable = layui.treeTable,
        $ = layui.$;

    /**
     * 声明全局变量
     */
    var _tableIns,
        _tableName,
        _callback,
        _title,
        _width = 0,
        _height = 0,
        _isClose = false,
        _isDbclick;

    /**
     * 自定义模块接口对象
     */
    var func = {

        /**
         * TABLE列表函数
         */
        tableIns: function (cols, tableName, callback = null, url = '', tableSort = false) {
            _tableName = tableName;
            _callback = callback;

            // 初始化网络请求URL
            if (!url || url == '') {
                url = cUrl + "/list";
            }

            // 初始化网络请求参数
            var param = $("#param").val();
            if (param) {
                param = JSON.parse(param);
                if ($.isArray(param)) {
                    for (var i in param) {
                        if (url.indexOf("?") >= 0) {
                            // 包含?
                            url += "&" + param[i];
                        } else {
                            // 不包含?
                            url += "?" + param[i];
                        }
                    }
                }
            }

            // 初始化TABLE组件
            _tableIns = table.render({
                elem: "#" + _tableName
                , url: url
                // , toolbar: '#toolbar_header'
                // , title: '用户数据表'
                // , totalRow: true
                , method: 'post'
                , cellMinWidth: 150
                // , page: true
                , page: {
                    // 限定条数   总数、计数  上一页     页     下一页    到第几页、跳
                    layout: ['refresh', 'prev', 'page', 'next', 'skip', 'count', 'limit'] //自定义分页布局
                    , curr: 1
                    , groups: 10 //显示 连续页码
                    , first: '首页'
                    , last: '尾页'
                }
                // //初始排序
                // , initSort: {
                //     field: 'id', //排序字段，对应 cols 设定的各字段名
                //     type: 'desc' //排序方式  asc: 升序、desc: 降序、null: 默认排序
                // }
                , height: "full-100"
                , limit: 20
                , limits: [20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200, 1000]
                , even: true // 开启隔行背景
                , cols: [cols]
                , loading: true
                , done: function (res, curr, count) {
                    // 新增监听table行双击事件
                    if (_isDbclick) {
                        var tbody = $('.layui-table-body').find("table").find("tbody");
                        var tr = tbody.children("tr");
                        tr.on('dblclick', function () {
                            var index = tbody.find(".layui-table-hover").data('index');
                            var obj = res.data[index];
                            common.edit(_title, obj.id, _width, _height);
                        });
                    }

                }
            });

            // 监听头工具栏事件
            table.on("toolbar(" + _tableName + ")", function (obj) {
                var checkStatus = table.checkStatus(obj.config.id);
                switch (obj.event) {
                    case 'getCheckData':
                        var data = checkStatus.data;
                        layer.alert(JSON.stringify(data));
                        break;
                    case 'getCheckLength':
                        var data = checkStatus.data;
                        layer.msg('选中了：' + data.length + ' 个');
                        break;
                    case 'isAll':
                        layer.msg(checkStatus.isAll ? '全选' : '未全选');
                        break;
                }
                ;
            });

            // 监听行工具事件
            table.on("tool(" + _tableName + ")", function (obj) {
                var data = obj.data
                    , layEvent = obj.event;

                if (layEvent === 'edit') {
                    // 编辑记录
                    common.edit(_title, data.id, _width, _height, [], function (index, type) {
                        if (type == 2) {
                            // 加载结束
                            $(".layui-laypage-btn").click()//刷新当前页
                        }
                    }, _isClose);
                } else if (layEvent === 'detail') {
                    // 记录详情
                    common.detail(_title, data.id, _width, _height, _isClose);
                } else if (layEvent === 'del') {
                    // 删除记录
                    common.delete(data.id, function (data, res) {
                        if (res) {
                            obj.del();
                        } else {
                        }
                    });
                } else if (layEvent === 'cache') {
                    // 重置缓存
                    common.cache(data.id);
                } else if (layEvent === 'copy') {
                    // 一键复制
                    common.copy(_title, data.id, _width, _height);
                } else {
                    // 其他操作,函数回调
                    if (_callback) {
                        _callback(layEvent, data);
                    }
                }
            });

            // 监听复选框
            table.on("checkbox(" + _tableName + ")", function (obj) {
                // console.log(obj.checked); //当前是否选中状态
                // console.log(obj.data); //选中行的相关数据
                // console.log(obj.type); //如果触发的是全选，则为：all，如果触发的是单选，则为：one
            });

            // 监听单元格编辑
            table.on('edit(' + _tableName + ')', function (obj) {
                var value = obj.value //得到修改后的值
                    , data = obj.data //得到所在行所有键值
                    , field = obj.field; //得到字段

                var json_data = {};
                json_data['id'] = data.id;
                json_data[field] = value;

                // JSON字符串
                var json_str = JSON.stringify(json_data);

                // JSON数据
                var json = JSON.parse(json_str);

                // 发起网络请求
                var url = cUrl + "/update";
                common.ajaxPost(url, json, function (res, success) {
                    // console.log("字段【" + field + "】:【" + value + "】值更新成功");
                }, '更新中...');

            });

            // 监听行单击事件
            table.on("row(" + _tableName + ")", function (obj) {
                // 标注选中样式
                obj.tr.addClass('layui-table-click').siblings().removeClass('layui-table-click');
                var data = obj.data;
            });

            // 监听排序事件
            if (tableSort) {
                table.on("sort(" + _tableName + ")", function (obj) {
                    // // 当前排序的字段名
                    // console.log(obj.field);
                    // // 当前排序类型：desc(降序)、asc(升序)、null(空对象，默认排序)
                    // console.log(obj.type);
                    // // 当前排序的 th对象
                    // console.log(this);

                    // 请求服务端进行动态排序
                    table.reload(_tableName, {
                        initSort: obj
                        , where: {
                            field: obj.field //排序字段
                            , order: obj.type //排序方式
                        }
                    });
                });
            }
            return this;

        },
        /**
         * treetable列表函数,
         */
        treetable: function (cols = [], tableName, isExpand = true, treeSpid = 0, treePidName = '', callback = null, url = '') {
            _tableName = tableName;
            // 初始化请求URL
            if (!url) {
                url = cUrl + "/list";
            }

            // 加载treetable
            var insTb = treeTable.render({
                elem: '#' + tableName, //表格id
                url: url,
                method: "POST",
                height: 'full-50',
                cellMinWidth: 80,
                // toolbar: 'default',
                tree: {
                    iconIndex: 1,
                    idName: 'id',
                    pidName: treePidName ? treePidName : "pid",
                    isPidData: true
                },
                cols: [cols],
                done: function (res, curr, count) {
                    // // res 可以获取文件的数据，或者是ajax请求的数据
                    // console.log(res);
                    // //得到当前页码
                    // console.log(curr);
                    // //得到数据总量
                    // console.log(count);
                    // 关闭加载
                    layer.closeAll('loading');
                },
                style: 'margin-top:0;'
            });

            // 工具条点击事件
            treeTable.on('tool(' + tableName + ')', function (obj) {
                var data = obj.data;
                var layEvent = obj.event;
                // 当前记录ID
                var id = data.id;
                if (layEvent === 'add') {
                    // 添加记录
                    common.edit(_title, 0, _width, _height, ['pid=' + id], function (index, type) {
                        if (type == 2) {
                            // // 树状结构刷新
                            // insTb.clearFilter();
                            // 刷新父页面
                            location.reload();
                        }
                    });
                } else if (layEvent === 'edit') {
                    // 修改记录
                    common.edit(_title, id, _width, _height, [], function (index, type) {
                        if (type == 2) {
                            // // 树状结构刷新
                            // insTb.clearFilter();
                            // 刷新父页面
                            location.reload();
                        }
                    });
                } else if (layEvent === 'addz') {
                    // 添加记录
                    common.edit(_title, 0, _width, _height, ['pid=' + id], function (index, type) {
                        if (type == 2) {
                            // // 树状结构刷新
                            // insTb.clearFilter();
                            // 刷新父页面
                            location.reload();
                        }
                    });
                } else if (layEvent === 'del') {
                    // 删除记录
                    common.delete(id, function (data, isSuc) {
                        if (isSuc) {
                            obj.del();
                        } else {
                        }
                    });
                } else {
                    // 其他操作(回调函数)
                    if (callback) {
                        callback(layEvent, id, 0);
                    }
                }
            });

            // 全部折叠
            $('#collapse').on('click', function () {
                insTb.foldAll();
                return false;
            });

            // 全部展开
            $('#expand').on('click', function () {
                insTb.expandAll();
                return false;
            });

            // 刷新页面
            $('#refresh').on('click', function () {
                insTb.refresh();
                return false;
            });

            // 搜索
            $('#search').click(function () {
                var keywords = $('#keywords').val();
                if (keywords) {
                    insTb.filterData(keywords);
                } else {
                    insTb.clearFilter();
                }
                return false;
            });
        },
        /**
         * 设置弹窗函数
         */
        setWin: function (title, width = 0, height = 0, isClose = false) {
            _title = title;
            _width = width;
            _height = height;
            _isClose = isClose;
            return this;
        },
        /**
         * 设置行双击事件
         */
        setDbclick: function (param) {
            _isDbclick = param ? param : true;
            return this;
        },
        /**
         * 模糊搜索函数
         */
        searchForm: function (searchForm, tableList) {

            // 搜索功能
            form.on("submit(" + searchForm + ")", function (data) {
                common.searchForm(table, data, tableList);
                return false;
            });
        },
        /**
         * TABLE复选框选中函数
         */
        getCheckData: function (tableName) {
            if (!tableName) {
                tableName = _tableName;
            }
            var checkStatus = table.checkStatus(tableName)
                , data = checkStatus.data;
            return data;

        },
        /**
         * 初始化日期组件(支持多组件初始化)
         */
        initDate: function (item, callback = null) {
            common.initDate(item, function (value, date) {
                if (callback) {
                    callback(value, date);
                }
            });
        },
        /**
         * 打开窗体函数
         */
        showWin: function (title, url, width = 0, height = 0, param = [], type = 2, btn = [], callback = null, _isClose = false) {
            common.showWin(title, url, width, height, param, type, btn, function (index, type) {
                if (callback) {
                    callback(index, type);
                }
            }, _isClose);
        },
        /**
         * 网络POST请求
         */
        ajaxPost: function (url, data, callback = null, msg = '处理中...') {
            common.ajaxPost(url, data, callback, msg);
        },
        /**
         * 网络GET请求
         */
        ajaxGet: function (url, data, callback = null, msg = '处理中...') {
            common.ajaxGet(url, data, callback, msg);
        },
        /**
         * Switch开关
         */
        formSwitch: function (name, url = '', callback = null) {
            common.formSwitch(name, url, function (data, res) {
                if (callback) {
                    callback(data, res);
                }
            });
        },
        /**
         * 上传文件
         */
        uploadFile: function (elem_id, callback = null, url = '', exts = 'xls|xlsx', size = 10240, data = {}) {
            common.uploadFile(elem_id, function (res, isSucc) {
                if (callback) {
                    callback(res, isSucc);
                }
            }, url, exts, size, data);
        },
    }

    /**
     * 表单验证函数
     */
    common.verify();

    /**
     * 提交表单
     */
    form.on('submit(submitForm)', function (data) {
        // // 带[]中括号下标的字段特殊处理
        // var nameArr = [];
        // var itemArr = [];
        // var param = data.field;
        // $.each(param, function (key, val) {
        //     // 正则验证字段是否存在中括号[]
        //     var regex = /\[|\]|【|】/g
        //     if (!regex.test(key)) {
        //         return;
        //     }
        //
        //     // 处理带括号[]的字段
        //     var regex1 = /\[(.+?)\]/g;   // [] 中括号及内容
        //
        //     // 获取括号及括号内容
        //     var content = key.match(regex1);
        //
        //     // 获取括号内容值
        //     var regex2 = "\\[(.+?)\\]";
        //     var item = key.match(regex2);
        //     val = item[1];
        //
        //     // 获取字段名
        //     var name = key.replace(content, "");
        //     // 字段名临时存储
        //     if ($.inArray(name, nameArr) < 0) {
        //         nameArr.push(name);
        //     }
        //
        //     // 字段名数组初始化
        //     if (!itemArr[name]) {
        //         itemArr[name] = [];
        //     }
        //     itemArr[name].push(val);
        // });
        // // 遍历数组
        // $.each(nameArr, function (i, name) {
        //     var item = [];
        //     $.each(itemArr[name], function (key, val) {
        //         item.push(val);
        //         // 移除指定元素
        //         delete param[name + "[" + val + "]"];
        //     });
        //     param[name] = item.join(",");
        // });
        common.submitForm(data.field, null, function (res, success) {
            console.log("保存成功回调");
        });
        return false;
    });

    /**
     * 关键词搜索
     */
    form.on("submit(searchForm)", function (data) {
        common.searchForm(table, data);
        return false;
    });

    /**
     * 操作按钮
     */
    $(".btnOption").click(function () {

        // 自定义参数
        var param = $(this).attr("data-param");
        if (param != null) {
            console.log(param);
            param = JSON.parse(param);
            console.log(param);
        }
        // 选择数据
        var data = func.getCheckData(_tableName);

        // 事件名称
        var layEvent = $(this).attr('lay-event');
        switch (layEvent) {
            case "add": {
                // 添加记录
                common.edit(_title, 0, _width, _height, param, function (index, type) {
                    if (type == 2) {
                        // 加载结束
                        location.reload();
                    }
                }, _isClose);
                break;
            }
            case "dall": {
                // 批量删除

                // 方法参数
                var item = {};
                item['title'] = "批量删除";
                item['url'] = cUrl + "/delete";
                item['data'] = data;
                item['confirm'] = true;
                item['type'] = "POST";

                // 执行方法
                common.batchFunc(item, function () {
                    _tableIns.reload();
                });
                break;
            }
            case "batchCache": {
                // 批量重置缓存

                // 方法参数
                var item = {};
                item['title'] = "批量重置缓存";
                item['url'] = cUrl + "/batchCache";
                item['data'] = data;
                item['confirm'] = true;
                item['type'] = "GET";

                // 执行方法
                common.batchFunc(item, function () {
                    _tableIns.reload();
                });
                break;
            }
            case "batchEnable": {
                // 批量启用

                // 方法参数
                var item = {};
                item['title'] = "批量启用状态";
                item['url'] = cUrl + "/batchStatus";
                item['param'] = param;
                item['data'] = data;
                item['form'] = "submitForm";
                item['confirm'] = true;
                item['show_tips'] = "处理中...";
                item['type'] = "POST";

                // 执行方法
                common.batchFunc(item, function () {
                    _tableIns.reload();
                });
                break;
            }
            case "batchDisable": {
                // 批量禁用

                // 方法参数
                var item = {};
                item['title'] = "批量禁用状态";
                item['url'] = cUrl + "/batchStatus";
                item['param'] = param;
                item['data'] = data;
                item['confirm'] = true;
                item['show_tips'] = "处理中...";
                item['type'] = "POST";

                // 执行方法
                common.batchFunc(item, function () {
                    _tableIns.reload();
                });
                break;
            }
            case "export": {
                // 导出Excel

                // 自定义参数
                var param = [];
                var item = $('.layui-form-item [name]').serializeArray();
                $.each(item, function () {
                    param.push(this.name + "=" + this.value)
                });

                // 方法参数
                var item = {};
                item['title'] = "导出数据";
                item['url'] = cUrl + "/export";
                item['data'] = data;
                item['confirm'] = true;
                item['type'] = "POST";
                item['show_tips'] = "数据准备中...";
                item['param'] = param;

                // 执行方法
                common.batchFunc(item, function (res, isSucc) {
                    // 下载文件
                    window.location.href = "/common/download?fileName=" + encodeURI(res.data) + "&isDelete=" + true;
                });
                break;
            }
            case "import": {
                // 导入Excel
                common.uploadFile('import', function (res, isSucc) {
                    // TODO...
                });
                break;
            }
        }
    });

    //关闭自身
    window.formClose = function () {
        // 先得到当前iframe层的索引
        var index = parent.layer.getFrameIndex(window.name);
        // 再执行关闭
        parent.layer.close(index);
    };

    /**
     * 输入自定义模块(此模块接口是对象)
     */
    exports('func', func);
});
