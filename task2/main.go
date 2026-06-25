package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"task2/counter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RPC_URL           string
	PRIVATE_KEY_OWNER string
	// PRIVATE_KEY_USER1 string
	CONTRACT_ADDRESS string
)

func init() {

	/**************************************************
	 *     1.环境搭建
	 *************************************************/

	RPC_URL = os.Getenv("SEPOLIA_RPC_URL")
	PRIVATE_KEY_OWNER = os.Getenv("SEPOLIA_PRIVATE_KEY_OWNER")
	// PRIVATE_KEY_USER1 = os.Getenv("SEPOLIA_PRIVATE_KEY_USER1")
	CONTRACT_ADDRESS = os.Getenv("SEPOLIA_CONTRACT_ADDRESS")
	if RPC_URL == "" || PRIVATE_KEY_OWNER == "" || CONTRACT_ADDRESS == "" {
		fmt.Fprintf(os.Stderr, "missing required environment variables: RPC_URL, SEPOLIA_PRIVATE_KEY_OWNER, or SEPOLIA_CONTRACT_ADDRESS\n")
		os.Exit(1)
	} else {
		fmt.Println("RPC_URL:", RPC_URL)
		fmt.Println("PRIVATE_KEY_OWNER:", PRIVATE_KEY_OWNER)
		fmt.Println("CONTRACT_ADDRESS:", CONTRACT_ADDRESS)
	}
}

func main() {

	client, err := ethclient.Dial(RPC_URL)
	if err != nil {
		panic(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	counter, err := counter.NewCounter(contractAddress, client)

	if err != nil {
		panic(err)
	}

	/********************* 查询当前计数器的值 *********************/

	// 查询当前计数器的值
	printNum(counter)

	/********************* 调用合约方法：IncBy(5) *********************/

	// options for the transaction
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY_OWNER)
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	opts.GasLimit = uint64(60000) // set gas limit
	opts.GasPrice = gasPrice      // set gas price
	opts.Nonce = big.NewInt(int64(nonce))

	tx, err := counter.IncBy(opts, big.NewInt(5))
	if err != nil {
		panic(err)
	}
	fmt.Println("num IncBy 5")
	fmt.Println("Transaction hash:", tx.Hash())

	// 等待区块生效
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("transaction receipt is nil")
	}
	fmt.Println("Transaction mined in block:", receipt.BlockNumber.Uint64())

	/********************* 查询当前计数器的值 *********************/

	printNum(counter)

}

func printNum(counter *counter.Counter) {

	num, _ := counter.Num(&bind.CallOpts{})

	fmt.Println("Current counter value:", num)
}
