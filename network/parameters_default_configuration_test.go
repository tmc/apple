package network

import "testing"

func TestParametersCreateSecureTCPNilUsesDefaultConfiguration(t *testing.T) {
	params := NWParametersCreateSecureTCP(nil, nil)
	if params.ID == 0 {
		t.Fatal("NWParametersCreateSecureTCP(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}

func TestParametersCreateSecureUDPNilUsesDefaultConfiguration(t *testing.T) {
	params := NWParametersCreateSecureUDP(nil, nil)
	if params.ID == 0 {
		t.Fatal("NWParametersCreateSecureUDP(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}
