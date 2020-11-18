# Golang-pangu

[简明教程](https://github.com/ruilisi/go-pangu/blob/master/document/%E7%AE%80%E6%98%8E%E6%95%99%E7%A8%8B.md)<br>
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
|[args](https://github.com/ruilisi/go-pangu/tree/master/args)|包含获取url的params的函数|
|[conf](https://github.com/ruilisi/go-pangu/tree/master/conf)|获取配置文件的函数|
|[controller](https://github.com/ruilisi/go-pangu/tree/master/controller)|router使用的handler控件，包含各种操作具体内容|
|[db](https://github.com/ruilisi/go-pangu/tree/master/db)|db操作，像是打开数据库|
|[jwt](https://github.com/ruilisi/go-pangu/tree/master/jwt)|jwt相关内容 包含生成jwt与验证jwt的函数|
|main.go|程序主函数，执行时增加-db参数可选择不同的内容，create创建数据库，migrate更新表结构，drop删除数据库|
|[middleware](https://github.com/ruilisi/go-pangu/tree/master/middleware)|中间件，验证token是否正确|
|[models](https://github.com/ruilisi/go-pangu/tree/master/models)|基础的结构以及一些基本的数据库操作|
|[params](https://github.com/ruilisi/go-pangu/tree/master/params)|数据绑定的结构|
|[redis](https://github.com/ruilisi/go-pangu/tree/master/redis)|包含连接redis和redis操作函数|
|[router](https://github.com/ruilisi/go-pangu/tree/master/routers)|路由|
|[test](https://github.com/ruilisi/go-pangu/tree/master/test)|测试|

## 开始运行
1. 安装postgres和redis数据库
2. 配置根目录下的 **application.yml** 配置文件
3. go run 运行 main.go
4. 在浏览器打开 `http://localhost:3002/ping` 会显示pong，表明服务成功部署



## Api 样例

* ### sign_up
  Post `http://localhost:3002/users/sign_up`

  params: email, password, password_confirm

  用户注册

* ### sign_in
    Post `http://localhost:3002/users/sign_in`

    params: email, password

  用户登录，成功后会在头部返回Authorization，这是后续用户接口需要的token

* ### auth_ping
    Get `http://localhost:3002/auth_ping`

  需要user的token的接口

* ### change_password
    Post `http://localhost:3002/users/change_password`

  修改用户密码，需要user的token

## 其他公开库
[Rails-pangu](https://github.com/ruilisi/rails-pangu) 基于 **Rails 6(API Only)** 框架搭建的一站式服务开发的技术解决方案

## 使用Go搭建后端的产品
|产品|描述|
|----|-----|
|[eSheep](https://esheep.io/)|电子绵羊eSheep是一款网络加速器，它可以帮助身在海外的您极速连接中国的视频音乐网站。|
|[cs-server](https://excitingfrog.gitbook.io/im-api/)|客服服务（未来上线）|
|soda-server|未来上线|

## 执照


代码和文档版权归2019年[MIT许可](https://github.com/ruilisi/go-pangu/blob/master/LICENSE)下发布的[Golang-pangu Authors](https://github.com/ruilisi/go-pangu/graphs/contributors) 和 [Ruilisi Network](https://ruilisi.co/)所拥有。
<table frame=void>
<tr>
<td >
<img src="logo.png" width="100px;" alt="ruilisi"/>
</td>
</tr>
</table>

## Contributors
致谢 ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->

<table>
  <tr>
    <td align="center"><a href="https://github.com/hophacker"><img src="https://avatars2.githubusercontent.com/u/3121413?v=4" width="100px;" alt="hophacker"/><br /><sub><b>hophacker</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=hophacker" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=hophacker" title="Documentation">📖</a> <a href="#infra-hophacker" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
    <td align="center"><a href="https://github.com/tony2100"><img src="https://avatars0.githubusercontent.com/u/37767017?v=4" width="100px;" alt="tony"/><br /><sub><b>Tony</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=tony2100" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=tony2100" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Ganggou"><img src="https://avatars1.githubusercontent.com/u/41427297?s=400&u=5cc6b0dfa214bc5671f849b3ee94acf597c2c6f4&v=4" width="100px;" alt="Ganggou"/><br /><sub><b>Ganggou</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Ganggou" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Ganggou" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/ExcitingFrog"><img src="https://avatars2.githubusercontent.com/u/25655802?s=460&u=23017079e78e3c3bfa57a14bc369607b1b23c470&v=4" width="100px;" alt="Xingo"/><br /><sub><b>ExcitingFrog</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=ExcitingFrog" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=ExcitingFrog" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Leo7991"><img src="https://avatars1.githubusercontent.com/u/67139714?s=460&u=278212a0d4d8ca824219adcd932dc85d2fd5ae24&v=4" width="100px;" alt="Leo7991"/><br /><sub><b>Leo7991</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Leo7991" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Leo7991" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Daxigua443"><img src="https://avatars1.githubusercontent.com/u/62984061?s=460&u=375eab6d59b2087058c1a30210f8646281971ff7&v=4" width="100px;" alt="Daxigua443"/><br /><sub><b>Daxigua443</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Daxigua443" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Daxigua443" title="Documentation">📖</a></td>
          <td align="center"><a href="https://github.com/Soryu23"><img src="https://avatars0.githubusercontent.com/u/67567977?s=460&u=fea632ad315bcdcfeff4de7ac5e2482b249929ac&v=4" width="100px;" alt="Soryu23"/><br /><sub><b>Soryu23</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Soryu23" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Soryu23" title="Documentation">📖</a></td>

  </tr>
</table>
<!-- ALL-CONTRIBUTORS-LIST:END -->

该项目遵循[贡献者](https://github.com/all-contributors/all-contributors)规范。欢迎任何形式的捐助！

