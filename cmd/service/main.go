package main

import (
	"context"
	"log"
	"math/big"

	"github.com/w3learn/internal/config"
	"github.com/w3learn/internal/service"
	"github.com/w3learn/pkg/client"
)

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

	//// 质押 & 查询
	//amount, ok := new(big.Int).SetString("25000000000000000000000000", 10)
	//
	//if !ok {
	//	log.Fatalf("Failed to set amount")
	//	return
	//}
	//
	//err = tokenService.Stake(amount, 3)
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

	// 解质押 & 查询
	err = tokenService.Withdraw(big.NewInt(6))

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
}
