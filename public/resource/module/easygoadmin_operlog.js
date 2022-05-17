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
 * 操作日志
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
            , {field: 'title', width: 100, title: '模块标题', align: 'center'}
            , {field: 'businessType', width: 100, title: '业务类型', align: 'center', templet(d) {
                var cls = "";
                if (d.businessType == 1) {
                    // 新增
                    cls = "layui-btn-normal";
                } else if (d.businessType == 2) {
                    // 修改
                    cls = "layui-btn-danger";
                } else if (d.businessType == 3) {
                    // 删除
                    cls = "layui-btn-warm";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.businessTypeName+'</span>';
            }}
            , {field: 'method', width: 100, title: '方法名称', align: 'center'}
            , {field: 'requestMethod', width: 100, title: '请求方式', align: 'center'}
            , {field: 'operatorType', width: 100, title: '操作类别', align: 'center', templet(d) {
                var cls = "";
                if (d.operatorType == 1) {
                    // 后台用户
                    cls = "layui-btn-normal";
                } else if (d.operatorType == 2) {
                    // 手机端用户
                    cls = "layui-btn-danger";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.operatorTypeName+'</span>';
            }}
            , {field: 'operName', width: 100, title: '操作人员', align: 'center'}
            , {field: 'operUrl', width: 100, title: '请求URL', align: 'center'}
            , {field: 'operIp', width: 130, title: '主机地址', align: 'center'}
            , {field: 'operLocation', width: 100, title: '操作地点', align: 'center'}
            , {field: 'operParam', width: 100, title: '请求参数', align: 'center'}
            , {field: 'jsonResult', width: 100, title: '返回参数', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                    var cls = "";
                    if (d.status == 1) {
                        // 正常
                        cls = "layui-btn-normal";
                    } else if (d.status == 2) {
                        // 异常
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.statusName+'</span>';
                }}
            , {field: 'errorMsg', width: 100, title: '错误消息', align: 'center'}
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center'}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

    }
});
