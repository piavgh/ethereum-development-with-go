package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x1F4CAD9E864bCDEe21E5242049b0b730Ef78A951")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode))
}
