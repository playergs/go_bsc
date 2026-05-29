package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/w3learn/internal/config"
	"github.com/w3learn/internal/service"
	"github.com/w3learn/pkg/client"
)

var fromBlock = big.NewInt(110180100)

func main() {
	appConfig, err := config.LoadAppConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("App Config: %#v\n", appConfig)

	c, err := client.NewClient(context.Background(), &appConfig.Blockchain.Ethereum)

	if err != nil {
		log.Fatal(err)
	}

	defer c.Client.Close()

	log.Printf("client: %#v\n", c)

	tokenService := service.NewTokenService(c)

	// 查询
	balance, err := tokenService.BalanceOf(c.PublicAddress)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("balance: %#v\n", balance)

	totalSupply, err := tokenService.TotalSupply()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("totalSupply: %#v\n", totalSupply)

	err = tokenService.QueryHistoricalEvents(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = tokenService.SubscribeNewBlocks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = tokenService.SubscribeEvents(context.Background(), fromBlock)

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	//质押 & 查询
	//amount, ok := new(big.Int).SetString("25000000000000000000000000", 10)
	//
	//if !ok {
	//	log.Fatalf("Failed to set amount")
	//	return
	//}
	//
	//err = tokenService.Stake(context.Background(), amount, 0)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//balance, err = tokenService.BalanceOf(c.PublicAddress)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf("balance: %#v\n", balance)
	//
	//totalSupply, err = tokenService.TotalSupply()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf("totalSupply: %#v\n", totalSupply)
	//
	//time.Sleep(40 * time.Second)

	//解质押 & 查询
	err = tokenService.Withdraw(context.Background(), big.NewInt(12))

	if err != nil {
		log.Fatal(err)
	}

	balance, err = tokenService.BalanceOf(c.PublicAddress)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("balance: %#v\n", balance)

	totalSupply, err = tokenService.TotalSupply()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("totalSupply: %#v\n", totalSupply)

	// 防止主协程退出
	select {}
}
