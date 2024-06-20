// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import "github.com/theodorsm/dlts/v2/pkg/protocol"

func defaultCompressionMethods() []*protocol.CompressionMethod {
	return []*protocol.CompressionMethod{
		{},
	}
}
