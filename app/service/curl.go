package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// sendGet CURL GET
func sendGet(apiURL string, header map[string]string, param map[string]interface{}) (body []byte, apiErr error) {
	client := &http.Client{}
	// 建立一個請求
	reqest, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}
	// 組參數
	q := reqest.URL.Query()
	for pk, pv := range param {
		paramV := reflect.ValueOf(pv)
		if paramV.Kind() == reflect.Slice {
			for i := 0; i < paramV.Len(); i++ {
				value := paramV.Index(i)
				q.Add(pk, fmt.Sprintf("%v", value))
			}
			continue
		}
		q.Add(pk, fmt.Sprintf("%v", paramV))
	}
	reqest.URL.RawQuery = q.Encode()

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}

// sendPost CURL POST
func sendPost(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr error) {
	// 組參數 - FormData, RawData
	var requestData string
	if rawData {
		byteData, err := jsoniter.Marshal(param)
		if err != nil {
			return nil, err
		}
		requestData = string(byteData)
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(requestData))
	if err != nil {
		return nil, err
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}

// sendPut CURL PUT
func sendPut(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr error) {
	// 組參數 - FormData, RawData
	var requestData string
	if rawData {
		byteData, err := jsoniter.Marshal(param)
		if err != nil {
			return nil, err
		}
		requestData = string(byteData)
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPut, apiURL, strings.NewReader(requestData))
	if err != nil {

		return nil, err
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}

// sendDelete CURL Delete
func sendDelete(apiURL string, header map[string]string, param map[string]interface{}, rawData bool) (body []byte, apiErr error) {
	// 組參數 - FormData, RawData
	var requestData string
	if rawData {
		byteData, err := jsoniter.Marshal(param)
		if err != nil {
			return nil, err
		}
		requestData = string(byteData)
	} else {
		form := url.Values{}
		for pk, pv := range param {
			paramV := reflect.ValueOf(pv)
			if paramV.Kind() == reflect.Slice {
				for i := 0; i < paramV.Len(); i++ {
					value := paramV.Index(i)
					form.Add(pk, fmt.Sprintf("%v", value))
				}
				continue
			}
			form.Add(pk, fmt.Sprintf("%v", paramV))
		}
		requestData = form.Encode()
	}

	// 建立一個請求
	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodDelete, apiURL, strings.NewReader(requestData))
	if err != nil {
		return nil, err
	}

	// 組Header
	for hk, hv := range header {
		reqest.Header.Add(hk, hv)
	}

	// 執行
	resp, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}
