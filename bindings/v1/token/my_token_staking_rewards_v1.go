// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyTokenStakingRewardsV1MetaData contains all meta data concerning the MyTokenStakingRewardsV1 contract.
var MyTokenStakingRewardsV1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardRates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period_\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeIdToInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakedIds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"stakeId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidStake\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}]},{\"type\":\"error\",\"name\":\"StakeAlreadyClaimed\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StakeNotFound\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StakeNotMatured\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawRewardsFailed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawStakedFailed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawalInvalidStaker\",\"inputs\":[{\"name\":\"realStaker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60803461010c57601f6109b938819003918201601f19168301916001600160401b038311848410176101105780849260209460405283398101031261010c57516001600160a01b0381169081900361010c575f80546001600160a01b031916919091178155600260205260057fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b55600a7fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e055600f7f679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c556003905260147f88601476d11616a71c5be67555bd1dff4b1cbf21533d2669b768b61518cfe1c35560405161089490816101258239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c90816310087fb1146104b75750806310c09bbd1461041357806318160ddd146103f65780631d62ebd9146103be5780632e1a7d4d1461016b578063541cb57f1461013a57806370a082311461010257806396bfdcfe146100aa5763fc0c546a1461007f575f80fd5b346100a6575f3660031901126100a6575f546040516001600160a01b039091168152602090f35b5f80fd5b346100a65760403660031901126100a6576100c3610759565b6001600160a01b03165f908152600360205260409020805460243591908210156100a6576020916100f391610786565b90549060031b1c604051908152f35b346100a65760203660031901126100a6576001600160a01b03610123610759565b165f526006602052602060405f2054604051908152f35b346100a65760203660031901126100a65760043560048110156100a65761016260209161076f565b54604051908152f35b346100a65760203660031901126100a657600435805f52600460205260405f2060018060a01b036001820154169081156103ab573382036103945760078101805460ff8116610381576005830154804210610367575060019060ff19161790555f6020610258600660028501946101e586546005546107f6565b60055585548786528285526101ff604087209182546107f6565b90550180548685526007845261021a604086209182546107f6565b90558354905460405163a9059cbb60e01b81526001600160a01b03888116600483015260248201929092529485939190921691839182906044820190565b03925af1908115610327575f91610348575b5015610332575f8054915460405163a9059cbb60e01b81526001600160a01b0385811660048301526024820192909252926020928492604492849291165af1908115610327575f916102f8575b50156102e3577f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d55f80a3005b6309d6b46560e21b5f5260045260245260445ffd5b61031a915060203d602011610320575b61031281836107af565b8101906107d1565b836102b7565b503d610308565b6040513d5f823e3d90fd5b5063ab33ef4360e01b5f5260045260245260445ffd5b610361915060203d6020116103205761031281836107af565b8461026a565b85639f04f29f60e01b5f526004524260245260445260645ffd5b84633497c6d760e11b5f5260045260245ffd5b50631bb06b9d60e11b5f526004523360245260445ffd5b826351651db960e01b5f5260045260245ffd5b346100a65760203660031901126100a6576001600160a01b036103df610759565b165f526007602052602060405f2054604051908152f35b346100a6575f3660031901126100a6576020600554604051908152f35b346100a65760203660031901126100a6576004355f52600460205260405f20805460018060a01b0360018301541660028301549260ff60038201541693600482015460058301549160ff6007600686015495015416946040519687526020870152604086015260048610156104a357610100956060860152608085015260a084015260c0830152151560e0820152f35b634e487b7160e01b5f52602160045260245ffd5b346100a65760403660031901126100a657600435906024359060048210156100a6578215610746575f80546323b872dd60e01b835233600484015230602484015260448301859052602091839160649183916001600160a01b03165af1908115610327575f91610727575b50156107105761053181610803565b9161053b8261076f565b5480820290828204036106fc5760649004916001545f1981146106fc57600161056a91019485600155426107e9565b90604051610100810181811067ffffffffffffffff8211176106e85760405285815260208101923384526040820190858252606083019384526080830142815260a0840191825260c084019288845260e08501965f88528a5f52600460205260405f2095518655600186019060018060a01b039051166bffffffffffffffffffffffff60a01b8254161790555160028501556003840194519460048610156104a35760079560ff80198354169116179055516004840155516005830155516006820155019051151560ff80198354169116179055335f52600360205260405f208054680100000000000000008110156106e85761066c91600182018155610786565b81549060031b9085821b915f19901b191617905561068c816005546107e9565b600555335f5260066020526106a660405f209182546107e9565b9055335f5260076020526106bf60405f209182546107e9565b9055337f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d5f80a3005b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5063eb15156960e01b5f523360045260245260445ffd5b610740915060203d6020116103205761031281836107af565b83610522565b8263222d164360e21b5f5260045260245ffd5b600435906001600160a01b03821682036100a657565b60048110156104a3575f52600260205260405f2090565b805482101561079b575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176106e857604052565b908160209103126100a6575180151581036100a65790565b919082018092116106fc57565b919082039182116106fc57565b6004811080156104a35781610819575050601e90565b60018203610828575050605a90565b6002820361083757505060b490565b6003820361084757505061016d90565b633a71614b60e21b5f52156104a35760045260245ffdfea2646970667358221220be2e4d00ab42287dbcfee4a0551cfe9ef2e58d7765e2535562f770933c575c1e64736f6c63430008230033",
}

