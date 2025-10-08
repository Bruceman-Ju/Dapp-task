package task1

import (
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EthClient(requestType uint) (*ethclient.Client, error) {
	if requestType == HTTP {
		client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/4gZRRWrE9vi0drk3L37D8")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return client, err
	}
	if requestType == WebSocket {
		client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/4gZRRWrE9vi0drk3L37D8")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return client, err
	}
	return nil, errors.New("invalid request type for eth client")
}

const (
	HTTP = iota
	WebSocket
)
