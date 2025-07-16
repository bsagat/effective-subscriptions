package httputils

import (
	"net/http"
	"submanager/internal/domain"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, err error) {
	errMessage := struct {
		Code    int    `json:"code"`
		Message string `json:"error"`
	}{
		Code:    code,
		Message: err.Error(),
	}
	ctx.JSON(errMessage.Code, &errMessage)
}

func SendMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
}

func GetStatus(err error) int {
	switch err {
	case domain.ErrSubNotUnique:
		return http.StatusConflict
	case domain.ErrSubsNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