// MyTokenStakingRewardsV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyTokenStakingRewardsV1MetaData.ABI instead.
var MyTokenStakingRewardsV1ABI = MyTokenStakingRewardsV1MetaData.ABI

// MyTokenStakingRewardsV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyTokenStakingRewardsV1MetaData.Bin instead.
var MyTokenStakingRewardsV1Bin = MyTokenStakingRewardsV1MetaData.Bin

// DeployMyTokenStakingRewardsV1 deploys a new Ethereum contract, binding an instance of MyTokenStakingRewardsV1 to it.
func DeployMyTokenStakingRewardsV1(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *MyTokenStakingRewardsV1, error) {
	parsed, err := MyTokenStakingRewardsV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyTokenStakingRewardsV1Bin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyTokenStakingRewardsV1{MyTokenStakingRewardsV1Caller: MyTokenStakingRewardsV1Caller{contract: contract}, MyTokenStakingRewardsV1Transactor: MyTokenStakingRewardsV1Transactor{contract: contract}, MyTokenStakingRewardsV1Filterer: MyTokenStakingRewardsV1Filterer{contract: contract}}, nil
}

// MyTokenStakingRewardsV1 is an auto generated Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1 struct {
	MyTokenStakingRewardsV1Caller     // Read-only binding to the contract
	MyTokenStakingRewardsV1Transactor // Write-only binding to the contract
	MyTokenStakingRewardsV1Filterer   // Log filterer for contract events
}

// MyTokenStakingRewardsV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenStakingRewardsV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenStakingRewardsV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyTokenStakingRewardsV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenStakingRewardsV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTokenStakingRewardsV1Session struct {
	Contract     *MyTokenStakingRewardsV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MyTokenStakingRewardsV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTokenStakingRewardsV1CallerSession struct {
	Contract *MyTokenStakingRewardsV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MyTokenStakingRewardsV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTokenStakingRewardsV1TransactorSession struct {
	Contract     *MyTokenStakingRewardsV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MyTokenStakingRewardsV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1Raw struct {
	Contract *MyTokenStakingRewardsV1 // Generic contract binding to access the raw methods on
}

// MyTokenStakingRewardsV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1CallerRaw struct {
	Contract *MyTokenStakingRewardsV1Caller // Generic read-only contract binding to access the raw methods on
}

// MyTokenStakingRewardsV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1TransactorRaw struct {
	Contract *MyTokenStakingRewardsV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyTokenStakingRewardsV1 creates a new instance of MyTokenStakingRewardsV1, bound to a specific deployed contract.
func NewMyTokenStakingRewardsV1(address common.Address, backend bind.ContractBackend) (*MyTokenStakingRewardsV1, error) {
	contract, err := bindMyTokenStakingRewardsV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1{MyTokenStakingRewardsV1Caller: MyTokenStakingRewardsV1Caller{contract: contract}, MyTokenStakingRewardsV1Transactor: MyTokenStakingRewardsV1Transactor{contract: contract}, MyTokenStakingRewardsV1Filterer: MyTokenStakingRewardsV1Filterer{contract: contract}}, nil
}

// NewMyTokenStakingRewardsV1Caller creates a new read-only instance of MyTokenStakingRewardsV1, bound to a specific deployed contract.
func NewMyTokenStakingRewardsV1Caller(address common.Address, caller bind.ContractCaller) (*MyTokenStakingRewardsV1Caller, error) {
	contract, err := bindMyTokenStakingRewardsV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1Caller{contract: contract}, nil
}

// NewMyTokenStakingRewardsV1Transactor creates a new write-only instance of MyTokenStakingRewardsV1, bound to a specific deployed contract.
func NewMyTokenStakingRewardsV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MyTokenStakingRewardsV1Transactor, error) {
	contract, err := bindMyTokenStakingRewardsV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1Transactor{contract: contract}, nil
}

