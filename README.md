# utility_go

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/ayflying/utility_go.svg)](https://github.com/ayflying/utility_go/stargazers)

一个功能丰富的Go语言工具类库，提供支付集成、缓存管理、排名算法、数据库操作、系统管理等核心功能模块，适用于各类Go项目开发。

## ✨ 核心特性

- **🔧 多模块集成**: 支付、缓存、排名、数据库、系统管理等完整解决方案
- **⚡ 高性能**: 基于Redis的高性能排名算法和缓存管理
- **📱 支付支持**: 集成Apple Pay、Google Play Store、支付宝、微信支付等主流支付平台
- **🛠️ CLI工具**: 提供代码生成器，快速创建模块文件
- **📊 监控告警**: 集成Prometheus监控和系统日志管理
- **🌐 多存储支持**: 支持S3对象存储、Elasticsearch搜索引擎

## 📦 安装

```bash
go get github.com/ayflying/utility_go
```

## 🏗️ 项目结构

```
utility_go/
├── api/                    # API接口定义
│   ├── admin/             # 管理后台API
│   ├── callback/          # 回调接口
│   ├── system/            # 系统API
│   └── pkg/               # 包相关API
├── cmd/                   # CLI命令工具
│   ├── make.go            # 代码生成器
│   ├── load.go            # 配置加载
│   ├── update.go          # 更新工具
│   └── middleware.go      # 中间件工具
├── controller/            # 控制器层
│   └── callback/          # 回调控制器
├── service/               # 服务层
│   ├── casdoor.go         # 认证服务
│   ├── game_act.go        # 游戏活动服务
│   ├── game_kv.go         # 键值存储服务
│   ├── ip_2_region.go     # IP地理位置服务
│   ├── log_data.go        # 日志数据服务
│   ├── os.go              # 操作系统服务
│   ├── system_cron.go     # 定时任务服务
│   └── system_log.go      # 系统日志服务
├── internal/              # 内部核心逻辑
│   ├── boot/              # 启动初始化
│   ├── game/              # 游戏逻辑
│   ├── logic/             # 业务逻辑
│   │   ├── casdoor/       # 认证逻辑
│   │   ├── gameAct/       # 游戏活动
│   │   ├── gameKv/        # 游戏键值存储
│   │   ├── ip2region/     # IP区域查询
│   │   ├── logData/       # 日志处理
│   │   ├── os/            # 系统操作
│   │   ├── systemCron/    # 定时任务
│   │   └── systemLog/     # 系统日志
│   └── model/             # 数据模型
│       ├── do/            # 数据对象
│       └── entity/        # 实体定义
├── package/               # 功能包
│   ├── aycache/           # 缓存管理
│   ├── excel/             # Excel处理
│   ├── gamelog/           # 游戏日志
│   ├── pay/               # 支付集成
│   │   ├── alipay/        # 支付宝
│   │   ├── apple/         # Apple Pay
│   │   ├── google/        # Google Play
│   │   ├── wechat/        # 微信支付
│   │   ├── playstore/     # Play Store
│   │   └── ...            # 其他支付平台
│   ├── rank/              # 排名算法
│   └── s3/                # S3存储
├── pkg/                   # 公共包
│   ├── aycache/           # 缓存包
│   ├── config/            # 配置包
│   ├── elasticsearch/     # ES包
│   ├── notice/            # 通知包
│   ├── rank/              # 排名包
│   └── s3/                # S3包
├── tools/                 # 工具函数
│   ├── random.go          # 随机数
│   ├── redis.go           # Redis操作
│   ├── time.go            # 时间处理
│   └── tools.go           # 通用工具
├── utility.go             # 主入口
├── go.mod                 # 模块定义
└── README.md              # 说明文档
```

## 🚀 快速开始

### 1. 初始化项目

```go
package main

import (
    _ "github.com/ayflying/utility_go"
    "github.com/gogf/gf/v2/frame/g"
)

func main() {
    g.Log().Info(nil, "项目启动成功")
}
```

### 2. 使用CLI工具生成代码

```bash
# 创建活动模块
go run main.go make -m act -i 1

# 创建逻辑模块
go run main.go make -m logic -n test

# 创建配置文件
go run main.go make -m config -n test

# 创建Socket模块
go run main.go make -m socket -n test
```

## 📚 核心模块详解

### 🔐 支付模块 (package/pay)

支持多种支付平台集成：

#### Google Play Store
```go
import "github.com/ayflying/utility_go/package/pay/playstore"

// 创建客户端
client, err := playstore.New(jsonKey)

// 验证签名
isValid, err := playstore.VerifySignature(publicKey, receipt, signature)
```

#### Apple Pay
```go
import "github.com/ayflying/utility_go/package/pay/apple"

// 处理Apple支付验证
```

#### 支付宝/微信支付
```go
import (
    "github.com/ayflying/utility_go/package/pay/alipay"
    "github.com/ayflying/utility_go/package/pay/wechat"
)
```

### 🏆 排名模块 (package/rank)

基于Redis的高性能排行榜实现：

```go
import "github.com/ayflying/utility_go/pkg/rank"

// 创建排行榜
rankMod := rank.New()
leaderboard := rankMod.CreateF64CountRank("season_1")

// 增加分数
curScore, err := leaderboard.IncrScore(userID, 100)

// 设置分数
err := leaderboard.SetScore(userID, 500)

// 获取排名信息
rankInfo, err := leaderboard.GetRankInfosNotTs(0, 10)
```

### 💾 缓存模块 (package/aycache)

```go
import "github.com/ayflying/utility_go/package/aycache"

// 使用缓存
cache := aycache.New()
```

### 📊 数据库操作

#### Elasticsearch
```go
import "github.com/ayflying/utility_go/package/elasticsearch"

// ES操作
```

#### MySQL (基于GoFrame)
```go
import "github.com/gogf/gf/v2/database/gdb"
```

### 🔄 定时任务 (systemCron)

```go
import "github.com/ayflying/utility_go/service"

// 添加定时任务
service.SystemCron().AddCronV2(v1.CronType_HOUR, func(ctx context.Context) error {
    // 执行任务
    return nil
}, true)
```

### 📝 日志管理 (systemLog)

```go
import "github.com/ayflying/utility_go/service"

// 记录日志
service.SystemLog().Info("操作成功")
service.SystemLog().Error("操作失败: %v", err)
```

### 🌐 IP地理位置查询 (ip2region)

```go
import "github.com/ayflying/utility_go/service"

// 查询IP位置
ipInfo, err := service.Ip2Region().Search("192.168.1.1")
```

### 📁 S3存储 (s3)

```go
import "github.com/ayflying/utility_go/package/s3"

// 列出存储桶
buckets, err := s3.ListBuckets()

// 上传文件
err := s3.UploadFile("bucket-name", "file-key", fileData)
```

### 📑 Excel处理 (excel)

```go
import "github.com/ayflying/utility_go/package/excel"

// 读取Excel
data, err := excel.Read("file.xlsx")

// 导出Excel
err := excel.Export(data, "output.xlsx")
```

## 🛠️ 配置管理

项目使用GoFrame框架的配置管理：

```go
import (
    "github.com/ayflying/utility_go/config"
    "github.com/gogf/gf/v2/frame/g"
)

// 获取配置
cfg := config.Cfg{}
dbConfig := g.Cfg().Get("database")
```

## 🔧 工具函数

### 时间处理 (tools/time.go)
```go
import "github.com/ayflying/utility_go/tools"

// 获取本周开始时间
weekStart := tools.GetWeekStart()

// 计算天数差
days := tools.DiffDays(startTime, endTime)
```

### Redis操作 (tools/redis.go)
```go
// 批量获取数据
data, err := tools.ScanRedis(pattern)
```

### 通用工具 (tools/tools.go)
```go
// 字符串转道具类型
items := tools.Tools.Spilt2Item("1|10|2|20")

// 切片转道具类型
items := tools.Tools.Slice2Item([]int64{1, 10, 2, 20})

// 道具格式转Map
itemMap := tools.Tools.Items2Map(items)
```

## 🎯 使用示例

### 完整的支付验证流程
```go
package main

import (
    "fmt"
    "github.com/ayflying/utility_go/package/pay/playstore"
)

func main() {
    // Google Play Store应用内购买验证
    purchaseData := "purchase_data_string"
    signature := "signature_string"
    publicKey := "base64_encoded_public_key"
    
    isValid, err := playstore.VerifySignature(publicKey, []byte(purchaseData), signature)
    if err != nil {
        fmt.Println("验证失败:", err)
        return
    }
    
    if isValid {
        fmt.Println("支付验证成功")
    }
}
```

### 排行榜使用示例
```go
package main

import (
    "fmt"
    "github.com/ayflying/utility_go/pkg/rank"
)

func main() {
    rankMod := rank.New()
    leaderboard := rankMod.CreateF64CountRank("game_season_1")
    
    // 用户得分
    score, err := leaderboard.IncrScore(1001, 50)
    if err != nil {
        fmt.Println("更新分数失败:", err)
        return
    }
    
    fmt.Printf("当前分数: %.0f\n", score)
    
    // 获取前10名
    top10, err := leaderboard.GetRankInfosNotTs(0, 9)
    if err != nil {
        fmt.Println("获取排名失败:", err)
        return
    }
    
    for i, info := range top10 {
        fmt.Printf("第%d名: 用户%d - 分数%d\n", i+1, info.Id, info.Score)
    }
}
```

## ⚙️ 开发指南

### 代码规范
- 使用GoFrame框架的最佳实践
- 遵循Go语言命名规范
- 添加必要的注释和文档
- 使用错误处理和日志记录

### 项目启动流程
1. 配置文件加载
2. 数据库连接初始化
3. 缓存系统初始化
4. 定时任务注册
5. 服务启动监听

### 注意事项
- ⚠️ **自动生成文件**: 使用CLI工具生成的文件包含 `// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.` 注释，请勿手动修改
- 🔒 **版本兼容性**: 建议使用Go 1.24+版本
- 📦 **依赖管理**: 使用go mod管理依赖

## 🔄 更新维护

- **定期更新**: 修复已知问题，添加新功能
- **性能优化**: 持续优化性能表现
- **安全补丁**: 及时修复安全漏洞

## 📄 许可证

本项目采用 [MIT License](LICENSE) 许可证。

## 🤝 贡献指南

欢迎提交Issue和Pull Request来贡献代码！

1. Fork本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启Pull Request

## 📞 联系方式

- **GitHub**: https://github.com/ayflying/utility_go
- **Gitee**: https://gitea.adesk.com/public_project/utility_go

## 🙏 致谢

感谢所有贡献者和开源社区的支持！

---

**utility_go** © 2025 - Made with ❤️ by [ayflying]
