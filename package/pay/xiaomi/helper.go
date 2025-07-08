package xiaomi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// SignatureHelper 签名辅助类
type SignatureHelper struct{}

// hmacSHA1 计算HMAC-SHA1哈希值
func hmacSHA1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Sign 计算hmac-sha1签名
func (m *MiPay) Sign(params map[string]string, secretKey string) string {
	if _, ok := params["signature"]; ok {
		delete(params, "signature")
	}
	for k, v := range params {
		if v == "" || v == "0" {
			delete(params, k)
		}
	}
	sortString := m.buildSortString(params)
	signature := hmacSHA1(sortString, secretKey)
	return signature
}

// VerifySignature 验证签名
func (m *MiPay) VerifySignature(params map[string]string, signature, secretKey string) bool {
	tmpSign := m.Sign(params, secretKey)
	return tmpSign == signature
}

// buildSortString 构造排序字符串
func (m *MiPay) buildSortString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	// 按键排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建排序字符串
	var fields []string
	for _, k := range keys {
		fields = append(fields, fmt.Sprintf("%s=%s", k, params[k]))
	}

	return strings.Join(fields, "&")
}
