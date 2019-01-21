package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"./token"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ebaf1785cc1b4f319e0ff07f26cadae8")

	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x997e5f74F4C398A804a4a0AD1946DC8c95870c1e")
	bal, err := instance.BalanceOf(&bind.TransactOpts{}, address)

	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("jedl")

	name, err := instance.Name(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.Hash().String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
