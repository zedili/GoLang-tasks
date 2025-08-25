package api

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (ctl *BaseController) Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func (ctl *BaseController) Fail(ctx *gin.Context, message string, err error) {
	response := ErrorResponse{
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
		fmt.Println("接口失败：", err)
	}
	ctx.JSON(http.StatusInternalServerError, response)
}

func (ctl *BaseController) GetCurrentUser(ctx *gin.Context) (*model.User, bool) {
	user, exists := ctx.Get("user")
	if !exists {
		return nil, false
	}
	if userObj, ok := user.(*model.User); ok {
		fmt.Println("jwt用户信息：", userObj)
		return userObj, ok
	}
	return nil, false

}
