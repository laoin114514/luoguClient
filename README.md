# LuoguSDK

洛谷 (luogu.com.cn) 平台的 Go SDK，覆盖认证、题目、记录、题单等核心功能。

## 安装

```bash
go get github.com/laoin114514/luoguClient
```

## 快速开始

```go
package main

import (
    "fmt"
    luogu "github.com/laoin114514/luoguClient"
)

func main() {
    client, _ := luogu.NewClient()

    // 登录（首次需要手动输入验证码）
    if !client.Auth.IsAuthenticated() {
        client.Auth.RefreshCSRF()
        client.Auth.LoginWithSolver("username", "password", mySolver)
    }

    // 获取题目
    problem, _ := client.Problem.Get("P1001")
    fmt.Println(problem.Title) // A+B Problem

    // 搜索题目
    results, _ := client.Problem.Search(luogu.SearchParams{
        Keyword: "排序", Page: 1,
    })

    // 查看提交记录
    records, _ := client.Record.GetList(luogu.RecordListParams{
        User:   1582049,
        Status: luogu.StatusAccepted,
        Page:   1,
    })

    // 获取题单
    list, _ := client.Training.GetList(luogu.TrainingListParams{Page: 1})
    detail, _ := client.Training.GetDetail(list.Trainings[0].ID)
}
```

## Cookie 管理

Cookie 只保存在内存中，SDK 不读写任何文件。是否持久化、存到哪里（文件、数据库、Redis 等）由调用方自行决定。

```go
client, _ := luogu.NewClient()

// 从调用方自己的存储恢复登录态
if data, err := os.ReadFile("cookies.json"); err == nil {
    client.ImportCookies(data)
}

// 登录成功后导出，由调用方自行保存
data, _ := client.ExportCookies()
os.WriteFile("cookies.json", data, 0600)

// 清空内存中的 cookie（例如登出后）
client.ClearCookies()
```

## Context 控制

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
client, _ := luogu.NewClient(luogu.WithContext(ctx))
client.Problem.Get("P1001") // 5 秒超时自动中断
```

## 内置常量

```go
// 提交状态（通过实测抓包验证）
luogu.StatusCompiling   // 2  编译中
luogu.StatusAccepted    // 12 通过
luogu.StatusUnaccepted  // 14 未通过

// 测试点状态
luogu.TestCaseMLEorTLE    // 4  资源超限
luogu.TestCaseWrongAnswer // 6  答案错误
luogu.TestCaseAccepted    // 12 通过

// 编程语言
luogu.LangGo    // 14
luogu.LangCPP14 // 28
```

## 项目结构

```
luoguClient/
├── client.go        # Client 核心、HTTP 请求、配置项
├── auth.go          # AuthService 认证
├── problem.go       # ProblemService 题目
├── record.go        # RecordService 提交记录
├── training.go      # TrainingService 题单
├── types.go         # 所有公开类型
├── constants.go     # 状态/语言常量
├── errors.go        # 错误类型
├── cookiestore.go   # Cookie 导出/导入（仅内存存储）
├── retry.go         # 重试逻辑
└── example/main.go  # 使用示例
```

## 错误类型

```go
&luogu.AuthError{Code: 400, Message: "login failed"}
&luogu.CSRFError{Err: ...}
&luogu.NetworkError{Err: ...}
&luogu.UnauthorizedError{}
```

所有自定义错误支持 `errors.Unwrap()`。

## 许可证

MIT
