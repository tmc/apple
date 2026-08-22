package mach

import (
	"strings"
	"testing"
	"time"
)

// TestPortRightDispositions is the leak test design/x-mach.md requires:
// send a right with each disposition and assert the sender's right count
// afterward, in both polarities — CopySend must keep the sender's right,
// MoveSend must consume it.
func TestPortRightDispositions(t *testing.T) {
	recv, err := NewPort()
	if err != nil {
		t.Fatal(err)
	}
	defer recv.DestroyReceive()

	for _, tt := range []struct {
		name          string
		disp          Disposition
		wantAfterSend int // sender's send-right refs on the payload after send
	}{
		{"CopySend keeps the sender's right", CopySend, 1},
		{"MoveSend consumes the sender's right", MoveSend, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := NewPort()
			if err != nil {
				t.Fatal(err)
			}
			defer payload.DestroyReceive()
			if err := payload.MakeSendRight(); err != nil {
				t.Fatal(err)
			}
			if refs, _ := payload.Refs(RightSend); refs != 1 {
				t.Fatalf("precondition: send refs = %d, want 1", refs)
			}

			if err := Send(recv, MakeSend, 42, []PortRight{{payload, tt.disp}}, []byte("tensor"), time.Second); err != nil {
				t.Fatal(err)
			}
			refs, err := payload.Refs(RightSend)
			if err != nil {
				t.Fatal(err)
			}
			if refs != tt.wantAfterSend {
				t.Errorf("after send with %v: sender's send refs = %d, want %d", tt.disp, refs, tt.wantAfterSend)
			}

			m, err := Receive(recv, time.Second)
			if err != nil {
				t.Fatal(err)
			}
			if m.Header.ID != 42 {
				t.Errorf("id = %d, want 42", m.Header.ID)
			}
			if !strings.HasPrefix(string(m.Body), "tensor") {
				t.Errorf("body = %q, want \"tensor\"", m.Body)
			}
			if len(m.Ports) != 1 {
				t.Fatalf("got %d ports, want 1", len(m.Ports))
			}
			// The received right is ours to balance.
			if err := m.Ports[0].Deallocate(); err != nil {
				t.Errorf("deallocate received right: %v", err)
			}
			if tt.disp == CopySend {
				if err := payload.Deallocate(); err != nil {
					t.Errorf("deallocate kept right: %v", err)
				}
			}
		})
	}
}

func TestBootstrapRoundTrip(t *testing.T) {
	const name = "com.tmc.apple.x.mach.test"
	p, err := NewPort()
	if err != nil {
		t.Fatal(err)
	}
	defer p.DestroyReceive()
	if err := p.MakeSendRight(); err != nil {
		t.Fatal(err)
	}
	defer p.Deallocate()

	if err := BootstrapRegister(name, p); err != nil {
		t.Fatal(err)
	}
	got, err := BootstrapLookUp(name)
	if err != nil {
		t.Fatal(err)
	}
	defer got.Deallocate()

	// Prove the looked-up right reaches our receive right.
	if err := Send(got, CopySend, 7, nil, []byte("hi"), time.Second); err != nil {
		t.Fatal(err)
	}
	m, err := Receive(p, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if m.Header.ID != 7 || !strings.HasPrefix(string(m.Body), "hi") {
		t.Errorf("got id=%d body=%q, want id=7 body=\"hi\"", m.Header.ID, m.Body)
	}
}
