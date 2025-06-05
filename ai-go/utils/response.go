package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 标准响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

// SuccessWithMessage 返回带自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// Fail 返回失败响应
func Fail(c *gin.Context, httpStatus int, code int, message string, err error) {
	response := Response{
		Code:    code,
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
	}
	c.JSON(httpStatus, response)
}

// BadRequest 返回400错误
func BadRequest(c *gin.Context, message string, err error) {
	Fail(c, http.StatusBadRequest, 400, message, err)
}

// Unauthorized 返回401错误
func Unauthorized(c *gin.Context, message string, err error) {
	Fail(c, http.StatusUnauthorized, 401, message, err)
}

// NotFound 返回404错误
func NotFound(c *gin.Context, message string, err error) {
	Fail(c, http.StatusNotFound, 404, message, err)
}

// InternalError 返回500错误
func InternalError(c *gin.Context, message string, err error) {
	Fail(c, http.StatusInternalServerError, 500, message, err)
}
