package common

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func Subscription() {
	// 创建WebSocket客户端（连接到DevNet）
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic("连接失败: " + err.Error())
	}

	// 将字符串转换为 solana.PublicKey
	programID := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

	// 订阅程序日志（示例：SPL Token程序）
	sub, err := wsClient.LogsSubscribeMentions(
		programID,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		panic("订阅失败: " + err.Error())
	}
	defer sub.Unsubscribe()

	// 实时处理事件
	for {
		log, err := sub.Recv(context.Background())
		if err != nil {
			fmt.Println("监听错误: ", err)
			return
		}

		fmt.Println("[程序日志] 签名: ", log.Value.Signature)
		fmt.Println("      日志内容: ", log.Value.Logs)
	}
}
