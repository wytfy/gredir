package config

type CertFile struct {
	PublicKey  string `mapstructure:"public_file"`
	PrivateKey string `mapstructure:"private_file"`
}
