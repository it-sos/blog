#!/bin/bash

basedir=$(cd "$(dirname "$0")" && pwd)
. $basedir/common.sh
PROJECT=$basedir/run/backend/$PROJECT_NAME

options="{full|backend|frontend} version (eq: v1.0.0)"
version=$2

# 检查buildkitd是否启用
if [ "$(ps -ef|grep buildkitd|grep -v grep|wc -l)" -lt 1 ]; then
    echo "Do the following:"
    echo "sudo buildkitd --oci-worker=false --containerd-worker=true &"
    echo "fg"
    echo "maybe input password"
    echo "ctrl + z"
    echo "bg %1"
    exit 1
fi

# 检查当前版本
if [ -z "$version" ];then 
    echo "Usage: $0 $options"
    echo "frontend current version: "$(grep 'itsos/blog-frontend:' $basedir/k8s/blog-frontend.yaml|awk -F : '{print $3}')
    echo "backend current version: "$(grep 'itsos/blog-backend:' $basedir/k8s/blog-backend.yaml|awk -F : '{print $3}')
    exit 1
fi

# 更新git
cd $basedir && git pull && cd - > /dev/null
if [ "$?" -gt "0" ];then
    exit 1
fi

run() {
  case $1 in
  "full")
    run frontend
    run backend
  ;;
  "backend")
    # 如果只是配置变更执行这个：
    # kubectl create configmap blog-backend-config --from-file=./config/config.yaml
    # kubectl edit configmap blog-backend-config

    run build_linux
    mkdir -p $basedir/run/backend/web
    rm -rf $basedir/run/backend/web/*
    cp -r $basedir/../src/backend/web/views run/backend/web
    sudo buildctl build \
        --frontend=dockerfile.v0 \
        --local context=. \
        --local dockerfile=$basedir/k8s/backend/ \
        --output type=image,name=docker.io/itsos/blog-backend:$version
    echo "sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-backend:$version"
    sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-backend:$version
    sed -i "s#blog-backend:v.*#blog-backend:$version#" $basedir/k8s/blog-backend.yaml
    kubectl apply -f k8s/blog-backend.yaml
  ;;
  "frontend")
    mkdir -p $basedir/run/frontend
    rm -rf $basedir/run/frontend/* $basedir/../src/frontend/dist/*
    cd $basedir/../src/frontend/ && yarn install && yarn build && cp -r dist/* $basedir/run/frontend && cd -
    sudo buildctl build \
        --frontend=dockerfile.v0 \
        --local context=. \
        --local dockerfile=$basedir/k8s/frontend/ \
        --output type=image,name=docker.io/itsos/blog-frontend:$version
    echo "sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-frontend:$version"
    sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-frontend:$version
    sed -i "s#blog-frontend:v.*#blog-frontend:$version#" $basedir/k8s/blog-frontend.yaml
    kubectl apply -f k8s/blog-frontend.yaml
  ;;
  "build_linux")
    build linux
  ;;
  "build_windows")
    build windows
  ;;
  "build_darwin")
    build darwin
  ;;
  *)
    echo "Usage: $0 $options"
  ;;
  esac
}

run $1
