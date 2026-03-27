//go:build darwin

package ane

import (
	"errors"
	"testing"
)

func TestConnectDaemonClose(t *testing.T) {
	d, err := ConnectDaemon()
	if errors.Is(err, ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()

	// Double close should be safe.
	if err := d.Close(); err != nil {
		t.Errorf("second close returned error: %v", err)
	}
}

func TestDaemonClosedErrors(t *testing.T) {
	d, err := ConnectDaemon()
	if errors.Is(err, ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	d.Close()

	if err := d.CompileModel(nil, nil); !errors.Is(err, ErrDaemonClosed) {
		t.Errorf("CompileModel after close: got %v, want ErrDaemonClosed", err)
	}
	if err := d.LoadModel(nil, nil); !errors.Is(err, ErrDaemonClosed) {
		t.Errorf("LoadModel after close: got %v, want ErrDaemonClosed", err)
	}
	if err := d.PurgeModelHash("test"); !errors.Is(err, ErrDaemonClosed) {
		t.Errorf("PurgeModelHash after close: got %v, want ErrDaemonClosed", err)
	}
}

func TestDaemonOptions(t *testing.T) {
	tests := []struct {
		name string
		opts *DaemonOptions
		want uint32
	}{
		{"nil", nil, 21},
		{"zero", &DaemonOptions{}, 21},
		{"custom", &DaemonOptions{QoS: 42}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.opts.qos()
			if got != tt.want {
				t.Errorf("qos() = %d, want %d", got, tt.want)
			}
		})
	}
}
