### supervisorctl管理
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