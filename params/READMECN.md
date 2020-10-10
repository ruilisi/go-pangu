## params
存放数据绑定所需要的数据格式，虽然可以使用args下的Param函数。但当数据比较多时建议使用shouldbind数据绑定
```go
if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
```

其内容为
```go
type SignIn struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
```
json,form,xml 设置不同的获取方式

