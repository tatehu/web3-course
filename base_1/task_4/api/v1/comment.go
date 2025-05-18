package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal_blog/model"
	"strconv"
)

// 创建评论
func CreateComment(c *gin.Context) {
	postID := c.Param("id")
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, _ := c.Get("userID")
	comment.UserID = uid.(uint)
	comment.PostID = stringToUint(postID)
	if err := model.GetDB().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// 获取某篇文章的所有评论
func GetComments(c *gin.Context) {
	postID := c.Param("id")
	var comments []model.Comment
	if err := model.GetDB().Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// 工具函数
func stringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
