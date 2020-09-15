// Author: yangzq80@gmail.com
// Date: 2020-09-07
//
package server

import (
	"encoding/json"
	"fmt"
	"github.com/yusys-cloud/analog-network/conf"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type Server struct {
	ApiPort string
	//Proxies  []*conf.Proxy
	Proxies map[string]*ProxyServer
	conf    *conf.Conf
}
type ProxyServer struct {
	proxy *conf.Proxy
	rnd   *rand.Rand
}

func NewServer(cnf *conf.Conf) *Server {
	return &Server{
		ApiPort: cnf.ApiPort,
		Proxies: make(map[string]*ProxyServer),
		conf:    cnf,
	}
}

func (server *Server) Start() {
	server.startProxies()
	server.startApiServer()
}

func (server *Server) startProxies() {
	for _, cp := range server.conf.Proxies {
		log.Println("Starting proxy ", cp.Target, cp.Port)

		proxyServer := &ProxyServer{
			proxy: cp,
			rnd:   rand.New(rand.NewSource(time.Now().Unix())),
		}
		server.Proxies[cp.Target] = proxyServer

		go proxyServer.tcpListen()
	}
}

func (server *Server) startApiServer() {
	log.Println("API server refresh ---> curl localhost:" + server.ApiPort + "/apply")
	http.HandleFunc("/apply", func(writer http.ResponseWriter, request *http.Request) {
		//TODO 刷新proxies
		server.RefreshProxies()

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("success \n"))
	})
	log.Fatal(http.ListenAndServe(":"+server.ApiPort, nil))
}

func (proxyServer *ProxyServer) tcpListen() {
	ln, err := net.Listen("tcp", "0.0.0.0:"+proxyServer.proxy.Port)
	if err != nil {
		fmt.Println("tcp_listen:", err)
		return
	}
	defer ln.Close()
	for {
		localConn, err := ln.Accept() //接受tcp客户端连接，并返回新的套接字进行通信
		if err != nil {
			fmt.Println("Accept:", err)
			return
		}
		go proxyServer.tcpHandle(localConn)
	}
}

func (proxyServer *ProxyServer) tcpHandle(localConn net.Conn) {

	remoteConn, err := net.Dial("tcp", proxyServer.proxy.Target) //连接目标服务器

	if err != nil {
		log.Fatal(err)
		return
	}
	go proxyServer.CtrlCopyBuffer(remoteConn, localConn, proxyServer.proxy.Ctl.In, nil)
	go proxyServer.CtrlCopyBuffer(localConn, remoteConn, proxyServer.proxy.Ctl.In, nil)
}

func ReadConfig() *conf.Conf {
	file := "config.json"

	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("read config file <%s> failure. err:%+v", file, err)
	}

	cnf := &conf.Conf{}
	err = json.Unmarshal(data, cnf)
	if err != nil {
		log.Fatalf("parse config file <%s> failure. error:%+v", file, err)
	}

	return cnf
}

func (server *Server) RefreshProxies() error {

	conf := ReadConfig()

	for _, cp := range conf.Proxies {
		server.Proxies[cp.Target].proxy.Ctl.CopyFrom(cp.Ctl)
	}

	b, _ := json.MarshalIndent(conf, "", "   ")
	//b,_ := json.Marshal(conf)

	log.Println("New proxies ------> \n", string(b))

	return nil
}
