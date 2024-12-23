package utility

import (
	"math/rand"
	"time"
)

// rands 结构体用于封装 rand.Rand 实例，以提供随机数生成功能。
// 该结构体目前不包含锁，因此在多线程环境下使用时应注意同步问题。
type rands struct {
	r *rand.Rand
	// lock sync.Mutex
}

// Rand 是一个全局的 rands 实例，用于在整个程序中生成随机数。
// 它使用当前时间的毫秒值作为随机源，以确保每次程序运行时都能获得不同的随机数序列。
var Rand = rands{
	r: rand.New(rand.NewSource(time.Now().UnixMilli())),
}

// RandByArrInt 函数从一个整数数组中按权重选择一个索引，并返回该索引。
// 权重是数组中相应元素的值。该函数通过计算累积和来确定选择的索引。
// 参数 v 是一个泛型参数，限制为实现了 Number 接口的类型。
// 返回值是一个整数，表示在数组中的索引。
func RandByArrInt[v Number](s []v) int {
	sv := 0
	for i := range s {
		sv += int(s[i])
	}
	r := Rand.Intn(sv)
	var all v
	for i := range s {
		all += s[i]
		if all > v(r) {
			return i
		}
	}
	return 0
}

// Intn 方法通过给定的整数 i 生成一个 0 到 i-1 之间的随机数。
// 如果 i 为0，则会触发 panic。
// 参数 i 是一个整数，表示生成随机数的上限（不包含）。
// 返回值 ret 是一个在 0 到 i-1 范围内的随机整数。
func (r rands) Intn(i int) (ret int) {
	if i == 0 {
		panic(1)
	}
	return rand.Intn(i)
}
