package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_KEY    = "2ZwofljhHha9NulOKxeYxc94"
	SECRET_KEY = "ZC9l9I8oZgNoCa4fNelw8s0RjlkogZbD"
)

// GetAccessToken 获取百度API的访问令牌
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return ""
	}
	token, ok := result["access_token"].(string)
	if !ok {
		return ""
	}
	return token
}

func IsTextValid(text string) (bool, string, error) {
	token := GetAccessToken()
	if token == "" {
		return false, "", fmt.Errorf("获取访问令牌失败")
	}

	url := fmt.Sprintf("https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined?access_token=%s", token)
	payload := strings.NewReader("text=" + text)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return false, "", fmt.Errorf("创建请求失败: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "", fmt.Errorf("读取响应失败: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查结论类型
	conclusionType, ok := result["conclusionType"].(float64)
	if !ok {
		return false, "", fmt.Errorf("无效的响应格式: 缺少conclusionType字段")
	}

	// 1:合规，2:不合规，3:需要人工审核
	if conclusionType == 1 {
		return true, "消息合规", nil
	}

	// 获取具体的违规原因
	//conclusion, _ := result["conclusion"].(string)
	return false, "消息不合规", nil
}
