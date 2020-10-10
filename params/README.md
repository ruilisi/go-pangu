## params
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/params/READMECN.md)<br>
this folder contain struct which used in data shouldbind.<br>
though the function `Param()` in args folder can get parameters. Struct binding can work better when json data contain many params<br>
```go
if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
```


```go
type SignIn struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
```
json,form,xml settings can bind data in different methods
