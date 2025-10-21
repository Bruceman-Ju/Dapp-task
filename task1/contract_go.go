package task1

import (
	"context"
	counter "dapp-task/task1/contracts/generated_go"
	"dapp-task/utils"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
具体任务
	编写智能合约
		1. 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
		2. 编译智能合约，生成 ABI 和字节码文件。
	使用 abigen 生成 Go 绑定代码
		1. 安装 abigen 工具。
		2. 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
	使用生成的 Go 绑定代码与合约交互
		1. 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
		2. 调用合约的方法，例如增加计数器的值。
		3. 输出调用结果。
*/

// DeployContractCounter Sepolia 上部署合约
func DeployContractCounter() {

	client, err := EthClient(HTTP)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer client.Close()

	_, privateKeyStr := utils.AddressAndPrivateKey()

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	// Retrieve the current chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve chain ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, tx, _, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

}

func CallAddCount() {

	client, err := EthClient(HTTP)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer client.Close()

	_, privateKeyStr := utils.AddressAndPrivateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xf6db9e14e69134dae0ab14569f12b489ced5a259")

	counterInstance, err := counter.NewCounter(tokenAddress, client)
	if err != nil {
		log.Fatalf(err.Error())
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("创建交易认证失败: ", err)
	}

	_, err = counterInstance.AddCount(transactOpts, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	waitSeconds := 15

	fmt.Println("waiting", waitSeconds, "seconds for contract execution......")

	for i := 0; i < waitSeconds; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(waitSeconds-i, " seconds remaining")
	}

	result, err := counterInstance.CurrentCount(&bind.CallOpts{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Current Count: %s\n", result)
}
