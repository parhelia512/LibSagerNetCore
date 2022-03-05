/*
Copyright (C) 2021 by nekohasekai <sekai@neko.services>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

// https://github.com/SagerNet/libping/blob/593b070fbd74a44e068e4319fe6c5863bd110698/ping.go#L1-L100

package libcore

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"strings"
	"time"
)

const payload = "abcdefghijklmnopqrstuvwabcdefghi"

func icmpPing(address string, timeout int32) (int32, error) {
	i := net.ParseIP(address)
	if i == nil {
		return 0, fmt.Errorf("unable to parse ip %s", address)
	}
	var err error
	v6 := i.To4() == nil
	var fd int
	if !v6 {
		fd, err = unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_ICMP)
	} else {
		fd, err = unix.Socket(unix.AF_INET6, unix.SOCK_DGRAM, unix.IPPROTO_ICMPV6)
	}

	f := os.NewFile(uintptr(fd), "dgram")
	if err != nil {
		return 0, newError("create file from fd").Base(err)
	}

	conn, err := net.FilePacketConn(f)
	if err != nil {
		return 0, newError("create conn").Base(err)
	}

	defer func(conn net.PacketConn) {
		_ = conn.Close()
	}(conn)

	start := time.Now()
	for seq := 1; timeout > 0; seq++ {
		var sockTo int32
		if timeout > 1000 {
			sockTo = 1000
		} else {
			sockTo = timeout
		}
		timeout -= sockTo

		err := conn.SetReadDeadline(time.Now().Add(time.Duration(sockTo) * time.Millisecond))
		if err != nil {
			return 0, newError("set read timeout").Base(err)
		}

		msg := icmp.Message{
			Body: &icmp.Echo{
				ID:   0xDBB,
				Seq:  seq,
				Data: []byte(payload),
			},
		}
		if !v6 {
			msg.Type = ipv4.ICMPTypeEcho
		} else {
			msg.Type = ipv6.ICMPTypeEchoRequest
		}

		data, err := msg.Marshal(nil)
		if err != nil {
			return 0, newError("make icmp message").Base(err)
		}

		_, err = conn.WriteTo(data, &net.UDPAddr{
			IP:   i,
			Port: 0,
		})
		if err != nil {
			return 0, newError("write icmp message").Base(err)
		}

		_, _, err = conn.ReadFrom(data)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				continue
			}

			return 0, newError("read icmp message").Base(err)
		}

		return int32(time.Since(start).Milliseconds()), nil
	}

	return -1, nil
}
