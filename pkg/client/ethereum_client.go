package client

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/w3learn/internal/config"
)

type EthereumClient struct {
	Client        *ethclient.Client
	ChainID       *big.Int
	PublicAddress common.Address
	Auth          *bind.TransactOpts
	TokenAddress  common.Address
	StakeAddress  common.Address
}

func NewClient(
	ctx context.Context,
	cfg *config.EthereumConfig,
) (*EthereumClient, error) {
	// 创建Ethereum Client
	c := &EthereumClient{}

	// 连接RPC
	var err error = nil
	c.Client, err = ethclient.DialContext(ctx, cfg.RpcUrl)

	if err != nil {
		return nil, err
	}

	// 私钥字符串通用处理
	pkStr := cfg.PrivateKey

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

	// 获取公钥地址
	c.PublicAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("Public key: %s\n", c.PublicAddress.Hex())

	// 获取链ID
	c.ChainID, err = c.Client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	log.Printf("Chain ID: %d\n", c.ChainID)

	// 获取交易对象
	c.Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, c.ChainID)

	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// 合约地址
	c.TokenAddress = common.HexToAddress(cfg.Contracts.TokenAddress)
	c.StakeAddress = common.HexToAddress(cfg.Contracts.StakeAddress)

	return c, nil
}
