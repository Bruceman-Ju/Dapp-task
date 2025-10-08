package task1

import (
	"crypto/ecdsa"
	"dapp-task/utils"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
)

/*
任务 1：区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
具体任务
	环境搭建
		安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
		注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
	查询区块
		1. 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
		2. 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
		3. 输出查询结果到控制台。
	发送交易
		1. 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
		2. 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
		3. 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
		4. 对交易进行签名，并将签名后的交易发送到网络。
		5. 输出交易的哈希值。
*/

// BlockInfo 查询区块
func BlockInfo() {

	client, err := EthClient(HTTP)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	blockNumber := big.NewInt(9318079)

	header, errHeader := client.HeaderByNumber(context.Background(), blockNumber)
	if errHeader != nil {
		log.Fatal(errHeader)
	}
	fmt.Println("当前区块头：", header.Number.Uint64())
	fmt.Println("当前区块头时间戳：", header.Time)
	fmt.Println("当前区块头难度", header.Difficulty)
	fmt.Println("当前区块头 Hash", header.Hash().Hex())
	fmt.Println("当前区块头 Nonce：", header.Nonce)
	fmt.Println("当前区块头 ParentHash：", header.ParentHash.Hex())
	fmt.Println("当前区块头 矿工：", header.Coinbase.Hex())
	fmt.Println("当前区块头 GasLimit：", header.GasLimit)
	fmt.Println("当前区块头燃料费用：", header.GasUsed)

	fmt.Println("----------------")

	block, errBlock := client.BlockByNumber(context.Background(), blockNumber)
	if errBlock != nil {
		log.Fatal(errBlock)
	}
	fmt.Println("当前区块：", block.NumberU64())
	fmt.Println("当前区块时间戳：", block.Time())
	fmt.Println("当前区块难度：", block.Difficulty().Uint64())
	fmt.Println("当前区块 Hash", block.Hash().Hex())
	fmt.Println("当前区块交易数：", len(block.Transactions()))
	fmt.Println("当前区块父区块 Hash：", block.ParentHash().Hex())
	fmt.Println("当前区块矿工：", block.Coinbase().Hex())
	fmt.Println("当前区块 Nonce：", block.Nonce())
	fmt.Println("当前区块 GasLimit：", block.GasLimit())
	fmt.Println("当前区块燃料费用：", block.GasUsed())

	count, errCount := client.TransactionCount(context.Background(), block.Hash())
	if errCount != nil {
		log.Fatal(errCount)
	}
	fmt.Println("当前区块交易数：", count)
}

func EthTransaction() {

	client, err := EthClient(HTTP)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	_, privateKeyStr := utils.AddressAndPrivateKey()

	privateKey, errPrivateKey := crypto.HexToECDSA(privateKeyStr)
	if errPrivateKey != nil {
		log.Fatal(errPrivateKey)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("From address:", fromAddress)

	nonce, errNonce := client.PendingNonceAt(context.Background(), fromAddress)
	if errNonce != nil {
		log.Fatal(errNonce)
	}
	fmt.Println("交易随机数：", nonce)

	precision := new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil)
	value := new(big.Int).Mul(big.NewInt(1), precision)

	gasLimit := uint64(21000)

	gasPrice, errGasPrice := client.SuggestGasPrice(context.Background())
	if errGasPrice != nil {
		log.Fatal(errGasPrice)
	}

	toAddress := common.HexToAddress("0x6EaC8F7C42fe83285fFaDE11911028D7cD421f28")

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainId, errChainId := client.NetworkID(context.Background())
	if errChainId != nil {
		log.Fatal(errChainId)
	}

	signedTx, errSignTx := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if errSignTx != nil {
		log.Fatal(errSignTx)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

}
