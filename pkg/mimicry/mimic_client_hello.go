// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package mimicry provides replay capabilities of captured handshakes for anti-fingerprinting purposes.
package mimicry

import (
	"encoding/hex"
	"errors"

	"github.com/pion/dtls/v2/pkg/protocol/handshake"
)

var errBufferTooSmall = errors.New("buffer is too small") //nolint:goerr113

//nolint:revive
type MimickedClientHello struct {
	ClientHelloFingerprint ClientHelloFingerprint
	Random                 handshake.Random
	SessionID              []byte
	Cookie                 []byte
}

//nolint:revive
func (m MimickedClientHello) Type() handshake.Type {
	return handshake.TypeClientHello
}

//nolint:revive
func (m *MimickedClientHello) Marshal() ([]byte, error) {
	var out []byte

	fingerprints := getClientHelloFingerprints()

	if len(fingerprints) < 1 {
		return out, errors.New("no fingerprints available") //nolint:goerr113
	}

	fingerprint := m.ClientHelloFingerprint

	if string(fingerprint) == "" {
		fingerprint = fingerprints[len(fingerprints)-1]
	}

	data, err := hex.DecodeString(string(fingerprint))
	if err != nil {
		err = errors.New("mimicry: failed to decode mimicry hexstring") //nolint:goerr113
	}

	if len(data) <= 2 {
		return out, errors.New("mimicked fingerprint is too short") //nolint:goerr113
	}

	// Major and minor version
	currOffset := 2
	out = append(out, data[:currOffset]...)

	rb := m.Random.MarshalFixed()
	out = append(out, rb[:]...)

	// Skip past random
	currOffset += 32

	currOffset++
	if len(data) <= currOffset {
		return out, errBufferTooSmall
	}
	n := int(data[currOffset-1])
	if len(data) <= currOffset+n {
		return out, errBufferTooSmall
	}
	mimickedSessionID := append([]byte{}, data[currOffset:currOffset+n]...)
	currOffset += len(mimickedSessionID)

	currOffset++
	if len(data) <= currOffset {
		return out, errBufferTooSmall
	}
	n = int(data[currOffset-1])
	if len(data) <= currOffset+n {
		return out, errBufferTooSmall
	}
	mimickedCookie := append([]byte{}, data[currOffset:currOffset+n]...)
	currOffset += len(mimickedCookie)

	out = append(out, byte(len(m.SessionID)))
	out = append(out, m.SessionID...)

	out = append(out, byte(len(m.Cookie)))
	out = append(out, m.Cookie...)

	out = append(out, data[currOffset:]...)

	return out, err
}

//nolint:revive
func (m *MimickedClientHello) Unmarshal(data []byte) error { return nil }
