package iosrestore

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/tmc/apple/x/plist"
)

const asrChunkSize = 128 << 10
const asrMessageLimit = 1 << 20

// SendImage serves ASR validation requests, then sends size bytes from image.
// ASR uses unframed XML control messages and optional SHA-1 payload checksums.
// A nil error means the payload was written, not that the device restored it.
//
// The caller must exclusively own conn, bind it to the intended device, and close
// it after this call. SendImage sets and clears its deadlines but never closes it.
// Context cancellation interrupts network I/O; image must return promptly from
// ReadAt and remain unchanged throughout the transfer. No payload is retried
// after a write error, since the receiver may already have consumed part of it.
func SendImage(ctx context.Context, conn net.Conn, image io.ReaderAt, size int64) (result error) {
	if conn == nil || image == nil || size <= 0 {
		return fmt.Errorf("asr requires a connection and nonempty image")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	deadline, _ := ctx.Deadline()
	if err := conn.SetDeadline(deadline); err != nil {
		return fmt.Errorf("set asr deadline: %w", err)
	}
	finished := make(chan struct{})
	stop := context.AfterFunc(ctx, func() { conn.SetDeadline(time.Now()); close(finished) })
	defer func() {
		if !stop() {
			<-finished
		}
		if err := ctx.Err(); err != nil {
			result = err
		}
		result = errors.Join(result, conn.SetDeadline(time.Time{}))
	}()
	reader := bufio.NewReader(conn)
	message, err := readASR(reader)
	if err != nil {
		return fmt.Errorf("read asr initiate: %w", err)
	}
	if message["Command"] != "Initiate" {
		return fmt.Errorf("expected asr initiate")
	}
	checksums := false
	for {
		switch message["Command"] {
		case "Initiate":
			if value, ok := message["Checksum Chunks"]; ok {
				checksums, ok = value.(bool)
				if !ok {
					return fmt.Errorf("invalid asr checksum negotiation")
				}
			}
			info := map[string]any{
				"FEC Slice Stride": 40, "Packet Payload Size": 1450, "Packets Per FEC": 25,
				"Payload": map[string]any{"Port": 1, "Size": size}, "Stream ID": 1, "Version": 1,
			}
			if checksums {
				info["Checksum Chunk Size"] = asrChunkSize
			}
			data, err := plist.Marshal(info, plist.FormatXML)
			if err != nil {
				return err
			}
			if err := writeAll(conn, data); err != nil {
				return fmt.Errorf("write asr packet info: %w", err)
			}
		case "OOBData":
			offset, ok := asrNumber(message["OOB Offset"])
			if !ok {
				return fmt.Errorf("invalid asr oob offset")
			}
			length, ok := asrNumber(message["OOB Length"])
			if !ok {
				return fmt.Errorf("invalid asr oob length")
			}
			if offset > uint64(size) || length > uint64(size)-offset {
				return fmt.Errorf("asr oob range exceeds image")
			}
			if err := sendASRRange(ctx, conn, image, int64(offset), int64(length), false); err != nil {
				return fmt.Errorf("send asr oob data: %w", err)
			}
		case "Payload":
			if err := sendASRRange(ctx, conn, image, 0, size, checksums); err != nil {
				return fmt.Errorf("send asr payload: %w", err)
			}
			return nil
		default:
			return fmt.Errorf("unexpected asr command %v", message["Command"])
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		message, err = readASR(reader)
		if err != nil {
			return fmt.Errorf("read asr command: %w", err)
		}
	}
}

func sendASRRange(ctx context.Context, w io.Writer, image io.ReaderAt, offset, length int64, checksums bool) error {
	buffer := make([]byte, asrChunkSize+sha1.Size)
	for length > 0 {
		if err := ctx.Err(); err != nil {
			return err
		}
		size := int(min(length, asrChunkSize))
		n, err := image.ReadAt(buffer[:size], offset)
		if n != size {
			if err == nil {
				err = io.ErrUnexpectedEOF
			}
			return err
		}
		if err != nil && err != io.EOF {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		data := buffer[:size]
		if checksums {
			digest := sha1.Sum(data) // ASR wire checksum, not a content-authentication policy.
			copy(buffer[size:], digest[:])
			data = buffer[:size+sha1.Size]
		}
		if err := writeAll(w, data); err != nil {
			return err
		}
		offset += int64(size)
		length -= int64(size)
	}
	return nil
}

func asrNumber(value any) (uint64, bool) {
	switch n := value.(type) {
	case uint64:
		return n, true
	case int64:
		if n >= 0 {
			return uint64(n), true
		}
	}
	return 0, false
}

// Giving encoding/xml a ByteReader prevents it from reading into the next plist.
// The shared buffered reader retains any bytes coalesced by the transport.
type asrXMLReader struct {
	r    io.ByteReader
	data []byte
}

func (r *asrXMLReader) ReadByte() (byte, error) {
	if len(r.data) >= asrMessageLimit {
		return 0, fmt.Errorf("asr control message exceeds limit")
	}
	b, err := r.r.ReadByte()
	if err == nil {
		r.data = append(r.data, b)
	}
	return b, err
}
func (r *asrXMLReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	p[0] = b
	return 1, nil
}
func readASR(reader io.ByteReader) (map[string]any, error) {
	r := &asrXMLReader{r: reader}
	decoder := xml.NewDecoder(r)
	depth := 0
	values := 0
	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if depth == 0 && (t.Name.Local != "plist" || t.Name.Space != "") {
				return nil, fmt.Errorf("asr message is not an XML plist")
			}
			if depth == 1 {
				values++
				if values > 1 || t.Name.Local != "dict" {
					return nil, fmt.Errorf("asr plist must contain one dictionary")
				}
			}
			depth++
			if depth > 64 {
				return nil, fmt.Errorf("asr XML nesting exceeds limit")
			}
		case xml.EndElement:
			depth--
			if depth == 0 {
				if values != 1 {
					return nil, fmt.Errorf("asr plist dictionary is missing")
				}
				b, err := r.ReadByte()
				if err != nil {
					return nil, err
				}
				if b == '\r' {
					b, err = r.ReadByte()
					if err != nil {
						return nil, err
					}
				}
				if b != '\n' {
					return nil, fmt.Errorf("asr plist is missing its newline terminator")
				}
				return dictionary(r.data)
			}
		case xml.CharData:
			if depth == 0 && strings.TrimSpace(string(t)) != "" {
				return nil, fmt.Errorf("unexpected data before asr plist")
			}
		}
	}
}
