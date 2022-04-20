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
 * 消息管理
 * @author 半城风雨
 * @since 2021/7/26
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
            , {field: 'title', width: 200, title: '消息标题', align: 'center'}
            , {field: 'content', width: 300, title: '消息内容', align: 'center'}
            , {field: 'receiver', width: 100, title: '接收人', align: 'center'}
            , {field: 'type', width: 100, title: '发送方式', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 系统
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 短信
                    cls = "layui-btn-danger";
                } else if (d.type == 3) {
                    // 邮件
                    cls = "layui-btn-warm";
                } else if (d.type == 4) {
                    // 微信
                    cls = "layui-btn-primary";
                } else if (d.type == 5) {
                    // 推送
                    cls = "layui-btn-disabled";
                }

				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'sendTime', width: 180, title: '发送时间', align: 'center'}
            , {field: 'sendStatus', width: 100, title: '发送状态', align: 'center', templet(d) {
                var cls = "";
                if (d.sendStatus == 1) {
                    // 已发送
                    cls = "layui-btn-normal";
                } else if (d.sendStatus == 2) {
                    // 未发送
                    cls = "layui-btn-danger";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.sendStatusName+'</span>';
            }}
            , {field: 'sendNum', width: 100, title: '发送次数', align: 'center'}
            , {field: 'note', width: 200, title: '发送备注', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("系统消息");

    }
});
