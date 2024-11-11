# Gin 開發框架

## go
```
go mod init
go mod tidy
swag init -d ./,./internal/controller

go run main.go

go build -o app
```

## 分層設計

```
/
|- config           # 設定檔         (格式勿動，可修value)
|- docs                 # swagger        (勿動)
|- internal             # 主程式
|   |- controller       # api 路由       (開發者)
|   |   |- v1           # api 版本迭代
|   |   |   |- 01-post-xxx.go    # api
|   |   |- v2
|   |- init             # 起始化          (管理者)
|   |   |- init.go      # 起始化          (管理者)
|   |   |- struct.go    # 起始化          (管理者)
|   |- lib              # 中間件          (管理者)
|   |- model            # 資料結構        (開發者)
|   |   |- 01-struct.go
|   |- service          # 業務流程        (開發者)
|   |   |- 01-service
|- manifest             # 交付項目
|   |- test             # unit test      (開發者)
|- view                 # 靜態檔案
|   |- css
|   |- html
|   |- image
|   |- js
|- go.mod               # 依賴管理        (勿動)
|- go.sum               # lib            (勿動)
|- main.go              # 入口點          (開發者)
|- Dockerfile           # 容器化腳本       (管理者)
|- README.md            # 需編輯API說明    (開發者)

```

以下是API說明時的格式

# xxx Project

## List

* xxx API
    * Summary
    * input
    ```
    {
        "age": int,
        "email": str,
        "name": "str"
    }
    ```
    * output
    ```
    {
        "ID": int,
        "CreatedAt": timestamp,
        "UpdatedAt": timestamp,
        "DeletedAt": timestamp,
        "name": str,
        "email": str,
        "age": int,
        "create_time": timestamp
    }
    ```
* xxx API
    * Summary
    * input
    ```
    
    ```
    * output
    ```
    
    ```
