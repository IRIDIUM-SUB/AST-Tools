Examples
---
这里是针对`runc`中一个比较小的模块,`runc/contrib/cmd/recvtty/recvtty.go`的一个原理验证测试,比较简单,没有很多的工程层级.但是作为原理验证,具有`arg parse`功能并且确定了完整工具中的输入/输出格式.
可以参考[AST-Viewer](https://github.com/yuroyoro/goast-viewer)的写法.主要是做一个数据传输格式的设计.
NOTE:
导入第三方库后运行`go mod tidy`
`Logrus`的日志级别:
```go

log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
log.Fatal("Bye.")   // 记完日志后会调用os.Exit(1) 
log.Panic("I'm bailing.")  // 记完日志后会调用 panic()
```
结构化记录日志应该如下所示:
```go
log.WithFields(log.Fields{
  "event": event,
  "topic": topic,
  "key": key,
}).Fatal("Failed to send event")
```
# 代码类型清单
见`go/ast`的`type`类别.
- [x] Return(`*ast.ReturnStmt`)
- [x] Function Declaration(`FuncDecl`)
- [ ] Expressions(`ExprStmt`)
- [x] Array(`ArrayType`)
