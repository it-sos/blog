#!/bin/bash

basedir=$(cd "$(dirname "$0")" && pwd)
. $basedir/common.sh
PROJECT=$basedir/run/$PROJECT_NAME


options="{full|backend}"

run() {
  case $1 in
  "full")
  ;;
  "backend")
    run build_linux
  ;;
  "build_linux")
    build linux
    mkdir -p $basedir/run/web && cp -r $basedir/../src/backend/web/views run/web
    buildctl build \
        --frontend=dockerfile.v0 \
        --local context=. \
        --local dockerfile=$basedir/k8s/backend/ \
        --output type=image,name=docker.io/itsos/blog_backend:v1.0.4,push=true
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
