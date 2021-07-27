/**
 * 系统人员
 * @auth 鲲鹏
 * @date 2020-04-20
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
            , {field: 'realname', width: 100, title: '真实姓名', align: 'center'}
            , {field: 'gender', width: 60, title: '性别', align: 'center', templet(d) {
                var cls = "";
                if (d.gender == 1) {
                    // 男
                    cls = "layui-btn-normal";
                } else if (d.gender == 2) {
                    // 女
                    cls = "layui-btn-danger";
                } else if (d.gender == 3) {
                    // 保密
                    cls = "layui-btn-warm";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.genderName+'</span>';
            }}
            , {field: 'avatar', width: 80, title: '头像', align: 'center', templet: function (d) {
                var avatarStr = "";
                if (d.avatarUrl) {
                    avatarStr = '<a href="' + d.avatarUrl + '" target="_blank"><img src="' + d.avatarUrl + '" height="26" /></a>';
                }
                return avatarStr;
              }
            }
            , {field: 'mobile', width: 130, title: '手机号码', align: 'center'}
            , {field: 'email', width: 200, title: '邮箱地址', align: 'center'}
            , {field: 'birthday', width: 120, title: '出生日期', align: 'center'}
            , {field: 'deptName', width: 200, title: '所属部门', align: 'center'}
            , {field: 'levelName', width: 120, title: '职级名称', align: 'center'}
            , {field: 'positionName', width: 120, title: '岗位名称', align: 'center'}
            , {field: 'cityName', width: 200, title: '所属城市', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: '#statusTpl'}
            , {field: 'sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'loginNum', width: 100, title: '登录次数', align: 'center'}
            , {field: 'loginIp', width: 100, title: '最近登录IP', align: 'center'}
            , {field: 'loginTime', width: 180, title: '最近登录时间', align: 'center'}
            , {field: 'createUserName', width: 100, title: '添加人', align: 'center'}
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center'}
            , {field: 'updateUserName', width: 100, title: '更新人', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("用户");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
