/**
 * 网络请求类
 * @author 牧羊人
 * @date 2019/5/16
 * 参考网址：https://www.cnblogs.com/mophy/p/8428622.html
 */
layui.define(['layer'], function (exports) {
    "use strict";
    var layer = layui.layer,
        $ = layui.$;

    /**
     * 自定义ajax网络请求模块
     */
    var ajax = {
        /**
         * 创建异步请求对象方法
         */
        createXHR: function () {
            if (window.XMLHttpRequest) { // IE7+、Firefox、Opera、Chrome 和Safari
                return new XMLHttpRequest();
            } else if (window.ActiveXObject) { // IE6 及以下
                var versions = ['MSXML2.XMLHttp', 'Microsoft.XMLHTTP'];
                for (var i = 0, len = versions.length; i < len; i++) {
                    try {
                        return new ActiveXObject(version[i]);
                        break;
                    } catch (e) {
                        // 跳过
                    }
                }
            } else {
                throw new Error('浏览器不支持XHR对象！');
            }
        },
        /**
         * 初始化数据方法
         */
        init: function (obj) {
            // 初始化数据
            var objAdapter = {
                method: 'get',
                data: {},
                success: function () {
                },
                complete: function () {
                },
                error: function (s) {
                    alert('status:' + s + 'error!');
                },
                async: true
            }
            // 通过使用JS随机字符串解决IE浏览器第二次默认获取缓存的问题
            ajax.url = obj.url + '?rand=' + Math.random();
            ajax.method = obj.method || objAdapter.method;
            ajax.data = ajax.params(obj.data) || ajax.params(objAdapter.data);
            ajax.async = obj.async || objAdapter.async;
            ajax.complete = obj.complete || objAdapter.complete;
            ajax.success = obj.success || objAdapter.success;
            ajax.error = obj.error || objAdapter.error;
        },
        /**
         * Ajax异步调用
         */
        ajax: function (obj) {
            ajax.method = obj.method || 'get';
            if (obj.method === 'post') {
                ajax.post(obj);
            } else {
                ajax.get(obj);
            }
        },
        /**
         * Ajax请求POST
         */
        post: function (obj) {
            var xhr = ajax.createXHR(); // 创建XHR对象
            ajax.init(obj);
            ajax.method = 'post';
            if (ajax.async === true) { // true表示异步，false表示同步
                // 使用异步调用的时候，需要触发readystatechange 事件
                xhr.onreadystatechange = function () {
                    if (xhr.readyState == 4) { // 判断对象的状态是否交互完成
                        ajax.callback(obj, this); // 回调
                    }
                };
            }
            // 在使用XHR对象时，必须先调用open()方法，
            // 它接受三个参数：请求类型(get、post)、请求的URL和表示是否异步。
            xhr.open(ajax.method, ajax.url, ajax.async);
            // post方式需要自己设置http的请求头，来模仿表单提交。
            // 放在open方法之后，send方法之前。
            xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
            xhr.send(ajax.data); // post方式将数据放在send()方法里
            if (ajax.async === false) { // 同步
                ajax.callback(obj, this); // 回调
            }
        },
        /**
         * Ajax请求GET
         */
        get: function (obj) {
            var xhr = ajax.createXHR(); // 创建XHR对象
            ajax.init(obj);
            if (ajax.async === true) { // true表示异步，false表示同步
                // 使用异步调用的时候，需要触发readystatechange 事件
                xhr.onreadystatechange = function () {
                    if (xhr.readyState == 4) { // 判断对象的状态是否交互完成
                        ajax.callback(obj, this); // 回调
                    }
                };
            }
            // 若是GET请求，则将数据加到url后面
            ajax.url += ajax.url.indexOf('?') == -1 ? '?' + ajax.data : '&' + ajax.data;
            // 在使用XHR对象时，必须先调用open()方法，
            // 它接受三个参数：请求类型(get、post)、请求的URL和表示是否异步。
            xhr.open(ajax.method, ajax.url, ajax.async);
            xhr.send(null); // get方式则填null
            if (ajax.async === false) { // 同步
                ajax.callback(obj, this); // 回调
            }
        },
        /**
         * 请求成功后,回调方法
         */
        callback: function (obj, xhr) {
            if (xhr.status == 200) { // 判断http的交互是否成功，200表示成功
                obj.success(xhr.responseText); // 回调传递参数
            } else {
                alert('获取数据错误！错误代号：' + xhr.status + '，错误信息：' + xhr.statusText);
            }
        },
        /**
         * 数据转换
         */
        params: function (data) {
            var arr = [];
            for (var i in data) {
                // 特殊字符传参产生的问题可以使用encodeURIComponent()进行编码处理
                arr.push(encodeURIComponent(i) + '=' + encodeURIComponent(data[i]));
            }
            return arr.join('&');
        }
    };

    /**
     * 输入自定义模块(此模块接口是对象)
     */
    exports('ajax', ajax);
});