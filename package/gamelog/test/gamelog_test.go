package test

import (
	"strings"
	"testing"
	"time"

	"github.com/ayflying/utility_go/package/gamelog"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/google/uuid"
)

func TestGamelog(t *testing.T) {
	glsdk, err := gamelog.INIT(&gamelog.SDKConfig{
		// 必填
		Pid:     "test5",                     // 项目ID
		BaseUrl: "http://47.76.178.47:10101", // 香港测试服上报地址
		// BaseUrl:       "http://127.0.0.1:10101", // 本次测试上报地址
		ReportSk:      "sngame2025", // xor混淆key
		FlushInterval: 5,            // 上报间隔
		DiskBakPath:   "gamelog",    // 本地磁盘备份, 用于意外情况下临时保存日志, 请确保该目录持久化(容器内要挂载). 每次启动时或每N次上报时加载到失败队列
		// 可填
		RetryN:       2,   // 默认每10次, 上传检查一次磁盘的失败数据
		ChanSize:     500, // 默认1000, 信道size
		SendSaveType: 2,   // 发送存储类型, 默认不设置为0代表文件存储, 2代表走kafka可实同步日志
	})

	// 随机测试事件和属性
	events := []string{"e1", "e2", "e3", "e4"}
	pms := []map[string]any{
		{"a": "1"},
		{"a": "2"},
		{"a": "3"},
		{"a": "4"},
	}
	if err != nil {
		t.Fatal(err)
	}
	gtest.C(t, func(t *gtest.T) {
		go func() {
			for {
				uuidval, _ := uuid.NewUUID()
				randUid := strings.ReplaceAll(uuidval.String(), "-", "")
				glsdk.LogLtz(randUid, events[grand.Intn(len(events))], pms[grand.Intn(len(pms))])
				time.Sleep(time.Millisecond * 100)
			}
		}()
		time.Sleep(time.Second * 14)
		// 模拟等待信号后优雅关闭
		glsdk.Shutdown()
	})
}
