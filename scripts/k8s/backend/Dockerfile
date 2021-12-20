FROM alpine:latest
MAINTAINER itsos<peng.yu@qjfu.cn>
WORKDIR /data
COPY run /data
EXPOSE 9080
ENTRYPOINT ["./studynotes", "-c", "/data/config"]
