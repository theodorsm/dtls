### Fingerprint mimicking

This submodule offers mimicking of the client hello message which is a anti-fingerprint feature for applications that needs censorship-resistance.

Some sample fingerprints of common browsers are given in `fingerprints.go` (*Note*: the fingerprint is from a webrtc application). Using updated fingerprints is highly recommended, which can be found by an external package such as [dtls-webrtc-fingerprint](https://github.com/theodorsm/dtls-webrtc-fingerprint).

Example with BYOF (bring your own fingerprints):

```go
import  (
  "github.com/pion/dtls/v2/pkg/mimicry"
  "github.com/theodorsm/dtls-webrtc-fingerprint/pkg/fingerprints"
)


// Get recent fingerprints
fingerprints := fingerprints.GetClientHelloFingerprints()

// Choose the freshest fingerprint
fingerprint := fingerprints[len(fingerprints)-1]

cfg := &dtls.Config{
  Certificates:           []tls.Certificate{cert},
	InsecureSkipVerify:     true,
	MimicryEnabled:         true,
	ClientHelloFingerprint: mimicry.ClientHelloFingerprint(fingerprint),
}
```
