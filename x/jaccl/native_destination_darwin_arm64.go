//go:build darwin && arm64

package jaccl

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/tmc/apple/rdma"
	xrdma "github.com/tmc/apple/x/rdma"
)

const maxPSN = 1<<24 - 1

// nativeDestination is the fixed route record exchanged over the control
// plane before a QP transitions to RTR.
type nativeDestination struct {
	QPN       uint32
	PSN       uint32
	LID       uint16
	GIDIndex  int
	GID       rdma.IbvGID
	ActiveMTU int32
	LinkLayer uint8
}

func (l *nativeLink) localDestination() (nativeDestination, error) {
	if l == nil || l.device == nil || l.qp == 0 {
		return nativeDestination{}, fmt.Errorf("local destination: %w", ErrProtocol)
	}
	psn, err := randomPSN()
	if err != nil {
		return nativeDestination{}, err
	}
	d := nativeDestination{
		QPN:       rdma.Ibv_qp_num(l.qp),
		PSN:       psn,
		LID:       l.device.port.LID,
		GIDIndex:  l.device.route.Index,
		GID:       l.device.route.GID,
		ActiveMTU: l.device.port.ActiveMTU,
		LinkLayer: l.device.port.LinkLayer,
	}
	if err := d.validate(); err != nil {
		return nativeDestination{}, err
	}
	return d, nil
}

func randomPSN() (uint32, error) {
	var data [4]byte
	if _, err := cryptorand.Read(data[:]); err != nil {
		return 0, fmt.Errorf("generate PSN: %w", err)
	}
	psn := binary.LittleEndian.Uint32(data[:]) & maxPSN
	if psn == 0 {
		psn = 1
	}
	return psn, nil
}

func (d nativeDestination) validate() error {
	if d.QPN == 0 {
		return fmt.Errorf("destination QPN is zero: %w", ErrProtocol)
	}
	if d.PSN > maxPSN {
		return fmt.Errorf("destination PSN %d: %w", d.PSN, ErrProtocol)
	}
	if d.LID == 0 && xrdma.IsZeroGID(d.GID) {
		return fmt.Errorf("destination LID and GID are zero: %w", ErrProtocol)
	}
	if !xrdma.IsZeroGID(d.GID) && (d.GIDIndex < 0 || d.GIDIndex > 255) {
		return fmt.Errorf("destination GID index %d: %w", d.GIDIndex, ErrProtocol)
	}
	if xrdma.MTUBytes(d.ActiveMTU) == 0 {
		return fmt.Errorf("destination active MTU %d: %w", d.ActiveMTU, ErrProtocol)
	}
	return nil
}

func encodeNativeDestination(d nativeDestination) ([]byte, error) {
	if err := d.validate(); err != nil {
		return nil, err
	}
	data, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("encode native destination: %w", err)
	}
	if len(data) > maxControlFrameBytes {
		return nil, fmt.Errorf("native destination length %d: %w", len(data), ErrProtocol)
	}
	return data, nil
}

func decodeNativeDestination(data []byte) (nativeDestination, error) {
	var destination nativeDestination
	if len(data) == 0 || len(data) > maxControlFrameBytes {
		return nativeDestination{}, fmt.Errorf("native destination length %d: %w", len(data), ErrProtocol)
	}
	if err := json.Unmarshal(data, &destination); err != nil {
		return nativeDestination{}, fmt.Errorf("decode native destination: %w", err)
	}
	if err := destination.validate(); err != nil {
		return nativeDestination{}, err
	}
	return destination, nil
}

func (l *nativeLink) initialize() error {
	if l == nil || l.device == nil || l.qp == 0 {
		return fmt.Errorf("initialize queue pair: %w", ErrProtocol)
	}
	attr := rdma.IbvQPAttr{
		QPState:       rdma.IBV_QPS_INIT,
		PortNum:       l.device.portNum,
		PKeyIndex:     0,
		QPAccessFlags: rdma.IBV_ACCESS_LOCAL_WRITE,
	}
	mask := rdma.IBV_QP_STATE | rdma.IBV_QP_PKEY_INDEX | rdma.IBV_QP_PORT | rdma.IBV_QP_ACCESS_FLAGS
	return l.modify("INIT", &attr, mask)
}

func (l *nativeLink) connect(ctx context.Context, local, remote nativeDestination) error {
	if err := local.validate(); err != nil {
		return fmt.Errorf("validate local destination: %w", err)
	}
	if err := remote.validate(); err != nil {
		return fmt.Errorf("validate remote destination: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("transition queue pair to RTR: %w", err)
	}
	attr, mask, err := nativeRTRAttr(l.device.portNum, local, remote)
	if err != nil {
		return fmt.Errorf("build RTR attributes: %w", err)
	}
	if err := l.modify("RTR", &attr, mask); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("transition queue pair to RTS: %w", err)
	}
	rts := rdma.IbvQPAttr{QPState: rdma.IBV_QPS_RTS, SQPSN: local.PSN}
	return l.modify("RTS", &rts, rdma.IBV_QP_STATE|rdma.IBV_QP_SQ_PSN)
}

func nativeRTRAttr(port uint8, local, remote nativeDestination) (rdma.IbvQPAttr, int, error) {
	return xrdma.RTRAttr(xrdma.LocalQP{
		PortNum:   port,
		GIDIndex:  local.GIDIndex,
		ActiveMTU: local.ActiveMTU,
		LinkLayer: local.LinkLayer,
	}, xrdma.RemoteQP{
		LID:       remote.LID,
		QPN:       remote.QPN,
		PSN:       remote.PSN,
		GIDIndex:  remote.GIDIndex,
		GID:       remote.GID,
		UseGlobal: !xrdma.IsZeroGID(remote.GID),
		ActiveMTU: remote.ActiveMTU,
	}, xrdma.RTRPolicy{MaxPathMTU: rdma.IBV_MTU_1024})
}

func (l *nativeLink) modify(state string, attr *rdma.IbvQPAttr, mask int) error {
	rc, err := rdma.IbvModifyQpAttr(l.qp, attr, mask)
	if err := nativeProviderRC("transition queue pair to "+state, rc, err); err != nil {
		return err
	}
	return nil
}
