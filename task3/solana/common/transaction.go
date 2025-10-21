package common

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/btcsuite/btcutil/base58"
)

// SendTransaction 函数用于创建并向Solana网络发送一个转账交易
func SendTransaction() {
	// 创建一个连接到Solana Devnet网络的客户端
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// 获取最新的区块哈希，这是构建交易所需的参数
	// recentBlockHash 用于防止交易重放攻击
	recentBlockHashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		// 如果获取区块哈希失败，记录错误并终止程序
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	// 从私钥生成 Base58 字符串，由于没有前端，所以就这么硬编码代替了。
	// 正常应该是前端请求后端，后端请求用户使用私钥获取 base58 字符串给前端。
	// 前端用 base58 字符串请求后端签名，转账
	privateKeyBytesFrom := base58.Encode([]byte{115, 35, 69, 158, 171, 175, 101, 149, 6, 11, 84, 87, 177, 83, 11, 98, 20, 94, 163, 249, 122, 112, 244, 20, 131, 75, 163, 32, 103, 60, 165, 192, 7, 147, 75, 207, 139, 73, 227, 117, 109, 58, 242, 87, 118, 247, 25, 131, 82, 205, 175, 202, 123, 126, 35, 75, 40, 183, 193, 252, 73, 7, 5, 124})
	privateKeyBytesTo := base58.Encode([]byte{100, 52, 187, 121, 140, 8, 101, 180, 224, 204, 60, 181, 15, 11, 206, 248, 158, 219, 30, 224, 118, 12, 33, 140, 236, 148, 41, 15, 36, 137, 147, 7, 117, 110, 5, 29, 252, 62, 142, 34, 155, 229, 108, 12, 136, 32, 225, 93, 4, 186, 148, 97, 60, 24, 103, 45, 2, 245, 193, 225, 204, 114, 241, 210})

	fromAccount, err := types.AccountFromBase58(privateKeyBytesFrom)
	if err != nil {
		log.Fatal("from account err", err)
	}

	toAccount, err := types.AccountFromBase58(privateKeyBytesTo)
	if err != nil {
		log.Fatal("to account err", err)
	}

	// 创建一个新的交易，包含转账指令
	tx, err := types.NewTransaction(types.NewTransactionParam{
		// 指定交易签名者，包括费用支付者和转账发起者
		Signers: []types.Account{fromAccount},
		// 创建交易消息
		Message: types.NewMessage(types.NewMessageParam{
			// 指定费用支付者的公钥
			FeePayer: fromAccount.PublicKey,
			// 使用获取到的最新区块哈希
			RecentBlockhash: recentBlockHashResponse.Blockhash,
			// 定义交易指令列表，这里是一个系统转账指令
			Instructions: []types.Instruction{
				// 创建系统转账指令，从Alice账户向指定地址转账
				system.Transfer(system.TransferParam{
					// 转账发起方的公钥
					From: fromAccount.PublicKey,
					// 转账接收方的公钥（这里是费用支付者地址）
					To: toAccount.PublicKey,
					// 转账金额，1e8 = 100,000,000 lamports = 0.1 SOL
					Amount: 1e8, // 0.1 SOL
				}),
			},
		}),
	})
	if err != nil {
		// 如果创建交易失败，记录错误并终止程序
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	// 将交易发送到Solana网络
	txHash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		// 如果发送交易失败，记录错误并终止程序
		log.Fatalf("failed to send tx, err: %v", err)
	}

	// 打印交易哈希，用于在区块链浏览器上查询交易状态
	fmt.Println("Transaction Hash:", txHash)
}
