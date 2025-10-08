package main

import (
	"dapp-task/utils"
	"fmt"
)

func main() {
	fmt.Println("Dapp tasks")
	fmt.Println("---------------------------")

	fmt.Println("任务 1：查询区块")
	//task1.BlockInfo()
	fmt.Println("---------------------------")

	fmt.Println("任务 1：发送 ETH 交易")
	//task1.EthTransaction()
	fmt.Println("---------------------------")

	fmt.Println("任务 2：合约代码生成，部署合约")
	//task1.DeployContractCounter()
	fmt.Println("---------------------------")

	fmt.Println("任务 2：调用合约方法 addCount")
	//task1.CallAddCount()

	utils.AddressAndPrivateKey()

}
