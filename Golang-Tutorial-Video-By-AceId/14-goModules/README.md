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
