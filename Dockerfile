FROM golang:alpine AS builder

# 构建可执行文件
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

WORKDIR /build
ADD . .
RUN go build -o main


FROM alpine
WORKDIR /app
COPY --from=builder /build/main /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add tzdata
CMD ["./main"]