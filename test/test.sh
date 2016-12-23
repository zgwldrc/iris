#!/bin/bash

#admin 登陆
curl -i -XPOST localhost:8080/login -d'{
	"name":"admin","password":"admin"
}' -c /tmp/cookie


#新增一个账户信息
curl -i -XPOST localhost:8080/account -d'{
	"account": "zgwldrc@163.com",
	"password": "test",
	"endpoint": {
		"url": "https://mail.163.com"
	},
	"account_type_id": 1
}' -b /tmp/cookie

#新增一个账户信息
curl -i -XPOST localhost:8080/account -d'{
	"account": "zgwldrc@162.com",
	"password": "test",
	"endpoint": {
		"url": "https://mail.162.com"
	},
	"account_type_id": 1
}' -b /tmp/cookie

#得到一个账户信息
curl -i -XGET -b /tmp/cookie localhost:8080/account/1

#得到我的账户列表
curl -i -XGET -b /tmp/cookie localhost:8080/account
