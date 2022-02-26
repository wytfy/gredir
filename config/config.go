package config

type Config struct {
	Proxies []Proxy `yaml:"proxies"`
	Log     string  `yaml:"log"`
}
