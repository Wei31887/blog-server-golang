# blog-server-golang

使用Gin + Vue3 + Element Plus 的前後端分離部落格 Blog 專案。採用 Gin 框架架設後端 server ，支援中間件 JWT 認證、日誌管理，資料庫採用 Postgresql 架設，關聯式管理部落客文章內容、文章分類、tags、留言等。

- **[Blog Front-end Project](https://github.com/Wei31887/blog-front-vue)**
- **[Admin Front-end Project](https://github.com/Wei31887/blog-admin-vue)**

## Features

- RESTful API 特色
- 使用 Gin 框架，提供 middleware 功能，如 blogger用戶認證、日誌、跨域等
- 基於 GORM 實現 model 介面，存取資料庫
- 基於 gRPC 實現用戶登入功能 api 
- 基於 JWT token-based 認證原理實現 admin 用戶登入功能
- 基於快取資料庫 Redis 實現登出後用戶 token 失效功能
- TODO: Dockerfile
- TODO: API unit test
- TODO: 系統上線

## 技術

1. [gin-gonic/gin](https://github.com/gin-gonic/gin) Gin 框架實現 RESTful API 及 middleware 處理
2. [jwt-go](https://github.com/dgrijalva/jwt-go) 實現 JWT 認證
3. [golang-migrate/migrate](https://github.com/golang-migrate/migrate) 實現資料庫遷移
4. [gorm.io/gorm](https://gorm.io) 實現存取資料庫
5. [viper](https://github.com/spf13/viper) 實現 config 初始化
6. [uber-go/zap](https://github.com/uber-go/zap) zap實現日誌log管理
7. [go-redis] 使用 Redis 實現 JWT 登出後 token 失效功能

## 環境

需要在local 安装 [go] [gin] [node](http://nodejs.org/) 和 [git](https://git-scm.com/) 

## Local Development

### 目錄創建

```bash
# 新建目錄
mkdir blog_golang_vue
cd blog_golang_vue
```

### Git clone

```bash
# clone server 程式碼
git clone https://github.com/Wei31887/blog-server-golang.git

# clone blog front-end 程式碼
git clone https://github.com/Wei31887/blog-front-vue.git

# clone admin front-end 程式碼
git clone https://github.com/Wei31887/blog-admin-vue.git
```

### 開始運行

#### 初始化運行

```bash
# 進入 blog-server-glang 專案後端 server 部分
cd ./server

go mod tidy

go build

# 修改資料庫配置 
# path:  config.yaml
vi config.yaml 
```

#### 以 Docker 初始化資料庫並運行程式

``` bash
# 首次須先配置資料庫
# 透過docker image 使用postgress image啟用
# 相關指令已寫入Makefile
$ make postgres

# 創建 Database
$ make createdb

# 創建 Table
$ make migrateup

# 運行main.go
$ go run main.go
```
