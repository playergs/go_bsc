package config

type EthereumConfig struct {
	Network    string          `mapstructure:"network" json:"network"`
	RpcUrl     string          `mapstructure:"rpc_url" json:"rpc_url"`
	PrivateKey string          `mapstructure:"private_key" json:"private_key"`
	Contracts  ContractsConfig `mapstructure:"contracts" json:"contracts"`
}

type ContractsConfig struct {
	TokenAddress string `mapstructure:"token_address" json:"token_address"`
	StakeAddress string `mapstructure:"stake_address" json:"stake_address"`
}
