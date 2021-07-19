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

安装canal
```shell
mkdir -p /usr/local/canal && cd /usr/local/canal/ && \
    wget https://github.com/alibaba/canal/releases/download/canal-1.1.5/canal.deployer-1.1.5.tar.gz && \
    tar zxvf canal.deployer-1.1.5.tar.gz

# 配置
# conf/canal.properties
...

# conf/example/instance.properties
# 不能与从库的其他id重复
canal.instance.mysql.slaveId = 1234
# mysql 地址/slave账号/密码
canal.instance.master.address=127.0.0.1:3306
canal.instance.dbUsername=canal
canal.instance.dbPassword=)(*cdgasf,23
canal.instance.connectionCharset = UTF-8
```
> 解释zookeeper实现canal HA机制的好文：https://segmentfault.com/a/1190000023297973

授权 canal 连接mysql账号具备作为mysql slave的权限
```mysql
CREATE USER canal IDENTIFIED BY ')(*cdgasf,23';
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
FLUSH PRIVILEGES;
```

