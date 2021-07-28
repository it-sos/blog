### 配置放置到 ./config.rc ，若不创建，运行脚本将自行创建
```shell
tee ./config.rc <<EOF
USER=user
PASSWORD=password
# 项目远程目录
GO_PROJ_DIST=/go_proj_dist
# 静态资源远程目录
STATIC_PROJ_DIST=/static_proj_dist
EOF
```

### supervisor 部署方案（适用 debian）
> 1. 通过 sftp 上送 supervisor 相关的配置文件到指定目录
> 1. 通过 ssh 执行远程命令，检查软件环境 supervisor，不存在则安装，存在则执行部署
```shell
./supervisor.sh
```

### study-notes 常规部署方案
> 1. 生产配置放置到 ./run/config.yaml
> 1. 本地生成可执行文件和静态资源包，通过 sftp 上送到相关目录
> 1. 通过 ssh 执行远程命令，重启 nginx 服务
```shell
./run.sh
```