// NewMyTokenStakingRewardsV1Filterer creates a new log filterer instance of MyTokenStakingRewardsV1, bound to a specific deployed contract.
func NewMyTokenStakingRewardsV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MyTokenStakingRewardsV1Filterer, error) {
	contract, err := bindMyTokenStakingRewardsV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1Filterer{contract: contract}, nil
}

// bindMyTokenStakingRewardsV1 binds a generic wrapper to an already deployed contract.
func bindMyTokenStakingRewardsV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyTokenStakingRewardsV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyTokenStakingRewardsV1.Contract.MyTokenStakingRewardsV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.MyTokenStakingRewardsV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.MyTokenStakingRewardsV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyTokenStakingRewardsV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.BalanceOf(&_MyTokenStakingRewardsV1.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.BalanceOf(&_MyTokenStakingRewardsV1.CallOpts, account)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) RewardOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "rewardOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) RewardOf(account common.Address) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.RewardOf(&_MyTokenStakingRewardsV1.CallOpts, account)
}

// RewardOf is a free data retrieval call binding the contract method 0x1d62ebd9.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) RewardOf(account common.Address) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.RewardOf(&_MyTokenStakingRewardsV1.CallOpts, account)
}

// RewardRates is a free data retrieval call binding the contract method 0x541cb57f.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) RewardRates(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "rewardRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRates is a free data retrieval call binding the contract method 0x541cb57f.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) RewardRates(arg0 uint8) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.RewardRates(&_MyTokenStakingRewardsV1.CallOpts, arg0)
}

// RewardRates is a free data retrieval call binding the contract method 0x541cb57f.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) RewardRates(arg0 uint8) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.RewardRates(&_MyTokenStakingRewardsV1.CallOpts, arg0)
}

// StakeIdToInfo is a free data retrieval call binding the contract method 0x10c09bbd.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) StakeIdToInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	Staker        common.Address
	StakedAmount  *big.Int
	Period        uint8
	StartTime     *big.Int
	EndTime       *big.Int
	RewardsAmount *big.Int
	Claimed       bool
}, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "stakeIdToInfo", arg0)

	outstruct := new(struct {
		Id            *big.Int
		Staker        common.Address
		StakedAmount  *big.Int
		Period        uint8
		StartTime     *big.Int
		EndTime       *big.Int
		RewardsAmount *big.Int
		Claimed       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakedAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Period = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardsAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// StakeIdToInfo is a free data retrieval call binding the contract method 0x10c09bbd.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) StakeIdToInfo(arg0 *big.Int) (struct {
	Id            *big.Int
	Staker        common.Address
	StakedAmount  *big.Int
	Period        uint8
	StartTime     *big.Int
	EndTime       *big.Int
	RewardsAmount *big.Int
	Claimed       bool
}, error) {
	return _MyTokenStakingRewardsV1.Contract.StakeIdToInfo(&_MyTokenStakingRewardsV1.CallOpts, arg0)
}

// StakeIdToInfo is a free data retrieval call binding the contract method 0x10c09bbd.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) StakeIdToInfo(arg0 *big.Int) (struct {
	Id            *big.Int
	Staker        common.Address
	StakedAmount  *big.Int
	Period        uint8
	StartTime     *big.Int
	EndTime       *big.Int
	RewardsAmount *big.Int
	Claimed       bool
}, error) {
	return _MyTokenStakingRewardsV1.Contract.StakeIdToInfo(&_MyTokenStakingRewardsV1.CallOpts, arg0)
}

// StakedIds is a free data retrieval call binding the contract method 0x96bfdcfe.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) StakedIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "stakedIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedIds is a free data retrieval call binding the contract method 0x96bfdcfe.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) StakedIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.StakedIds(&_MyTokenStakingRewardsV1.CallOpts, arg0, arg1)
}

