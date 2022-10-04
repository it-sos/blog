#!/bin/sh

if [ ! -e "rizla" ];then
    go install github.com/kataras/rizla
fi
rizla cmd/blog/main.go -w $PWD
