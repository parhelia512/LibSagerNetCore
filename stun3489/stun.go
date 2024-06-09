package stun3489

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"libcore/socks"
)

//go:generate go run ../errorgen

func connect(addrStr string, socksPort int) (net.PacketConn, *net.UDPAddr, error) {
	addr, err := net.ResolveUDPAddr("udp", addrStr)
	if err != nil {
		return nil, nil, newError("failed to resolve server address ", addrStr).Base(err)
	}

	logrus.Info(newError("connecting to STUN server: ", addrStr))

	var natConn net.PacketConn

	tcpConn, err := net.Dial("tcp", fmt.Sprint("127.0.0.1:", socksPort))
	if err == nil {
		natConn = socks.NewSocksConn(tcpConn, addrStr)
	}

	if natConn == nil {
		natConn, err = net.ListenUDP("udp", nil)
		if err != nil {
			return nil, nil, newError("failed to listen udp").Base(err)
		}
	}

	logrus.Info(newError("local address: ", natConn.LocalAddr()))
	logrus.Info(newError("remote address: ", addr))

	return natConn, addr, nil
}

// RFC 3489
func Test(addrStr string, socksPort int) (natType NATType, err error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	var mapTestConn net.PacketConn
	var addr *net.UDPAddr
	newConn := func() error {
		if err == nil {
			mapTestConn, addr, err = connect(addrStr, socksPort)
			if err != nil {
				e := newError("error creating STUN connection").Base(err)
				logrus.Warn(e)
				return e
			}
		}
		return err
	}
	if newConn() == nil {
		client := NewClient(mapTestConn, addrStr)
		natType, err = client.Discover(mapTestConn, addr)
		return
	}
	return
}
