package post

import (
	"github.com/gin-gonic/gin"
	"go.example/blog/response"
	"net/http"
)

func Query(c *gin.Context) {
	all := QueryAll()
	if len(all) == 0 {
		c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
		return
	}
	c.JSON(http.StatusOK, response.Success("success", all))
}

func Save(c *gin.Context) {
	var newPost post
	err := c.ShouldBindJSON(&newPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}

	existingPostByTitle := QueryByTitle(newPost.Title)
	if existingPostByTitle != nil {
		c.JSON(http.StatusOK, response.Fail("标题已存在"))
		return
	}

	affectedRows, err := SavePost(newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Fail("保存失败"))
		return
	}
	c.JSON(http.StatusOK, response.Success("保存成功", affectedRows))
	return
}
