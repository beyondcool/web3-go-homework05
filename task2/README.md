# 任务 2：合约代码生成 任务目标

使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。具体任务

    编写智能合约
        使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
        编译智能合约，生成 ABI 和字节码文件。
    使用 abigen 生成 Go 绑定代码
        安装 abigen 工具。
        使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
    使用生成的 Go 绑定代码与合约交互
        编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
        调用合约的方法，例如增加计数器的值。
        输出调用结果。

# 实现

## 生成ABI、go文件

见 `smart-contracts/solc-abi-abigen-go/readme.md`

## 编写调用代码

## 运行

### 配置环境变量

```shell
export SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID
export SEPOLIA_PRIVATE_KEY_OWNER=0xYOUR_PRIVATE_KEY
export SEPOLIA_CONTRACT_ADDRESS=0x

```

### 运行代码

```shell
task2 (main)$ 
task2 (main)$ go run main.go 
RPC_URL: https://sepolia.infura.io/v3/xxxxxxxxx
PRIVATE_KEY_OWNER: xxxxx
CONTRACT_ADDRESS: 0x5D6344f67Ae6d28890fC3BBFD9b8F9e6FdFB2Da2
Current counter value: 10
num IncBy 5
Transaction hash: 0xdf3da80f5351d391c0cb7df7f1627ea0a590b2addb9d11555a4cab4eb67df3e6
Transaction mined in block: 11136293
Current counter value: 15
task2 (main)$ 

```
