package network

import (
	"testing"
	"time"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
)

func TestConnectionStateChangedHandlerExposesTypedError(t *testing.T) {
	endpoint := Nw_endpoint_create_host("127.0.0.1", "1")
	if endpoint.ID == 0 {
		t.Fatal("Nw_endpoint_create_host returned nil")
	}
	t.Cleanup(endpoint.Release)

	params := Nw_parameters_create_secure_tcp(nil, nil)
	if params.ID == 0 {
		t.Fatal("Nw_parameters_create_secure_tcp returned nil")
	}
	t.Cleanup(params.Release)

	connection := Nw_connection_create(endpoint, params)
	if connection.ID == 0 {
		t.Fatal("Nw_connection_create returned nil")
	}
	t.Cleanup(connection.Release)
	t.Cleanup(func() { Nw_connection_cancel(connection) })

	errc := make(chan Nw_error_t, 1)
	Nw_connection_set_queue(connection, dispatch.GetGlobalQueue(dispatch.QOSDefault))
	Nw_connection_set_state_changed_handler(connection, func(state NwConnectionState, nwErr Nw_error_t) {
		if nwErr.IsZero() {
			return
		}
		switch state {
		case Nw_connection_state_waiting, Nw_connection_state_failed:
			select {
			case errc <- nwErr:
			default:
			}
		}
	})

	Nw_connection_start(connection)

	select {
	case nwErr := <-errc:
		if got := nwErr.Error(); got == "" {
			t.Fatal("Nw_error_t.Error() returned empty string")
		}
		if got := nwErr.Code(); got == 0 {
			t.Fatal("Nw_error_t.Code() returned 0")
		}
		if got := nwErr.Domain(); got == Nw_error_domain_invalid {
			t.Fatalf("Nw_error_t.Domain() = %v, want a real domain", got)
		}
		if got := nwErr.DomainString(); got == "" {
			t.Fatal("Nw_error_t.DomainString() returned empty string")
		}
		cfErr := nwErr.CopyCFError()
		if cfErr == 0 {
			t.Fatal("Nw_error_t.CopyCFError() returned nil")
		}
		corefoundation.CFRelease(corefoundation.CFTypeRef(cfErr))
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for a real network error")
	}
}
