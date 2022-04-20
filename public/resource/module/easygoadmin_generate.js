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
 * 代码生成器
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //【声明变量】
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'name', width: 150, title: '表名', align: 'center'}
            , {field: 'engine', width: 100, title: '引擎', align: 'center'}
            , {field: 'version', width: 100, title: '版本', align: 'center'}
            , {field: 'collation', width: 180, title: '编码', align: 'center'}
            , {field: 'rows', width: 100, title: '记录数', align: 'center'}
            , {field: 'data_length', width: 100, title: '大小', align: 'center'}
            , {field: 'auto_increment', width: 100, title: '自增索引', align: 'center'}
            , {field: 'comment', width: 150, title: '表备注', align: 'center'}
            , {
                field: '', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '未备份';
                }
            }
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center', sort: true}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center', sort: true}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'generate') {
                // 一键生成模块
                layer.confirm('您确定要生成表【' + data.name + '】对应的模块吗？生成后将会覆盖已有的模块文件！', {
                    icon: 3,
                    skin: 'layer-ext-moon',
                    btn: ['确认', '取消'] //按钮
                }, function (index) {
                    // layer.msg("一键生成模块");
                    // 调用内部方法
                    var url = cUrl + "/generate";
                    func.ajaxPost(url, {"name": data.name, "comment": data.comment}, function (data, flag) {
                        // 关闭弹窗
                        layer.close(index);
                    }, '模块文件生成中。。。');
                });
            }
        });

        // 批量生成
        $("#batchGenerate").click(function () {
            // 选择数据
            var data = func.getCheckData();
            if (data.length == 0) {
                layer.msg("请选择数据表", {icon: 5});
                return false;
            }
            layer.msg("批量生成功能研发中,请耐心等待...")
        });
    }
});
