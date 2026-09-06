package jaccl

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	protocolMagic        = "apple/jaccl"
	protocolVersion      = 1
	maxControlFrameBytes = 8 << 10
)

type frameKind uint8

const (
	frameHello frameKind = iota + 1
	frameDestination
	frameReady
	frameOperation
	frameCredit
	frameAbort
)

type controlFrame struct {
	Magic       string    `json:"magic"`
	Version     int       `json:"version"`
	Kind        frameKind `json:"kind"`
	GroupID     string    `json:"group_id"`
	Incarnation uint64    `json:"incarnation"`
	Source      int       `json:"source"`
	Destination int       `json:"destination"`
	Epoch       uint32    `json:"epoch"`
	Payload     []byte    `json:"payload,omitempty"`
}

// readControlFrameContext closes conn if cancellation wins. Control readers are
// serialized by nativeBackend, so closing the connection cannot interrupt an
// unrelated frame; it deliberately makes the group unusable rather than
// leaving a caller blocked after its context was canceled.
func readControlFrameContext(ctx context.Context, conn net.Conn) (controlFrame, error) {
	if conn == nil {
		return controlFrame{}, fmt.Errorf("read control frame: %w", ErrClosed)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	type result struct {
		frame controlFrame
		err   error
	}
	resultc := make(chan result, 1)
	go func() {
		frame, err := readControlFrame(conn)
		resultc <- result{frame: frame, err: err}
	}()
	select {
	case result := <-resultc:
		return result.frame, result.err
	case <-ctx.Done():
		_ = conn.Close()
		return controlFrame{}, ctx.Err()
	}
}

func writeControlFrame(w io.Writer, f controlFrame) error {
	if err := f.validate(); err != nil {
		return err
	}
	payload, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("marshal control frame: %w", err)
	}
	if len(payload) > maxControlFrameBytes {
		return fmt.Errorf("control frame length %d: %w", len(payload), ErrProtocol)
	}
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(len(payload)))
	if err := writeFull(w, header[:]); err != nil {
		return fmt.Errorf("write control frame header: %w", err)
	}
	if err := writeFull(w, payload); err != nil {
		return fmt.Errorf("write control frame payload: %w", err)
	}
	return nil
}

func readControlFrame(r io.Reader) (controlFrame, error) {
	var header [4]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return controlFrame{}, fmt.Errorf("read control frame header: %w", err)
	}
	n := binary.BigEndian.Uint32(header[:])
	if n > maxControlFrameBytes {
		return controlFrame{}, fmt.Errorf("control frame length %d: %w", n, ErrProtocol)
	}
	payload := make([]byte, n)
	if _, err := io.ReadFull(r, payload); err != nil {
		return controlFrame{}, fmt.Errorf("read control frame payload: %w", err)
	}
	var f controlFrame
	if err := json.Unmarshal(payload, &f); err != nil {
		return controlFrame{}, fmt.Errorf("decode control frame: %w", errors.Join(ErrProtocol, err))
	}
	if err := f.validate(); err != nil {
		return controlFrame{}, err
	}
	return f, nil
}

func (f controlFrame) validate() error {
	if f.Magic != protocolMagic {
		return fmt.Errorf("control frame magic %q: %w", f.Magic, ErrProtocol)
	}
	if f.Version != protocolVersion {
		return fmt.Errorf("control frame version %d: %w", f.Version, ErrProtocol)
	}
	if f.Kind < frameHello || f.Kind > frameAbort {
		return fmt.Errorf("control frame kind %d: %w", f.Kind, ErrProtocol)
	}
	if f.GroupID == "" {
		return fmt.Errorf("control frame group id is empty: %w", ErrProtocol)
	}
	if f.Source < 0 || f.Destination < 0 || f.Source == f.Destination {
		return fmt.Errorf("control frame source=%d destination=%d: %w", f.Source, f.Destination, ErrProtocol)
	}
	if len(f.Payload) > maxControlFrameBytes {
		return fmt.Errorf("control frame payload length %d: %w", len(f.Payload), ErrProtocol)
	}
	return nil
}

func writeFull(w io.Writer, p []byte) error {
	for len(p) > 0 {
		n, err := w.Write(p)
		if n > 0 {
			p = p[n:]
		}
		if err != nil {
			return err
		}
		if n == 0 {
			return io.ErrShortWrite
		}
	}
	return nil
}
