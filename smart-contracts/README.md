# 部署到Sepolia测试网络

## 设置环境变量

``` shell 
# Sepolia Testnet Configuration
export SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID
export SEPOLIA_PRIVATE_KEY=0xYOUR_PRIVATE_KEY
```



## 执行部署

``` shell 
# 执行部署
npx hardhat run scripts/deploy-sepolia.ts

```

<b>示例：</b>
``` shell
smart-contracts (main)$ npx hardhat run scripts/deploy-sepolia.ts

Deploying Counter to Sepolia...
Counter deployed to: 0x5D6344f67Ae6d28890fC3BBFD9b8F9e6FdFB2Da2
Transaction hash: 0xaf8bfc8460269bf4dba70af05eb33216802c91ec2cdb7b417c117d5894fd215f
smart-contracts (main)$ 

```
