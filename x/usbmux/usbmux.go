package usbmux

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/tmc/apple/x/plist"
)

// Client selects a usbmuxd endpoint. Empty Network and Address select unix and
// /var/run/usbmuxd. Use Network tcp with an explicit Address for a remote daemon.
type Client struct{ Network, Address string }

// Device describes a daemon attachment. ID is transient; Serial identifies the
// device. ConnectionType distinguishes USB from Network attachments.
type Device struct {
	ID             uint32 `plist:"DeviceID"`
	Serial         string `plist:"SerialNumber"`
	ConnectionType string `plist:"ConnectionType"`
}

// List returns currently attached devices, including network attachments.
func (c Client) List(ctx context.Context) ([]Device, error) {
	conn, err := c.open(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	stop := context.AfterFunc(ctx, func() { conn.Close() })
	defer stop()
	response, err := exchange(conn, map[string]any{"MessageType": "ListDevices"})
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if err != nil {
		return nil, err
	}
	var result struct {
		DeviceList []struct {
			ID         uint32 `plist:"DeviceID"`
			Properties Device `plist:"Properties"`
		} `plist:"DeviceList"`
	}
	if _, ok := response["DeviceList"]; !ok {
		return nil, fmt.Errorf("usbmux list: missing device list")
	}
	data, err := plist.Marshal(response, plist.FormatXML)
	if err != nil {
		return nil, err
	}
	if _, err := plist.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("usbmux list: %w", err)
	}
	devices := make([]Device, 0, len(result.DeviceList))
	for _, entry := range result.DeviceList {
		d := entry.Properties
		d.ID = entry.ID
		if d.ID == 0 || d.Serial == "" || d.ConnectionType == "" {
			return nil, fmt.Errorf("usbmux list: incomplete device")
		}
		devices = append(devices, d)
	}
	return devices, nil
}

// Dial connects to a device's TCP port through usbmuxd. Context controls the
// handshake only. The caller owns the returned connection and its deadlines.
func (c Client) Dial(ctx context.Context, deviceID uint32, port uint16) (net.Conn, error) {
	if deviceID == 0 || port == 0 {
		return nil, fmt.Errorf("usbmux connect: invalid device or port")
	}
	conn, err := c.open(ctx)
	if err != nil {
		return nil, err
	}
	finished := make(chan struct{})
	stop := context.AfterFunc(ctx, func() { conn.Close(); close(finished) })
	response, err := exchange(conn, map[string]any{"MessageType": "Connect", "DeviceID": deviceID, "PortNumber": port<<8 | port>>8})
	if !stop() {
		<-finished
	}
	if ctx.Err() != nil {
		err = ctx.Err()
	}
	if err == nil {
		err = checkResult(response)
	}
	if err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

func (c Client) open(ctx context.Context) (net.Conn, error) {
	network, address := c.Network, c.Address
	if network == "" {
		network = "unix"
	}
	if address == "" {
		if network != "unix" {
			return nil, fmt.Errorf("usbmux: address required for %s", network)
		}
		address = "/var/run/usbmuxd"
	}
	conn, err := (&net.Dialer{}).DialContext(ctx, network, address)
	if err != nil {
		return nil, fmt.Errorf("connect usbmuxd: %w", err)
	}
	return conn, nil
}

const maxMessage = 16 << 20

func exchange(conn net.Conn, request map[string]any) (map[string]any, error) {
	request["ClientVersionString"] = "tmc/apple"
	request["ProgName"] = "usbmux"
	request["kLibUSBMuxVersion"] = 3
	data, err := plist.Marshal(request, plist.FormatXML)
	if err != nil {
		return nil, err
	}
	packet := make([]byte, 16+len(data))
	binary.LittleEndian.PutUint32(packet, uint32(len(packet)))
	binary.LittleEndian.PutUint32(packet[4:], 1)
	binary.LittleEndian.PutUint32(packet[8:], 8)
	binary.LittleEndian.PutUint32(packet[12:], 1)
	copy(packet[16:], data)
	for len(packet) > 0 {
		n, e := conn.Write(packet)
		if e != nil {
			return nil, fmt.Errorf("write usbmux request: %w", e)
		}
		if n == 0 {
			return nil, io.ErrShortWrite
		}
		packet = packet[n:]
	}
	var header [16]byte
	if _, err := io.ReadFull(conn, header[:]); err != nil {
		return nil, fmt.Errorf("read usbmux header: %w", err)
	}
	size := binary.LittleEndian.Uint32(header[:])
	if size < 16 || size > maxMessage || binary.LittleEndian.Uint32(header[4:]) != 1 || binary.LittleEndian.Uint32(header[8:]) != 8 || binary.LittleEndian.Uint32(header[12:]) != 1 {
		return nil, fmt.Errorf("invalid usbmux response header")
	}
	data = make([]byte, int(size)-16)
	if _, err := io.ReadFull(conn, data); err != nil {
		return nil, fmt.Errorf("read usbmux response: %w", err)
	}
	value, err := plist.ParseBytes(data)
	if err != nil {
		return nil, fmt.Errorf("parse usbmux response: %w", err)
	}
	response, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("usbmux response is not a dictionary")
	}
	if response["MessageType"] == "Result" {
		if err := checkResult(response); err != nil {
			return nil, err
		}
	}
	return response, nil
}

func checkResult(response map[string]any) error {
	if response["MessageType"] != "Result" {
		return fmt.Errorf("usbmux: expected result response")
	}
	number, ok := response["Number"].(int64)
	if !ok {
		return fmt.Errorf("usbmux: missing result number")
	}
	if number != 0 {
		return fmt.Errorf("usbmux: daemon returned error %d", number)
	}
	return nil
}
