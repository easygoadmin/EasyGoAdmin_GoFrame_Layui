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
 * 通知公告
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'title', width: 350, title: '通知标题', align: 'center'}
            , {field: 'source', width: 100, title: '通知来源', align: 'center', templet(d) {
                var cls = "";
                if (d.source == 1) {
                    cls = "layui-btn-normal";
                } else if (d.source == 2) {
                    cls = "layui-btn-primary"
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.sourceName+'</span>';
            }}
            , {field: 'isTop', width: 100, title: '是否置顶', align: 'center', templet(d) {
                if (d.isTop == 1) {
                    // 已置顶
                    return '<span class="layui-btn layui-btn-primary layui-btn-xs">已置顶</span>';
                } else if (d.isTop == 2) {
                    // 未置顶
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">未置顶</span>';
                }
            }}
            , {field: 'status', width: 100, title: '发布状态', align: 'center', templet(d) {
                if (d.status == 1) {
                    // 草稿箱
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">已发布</span>';
                } else if (d.status == 2) {
                    // 立即发布
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">待发布</span>';
                }
            }}
            , {field: 'browse', width: 100, title: '阅读量', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("通知公告");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
