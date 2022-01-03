// Deprecated: This package was moved to github.com/libp2p/go-libp2p-testing/netutil
package testutil

import (
	"testing"

	tnet "github.com/libp2p/go-libp2p-testing/net"
	"github.com/libp2p/go-libp2p-testing/netutil"
)

// TestBogusPrivateKey is a key used for testing (to avoid expensive keygen)
// Deprecated: use the corresponding type in github.com/libp2p/go-libp2p-testing/netutil.
type TestBogusPrivateKey = netutil.TestBogusPrivateKey

// TestBogusPublicKey is a key used for testing (to avoid expensive keygen)
// Deprecated: use the corresponding type in github.com/libp2p/go-libp2p-testing/netutil.
type TestBogusPublicKey = netutil.TestBogusPublicKey

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusPrivateKey() (TestBogusPrivateKey, error) {
	return netutil.RandTestBogusPrivateKey()
}

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusPublicKey() (TestBogusPublicKey, error) {
	return netutil.RandTestBogusPublicKey()
}

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusPrivateKeyOrFatal(t *testing.T) TestBogusPrivateKey {
	return netutil.RandTestBogusPrivateKeyOrFatal(t)
}

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusPublicKeyOrFatal(t *testing.T) TestBogusPublicKey {
	return netutil.RandTestBogusPublicKeyOrFatal(t)
}

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusIdentity() (tnet.Identity, error) {
	return netutil.RandTestBogusIdentity()
}

// Deprecated: use the corresponding function in github.com/libp2p/go-libp2p-testing/netutil.
func RandTestBogusIdentityOrFatal(t *testing.T) tnet.Identity {
	return netutil.RandTestBogusIdentityOrFatal(t)
}