// StakedIds is a free data retrieval call binding the contract method 0x96bfdcfe.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) StakedIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.StakedIds(&_MyTokenStakingRewardsV1.CallOpts, arg0, arg1)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) Token() (common.Address, error) {
	return _MyTokenStakingRewardsV1.Contract.Token(&_MyTokenStakingRewardsV1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) Token() (common.Address, error) {
	return _MyTokenStakingRewardsV1.Contract.Token(&_MyTokenStakingRewardsV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyTokenStakingRewardsV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) TotalSupply() (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.TotalSupply(&_MyTokenStakingRewardsV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1CallerSession) TotalSupply() (*big.Int, error) {
	return _MyTokenStakingRewardsV1.Contract.TotalSupply(&_MyTokenStakingRewardsV1.CallOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Transactor) Stake(opts *bind.TransactOpts, amount *big.Int, period_ uint8) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.contract.Transact(opts, "stake", amount, period_)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) Stake(amount *big.Int, period_ uint8) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.Stake(&_MyTokenStakingRewardsV1.TransactOpts, amount, period_)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1TransactorSession) Stake(amount *big.Int, period_ uint8) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.Stake(&_MyTokenStakingRewardsV1.TransactOpts, amount, period_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Transactor) Withdraw(opts *bind.TransactOpts, stakeId_ *big.Int) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.contract.Transact(opts, "withdraw", stakeId_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Session) Withdraw(stakeId_ *big.Int) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.Withdraw(&_MyTokenStakingRewardsV1.TransactOpts, stakeId_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId_) returns()
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1TransactorSession) Withdraw(stakeId_ *big.Int) (*types.Transaction, error) {
	return _MyTokenStakingRewardsV1.Contract.Withdraw(&_MyTokenStakingRewardsV1.TransactOpts, stakeId_)
}

// MyTokenStakingRewardsV1StakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1StakedIterator struct {
	Event *MyTokenStakingRewardsV1Staked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyTokenStakingRewardsV1StakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenStakingRewardsV1Staked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenStakingRewardsV1Staked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyTokenStakingRewardsV1StakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenStakingRewardsV1StakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenStakingRewardsV1Staked represents a Staked event raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1Staked struct {
	Staker  common.Address
	StakeId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address, stakeId []*big.Int) (*MyTokenStakingRewardsV1StakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _MyTokenStakingRewardsV1.contract.FilterLogs(opts, "Staked", stakerRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1StakedIterator{contract: _MyTokenStakingRewardsV1.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *MyTokenStakingRewardsV1Staked, staker []common.Address, stakeId []*big.Int) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _MyTokenStakingRewardsV1.contract.WatchLogs(opts, "Staked", stakerRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenStakingRewardsV1Staked)
				if err := _MyTokenStakingRewardsV1.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) ParseStaked(log types.Log) (*MyTokenStakingRewardsV1Staked, error) {
	event := new(MyTokenStakingRewardsV1Staked)
	if err := _MyTokenStakingRewardsV1.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyTokenStakingRewardsV1WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1WithdrawnIterator struct {
	Event *MyTokenStakingRewardsV1Withdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyTokenStakingRewardsV1WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenStakingRewardsV1Withdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenStakingRewardsV1Withdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyTokenStakingRewardsV1WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenStakingRewardsV1WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenStakingRewardsV1Withdrawn represents a Withdrawn event raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1Withdrawn struct {
	Staker  common.Address
	StakeId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) FilterWithdrawn(opts *bind.FilterOpts, staker []common.Address, stakeId []*big.Int) (*MyTokenStakingRewardsV1WithdrawnIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _MyTokenStakingRewardsV1.contract.FilterLogs(opts, "Withdrawn", stakerRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenStakingRewardsV1WithdrawnIterator{contract: _MyTokenStakingRewardsV1.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *MyTokenStakingRewardsV1Withdrawn, staker []common.Address, stakeId []*big.Int) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _MyTokenStakingRewardsV1.contract.WatchLogs(opts, "Withdrawn", stakerRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenStakingRewardsV1Withdrawn)
				if err := _MyTokenStakingRewardsV1.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed staker, uint256 indexed stakeId)
func (_MyTokenStakingRewardsV1 *MyTokenStakingRewardsV1Filterer) ParseWithdrawn(log types.Log) (*MyTokenStakingRewardsV1Withdrawn, error) {
	event := new(MyTokenStakingRewardsV1Withdrawn)
	if err := _MyTokenStakingRewardsV1.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
