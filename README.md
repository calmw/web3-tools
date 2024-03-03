#### 介绍

- Web3开发辅助工具

#### 安装

- 如果有go运行环境，将GOPATH下的bin目录加入环境变量
    - 执行``` go install github.com/calmw/web3-tools@v0.0.4 ```

- linux/mac

```shell
    # ADM
    wget https://github.com/calmw/web3-tools/releases/download/v0.0.4/web3-tools_amd64
    ln -s web3-tools_amd64 /usr/local/bin/web3
    # Apple 
    wget https://github.com/calmw/web3-tools/releases/download/v0.0.4/web3-tools_arm64
    ln -s web3-tools_arm64 /usr/local/bin/web3
```

- windows
    - 直接下载最新的二进制文件使用,下载地址：https://github.com/calmw/web3-tools/releases/download/v0.0.4/web3-tools.exe

#### 使用

- 函数选择器，Keccak256加密, [函数选择器,Keccak256加密使用手册](./docs/usage/sign_fn_selector.md)
- ERC20手册: [ERC20手册](./docs/usage/erc20.md)
- Native Coin(ETH)使用手册:[Native Coin使用手册](./docs/usage/native_coin.md)
- Hardhat增强工具：[Native Coin使用手册](./docs/usage/hardhat_tools.md)

