#!/bin/sh

# 检查静态资源
if [ ! -d web/views/static ];then
    cd ../frontend/ && yarn build:beta
fi

# 启动服务
which rizla > /dev/null 2>&1
if [ $? -gt 0 ];then
    go install github.com/kataras/rizla
fi
rizla cmd/blog/main.go -w $PWD
