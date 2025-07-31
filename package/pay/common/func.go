package common

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
	"sort"
	"strings"
)

// FormatPublicKey 将原始公钥字符串格式化为标准PEM格式的公钥
// 功能：为原始公钥添加PEM头部和尾部，并按64字符长度拆分换行，符合PKCS#8标准格式要求
// 参数 publicKey: 原始未格式化的公钥字符串（通常为Base64编码且无换行）
// 返回值: 格式化后的PEM格式公钥字符串
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	// 写入PEM格式头部
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")

	// 定义每行公钥的标准长度（PEM格式要求64字符/行）
	rawLen := 64
	keyLen := len(publicKey)
	// 计算需要拆分的总行数（向上取整）
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++ // 若有余数则增加一行
	}

	// 按行拆分并写入公钥内容
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			// 最后一行取剩余所有字符（处理不足64字符的情况）
			buffer.WriteString(publicKey[start:])
		} else {
			// 非最后行取固定64字符
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n') // 每行结束添加换行符
		start += rawLen
		end = start + rawLen
	}

	// 写入PEM格式尾部
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// ParseNotifyToBodyMap 将HTTP请求中的表单数据解析为键值对映射
// 功能：解析请求表单数据，提取单值字段并转换为map[string]interface{}格式
// 参数 req: 包含表单数据的HTTP请求对象
// 返回值: 解析后的键值对映射(bm)和可能的错误(err)
func ParseNotifyToBodyMap(req *http.Request) (bm map[string]interface{}, err error) {
	// 解析请求表单数据，若失败则返回错误
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	// 获取解析后的表单数据（key为字段名，value为字符串切片形式的字段值）
	var form map[string][]string = req.Form
	// 初始化结果映射，预分配容量（表单字段数+1，预留扩展空间）
	bm = make(map[string]interface{}, len(form)+1)
	// 遍历表单字段，仅保留单值字段（忽略多值字段）
	for k, v := range form {
		if len(v) == 1 {
			bm[k] = v[0]
		}
	}
	return
}

// BuildSignStr 根据传入的g.Map构建签名字符串
// 规则：对所有非空值的键进行字母排序后，按"key=value&"格式拼接，最后去除末尾的"&"
// 参数 bm: 包含键值对的g.Map
// 返回值: 构建好的签名字符串和可能的错误
func BuildSignStr(bm g.Map) (string, error) {
	var (
		buf     strings.Builder
		keyList []string
	)
	// 收集所有键名
	for k := range bm {
		keyList = append(keyList, k)
	}
	// 对键名进行字母排序
	sort.Strings(keyList)
	// 遍历排序后的键，拼接非空值的键值对
	for _, k := range keyList {
		if v := bm[k]; v != "" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(gconv.String(v))
			buf.WriteByte('&')
			// 去除末尾多余的'&'字符
			// 检查是否有有效的键值对被拼接
		}
	}
	if buf.Len() <= 0 {
		return "", errors.New("length is error")
	}
	return buf.String()[:buf.Len()-1], nil
}
