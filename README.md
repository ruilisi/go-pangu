# Golang-pangu
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/READMECN.md)

Go-pangu is a Go boilerplate which follows cutting-edge solutions already adopted by the industry,  JWT(JSON Web Tokens), Postgres, Redis, Docker,  Gin, Ginkgo, Gorm. It is a solid production-ready starting point for your new backend projects.

## Features
Golang-pangu is based on following tools

|name|description|
|------|--------|
|[Go](https://github.com/golang/go)|an open source programming language that makes it easy to build simple, reliable, and efficient software.|
|[Gin](https://github.com/gin-gonic/gin)|web struct based on Go, flexible middleware，strong data binding and outstanding performance.|
|[Gorm](https://github.com/go-gorm/gorm)|The fantastic ORM library for Golang aims to be developer friendly.|
|[Ginkgo](https://github.com/onsi/ginkgo)|Ginkgo builds on Go's testing package, allowing expressive Behavior-Driven Development ("BDD") style tests.|
|[JWT](https://jwt.io/)|JSON Web Tokens. An open, industry standard RFC 7519 method for representing claims securely between two parties.|
|[Postgres](https://www.postgresql.org/)|The world's most advanced open source relational database|
|[Redis](https://redis.io/)|An open source (BSD licensed), in-memory data structure store, used as a database, cache and message broker.|
|[Docker](https://www.docker.com/)|Docker is a tool designed to make it easier to create, deploy, and run applications by using containers.|

## Struct
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

|file|function|
|------|--------|
|application.yml|config file|
|[args](https://github.com/ruilisi/go-pangu/tree/master/args)|functions which can fetch params from request url|
|[conf](https://github.com/ruilisi/go-pangu/tree/master/conf)|functions which can get configurations|
|[controller](https://github.com/ruilisi/go-pangu/tree/master/controller)|handlers|
|[db](https://github.com/ruilisi/go-pangu/tree/master/db)|database operations like migrating database|
|[jwt](https://github.com/ruilisi/go-pangu/tree/master/jwt)|fuctions to create and check jwt|
|main.go|main function.Call function with "--db" parameter, "create" to create database, "migrate" to migrate tables, "dorp" to delete database|
|[middleware](https://github.com/ruilisi/go-pangu/tree/master/middleware)|middleware|
|[models](https://github.com/ruilisi/go-pangu/tree/master/models)|base models and basic database operations|
|[params](https://github.com/ruilisi/go-pangu/tree/master/params)|struct used in data binding|
|[redis](https://github.com/ruilisi/go-pangu/tree/master/redis)|redis connection and operations|
|[router](https://github.com/ruilisi/go-pangu/tree/master/routers)|router|
|[test](https://github.com/ruilisi/go-pangu/tree/master/test)|test|


## Start

1. install postgres, redis
2. config application.yml
3. go run main.go
4. open `http://localhost:3002/ping` in web browser, and then you will get a "pong" response

## Api examples

* ### sign_up

  Post `http://localhost:3002/users/sign_up`

  params: email, password, password_confirm

  Register user

* ### sign_in

  Post `http://localhost:3002/users/sign_in`

  params: email, password

  You will get a header with authorization parameter from response after logging in successfully

* ### auth_ping

  Get `http://localhost:3002/auth_ping`

  Should add a valid user token to request this api

* ### change_password

  Post `http://localhost:3002/users/change_password`

  params: origin_password, password, password_confirm

  Modify user's password, which needs authorization


## other public library
  [Rails-pangu](https://github.com/ruilisi/rails-pangu) is a Rails 6(API Only) boilerplate which follows cutting-edge solutions already adopted by the industry, notablly, Devise, JWT(JSON Web Tokens), Postgres, Redis, Docker, Rspec, RuboCop, CircleCI. It is a solid production-ready starting point for your new backend projects.

## Projects using Go-pangu
  |product|description|
  |----|-----|
  |[eSheep](https://esheep.io/)|Network booster which helps global users access better entertainment content from China.|
  |[cs-server](https://excitingfrog.gitbook.io/im-api/)|agent server（Comming soon）|

## License
Code and documentation copyright 2020 the [Golang-pangu Authors](https://github.com/ruilisi/go-pangu/graphs/contributors) and [ruilisi Network](https://ruilisi.co/) Code released under the [MIT License](https://github.com/ruilisi/go-pangu/blob/master/LICENSE).
<table frame=void>
<tr>
<td >
<img src="logo.png" width="100px;" alt="hophacker"/>
</td>
</tr>
</table>

## Contributors

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->

<table>
  <tr>
    <td align="center"><a href="https://paiyou.co/"><img src="https://avatars2.githubusercontent.com/u/3121413?v=4" width="100px;" alt="hophacker"/><br /><sub><b>hophacker</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=hophacker" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=hophacker" title="Documentation">📖</a> <a href="#infra-hophacker" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
    <td align="center"><a href="https://github.com/caibiwsq"><img src="https://avatars0.githubusercontent.com/u/37767017?v=4" width="100px;" alt="caibiwsq"/><br /><sub><b>caibiwsq</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=caibiwsq" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=caibiwsq" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Ganggou"><img src="https://avatars1.githubusercontent.com/u/41427297?s=400&u=5cc6b0dfa214bc5671f849b3ee94acf597c2c6f4&v=4" width="100px;" alt="Ganggou"/><br /><sub><b>Ganggou</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Ganggou" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Ganggou" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/ExcitingFrog"><img src="https://avatars2.githubusercontent.com/u/25655802?s=460&u=23017079e78e3c3bfa57a14bc369607b1b23c470&v=4" width="100px;" alt="ExcitingFrog"/><br /><sub><b>ExcitingFrog</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=ExcitingFrog" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=ExcitingFrog" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Leo7991"><img src="https://avatars1.githubusercontent.com/u/67139714?s=460&u=278212a0d4d8ca824219adcd932dc85d2fd5ae24&v=4" width="100px;" alt="Leo7991"/><br /><sub><b>Leo7991</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Leo7991" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Leo7991" title="Documentation">📖</a></td>
        <td align="center"><a href="https://github.com/Daxigua443"><img src="https://avatars1.githubusercontent.com/u/62984061?s=460&u=375eab6d59b2087058c1a30210f8646281971ff7&v=4" width="100px;" alt="Daxigua443"/><br /><sub><b>Daxigua443</b></sub></a><br /><a href="https://github.com/ruilisi/golang-pangu/commits?author=Daxigua443" title="Code">💻</a> <a href="https://github.com/ruilisi/golang-pangu/commits?author=Daxigua443" title="Documentation">📖</a></td>

  </tr>
</table>


<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

