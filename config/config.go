package config

type Config struct {
	Proxies  []Proxy  `mapstructure:"proxies"`
	Log      string   `mapstructure:"log"`
	CertFile CertFile `mapstructure:"tls"`
}
