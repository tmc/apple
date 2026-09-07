//go:build darwin

package irecovery

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestParseSerial(t *testing.T) {
	for _, tt := range []struct {
		name, serial string
		want         uint64
		fail         bool
	}{
		{"research", "CPID:FE01 ECID:1234 SRTG:[iBoot]", 0x1234, false},
		{"large", "ECID:FFFFFFFFFFFFFFFF", ^uint64(0), false},
		{"missing", "SRTG:[iBoot]", 0, false},
		{"malformed", "ECID:xyz", 0, true},
		{"duplicate", "ECID:1 ECID:2", 0, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			d, err := parseSerial(tt.serial)
			if (err != nil) != tt.fail || (!tt.fail && d.ECID != tt.want) {
				t.Fatal(d, err)
			}
		})
	}
}
func TestTransferTimeout(t *testing.T) {
	if n, err := transferTimeout(context.Background()); err != nil || n != 1000 {
		t.Fatal(n, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if n, err := transferTimeout(ctx); err != nil || n == 0 || n > 10 {
		t.Fatal(n, err)
	}
}
func TestClosed(t *testing.T) {
	var c Conn
	if _, err := c.Control(context.Background(), 0, 0, 0, 0, nil); err == nil {
		t.Fatal("closed control succeeded")
	}
	if _, err := c.BulkWrite(context.Background(), 4, nil); err == nil {
		t.Fatal("closed bulk succeeded")
	}
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}
func ExampleDiscover() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Discover(ctx, "libusb.dylib")
	fmt.Println(err)
	// Output: context canceled
}
func ExampleOpen() {
	_, err := Open(context.Background(), "libusb.dylib", 0)
	fmt.Println(err)
	// Output: nonzero ECID is required
}
func ExampleConn_Close() {
	var conn Conn
	fmt.Println(conn.Close())
	// Output: <nil>
}
func ExampleConn_Control() {
	var conn Conn
	_, err := conn.Control(context.Background(), 0xA1, 3, 0, 0, make([]byte, 6))
	fmt.Println(err)
	// Output: recovery connection is closed
}
func ExampleConn_BulkWrite() {
	var conn Conn
	_, err := conn.BulkWrite(context.Background(), 4, []byte("payload"))
	fmt.Println(err)
	// Output: recovery connection is closed
}
