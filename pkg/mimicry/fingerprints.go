// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package mimicry

//nolint:revive
type ClientHelloFingerprint string

// These fingerprints are added automatically generated and added by the 'fingerprint' workflow
// The first byte should correspond to the DTLS version in a handshake message
const (
	Mozilla_Firefox_125_0 ClientHelloFingerprint = "fefda9d2765eaa8856d81195252cd120135b1f1a67c059a5b73d1bd3ed02c07421b800000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200" //nolint:revive,stylecheck
)

func getClientHelloFingerprints() []ClientHelloFingerprint {
	return []ClientHelloFingerprint{
		Mozilla_Firefox_125_0, //nolint:revive,stylecheck
	}
}
