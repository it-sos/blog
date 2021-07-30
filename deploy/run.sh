#!/bin/bash

# 构建与部署


if [ ! -x "/usr/bin/expect" ]; then
  echo "The executable command expect is required."
  exit 1
fi

basedir=$(dirname $0)

if [ ! -e "$basedir/configrc" ]; then
  echo "<<<<<<<<<<<<<<"
  echo "Please modify the correct configuration item. vim $basedir/configrc."
  echo ">>>>>>>>>>>>>>"
  echo "=============="
  tee $basedir/configrc <<EOF
## 常规配置
# 远程服务器地址
HOST=host
# 远程服务器端口
PORT=port
# 远程服务器用户
USER=user
# 远程服务器密码
PASSWORD=password

## 如果在 /ssh/ssh/ssh_config 进行过配置且已经公钥授权（无密码登录），此项适用于通过跳板机远程部署
# 如果此处配置了，将忽略上方的配置
JUMP_HOST_NAME=jumphostname

# 项目名
PROJECT_NAME=project_name
# 项目远程目录
GO_PROJ_DIST=/go_proj_dist
# 静态资源远程目录
STATIC_PROJ_DIST=/static_proj_dist
EOF
  echo "=============="
  exit 1
fi

# 加载配置
. $basedir/configrc

# 执行远程命令
remoteShell() {
  echo $(/usr/bin/expect <<-EOF
set timeout -1
if { "$JUMP_HOST_NAME" == "" } {
  spawn ssh $2 -p$PORT $USER@$HOST "$1"
} else {
  spawn ssh $2 $JUMP_HOST_NAME "$1"
}
if { "$2" == "-t" } {
  expect {
    "请输入密码" { send "$PASSWORD\r" }
    "password:" { send "$PASSWORD\r" }
  }
}
expect eof
EOF)
}

# 上传文件
remoteSftp() {
  /usr/bin/expect <<-EOF
set timeout -1
if { "$JUMP_HOST_NAME" == "" } {
  spawn sftp -p$PORT $USER@$HOST
} else {
  spawn sftp $JUMP_HOST_NAME
}
expect "sftp>"
send "put $1 $2\r"
expect "sftp>"
send "bye\r"
expect eof
EOF
}

# 验证预期关键字是否存在
checkExists() {
  echo $(echo $1|grep -ow "{succ}"|wc -l)
}

PROJECT=$basedir/run/$PROJECT_NAME

build() {
  CGO_ENABLED=0 GOOS=$1 go build -v -o $PROJECT $basedir/../cmd/$PROJECT_NAME"_sock/"
}

run() {
  case $1 in
  "first_linux")
    # 第一次部署
    run cover_supervisor
    run linux
  ;;
  "linux")
    run build_linux_go
    run cover_views
    run cover_config
    run restart_supervisor
    run build_vue
    run restart_nginx
  ;;
  "cover_views")
    # 上送静态模板
    remoteShell "mkdir -p $GO_PROJ_DIST/web && rm -rf $GO_PROJ_DIST/web/*"
    remoteSftp "-r $basedir/../web/views" $GO_PROJ_DIST/web
  ;;
  "cover_supervisor")
    # 检查 supervisor 是否安装
    ok=$(remoteShell "which supervisorctl && echo {succ}")
    ok=$(echo $ok|grep -ow "{succ}"|wc -l)
    if [ "$ok" -lt "2" ]; then
      echo "Not found supervisorctl. Please apt-get install supervisor."
      exit 1
    fi
    # 生成配置
    tee $basedir/supervisor/supervisord.conf <<EOF
[program:$PROJECT_NAME]
command=$GO_PROJ_DIST/bin/$PROJECT_NAME -w $GO_PROJ_DIST
user=$USER
autostart=true
autorestart=true
redirect_stderr=true
stdout_logfile=$GO_PROJ_DIST/logs/$PROJECT_NAME.log
stopsignal=15
stopasgroup=true
killasgroup=true
EOF
    # 上送配置
    remoteSftp $basedir/supervisor/supervisord.conf /tmp/supervisord.conf
    # 配置文件归位
    remoteShell "sudo mv /tmp/supervisord.conf /etc/supervisor/conf.d/$PROJECT_NAME.conf" "-t"
  ;;
  "restart_supervisor")
    # 重启 supervisor
    remoteShell "sudo systemctl restart supervisor" "-t"
  ;;
  "restart_nginx")
    # 重启 nginx
    remoteShell "sudo systemctl restart nginx" "-t"
  ;;
  "cover_config")
    # 上送配置文件
    remoteSftp $basedir/run/config.yaml /tmp/config.yaml
    # 转移至项目目录
    remoteShell "mkdir -p $GO_PROJ_DIST && mv /tmp/config.yaml $GO_PROJ_DIST/config.yaml"
  ;;
  "build_vue")
    # vue构建与上送
    cd $basedir/../static && npm run build && cd -
    remoteShell "mkdir -p $STATIC_PROJ_DIST && rm -rf $STATIC_PROJ_DIST/*"
    remoteSftp "-r $basedir/../static/dist/*" $STATIC_PROJ_DIST
  ;;
  "build_linux_go")
    # go构建与上送
    build linux
    # 不管目录是否存在仍创建
    remoteShell "mkdir -p $GO_PROJ_DIST/{bin,logs}"
    # 上送到指定目录
    PROJ_NAME=$PROJECT_NAME.`date +%s`
    remoteSftp $PROJECT $GO_PROJ_DIST/bin/$PROJ_NAME
    # 创建执行程序软连接
    remoteShell "ln -sf $GO_PROJ_DIST/bin/$PROJ_NAME $GO_PROJ_DIST/bin/$PROJECT_NAME"
  ;;
  "windows")
    build windows
    echo "仅做演示..."
  ;;
  "darwin")
    build darwin
    echo "仅做演示..."
  ;;
  *)
    echo "Usage: $0 linux|windows|darwin"
  ;;
  esac
}

run $1