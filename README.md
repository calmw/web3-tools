#### 介绍

- Web3开发辅助工具

#### 安装

- 1 如果有go运行环境 ``` go install github.com/calmw/web3-tools@v0.0.2 ```

#### 使用

- 例子：

``` shell
# 获取函数选择器的值,返回16进制字符串
web3-tools getFnSelector 'adminSetClaimType(uint256,uint256)'

output:
    函数签名:                     adminSetClaimType(uint256,uint256)
    函数选择器:                   0xba1c6a10
    
# 对字符串进行keccak256加密，返回16进制字符串
web3-tools keccak256FromStringToHex 'some string'
    
# 获取ERC20代币余额
web3-tools  erc20Balance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E
    
# ERC20代币Approve授权
web3-tools erc20Approve --rpc https://lisbon-testnet-rpc.matchtest.co --chainId 698 --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --key 0x88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e --spender 0x7D05883B19cD14c216AD82222c48bCD4eE98914c --amount 10000
    
# 查询ERC20代币Allowance
web3-tools erc20Approve erc20Allowance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --owner 0x360C815e8C5F130913113801D0c57611Ee95723A  --spender 0x7D05883B19cD14c216AD82222c48bCD4eE98914c



```