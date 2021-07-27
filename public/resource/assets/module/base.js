/**
 * JS公共函数类
 * @author 鲲鹏
 * @date 2019/6/11
 */
layui.define(['jquery'], function (exports) {
    "use strict";

    // 声明变量
    var $ = layui.$;

    var base = {
        /**
         * 判断字符串是否为空
         */
        isEmpty: function (str) {
            if (str == null || typeof str == "undefined" || str == "") {
                return true;
            }
            return false;
        },
        /**
         * 邮箱格式验证
         */
        isEmail: function (str) {
            var reg = /^[a-z0-9]([a-z0-9\\.]*[-_]{0,4}?[a-z0-9-_\\.]+)*@([a-z0-9]*[-_]?[a-z0-9]+)+([\.][\w_-]+){1,5}$/i;
            if (reg.test(str)) {
                return true;
            } else {
                return false;
            }
        },
        /**
         * 手机号格式验证
         */
        isMobile: function (tel) {
            var reg = /(^1[3|4|5|7|8][0-9]{9}$)/;
            if (reg.test(tel)) {
                return true;
            } else {
                return false;
            }
            ;
        },
        upCase: function (str) {
            if (comm.isEmpty(str)) {
                return;
            }
            return str.substring(0, 1).toUpperCase() + str.substring(1);
        },
        /**
         * 金额数字转大写
         */
        upDigit: function (num) {
            var fraction = ['角', '分', '厘'];
            var digit = ['零', '壹', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'];
            var unit = [
                ['元', '万', '亿'],
                ['', '拾', '佰', '仟']
            ];
            var head = num < 0 ? '欠人民币' : '人民币';
            num = Math.abs(num);
            var s = '';
            for (var i = 0; i < fraction.length; i++) {
                s += (digit[Math.floor(num * 10 * Math.pow(10, i)) % 10] + fraction[i]).replace(/零./, '');
            }
            s = s || '整';
            num = Math.floor(num);
            for (var i = 0; i < unit[0].length && num > 0; i++) {
                var p = '';
                for (var j = 0; j < unit[1].length && num > 0; j++) {
                    p = digit[num % 10] + unit[1][j] + p;
                    num = Math.floor(num / 10);
                }
                s = p.replace(/(零.)*零$/, '').replace(/^$/, '零') + unit[0][i] + s;
                //s = p + unit[0][i] + s;
            }
            return head + s.replace(/(零.)*零元/, '元').replace(/(零.)+/g, '零').replace(/^整$/, '零元整');
        },
        /**
         * 设置cookie
         */
        setCookie: function (name, value, iDay) {
            var oDate = new Date();
            oDate.setDate(oDate.getDate() + iDay);
            document.cookie = name + '=' + value + ';expires=' + oDate;
        },
        /**
         * 获取cookie
         */
        getCookie: function (name) {
            var arr = document.cookie.split('; ');
            for (var i = 0; i < arr.length; i++) {
                var arr2 = arr[i].split('=');
                if (arr2[0] == name) {
                    return arr2[1];
                }
            }
            return '';
        }
        /**
         * 删除Cookie
         */
        , removeCookie: function (name) {
            this.setCookie(name, 1, -1);
        },
        /**
         * 显示
         */
        show: function (obj) {
            var blockArr = ['div', 'li', 'ul', 'ol', 'dl', 'table', 'article', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'p', 'hr', 'header', 'footer', 'details', 'summary', 'section', 'aside', '']
            if (blockArr.indexOf(obj.tagName.toLocaleLowerCase()) === -1) {
                obj.style.display = 'inline';
            } else {
                obj.style.display = 'block';
            }
        },
        /**
         * 隐藏
         */
        hide: function (obj) {
            obj.style.display = "none";
        },
        /**
         * Ajax网络请求
         */
        ajax: function (obj) {
            obj = obj || {};
            obj.type = obj.type.toUpperCase() || 'POST';
            obj.url = obj.url || '';
            obj.async = obj.async || true;
            obj.data = obj.data || null;
            obj.success = obj.success || function () {
            };
            obj.error = obj.error || function () {
            };
            var xmlHttp = null;
            if (XMLHttpRequest) {
                xmlHttp = new XMLHttpRequest();
            } else {
                xmlHttp = new ActiveXObject('Microsoft.XMLHTTP');
            }
            var params = [];
            for (var key in obj.data) {
                params.push(key + '=' + obj.data[key]);
            }
            var postData = params.join('&');
            if (obj.type.toUpperCase() === 'POST') {
                xmlHttp.open(obj.type, obj.url, obj.async);
                xmlHttp.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded;charset=utf-8');
                xmlHttp.send(postData);
            } else if (obj.type.toUpperCase() === 'GET') {
                xmlHttp.open(obj.type, obj.url + '?' + postData, obj.async);
                xmlHttp.send(null);
            }
            xmlHttp.onreadystatechange = function () {
                if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
                    obj.success(xmlHttp.responseText);
                } else {
                    obj.error(xmlHttp.responseText);
                }
            };
        },
        /**
         * 数据类型判断
         * 案例：istype([],'array')
         */
        istype: function (o, type) {
            if (type) {
                var _type = type.toLowerCase();
            }
            switch (_type) {
                case 'string':
                    return Object.prototype.toString.call(o) === '[object String]';
                case 'number':
                    return Object.prototype.toString.call(o) === '[object Number]';
                case 'boolean':
                    return Object.prototype.toString.call(o) === '[object Boolean]';
                case 'undefined':
                    return Object.prototype.toString.call(o) === '[object Undefined]';
                case 'null':
                    return Object.prototype.toString.call(o) === '[object Null]';
                case 'function':
                    return Object.prototype.toString.call(o) === '[object Function]';
                case 'array':
                    return Object.prototype.toString.call(o) === '[object Array]';
                case 'object':
                    return Object.prototype.toString.call(o) === '[object Object]';
                case 'nan':
                    return isNaN(o);
                case 'elements':
                    return Object.prototype.toString.call(o).indexOf('HTML') !== -1
                default:
                    return Object.prototype.toString.call(o)
            }
        },
        /**
         * 关键字加标签（多个关键词用空格隔开）
         * 案例：findKey('守侯我oaks接到了来自下次你离开快乐吉祥留在开城侯','守侯 开','i')
         */
        findKey: function (str, key, el) {
            var arr = null,
                regStr = null,
                content = null,
                Reg = null,
                _el = el || 'span';
            arr = key.split(/\s+/);
            //alert(regStr); //    如：(前端|过来)
            regStr = this.createKeyExp(arr);
            content = str;
            //alert(Reg);//        /如：(前端|过来)/g
            Reg = new RegExp(regStr, "g");
            //过滤html标签 替换标签，往关键字前后加上标签
            content = content.replace(/<\/?[^>]*>/g, '')
            return content.replace(Reg, "<" + _el + ">$1</" + _el + ">");
        },
        /**
         * 获取URL参数
         * 调用：get_url_param('http://xxxx?draftId=122000011938')
         */
        get_url_param: function (url) {
            url = url ? url : window.location.href;
            var _pa = url.substring(url.indexOf('?') + 1),
                _arrS = _pa.split('&'),
                _rs = {};
            for (var i = 0, _len = _arrS.length; i < _len; i++) {
                var pos = _arrS[i].indexOf('=');
                if (pos == -1) {
                    continue;
                }
                var name = _arrS[i].substring(0, pos),
                    value = window.decodeURIComponent(_arrS[i].substring(pos + 1));
                _rs[name] = value;
            }
            return _rs;
        },
        /**
         * 设置URL参数
         * 调用：set_url_param({'a':1,'b':2})
         */
        set_url_param: function (obj) {
            var _rs = [];
            for (var p in obj) {
                if (obj[p] != null && obj[p] != '') {
                    _rs.push(p + '=' + obj[p])
                }
            }
            return _rs.join('&');
        },
        /**
         * 随机产生颜色
         */
        random_color: function () {
            //randomNumber是下面定义的函数
            //写法1
            //return 'rgb(' + this.randomNumber(255) + ',' + this.randomNumber(255) + ',' + this.randomNumber(255) + ')';

            //写法2
            return '#' + Math.random().toString(16).substring(2).substr(0, 6);

            //写法3
            //var color='#',_index=this.randomNumber(15);
            //for(var i=0;i<6;i++){
            //color+='0123456789abcdef'[_index];
            //}
            //return color;
        },
        /**
         * 随机返回一定范围的数字
         */
        random_number: function (n1, n2) {
            //randomNumber(5,10)
            //返回5-10的随机整数，包括5，10
            if (arguments.length === 2) {
                return Math.round(n1 + Math.random() * (n2 - n1));
            }
            //randomNumber(10)
            //返回0-10的随机整数，包括0，10
            else if (arguments.length === 1) {
                return Math.round(Math.random() * n1)
            }
            //randomNumber()
            //返回0-255的随机整数，包括0，255
            else {
                return Math.round(Math.random() * 255)
            }
        },
        /**
         * 数字排序
         * 调用：array_sort(arr,'a,b')a是第一排序条件，b是第二排序条件
         */
        array_sort: function (arr, sort) {
            if (!sort) {
                return arr
            }
            var _sort = sort.split(',').reverse(),
                _arr = arr.slice(0);
            for (var i = 0, len = _sort.length; i < len; i++) {
                _arr.sort(function (n1, n2) {
                    return n1[_sort[i]] - n2[_sort[i]]
                })
            }
            return _arr;
        }
    };

    /**
     * 输出模块(此模块接口是对象)
     */
    exports('base', base);
});
