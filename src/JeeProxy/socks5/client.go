package socks5

import (
	"net"
	"errors"
	"log"
	"strconv"
	"io"
)

type Client struct {
	conn net.Conn
}

//转发
func (t *Client) Proxy() error {
	buff := make([]byte, 262)
	reply := []byte{0x05, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x22, 0x22}
	if _, err := io.ReadFull(t.conn, buff[:2]); err != nil {
		return err
	}
	if buff[0] != 0x05 {
		reply[1] = 0xff
		t.conn.Write(reply[:2])
		return errors.New("客户端协议错误")
	}
	numMethod := buff[1]
	if _, err := io.ReadFull(t.conn, buff[:numMethod]); err != nil {
		return err
	}
	reply[1] = 0
	if _, err := t.conn.Write(reply[:2]); err != nil {
		return err
	}

	if _, err := io.ReadFull(t.conn, buff[:4]); err != nil {
		return err
	}
	if buff[1] != 1 {
		reply[1] = 7
		t.conn.Write(reply)
		return errors.New("客户端命令不支持")
	}

	addressType := buff[3]
	addressLen := 0
	switch addressType {
	case 1:
		addressLen = net.IPv4len
	case 4:
		addressLen = net.IPv6len
	case 3:
		if _, err := io.ReadFull(t.conn, buff[:1]); err != nil {
			return err
		}
		addressLen = int(buff[0])
	default:
		reply[1] = 8
		t.conn.Write(reply)
		return errors.New("目标主机地址类型不在协议规定范围")
	}
	host := make([]byte, addressLen)
	if _, err := io.ReadFull(t.conn, host); err != nil {
		return err
	}
	if _, err := io.ReadFull(t.conn, buff[:2]); err != nil {
		return err
	}
	hostStr := ""
	switch addressType {
	case 1, 4:
		ip := net.IP(host)
		hostStr = ip.String()
	case 3:
		hostStr = string(host)
	}
	port := uint16(buff[0])<<8 | uint16(buff[1])
	if port < 1 || port > 0xffff {
		reply[1] = 4
		t.conn.Write(reply)
		return errors.New("目标端口错误")
	}
	portStr := strconv.Itoa(int(port))
	hostStr = net.JoinHostPort(hostStr, portStr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", hostStr)
	if err != nil {
		reply[1] = 5
		t.conn.Write(reply)
		return err
	}
	targetConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		reply[1] = 5
		t.conn.Write(reply)
		return err
	}
	log.Println("目标服务器：" + hostStr)
	defer targetConn.Close()
	reply[1] = 0
	if _, err := t.conn.Write(reply); err != nil {
		return err
	}

	go func() {
		defer t.conn.Close()
		defer targetConn.Close()
		io.Copy(t.conn, targetConn)
	}()

	io.Copy(targetConn, t.conn)
	return nil
}

func NewSocks5Client(conn net.Conn) {
	defer conn.Close()
	c := &Client{
		conn: conn,
	}
	err := c.Proxy()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
