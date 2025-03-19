package dy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func URLEncode(req map[string]string, URL string) string {
	encode := url.Values{}
	for k, v := range req {
		encode.Add(k, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("%s?%s", URL, encode.Encode())
}

func Request(ctx context.Context, method string, headers map[string]string, params map[string]string, url string) (interface{}, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	var reqBody io.Reader
	// 如果是GET请求，将参数编码到URL中
	if method == http.MethodGet {
		url = URLEncode(params, url)
	} else {
		// 对于POST请求，将参数编码为JSON
		jsonData, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	// 尝试将响应体解析为JSON
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		// 如果解析失败，返回原始字符串
		return string(body), nil
	}

	return result, nil
}
