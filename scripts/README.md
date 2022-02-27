### 部署脚本
```shell
# 运行命令可参见选项
./run.sh
# 第一次部署执行，根据提示操作
./run.sh deploy_full
```

### supervisor 管理指令
```shell
# 进程们状态
supervisorctl status
# 停止进程
supervisorctl stop monitor
# 启动进程
supervisorctl start monitor
# 重启进程
supervisorctl restart monitor
# 重新加载配置，不会启动新添加的程序
supervisorctl reread
# 重启配置文件修改过的程序
supervisorctl update
```

### nginx go sock 方式配置
```shell
upstream studynotes {
    server unix:/tmp/studynotes.sock;
    keepalive 300;
}

location / {                                                                                                                                               
    try_files $uri $uri/ /index.html;                                                                                                                      
}

location ^~ /api {
    rewrite /api/(.+)$ /$1 break;
    proxy_pass http://studynotes;
    proxy_redirect     off;
    proxy_set_header   Host             $host;
    proxy_set_header   X-Real-IP        $remote_addr;
    proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
    proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;
    proxy_max_temp_file_size 0;
    proxy_connect_timeout      90;
    proxy_send_timeout         90;
    proxy_read_timeout         90;
    proxy_buffer_size          4k;
    proxy_buffers              4 32k;
    proxy_busy_buffers_size    64k;
    proxy_temp_file_write_size 64k;
}
```

### 镜像构建
```shell
sudo buildkitd --oci-worker=false --containerd-worker=true &
sh dockerfile.sh 
```
