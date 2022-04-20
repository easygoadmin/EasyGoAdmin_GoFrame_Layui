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
 * 登录
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['layer', 'form', 'index'], function () {
    var $ = layui.jquery;
    var layer = layui.layer;
    var form = layui.form;
    var index = layui.index;
    $('.login-wrapper').removeClass('layui-hide');

    // 登录事件
    form.on('submit(loginSubmit)', function (data) {
        // 设置按钮文字“登录中...”及禁止点击状态
        $(data.elem).attr('disabled', true).text('登录中。。。');

        // 网络请求
        var loadIndex = layer.load(2);
        $.ajax({
            type: "POST",
            url: '/login',
            data: JSON.stringify(data.field),
            contentType: "application/json",
            dataType: "json",
            beforeSend: function () {
                // TODO...
            },
            success: function (res) {
                layer.close(loadIndex);
                if (res.code == 0) {
                    // 清除Tab记忆
                    index.clearTabCache();

                    // 设置登录成功状态
                    $(data.elem).attr('disabled', true).text('登录成功');

                    // 提示语
                    layer.msg('登录成功', {
                        icon: 1,
                        time: 1500
                    });

                    // 延迟3秒
                    setTimeout(function () {
                        // 跳转后台首页
                        window.location.href = "/index";
                    }, 2000);

                    return false;
                } else {
                    // 错误信息
                    layer.msg(res.msg, {icon: 2, anim: 6});
                    // 刷新验证码
                    $('img.login-captcha').click(function () {
                        this.src = '/captcha?t=' + (new Date).getTime();
                    }).trigger('click');

                    // 延迟3秒恢复可登录状态
                    setTimeout(function () {
                        // 设置按钮状态为登录”
                        var login_text = $(data.elem).text().replace('中。。。', '');
                        // 设置按钮为可点击状态
                        $(data.elem).text(login_text).removeAttr('disabled');
                    }, 1000);
                }
            },
            error: function () {
                layer.msg("AJAX请求异常");
            }
        });
        return false;
    });

    // 获取图片验证码
    $('img.login-captcha').click(function () {
        var url = "/captcha?t=" + (new Date).getTime();
        $.ajax({
            type: "get",
            url: url,
            success: function (res) {
                if (res.code == 0) {
                    this.src = res.data;
                    $("#imgcode").attr("src", res.data);
                    $("#idkey").val(res.idkey);
                }
            }
        });
    }).trigger('click');

});