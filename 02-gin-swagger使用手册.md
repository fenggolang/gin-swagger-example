#### gin-swagger 地址
```html
https://github.com/swaggo/gin-swagger
```
---
#### gin-swagger 概述
```markdown
# swag
Swag将Go注释转换为Swagger Documentation 2.0。我们为流行的Go Web(gin、echo、net/http)框
架创建了各种插件。这允许您快速与现有的Go项目集成（使用Swagger UI）。
# gin-swagger
gin-swagger是一个使用Swagger 2.0自动生成RESTful API文档的gin middleware。

Swag是一个框架，而gin-swagger是一个基于Swag框架写的一个中间件(也可以叫做插件)
```
---
#### gin-swagger　用法
```markdown
# 1. 在你的源代码中添加注释
声明注释格式:https://swaggo.github.io/swaggo.io/declarative_comments_format/
# 2. 下载Swag
go get -u -v github.com/swaggo/swag/cmd/swag
#3. 在你的项目根目录下运行swag初始化命令,注意根目录下需要有main.go文件，swag将解析注释和生成必要的文件(生成docs文件夹和docs/doc.go文件)
```go
// 初始化生成docs文档和目录
swag init
// main.go文件导入下面docs目录
_ "gin-swagger-example/celler/docs"

```
# 4. 下载gin-swagger
go get -u -v github.com/swaggo/gin-swagger
go get -u -v github.com/swaggo/gin-swagger/swaggerFiles

go get -u -v github.com/satori/go.uuid
# 然后在你的代码中导入如下头文件行
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
```
---

#### gin-swagger典型的例子
[swaggo+gin](https://github.com/swaggo/swag/tree/master/example)  
[gin-swagger-example](https://github.com/swaggo/gin-swagger/tree/master/example)
