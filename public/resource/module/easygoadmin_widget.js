/**
 * 职级
 * @auth 鲲鹏
 * @date 2020-04-20
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    // 下拉菜单事件
    func.dropdownClick("demo", function (event) {
        layer.msg("你点击了test1：" + event);
        console.log("参数：" + event)
    });

    func.popMenu("popmenu", "modify|openvip", function (layEvent) {
        if (layEvent === 'modify') {
            layer.msg("您点击了修改")
        } else if (layEvent === 'openvip') {
            layer.msg("您点击了开通会员");
        }
    });
});
