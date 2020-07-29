# Golang-pangu

Golang-pangu is based on [Gin](https://github.com/gin-gonic/gin), [jwt-go](https://github.com/dgrijalva/jwt-go), [gorm](https://github.com/go-gorm/gorm), it realize single device multi-segment login.

## Start

1.install postgres, redis

2.config application.yml

3.go run main.go

4.open `http://localhost:3000/ping` in web browser, and then you will get a "pong" response

## Api example

* ### sign_up

  Post `http://localhost:3002/sign_up` 

  params: email, password, password_confirm

* ### sign_in

  Post `http://localhost:3002/sign_in` 

  params: email, password, device_type

  when sign_in success, you will get a Authorization header from response

* ### auth_ping

  Get `http://localhost:3002/auth_ping` 

  params: should add a valid Authorization header to request this api

* ### change_password

  Post `http://localhost:3002/change_password` 

  params: origin_password, password, password_confirm
