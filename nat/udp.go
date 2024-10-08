package nat

import (
	"github.com/v2fly/v2ray-core/v5/common/buf"
	v2rayNet "github.com/v2fly/v2ray-core/v5/common/net"
	"gvisor.dev/gvisor/pkg/buffer"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/checksum"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
	"libcore/comm"
)

func (t *SystemTun) processIPv4UDP(cache *buf.Buffer, ipHdr header.IPv4, hdr header.UDP) {
	sourceAddress := ipHdr.SourceAddress()
	destinationAddress := ipHdr.DestinationAddress()
	sourcePort := hdr.SourcePort()
	destinationPort := hdr.DestinationPort()

	source := v2rayNet.Destination{
		Address: v2rayNet.IPAddress(sourceAddress.AsSlice()),
		Port:    v2rayNet.Port(sourcePort),
		Network: v2rayNet.Network_UDP,
	}
	destination := v2rayNet.Destination{
		Address: v2rayNet.IPAddress(destinationAddress.AsSlice()),
		Port:    v2rayNet.Port(destinationPort),
		Network: v2rayNet.Network_UDP,
	}

	ipHdr.SetDestinationAddress(sourceAddress)
	hdr.SetDestinationPort(sourcePort)

	headerLength := ipHdr.HeaderLength()
	headerCache := buf.NewWithSize(int32(t.mtu))
	headerCache.Write(ipHdr[:headerLength+header.UDPMinimumSize])

	cache.Advance(int32(headerLength + header.UDPMinimumSize))
	go t.handler.NewPacket(source, destination, cache, func(bytes []byte, addr *v2rayNet.UDPAddr) (int, error) {
		index := headerCache.Len()
		newHeader := headerCache.Extend(index)
		copy(newHeader, headerCache.Bytes())
		headerCache.Advance(index)

		defer func() {
			headerCache.Clear()
			headerCache.Resize(0, index)
		}()

		var newSourceAddress tcpip.Address
		var newSourcePort uint16

		if addr != nil {
			newSourceAddress = tcpip.AddrFromSlice(addr.IP)
			newSourcePort = uint16(addr.Port)
		} else {
			newSourceAddress = destinationAddress
			newSourcePort = destinationPort
		}

		newIpHdr := header.IPv4(newHeader)
		newIpHdr.SetSourceAddress(newSourceAddress)
		newIpHdr.SetTotalLength(uint16(int(headerCache.Len()) + len(bytes)))
		newIpHdr.SetChecksum(0)
		newIpHdr.SetChecksum(^newIpHdr.CalculateChecksum())

		udpHdr := header.UDP(headerCache.BytesFrom(headerCache.Len() - header.UDPMinimumSize))
		udpHdr.SetSourcePort(newSourcePort)
		udpHdr.SetLength(uint16(header.UDPMinimumSize + len(bytes)))
		udpHdr.SetChecksum(0)
		udpHdr.SetChecksum(^udpHdr.CalculateChecksum(checksum.Checksum(bytes, header.PseudoHeaderChecksum(header.UDPProtocolNumber, newSourceAddress, sourceAddress, uint16(header.UDPMinimumSize+len(bytes))))))

		payload := buffer.MakeWithData(newHeader)
		payload.Append(buffer.NewViewWithData(bytes))

		pkt := stack.NewPacketBuffer(stack.PacketBufferOptions{
			Payload: payload,
		})
		if err := t.writeRawPacket(pkt); err != nil {
			return 0, newError(err.String())
		}

		return len(bytes), nil
	}, comm.Closer(headerCache.Release))
}

func (t *SystemTun) processIPv6UDP(cache *buf.Buffer, ipHdr header.IPv6, hdr header.UDP) {
	sourceAddress := ipHdr.SourceAddress()
	destinationAddress := ipHdr.DestinationAddress()
	sourcePort := hdr.SourcePort()
	destinationPort := hdr.DestinationPort()

	source := v2rayNet.Destination{
		Address: v2rayNet.IPAddress(sourceAddress.AsSlice()),
		Port:    v2rayNet.Port(sourcePort),
		Network: v2rayNet.Network_UDP,
	}
	destination := v2rayNet.Destination{
		Address: v2rayNet.IPAddress(destinationAddress.AsSlice()),
		Port:    v2rayNet.Port(destinationPort),
		Network: v2rayNet.Network_UDP,
	}

	ipHdr.SetDestinationAddress(sourceAddress)
	hdr.SetDestinationPort(sourcePort)

	headerLength := uint16(len(ipHdr)) - ipHdr.PayloadLength()
	headerCache := buf.NewWithSize(int32(t.mtu))
	headerCache.Write(ipHdr[:headerLength+header.UDPMinimumSize])

	cache.Advance(int32(headerLength + header.UDPMinimumSize))
	go t.handler.NewPacket(source, destination, cache, func(bytes []byte, addr *v2rayNet.UDPAddr) (int, error) {
		index := headerCache.Len()
		newHeader := headerCache.Extend(index)
		copy(newHeader, headerCache.Bytes())
		headerCache.Advance(index)

		defer func() {
			headerCache.Clear()
			headerCache.Resize(0, index)
		}()

		var newSourceAddress tcpip.Address
		var newSourcePort uint16

		if addr != nil {
			newSourceAddress = tcpip.AddrFromSlice(addr.IP)
			newSourcePort = uint16(addr.Port)
		} else {
			newSourceAddress = destinationAddress
			newSourcePort = destinationPort
		}

		newIpHdr := header.IPv6(newHeader)
		newIpHdr.SetSourceAddress(newSourceAddress)
		newIpHdr.SetPayloadLength(uint16(header.UDPMinimumSize + len(bytes)))

		udpHdr := header.UDP(headerCache.BytesFrom(headerCache.Len() - header.UDPMinimumSize))
		udpHdr.SetSourcePort(newSourcePort)
		udpHdr.SetLength(uint16(header.UDPMinimumSize + len(bytes)))
		udpHdr.SetChecksum(0)
		udpHdr.SetChecksum(^udpHdr.CalculateChecksum(checksum.Checksum(bytes, header.PseudoHeaderChecksum(header.UDPProtocolNumber, newSourceAddress, sourceAddress, uint16(header.UDPMinimumSize+len(bytes))))))

		payload := buffer.MakeWithData(newHeader)
		payload.Append(buffer.NewViewWithData(bytes))

		pkt := stack.NewPacketBuffer(stack.PacketBufferOptions{
			Payload: payload,
		})
		if err := t.writeRawPacket(pkt); err != nil {
			return 0, newError(err.String())
		}

		return len(bytes), nil
	}, comm.Closer(headerCache.Release))
}
