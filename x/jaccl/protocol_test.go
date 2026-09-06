package jaccl

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestControlFrameRoundTrip(t *testing.T) {
	want := controlFrame{
		Magic:       protocolMagic,
		Version:     protocolVersion,
		Kind:        frameDestination,
		GroupID:     "group",
		Incarnation: 7,
		Source:      0,
		Destination: 1,
		Epoch:       2,
		Payload:     []byte("payload"),
	}
	var buf bytes.Buffer
	if err := writeControlFrame(&buf, want); err != nil {
		t.Fatal(err)
	}
	got, err := readControlFrame(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("frame = %#v, want %#v", got, want)
	}
}

func TestReadControlFrameContextCancels(t *testing.T) {
	left, right := net.Pipe()
	defer right.Close()
	ctx, cancel := context.WithCancel(context.Background())
	returned := make(chan error, 1)
	go func() {
		_, err := readControlFrameContext(ctx, left)
		returned <- err
	}()
	cancel()
	select {
	case err := <-returned:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("read error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("canceled read did not return")
	}
}

func TestControlFrameRejectsOversize(t *testing.T) {
	var buf bytes.Buffer
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], maxControlFrameBytes+1)
	buf.Write(header[:])
	if _, err := readControlFrame(&buf); !errors.Is(err, ErrProtocol) {
		t.Fatalf("read error = %v, want ErrProtocol", err)
	}
}

func TestControlFrameRejectsIdentityMismatch(t *testing.T) {
	f := controlFrame{Magic: "wrong", Version: protocolVersion, Kind: frameHello, GroupID: "group", Source: 0, Destination: 1}
	if err := f.validate(); !errors.Is(err, ErrProtocol) {
		t.Fatalf("validate error = %v, want ErrProtocol", err)
	}
}
