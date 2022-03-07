package main

import (
	"crypto/rand"
	"crypto/tls"
	"github.com/coreos/go-systemd/daemon"
	"github.com/wytfy/gredir/config"
	"github.com/wytfy/gredir/global"
	"github.com/wytfy/gredir/handler"
	"github.com/wytfy/gredir/initialize"
	"log"
	"sync"
	"time"
)

func main() {
	// 读取配置文件
	initialize.InitViper()
	initialize.InitLogrus(global.CONF.Log)
	// 守护进程
	if _, err := daemon.SdNotify(false, "READY=1"); err != nil {
		global.LOGGER.Warn("notification supported, but failed")
	}
	// tls配置
	cert := global.CONF.CertFile
	crt, err := tls.LoadX509KeyPair(cert.PublicKey, cert.PrivateKey)
	if err != nil {
		global.LOGGER.Fatalln(err)
	}
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	// 对每一个配置启动一个代理
	proxies := global.CONF.Proxies
	var wg sync.WaitGroup
	wg.Add(len(proxies))
	for _, proxy := range proxies {
		go func(p config.Proxy) {
			defer wg.Done()
			// listen, err := net.Listen("tcp", p.LocalAddr)
			listen, err := tls.Listen("tcp", p.LocalAddr, tlsConfig)
			if err != nil {
				global.LOGGER.Errorf("[%v] Fail to listen the port: %v", p.User, err)
			}
			global.LOGGER.Infof("[%s] listen on localhost: %v.", p.User, p.LocalAddr)
			for {
				conn, err := listen.Accept()
				if err != nil {
					global.LOGGER.Errorf("[%s] connection failed", p.User)
					continue
				}
				go handler.TcpHandler(conn, &p)
			}
		}(proxy)
	}
	wg.Wait()
	log.Printf("exit.")
}
