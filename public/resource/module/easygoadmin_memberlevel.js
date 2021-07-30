/**
 * 会员等级
 * @author 牧羊人
 * @since 2020/7/4
 */
layui.use(['func'], function () {

    //【声明变量】
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'name', width: 200, title: '分组名称', align: 'center'}
            , {field: 'sort', width: 120, title: '排序', align: 'center'}
            , {field: 'createTime', width: 200, title: '创建时间', align: 'center', sort: true}
            , {field: 'updateTime', width: 200, title: '更新时间', align: 'center', sort: true}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("会员等级", 450, 250);
    }
});
