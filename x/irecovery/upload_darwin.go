//go:build darwin

package irecovery

import (
	"context"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"time"
)

// StatusError reports the device's DFU status and state bytes.
type StatusError struct{ Status, State byte }

func (e *StatusError) Error() string {
	return fmt.Sprintf("dfu status %d in state %d", e.Status, e.State)
}

// Upload transfers an image to the selected DFU or recovery endpoint. It holds
// exclusive access to the connection for the entire transfer and fails on short
// writes without retrying them. The caller must not modify data during the call.
//
// In DFU mode, Upload appends the Apple suffix and CRC and waits for download-idle
// after each block. A non-idle initial state is cleared or aborted, then reported
// as an error requiring an explicit retry. Manifestation notification, USB reset and matching-device
// reconnection are separate steps; success here only establishes image transfer.
// In recovery mode, Upload sends bulk data and any required terminating ZLP.
// It does not execute the uploaded image. The operation has a two-minute ceiling
// or the caller's earlier deadline. Physical-device compatibility is unqualified.
func (c *Conn) Upload(ctx context.Context, data []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if len(data) == 0 {
		return fmt.Errorf("empty recovery image")
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	if err := c.lock(ctx); err != nil {
		return err
	}
	defer c.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.library == nil {
		return fmt.Errorf("recovery connection is closed")
	}
	c.dfuBlocks = 0
	switch c.info.Mode {
	case "recovery":
		return c.uploadRecovery(ctx, data)
	case "dfu":
		// The later manifestation notification needs the next 16-bit block ID.
		if len(data) > 0xffff*0x800 {
			return fmt.Errorf("dfu image exceeds transfer limit")
		}
		if err := c.uploadDFU(ctx, data); err != nil {
			return err
		}
		c.dfuBlocks = uint16((len(data) + 0x7ff) / 0x800)
		return nil
	default:
		return fmt.Errorf("unsupported upload mode %q", c.info.Mode)
	}
}

func (c *Conn) uploadRecovery(ctx context.Context, data []byte) error {
	if err := c.controlExact(ctx, 0x41, 0, 0, nil); err != nil {
		return fmt.Errorf("begin recovery upload: %w", err)
	}
	for offset := 0; offset < len(data); offset += 0x8000 {
		chunk := data[offset:min(offset+0x8000, len(data))]
		n, err := c.bulkWrite(ctx, 4, chunk)
		if err != nil {
			return fmt.Errorf("upload recovery offset %d: %w", offset, err)
		}
		if n != len(chunk) {
			return fmt.Errorf("upload recovery offset %d: %w", offset, io.ErrShortWrite)
		}
	}
	if len(data)%512 == 0 {
		n, err := c.bulkWrite(ctx, 4, nil)
		if err != nil {
			return fmt.Errorf("terminate recovery upload: %w", err)
		}
		if n != 0 {
			return fmt.Errorf("terminate recovery upload: unexpected byte count %d", n)
		}
	}
	return nil
}

func (c *Conn) uploadDFU(ctx context.Context, data []byte) error {
	var state [1]byte
	if err := c.controlExact(ctx, 0xa1, 5, 0, state[:]); err != nil {
		return fmt.Errorf("read dfu state: %w", err)
	}
	if state[0] != 2 {
		request := byte(6)
		if state[0] == 10 {
			request = 4
		}
		if err := c.controlExact(ctx, 0x21, request, 0, nil); err != nil {
			return fmt.Errorf("recover initial dfu state %d: %w", state[0], err)
		}
		return fmt.Errorf("recovered initial dfu state %d; retry upload explicitly", state[0])
	}
	footer := dfuFooter(data)
	for offset, block := 0, uint16(0); offset < len(data); offset, block = offset+0x800, block+1 {
		chunk := data[offset:min(offset+0x800, len(data))]
		if offset+len(chunk) == len(data) {
			if len(chunk)+len(footer) > 0x800 {
				// Apple recovery sends a separate suffix with the same final block ID.
				if err := c.controlExact(ctx, 0x21, 1, block, chunk); err != nil {
					return fmt.Errorf("upload dfu block %d: %w", block, err)
				}
				chunk = nil
			}
			packet := make([]byte, 0, len(chunk)+len(footer))
			packet = append(packet, chunk...)
			packet = append(packet, footer[:]...)
			chunk = packet
		}
		if err := c.controlExact(ctx, 0x21, 1, block, chunk); err != nil {
			return fmt.Errorf("upload dfu block %d: %w", block, err)
		}
		if err := c.waitDownload(ctx); err != nil {
			return fmt.Errorf("complete dfu block %d: %w", block, err)
		}
	}
	return nil
}

func dfuFooter(data []byte) [16]byte {
	footer := [16]byte{0xff, 0xff, 0xff, 0xff, 0xac, 0x05, 0x00, 0x01, 0x55, 0x46, 0x44, 0x10}
	crc := crc32.Update(crc32.ChecksumIEEE(data), crc32.IEEETable, footer[:12])
	binary.LittleEndian.PutUint32(footer[12:], ^crc)
	return footer
}

func (c *Conn) controlExact(ctx context.Context, kind, request byte, block uint16, data []byte) error {
	n, err := c.control(ctx, kind, request, block, 0, data)
	if err != nil {
		return err
	}
	if n != len(data) {
		if kind&0x80 != 0 {
			return io.ErrUnexpectedEOF
		}
		return io.ErrShortWrite
	}
	return nil
}
func (c *Conn) waitDownload(ctx context.Context) error {
	for {
		var response [6]byte
		if err := c.controlExact(ctx, 0xa1, 3, 0, response[:]); err != nil {
			return err
		}
		status, state := response[0], response[4]
		if status != 0 || state != 3 && state != 4 && state != 5 {
			return &StatusError{Status: status, State: state}
		}
		delay := time.Duration(uint32(response[1])|uint32(response[2])<<8|uint32(response[3])<<16) * time.Millisecond
		if delay > 0 || state != 5 {
			timer := time.NewTimer(max(delay, time.Millisecond))
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
		}
		if state == 5 {
			return nil
		}
	}
}
