// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// MyTokenStakingRewardsV1MetaData contains all meta data concerning the MyTokenStakingRewardsV1 contract.
var MyTokenStakingRewardsV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardRates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period_\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeIdToInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakedIds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"stakeId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidStake\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMyTokenStakingRewardsV1.StakingPeriod\"}]},{\"type\":\"error\",\"name\":\"StakeAlreadyClaimed\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StakeNotFound\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StakeNotMatured\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawRewardsFailed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawStakedFailed\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawalInvalidStaker\",\"inputs\":[{\"name\":\"realStaker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	ID:  "MyTokenStakingRewardsV1",
	Bin: "0x60803461010c57601f6109b938819003918201601f19168301916001600160401b038311848410176101105780849260209460405283398101031261010c57516001600160a01b0381169081900361010c575f80546001600160a01b031916919091178155600260205260057fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b55600a7fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e055600f7f679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c556003905260147f88601476d11616a71c5be67555bd1dff4b1cbf21533d2669b768b61518cfe1c35560405161089490816101258239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c90816310087fb1146104b75750806310c09bbd1461041357806318160ddd146103f65780631d62ebd9146103be5780632e1a7d4d1461016b578063541cb57f1461013a57806370a082311461010257806396bfdcfe146100aa5763fc0c546a1461007f575f80fd5b346100a6575f3660031901126100a6575f546040516001600160a01b039091168152602090f35b5f80fd5b346100a65760403660031901126100a6576100c3610759565b6001600160a01b03165f908152600360205260409020805460243591908210156100a6576020916100f391610786565b90549060031b1c604051908152f35b346100a65760203660031901126100a6576001600160a01b03610123610759565b165f526006602052602060405f2054604051908152f35b346100a65760203660031901126100a65760043560048110156100a65761016260209161076f565b54604051908152f35b346100a65760203660031901126100a657600435805f52600460205260405f2060018060a01b036001820154169081156103ab573382036103945760078101805460ff8116610381576005830154804210610367575060019060ff19161790555f6020610258600660028501946101e586546005546107f6565b60055585548786528285526101ff604087209182546107f6565b90550180548685526007845261021a604086209182546107f6565b90558354905460405163a9059cbb60e01b81526001600160a01b03888116600483015260248201929092529485939190921691839182906044820190565b03925af1908115610327575f91610348575b5015610332575f8054915460405163a9059cbb60e01b81526001600160a01b0385811660048301526024820192909252926020928492604492849291165af1908115610327575f916102f8575b50156102e3577f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d55f80a3005b6309d6b46560e21b5f5260045260245260445ffd5b61031a915060203d602011610320575b61031281836107af565b8101906107d1565b836102b7565b503d610308565b6040513d5f823e3d90fd5b5063ab33ef4360e01b5f5260045260245260445ffd5b610361915060203d6020116103205761031281836107af565b8461026a565b85639f04f29f60e01b5f526004524260245260445260645ffd5b84633497c6d760e11b5f5260045260245ffd5b50631bb06b9d60e11b5f526004523360245260445ffd5b826351651db960e01b5f5260045260245ffd5b346100a65760203660031901126100a6576001600160a01b036103df610759565b165f526007602052602060405f2054604051908152f35b346100a6575f3660031901126100a6576020600554604051908152f35b346100a65760203660031901126100a6576004355f52600460205260405f20805460018060a01b0360018301541660028301549260ff60038201541693600482015460058301549160ff6007600686015495015416946040519687526020870152604086015260048610156104a357610100956060860152608085015260a084015260c0830152151560e0820152f35b634e487b7160e01b5f52602160045260245ffd5b346100a65760403660031901126100a657600435906024359060048210156100a6578215610746575f80546323b872dd60e01b835233600484015230602484015260448301859052602091839160649183916001600160a01b03165af1908115610327575f91610727575b50156107105761053181610803565b9161053b8261076f565b5480820290828204036106fc5760649004916001545f1981146106fc57600161056a91019485600155426107e9565b90604051610100810181811067ffffffffffffffff8211176106e85760405285815260208101923384526040820190858252606083019384526080830142815260a0840191825260c084019288845260e08501965f88528a5f52600460205260405f2095518655600186019060018060a01b039051166bffffffffffffffffffffffff60a01b8254161790555160028501556003840194519460048610156104a35760079560ff80198354169116179055516004840155516005830155516006820155019051151560ff80198354169116179055335f52600360205260405f208054680100000000000000008110156106e85761066c91600182018155610786565b81549060031b9085821b915f19901b191617905561068c816005546107e9565b600555335f5260066020526106a660405f209182546107e9565b9055335f5260076020526106bf60405f209182546107e9565b9055337f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d5f80a3005b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5063eb15156960e01b5f523360045260245260445ffd5b610740915060203d6020116103205761031281836107af565b83610522565b8263222d164360e21b5f5260045260245ffd5b600435906001600160a01b03821682036100a657565b60048110156104a3575f52600260205260405f2090565b805482101561079b575f5260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176106e857604052565b908160209103126100a6575180151581036100a65790565b919082018092116106fc57565b919082039182116106fc57565b6004811080156104a35781610819575050601e90565b60018203610828575050605a90565b6002820361083757505060b490565b6003820361084757505061016d90565b633a71614b60e21b5f52156104a35760045260245ffdfea2646970667358221220be2e4d00ab42287dbcfee4a0551cfe9ef2e58d7765e2535562f770933c575c1e64736f6c63430008230033",
}

