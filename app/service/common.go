package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// APIResult 回傳API格式
type APIResult struct {
	ErrorMsg string      `json:"error_msg"`
	Result   interface{} `json:"result"`
}

// validateStruct 驗證struct規則
func validateStruct(req interface{}) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

// CatchError 回傳不可預期的錯誤
func catchError(c *gin.Context) {
	if err := recover(); err != nil {
		// 回傳不可預期的錯誤
		c.JSON(http.StatusBadRequest, err)
	}
}

// Success 回傳成功API
func success(result interface{}) *APIResult {
	res := &APIResult{
		ErrorMsg: "",
		Result:   []string{},
	}

	if result != "" && result != nil {
		res.Result = result
		return res
	}

	return res
}

// Fail 回傳失敗API
func fail(err error) *APIResult {
	res := &APIResult{}

	res.ErrorMsg = err.Error()
	res.Result = []string{}

	return res
}
