#### gin-swagger介绍
```markdown
 使用Swagger 2.0自动生成RESTful API文档的gin中间件。
 ```
---

#### 安装gin-swagger
如果下载包有问题就去这个网站下载打包解压安装  　　　　
https://www.golangtc.com/download/package
```markdown
# 需要手动clone到$GOPATH/src/golang.org/x目录
git clone https://github.com/golang/text.git  
git clone https://github.com/golang/net.git  
git clone https://github.com/golang/tools.git

# 使用以下命令下载Swag for Go
# https://github.com/swaggo/swag
go get github.com/swaggo/swag/cmd/swag

# 使用以下方法下载gin-swagger
# https://github.com/swaggo/gin-swagger
go get -u -v github.com/swaggo/gin-swagger  
go get -u -v github.com/swaggo/gin-swagger/swaggerFiles
```
---

#### go的热更新库下载
```bash
go get -u -v github.com/pilu/fresh
```