#!/bin/bash

# configrc 配置

config=$basedir/config

if [ ! -e "$config/configrc" ]; then
  is_exit=1
  mkdir -p $config
  echo "<<<<<<<<<<<<<<"
  echo "Please modify the configuration file and execute again: $config/configrc"
  echo ">>>>>>>>>>>>>>"
  cat > $config/configrc <<EOF
## 常规配置
# 远程服务器地址
HOST=host
# 远程服务器ssh端口
PORT=port
# 远程服务器用户
USER=user
# 远程服务器密码
PASSWORD=password

## 如果在 /ssh/ssh/ssh_config 进行过配置且已经公钥授权（无密码登录），此项适用于通过跳板机远程部署
# 如果此处配置了，将忽略远程服务器地址、端口等配置
JUMP_HOST_NAME=jump_host_name

# 项目名
PROJECT_NAME=studynotes
# 项目远程目录
GO_PROJ_DIST=/data1/supervisord/study-notes/go
# 静态资源远程目录
STATIC_PROJ_DIST=/data1/supervisord/study-notes/html
EOF
fi

# config.yaml 配置
if [ ! -e "$config/config.yaml" ]; then
  is_exit=1
  echo "<<<<<<<<<<<<<<"
  echo "Please modify the configuration file and execute again: $config/config.yaml"
  echo ">>>>>>>>>>>>>>"
  cp $basedir/../src/backend/config.yaml $config/config.yaml
fi

# 加载配置
. $config/configrc

PROJECT=$basedir/run/$PROJECT_NAME

unlinkProject() {
  if [ -e "$PROJECT" ];then
    rm $PROJECT
  fi
}

export PATH=$PATH:/usr/local/go/bin/
go env -w GOPROXY="https://goproxy.cn,direct"

build() {
  unlinkProject
  mkdir -p $basedir/run
  cd $basedir/../src/backend/cmd/$PROJECT_NAME && \
      go get -v ./... && \
      CGO_ENABLED=0 GOOS=$1 go build -v -o $PROJECT && cd -
}

build_sock() {
  unlinkProject
  mkdir -p $basedir/run
  cd $basedir/../src/backend/cmd/$PROJECT_NAME"_sock/" && \
      CGO_ENABLED=0 GOOS=$1 go build -v -o $PROJECT && cd -
}
