## params
[中文文档](https://github.com/ruilisi/go-pangu/blob/master/params/READMECN.md)
this folder contain struct which used in data shouldbind.
though param function in args folder can get param,struct bind can work better when json contain many params
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
