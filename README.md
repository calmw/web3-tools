#### 介绍

- Web3开发辅助工具

#### 安装

- 1 如果有go运行环境 ``` go install github.com/calmw/web3-tools@v0.0.2 ```

#### 使用

- 例子：

``` shell
# 获取函数选择器的值
web3-tools getFnSelector 'adminSetClaimType(uint256,uint256)'

output:
    函数签名:                     adminSetClaimType(uint256,uint256)
    函数选择器:                   0xba1c6a10
    
# 对字符串进行keccak256加密，返回16进制字符串
web3-tools keccak256FromStringToHex 'some string'


```