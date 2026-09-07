package usbmux

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/tmc/apple/x/plist"
)

func daemon(t *testing.T, serve func(net.Conn)) Client {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { listener.Close() })
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(3 * time.Second))
		serve(conn)
	}()
	return Client{Network: "tcp", Address: listener.Addr().String()}
}
func request(conn net.Conn) (map[string]any, error) {
	var h [16]byte
	if _, err := io.ReadFull(conn, h[:]); err != nil {
		return nil, err
	}
	n := binary.LittleEndian.Uint32(h[:])
	if n < 16 || n > maxMessage {
		return nil, fmt.Errorf("bad size")
	}
	b := make([]byte, n-16)
	if _, err := io.ReadFull(conn, b); err != nil {
		return nil, err
	}
	v, err := plist.ParseBytes(b)
	if err != nil {
		return nil, err
	}
	return v.(map[string]any), nil
}
func reply(conn net.Conn, value any) {
	b, _ := plist.Marshal(value, plist.FormatXML)
	h := make([]byte, 16)
	binary.LittleEndian.PutUint32(h, uint32(16+len(b)))
	binary.LittleEndian.PutUint32(h[4:], 1)
	binary.LittleEndian.PutUint32(h[8:], 8)
	binary.LittleEndian.PutUint32(h[12:], 1)
	conn.Write(append(h, b...))
}
func TestList(t *testing.T) {
	c := daemon(t, func(conn net.Conn) {
		r, err := request(conn)
		if err != nil || r["MessageType"] != "ListDevices" {
			t.Error(r, err)
			return
		}
		reply(conn, map[string]any{"DeviceList": []any{map[string]any{"DeviceID": 7, "Properties": map[string]any{"SerialNumber": "abc", "ConnectionType": "USB"}}}})
	})
	d, err := c.List(context.Background())
	if err != nil || len(d) != 1 || d[0].ID != 7 || d[0].Serial != "abc" {
		t.Fatal(d, err)
	}
}
func TestDialStreamAndPort(t *testing.T) {
	c := daemon(t, func(conn net.Conn) {
		r, err := request(conn)
		if err != nil || r["PortNumber"] != int64(0x7ef2) {
			t.Error(r, err)
			return
		}
		reply(conn, map[string]any{"MessageType": "Result", "Number": 0})
		conn.Write([]byte("hello"))
	})
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := c.Dial(ctx, 7, 62078)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	cancel()
	b := make([]byte, 5)
	if _, err := io.ReadFull(conn, b); err != nil || string(b) != "hello" {
		t.Fatal(string(b), err)
	}
}
func TestCancellation(t *testing.T) {
	c := daemon(t, func(conn net.Conn) { io.Copy(io.Discard, conn) })
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if _, err := c.List(ctx); err != context.DeadlineExceeded {
		t.Fatal(err)
	}
}
func TestInvalidHeaders(t *testing.T) {
	for _, size := range []uint32{0, 15, maxMessage + 1} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			c := daemon(t, func(conn net.Conn) {
				request(conn)
				h := make([]byte, 16)
				binary.LittleEndian.PutUint32(h, size)
				conn.Write(h)
			})
			if _, err := c.List(context.Background()); err == nil {
				t.Fatal("accepted invalid header")
			}
		})
	}
}
func ExampleClient_List() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := (Client{}).List(ctx)
	fmt.Println(err != nil)
	// Output: true
}
func ExampleClient_Dial() {
	_, err := (Client{}).Dial(context.Background(), 0, 0)
	fmt.Println(err)
	// Output: usbmux connect: invalid device or port
}
