package main

import (
	"net"
	"testing"
	"time"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"1", 1},
		{"64K", 64 * 1024},
		{"1M", 1024 * 1024},
	}
	for _, tt := range tests {
		if got := parseSize(tt.in); got != tt.want {
			t.Fatalf("parseSize(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestPatternRoundTrip(t *testing.T) {
	for _, pattern := range []string{"stream", "pingpong", "duplex"} {
		if got := patternName(patternID(pattern)); got != pattern {
			t.Fatalf("patternName(patternID(%q)) = %q", pattern, got)
		}
	}
}

func TestSummarizeLatency(t *testing.T) {
	got := summarizeLatency([]time.Duration{
		10 * time.Microsecond,
		20 * time.Microsecond,
		30 * time.Microsecond,
		40 * time.Microsecond,
	})
	if got == nil {
		t.Fatal("summarizeLatency returned nil")
	}
	if got.Count != 4 {
		t.Fatalf("Count = %d, want 4", got.Count)
	}
	if got.Min != "10µs" || got.P50 != "20µs" || got.Max != "40µs" {
		t.Fatalf("summary = %#v", got)
	}
}

func TestRDMAReadinessNames(t *testing.T) {
	if got := rdmaNetInterface("rdma_en2"); got != "en2" {
		t.Fatalf("rdmaNetInterface = %q, want en2", got)
	}
	if got := portStateName(1); got != "PORT_DOWN" {
		t.Fatalf("portStateName = %q, want PORT_DOWN", got)
	}
	if got := linkLayerName(100); got != "Thunderbolt" {
		t.Fatalf("linkLayerName = %q, want Thunderbolt", got)
	}
	if got := mtuBytes(5); got != 4096 {
		t.Fatalf("mtuBytes = %d, want 4096", got)
	}
}

func TestRunTCPLocalhost(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 3; i++ {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			serveConn(c, true)
		}
	}()

	stream := runTCP(ln.Addr().String(), "stream", 1024, 50*time.Millisecond)
	if stream.Error != "" {
		t.Fatalf("stream error: %v", stream.Error)
	}
	if stream.Bytes == 0 || stream.BytesPerSec == 0 {
		t.Fatalf("stream result missing bytes/rate: %#v", stream)
	}

	ping := runTCP(ln.Addr().String(), "pingpong", 64, 50*time.Millisecond)
	if ping.Error != "" {
		t.Fatalf("pingpong error: %v", ping.Error)
	}
	if ping.Latency == nil || ping.Latency.Count == 0 {
		t.Fatalf("pingpong missing latency: %#v", ping)
	}

	duplex := runTCP(ln.Addr().String(), "duplex", 1024, 50*time.Millisecond)
	if duplex.Error != "" {
		t.Fatalf("duplex error: %v", duplex.Error)
	}
	if duplex.Bytes == 0 || duplex.BytesPerSec == 0 {
		t.Fatalf("duplex result missing bytes/rate: %#v", duplex)
	}

	<-done
}
