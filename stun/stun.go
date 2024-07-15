package stun

import (
	"context"
	"net"
	"strconv"

	"github.com/ccding/go-stun/stun"
	"github.com/wzshiming/socks5"
)

func setupPacketConn(socksPort int) (net.PacketConn, bool, error) {
	dialer, err := socks5.NewDialer("socks5h://127.0.0.1:" + strconv.Itoa(socksPort))
	if err != nil {
		return nil, false, err
	}
	var packetConn net.PacketConn
	useSOCKS5 := false
	conn, err := dialer.Dial("udp", "0.0.0.0:0")
	if err == nil {
		packetConn = conn.(*socks5.UDPConn)
		useSOCKS5 = true
	} else {
		packetConn, err = net.ListenUDP("udp", nil)
	}
	if err != nil {
		return nil, false, err
	}
	return packetConn, useSOCKS5, nil
}

func resolveDNS(host string, dnsPort int) (net.IP, error) {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dialer := new(net.Dialer)
			return dialer.DialContext(ctx, network, "127.0.0.1:"+strconv.Itoa(dnsPort))
		},
	}
	ips, err := resolver.LookupIP(context.Background(), "ip", host)
	if err != nil {
		return nil, err
	}
	return ips[0], nil
}

// RFC 5780
func Test(addrStr string, socksPort int, dnsPort int) (*stun.NATBehavior, error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	host, port, err := net.SplitHostPort(addrStr)
	if err != nil {
		return nil, err
	}
	packetConn, useSOCKS5, err := setupPacketConn(socksPort)
	if err != nil {
		return nil, err
	}
	if useSOCKS5 && net.ParseIP(host) == nil {
		ip, err := resolveDNS(host, dnsPort)
		if err != nil {
			return nil, err
		}
		addrStr = net.JoinHostPort(ip.String(), port)
	}
	client := stun.NewClientWithConnection(packetConn)
	client.SetServerAddr(addrStr)
	return client.BehaviorTest()
}

// RFC 3489
func TestLegacy(addrStr string, socksPort int, dnsPort int) (*stun.NATType, *stun.Host, error) {
	if addrStr == "" {
		addrStr = "stun.syncthing.net:3478"
	}
	host, port, err := net.SplitHostPort(addrStr)
	if err != nil {
		return nil, nil, err
	}
	packetConn, useSOCKS5, err := setupPacketConn(socksPort)
	if err != nil {
		return nil, nil, err
	}
	if useSOCKS5 && net.ParseIP(host) == nil {
		ip, err := resolveDNS(host, dnsPort)
		if err != nil {
			return nil, nil, err
		}
		addrStr = net.JoinHostPort(ip.String(), port)
	}
	client := stun.NewClientWithConnection(packetConn)
	client.SetServerAddr(addrStr)
	natType, addr, err := client.Discover()
	return &natType, addr, err
}
