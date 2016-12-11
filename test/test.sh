#!/bin/bash
#admin登录
curl localhost:8080/login -d'{"Name":"admin","Password":"admin"}' -i -c cookie	
curl localhost:8080/login -d'{"Name":"xiayu","Password":"123"}' -i -c cookie
#登出
curl -XDELETE localhost:8080/login -i -b cookie
#新增用户
curl -i -XPUT localhost:8080/user -d'{"Name":"xiayu","Password":"123"}' -b cookie
curl -i -XPUT localhost:8080/User -d'{"Name":"xiayu2","Password":"123"}' -b cookie
curl -i -XPUT localhost:8080/User -d'{"Name":"xiayu3","Password":"123"}' -b cookie
#删除用户
curl -i -XDELETE localhost:8080/User -d'{"Name":"xiayu"}' -b cookie
