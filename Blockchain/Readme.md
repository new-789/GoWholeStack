## 使用方法
#### 一、使用下列命令将源码克隆到本地：
```bash
git clone https://github.com/new-789/GoWholeStack/tree/master/Blockchain
```
#### 二、运行 `run.sh` 脚本文件，运行后会在当前目录生成以下文件，并输出如下内容：
```bash
rm: 无法删除'blockChain': 没有那个文件或目录
rm: 无法删除'*.db': 没有那个文件或目录
开始挖矿。。。。。。。
挖矿成功, hash:000005e98ee4061bf9604cbf8280df62ca79c9bcbe6c404c0fc3e4dc3dca7cde, nonce: 1782963

        printChan                               "正向打印区块链信息"
        printChainR                             "反向打印区块链信息"
        getBalance --address ADDRESS "获取指定地址的余额"
        send FROM TO AMOUNT MINER DATA "由  from 转 amount 给 to，由 miner 挖矿，同时写入 data"
        newWallet                               "创建一个新的钱包(私钥公钥对)"
        listAddresses                   "列举所有钱包地址"
```
- 2.1：blockChain ：编译后的区块链程序二进制文件，使用该文件加上参数即可使用本程序
- 2.2：blockChain.db ：区块链数据存储文件
#### 三、使用 `./blockChain newWallet` 生成几个新的钱包地址
#### 四、使用 `/blockChain listAddresses` 命令查看已存在的钱包地址
#### 五、通过 `./backChain send --address` 命令使用创世快 `1LFEaXkHpDKvAu6WtfTqPwSs15EyAPg6nF` 
给指定的钱包地址进行转账和转账金额，并指定另一个钱包地址进行挖矿。

#### 更详细的使用参见初次运行时的命令帮助，本程序为区块链练习代码，功能尚不完善，仅供娱乐切勿当真！！
