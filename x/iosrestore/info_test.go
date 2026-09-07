package iosrestore

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestQueryInfo(t *testing.T) {
	for _, ecid := range []any{uint64(42), uint64(1)<<63 | 17, int64(-1)} {
		t.Run(fmt.Sprint(ecid), func(t *testing.T) {
			host, peer := net.Pipe()
			defer host.Close()
			defer peer.Close()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			done := make(chan error, 1)
			go func() {
				request, err := Receive(peer)
				if err != nil {
					done <- err
					return
				}
				if request["Request"] != "QueryType" {
					done <- fmt.Errorf("wrong query: %v", request)
					return
				}
				if err := Send(peer, map[string]any{"Type": "com.apple.mobile.restored", "RestoreProtocolVersion": 15}); err != nil {
					done <- err
					return
				}
				request, err = Receive(peer)
				if err != nil {
					done <- err
					return
				}
				if request["Request"] != "QueryValue" || request["QueryKey"] != "HardwareInfo" {
					done <- fmt.Errorf("wrong hardware query: %v", request)
					return
				}
				done <- Send(peer, map[string]any{"HardwareInfo": map[string]any{"UniqueChipID": ecid}})
			}()
			info, err := QueryInfo(ctx, host)
			if err != nil {
				t.Fatal(err)
			}
			var want uint64
			switch n := ecid.(type) {
			case uint64:
				want = n
			case int64:
				want = uint64(n)
			}
			if info.ECID != want || info.ProtocolVersion != 15 {
				t.Fatal(info)
			}
			if err := <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestQueryInfoRejectsOtherService(t *testing.T) {
	host, peer := net.Pipe()
	defer host.Close()
	defer peer.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() { Receive(peer); Send(peer, map[string]any{"Type": "com.apple.mobile.lockdown"}) }()
	_, err := QueryInfo(ctx, host)
	if !errors.Is(err, ErrNotRestored) {
		t.Fatal(err)
	}
}

func ExampleQueryInfo() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	host, peer := net.Pipe()
	defer host.Close()
	defer peer.Close()
	_, err := QueryInfo(ctx, host)
	fmt.Println(errors.Is(err, context.Canceled))
	// Output: true
}

func TestQueryInfoMalformedType(t *testing.T) {
	host, peer := net.Pipe()
	defer host.Close()
	defer peer.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() { Receive(peer); Send(peer, map[string]any{"Error": "InvalidRequest"}) }()
	_, err := QueryInfo(ctx, host)
	if err == nil || errors.Is(err, ErrNotRestored) {
		t.Fatal("silently skipped malformed peer", err)
	}
}
