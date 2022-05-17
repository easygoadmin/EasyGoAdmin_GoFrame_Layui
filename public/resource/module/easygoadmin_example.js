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
 * 演示一管理
 * @author 半城风雨
 * @since 2021/08/07
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
        
             
			, {field: 'name', width: 100, title: '测试名称', align: 'center'}
            
        
            
            , {field: 'avatar', width: 90, title: '头像', align: 'center', templet: function (d) {
                    var avatarStr = "";
                    if (d.avatar) {
                        avatarStr = '<a href="' + d.avatar + '" target="_blank"><img src="' + d.avatar + '" height="26" /></a>';
                    }
                    return avatarStr;
                }
            }
			
        
             
			, {field: 'content', width: 100, title: '内容', align: 'center'}
            
        
            
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|停用" lay-filter="Status" '+(d.status==1 ? 'checked' : '')+'>';
            }}
            
        
            
			, {field: 'type', width: 100, title: '类型', align: 'center', templet(d) {
					
					if (d.type == 1) {
						// 京东
						return '<span class="layui-btn layui-btn-normal layui-btn-xs">京东</span>';
					} 
					
					else if (d.type == 2) {
						// 淘宝
						return '<span class="layui-btn layui-btn-danger layui-btn-xs">淘宝</span>';
					} 
					
					else if (d.type == 3) {
						// 拼多多
						return '<span class="layui-btn layui-btn-warm layui-btn-xs">拼多多</span>';
					} 
					
					else if (d.type == 4) {
						// 唯品会
						return '<span class="layui-btn layui-btn-primary layui-btn-xs">唯品会</span>';
					} 
					
				}
			}
			
        
            
            , {field: 'is_vip', width: 100, title: '是否VIP', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="is_vip" value="' + d.id + '" lay-skin="switch" lay-text="是|否" lay-filter="IsVip" '+(d.isVip==1 ? 'checked' : '')+'>';
            }}
            
        
             
			, {field: 'sort', width: 100, title: '排序号', align: 'center'}
            
        
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        
        func.setWin("演示一", 750, 450);
        

    
		
	
		
	
		
	
		
		//【设置状态】
        func.formSwitch('Status', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
		
		//【设置是否VIP】
        func.formSwitch('IsVip', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
    }
});
