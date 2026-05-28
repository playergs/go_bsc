package service

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/w3learn/bindings/v2/token"
	"github.com/w3learn/pkg/client"
)

var _ TokenService = (*tokenService)(nil)

type TokenService interface {
	BalanceOf(address common.Address) (*big.Int, error)
	TotalSupply() (*big.Int, error)
	Stake(amount *big.Int, period uint8) error
	Withdraw(stakeId *big.Int) error
}

type tokenService struct {
	c          *client.EthereumClient
	definition *token.MyTokenStakingRewardsV1
	contract   *bind.BoundContract
}

func NewTokenService(client *client.EthereumClient) TokenService {
	service := &tokenService{c: client}

	// ABI 包装对象
	service.definition = token.NewMyTokenStakingRewardsV1()

	// 合约对象
	service.contract = service.definition.Instance(client.Client, client.StakeAddress)

	return service
}

func (t *tokenService) BalanceOf(address common.Address) (*big.Int, error) {
	var results []any

	if t.contract.Call(&bind.CallOpts{}, &results, "balanceOf", address) != nil {
		return nil, errors.New("failed to call balanceOf")
	}

	balance := abi.ConvertType(results[0], new(big.Int)).(*big.Int)

	log.Printf("balance: %#v", balance)

	return balance, nil
}

func (t *tokenService) TotalSupply() (*big.Int, error) {
	var results []any

	if t.contract.Call(&bind.CallOpts{}, &results, "totalSupply") != nil {
		return nil, errors.New("failed to call totalSupply")
	}

	totalSupply := abi.ConvertType(results[0], new(big.Int)).(*big.Int)

	log.Printf("totalSupply: %#v", totalSupply)

	return totalSupply, nil
}

func (t *tokenService) Stake(amount *big.Int, period uint8) error {
	tx, err := t.contract.Transact(t.c.Auth, "stake", amount, period)

	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), t.c.Client, tx.Hash())

	if err != nil {
		return err
	}

	log.Printf("stake receipt: %#v", receipt)
	return nil
}

func (t *tokenService) Withdraw(stakeId *big.Int) error {
	tx, err := t.contract.Transact(t.c.Auth, "withdraw", stakeId)

	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), t.c.Client, tx.Hash())

	if err != nil {
		return err
	}

	log.Printf("stake receipt: %#v", receipt)
	return nil
}
