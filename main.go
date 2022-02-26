package main

import (
	"github.com/coreos/go-systemd/daemon"
	"github.com/wytfy/gredir/config"
	"github.com/wytfy/gredir/global"
	"github.com/wytfy/gredir/handler"
	"github.com/wytfy/gredir/initialize"
	"log"
	"net"
	"sync"
)

func main() {
	// 读取配置文件
	initialize.InitViper()
	initialize.InitLogrus(global.CONF.Log)
	// 守护进程
	if _, err := daemon.SdNotify(false, "READY=1"); err != nil {
		global.LOGGER.Warn("notification supported, but failed")
	}
	// 对每一个配置启动一个代理
	proxies := global.CONF.Proxies
	var wg sync.WaitGroup
	wg.Add(len(proxies))
	for _, proxy := range proxies {
		go func(p config.Proxy) {
			defer wg.Done()
			listen, err := net.Listen("tcp", p.LocalAddr)
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
