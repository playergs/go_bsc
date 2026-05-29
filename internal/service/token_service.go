package service

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/w3learn/bindings/v2/token"
	"github.com/w3learn/pkg/client"
)

var _ TokenService = (*tokenService)(nil)

type TokenService interface {
	BalanceOf(address common.Address) (*big.Int, error)
	TotalSupply() (*big.Int, error)
	Stake(ctx context.Context, amount *big.Int, period uint8) error
	Withdraw(ctx context.Context, stakeId *big.Int) error

	QueryHistoricalEvents(ctx context.Context) error
	SubscribeEvents(ctx context.Context, fromBlock *big.Int) error
	SubscribeNewBlocks(ctx context.Context) error
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

func (t *tokenService) Stake(ctx context.Context, amount *big.Int, period uint8) error {
	tx, err := t.contract.Transact(t.c.Auth, "stake", amount, period)

	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, t.c.Client, tx.Hash())

	if err != nil {
		return err
	}

	log.Printf("stake receipt: %#v", receipt)
	return nil
}

func (t *tokenService) Withdraw(ctx context.Context, stakeId *big.Int) error {
	tx, err := t.contract.Transact(t.c.Auth, "withdraw", stakeId)

	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, t.c.Client, tx.Hash())

	if err != nil {
		return err
	}

	log.Printf("stake receipt: %#v", receipt)
	return nil
}

func (t *tokenService) QueryHistoricalEvents(ctx context.Context) error {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(110180100),
		ToBlock:   nil,
		Addresses: []common.Address{t.c.StakeAddress},
		Topics: [][]common.Hash{{
			crypto.Keccak256Hash([]byte("Staked(address,uint256)")),
			crypto.Keccak256Hash([]byte("Withdrawn(address,uint256)")),
		}},
	}

	logs, err := t.c.Client.FilterLogs(ctx, query)

	if err != nil {
		return err
	}

	log.Printf("==========历史Log共 %d 条==========\n", len(logs))
	for _, eLog := range logs {
		switch eLog.Topics[0].Hex() {
		case crypto.Keccak256Hash([]byte("Staked(address,uint256)")).Hex():
			event, err := t.definition.UnpackStakedEvent(&eLog)
			if err != nil {
				return err
			}
			log.Printf("Staked Event: %#v", event)
		case crypto.Keccak256Hash([]byte("Withdrawn(address,uint256)")).Hex():
			event, err := t.definition.UnpackWithdrawnEvent(&eLog)
			if err != nil {
				return err
			}
			log.Printf("Withdrawn Event: %#v", event)
		}
	}
	log.Printf("==================================\n")

	return nil
}

func (t *tokenService) SubscribeNewBlocks(ctx context.Context) error {

	headerCh := make(chan *types.Header)

	sub, err := t.c.Client.SubscribeNewHead(ctx, headerCh)

	if err != nil {
		log.Printf("SubscribeNewBlocks failed: %s", err)
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case err := <-sub.Err():
				log.Printf("SubscribeNewBlocks error: %s", err)
				return
			case header := <-headerCh:
				log.Printf("New block: block number: #%d,  Hash: %s, Time: %d, all: %#v", header.Number, header.Hash().Hex(), header.Time, header)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (t *tokenService) SubscribeEvents(ctx context.Context, fromBlock *big.Int) error {
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: []common.Address{t.c.StakeAddress},
		Topics: [][]common.Hash{{
			crypto.Keccak256Hash([]byte("Staked(address,uint256)")),
			crypto.Keccak256Hash([]byte("Withdrawn(address,uint256)")),
		}},
	}

	logsCh := make(chan types.Log)

	sub, err := t.c.Client.SubscribeFilterLogs(ctx, query, logsCh)

	if err != nil {
		return err
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("SubscribeEvents error: %s", err)
				return
			case <-ctx.Done():
				return
			case logs := <-logsCh:
				switch logs.Topics[0].Hex() {
				case crypto.Keccak256Hash([]byte("Staked(address,uint256)")).Hex():
					event, err := t.definition.UnpackStakedEvent(&logs)
					if err != nil {
						log.Printf("UnpackStakedEvent error: %s", err)
					}
					log.Printf("staked event: %#v", event)
				case crypto.Keccak256Hash([]byte("Withdrawn(address,uint256)")).Hex():
					event, err := t.definition.UnpackWithdrawnEvent(&logs)
					if err != nil {
						log.Printf("UnpackWithdrawnEvent error: %s", err)
					}
					log.Printf("withdrawn event: %#v", event)
				}
			}
		}
	}()
	_ = query
	return nil
}
