package handlers

import (
	"net/http"
	"strconv"
	"test4/dbfiles"

	"github.com/gin-gonic/gin"
)

var (
	db1 = dbfiles.InitDB()
)

// CreateComment 创建评论，已认证的用户可以对文章发表评论
func CreateComment(c *gin.Context) {
	// 从上下文中获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	// 从URL参数中获取文章ID
	postID := c.Query("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID参数"})
		return
	}

	// 转换文章ID为整数
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 解析请求参数
	var request struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 检查文章是否存在
	var post dbfiles.Post
	if err := db1.First(&post, postIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 创建评论
	comment := dbfiles.Comment{
		Content: request.Content,
		UserID:  userID.(uint),
		PostID:  uint(postIDInt),
	}

	if err := db1.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败: " + err.Error()})
		return
	}

	// 预加载用户信息返回
	db1.Preload("User").First(&comment, comment.ID)

	// 返回创建成功的评论信息
	c.JSON(http.StatusCreated, gin.H{
		"message": "评论创建成功",
		"comment": comment,
	})
}

// GetComments 获取某篇文章的所有评论列表
func GetComments(c *gin.Context) {
	// 从URL参数中获取文章ID
	postID := c.Query("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID参数"})
		return
	}

	// 转换文章ID为整数
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 检查文章是否存在
	var post dbfiles.Post
	if err := db.First(&post, postIDInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 查询该文章的所有评论，预加载用户信息
	var comments []dbfiles.Comment
	if err := db.Preload("User").Where("post_id = ?", postIDInt).Order("created_at ASC").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论列表失败: " + err.Error()})
		return
	}

	// 返回评论列表
	c.JSON(http.StatusOK, gin.H{
		"message":  "获取评论列表成功",
		"comments": comments,
	})
}
