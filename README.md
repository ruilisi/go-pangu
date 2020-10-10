# Golang-pangu
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/READMECN.md)

Go-pangu is a Go boilerplate which follows cutting-edge solutions already adopted by the industry,  JWT(JSON Web Tokens), Postgres, Redis, Docker,  Gin, Ginkgo, Gorm. It is a solid production-ready starting point for your new backend projects.

## Features
Golang-pangu is based on following tools

|name|description|
|------|--------|
|[Go](https://github.com/golang/go)|an open source programming language that makes it easy to build simple, reliable, and efficient software.|
|[Gin](https://github.com/gin-gonic/gin)|web struct based on Go,Flexible middleware，strong data bind and high performance.|
|[Gorm](https://github.com/go-gorm/gorm)|The fantastic ORM library for Golang aims to be developer friendly.|
|[Ginkgo](https://github.com/onsi/ginkgo)|Ginkgo builds on Go's testing package, allowing expressive Behavior-Driven Development ("BDD") style tests.|
|[JWT](https://jwt.io/)|JSON Web Tokens.an open, industry standard RFC 7519 method for representing claims securely between two parties.|
|[Postgres](https://www.postgresql.org/)|The World's Most Advanced Open Source Relational Database|
|[Redis](https://redis.io/)|an open source (BSD licensed), in-memory data structure store, used as a database, cache and message broker.|
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
|<font color=Blue>args</font>|include functions which can get params from request url|
|<font color=Blue>conf</font>|include functions which can get config file content|
|<font color=Blue>controller</font>|handlers|
|<font color=Blue>db</font>|database operate|
|<font color=Blue>db</font>|include create and verify jwt fuction|
|main.go|main function, with db param|
|<font color=Blue>middleware</font>|middleware|
|<font color=Blue>models</font>|base models |
|<font color=Blue>params</font>|struct used in data bind|
|<font color=Blue>redis</font>|redis operate functions|
|<font color=Blue>router</font>|router|
|<font color=Blue>test</font>|test|


## Start

1. install postgres, redis
2. config application.yml
3. go run main.go
4. open `http://localhost:3000/ping` in web browser, and then you will get a "pong" response

## Api examples

* ### sign_up

  Post `http://localhost:3002/users/sign_up`

  params: email, password, password_confirm

* ### sign_in

  Post `http://localhost:3002/users/sign_in`

  params: email, password

  when sign_in success, you will get a Authorization header from response

* ### auth_ping

  Get `http://localhost:3002/auth_ping`

  params: should add a valid Authorization header to request this api

* ### change_password

  Post `http://localhost:3002/users/change_password`

  params: origin_password, password, password_confirm


## other public library
  [Rails-pangu](https://github.com/ruilisi/rails-pangu) is a Rails 6(API Only) boilerplate which follows cutting-edge solutions already adopted by the industry, notablly, Devise, JWT(JSON Web Tokens), Postgres, Redis, Docker, Rspec, RuboCop, CircleCI. It is a solid production-ready starting point for your new backend projects.

## Projects using Go-pangu
  |product|description|
  |----|-----|
  |[eSheep](https://esheep.io/)|Network booster which helps global users access better entertainment content from China.|
  |||

## License

## Contributors
