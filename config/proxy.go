package config

type Proxy struct {
	Id         uint32 `mapstructure:"id"`
	LocalAddr  string `mapstructure:"local_addr"`
	RemoteAddr string `mapstructure:"remote_addr"`
	User       string `mapstructure:"user"`
	Date       string `mapstructure:"date"`
	Tls        bool   `mapstructure:"tls"`
}
