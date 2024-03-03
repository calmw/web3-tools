### ERC20使用手册

- #### 获取ERC20代币余额

``` shell
web3 erc20Balance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570
--wallet 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E
```

- #### ERC20代币Approve授权

``` shell
web3 erc20Approve --rpc https://lisbon-testnet-rpc.matchtest.co --chainId 698 --token
0x3eE243ff68074502b1D9D65443fa97b99f634570 --key 0x88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e
--spender 0x7D05883B19cD14c216AD82222c48bCD4eE98914c --amount 10000
```

- #### 查询ERC20代币Allowance

``` shell
web3 erc20Allowance --rpc https://lisbon-testnet-rpc.matchtest.co  --token 0x3eE243ff68074502b1D9D65443fa97b99f634570
--owner 0x360C815e8C5F130913113801D0c57611Ee95723A --spender 0x7D05883B19cD14c216AD82222c48bCD4eE98914c
```

- #### ERC20代币转账

``` shell
web3 erc20Transfer --rpc https://lisbon-testnet-rpc.matchtest.co --chainId 698 --token
0x3eE243ff68074502b1D9D65443fa97b99f634570 --key 0x88efa0c968693a034301450d450c7169b3f608966977a4e3b5b0672cb82c597e --to
0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E --transferAmount 1
```
