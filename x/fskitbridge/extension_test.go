package fskitbridge

import (
	"errors"
	"testing"
)

// TestExtensionInitRetries covers the retryable-init contract: a failed Init
// is not sticky (the next call retries the factory) and LastError reports the
// most recent failure, but once a build succeeds the factory is not called
// again. A nil *Server is used as the built value; Init only checks for a
// non-nil error, so the factory's success path does not need a live class.
func TestExtensionInitRetries(t *testing.T) {
	wantErr := errors.New("class not registered")
	var calls int
	fail := true
	e := NewExtension(ReplyBlockShims{}, func() (*Server, error) {
		calls++
		if fail {
			return nil, wantErr
		}
		return &Server{}, nil
	})

	if err := e.Init(); !errors.Is(err, wantErr) {
		t.Fatalf("Init() error = %v, want %v", err, wantErr)
	}
	if got := e.LastError(); !errors.Is(got, wantErr) {
		t.Fatalf("LastError() = %v, want %v", got, wantErr)
	}
	if calls != 1 {
		t.Fatalf("factory calls = %d, want 1", calls)
	}

	// A second Init retries rather than returning the cached failure.
	if err := e.Init(); !errors.Is(err, wantErr) {
		t.Fatalf("retry Init() error = %v, want %v", err, wantErr)
	}
	if calls != 2 {
		t.Fatalf("factory calls after retry = %d, want 2", calls)
	}

	// Once the factory succeeds, Init caches the Server and stops calling it.
	fail = false
	if err := e.Init(); err != nil {
		t.Fatalf("Init() after recovery error = %v, want nil", err)
	}
	if e.LastError() != nil {
		t.Fatalf("LastError() after recovery = %v, want nil", e.LastError())
	}
	if err := e.Init(); err != nil {
		t.Fatalf("second Init() after recovery error = %v, want nil", err)
	}
	if calls != 3 {
		t.Fatalf("factory calls after recovery = %d, want 3 (built once)", calls)
	}
}
