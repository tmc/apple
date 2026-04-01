package network

import "testing"

func TestParametersCreateSecureTCPNilUsesDefaultConfiguration(t *testing.T) {
	params := Nw_parameters_create_secure_tcp(nil, nil)
	if params.ID == 0 {
		t.Fatal("Nw_parameters_create_secure_tcp(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}

func TestParametersCreateSecureUDPNilUsesDefaultConfiguration(t *testing.T) {
	params := Nw_parameters_create_secure_udp(nil, nil)
	if params.ID == 0 {
		t.Fatal("Nw_parameters_create_secure_udp(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}
