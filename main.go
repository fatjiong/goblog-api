package main

import (
	"github.com/fatjiong/goblog-api/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	//文章列表
	router.GET("/article",controllers.ArticleListGet)
	router.GET("/article/:id",controllers.ArticleDetailGet)
	router.DELETE("/article/:id",controllers.ArticleDelete)
	router.PUT("/article/:id",controllers.ArticlePut)
	router.Run(":8080")
}
