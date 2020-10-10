## router
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/routers/READMECN.md)
set port monitor
```go
router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
```
set router group
```go
users := router.Group("/users")
```
set middleware
```go
users.Use(middleware.Auth("user"))
```
