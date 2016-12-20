#!/bin/bash

#登录
curl localhost:8080/login -d'{"name":"admin","password":"admin"}' -i -c /tmp/cookie

#新增用户
curl -i -XPOST localhost:8080/user -d'{"Name":"xiayu","Password":"123"}' -b /tmp/cookie
#删除用户
curl -i -XDELETE localhost:8080/user/2 -d'{"Name":"xiayu"}' -b /tmp/cookie

#获得帐户类型列表
curl -i localhost:8080/account-type -b /tmp/cookie
#新增一个账户信息
curl -i XPOST localhost:8080/account -d'{}'

#登出
curl -XDELETE localhost:8080/login -i -b /tmp/cookie
