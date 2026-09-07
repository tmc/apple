package iosrestore

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"
)

// ErrNotRestored means the peer identified itself as another service.
var ErrNotRestored = errors.New("peer is not the restored service")

// Info describes a live restored service. ECID is the hardware UniqueChipID,
// independent of usbmux's transient attachment number and serial string.
type Info struct {
	ECID            uint64         `json:"ecid"`
	ProtocolVersion uint64         `json:"protocolVersion"`
	Hardware        map[string]any `json:"hardware"`
}

// QueryInfo queries the service type, protocol version and hardware identity.
// The caller exclusively owns conn. This call sets and clears its deadlines,
// does not close it, and requires the caller to discard it after an I/O error.
func QueryInfo(ctx context.Context, conn net.Conn) (info Info, result error) {
	if conn == nil {
		return info, fmt.Errorf("restore connection is required")
	}
	if err := ctx.Err(); err != nil {
		return info, err
	}
	deadline, _ := ctx.Deadline()
	if err := conn.SetDeadline(deadline); err != nil {
		return info, err
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
	if err := Send(conn, map[string]any{"Request": "QueryType", "Label": "tmc/apple"}); err != nil {
		return info, err
	}
	response, err := Receive(conn)
	if err != nil {
		return info, err
	}
	service, ok := response["Type"].(string)
	if !ok || service == "" {
		return info, fmt.Errorf("restore service type is missing or invalid")
	}
	if service != "com.apple.mobile.restored" {
		return info, ErrNotRestored
	}
	version, ok := asrNumber(response["RestoreProtocolVersion"])
	if !ok || version == 0 {
		return info, fmt.Errorf("invalid restore protocol version")
	}
	if err := Send(conn, map[string]any{"Request": "QueryValue", "QueryKey": "HardwareInfo", "Label": "tmc/apple"}); err != nil {
		return info, err
	}
	response, err = Receive(conn)
	if err != nil {
		return info, err
	}
	hardware, ok := response["HardwareInfo"].(map[string]any)
	if !ok {
		return info, fmt.Errorf("restore hardware information is missing")
	}
	// Older plist producers can represent the same 64 ECID bits as a signed value.
	var ecid uint64
	switch n := hardware["UniqueChipID"].(type) {
	case int64:
		ecid = uint64(n)
	case uint64:
		ecid = n
	default:
		return info, fmt.Errorf("invalid restore ECID")
	}
	if ecid == 0 {
		return info, fmt.Errorf("restore ECID is zero")
	}
	return Info{ECID: ecid, ProtocolVersion: version, Hardware: hardware}, nil
}
