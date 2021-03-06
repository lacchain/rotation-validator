package model

type ApplicationConfig struct {
	WSURL string `mapstructure:"wsURL"`
	RPCURL string `mapstructure:"rpcURL"`
	ContractAddress string `mapstructure:"contractAddress"`
	NodeKeyPath string `mapstructure:"nodeKeyPath"`
	NodeAddressPath string `mapstructure:"nodeAddressPath"`
	NodeAddress string `mapstructure:"nodeAddress"`
	Key string `mapstructure:"key"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
}