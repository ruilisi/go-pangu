## test 
### 安装
`go get -u github.com/onsi/ginkgo/ginkgo`

### 使用说明
直接运行<br>
ginkgo 测试当前文件下所有文件<br>
ginkgo -focus=SignIn 测试名为SignIn的Descibe下的所有测试
```go
var _ = Describe("SignIn", func() {
...  
})
```