// MyTokenStakingRewardsV1 is an auto generated Go binding around an Ethereum contract.
type MyTokenStakingRewardsV1 struct {
	abi abi.ABI
}

// NewMyTokenStakingRewardsV1 creates a new instance of MyTokenStakingRewardsV1.
func NewMyTokenStakingRewardsV1() *MyTokenStakingRewardsV1 {
	parsed, err := MyTokenStakingRewardsV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MyTokenStakingRewardsV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MyTokenStakingRewardsV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address token_) returns()
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackConstructor(token_ common.Address) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("", token_)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackBalanceOf(account common.Address) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRewardOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1d62ebd9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackRewardOf(account common.Address) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("rewardOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1d62ebd9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackRewardOf(account common.Address) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("rewardOf", account)
}

// UnpackRewardOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1d62ebd9.
//
// Solidity: function rewardOf(address account) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackRewardOf(data []byte) (*big.Int, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("rewardOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRewardRates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x541cb57f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackRewardRates(arg0 uint8) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("rewardRates", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardRates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x541cb57f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackRewardRates(arg0 uint8) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("rewardRates", arg0)
}

// UnpackRewardRates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x541cb57f.
//
// Solidity: function rewardRates(uint8 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackRewardRates(data []byte) (*big.Int, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("rewardRates", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10087fb1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stake(uint256 amount, uint8 period_) returns()
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackStake(amount *big.Int, period uint8) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("stake", amount, period)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10087fb1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stake(uint256 amount, uint8 period_) returns()
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackStake(amount *big.Int, period uint8) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("stake", amount, period)
}

// PackStakeIdToInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10c09bbd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackStakeIdToInfo(arg0 *big.Int) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("stakeIdToInfo", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStakeIdToInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10c09bbd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackStakeIdToInfo(arg0 *big.Int) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("stakeIdToInfo", arg0)
}

// StakeIdToInfoOutput serves as a container for the return parameters of contract
// method StakeIdToInfo.
type StakeIdToInfoOutput struct {
	Id            *big.Int
	Staker        common.Address
	StakedAmount  *big.Int
	Period        uint8
	StartTime     *big.Int
	EndTime       *big.Int
	RewardsAmount *big.Int
	Claimed       bool
}

// UnpackStakeIdToInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x10c09bbd.
//
// Solidity: function stakeIdToInfo(uint256 ) view returns(uint256 id, address staker, uint256 stakedAmount, uint8 period, uint256 startTime, uint256 endTime, uint256 rewardsAmount, bool claimed)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakeIdToInfo(data []byte) (StakeIdToInfoOutput, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("stakeIdToInfo", data)
	outstruct := new(StakeIdToInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Id = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Staker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakedAmount = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.Period = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.StartTime = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.EndTime = abi.ConvertType(out[5], new(big.Int)).(*big.Int)
	outstruct.RewardsAmount = abi.ConvertType(out[6], new(big.Int)).(*big.Int)
	outstruct.Claimed = *abi.ConvertType(out[7], new(bool)).(*bool)
	return *outstruct, nil
}

// PackStakedIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x96bfdcfe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackStakedIds(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("stakedIds", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStakedIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x96bfdcfe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackStakedIds(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("stakedIds", arg0, arg1)
}

// UnpackStakedIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x96bfdcfe.
//
// Solidity: function stakedIds(address , uint256 ) view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakedIds(data []byte) (*big.Int, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("stakedIds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function token() view returns(address)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackToken() []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("token")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function token() view returns(address)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackToken() ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("token")
}

// UnpackToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackToken(data []byte) (common.Address, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("token", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackTotalSupply() []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackTotalSupply() ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := myTokenStakingRewardsV1.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e1a7d4d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdraw(uint256 stakeId_) returns()
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) PackWithdraw(stakeId *big.Int) []byte {
	enc, err := myTokenStakingRewardsV1.abi.Pack("withdraw", stakeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e1a7d4d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdraw(uint256 stakeId_) returns()
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) TryPackWithdraw(stakeId *big.Int) ([]byte, error) {
	return myTokenStakingRewardsV1.abi.Pack("withdraw", stakeId)
}

// MyTokenStakingRewardsV1Staked represents a Staked event raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1Staked struct {
	Staker  common.Address
	StakeId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const MyTokenStakingRewardsV1StakedEventName = "Staked"

// ContractEventName returns the user-defined event name.
func (MyTokenStakingRewardsV1Staked) ContractEventName() string {
	return MyTokenStakingRewardsV1StakedEventName
}

// UnpackStakedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Staked(address indexed staker, uint256 indexed stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakedEvent(log *types.Log) (*MyTokenStakingRewardsV1Staked, error) {
	event := "Staked"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != myTokenStakingRewardsV1.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(MyTokenStakingRewardsV1Staked)
	if len(log.Data) > 0 {
		if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myTokenStakingRewardsV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MyTokenStakingRewardsV1Withdrawn represents a Withdrawn event raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1Withdrawn struct {
	Staker  common.Address
	StakeId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const MyTokenStakingRewardsV1WithdrawnEventName = "Withdrawn"

// ContractEventName returns the user-defined event name.
func (MyTokenStakingRewardsV1Withdrawn) ContractEventName() string {
	return MyTokenStakingRewardsV1WithdrawnEventName
}

// UnpackWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdrawn(address indexed staker, uint256 indexed stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackWithdrawnEvent(log *types.Log) (*MyTokenStakingRewardsV1Withdrawn, error) {
	event := "Withdrawn"
	if len(log.Topics) == 0 {
		return nil, bind.ErrNoEventSignature
	}
	if log.Topics[0] != myTokenStakingRewardsV1.abi.Events[event].ID {
		return nil, bind.ErrEventSignatureMismatch
	}
	out := new(MyTokenStakingRewardsV1Withdrawn)
	if len(log.Data) > 0 {
		if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myTokenStakingRewardsV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["InvalidStake"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackInvalidStakeError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["InvalidStakeAmount"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackInvalidStakeAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["InvalidStakingPeriod"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackInvalidStakingPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["StakeAlreadyClaimed"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackStakeAlreadyClaimedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["StakeNotFound"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackStakeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["StakeNotMatured"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackStakeNotMaturedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["WithdrawRewardsFailed"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackWithdrawRewardsFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["WithdrawStakedFailed"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackWithdrawStakedFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myTokenStakingRewardsV1.abi.Errors["WithdrawalInvalidStaker"].ID.Bytes()[:4]) {
		return myTokenStakingRewardsV1.UnpackWithdrawalInvalidStakerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MyTokenStakingRewardsV1InvalidStake represents a InvalidStake error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1InvalidStake struct {
	Staker      common.Address
	StakeAmount *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidStake(address staker, uint256 stakeAmount)
func MyTokenStakingRewardsV1InvalidStakeErrorID() common.Hash {
	return common.HexToHash("0xeb151569aee518edd58f462f1950978c1594ebfcd2688680103a68913a3661e7")
}

// UnpackInvalidStakeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidStake(address staker, uint256 stakeAmount)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackInvalidStakeError(raw []byte) (*MyTokenStakingRewardsV1InvalidStake, error) {
	out := new(MyTokenStakingRewardsV1InvalidStake)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "InvalidStake", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1InvalidStakeAmount represents a InvalidStakeAmount error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1InvalidStakeAmount struct {
	Amount *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidStakeAmount(uint256 amount)
func MyTokenStakingRewardsV1InvalidStakeAmountErrorID() common.Hash {
	return common.HexToHash("0x88b4590c99d4b7579c7ea1cc9a6d43ad19bffb1797e0f8c407de7af9e9339ce4")
}

// UnpackInvalidStakeAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidStakeAmount(uint256 amount)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackInvalidStakeAmountError(raw []byte) (*MyTokenStakingRewardsV1InvalidStakeAmount, error) {
	out := new(MyTokenStakingRewardsV1InvalidStakeAmount)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "InvalidStakeAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1InvalidStakingPeriod represents a InvalidStakingPeriod error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1InvalidStakingPeriod struct {
	Period uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidStakingPeriod(uint8 period)
func MyTokenStakingRewardsV1InvalidStakingPeriodErrorID() common.Hash {
	return common.HexToHash("0xe9c5852c758fbf9e2dc6f4cc2693a873dc41a48d458d15757ced00f3a34a14a5")
}

// UnpackInvalidStakingPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidStakingPeriod(uint8 period)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackInvalidStakingPeriodError(raw []byte) (*MyTokenStakingRewardsV1InvalidStakingPeriod, error) {
	out := new(MyTokenStakingRewardsV1InvalidStakingPeriod)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "InvalidStakingPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1StakeAlreadyClaimed represents a StakeAlreadyClaimed error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1StakeAlreadyClaimed struct {
	StakeId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StakeAlreadyClaimed(uint256 stakeId)
func MyTokenStakingRewardsV1StakeAlreadyClaimedErrorID() common.Hash {
	return common.HexToHash("0x692f8dae3df92885ba7078c6781cf0be038620cc9e3f27a0d318313b29df0be4")
}

// UnpackStakeAlreadyClaimedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StakeAlreadyClaimed(uint256 stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakeAlreadyClaimedError(raw []byte) (*MyTokenStakingRewardsV1StakeAlreadyClaimed, error) {
	out := new(MyTokenStakingRewardsV1StakeAlreadyClaimed)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "StakeAlreadyClaimed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1StakeNotFound represents a StakeNotFound error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1StakeNotFound struct {
	StakeId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StakeNotFound(uint256 stakeId)
func MyTokenStakingRewardsV1StakeNotFoundErrorID() common.Hash {
	return common.HexToHash("0x51651db92e27a2580851d046203ca08014885282f5c5eb809aac842583e97ca6")
}

// UnpackStakeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StakeNotFound(uint256 stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakeNotFoundError(raw []byte) (*MyTokenStakingRewardsV1StakeNotFound, error) {
	out := new(MyTokenStakingRewardsV1StakeNotFound)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "StakeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1StakeNotMatured represents a StakeNotMatured error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1StakeNotMatured struct {
	StakeId     *big.Int
	CurrentTime *big.Int
	EndTime     *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StakeNotMatured(uint256 stakeId, uint256 currentTime, uint256 endTime)
func MyTokenStakingRewardsV1StakeNotMaturedErrorID() common.Hash {
	return common.HexToHash("0x9f04f29f58a55c8e0d66dbd210114e61766e6585dcf099bcf974cafd2a30f152")
}

// UnpackStakeNotMaturedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StakeNotMatured(uint256 stakeId, uint256 currentTime, uint256 endTime)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackStakeNotMaturedError(raw []byte) (*MyTokenStakingRewardsV1StakeNotMatured, error) {
	out := new(MyTokenStakingRewardsV1StakeNotMatured)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "StakeNotMatured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1WithdrawRewardsFailed represents a WithdrawRewardsFailed error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1WithdrawRewardsFailed struct {
	Staker  common.Address
	StakeId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WithdrawRewardsFailed(address staker, uint256 stakeId)
func MyTokenStakingRewardsV1WithdrawRewardsFailedErrorID() common.Hash {
	return common.HexToHash("0xab33ef435b52ec31b36264d48ddb10d0b2bac8b7e10c2449a7f795afe738b620")
}

// UnpackWithdrawRewardsFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WithdrawRewardsFailed(address staker, uint256 stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackWithdrawRewardsFailedError(raw []byte) (*MyTokenStakingRewardsV1WithdrawRewardsFailed, error) {
	out := new(MyTokenStakingRewardsV1WithdrawRewardsFailed)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "WithdrawRewardsFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1WithdrawStakedFailed represents a WithdrawStakedFailed error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1WithdrawStakedFailed struct {
	Staker  common.Address
	StakeId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WithdrawStakedFailed(address staker, uint256 stakeId)
func MyTokenStakingRewardsV1WithdrawStakedFailedErrorID() common.Hash {
	return common.HexToHash("0x275ad194e27cc04445df3a88b2409ce76446bfe552f815e1cafe3cacd7aea7fc")
}

// UnpackWithdrawStakedFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WithdrawStakedFailed(address staker, uint256 stakeId)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackWithdrawStakedFailedError(raw []byte) (*MyTokenStakingRewardsV1WithdrawStakedFailed, error) {
	out := new(MyTokenStakingRewardsV1WithdrawStakedFailed)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "WithdrawStakedFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyTokenStakingRewardsV1WithdrawalInvalidStaker represents a WithdrawalInvalidStaker error raised by the MyTokenStakingRewardsV1 contract.
type MyTokenStakingRewardsV1WithdrawalInvalidStaker struct {
	RealStaker common.Address
	Sender     common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WithdrawalInvalidStaker(address realStaker, address sender)
func MyTokenStakingRewardsV1WithdrawalInvalidStakerErrorID() common.Hash {
	return common.HexToHash("0x3760d73ae9db80ba6fb2adc9d0937f14d82f00cfe1bc61acacfe096b7fb4d56e")
}

// UnpackWithdrawalInvalidStakerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WithdrawalInvalidStaker(address realStaker, address sender)
func (myTokenStakingRewardsV1 *MyTokenStakingRewardsV1) UnpackWithdrawalInvalidStakerError(raw []byte) (*MyTokenStakingRewardsV1WithdrawalInvalidStaker, error) {
	out := new(MyTokenStakingRewardsV1WithdrawalInvalidStaker)
	if err := myTokenStakingRewardsV1.abi.UnpackIntoInterface(out, "WithdrawalInvalidStaker", raw); err != nil {
		return nil, err
	}
	return out, nil
}
