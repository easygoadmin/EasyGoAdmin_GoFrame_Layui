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
 * 城市管理
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
              {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '城市名称', align: 'left'}
            , {field: 'level', width: 100, title: '城市级别', align: 'center', templet(d) {
                var cls = "";
                var levelStr = ""
                if (d.level == 1) {
                    // 省份
                    cls = "layui-btn-normal";
                    levelStr = "省份"
                } else if (d.level == 2) {
                    // 市区
                    cls = "layui-btn-danger";
                    levelStr = "市区"
                } else if (d.level == 3) {
                    // 区县
                    cls = "layui-btn-warm";
                    levelStr = "区县"
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+levelStr+'</span>';
            }}
            , {field: 'citycode', width: 150, title: '城市编号（区号）', align: 'center'}
            , {field: 'pAdcode', width: 150, title: '父级地理编号', align: 'center'}
            , {field: 'adcode', width: 150, title: '地理编号', align: 'center'}
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'createUserName', width: 100, title: '添加人', align: 'center'}
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center'}
            , {width: 200, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("城市",700, 400);

    }
});
