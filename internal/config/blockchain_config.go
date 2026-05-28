package config

type BlockchainConfig struct {
	Ethereum EthereumConfig `mapstructure:"ethereum" json:"ethereum"`
}
