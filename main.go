package main

import (
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	// 从错误中恢复
	defer func() {
		if p := recover(); p != nil {
			log.Fatalf("internal error: %v", p)
		}
	}()
	// 读取配置文件
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	if err :=config.ReadInConfig(); err != nil {
		log.Fatalf("Fail to read the configure file.")
	}
	proxies := &Proxies{}
	if err := config.Unmarshal(proxies); err != nil {
		log.Fatalf("fail to unmarshal config")
	}
	// 生成pidfile
	pidfile := "./gredir.pid"
	if _, err := os.Stat(pidfile); err != nil {
		os.Remove(pidfile)
	}
	ioutil.WriteFile(pidfile, []byte(strconv.Itoa(os.Getpid())), 0644)
	// logger配置
	logFile, err := os.OpenFile("./gredir.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("faile to open the log file.")
	}
	log.SetOutput(logFile)
	// 对每一个配置启动一个代理
	var wg sync.WaitGroup
	wg.Add(len(proxies.Proxies))
	for _, proxy := range proxies.Proxies {
		go func(p Proxy) {
			defer wg.Done()
			listen, err := net.Listen("tcp", p.LocalAddr)
			if err != nil {
				log.Println(err)
				log.Fatalf("[%s] Fail to listen the port", p.User)
			}
			log.Printf("[%s] listen on localhost:%v.", p.User, p.LocalAddr)
			for {
				conn, err := listen.Accept()
				if err != nil {
					log.Printf("[%s] connection failed.", p.User)
					continue
				}
				go handler(conn, &p)
			}
		}(proxy)
	}
	wg.Wait()
	log.Printf("exit.")
}

func handler(localConn net.Conn, px *Proxy) {
	var wg sync.WaitGroup
	remoteConn, err := net.Dial("tcp", px.RemoteAddr)
	if err != nil {
		localConn.Close()
		log.Printf("[%s] failed to dial", px.User)
	}
	wg.Add(2)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		if _,err := io.Copy(remote, local); err != nil {
			log.Printf("[%s] fail to sent the tcp message", px.User)
			remote.Close()
			return
		}
		log.Printf("[%s] sent the tcp message", px.User)
		remote.Close()
	}(localConn, remoteConn)
	go func(local net.Conn, remote net.Conn) {
		defer wg.Done()
		if _,err := io.Copy(local, remote); err != nil {
			log.Printf("[%s] fail to receive the tcp message", px.User)
			return
		}
		log.Printf("[%s] receive the tcp message", px.User)
	}(localConn, remoteConn)
	wg.Wait()
}

type Proxy struct {
	Id int32 `yaml:"id"`
	LocalAddr string `yaml:"localAddr"`
	RemoteAddr string `yaml:"remoteAddr"`
	User string `yaml:"user"`
}

type Proxies struct {
	Proxies []Proxy `yaml:"proxies"`
}