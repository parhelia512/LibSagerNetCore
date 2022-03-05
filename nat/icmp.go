package nat

import (
	"gvisor.dev/gvisor/pkg/tcpip/header"
)

func (t *SystemTun) processICMPv4(ipHdr header.IPv4, hdr header.ICMPv4) bool {
	if hdr.Type() != header.ICMPv4Echo || hdr.Code() != header.ICMPv4UnusedCode {
		return false
	}

	sourceAddress := ipHdr.SourceAddress()
	ipHdr.SetSourceAddress(ipHdr.DestinationAddress())
	ipHdr.SetDestinationAddress(sourceAddress)
	ipHdr.SetChecksum(0)
	ipHdr.SetChecksum(^ipHdr.CalculateChecksum())

	hdr.SetType(header.ICMPv4EchoReply)
	hdr.SetChecksum(0)
	hdr.SetChecksum(header.ICMPv4Checksum(hdr, 0))
	t.writeBuffer(ipHdr)
	return false
}

func (t *SystemTun) processICMPv6(ipHdr header.IPv6, hdr header.ICMPv6) bool {
	if hdr.Type() != header.ICMPv6EchoRequest || hdr.Code() != header.ICMPv6UnusedCode {
		return false
	}

	sourceAddress := ipHdr.SourceAddress()
	ipHdr.SetSourceAddress(ipHdr.DestinationAddress())
	ipHdr.SetDestinationAddress(sourceAddress)

	hdr.SetType(header.ICMPv6EchoReply)
	hdr.SetChecksum(0)
	hdr.SetChecksum(header.ICMPv6Checksum(header.ICMPv6ChecksumParams{
		Header: hdr,
		Src:    ipHdr.SourceAddress(),
		Dst:    ipHdr.DestinationAddress(),
	}))
	t.writeBuffer(ipHdr)
	return false
}
