package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./store"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x1F4CAD9E864bCDEe21E5242049b0b730Ef78A951")

	instance, err := store.NewStore(address, client)

	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	items, err := instance.Items(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}