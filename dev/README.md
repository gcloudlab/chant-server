# IBB 开发手册

## 技术栈
语言：Golang 数据库：MongoDB 框架：GIN 协议：Websocket


## 依赖安装

```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
```

## 环境搭建

### 开发环境

操作系统: Windows 11, Linux (VSC Remote - Containers) 

开发环境: VSCode, Docker Desktop (MongoDB), Postman

### MongoDB (Docker)

```shell
docker run -d --network ibb-network --name ibb-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -p 27017:27017 mongo
```

## 参考

https://github.com/gorilla/websocket