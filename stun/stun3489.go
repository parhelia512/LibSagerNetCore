package stun

import (
	"fmt"
	"net"

	"github.com/ccding/go-stun/stun"
	"github.com/sirupsen/logrus"

	"libcore/clash/transport/socks5"
)

//go:generate go run ../errorgen

func connectLegacy(addrStr string, socksPort int) (net.PacketConn, error) {
	addr, err := net.ResolveUDPAddr("udp", addrStr)
	if err != nil {
		return nil, newError("failed to resolve server address ", addrStr).Base(err)
	}
	logrus.Info(newError("connecting to STUN server: ", addrStr))
	var mapTestConn net.PacketConn
	socksConn, err := net.Dial("tcp", fmt.Sprint("127.0.0.1:", socksPort))
	if err == nil {
		handshake, err := socks5.ClientHandshake(socksConn, socks5.ParseAddr("0.0.0.0:0"), socks5.CmdUDPAssociate, nil)
		if err != nil {
			logrus.Warn(newError("failed to do udp associate handshake").Base(err))
		}
		udpConn, err := net.DialUDP("udp", nil, handshake.UDPAddr())
		if err == nil {
			mapTestConn = &socksPacketConn{udpConn, socksConn}
		}
	}

	if mapTestConn == nil {
		mapTestConn, err = net.ListenUDP("udp", nil)
		if err != nil {
			return nil, newError("failed to listen udp").Base(err)
		}
	}

	logrus.Info(newError("local address: ", mapTestConn.LocalAddr()))
	logrus.Info(newError("remote address: ", addr))

	return mapTestConn, nil
}

// RFC 3489
func TestLegacy(addrStr string, socksPort int) (natType stun.NATType, err error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	var mapTestConn net.PacketConn
	newConn := func() error {
		if err == nil {
			mapTestConn, err = connectLegacy(addrStr, socksPort)
			if err != nil {
				e := newError("error creating STUN connection").Base(err)
				logrus.Warn(e)
				return e
			}
		}
		return err
	}
	if newConn() == nil {
		client := stun.NewClientWithConnection(mapTestConn)
		client.SetServerAddr(addrStr)
		natType, _, err = client.Discover()
		return
	}
	return
}
