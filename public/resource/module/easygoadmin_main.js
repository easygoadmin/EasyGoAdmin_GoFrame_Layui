// +----------------------------------------------------------------------
// | RXThink [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2017-2019 http://rxthink.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: 牧羊人 <rxthink@gmail.com>
// +----------------------------------------------------------------------

/**
 * @name 后台首页
 */
layui.define(['jquery', 'larry'], function (exports) {
    "use strict";
    var $ = layui.$,
        larry = layui.larry,
        larryms = layui.larryms;

    $('#closeInfo').on('click', function () {
        $('#Minfo').hide();
    });

    $('#shortcut a').off('click').on('click', function () {
        var data = {
            href: $(this).data('url'),
            icon: $(this).data('icon'),
            ids: $(this).data('id'),
            group: $(this).data('group'),
            title: $(this).children('.larry-value').children('cite').text()
        }
        // larry.addTab(data);
    });

    // 引入larry-panel操作
    larry.panel();

//    var myDate = new Date(),
//           year = myDate.getFullYear(),
//           month = myDate.getMonth()+1,
//           day = myDate.getDate(),
//           time = myDate.toLocaleTimeString();
//       $('#weather').html('您好，现在时间是：'+year+'年'+month+'月'+day+'日 '+time);
//       $('#versionT').text(larryms.version);


    //引用百度图标JS
    larryms.plugin('echarts/echarts.min.js');
    larryms.plugin('echarts/theme/macarons.js', statisticsList);
    var themeName = "";//"macarons";

    //大数据统计入口
    function statisticsList() {

        //会员来源
        statistics1();

        //周会员增量
        statistics2();

        //数据统计
        statistics3();

    }

    /**
     * 会员来源
     */
    function statistics1() {
        var myChart = echarts.init(document.getElementById('larryCount1'), themeName);
        var data = genData(50);
        var option = {
            title: {
                text: '同名数量统计',
                subtext: '纯属虚构',
                x: 'center'
            },
            tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b} : {c} ({d}%)"
            },
            legend: {
                type: 'scroll',
                orient: 'vertical',
                right: 10,
                top: 20,
                bottom: 20,
                data: data.legendData,

                selected: data.selected
            },
            series: [
                {
                    name: '姓名',
                    type: 'pie',
                    radius: '55%',
                    center: ['40%', '50%'],
                    data: data.seriesData,
                    itemStyle: {
                        emphasis: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }
            ]
        };

        function genData(count) {
            var nameList = [
                '赵', '钱', '孙', '李', '周', '吴', '郑', '王', '冯', '陈', '褚', '卫', '蒋', '沈', '韩', '杨', '朱', '秦', '尤', '许', '何', '吕', '施', '张', '孔', '曹', '严', '华', '金', '魏', '陶', '姜', '戚', '谢', '邹', '喻', '柏', '水', '窦', '章', '云', '苏', '潘', '葛', '奚', '范', '彭', '郎', '鲁', '韦', '昌', '马', '苗', '凤', '花', '方', '俞', '任', '袁', '柳', '酆', '鲍', '史', '唐', '费', '廉', '岑', '薛', '雷', '贺', '倪', '汤', '滕', '殷', '罗', '毕', '郝', '邬', '安', '常', '乐', '于', '时', '傅', '皮', '卞', '齐', '康', '伍', '余', '元', '卜', '顾', '孟', '平', '黄', '和', '穆', '萧', '尹', '姚', '邵', '湛', '汪', '祁', '毛', '禹', '狄', '米', '贝', '明', '臧', '计', '伏', '成', '戴', '谈', '宋', '茅', '庞', '熊', '纪', '舒', '屈', '项', '祝', '董', '梁', '杜', '阮', '蓝', '闵', '席', '季', '麻', '强', '贾', '路', '娄', '危'
            ];
            var legendData = [];
            var seriesData = [];
            var selected = {};
            for (var i = 0; i < 50; i++) {
                name = Math.random() > 0.65
                    ? makeWord(4, 1) + '·' + makeWord(3, 0)
                    : makeWord(2, 1);
                legendData.push(name);
                seriesData.push({
                    name: name,
                    value: Math.round(Math.random() * 100000)
                });
                selected[name] = i < 6;
            }

            return {
                legendData: legendData,
                seriesData: seriesData,
                selected: selected
            };

            function makeWord(max, min) {
                var nameLen = Math.ceil(Math.random() * max + min);
                var name = [];
                for (var i = 0; i < nameLen; i++) {
                    name.push(nameList[Math.round(Math.random() * nameList.length - 1)]);
                }
                return name.join('');
            }
        };
        myChart.setOption(option);

        window.onresize = function () {
            myChart.resize();
        };
    }

    /**
     * 周会员增量
     */
    function statistics2() {
        var myChart = echarts.init(document.getElementById('larryCount2'), themeName);
        var xAxisData = [];
        var data1 = [];
        var data2 = [];
        for (var i = 0; i < 100; i++) {
            xAxisData.push('类目' + i);
            data1.push((Math.sin(i / 5) * (i / 5 - 10) + i / 6) * 5);
            data2.push((Math.cos(i / 5) * (i / 5 - 10) + i / 6) * 5);
        }

        var option = {
            title: {
                text: '柱状图动画延迟'
            },
            legend: {
                data: ['bar', 'bar2'],
                align: 'left'
            },
            toolbox: {
                // y: 'bottom',
                feature: {
                    magicType: {
                        type: ['stack', 'tiled']
                    },
                    dataView: {},
                    saveAsImage: {
                        pixelRatio: 2
                    }
                }
            },
            tooltip: {},
            xAxis: {
                data: xAxisData,
                silent: false,
                splitLine: {
                    show: false
                }
            },
            yAxis: {},
            series: [{
                name: 'bar',
                type: 'bar',
                data: data1,
                animationDelay: function (idx) {
                    return idx * 10;
                }
            }, {
                name: 'bar2',
                type: 'bar',
                data: data2,
                animationDelay: function (idx) {
                    return idx * 10 + 100;
                }
            }],
            animationEasing: 'elasticOut',
            animationDelayUpdate: function (idx) {
                return idx * 5;
            }
        };
        myChart.setOption(option);

        window.onresize = function () {
            myChart.resize();
        };
    }


    /**
     * 数据统计
     */
    function statistics3() {
        var myChart = echarts.init(document.getElementById('larryCount3'), themeName),
            option = {
                tooltip: {
                    trigger: 'axis'
                },
                toolbox: {
                    show: true,
                    y: 'bottom',
                    feature: {
                        mark: {show: true},
                        dataView: {show: true, readOnly: false},
                        magicType: {show: true, type: ['line', 'bar', 'stack', 'tiled']},
                        restore: {show: true},
                        saveAsImage: {show: true}
                    }
                },
                calculable: true,
                legend: {
                    data: ['直接访问', '邮件营销', '联盟广告', '视频广告', '搜索引擎', '百度', '谷歌', '必应', '其他']
                },
                xAxis: [
                    {
                        type: 'category',
                        splitLine: {show: false},
                        data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
                    }
                ],
                yAxis: [
                    {
                        type: 'value',
                        position: 'right'
                    }
                ],
                series: [
                    {
                        name: '直接访问',
                        type: 'bar',
                        data: [320, 332, 301, 334, 390, 330, 320]
                    },
                    {
                        name: '邮件营销',
                        type: 'bar',
                        tooltip: {trigger: 'item'},
                        stack: '广告',
                        data: [120, 132, 101, 134, 90, 230, 210]
                    },
                    {
                        name: '联盟广告',
                        type: 'bar',
                        tooltip: {trigger: 'item'},
                        stack: '广告',
                        data: [220, 182, 191, 234, 290, 330, 310]
                    },
                    {
                        name: '视频广告',
                        type: 'bar',
                        tooltip: {trigger: 'item'},
                        stack: '广告',
                        data: [150, 232, 201, 154, 190, 330, 410]
                    },
                    {
                        name: '搜索引擎',
                        type: 'line',
                        data: [862, 1018, 964, 1026, 1679, 1600, 1570]
                    },

                    {
                        name: '搜索引擎细分',
                        type: 'pie',
                        tooltip: {
                            trigger: 'item',
                            formatter: '{a} <br/>{b} : {c} ({d}%)'
                        },
                        center: [160, 130],
                        radius: [0, 50],
                        itemStyle: {
                            normal: {
                                labelLine: {
                                    length: 20
                                }
                            }
                        },
                        data: [
                            {value: 1048, name: '百度'},
                            {value: 251, name: '谷歌'},
                            {value: 147, name: '必应'},
                            {value: 102, name: '其他'}
                        ]
                    }
                ]
            };

        myChart.setOption(option);

        window.onresize = function () {
            myChart.resize();
        }
    }

    exports('main', {});
});