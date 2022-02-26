package handler

import (
	"github.com/wytfy/gredir/config"
	"github.com/wytfy/gredir/global"
	"io"
	"net"
	"sync"
)

func TcpHandler(localConn net.Conn, px *config.Proxy) {
	global.LOGGER.Infof("[%s] remote addr: %v", px.User, localConn.RemoteAddr())
	var wg sync.WaitGroup
	remoteConn, err := net.Dial("tcp", px.RemoteAddr)
	if err != nil {
		if err := localConn.Close(); err != nil {
			global.LOGGER.Errorf("[%v] failed to close local connection", px.User)
		}
		global.LOGGER.Errorf("[%s] failed to dial", px.User)
	}
	wg.Add(2)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		if _, err := io.Copy(remote, local); err != nil {
			global.LOGGER.Errorf("[%s] fail to sent the tcp message", px.User)
			if err := remote.Close(); err != nil {
				global.LOGGER.Errorf("[%v] failed to close remote connection", px.User)
			}
			return
		}
		global.LOGGER.Infof("[%s] sent the tcp message", px.User)
		if err := remote.Close(); err != nil {
			global.LOGGER.Errorf("[%v] failed to close remote connection", px.User)
		}
	}(localConn, remoteConn)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		if _, err := io.Copy(local, remote); err != nil {
			global.LOGGER.Errorf("[%s] fail to receive the tcp message", px.User)
			return
		}
		global.LOGGER.Infof("[%s] receive the tcp message", px.User)
	}(localConn, remoteConn)
	wg.Wait()
}
