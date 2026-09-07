package iosrestore

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tmc/apple/x/plist"
)

const maxMessage = 16 << 20

// Send writes a length-prefixed plist dictionary to a restored service stream.
// The caller serializes writes and configures transport deadlines.
func Send(w io.Writer, message map[string]any) error {
	data, err := plist.Marshal(message, plist.FormatXML)
	if err != nil {
		return err
	}
	if len(data) > maxMessage {
		return fmt.Errorf("restore message exceeds limit")
	}
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(len(data)))
	if err := writeAll(w, header[:]); err != nil {
		return err
	}
	return writeAll(w, data)
}

// Receive reads one restored service dictionary, bounded to 16 MiB. The caller
// must discard the stream after any framing or decoding error.
func Receive(r io.Reader) (map[string]any, error) {
	var header [4]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return nil, fmt.Errorf("read restore header: %w", err)
	}
	size := binary.BigEndian.Uint32(header[:])
	if size == 0 || size > maxMessage {
		return nil, fmt.Errorf("invalid restore message size %d", size)
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, fmt.Errorf("read restore message: %w", err)
	}
	return dictionary(data)
}

// Tickets sends a caller-prepared personalization request to Apple's TSS
// service. A nil client uses http.DefaultClient; an empty endpoint uses HTTPS.
// Requests must contain the correct build identity, nonces and device fields.
// This function checks the protocol status but cannot qualify those inputs.
func Tickets(ctx context.Context, client *http.Client, endpoint string, request map[string]any) (map[string]any, error) {
	if endpoint == "" {
		endpoint = "https://gs.apple.com/TSS/controller?action=2"
	}
	if client == nil {
		client = http.DefaultClient
	}
	data, err := plist.Marshal(request, plist.FormatXML)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request signing tickets: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("signing server HTTP status %d", response.StatusCode)
	}
	data, err = io.ReadAll(io.LimitReader(response.Body, maxMessage+1))
	if err != nil {
		return nil, err
	}
	if len(data) > maxMessage {
		return nil, fmt.Errorf("signing response exceeds limit")
	}
	prefix, payload, hasPayload := strings.Cut(string(data), "&REQUEST_STRING=")
	values, err := url.ParseQuery(prefix)
	if err != nil {
		return nil, fmt.Errorf("parse signing status: %w", err)
	}
	status := values["STATUS"]
	if len(status) != 1 {
		return nil, fmt.Errorf("signing response missing unique status")
	}
	code, err := strconv.Atoi(status[0])
	if err != nil {
		return nil, fmt.Errorf("invalid signing status")
	}
	if code != 0 {
		return nil, fmt.Errorf("signing server status %d: %s", code, values.Get("MESSAGE"))
	}
	if !hasPayload {
		return nil, fmt.Errorf("signing response missing ticket plist")
	}
	return dictionary([]byte(payload))
}

func dictionary(data []byte) (map[string]any, error) {
	value, err := plist.ParseBytes(data)
	if err != nil {
		return nil, fmt.Errorf("parse restore plist: %w", err)
	}
	result, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("restore plist is not a dictionary")
	}
	return result, nil
}
func writeAll(w io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		if n == 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
	}
	return nil
}
