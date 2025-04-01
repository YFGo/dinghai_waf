FROM golang:1.21 AS builder
WORKDIR /src

# 安装构建依赖
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# 复制整个项目到容器
COPY . /src
# 进入相应的服务目录
WORKDIR /src/wafconsole/app/user
# 使用 app_makefile 构建    
RUN cd /src/wafconsole && GOPROXY=https://goproxy.cn make -f app_makefile build APP_RELATIVE_PATH=app/user

# 第二阶段：运行环境
FROM debian:bookworm-slim

# 安装运行时依赖 + 时区
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    netbase \
    tzdata \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get autoremove -y && apt-get autoclean -y

# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 复制二进制文件
COPY --from=builder /src/bin /app

WORKDIR /app
EXPOSE 8002 9002
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]