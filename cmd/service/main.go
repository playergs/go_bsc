package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/w3learn/bindings/token"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// 初始化客户端
	client, err := ethclient.Dial(os.Getenv("BSC_TEST_RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the BSC network: %v", err)
	}

	defer client.Close()

	// 私钥字符串通用处理
	pkStr := os.Getenv("PRIVATE_KEY")

	if strings.HasPrefix(pkStr, "0x") {
		pkStr = pkStr[2:]
	}

	// 将字符串私钥转换成ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(pkStr)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取crypto.PublicKey公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 公钥地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("Public key: %s\n", fromAddress.Hex())

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	fmt.Printf("Chain ID: %d\n", chainID)

	fmt.Println("Go Ethereum SDK初始化完成")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// TODO: 发起交易时使用
	_ = auth

	contractAddress := common.HexToAddress(os.Getenv("ERC20_CONTRACT_STAKING_REWARDS"))

	fmt.Printf("Contract address: %s\n", contractAddress.Hex())

	myTokenStakingRewardsV1, err := token.NewMyTokenStakingRewardsV1(contractAddress, client)

	if err != nil {
		log.Fatalf("Failed to create staking reward v1: %v", err)
	}

	totalSupply, err := myTokenStakingRewardsV1.TotalSupply(&bind.CallOpts{Context: context.Background()})

	if err != nil {
		log.Fatalf("Failed to get total supply: %v", err)
	}

	fmt.Printf("Total supply: %s\n", totalSupply.String())

	info, err := myTokenStakingRewardsV1.StakeIdToInfo(&bind.CallOpts{Context: context.Background()}, big.NewInt(1))

	if err != nil {
		log.Fatalf("Failed to get staking info: %v", err)
	}

	fmt.Printf("info: %#v\n", info)

	//amount, ok := new(big.Int).SetString("25000000000000000000000000", 10)
	//
	//if !ok {
	//	log.Fatalf("Failed to set amount")
	//	return
	//}

	// 质押
	//stakingError := stake(context.Background(), client, myTokenStakingRewardsV1, auth, amount, 3)
	//
	//if stakingError != nil {
	//	log.Fatalf("Failed to stake: %v", stakingError)
	//	return
	//}

	// 查询质押量
	balance, err := balanceOf(myTokenStakingRewardsV1, fromAddress)

	fmt.Printf("Balance: %s\n", balance.String())

	// 解质押
	withdrawError := withdraw(context.Background(), client, myTokenStakingRewardsV1, auth, big.NewInt(5))

	if withdrawError != nil {
		log.Fatalf("Failed to withdraw: %v", withdrawError)
		return
	}

	// 查询质押量
	balance, err = balanceOf(myTokenStakingRewardsV1, fromAddress)

	fmt.Printf("Balance: %s\n", balance.String())
}

func balanceOf(
	myTokenStakingRewardsV1 *token.MyTokenStakingRewardsV1,
	address common.Address,
) (*big.Int, error) {
	balance, err := myTokenStakingRewardsV1.BalanceOf(&bind.CallOpts{Context: context.Background()}, address)

	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
		return big.NewInt(0), err
	}

	return balance, nil
}

func stake(
	ctx context.Context,
	client *ethclient.Client,
	myTokenStakingRewardsV1 *token.MyTokenStakingRewardsV1,
	auth *bind.TransactOpts,
	amount *big.Int,
	period uint8,
) error {
	transaction, err := myTokenStakingRewardsV1.Stake(auth, amount, period)

	if err != nil {
		log.Fatalf("Failed to stake: %v", err)
		return err
	}

	fmt.Printf("transaction: %#v\n", transaction)

	receipt, err := bind.WaitMined(ctx, client, transaction)

	if err != nil {
		log.Fatalf("Failed to wait transaction: %v", err)
		return err
	}

	fmt.Printf("receipt: %#v\n", receipt)

	return nil
}

func withdraw(
	ctx context.Context,
	client *ethclient.Client,
	myTokenStakingRewardsV1 *token.MyTokenStakingRewardsV1,
	auth *bind.TransactOpts,
	stakeId *big.Int,
) error {
	transaction, err := myTokenStakingRewardsV1.Withdraw(auth, stakeId)

	if err != nil {
		log.Fatalf("Failed to withdraw: %v", err)
		return err
	}

	fmt.Printf("transaction: %#v\n", transaction)

	receipt, err := bind.WaitMined(ctx, client, transaction)

	if err != nil {
		log.Fatalf("Failed to wait transaction: %v", err)
		return err
	}

	fmt.Printf("receipt: %#v\n", receipt)

	return nil
}
