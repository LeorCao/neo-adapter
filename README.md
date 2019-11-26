# neo-adapter

neo-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 项目依赖库

- [go-owcrypt](https://github.com/blocktree/go-owcrypt.git)
- [go-owcdrivers](https://github.com/blocktree/.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建NEO.ini文件，编辑如下内容：

```ini

isScan = true
# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url, if RPC Server Type = 0, use neo core full node
serverAPI = "http://127.0.0.1:30333"
# node api url, if RPC Server Type = 1, use neo insight-api
#serverAPI = "http://127.0.0.1::20003/insight-api/"
# RPC Authentication Username
rpcUser = "user-name"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = true
# support segWit
supportSegWit = true
# minimum transaction fees
minFees = "0.00000001"

```
