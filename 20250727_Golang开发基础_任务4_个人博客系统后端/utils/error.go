package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// APIError 定义API错误结构
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse 统一错误响应格式
type ErrorResponse struct {
	Error APIError `json:"error"`
}

// Logger 全局日志实例
var Logger = logrus.New()

// HandleError 统一错误处理函数
func HandleError(c *gin.Context, statusCode int, message string, err error) {
	// 记录错误日志
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"status":  statusCode,
			"message": message,
			"error":   err.Error(),
		}).Error("API Error")
	} else {
		Logger.WithFields(logrus.Fields{
			"status":  statusCode,
			"message": message,
		}).Error("API Error")
	}

	// 返回错误响应
	c.JSON(statusCode, ErrorResponse{
		Error: APIError{
			Code:    statusCode,
			Message: message,
		},
	})
}

// LogInfo 记录信息日志
func LogInfo(message string, fields map[string]interface{}) {
	if fields != nil {
		Logger.WithFields(fields).Info(message)
	} else {
		Logger.Info(message)
	}
}

// LogError 记录错误日志
func LogError(message string, err error, fields map[string]interface{}) {
	if fields != nil {
		Logger.WithFields(fields).Error(message, err)
	} else {
		Logger.Error(message, err)
	}
}
