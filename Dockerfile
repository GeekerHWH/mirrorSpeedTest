# 使用一个基础的镜像，可以根据你的项目需要选择合适的基础镜像
FROM golang:1.22.0-alpine3.19

# 设置工作目录
WORKDIR /app

# 将本地文件复制到容器中
COPY . .

# 构建应用
RUN go build -o mirrorSpeedTest main/main.go

# 定义容器启动时执行的命令
CMD ["./mirrorSpeedTest -h"]
