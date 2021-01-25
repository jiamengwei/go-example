package main

import (
	"github.com/gin-gonic/gin"
	"go.example/blog/category"
	"go.example/blog/post"
	"go.example/blog/response"
	"net/http"
)

func posts(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
}

func comments(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
}

func categories(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
}

func main() {
	router := gin.Default()
	defer router.Run()

	postRouter := router.Group("/posts")
	postRouter.GET("", post.Query)
	postRouter.POST("", post.Save)
	postRouter.PUT("", post.Edit)
	postRouter.DELETE("", post.Delete)

	commentRouter := router.Group("/comments")
	commentRouter.GET("", comments)

	categoryRouter := router.Group("/categories")
	categoryRouter.GET("", category.Query)
	categoryRouter.POST("", category.Save)
	categoryRouter.DELETE("", category.Delete)
}
