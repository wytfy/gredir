package config

type Proxy struct {
	Id         uint32 `yaml:"id"`
	LocalAddr  string `yaml:"localAddr"`
	RemoteAddr string `yaml:"remoteAddr"`
	User       string `yaml:"user"`
	Date       string `yaml:"date"`
}
