package main

import (
	"github.com/kataras/iris"
	"testing"
)

//后台账号列表
func TestUsersList(t *testing.T) {

	//设置测试数据表
	//测试前后会自动创建和删除表
	SetTestTableName("users")

	users := []*Users{testAdminUser}

	//发起 http 请求
	//Url        string      //测试路由
	//Object     interface{} //发送的json 对象
	//StatusCode int         //返回的 http 状态码
	//Status     bool        //返回的状态
	//Msg        string      //返回提示信息
	//Data       interface{} //返回数据
	bc := BaseCase{"/v1/admin/users", nil, iris.StatusOK, true, "操作成功", users}
	bc.get(t)

}