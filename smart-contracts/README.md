# 部署到Sepolia测试网络

## 编辑.env 内容 
在smart-contracts目录中创建“.env”，内容格式：

``` shell 

# Sepolia Testnet Configuration
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID
SEPOLIA_PRIVATE_KEY=0xYOUR_PRIVATE_KEY
```



## 执行部署

``` shell 
# 执行部署
npx hardhat run scripts/deploy-sepolia.ts

```
