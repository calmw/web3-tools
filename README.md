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
web3-tools erc20Approve --key 88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e --chainId 698 --spender 0x20142400c97AD2c0A74d4C0400e482a973087033  --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570 --wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E --amount 100


```