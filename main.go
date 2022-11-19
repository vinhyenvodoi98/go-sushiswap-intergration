package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	uniswap "github.com/vinhyenvodoi98/go-sushiswap-intergration/SmartContract/uniswap"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main () {
	godotenv.Load(".env")
	client, err := ethclient.Dial(os.Getenv("NETWORK_ENDPOINTS"))
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0xE9a889E6963f122a98f8083d951c71329c726c0A")
    instance, err := uniswap.NewUniswap(address, client)
    if err != nil {
        log.Fatal(err)
    }

    token0, err := instance.Token0(nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(token0) // "1.0"
}
