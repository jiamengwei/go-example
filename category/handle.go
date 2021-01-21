package category

import (
	"github.com/gin-gonic/gin"
	"go.example/blog/response"
	"log"
	"net/http"
	"strconv"
)

func Save(c *gin.Context) {
	var category category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("参数异常", err)
		c.JSON(http.StatusBadRequest, response.Fail(err.Error()))
		return
	}

	existingByName := queryByName(category.Name)
	if existingByName != nil {
		c.JSON(http.StatusOK, response.Fail("分类名已存在"))
		return
	}

	id, err := saveCategory(category)
	if err != nil {
		log.Println("分类保存失败", err)
		c.JSON(http.StatusOK, response.Fail("分类保存失败"))
		return
	}

	c.JSON(http.StatusOK, response.Success("success", id))
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.Fail("参数错误"))
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Fail("参数类型错误"))
	}
	byId := QueryById(idInt)
	if byId == nil {
		c.JSON(http.StatusOK, response.Fail("分类不存在"))
		return
	}
	rowsAffected, err := deleteById(id)

	if err != nil {
		c.JSON(http.StatusOK, response.Fail("删除失败"))
		return
	}
	c.JSON(http.StatusOK, response.Success("删除成功", rowsAffected))
}
