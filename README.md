# IBB 聊天社区

IBB 即时通讯后台

技术栈: Go (GIN), MongoDB, Websocket

## 调试



## 核心包

https://github.com/gorilla/websocket

## 扩展安装

```shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go
go get github.com/jordan-wright/email
```

## Docker 安装 mongoDB

```shell
docker run -d --network some-network --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```

## 功能列表

### 1.用户模块
 
- [x] 密码登录
- [x] 邮箱注册(验证码)
- [x] 用户详情

### 2.通讯模块
 
- [x] 一对一通讯
- [x] 多对多通讯
- [x] 消息列表
- [x] 聊天记录列表

## License

[GPL](LICENSE)