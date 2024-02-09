# 使用一个基础的镜像，可以根据你的项目需要选择合适的基础镜像
FROM golang:1.21.5-bookworm

# 设置工作目录
WORKDIR /app

# 将本地文件复制到容器中
COPY . .

# 在容器中执行安装依赖等操作，根据你的项目需要修改
# RUN go mod download

# 构建你的应用
# RUN go build -o myapp

# 暴露端口，如果你的应用需要监听端口的话，根据实际情况修改
# EXPOSE 8080

# 定义容器启动时执行的命令
# CMD ["./myapp"]
