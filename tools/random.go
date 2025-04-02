package tools

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"math/rand"
	"time"
)

var (
	Rand *randMod
)

type randMod struct {
}

// RandomAll 按权重随机选取 N 个不重复的元素
func (m *randMod) RandomAll(data map[int]int, n int) []int {
	if n > len(data) {
		n = len(data)
	}
	rand.Seed(time.Now().UnixNano())
	// 复制权重映射，避免修改原始数据
	remainingWeights := make(map[int]int)
	for k, v := range data {
		remainingWeights[k] = v
	}
	result := make([]int, 0, n)

	for i := 0; i < n; i++ {
		totalWeight := 0
		// 计算剩余元素的总权重
		for _, weight := range remainingWeights {
			totalWeight += weight
		}
		if totalWeight == 0 {
			break
		}
		// 生成一个 0 到总权重之间的随机数
		randomNum := rand.Intn(totalWeight)
		currentWeight := 0
		for key, weight := range remainingWeights {
			currentWeight += weight
			if randomNum < currentWeight {
				// 将选中的元素添加到结果切片中
				result = append(result, key)
				// 从剩余权重映射中移除选中的元素
				delete(remainingWeights, key)
				break
			}
		}
	}
	return result
}

// RandByArrInt 根据传入的 interface 切片中的整数值按权重随机返回一个索引
// 参数 s: 一个包含整数的 interface 切片，切片中的每个元素代表一个权重
// 返回值: 随机选中的元素的索引
func (m *randMod) RandByArrInt(_s interface{}) int {
	// 初始化总权重为 0
	sv := 0
	s := gconv.Ints(_s)

	// 遍历切片，累加每个元素的权重
	for i := range s {
		sv += gconv.Int(s[i])
	}
	if sv < 1 {
		return 0
	}

	// 使用 grand.Intn 生成一个 0 到总权重之间的随机数
	r := grand.Intn(sv)
	// 初始化当前累加的权重为 0
	var all int
	// 再次遍历切片，累加权重
	for i := range s {
		all += s[i]
		// 如果当前累加的权重大于随机数，则返回当前索引
		if all > r {
			return i
		}
	}
	// 如果没有找到符合条件的索引，返回 0
	return 0
}
