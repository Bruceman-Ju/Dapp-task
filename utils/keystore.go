package utils

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
)

func AddressAndPrivateKey() (string, string) {
	// Read key from file.
	// 下面的路径是相对于项目根目录
	keyJson, err := os.ReadFile("keystore/UTC--2025-10-08T15-43-48.880806000Z--7772936d8812dfc65f7f4135727f480f6db81bf6")
	if err != nil {
		utils.Fatalf("Failed to read the keyfile at ../keystore: %v", err)
	}

	// Decrypt key with passphrase.
	passphrase := ""
	key, err := keystore.DecryptKey(keyJson, passphrase)
	if err != nil {
		utils.Fatalf("Error decrypting key: %v", err)
	}

	address := key.Address.Hex()
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	fmt.Println(fmt.Sprintf("address: %s,\nprivateKye: %s", address, privateKey))

	return address, privateKey
}
