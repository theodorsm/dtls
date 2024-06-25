module github.com/pion/dtls/v2

require (
	github.com/pion/logging v0.2.2
	github.com/pion/transport/v3 v3.0.2
	golang.org/x/crypto v0.24.0
	golang.org/x/net v0.26.0
)

require (
	github.com/google/gopacket v1.1.19 // indirect
	github.com/theodorsm/covert-dtls v0.0.0-20240620110753-528715b4b459 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
replace github.com/theodorsm/covert-dtls v0.0.0-20240620110753-528715b4b459 => ../covertDTLS

go 1.19
