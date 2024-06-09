package socks

import (
	"errors"
	"net"

	"libcore/clash/transport/socks5"
)

func NewSocksConn(tcpConn net.Conn, addrStr string) *PacketConn {
	handshake, err := socks5.ClientHandshake(tcpConn, socks5.ParseAddr(addrStr), socks5.CmdUDPAssociate, nil)
	if err != nil {
		return nil
	}
	udpConn, err := net.DialUDP("udp", nil, handshake.UDPAddr())
	if err != nil {
		return nil
	}
	return &PacketConn{
		UDPConn: udpConn,
		TCPConn: tcpConn,
	}
}

type PacketConn struct {
	*net.UDPConn
	TCPConn net.Conn
}

func (uc *PacketConn) WriteTo(b []byte, addr net.Addr) (n int, err error) {
	packet, err := socks5.EncodeUDPPacket(socks5.ParseAddrToSocksAddr(addr), b)
	if err != nil {
		return
	}
	return uc.UDPConn.Write(packet)
}

func (uc *PacketConn) ReadFrom(b []byte) (int, net.Addr, error) {
	_, _, err := uc.UDPConn.ReadFrom(b)
	if err != nil {
		return 0, nil, err
	}
	addr, payload, err := socks5.DecodeUDPPacket(b)
	if err != nil {
		return 0, nil, err
	}

	udpAddr := addr.UDPAddr()
	if udpAddr == nil {
		return 0, nil, errors.New("parse udp addr error")
	}

	// due to DecodeUDPPacket is mutable, record addr length
	copy(b, payload)
	return len(payload), udpAddr, nil
}

func (uc *PacketConn) Close() error {
	uc.TCPConn.Close()
	return uc.UDPConn.Close()
}
