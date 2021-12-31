FROM alpine:latest
MAINTAINER itsos<peng.yu@qjfu.cn>

# 设置ca-certificates允许以https方式连接外部网络；安装tzdata支持时区相关功能
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    tzdata \
    && update-ca-certificates 2>/dev/null || true

# 设置默认时区
ENV TZ Asia/Shanghai
WORKDIR /data
COPY run/backend /data
EXPOSE 9080
ENTRYPOINT ["./studynotes", "-c", "/data/config"]
