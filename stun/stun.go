package stun

import (
	"fmt"
	"net"

	"github.com/ccding/go-stun/stun"
	"github.com/sirupsen/logrus"

	"libcore/clash/transport/socks5"
)

//go:generate go run ../errorgen

func connect(addrStr string, socksPort int) (net.PacketConn, error) {
	addr, err := net.ResolveUDPAddr("udp", addrStr)
	if err != nil {
		return nil, newError("failed to resolve server address ", addrStr).Base(err)
	}
	logrus.Info(newError("connecting to STUN server: ", addrStr))
	var packetConn net.PacketConn
	socksConn, err := net.Dial("tcp", fmt.Sprint("127.0.0.1:", socksPort))
	if err == nil {
		handshake, err := socks5.ClientHandshake(socksConn, socks5.ParseAddr("0.0.0.0:0"), socks5.CmdUDPAssociate, nil)
		if err != nil {
			logrus.Warn(newError("failed to do udp associate handshake").Base(err))
		}
		udpConn, err := net.DialUDP("udp", nil, handshake.UDPAddr())
		if err == nil {
			packetConn = &socksPacketConn{udpConn, socksConn}
		}
	}

	if packetConn == nil {
		packetConn, err = net.ListenUDP("udp", nil)
		if err != nil {
			return nil, newError("failed to listen udp").Base(err)
		}
	}

	logrus.Info(newError("local address: ", packetConn.LocalAddr()))
	logrus.Info(newError("remote address: ", addr))

	return packetConn, nil
}

func newConn(addrStr string, socksPort int) (net.PacketConn, error) {
	packetConn, err := connect(addrStr, socksPort)
	if err != nil {
		e := newError("error creating STUN connection").Base(err)
		logrus.Warn(e)
		return nil, e
	}
	return packetConn, nil
}

// RFC 5780
func Test(addrStr string, socksPort int) (*stun.NATBehavior, error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	if packetConn, err := newConn(addrStr, socksPort); err == nil {
		client := stun.NewClientWithConnection(packetConn)
		client.SetServerAddr(addrStr)
		return client.BehaviorTest()
	} else {
		return &stun.NATBehavior{
			MappingType:   stun.BehaviorTypeUnknown,
			FilteringType: stun.BehaviorTypeUnknown,
		}, err
	}
}

// RFC 3489
func TestLegacy(addrStr string, socksPort int) (stun.NATType, *stun.Host, error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	if packetConn, err := newConn(addrStr, socksPort); err == nil {
		client := stun.NewClientWithConnection(packetConn)
		client.SetServerAddr(addrStr)
		return client.Discover()
	} else {
		return stun.NATError, nil, err
	}
}
