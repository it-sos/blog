#!/bin/bash

basedir=$(cd "$(dirname "$0")" && pwd)
. $basedir/common.sh
PROJECT=$basedir/run/backend/$PROJECT_NAME


options="{full|backend|frontend} [version(default=v1.0.0)]"
version=$2
if [ -z "$2" ]; then
    version="v1.0.0"
fi

run() {
  case $1 in
  "full")
    run frontend
    run backend
  ;;
  "backend")
    run build_linux
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
    sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-frontend:$version
    sed -i "s#blog-frontend:v.*#blog-frontend:$version#" $basedir/k8s/blog-frontend.yaml
  ;;
  "build_linux")
    build linux
    mkdir -p $basedir/run/backend/web
    rm -rf $basedir/run/backend/web/*
    cp -r $basedir/../src/backend/web/views run/backend/web
    sudo buildctl build \
        --frontend=dockerfile.v0 \
        --local context=. \
        --local dockerfile=$basedir/k8s/backend/ \
        --output type=image,name=docker.io/itsos/blog-backend:$version
    sudo ctr -n buildkit i push -u itsos docker.io/itsos/blog-backend:$version
    sed -i "s#blog-backend:v.*#blog-backend:$version#" $basedir/k8s/blog-backend.yaml
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
