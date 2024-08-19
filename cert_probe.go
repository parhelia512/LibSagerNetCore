// modified from https://github.com/xchacha20-poly1305/TLS-scribe, license: MIT

package libcore

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/wzshiming/socks5"
)

func ProbeCertTLS(ctx context.Context, address, sni string, port int32) ([]*x509.Certificate, error) {
	socks5Dialer, _ := socks5.NewDialer("socks5h://127.0.0.1:" + strconv.Itoa(int(port)))
	conn, err := socks5Dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		dialer := &net.Dialer{}
		conn, err = dialer.DialContext(ctx, "tcp", address)
	}
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tlsConn := tls.Client(conn, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"h2", "http/1.1"},
		ServerName:         sni,
	})
	err = tlsConn.HandshakeContext(ctx)
	if err != nil {
		return nil, err
	}
	defer tlsConn.Close()
	return tlsConn.ConnectionState().PeerCertificates, nil
}

type udpAddr struct {
	address string
}

func (a *udpAddr) Network() string {
	return "udp"
}

func (a *udpAddr) String() string {
	return a.address
}

func ProbeCertQUIC(ctx context.Context, address, sni string, socksPort int32) ([]*x509.Certificate, error) {
	socks5Dialer, _ := socks5.NewDialer("socks5h://127.0.0.1:" + strconv.Itoa(int(socksPort)))
	var packetConn net.PacketConn
	var addr net.Addr
	conn, err := socks5Dialer.DialContext(ctx, "udp", address)
	if err != nil {
		packetConn, err = net.ListenUDP("udp", nil)
		if err != nil {
			return nil, err
		}
		defer packetConn.Close()
		addr, err = net.ResolveUDPAddr("udp", address)
		if err != nil {
			return nil, err
		}
	} else {
		defer conn.Close()
		packetConn = conn.(*socks5.UDPConn)
		addr = &udpAddr{address: address}
	}
	quicConn, err := quic.Dial(ctx, packetConn, addr, &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"h3"},
		ServerName:         sni,
	}, &quic.Config{Versions: []quic.Version{quic.Version2, quic.Version1}})
	if err != nil {
		return nil, err
	}
	defer quicConn.CloseWithError(0x00, "")
	return quicConn.ConnectionState().TLS.PeerCertificates, nil
}

func ProbeCert(address, sni, protocol string, socksPort int32) (cert string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var certs []*x509.Certificate
	switch protocol {
	case "tls":
		certs, err = ProbeCertTLS(ctx, address, sni, socksPort)
	case "quic":
		certs, err = ProbeCertQUIC(ctx, address, sni, socksPort)
	default:
		err = newError("unknown protocol: ", protocol)
	}
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	for _, cert := range certs {
		err = pem.Encode(&builder, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
