AST Tool
---
试制AST分析工具，输出结构化的文件静态结构。使用AST库。
# 需求分析

需要提取的类别:
1. 包名`package main`(Red Highlight)
2. import list(Red Highlight)
3. 变量/常量赋值语句`var version = ""`(Blue Highlight)
4. 函数实现(函数名称和形参表)`func bail(err error)`(Green Highlight)
5. 复合赋值语句`ln, err := net.Listen("unix", path)`(Yellow Highlight)
6. **函数/方法调用**(调用名称,实参名称)(Pink Highlight)
    1. 常规调用,比如`fmt.Println("hello")`
    2. 复合语句中的调用.比如`unixconn, ok := conn.(*net.UnixConn)`
    3. 特殊调用,比如`return errors.New("failed to cast to unixconn")`

每个类别都需要提取的信息:
1. 文件名
2. 类别
3. 位置(方便定位)

# 需要解决的问题
- [ ] 分析深度.是否需要追到系统库中去.
- [ ] 性能优化.过于复杂的系统能否运行.
- [ ] 跨package的分析和处理问题.
- [ ] 功能扩展和交互优化.
- [ ] 模块的整理和重构.明显不能都扔在一个package里.
- [ ] 代码整理,标准化和可扩展.前期先只写一个Demo.