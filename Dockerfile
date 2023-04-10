# 使用 Golang 官方的基本映像
FROM golang:1.19-alpine as builder

# 設置工作目錄
WORKDIR /app

# 複製專案到容器中
COPY . .

# 編譯 Golang 程式
RUN go mod download && go build -o app

# 第二個階段，使用輕量級的 alpine 映像
FROM alpine:3.14

# 安裝 PostgreSQL 和 Redis 客戶端
RUN apk --no-cache add postgresql-client redis

# 複製從 builder 階段編譯的二進制檔
COPY --from=builder /app/app /app/app

# 設定環境變數
ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=mydb
ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

# 開放對外的端口
EXPOSE 8080

# 啟動 Golang Gin 後端伺服器
CMD ["/app/app"]
