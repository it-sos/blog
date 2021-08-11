# study notes

### 介绍
学习笔记 - blog

### 安装依赖
#### Mysql 全文索引 elasticsearch + kibana + canal + zookeeper + kafka
> kibana: https://www.elastic.co/guide/en/kibana/current/deb.html

安装 elasticsearch kibana
```shell
wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
sudo apt-get install apt-transport-https
echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-7.x.list
sudo apt-get update && sudo apt-get install elasticsearch kibana
```

启动 elasticsearch
```shell
# 因为 elasticsearch 的运行不能是root，因此可以设置一下用户再来运行
sudo chown -R guest /usr/share/elasticsearch/ /etc/elasticsearch/ /etc/default/elasticsearch /var/log/elasticsearch/

# 安装中文分词器
sudo ./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.13.3/elasticsearch-analysis-ik-7.13.3.zip

# 运行
/usr/share/elasticsearch/bin/elasticsearch -d # -d 作为守护进程后台运行

# port: 9200 ：是http访问端口返回json格式服务的基本信息
# port: 9300 ：是提供服务的端口，譬如canal-adapter 可以直连实现数据同步
```

启动 kibana
```shell
nohup /usr/share/kibana/bin/kibana serve > /tmp/kibana.log 2>&1 &
```

访问
> kibana: http://localhost:5601
>
> elasticsearch: http://localhost:9200

安装 zookeeper
```shell
wget https://downloads.apache.org/zookeeper/zookeeper-3.7.0/apache-zookeeper-3.7.0-bin.tar.gz && \
    tar zxvf tar zxvf apache-zookeeper-3.7.0-bin.tar.gz && \
    mv zookeeper-3.7.0 /usr/local/zookeeper && \
    sudo tee -a /etc/profile <<EOF
    # ZooKeeper Env
    export ZOOKEEPER_HOME=/usr/local/zookeeper
    export PATH=$PATH:$ZOOKEEPER_HOME/bin
    EOF

source /etc/profile

# 初次运行修改配置
mv $ZOOKEEPER_HOME/conf/zoo_sample.cfg $ZOOKEEPER_HOME/conf/zoo.cfg

# 创建文件
mkdir /usr/local/zookeeper/data /usr/local/zookeeper/logs

# 可以直接覆盖以下配置
sudo tee /usr/local/zookeeper/zoo.cfg <<EOF
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/usr/local/zookeeper/data
dataLogDir=/usr/local/zookeeper/logs
clientPort=2181
#maxClientCnxns=60
#autopurge.snapRetainCount=3
#autopurge.purgeInterval=1
#metricsProvider.className=org.apache.zookeeper.metrics.prometheus.PrometheusMetricsProvider
#metricsProvider.httpPort=7000
#metricsProvider.exportJvmInfo=true
EOF

# 承上，如果是多节点，增加如下配置
server.1=192.168.1.110:2888:3888
server.2=192.168.1.111:2888:3888
server.3=192.168.1.112:2888:3888
#master
echo "1">/usr/local/zookeeper/data/myid
#slave1
echo "2">/usr/local/zookeeper/data/myid
#slave2
echo "3">/usr/local/zookeeper/data/myid

cd $ZOOKEEPER_HOME/bin
# 启动zookeeper
./zkServer.sh  start
# 查看状态
./zkServer.sh  status
# 停用服务
./zkServer.sh  stop

# 验证 zookeeper 服务
telnet 127.0.0.1 2181
```

安装 Kafka
```shell
wget https://downloads.apache.org/kafka/2.8.0/kafka_2.13-2.8.0.tgz && \
    mkdir  -p /usr/local/kafka && \
    cp kafka_2.13-2.8.0.tgz  /usr/local/kafka && \
    tar -zxvf kafka_2.13-2.8.0.tgz -C /usr/local/kafa && \
    rm -f kafka_2.13-2.8.0.tgz 

# 修改配置文件
vim /usr/local/kafka/kafka_2.13-2.8.0/config/server.properties
# 修改参数
zookeeper.connect=192.168.1.110:2181
listeners=PLAINTEXT://:9092
advertised.listeners=PLAINTEXT://192.168.1.117:9092 # 本机ip
# ...

# 启动
cd /usr/local/kafka/kafka_2.13-2.8.0/ && \
    bin/kafka-server-start.sh  -daemon  config/server.properties

# 查看所有topic
cd /usr/local/kafka/kafka_2.13-2.8.0/ && \
    bin/kafka-topics.sh --list --zookeeper 192.168.1.110:2181
```

开启mysql binlog
```shell
# 修改mysql配置文件/etc/mysql/my.cnf，后重启mysqld
[mysqld]
log-bin=mysql-bin
binlog-format=ROW
server-id=1

# 验证
mysql> show variables like 'log_bin%'
mysql> show variables like 'binlog_format%'

# 授权 canal 连接mysql账号具备作为mysql slave的权限
CREATE USER canal IDENTIFIED BY ')(*cdgasf,23';
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
# 如果是 mysql 8.0 改变一下默认的认证方式
ALTER USER 'canal'@'%' IDENTIFIED WITH mysql_native_password BY ')(*cdgasf,23'; #更新一下用户密码
FLUSH PRIVILEGES;
```

安装canal
```shell
mkdir -p /usr/local/canal && cd /usr/local/canal/ && \
    wget https://github.com/alibaba/canal/releases/download/canal-1.1.5/canal.deployer-1.1.5.tar.gz && \
    tar zxvf canal.deployer-1.1.5.tar.gz

# 配置
# conf/canal.properties

# conf/example/instance.properties
# 不能与从库的其他id重复
canal.instance.mysql.slaveId = 1234
# mysql 地址/slave账号/密码
canal.instance.master.address=127.0.0.1:3306
canal.instance.dbUsername=canal
canal.instance.dbPassword=)(*cdgasf,23
canal.instance.connectionCharset = UTF-8
canal.instance.filter.regex=study_notes.* # 过滤的表名
```

> 解释zookeeper实现canal HA机制的好文：https://segmentfault.com/a/1190000023297973

安装 canal.adapter
```shell
mkdir -p /usr/local/canal-adapter && cd /usr/local/canal-adapter && \
    wget https://github.com/alibaba/canal/releases/download/canal-1.1.5/canal.adapter-1.1.5.tar.gz && \
    tar zxvf canal.adapter-1.1.5.tar.gz

# 配置
# conf/local/canal-adapter/application.yml
```
数据全量同步
> curl -XPOST http://127.0.0.1:8081/etl/es7/canal.yml
