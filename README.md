# blog-server-golang

使用Gin + Vue3 + Element Plus 的前後端分離部落格 Blog 專案，外加Admin後台管理系統，集中管理部落客文章內容、文章分類、tags、留言管理。

- [Blog Front-end Project](https://github.com/Wei31887/blog-front-vue)
- [Admin Front-end Project](https://github.com/Wei31887/blog-admin-vue)

## Features

- RESTful API 特色
- 使用 Gin 框架，提供 middleware 功能，如 admin用戶認證、日誌、跨域等
- 基於 JWT 認證的 Admin 用戶登入功能
- 基於 GORM 存取資料庫
- TODO: Dockerfile
- TODO: API unit test
- TODO: 系統上線

## 技術

1. Gin 框架實現 RESTful API 及 middleware 處理
2. jwt-go 實現 JWT 認證
3. go-migrate 實現資料庫遷移
4. gorm 實現存取資料庫
5. viper 實現 config 初始化
6. zap 實現日誌log

## 環境

需要在local 安装 [go] [gin] [node](http://nodejs.org/) 和 [git](https://git-scm.com/) 

## Local Development

### 目錄創建

```bash
# 新建目錄
mkdir blogserver
cd blogserver
```

### Git clone

```bash
# clone server 程式碼
git clone https://github.com/Wei31887/blog-server-golang.git

```

### 開始運行

#### 初始化運行

```bash
# 進入 blog-server-glang 後端
cd ./blog-server-glang

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
