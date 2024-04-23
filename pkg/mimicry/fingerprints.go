// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package mimicry

//nolint:revive
type ClientHelloFingerprint string

// These fingerprints are added automatically generated and added by the 'fingerprint' workflow
// The first byte should correspond to the DTLS version in a handshake message
const (
	Mozilla_Firefox_125_0_1             ClientHelloFingerprint = "fefdb22e77791af4658e6fa29c6d396dcdfe51f471e584744683e6c22bd251a2aba000000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200" //nolint:revive,stylecheck
	Google_Chrome_124_0_6367_60_unknown ClientHelloFingerprint = "fefdf77828e1c5bffb408e1ec59f4e4ee4eb30caa71ee64d9f9656e0f98c0d7947fd00000016c02bc02fcca9cca8c009c013c00ac014009c002f00350100004400170000ff01000100000a00080006001d00170018000b0002010000230000000d00140012040308040401050308050501080606010201000e0009000600010008000700"                                                                 //nolint:revive,stylecheck
	Mozilla_Firefox_125_0_2             ClientHelloFingerprint = "fefdd36f0e23957c7fe9b77c21b56d320ab0a75b9d9987795d1046d5296710b1725d00000010c02bc02fcca9cca8c00ac009c013c0140100006a00170000ff01000100000a00080006001d00170018000b000201000010001200100677656272746308632d776562727463000d0020001e040305030603020308040805080604010501060102010402050206020202001c00024000000e000b0008000700080001000200" //nolint:revive,stylecheck
)

func getClientHelloFingerprints() []ClientHelloFingerprint {
	return []ClientHelloFingerprint{
		Mozilla_Firefox_125_0_1,             //nolint:revive,stylecheck
		Google_Chrome_124_0_6367_60_unknown, //nolint:revive,stylecheck
		Mozilla_Firefox_125_0_2,             //nolint:revive,stylecheck
	}
}
