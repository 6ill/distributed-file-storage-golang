package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddr: ":4000",
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.TCPTransportOpts.ListenAddr, tcpOpts.ListenAddr)

	// Server
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
