# Chant 开发手册

## 技术栈

语言：Golang 数据库：MongoDB 框架：GIN 协议：Websocket

## 核心依赖

```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
go get github.com/jordan-wright/email
```

## 开发环境

操作系统: Windows 11, Linux (VSC Remote - Containers) 

开发环境: VSCode, Docker Desktop (MongoDB), Postman

### MongoDB (Docker)

```shell
docker run -d --network some-network --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```

### 调试

- API服务端口：8088
- Socket服务端口：8089


## 参考

https://github.com/gorilla/websocket