package network

import "testing"

func TestParametersCreateSecureTCPNilUsesDefaultConfiguration(t *testing.T) {
	params, err := TryNw_parameters_create_secure_tcp(nil, nil)
	if err != nil {
		t.Fatalf("TryNw_parameters_create_secure_tcp(nil, nil): %v", err)
	}
	if params.ID == 0 {
		t.Fatal("TryNw_parameters_create_secure_tcp(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}

func TestParametersCreateSecureUDPNilUsesDefaultConfiguration(t *testing.T) {
	params, err := TryNw_parameters_create_secure_udp(nil, nil)
	if err != nil {
		t.Fatalf("TryNw_parameters_create_secure_udp(nil, nil): %v", err)
	}
	if params.ID == 0 {
		t.Fatal("TryNw_parameters_create_secure_udp(nil, nil) returned nil")
	}
	t.Cleanup(params.Release)
}
