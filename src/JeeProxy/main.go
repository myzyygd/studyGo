package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"net"
	"JeeProxy/socks5"
)

type Config struct {
	Addr string
}

func main() {
	conf, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	runSOCKS5Server(*conf)
}

//配置文件加载
func LoadConfig(s string) (*Config, error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err = json.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}

//开启socks5服务
func runSOCKS5Server(conf Config) {
	log.Println("server run at " + conf.Addr)
	listener, err := net.Listen("tcp", conf.Addr)
	defer listener.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go socks5.NewSocks5Client(conn)
	}

}
