package tools

import (
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
		return nil
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
