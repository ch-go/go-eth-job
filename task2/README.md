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
