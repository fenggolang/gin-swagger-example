package main

import (
	_ "gin-swagger-example/docs" // docs由Swag 命令行生成，必须导入它
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"              // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // 植入swagger页面html文件
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// 创建路由
	//router := gin.New()
	router := gin.Default()

	//使用ginSwagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动一个http server
	// url: http://localhost:8080/swagger/index.html
	router.Run(":8080")
}
