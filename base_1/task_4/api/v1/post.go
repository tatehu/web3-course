package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/model"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, _ := c.Get("userID")
	post.UserID = uid.(uint)
	if err := model.GetDB().Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// 获取所有文章
func GetPosts(c *gin.Context) {
	var posts []model.Post
	if err := model.GetDB().Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// 获取单篇文章
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := model.GetDB().Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// 更新文章
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := model.GetDB().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	uid, _ := c.Get("userID")
	if post.UserID != uid.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "No permission"})
		return
	}
	var input model.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Title = input.Title
	post.Content = input.Content
	model.GetDB().Save(&post)
	c.JSON(http.StatusOK, post)
}

// 删除文章
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := model.GetDB().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	uid, _ := c.Get("userID")
	if post.UserID != uid.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "No permission"})
		return
	}
	model.GetDB().Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
