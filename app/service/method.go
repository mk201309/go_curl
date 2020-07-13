package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	Url    string                 `json:"url" validate:"required" `
	Header map[string]string      `json:"header" validate:"required"`
	Params map[string]interface{} `json:"params"`
}

type postRequest struct {
	getRequest
	Params map[string]interface{} `json:"params" validate:"required"`
}

type putRequest struct {
	postRequest
}

type deleteRequest struct {
	postRequest
}

func Get(c *gin.Context) {
	// 接不可預期的錯誤
	defer catchError(c)

	// 組參數
	req := getRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	// 驗證參數
	if err := validateStruct(req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	resp, err := GetCurl(req.Url, req.Header, req.Params)
	if err != nil {
		c.JSON(http.StatusOK, fail(err))
		return
	}

	c.JSON(http.StatusOK, success(resp))
}

func Post(c *gin.Context) {
	// 接不可預期的錯誤
	defer catchError(c)

	// 組參數
	req := postRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	// 驗證參數
	if err := validateStruct(req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	resp, err := PostCurl(req.Url, req.Header, req.Params)
	if err != nil {
		c.JSON(http.StatusOK, fail(err))
		return
	}

	c.JSON(http.StatusOK, success(resp))
}

func Put(c *gin.Context) {
	// 接不可預期的錯誤
	defer catchError(c)

	// 組參數
	req := putRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	// 驗證參數
	if err := validateStruct(req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	resp, err := PutCurl(req.Url, req.Header, req.Params)
	if err != nil {
		c.JSON(http.StatusOK, fail(err))
		return
	}

	c.JSON(http.StatusOK, success(resp))
}

func Delete(c *gin.Context) {
	// 接不可預期的錯誤
	defer catchError(c)

	// 組參數
	req := deleteRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	// 驗證參數
	if err := validateStruct(req); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusOK, fail(err))
		return
	}

	resp, err := DeleteCurl(req.Url, req.Header, req.Params)
	if err != nil {
		c.JSON(http.StatusOK, fail(err))
		return
	}

	c.JSON(http.StatusOK, success(resp))
}
