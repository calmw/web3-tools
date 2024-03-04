### 函数选择器，Keccak256加密使用手册

- ##### 获取函数选择器的值,返回16进制字符串

``` shell

web3 getFnSelector --funcSig "adminSetClaimType(uint256,uint256)"

output:
    函数签名:                     adminSetClaimType(uint256,uint256)
    函数选择器:                   0xba1c6a10
```   

- #### 对字符串进行keccak256加密，返回16进制字符串

``` shell
web3 keccak256FromStringToHex --someStr "some string"
```  
