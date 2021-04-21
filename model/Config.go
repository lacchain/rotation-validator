package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	ContractAddress string `mapstructure:"contractAddress"`
	NodeKeyPath string `mapstructure:"nodeKeyPath"`
	NodeAddressPath string `mapstructure:"nodeAddressPath"`
	NodeAddress string `mapstructure:"nodeAddress"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
}