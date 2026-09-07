//go:build darwin

package irecovery

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Info is a fresh observation from an open DFU or recovery connection. Nil
// optional fields were absent; they must not be interpreted as zero. CPFM bit 0
// is security mode and bit 1 is production mode; IBFL bit 2 indicates Image4.
// USB descriptors report identity but do not cryptographically attest it.
type Info struct {
	Device   Device
	BoardID  *uint32
	CPFM     *uint32
	IBFL     *uint32
	APNonce  []byte
	SEPNonce []byte
}

// ReadInfo rereads the serial and nonce string descriptors on this connection,
// checking ECID and chip identity against Open's selection. It reads NONC/SNON
// together from string descriptor 1. The caller must refresh observations after
// a boot transition or reconnect. Returned bytes do not alias connection state.
// ReadInfo neither sends recovery commands nor invalidates a pending DFU upload.
func (c *Conn) ReadInfo(ctx context.Context) (Info, error) {
	if err := c.lock(ctx); err != nil {
		return Info{}, err
	}
	defer c.mu.Unlock()
	if c.library == nil {
		return Info{}, fmt.Errorf("recovery connection is closed")
	}
	if c.info.serialIndex == 0 {
		return Info{}, fmt.Errorf("device has no serial descriptor")
	}
	langs, err := c.stringDescriptor(ctx, 0, 0)
	if err != nil {
		return Info{}, err
	}
	if len(langs) < 2 {
		return Info{}, fmt.Errorf("USB language descriptor is empty")
	}
	lang := binary.LittleEndian.Uint16(langs)
	serial, err := c.asciiDescriptor(ctx, c.info.serialIndex, lang)
	if err != nil {
		return Info{}, err
	}
	device, err := parseSerial(serial)
	if err != nil {
		return Info{}, err
	}
	if device.ECID == 0 || device.ECID != c.info.ECID || device.CPID == 0 || (c.info.CPID != 0 && device.CPID != c.info.CPID) {
		return Info{}, fmt.Errorf("USB identity changed since device selection")
	}
	device.ProductID, device.Mode, device.serialIndex = c.info.ProductID, c.info.Mode, c.info.serialIndex
	info, err := parseInfo(device)
	if err != nil {
		return Info{}, err
	}
	nonces := serial
	if c.info.serialIndex != 1 {
		nonces, err = c.asciiDescriptor(ctx, 1, lang)
		if err != nil {
			return Info{}, err
		}
	}
	info.APNonce, info.SEPNonce, err = parseNonces(nonces)
	if err != nil {
		return Info{}, err
	}
	if err := ctx.Err(); err != nil {
		return Info{}, err
	}
	return info, nil
}

func (c *Conn) stringDescriptor(ctx context.Context, index uint8, lang uint16) ([]byte, error) {
	var data [255]byte
	n, err := c.control(ctx, 0x80, 6, 0x0300|uint16(index), lang, data[:])
	if err != nil {
		return nil, fmt.Errorf("read USB string descriptor %d: %w", index, err)
	}
	if n < 2 || n > len(data) || int(data[0]) != n || data[1] != 3 || n%2 != 0 {
		return nil, fmt.Errorf("invalid USB string descriptor %d", index)
	}
	return data[2:n], nil
}
func (c *Conn) asciiDescriptor(ctx context.Context, index uint8, lang uint16) (string, error) {
	data, err := c.stringDescriptor(ctx, index, lang)
	if err != nil {
		return "", err
	}
	text := make([]byte, len(data)/2)
	for i := range text {
		ch := binary.LittleEndian.Uint16(data[2*i:])
		if ch == 0 || ch > 127 {
			return "", fmt.Errorf("USB identity descriptor %d is not ASCII", index)
		}
		text[i] = byte(ch)
	}
	return string(text), nil
}
func parseInfo(device Device) (Info, error) {
	info := Info{Device: device}
	for _, field := range strings.Fields(device.Serial) {
		key, value, ok := strings.Cut(field, ":")
		if !ok {
			continue
		}
		var target **uint32
		switch key {
		case "BDID":
			target = &info.BoardID
		case "CPFM":
			target = &info.CPFM
		case "IBFL":
			target = &info.IBFL
		default:
			continue
		}
		if *target != nil {
			return Info{}, fmt.Errorf("duplicate USB serial field %s", key)
		}
		n, err := strconv.ParseUint(value, 16, 32)
		if err != nil {
			return Info{}, fmt.Errorf("parse USB %s: %w", key, err)
		}
		number := uint32(n)
		*target = &number
	}
	return info, nil
}
func parseNonces(text string) (ap, sep []byte, err error) {
	for _, field := range strings.Fields(text) {
		key, value, ok := strings.Cut(field, ":")
		if !ok {
			continue
		}
		var target *[]byte
		switch key {
		case "NONC":
			target = &ap
		case "SNON":
			target = &sep
		default:
			continue
		}
		if *target != nil {
			return nil, nil, fmt.Errorf("duplicate USB nonce field %s", key)
		}
		if value == "" {
			return nil, nil, fmt.Errorf("empty USB nonce field %s", key)
		}
		*target, err = hex.DecodeString(value)
		if err != nil {
			return nil, nil, fmt.Errorf("parse USB %s: %w", key, err)
		}
	}
	return ap, sep, nil
}
