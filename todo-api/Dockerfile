# 1) Builder 階段：編譯 Go 程式
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o todo-api .

# 2) 執行階段：只留執行環境
FROM alpine:latest

# 建立非 root 使用者
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# 複製編譯好的二進位檔
COPY --from=builder /app/todo-api .

# 暴露端口
EXPOSE 8080

# 切換到非 root 使用者，最後执行
USER appuser

# 啟動程式
ENTRYPOINT ["./todo-api"]
