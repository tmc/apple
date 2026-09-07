//go:build darwin

package irecovery

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
)

// SendCommand sends one NUL-terminated iBoot command in recovery mode. Request
// is 0 for normal commands or 1 for the go command used after an iBEC upload.
// A transfer failure or short write leaves device state uncertain; do not replay
// automatically. Successful transmission does not prove command execution.
func (c *Conn) SendCommand(ctx context.Context, command string, request uint8) error {
	if err := checkCommand(command, request); err != nil {
		return err
	}
	if err := c.lock(ctx); err != nil {
		return err
	}
	defer c.mu.Unlock()
	return c.sendCommand(ctx, command, request)
}
func checkCommand(command string, request uint8) error {
	if len(command) == 0 || len(command) >= 256 || strings.ContainsAny(command, "\x00\r\n") {
		return fmt.Errorf("recovery command must contain 1–255 bytes without NUL or newlines")
	}
	if request > 1 {
		return fmt.Errorf("unsupported recovery command request %d", request)
	}
	return nil
}
func (c *Conn) sendCommand(ctx context.Context, command string, request uint8) error {
	if c.library == nil {
		return fmt.Errorf("recovery connection is closed")
	}
	if c.info.Mode != "recovery" {
		return fmt.Errorf("commands require recovery mode")
	}
	data := append([]byte(command), 0)
	n, err := c.control(ctx, 0x40, request, 0, 0, data)
	if err != nil {
		return fmt.Errorf("send recovery command: %w", err)
	}
	if n != len(data) {
		return fmt.Errorf("send recovery command: %w", io.ErrShortWrite)
	}
	return nil
}

// Getenv sends getenv and reads its bounded response while holding the connection
// lock. It returns an error on transfer failure or possible response truncation.
// A trailing NUL and zero padding are removed; other bytes are preserved.
func (c *Conn) Getenv(ctx context.Context, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("recovery variable name is required")
	}
	for _, ch := range name {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || strings.ContainsRune("._-", ch)) {
			return "", fmt.Errorf("invalid recovery variable name")
		}
	}
	command := "getenv " + name
	if err := checkCommand(command, 0); err != nil {
		return "", err
	}
	if err := c.lock(ctx); err != nil {
		return "", err
	}
	defer c.mu.Unlock()
	if err := c.sendCommand(ctx, command, 0); err != nil {
		return "", err
	}
	var data [255]byte
	n, err := c.control(ctx, 0xc0, 0, 0, 0, data[:])
	if err != nil {
		return "", fmt.Errorf("read recovery variable: %w", err)
	}
	if n < 0 || n > len(data) {
		return "", fmt.Errorf("invalid recovery variable byte count %d", n)
	}
	response := data[:n]
	if end := bytes.IndexByte(response, 0); end >= 0 {
		for _, b := range response[end:] {
			if b != 0 {
				return "", fmt.Errorf("nonzero data after recovery variable terminator")
			}
		}
		response = response[:end]
	} else if n == len(data) {
		return "", fmt.Errorf("recovery variable response may be truncated")
	}
	return string(response), nil
}
