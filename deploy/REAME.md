### 部署脚本
```shell
# 第一次部署执行
./run.sh first_linux 
```

### nginx go sock 方式配置
```shell
location / {
    # config.yaml 参考此项配置 sock:/tmp/studynotes.sock
    uwsgi_pass unix:///tmp/studynotes.sock;
    uwsgi_read_timeout 600;
    include /etc/nginx/uwsgi_params;
}
```