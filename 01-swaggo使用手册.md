#### swaggo 地址
```html
https://github.com/swaggo/swag
```
---
#### swaggo 概述
```markdown
Swag将Go注释转换为Swagger Documentation 2.0。我们为流行的Go Web(gin、echo、net/http)框
架创建了各种插件。这允许您快速与现有的Go项目集成（使用Swagger UI）。
```
---
#### swaggo 例子
例子：
[swaggo+gin](https://github.com/swaggo/swag/tree/master/example)
---
#### swaggo 入门
1.在API源代码中添加注释  
[请参阅声明注释格式](https://github.com/swaggo/swag#general-api-info)  
2.下载swag
```markdown
go get -u -v github.com/swaggo/swag/cmd/swag
```
3.运行swag初始化命令
```markdown
# 在swag int项目的根文件夹中运行，该文件夹包含该main.go文件，这将解析你的注释并生成所需的文件(docs文件夹和docs.go文件)
swag init
```
4.使用swag支持的插件之一,需要满足以下添加之一
```markdown
# a
确保导入生成的内容，docs/docs.go以便进行特定配置init,
# b
如果你的普通API注释不存在main.go,你可以通过为swag加-g参数明确指定
swag init -g http/api.go
```
---
#### swaggo 目前支持的框架
- [gin](https://github.com/swaggo/gin-swagger)
- [echo](https://github.com/swaggo/echo-swagger)
- [net/http](https://github.com/swaggo/http-swagger)