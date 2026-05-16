package network

import (
	"testing"
	"time"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
)

func TestConnectionStateChangedHandlerExposesTypedError(t *testing.T) {
	endpoint := NWEndpointCreateHost("127.0.0.1", "1")
	if endpoint.ID == 0 {
		t.Fatal("NWEndpointCreateHost returned nil")
	}
	t.Cleanup(endpoint.Release)

	params := NWParametersCreateSecureTCP(nil, nil)
	if params.ID == 0 {
		t.Fatal("NWParametersCreateSecureTCP returned nil")
	}
	t.Cleanup(params.Release)

	connection := NWConnectionCreate(endpoint, params)
	if connection.ID == 0 {
		t.Fatal("NWConnectionCreate returned nil")
	}
	t.Cleanup(connection.Release)
	t.Cleanup(func() { NWConnectionCancel(connection) })

	errc := make(chan NWError, 1)
	NWConnectionSetQueue(connection, dispatch.GetGlobalQueue(dispatch.QOSDefault))
	NWConnectionSetStateChangedHandler(connection, func(state NWConnectionState, nwErr NWError) {
		if nwErr.IsZero() {
			return
		}
		switch state {
		case NWConnectionStateWaiting, NWConnectionStateFailed:
			select {
			case errc <- nwErr:
			default:
			}
		}
	})

	NWConnectionStart(connection)

	select {
	case nwErr := <-errc:
		if got := nwErr.Error(); got == "" {
			t.Fatal("NWError.Error() returned empty string")
		}
		if got := nwErr.Code(); got == 0 {
			t.Fatal("NWError.Code() returned 0")
		}
		if got := nwErr.Domain(); got == NWErrorDomainInvalid {
			t.Fatalf("NWError.Domain() = %v, want a real domain", got)
		}
		if got := nwErr.DomainString(); got == "" {
			t.Fatal("NWError.DomainString() returned empty string")
		}
		cfErr := nwErr.CopyCFError()
		if cfErr == 0 {
			t.Fatal("NWError.CopyCFError() returned nil")
		}
		corefoundation.CFRelease(corefoundation.CFTypeRef(cfErr))
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for a real network error")
	}
}
