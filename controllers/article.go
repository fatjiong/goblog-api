package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleListGet(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"获取列表成功",
		"data":nil,
	})
}


func ArticleDetailGet(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"获取详情成功,id："+id,
		"data":nil,
	})
}

func ArticleDelete(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"删除详情成功,id："+id,
		"data":nil,
	})
}

func ArticlePut(c *gin.Context){
	id := c.PostForm("id")
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"编辑详情成功,id："+id,
		"data":nil,
	})
}