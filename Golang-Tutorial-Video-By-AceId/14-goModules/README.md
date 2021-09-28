## 使用 Go Module 初始化项目

保证 go env `GO111MODULE` 的值不为 `off`

1. 初始化项目，生成 `go.mod` 文件；
    ```
    go mod init <module_name>
    ```
2. 编写代码 `main.go`；
3. （可选）下载依赖，生成 `go.sum` 文件；
    ```
    go get
    ```
4. 执行
    ```
    go run main.go
    ```

## 修改项目模块的版本依赖关系

将 2.18.0 版本作为 2.19.0 版本使用

```
go mod edit -replace=github.com/gofiber/fiber/v2@v2.19.0=github.com/gofiber/fiber/v2@v2.18.0
```
