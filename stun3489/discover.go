// Copyright 2016 Cong Ding
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stun3489

import (
	"net"

	"github.com/sirupsen/logrus"
)

func (c *Client) Discover(conn net.PacketConn, addr *net.UDPAddr) (NATType, error) {
	// Perform test1 to check if it is under NAT.
	logrus.Info(newError("Do Test1"))
	logrus.Info(newError("Send To:", addr))
	resp, err := c.test1(conn, addr)
	if err != nil {
		return NATError, err
	}
	logrus.Info(newError("Received:", resp))
	if resp == nil {
		return NATBlocked, nil
	}
	// identical used to check if it is open Internet or not.
	identical := resp.identical
	// changedAddr is used to perform second time test1 and test3.
	changedAddr := resp.changedAddr
	// mappedAddr is used as the return value, its IP is used for tests
	mappedAddr := resp.mappedAddr
	// Make sure IP and port are not changed.
	if resp.serverAddr.IP() != addr.IP.String() ||
		resp.serverAddr.Port() != uint16(addr.Port) {
		return NATError, newError("Server error: response IP/port")
	}
	// if changedAddr is not available, use otherAddr as changedAddr,
	// which is updated in RFC 5780
	if changedAddr == nil {
		changedAddr = resp.otherAddr
	}
	// changedAddr shall not be nil
	if changedAddr == nil {
		return NATError, newError("Server error: no changed address.")
	}
	// Perform test2 to see if the client can receive packet sent from
	// another IP and port.
	logrus.Info(newError("Do Test2", resp))
	logrus.Info(newError("Send To:", addr))
	resp, err = c.test2(conn, addr)
	if err != nil {
		return NATError, err
	}
	logrus.Info(newError("Received:", resp))
	// Make sure IP and port are changed.
	if resp != nil &&
		(resp.serverAddr.IP() == addr.IP.String() ||
			resp.serverAddr.Port() == uint16(addr.Port)) {
		return NATError, newError("Server error: response IP/port")
	}
	if identical {
		if resp == nil {
			return SymmetricUDPFirewall, nil
		}
		return NATNone, nil
	}
	if resp != nil {
		return NATFull, nil
	}
	// Perform test1 to another IP and port to see if the NAT use the same
	// external IP.
	logrus.Info(newError("Do Test1"))
	logrus.Info(newError("Send To:", changedAddr))
	caddr, err := net.ResolveUDPAddr("udp", changedAddr.String())
	if err != nil {
		logrus.Info(newError("ResolveUDPAddr error: %v", err))
		return NATError, err
	}
	resp, err = c.test1(conn, caddr)
	if err != nil {
		return NATError, err
	}
	logrus.Info(newError("Received:", resp))
	// Make sure IP/port is not changed.
	if resp.serverAddr.IP() != caddr.IP.String() ||
		resp.serverAddr.Port() != uint16(caddr.Port) {
		return NATError, newError("Server error: response IP/port")
	}
	if mappedAddr.IP() == resp.mappedAddr.IP() && mappedAddr.Port() == resp.mappedAddr.Port() {
		// Perform test3 to see if the client can receive packet sent
		// from another port.
		logrus.Info(newError("Do Test3"))
		logrus.Info(newError("Send To:", caddr))
		resp, err = c.test3(conn, caddr)
		if err != nil {
			return NATError, err
		}
		logrus.Info(newError("Received:", resp))
		if resp == nil {
			return NATPortRestricted, nil
		}
		// Make sure IP is not changed, and port is changed.
		if resp.serverAddr.IP() != caddr.IP.String() ||
			resp.serverAddr.Port() == uint16(caddr.Port) {
			return NATError, newError("Server error: response IP/port")
		}
		return NATRestricted, nil
	}
	return NATSymmetric, nil
}
