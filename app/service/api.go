package service

import (
	"fmt"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

// ApiService Cypress API 專用
type ApiService struct {
	ErrorNotifyUnit string
}

var apiSingleton *ApiService
var apiOnce sync.Once

// ApiIns 獲得Rotate對象
func ApiIns(notify string) *ApiService {
	apiOnce.Do(func() {
		apiSingleton = &ApiService{
			ErrorNotifyUnit: notify,
		}
	})
	return apiSingleton
}

// GetCurl
func GetCurl(url string, header map[string]string) (data interface{}, apiErr error) {

	// 組參數
	param := map[string]interface{}{}

	// 執行
	content, apiErr := sendGet(url, header, param)
	if apiErr != nil {
		fmt.Printf("%+v\n", apiErr)
		return nil, apiErr
	}

	m1 := make(map[string]interface{})
	if err := jsoniter.Unmarshal(content, &m1); err != nil {
		fmt.Printf("%+v\n", content)
		return nil, err
	}

	return m1, nil
}

// PostCurl
func PostCurl(url string, header map[string]string, param map[string]interface{}) (data interface{}, apiErr error) {

	// 執行
	content, apiErr := sendPost(url, header, param, true)

	if apiErr != nil {
		fmt.Printf("%+v\n", apiErr)
		return nil, apiErr
	}

	m1 := make(map[string]interface{})
	if err := jsoniter.Unmarshal(content, &m1); err != nil {
		fmt.Printf("%+v\n", content)
		return nil, err
	}

	return m1, nil
}

// PutCurl
func PutCurl(url string, header map[string]string, param map[string]interface{}) (data interface{}, apiErr error) {

	// 執行
	content, apiErr := sendPut(url, header, param, true)

	if apiErr != nil {
		fmt.Printf("%+v\n", apiErr)
		return nil, apiErr
	}

	m1 := make(map[string]interface{})
	if err := jsoniter.Unmarshal(content, &m1); err != nil {
		fmt.Printf("%+v\n", content)
		return nil, err
	}

	return m1, nil
}

// DeleteCurl
func DeleteCurl(url string, header map[string]string, param map[string]interface{}) (data interface{}, apiErr error) {

	// 執行
	content, apiErr := sendDelete(url, header, param, true)

	if apiErr != nil {
		fmt.Printf("%+v\n", apiErr)
		return nil, apiErr
	}

	m1 := make(map[string]interface{})
	if err := jsoniter.Unmarshal(content, &m1); err != nil {
		fmt.Printf("%+v\n", content)
		return nil, err
	}

	return m1, nil
}
