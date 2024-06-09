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

const (
	magicCookie = 0x2112A442
	fingerprint = 0x5354554e
)

// NATType is the type of NAT described by int.
type NATType int

// NAT types.
const (
	NATError NATType = iota
	NATNone
	NATBlocked
	NATFull
	NATSymmetric
	NATRestricted
	NATPortRestricted
	SymmetricUDPFirewall
)

var natStr map[NATType]string

func init() {
	natStr = map[NATType]string{
		NATError:             "Test failed",
		NATNone:              "Not behind a NAT",
		NATBlocked:           "UDP is blocked",
		NATFull:              "Full cone NAT",
		NATSymmetric:         "Symmetric NAT",
		NATRestricted:        "Restricted NAT",
		NATPortRestricted:    "Port restricted NAT",
		SymmetricUDPFirewall: "Symmetric UDP firewall",
	}
}

func (nat NATType) String() string {
	if s, ok := natStr[nat]; ok {
		return s
	}
	return "Unknown"
}

const (
	attributeFamilyIPv4 = 0x01
	attributeFamilyIPV6 = 0x02
)

const (
	attributeMappedAddress       = 0x0001
	attributeChangeRequest       = 0x0003
	attributeSourceAddress       = 0x0004
	attributeChangedAddress      = 0x0005
	attributeXorMappedAddress    = 0x0020
	attributeXorMappedAddressExp = 0x8020
	attributeSoftware            = 0x8022
	attributeFingerprint         = 0x8028
	attributeOtherAddress        = 0x802c
)

const (
	typeBindingRequest = 0x0001
)
