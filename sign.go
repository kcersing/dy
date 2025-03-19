package dy

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"sort"
)

// SignV2 用于生活服务应用计算header中的签名(sha256) 新申请应用请使用该方法计算
func SignV2(clientSecret, body string, query map[string]string) string {

	keys := make([]string, 0)
	for k, _ := range query {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Sort(stringSlice(keys))

	str := clientSecret
	for _, k := range keys {
		val := query[k]
		str = fmt.Sprintf("%s&%s", str, fmt.Sprintf("%s=%s", k, val))
	}
	if body != "" {
		str = fmt.Sprintf("%s&http_body=%s", str, body)
	}

	result := Sha256(str)
	fmt.Printf("[Sign] v2 str:%s, result:%s\n", str, result)

	return result
}

// Sign 用于生活服务网站应用计算url签名(md5) 不再使用
func Sign(clientSecret, body string, query map[string]string) string {
	keys := make([]string, 0)
	for k, _ := range query {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Sort(stringSlice(keys))

	str := clientSecret
	for _, k := range keys {
		val := query[k]
		str = fmt.Sprintf("%s&%s", str, fmt.Sprintf("%s=%s", k, val))
	}
	if body != "" {
		str = fmt.Sprintf("%s&http_body=%s", str, body)
	}

	fmt.Printf("[Sign] str:%s\n", str)

	return MD5(str)
}

func MD5(str string) string {
	data := []byte(str) // 切片
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash) // 将[]byte转成16进制
	return md5str
}

func Sha256(str string) string {
	sum256 := sha256.Sum256([]byte(str))
	shaStr := fmt.Sprintf("%x", sum256) // 将[]byte转成16进制
	return shaStr
}

type stringSlice []string

func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
