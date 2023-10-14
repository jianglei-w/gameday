package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
}

// Result 构造返回体
func Result(data interface{}, msg string, status int, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Data:   data,
		Msg:    msg,
		Status: status,
	})
}

const (
	SUCCESS = 0 // 正确返回
	ERROR   = 1 // 错误返回
)

// 成功返回

// OK 单一返回操作成功,无返回值
func OK(c *gin.Context) {
	Result(map[string]interface{}{}, "操作成功", SUCCESS, c)
}

// OKWithMessage 成功返回消息
func OKWithMessage(msg string, c *gin.Context) {
	Result(map[string]interface{}{}, msg, SUCCESS, c)
}

// OKWithData 成功返回消息和数据
func OKWithData(data interface{}, msg string, c *gin.Context) {
	Result(data, msg, SUCCESS, c)
}

// 失败返回

// FailWithMessage 失败返回错误信息
func FailWithMessage(msg string, c *gin.Context) {
	Result(map[string]interface{}{}, msg, ERROR, c)
}
