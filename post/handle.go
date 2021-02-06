package post

import (
	"github.com/gin-gonic/gin"
	"go.example/blog/category"
	"go.example/blog/response"
	"log"
	"net/http"
	"strconv"
)

func Query(c *gin.Context) {
	qTitle := c.Query("q")
	all := QueryAll(qTitle)
	if len(all) == 0 {
		c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
		return
	}
	c.JSON(http.StatusOK, response.Success("success", all))
}

func FindByCategory(c *gin.Context) {
	qCategoryName := c.Param("name")
	postsByCategory := QueryByCategory(qCategoryName)
	if postsByCategory == nil || len(postsByCategory) == 0 {
		c.JSON(http.StatusOK, response.Success("success", response.EmptySlice()))
		return
	}
	c.JSON(http.StatusOK, response.Success("success", postsByCategory))
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

	existingCategoryById := category.QueryById(newPost.CategoryId)
	if existingCategoryById == nil {
		c.JSON(http.StatusOK, response.Fail("分类不存在"))
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

func Edit(c *gin.Context) {
	var newPost post
	err := c.ShouldBindJSON(&newPost)
	if err != nil {
		log.Print("err")
		c.JSON(http.StatusOK, response.Fail("更新失败"))
		return
	}

	postById := QueryById(newPost.Id)
	if postById == nil {
		c.JSON(http.StatusOK, response.Fail("文章不存在"))
		return
	}

	affected := UpdatePost(newPost)
	if affected < 1 {
		log.Println("修改失败")
		c.JSON(http.StatusOK, response.Fail("修改失败"))
		return
	}

	c.JSON(http.StatusOK, response.Success("修改成功", nil))
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("参数错误"))
		return
	}

	postById := QueryById(postId)
	if postById == nil {
		c.JSON(http.StatusOK, response.Fail("文章不存在"))
		return
	}

	affected := DeleteById(postId)
	if affected < 1 {
		c.JSON(http.StatusOK, response.Success("删除失败", nil))
		return
	}
	c.JSON(http.StatusOK, response.Success("删除成功", nil))
	return
}
