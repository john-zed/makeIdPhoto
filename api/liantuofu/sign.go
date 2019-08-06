package liantuofu

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func CreateSign(paras map[string]string, key string) string {
	res := ""
	var paraKeys []string
	for k := range paras {
		paraKeys = append(paraKeys, k)
	}
	// 按字符升序排序
	sort.Strings(paraKeys)
	for _, k := range paraKeys {
		if k != "sign" && k != "key" {
			res += k + "=" + paras[k] + "&"
		}
	}
	res += "key=" + key
	data := []byte(res)
	res = fmt.Sprintf("%x", md5.Sum(data))
	return strings.ToLower(res)
}
