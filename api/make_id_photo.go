package api

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const AppCode = "48e01692295348219ce26ac7af3931c2"
const URL = "https://alidphoto.aisegment.com/idphoto/make"
const FreeUrl = "http://zjz.market.alicloudapi.com/api/cut_check_pic"
const FreeAppCode = "48e01692295348219ce26ac7bf3921c2"
const DefaultSpec = "1"
const DefaultBk = "red"

func MakeIdPhoto(imgType, photoBase64, spec, bk string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	if strings.TrimSpace(spec) == "" {
		spec = DefaultSpec
	}
	if strings.TrimSpace(bk) == "" {
		bk = DefaultBk
	}
	data := "{\"type\": \""
	data = data + imgType + "\", \"photo\": \"" + photoBase64 + "\", \"spec\": \"" + spec + "\", \"bk\": \"" + bk + "\"}"

	req, err := http.NewRequest("POST", URL, strings.NewReader(data))
	if err != nil {
		return err.Error(), err
	}
	req.Header.Set("Authorization", "APPCODE "+AppCode)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)

	if err != nil {
		return err.Error(), err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func MakeIdPhotoFree(photoBase64, spec string) (string, error) {
	// 接口要求base64需要去除逗号及前面的类型信息
	photoBase64 = photoBase64[strings.Index(photoBase64, ",")+1:]
	client := &http.Client{Timeout: 20 * time.Second}
	if strings.TrimSpace(spec) == "" {
		spec = DefaultSpec
	}
	data := "{\"spec_id\": \""
	data = data + spec + "\", \"file\": \"" + photoBase64 + "\"}"

	req, err := http.NewRequest("POST", FreeUrl, strings.NewReader(data))
	if err != nil {
		return err.Error(), err
	}
	req.Header.Set("Authorization", "APPCODE "+FreeAppCode)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := client.Do(req)

	if err != nil {
		return err.Error(), err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
