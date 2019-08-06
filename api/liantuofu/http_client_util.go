package liantuofu

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultTimeout = 5 // 默认超时时间
)

func Post(url string, paras map[string]string, timeout int64) string {

	if timeout == 0 {
		timeout = DefaultTimeout
	}
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	req, err := http.NewRequest("POST", url, strings.NewReader(PostMapToStr(paras)))
	if err != nil {
		return err.Error()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

/**
将Map类型转换为查询字符串
*/
func PostMapToStr(postMap map[string]string) string {
	if nil == postMap {
		return ""
	} else {
		postValues := url.Values{}
		for postKey, postValue := range postMap {
			postValues.Set(postKey, postValue)
		}
		postDataStr := postValues.Encode()
		return postDataStr
	}
}
