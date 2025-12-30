## 任务 1：区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
 具体任务
1. 环境搭建
   - 安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
   - 注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
2. 查询区块
   - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
   - 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
   - 输出查询结果到控制台。
3. 发送交易
   - 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
   - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
   - 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
   - 对交易进行签名，并将签名后的交易发送到网络。
   - 输出交易的哈希值。

## 任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
 具体任务
1. 编写智能合约
   - 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
   - 编译智能合约，生成 ABI 和字节码文件。
2. 使用 abigen 生成 Go 绑定代码
   - 安装 abigen 工具。
   - 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
3. 使用生成的 Go 绑定代码与合约交互
   - 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
   - 调用合约的方法，例如增加计数器的值。
   - 输出调用结果。



# Task2 - Contract Binding

## 初始化go项目
- go mod init task/v
- go mod tidy  

## Steps
1. 编写计数器合约.
   - File: `task2/contract/Counter.sol`

2. 生成abi文件
   ```bash
   solc --abi --bin --overwrite -o task2/contract/build task2/contract/Counter.sol
   ```

3. 利用abigen生成合约代码.
   ```bash
   abigen --abi task2/contract/build/Counter.abi \
     --bin task2/contract/build/Counter.bin \
     --pkg contract --type Counter \
     --out task2/contract/counter.go
   ```
4. 部署合约
   ```bash
   export SEPOLIA_RPC_URL="https://sepolia.infura.io/v3/YOUR_KEY"
   export SEPOLIA_PRIVATE_KEY="YOUR_PRIVATE_KEY"
   go run ./task2/cmd/deploy
   ```
5. 调用合约.
   ```bash
   export SEPOLIA_RPC_URL="https://sepolia.infura.io/v3/YOUR_KEY"
   export SEPOLIA_PRIVATE_KEY="YOUR_PRIVATE_KEY"
   export COUNTER_CONTRACT_ADDRESS="0xYourContractAddress"
   go run ./task2/cmd/task2
   ```

## Notes
- The Go entry point is `task2/task2.go` and uses the generated binding in `task2/contract/counter.go`.
- The `count` read happens immediately; the increment may need confirmations before the updated value is visible.
