package main

import (
	"net"
	"log"
	"io"
	"errors"
	"strconv"
)

func handelConnection(conn *net.TCPConn) error {
	buff := make([]byte, 1024)
	replay := []byte{0x05, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x22, 0x22}
	_, err := io.ReadFull(conn, buff[:2])
	if err != nil {
		return err
	}
	if buff[0] != 0x05 {
		replay[1] = 0xff
		conn.Write(replay[:2])
		return errors.New("协议错误")
	}
	methodNum := buff[1]
	_, err = io.ReadFull(conn, buff[:methodNum])
	if err != nil {
		return err
	}
	replay[1] = 0x00
	_, err = conn.Write(replay[:2])
	if err != nil {
		return err
	}
	_, err = io.ReadFull(conn, buff[:4])
	if err != nil {
		return err
	}

	if buff[1] != 1 {
		replay[1] = 7
		return errors.New("服务端不支持此命令")
	}
	addType := buff[3]
	log.Println(addType)
	addressLen  := 0
	switch addType {
	case 1:
		addressLen = net.IPv4len
	case 3:
		_,err := io.ReadFull(conn,buff[:1])
		if err!=nil{
			return errors.New("地址解析错误")
		}
		addressLen = int(buff[0])
	case 4:
		addressLen = net.IPv6len
	default:
		return errors.New("地址解析错误")
	}
	log.Println(addressLen)
	host := make([]byte,addressLen)
	_,err =io.ReadFull(conn,host)
	if err!=nil{
		return err
	}
	_,err = io.ReadFull(conn,buff[:2])
	if err!=nil{
		return err
	}
	hostStr := ""
	switch addType {
	case 1,4:
		ip := net.IP(host)
		hostStr = ip.String()
	case 3:
		hostStr = string(host)
	}
	log.Println(uint16(buff[0])<<8)
	log.Println("port")
	log.Println(uint16(buff[1]))
	log.Println("port2")
	port := uint16(buff[0])<<8 | uint16(buff[1])
	if port < 1 || port > 0xffff {
		replay[1] = 4
		conn.Write(replay)
		return errors.New("目标端口错误")
	}
	portStr := strconv.Itoa(int(port))
	hostStr = net.JoinHostPort(hostStr, portStr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", hostStr)
	if err != nil {
		replay[1] = 5
		conn.Write(replay)
		return err
	}
	targetConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		replay[1] = 5
		conn.Write(replay)
		return err
	}
	log.Println("目标服务器：" + hostStr)
	defer targetConn.Close()
	replay[1] = 0
	if _, err := conn.Write(replay); err != nil {
		return err
	}

	go func() {
		defer conn.Close()
		defer targetConn.Close()
		io.Copy(conn, targetConn)
	}()

	io.Copy(targetConn, conn)
	return nil
}

func main() {
	localAddress := "127.0.0.1:5230"
	ip, _ := net.ResolveTCPAddr("tcp", localAddress)
	ln, err := net.ListenTCP("tcp", ip)
	defer ln.Close()
	if err != nil {
		log.Println("监听端口失败")
		log.Println(err)
		return
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			continue
		}
		go handelConnection(conn)
	}
}