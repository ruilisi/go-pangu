## router
设置本地监听端口
```go
router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
```
设置路由组
```go
users := router.Group("/users")
```
设置中间件
```go
users.Use(middleware.Auth("user"))
```

