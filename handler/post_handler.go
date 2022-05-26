package handler

import (
	"fmt"
	"net/http"
	"strconv"

	postSchema "svi-backend/schema"
	"svi-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *postHandler {
	return &postHandler{postService}
}

func (h *postHandler) PostArticleHandler(c *gin.Context) {
	var postRequest postSchema.PostRequest

	err := c.ShouldBindJSON(&postRequest)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	post, err := h.postService.CreateArticle(postRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"code":  400,
			"msg":   "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
		"code": 200,
		"msg":  "success",
	})
}

func (h *postHandler) FindByIDBookHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	result, err := h.postService.FindArticleById(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"msg":   "not found",
			"code":  404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}

func (h *postHandler) UpdateArticleHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	var postRequest postSchema.PostRequest

	err := c.ShouldBindJSON(&postRequest)

	errorMessages := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
			"code":   400,
			"msg":    "bad request",
		})
		return
	}

	post, err := h.postService.UpdateArticle(postRequest, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  404,
			"msg":   "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
		"code": 200,
		"msg":  "success",
	})
}

func (h *postHandler) DeleteArticleHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	_, err := h.postService.DeleteArticle(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"code":  404,
			"msg":   "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func (h *postHandler) FindAllArticleHandler(c *gin.Context) {
	page := c.Param("page")
	limit := c.Param("id")
	Page, _ := strconv.Atoi(page)
	Limit, _ := strconv.Atoi(limit)
	result, err := h.postService.FindAllArticle(Page, Limit)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"msg":   "not found",
			"code":  404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})

}
