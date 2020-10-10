# Golang-pangu
[English document](https://github.com/ruilisi/go-pangu/blob/master/README.md)<br>
Golang-pangu是一个用Go语言开发的一站式服务开发的技术解决方案，它整合了 JWT(JSON Web Tokens), Postgres, Redis, Docker, Gin, Ginkgo, Gorm等多项业界尖端技术，它是后端项目开发的起点，可作为开发者强有力的生产工具。

## 产品特性
Golang-pangu基于以下的工具

|名字|描述|
|------|--------|
|[Go](https://github.com/golang/go)|最近几年最为流行的新兴语言，简单的同时拥有极高的并发性能。|
|[Gin](https://github.com/gin-gonic/gin)|基于Go语言的web框架，方便灵活的中间件，强大的数据绑定，以及极高的性能|
|[Gorm](https://github.com/go-gorm/gorm)|Go语言的全功能ORM库，用于操作数据库|
|[Ginkgo](https://github.com/onsi/ginkgo)|Ginkgo是一个BDD风格的Go测试框架，旨在帮助你有效地编写富有表现力的全方位测试。|
|[JWT](https://jwt.io/)|JSON Web Tokens，是目前最流行的跨域认证解决方案。|
|[Postgres](https://www.postgresql.org/)|高性能开源数据库，当整体负载达到很高时依旧能有强大的性能|
|[Redis](https://redis.io/)|内存数据库，拥有极高的性能|
|[Docker](https://www.docker.com/)|开发、部署、运行应用的虚拟化平台|

## 整体结构
```
.
├── application.yml  
├── args 
│   ├── args.go
│   └── cmd.go
├── conf  
│   ├── conf_debug.go
│   ├── conf.go
│   └── conf_release.go
├── controller 
│   ├── application.go
│   ├── auth.go
│   ├── error.go
│   └── session.go
├── db  
│   └── db.go
├── dist
│   ├── go-pangu-amd64-debug-linux
│   └── go-pangu-amd64-release-linux
├── Dockerfile
├── go.mod
├── go.sum
├── jwt  
│   └── jwt.go
├── main.go 
├── Makefile  
├── middleware  
│   └── middleware.go
├── models  
│   ├── base_model.go
│   └── user.go
├── params  
│   └── params.go
├── README.md
├── redis 
│   └── redis.go
├── routers  
│   └── router.go
├── test 
│   ├── sign_in_test.go
│   └── test_suite_test.go
└── util
    └── util.go
```

|文件|功能|
|------|--------|
|application.yml|配置文件，包含基本信息|
|<font color=Blue>args</font>|包含获取params的函数|
|<font color=Blue>conf</font>|获取配置文件的函数|
|<font color=Blue>controller</font>|router使用的handler控件，包含各种操作具体内容|
|<font color=Blue>db</font>|db操作，像是打开数据库|
|<font color=Blue>db</font>|jwt相关内容 包含生成jwt与验证jwt的函数|
|main.go|程序主函数，执行时增加-db参数可选择不同的内容，create创建数据库，migrate更新表结构，drop删除数据库|
|<font color=Blue>middleware</font>|中间件，验证token是否正确|
|<font color=Blue>models</font>|基础的结构以及一些基本的数据库操作|
|<font color=Blue>params</font>|数据绑定的结构|
|<font color=Blue>redis</font>|包含连接redis和redis操作函数|
|<font color=Blue>router</font>|路由|
|<font color=Blue>test</font>|测试|

## 开始运行
1. 安装postgres和redis数据库
2. 配置根目录下的 **application.yml** 配置文件 
3. go run 运行 main.go
4. 在浏览器打开 `http://localhost:3000/ping` 会显示pong，表明服务成功部署



## Api 样例

* ### <div> <button style="background-color: rgb(38, 203, 124);border:0px;width:60px;height:25px;border-radius:10px;font-size:15px;color:white" >POST</button> Sign Up </div>
<pre><span style="color:grey">http://localhost:3002/users/sign_up</span><b>managers/sign_up</b></pre>
用户注册

* ### <div> <button style="background-color: rgb(38, 203, 124);border:0px;width:60px;height:25px;border-radius:10px;font-size:15px;color:white" >POST</button> Sign In </div>
<pre><span style="color:grey">http://localhost:3002/users/sign_in</span><b>managers/sign_up</b></pre>
用户登录，成功后会在头部返回Authorization，这是后续用户接口需要的token

* ### <div> <button style="background-color: rgb(56, 132, 255);border:0px;width:60px;height:25px;border-radius:10px;font-size:15px;color:white" >GET</button> Auth Ping </div>
<pre><span style="color:grey">http://localhost:3002/auth_ping</span><b>managers/sign_up</b></pre>
需要user的token的接口

* ### <div> <button style="background-color: rgb(38, 203, 124);border:0px;width:60px;height:25px;border-radius:10px;font-size:15px;color:white" >POST</button> Sign In </div>

<pre><span style="color:grey">http://localhost:3002/users/change_password</span><b>managers/Change Password</b></pre>
修改用户密码，需要user的token

## 其他公开库
[Rails-pangu](https://github.com/ruilisi/rails-pangu) 基于 **Rails 6(API Only)** 框架搭建的一站式服务开发的技术解决方案

## 使用Go搭建后端的产品
|产品|描述|
|----|-----|
|[eSheep](https://esheep.io/)|电子绵羊eSheep是一款网络加速器，它可以帮助身在海外的您极速连接中国的视频音乐网站。|
|||

## 执照

## Contributors

