package utility_go_test

import (
	//_ "github.com/ayflying/utility_go/internal/logic"

	"github.com/ayflying/utility_go/internal/boot"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

var (
	ctx = gctx.GetInitCtx()
)

func TestInit(t *testing.T) {
	g.Log().Debug(ctx, "开始调试了")
	// 初始化配置
	var err = boot.Boot()
	if err != nil {
		t.Error(err)
	}
}

//
//func TestLoadConfig(t *testing.T) {
//
//	tests := []struct {
//		name     string
//		filePath string
//		wantErr  bool
//	}{
//		{
//			name:     "valid config file",
//			filePath: "testdata/valid_config.json",
//			wantErr:  false,
//		},
//		{
//			name:     "non-existent file",
//			filePath: "nonexistent.json",
//			wantErr:  true,
//		},
//		{
//			name:     "invalid config format",
//			filePath: "testdata/invalid_config.json",
//			wantErr:  true,
//		},
//		{
//			name:     "empty file path",
//			filePath: "",
//			wantErr:  true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			_, err := config.Load(tt.filePath)
//			if tt.wantErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
