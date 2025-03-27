# utility_go：强大的Go语言工具类库

## 一、项目概述
`utility_go` 是一个功能丰富的Go语言工具类库，旨在助力开发者更高效地开发Go应用程序。它提供了一系列实用的功能模块，广泛涵盖数据库操作、缓存处理、支付接口集成、排名算法实现等多个方面，适用于各类Go项目。

## 二、安装方式
要使用 `utility_go` 工具类库，可使用 `go get` 命令进行安装。请确保你的Go环境已正确配置，且可访问互联网。

### 安装命令
```sh
go get github.com/ayflying/utility_go
```

### 安装验证
安装完成后，你可以在Go代码中导入 `utility_go` 相关的包，检查是否能够正常使用。示例如下：
```go
package main

import (
    // 这里导入utility_go中的某个具体包，根据实际使用情况调整
    "github.com/ayflying/utility_go/package/pay/apple"
)

func main() {
    // 进行一些简单的操作，如打印版本信息等
    // 这里假设apple包有一个Version函数
    // fmt.Println(apple.Version())
}
```

## 三、项目结构
`utility_go` 的项目结构设计清晰，各模块分工明确，便于开发者使用和扩展。以下是项目的主要目录结构及说明：

### 主要目录说明
- **api/**：存放与API相关的代码，包含不同模块的API接口定义，如 `admin`、`callback`、`system` 等。这些API接口用于处理各种外部请求，是项目与外部系统交互的重要入口。
- **service/**：服务层代码，包含各种业务逻辑的实现，如 `game_act.go`、`ip_2_region.go` 等。服务层负责处理具体的业务需求，将数据处理和业务逻辑封装在独立的函数或方法中，提高代码的可维护性和复用性。
- **controller/**：控制器层代码，主要用于处理HTTP请求和响应。控制器接收客户端的请求，调用相应的服务层方法进行处理，并将处理结果返回给客户端。
- **internal/**：内部包，包含项目的核心业务逻辑，如 `game`、`logic`、`model` 等模块。其中，`model` 模块定义了项目中使用的数据模型，如数据库表对应的结构体；`logic` 模块实现了各种业务逻辑的处理函数。
- **package/**：包含各种功能包，提供了丰富的工具和功能，具体如下：
  - **aycache**：缓存相关的功能包，提供了缓存操作的接口和实现，帮助开发者更方便地使用缓存技术，提高应用程序的性能。
  - **elasticsearch**：Elasticsearch相关的功能包，用于与Elasticsearch搜索引擎进行交互，实现数据的存储、检索和分析等功能。
  - **excel**：Excel处理相关的功能包，提供了Excel文件的读写操作接口，方便开发者处理Excel文件中的数据。
  - **pay**：支付相关的功能包，包含了与各种支付平台的接口集成，如 `apple`（苹果支付）、`playstore`（Google Play Store支付）等，支持应用内购买等支付功能。
  - **rand**：随机数相关的功能包，提供了生成各种随机数的函数，可用于测试、加密等场景。
  - **rank**：排名相关的功能包，实现了各种排名算法，如基于Redis的排行榜功能，可用于游戏排名、活动排名等场景。
  - **s3**：S3存储相关的功能包，用于与Amazon S3等云存储服务进行交互，实现文件的上传、下载、删除等操作。
- **tools/**：工具类代码，包含了一些常用的工具函数，如 `redis.go`（Redis操作相关）、`time.go`（时间处理相关）、`tools.go`（通用工具函数）等，方便开发者在项目中使用。

## 四、主要模块功能

### 4.1 pay/playstore
该模块主要用于与Google Play Store API交互，处理应用内购买相关的操作，提供以下主要功能：
- **创建客户端**：通过 `New` 函数创建并返回一个包含访问 `androidpublisher` API所需凭证的HTTP客户端，方便开发者与Google Play Store API进行通信。
- **使用自定义客户端**：`NewWithClient` 函数允许开发者使用自定义的HTTP客户端创建并返回一个包含访问 `androidpublisher` API所需凭证的HTTP客户端，增加了客户端的灵活性。
- **验证签名**：`VerifySignature` 函数用于验证应用内购买的签名，确保支付信息的安全性和合法性。

### 4.2 s3
`s3` 模块主要用于与S3存储服务进行交互，提供文件存储和管理的功能。其中，`ListBuckets` 函数可以列出S3存储桶的信息，方便开发者管理存储桶中的文件。

### 4.3 model
`model` 模块定义了项目中使用的数据模型，这些数据模型通常与数据库表相对应，用于数据的存储和操作。例如：
- **GameMailMass**：表示游戏邮件群发的数据模型，包含邮件的标题、内容、类型等信息。
- **GameBag**：表示游戏背包的数据模型，包含用户标识、道具数据、图鉴、手势等信息。
- **MemberSave**：（根据具体代码中的定义）可能表示用户会员信息的数据模型，用于存储用户的会员相关数据。

### 4.4 tools
`tools` 模块提供了一系列通用工具函数，涵盖时间处理、Redis操作、道具数据处理等功能：
- **时间处理**：`time.go` 文件中的函数可以进行时间计算，如获取本周开始时间、计算两个时间间隔天数等。
- **Redis操作**：`redis.go` 文件中的函数可以进行Redis的扫描操作，支持批量获取大量数据。
- **通用工具**：`tools.go` 文件中的函数提供了字符串处理、切片操作、道具数据合并等功能。

## 五、使用示例

### 5.1 支付模块示例
以下是一个使用 `pay/playstore` 模块验证应用内购买签名的示例代码：
```go
package main

import (
    "fmt"
    "github.com/ayflying/utility_go/package/pay/playstore"
)

func main() {
    // 初始化参数
    purchaseData := "your_purchase_data"
    signature := "your_signature"

    // 验证签名
    err := playstore.VerifySignature(purchaseData, signature)
    if err != nil {
        fmt.Println("Signature verification failed:", err)
        return
    }

    fmt.Println("Signature verification succeeded")
}
```

### 5.2 S3模块示例
以下是一个使用 `s3` 模块列出S3存储桶信息的示例代码：
```go
package main

import (
    "fmt"
    "github.com/ayflying/utility_go/package/s3"
)

func main() {
    // 列出S3存储桶信息
    buckets, err := s3.ListBuckets()
    if err != nil {
        fmt.Println("Failed to list buckets:", err)
        return
    }

    for _, bucket := range buckets {
        fmt.Println("Bucket name:", bucket.Name)
    }
}
```

### 5.3 工具模块示例
以下是一个使用 `tools` 模块处理道具数据的示例代码：
```go
package main

import (
    "fmt"
    "github.com/ayflying/utility_go/tools"
)

func main() {
    // 字符串转道具类型
    str := "1|10|2|20"
    result := tools.Tools.Spilt2Item(str)
    fmt.Println("Spilt2Item result:", result)

    // 切片换道具类型
    slice := []int64{1, 10, 2, 20}
    res := tools.Tools.Slice2Item(slice)
    fmt.Println("Slice2Item result:", res)

    // 道具格式转map
    list := tools.Tools.Items2Map(result)
    fmt.Println("Items2Map result:", list)
}
```

## 六、注意事项
- **自动生成文件**：项目中有部分代码文件是由GoFrame CLI工具生成并维护的，这些文件通常会标注有 `// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.` 注释。请不要手动修改这些文件，以免造成不必要的问题。
- **版本兼容性**：在使用 `utility_go` 工具类库时，请确保你的Go语言版本与工具类库的要求版本兼容。建议使用最新的Go语言版本，以获得更好的性能和稳定性。

## 七、更新和维护
`utility_go` 工具类库会定期进行更新和维护，以修复已知的问题、添加新的功能和优化性能。你可以关注项目的GitHub仓库（https://github.com/ayflying/utility_go）获取最新的更新信息。如果你在使用过程中遇到任何问题或有任何建议，欢迎提交Issue或Pull Request。

## 八、许可证信息
`utility_go` 工具类库遵循MIT许可证，你可以在项目的 `LICENSE` 文件中查看详细的许可证条款。请确保在使用该工具类库时遵守相关的许可证规定。