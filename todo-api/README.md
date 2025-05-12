# 📝 Todo API (Go)

一個使用 Go 語言建構的簡單待辦事項 RESTful API，後端採用 Microsoft SQL Server，具備分層式架構，便於維護與擴充。

## 📁 專案結構
```bash
todo-api/
├── go.mod
├── go.sum
├── main.go               // 專案入口，啟動伺服器
├── .env                  // 環境變數
├── config/               // 設定相關（.env 讀取、組態）
│   └── config.go
├── model/                // 資料模型（struct 定義）
│   └── todo.go
├── handler/              // HTTP 路由處理邏輯
│   └── todo_handler.go
├── service/              // 商業邏輯（CRUD、驗證邏輯）
│   └── todo_service.go
├── router/               // 路由註冊集中處理
│   └── router.go
├── middleware/           // Middleware（驗證、日誌等）
│   └── logger.go
└── util/                 // 工具、共用小功能
    └── response.go
```

## 🚀 快速開始

### 1️⃣ 安裝依賴套件

```bash
go mod tidy
go run main.go
```