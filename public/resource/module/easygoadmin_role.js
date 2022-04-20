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
 * 角色管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'admin', 'zTree'], function () {

    //声明变量
    var func = layui.func
        , admin = layui.admin
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'name', width: 200, title: '角色名称', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                return  '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" '+(d.status==1 ? 'checked' : '')+'>';
            }}
            , {field: 'sort', width: 100, title: '排序', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 250, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'permission') {
                admin.open({
                    title: '角色权限分配',
                    btn: ['保存', '取消'],
                    content: '<ul id="roleAuthTree" class="ztree"></ul>',
                    success: function (layero, dIndex) {
                        var loadIndex = layer.load(2);
                        $.get('/rolemenu/index', {roleId: data.id}, function (res) {
                            layer.close(loadIndex);
                            if (0 === res.code) {
                                $.fn.zTree.init($('#roleAuthTree'), {
                                    check: {enable: true},
                                    data: {simpleData: {enable: true}}
                                }, res.data);
                            } else {
                                layer.msg(res.msg, {icon: 2});
                            }
                        }, 'json');
                        // 超出一定高度滚动
                        $(layero).children('.layui-layer-content').css({'max-height': '300px', 'overflow': 'auto'});
                    },
                    yes: function (dIndex) {
                        var insTree = $.fn.zTree.getZTreeObj('roleAuthTree');
                        var checkedRows = insTree.getCheckedNodes(true);
                        var ids = [];
                        for (var i = 0; i < checkedRows.length; i++) {
                            ids.push(checkedRows[i].id);
                        }
                        func.ajaxPost("/rolemenu/save", {roleId: data.id, menuIds: ids.join(',')}, function (res, success) {
                            // 关闭窗体
                            layer.close(dIndex);
                        });
                    }
                });
            }
        });

        //【设置弹框】
        func.setWin("角色", 500, 350);

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
