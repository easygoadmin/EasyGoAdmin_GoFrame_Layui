/**
 * 短信日志
 * @auth 鲲鹏
 * @date 2020-05-04
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'mobile', width: 130, title: '手机号码', align: 'center'}
            , {field: 'type', width: 100, title: '发送类型', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 用户注册
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 修改密码
                    cls = "layui-btn-danger";
                } else if (d.type == 3) {
                    // 找回密码
                    cls = "layui-btn-warm";
                } else if (d.type == 4) {
                    // 换绑手机号验证
                    cls = "layui-btn-primary";
                } else if (d.type == 5) {
                    // 换绑手机号
                    cls = "layui-btn-disabled";
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'content', width: 250, title: '短信内容', align: 'center'}
            , {field: 'sign', width: 100, title: '项目签名', align: 'center'}
            , {field: 'templateCode', width: 100, title: '模板编码', align: 'center'}
            , {field: 'params', width: 100, title: '参数', align: 'center'}
            , {field: 'bizId', width: 120, title: '阿里云BizId', align: 'center'}
            , {field: 'code', width: 100, title: 'code编码', align: 'center'}
            , {field: 'message', width: 100, title: '错误信息', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                var cls = "";
                if (d.status == 1) {
                    // 成功
                    cls = "layui-btn-normal";
                } else if (d.status == 2) {
                    // 失败
                    cls = "layui-btn-danger";
                } else if (d.status == 3) {
                    // 待处理
                    cls = "layui-btn-warm";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.statusName+'</span>';
            }}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");
    }
});
