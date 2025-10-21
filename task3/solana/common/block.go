package common

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

func Block() {
	// 创建RPC客户端（连接到DevNet）
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// 获取最新区块
	recentBlock, err := c.GetBlock(context.Background(), 414503061) // 0表示最新区块
	if err != nil {
		panic("查询失败: " + err.Error())
	}

	fmt.Println("区块高度:", recentBlock.BlockHeight)
	fmt.Println("交易数量:", len(recentBlock.Transactions))
	fmt.Println("区块时间戳：", recentBlock.BlockTime)
	fmt.Println("区块 Hash：", recentBlock.Blockhash)

	for _, reward := range recentBlock.Rewards {
		fmt.Println("奖励类型：", reward.RewardType)
		fmt.Println("奖励 Lam ports", reward.Lamports)
		fmt.Println("奖励 Commission", reward.Commission)
		fmt.Println("奖励 PostBalances", reward.PostBalances)
		fmt.Println("奖励公钥：", reward.Pubkey)
	}

	//for _, tx := range recentBlock.Transactions {
	//	fmt.Println("交易 Meta", tx.Meta)
	//	fmt.Println("交易 账户 keys", tx.AccountKeys)
	//	fmt.Println("交易 签名", tx.Transaction.Signatures)
	//	fmt.Println("交易 Message", tx.Transaction.Message)
	//	fmt.Println("-------------------")
	//}
}
