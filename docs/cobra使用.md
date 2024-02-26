#### 生成代码

- 创建项目 ``` go mod init project_name ```
- cobra代码生成工具安装,参考文档，注意将go bin目录加入环境变量 ``` go install github.com/spf13/cobra-cli@latest ```
- 生成代码示例 ``` cobra-cli init --author "calm.wang@hotmail.com" --license MIT ```
- 添加命令

``` shell
# 添加config命令，命令名称使用驼峰命名法
cobra-cli add config  # go run main.go config

# 添加config命令的子命令。默认父命令是rootCmd
cobra-cli add create -p 'configCmd'  # go run main.go config create
```

- 修改父命令的信息，在root.go

#### 开发注意事项

- cmd中命令的 init flag 定义一次即可，多个文件中重复定义会报错
- erc20BalanceCmd.Flags()           // 全局有效
  erc20ApproveCmd.PersistentFlags() // 当前命令有效