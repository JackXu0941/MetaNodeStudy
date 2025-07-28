package handlers

import (
	"net/http"
	"strconv"
	"test4/dbfiles"

	"github.com/gin-gonic/gin"
)

var (
	db = dbfiles.InitDB()
)

// 1.实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
// CreatePost 创建文章
func CreatePost(c *gin.Context) {

	// 从上下文中获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	// 解析请求参数
	var request struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 创建文章
	post := dbfiles.Post{
		Title:   request.Title,
		Content: request.Content,
		UserID:  userID.(uint),
		// UserID: 1,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败: " + err.Error()})
		return
	}

	// 返回创建成功的文章信息
	c.JSON(http.StatusCreated, gin.H{
		"message": "文章创建成功",
		"post":    post,
	})

}

// 2.实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。

// GetPosts 获取所有文章列表
func GetPosts(c *gin.Context) {
	var posts []dbfiles.Post

	// 查询所有文章，预加载用户信息
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章列表成功",
		"posts":   posts,
	})
}

// GetPost 获取单个文章的详细信息
func GetPost(c *gin.Context) {
	// 从URL参数中获取文章ID
	postID := c.Query("id")

	// 转换文章ID为整数
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 查询指定ID的文章，预加载用户信息
	var post dbfiles.Post
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 查询文章的评论，预加载用户信息
	var comments []dbfiles.Comment
	if err := db.Preload("User").Where("post_id = ?", id).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败: " + err.Error()})
		return
	}

	// 将评论添加到文章结构中
	post.Comments = comments

	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章成功",
		"post":    post,
	})
}

// 3.实现文章的更新功能，只有文章的作者才能更新自己的文章。

// UpdatePost 更新文章，只有文章作者才能更新
func UpdatePost(c *gin.Context) {
	// 从上下文中获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	// 从URL参数中获取文章ID
	postID := c.Query("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID参数"})
		return
	}

	// 转换文章ID为整数
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 解析请求参数
	var request struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 查询要更新的文章
	var post dbfiles.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 验证当前用户是否为文章作者
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "您无权限更新此文章"})
		return
	}

	// 更新文章
	post.Title = request.Title
	post.Content = request.Content

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败: " + err.Error()})
		return
	}

	// 返回更新成功的文章信息
	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
		"post":    post,
	})
}

// 4 .实现文章的删除功能，只有文章的作者才能删除自己的文章。

// DeletePost 删除文章，只有文章作者才能删除
func DeletePost(c *gin.Context) {
	// 从上下文中获取当前用户信息
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return
	}

	// 从URL参数中获取文章ID
	postID := c.Query("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID参数"})
		return
	}

	// 转换文章ID为整数
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 查询要删除的文章
	var post dbfiles.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 验证当前用户是否为文章作者
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "您无权限删除此文章"})
		return
	}

	// 删除文章相关的评论
	if err := db.Where("post_id = ?", id).Delete(&dbfiles.Comment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除评论失败: " + err.Error()})
		return
	}

	// 删除文章
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败: " + err.Error()})
		return
	}

	// 返回删除成功的信息
	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
	})
}
