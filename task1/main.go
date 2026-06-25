package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RPC_URL           string
	PRIVATE_KEY_OWNER string
	PRIVATE_KEY_USER1 string
)

func init() {

	/**************************************************
	 *     1.环境搭建
	 *************************************************/

	RPC_URL = os.Getenv("SEPOLIA_RPC_URL")
	PRIVATE_KEY_OWNER = os.Getenv("SEPOLIA_PRIVATE_KEY_OWNER")
	PRIVATE_KEY_USER1 = os.Getenv("SEPOLIA_PRIVATE_KEY_USER1")

	if RPC_URL == "" || PRIVATE_KEY_OWNER == "" || PRIVATE_KEY_USER1 == "" {
		fmt.Fprintf(os.Stderr, "missing required environment variables: RPC_URL, SEPOLIA_PRIVATE_KEY_OWNER or SEPOLIA_PRIVATE_KEY_USER1\n")
		os.Exit(1)
	} else {
		fmt.Println("RPC_URL:", RPC_URL)
		fmt.Println("PRIVATE_KEY_OWNER:", PRIVATE_KEY_OWNER)
		fmt.Println("PRIVATE_KEY_USER1:", PRIVATE_KEY_USER1)
	}
}

func main() {

	/**************************************************
	 *     2.查询区块
	 	编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
		实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
		输出查询结果到控制台。
	 *************************************************/

	// Connect to the Sepolia client
	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		fmt.Println("Failed to connect to the Sepolia client")
		return
	}
	defer client.Close()

	maxBlockNum, _ := client.BlockNumber(context.Background())

	//query the block info by block number
	blockNumber := int64(maxBlockNum) // Replace with the desired block number
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		fmt.Println("Failed to retrieve block:", err)
		return
	}
	fmt.Printf("Block hash: %+v\n", block.Hash())
	fmt.Printf("Block timestamp: %+v\n", block.Time())
	fmt.Printf("Block transactions Len: %+v\n", len(block.Transactions()))

	/**************************************************
	 *     3.发送交易
			准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
			编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
			构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
			对交易进行签名，并将签名后的交易发送到网络。
			输出交易的哈希值。
	 *************************************************/

	// read private key from environment variable
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY_USER1)
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return
	}

	// get the public key and address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get the nonce for the account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(PRIVATE_KEY_OWNER)
	value := big.NewInt(100)  // 100 wei
	gasLimit := uint64(60000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("nonce: %d\n", nonce)
	fmt.Printf("toAddress: %s\n", toAddress.Hex())
	fmt.Printf("gasLimit: %d\n", gasLimit)
	fmt.Printf("gasPrice: %d\n", gasPrice)

	// create the transaction:
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

}
