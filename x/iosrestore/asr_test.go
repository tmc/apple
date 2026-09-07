package iosrestore

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

// Literal wire fixtures deliberately do not use the implementation's encoder.
func asrCommand(body string) string {
	return "<?xml version=\"1.0\"?><plist version=\"1.0\"><dict>" + body + "</dict></plist>\n"
}

const asrInitiate = "<key>Command</key><string>Initiate</string>"
const asrPayload = "<key>Command</key><string>Payload</string>"

func TestASRTransfer(t *testing.T) {
	for _, mode := range []string{"plain", "checksums", "renegotiate"} {
		t.Run(mode, func(t *testing.T) {
			host, device := net.Pipe()
			defer host.Close()
			defer device.Close()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			payload := make([]byte, 2*asrChunkSize+37)
			for i := range payload {
				payload[i] = byte(i * 17)
			}
			done := make(chan error, 1)
			go func() { done <- SendImage(ctx, host, bytes.NewReader(payload), int64(len(payload))) }()
			reader := bufio.NewReader(device)
			checksum := mode == "checksums"
			sendInfo := func(checksum bool) {
				wire := asrCommand(asrInitiate + fmt.Sprintf("<key>Checksum Chunks</key><%t/>", checksum))
				// Exercise arbitrary transport fragmentation.
				for i := range wire {
					if _, err := io.WriteString(device, wire[i:i+1]); err != nil {
						t.Fatal(err)
					}
				}
				info, err := readASR(reader)
				if err != nil {
					t.Fatal(err)
				}
				if info["Version"] != int64(1) || info["Stream ID"] != int64(1) || info["Packet Payload Size"] != int64(1450) || info["Packets Per FEC"] != int64(25) || info["FEC Slice Stride"] != int64(40) {
					t.Fatal(info)
				}
				p := info["Payload"].(map[string]any)
				if p["Size"] != int64(len(payload)) || p["Port"] != int64(1) {
					t.Fatal(p)
				}
				if checksum {
					if info["Checksum Chunk Size"] != int64(asrChunkSize) {
						t.Fatal(info)
					}
				} else if _, ok := info["Checksum Chunk Size"]; ok {
					t.Fatal(info)
				}
			}
			sendInfo(checksum)
			if mode == "renegotiate" {
				sendInfo(true)
				checksum = true
			}
			// Two coalesced requests prove buffered bytes survive plist boundaries.
			first := asrCommand("<key>Command</key><string>OOBData</string><key>OOB Offset</key><integer>131065</integer><key>OOB Length</key><integer>29</integer>")
			wire := first + asrCommand(asrPayload)
			if _, err := io.WriteString(device, wire); err != nil {
				t.Fatal(err)
			}
			oob := make([]byte, 29)
			if _, err := io.ReadFull(reader, oob); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(oob, payload[131065:131094]) {
				t.Fatal("wrong OOB range")
			}
			for offset := 0; offset < len(payload); offset += asrChunkSize {
				n := min(asrChunkSize, len(payload)-offset)
				got := make([]byte, n)
				if _, err := io.ReadFull(reader, got); err != nil {
					t.Fatal(err)
				}
				if !bytes.Equal(got, payload[offset:offset+n]) {
					t.Fatal("wrong payload")
				}
				if checksum {
					got := make([]byte, 20)
					if _, err := io.ReadFull(reader, got); err != nil {
						t.Fatal(err)
					}
					want := sha1.Sum(payload[offset : offset+n])
					if !bytes.Equal(got, want[:]) {
						t.Fatal("wrong checksum")
					}
				}
			}
			if err := <-done; err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestASRChecksumKnownVector(t *testing.T) {
	var out bytes.Buffer
	if err := sendASRRange(context.Background(), &out, strings.NewReader("abc"), 0, 3, true); err != nil {
		t.Fatal(err)
	}
	if got := hex.EncodeToString(out.Bytes()); got != "616263a9993e364706816aba3e25717850c26c9cd0d89d" {
		t.Fatal(got)
	}
}

func TestASRRejectCommands(t *testing.T) {
	for _, tt := range []struct{ name, body string }{
		{"negative", "<key>OOB Offset</key><integer>-1</integer><key>OOB Length</key><integer>1</integer>"},
		{"overflow", "<key>OOB Offset</key><integer>18446744073709551615</integer><key>OOB Length</key><integer>2</integer>"},
		{"past end", "<key>OOB Offset</key><integer>3</integer><key>OOB Length</key><integer>2</integer>"},
		{"wrong type", "<key>OOB Offset</key><string>0</string><key>OOB Length</key><integer>1</integer>"},
		{"missing", "<key>OOB Offset</key><integer>0</integer>"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			host, device := net.Pipe()
			defer host.Close()
			defer device.Close()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			done := make(chan error, 1)
			go func() { done <- SendImage(ctx, host, strings.NewReader("abcd"), 4) }()
			io.WriteString(device, asrCommand(asrInitiate))
			if _, err := readASR(bufio.NewReader(device)); err != nil {
				t.Fatal(err)
			}
			io.WriteString(device, asrCommand("<key>Command</key><string>OOBData</string>"+tt.body))
			if err := <-done; err == nil {
				t.Fatal("accepted invalid range")
			}
		})
	}
}

func TestASRCancellation(t *testing.T) {
	for _, phase := range []string{"initiate", "payload"} {
		t.Run(phase, func(t *testing.T) {
			host, device := net.Pipe()
			defer host.Close()
			defer device.Close()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			done := make(chan error, 1)
			go func() { done <- SendImage(ctx, host, strings.NewReader("abcdef"), 6) }()
			if phase == "payload" {
				io.WriteString(device, asrCommand(asrInitiate))
				readASR(bufio.NewReader(device))
				io.WriteString(device, asrCommand(asrPayload))
			}
			cancel()
			select {
			case err := <-done:
				if !errors.Is(err, context.Canceled) {
					t.Fatal(err)
				}
			case <-time.After(time.Second):
				t.Fatal("cancellation did not interrupt I/O")
			}
			// SendImage must neither close the connection nor leave a deadline callback behind.
			exchange := make(chan error, 1)
			go func() { _, err := host.Write([]byte{42}); exchange <- err }()
			var b [1]byte
			device.SetReadDeadline(time.Now().Add(time.Second))
			if _, err := io.ReadFull(device, b[:]); err != nil || b[0] != 42 {
				t.Fatal(b, err)
			}
			if err := <-exchange; err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestASRXMLBoundaries(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(asrCommand(asrInitiate) + asrCommand(asrPayload)))
	for _, want := range []string{"Initiate", "Payload"} {
		got, err := readASR(reader)
		if err != nil || got["Command"] != want {
			t.Fatal(got, err)
		}
	}
	for _, data := range []string{"<other/>", "<plist><dict>", strings.Repeat(" ", asrMessageLimit+1), "<plist>" + strings.Repeat("<array>", 65)} {
		if _, err := readASR(bufio.NewReader(strings.NewReader(data))); err == nil {
			t.Fatal("accepted malformed/oversized XML")
		}
	}
}

func TestASRShortImage(t *testing.T) {
	var out bytes.Buffer
	if err := sendASRRange(context.Background(), &out, strings.NewReader("short"), 0, 12, false); err == nil || out.Len() != 0 {
		t.Fatal(out.Len(), err)
	}
}

func ExampleSendImage() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	host, device := net.Pipe()
	defer host.Close()
	defer device.Close()
	err := SendImage(ctx, host, strings.NewReader("image"), 5)
	fmt.Println(errors.Is(err, context.Canceled))
	// Output: true
}

type asrShortWriter struct {
	data bytes.Buffer
	fail bool
}

func (w *asrShortWriter) Write(p []byte) (int, error) {
	n := min(len(p), 3)
	w.data.Write(p[:n])
	if w.fail {
		return n, io.ErrClosedPipe
	}
	return n, nil
}
func TestASRShortWrites(t *testing.T) {
	for _, fail := range []bool{false, true} {
		w := &asrShortWriter{fail: fail}
		err := sendASRRange(context.Background(), w, strings.NewReader("abcdefg"), 0, 7, false)
		if fail {
			if !errors.Is(err, io.ErrClosedPipe) || w.data.String() != "abc" {
				t.Fatal(w.data.String(), err)
			}
		} else if err != nil || w.data.String() != "abcdefg" {
			t.Fatal(w.data.String(), err)
		}
	}
}
