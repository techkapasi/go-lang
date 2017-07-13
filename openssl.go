package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"io"
	"math/big"
	"crypto/tls"
	"net"
	"strings"
)


var (
	// idReader is used for random id generation. This declaration allows us to
	// replace it for testing.
	idReader = cryptorand.Reader
)

// parameters for random identifier generation. We can tweak this when there is
// time for further analysis.
const (
	randomIDEntropyBytes = 17
	randomIDBase         = 36

	// To ensure that all identifiers are fixed length, we make sure they
	// get padded out or truncated to 25 characters.
	//
	// For academics,  f5lxx1zz5pnorynqglhzmsp33  == 2^128 - 1. This value
	// was calculated from floor(log(2^128-1, 36)) + 1.
	//
	// While 128 bits is the largest whole-byte size that fits into 25
	// base-36 characters, we generate an extra byte of entropy to fill
	// in the high bits, which would otherwise be 0. This gives us a more
	// even distribution of the first character.
	//
	// See http://mathworld.wolfram.com/NumberLength.html for more information.
	maxRandomIDLength = 25
)

// NewID generates a new identifier for use where random identifiers with low
// collision probability are required.
//
// With the parameters in this package, the generated identifier will provide
// ~129 bits of entropy encoded with base36. Leading padding is added if the
// string is less 25 bytes. We do not intend to maintain this interface, so
// identifiers should be treated opaquely.
func NewID() string {
	var p [randomIDEntropyBytes]byte

	if _, err := io.ReadFull(idReader, p[:]); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}

	p[0] |= 0x80 // set high bit to avoid the need for padding
	return (&big.Int{}).SetBytes(p[:]).Text(randomIDBase)[1 : maxRandomIDLength+1]
}



// Clone returns a clone of tls.Config. This function is provided for
// compatibility for go1.7 that doesn't include this method in stdlib.
func Clone(c *tls.Config) *tls.Config {
	return &tls.Config{
		Rand:                        c.Rand,
		Time:                        c.Time,
		Certificates:                c.Certificates,
		NameToCertificate:           c.NameToCertificate,
		GetCertificate:              c.GetCertificate,
		RootCAs:                     c.RootCAs,
		NextProtos:                  c.NextProtos,
		ServerName:                  c.ServerName,
		ClientAuth:                  c.ClientAuth,
		ClientCAs:                   c.ClientCAs,
		InsecureSkipVerify:          c.InsecureSkipVerify,
		CipherSuites:                c.CipherSuites,
		PreferServerCipherSuites:    c.PreferServerCipherSuites,
		SessionTicketsDisabled:      c.SessionTicketsDisabled,
		SessionTicketKey:            c.SessionTicketKey,
		ClientSessionCache:          c.ClientSessionCache,
		MinVersion:                  c.MinVersion,
		MaxVersion:                  c.MaxVersion,
		CurvePreferences:            c.CurvePreferences,
		DynamicRecordSizingDisabled: c.DynamicRecordSizingDisabled,
		Renegotiation:               c.Renegotiation,
	}
}

package listeners

import (
"crypto/tls"
"fmt"
"net"
"strings"

"github.com/Microsoft/go-winio"
"github.com/docker/go-connections/sockets"
)

// Init creates new listeners for the server.
func Init(proto, addr, socketGroup string, tlsConfig *tls.Config) ([]net.Listener, error) {
	ls := []net.Listener{}

	switch proto {
	case "tcp":
		l, err := sockets.NewTCPSocket(addr, tlsConfig)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)

	case "npipe":
		// allow Administrators and SYSTEM, plus whatever additional users or groups were specified
		sddl := "D:P(A;;GA;;;BA)(A;;GA;;;SY)"
		if socketGroup != "" {
			for _, g := range strings.Split(socketGroup, ",") {
				sid, err := winio.LookupSidByName(g)
				if err != nil {
					return nil, err
				}
				sddl += fmt.Sprintf("(A;;GRGW;;;%s)", sid)
			}
		}
		c := winio.PipeConfig{
			SecurityDescriptor: sddl,
			MessageMode:        true,  // Use message mode so that CloseWrite() is supported
			InputBufferSize:    65536, // Use 64KB buffers to improve performance
			OutputBufferSize:   65536,
		}
		l, err := winio.ListenPipe(addr, &c)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)

	default:
		return nil, fmt.Errorf("invalid protocol format: windows only supports tcp and npipe")
	}

	return ls, nil
}


func main () {
	id := NewID()
	println("ID is " + id)
	//cfg *tls.Config{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "" }
	//println(Clone(cfg))

	Init()

}