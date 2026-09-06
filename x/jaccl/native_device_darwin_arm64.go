//go:build darwin && arm64

package jaccl

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/rdma"
	xrdma "github.com/tmc/apple/x/rdma"
)

// nativeDevice owns the portion of the verbs chain shared by every peer QP on
// one local device. The device list stays open until the context closes because
// its handles are provider-owned pointers.
type nativeDevice struct {
	list    *rdma.DeviceList
	name    string
	context rdma.RDMAContext
	pd      rdma.RDMAPD
	portNum uint8
	port    rdma.IbvPortAttr
	route   xrdma.RouteGID

	cleanup closeStack
}

func openNativeDevice(cfg Config) (*nativeDevice, error) {
	if !rdma.Available() {
		return nil, fmt.Errorf("Apple RDMA provider: %w", ErrUnsupported)
	}
	list, err := rdma.OpenDeviceList()
	if err != nil {
		return nil, fmt.Errorf("list RDMA devices: %w", err)
	}
	devices := list.Devices()
	if cfg.Device != "" {
		devices = namedDevice(devices, cfg.Device)
	}
	if len(devices) == 0 {
		_ = list.Close()
		return nil, fmt.Errorf("RDMA device %q: %w", cfg.Device, ErrUnsupported)
	}

	var candidates []*nativeDevice
	var candidateErrs []error
	for _, device := range devices {
		candidate, err := openNativeDeviceCandidate(device, cfg.Port)
		if err != nil {
			candidateErrs = append(candidateErrs, fmt.Errorf("device %q: %w", device.Name, err))
			continue
		}
		candidates = append(candidates, candidate)
	}
	if len(candidates) != 1 {
		for _, candidate := range candidates {
			_ = candidate.Close()
		}
		_ = list.Close()
		if len(candidates) == 0 {
			return nil, fmt.Errorf("no safe RDMA device: %w", errors.Join(append(candidateErrs, ErrUnsupported)...))
		}
		return nil, fmt.Errorf("%d safe RDMA devices; set Config.Device: %w", len(candidates), ErrInvalidConfig)
	}
	candidates[0].list = list
	return candidates[0], nil
}

func namedDevice(devices []rdma.Device, name string) []rdma.Device {
	for _, device := range devices {
		if device.Name == name {
			return []rdma.Device{device}
		}
	}
	return nil
}

func openNativeDeviceCandidate(device rdma.Device, portNumber uint8) (*nativeDevice, error) {
	context, err := device.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	if context == 0 {
		return nil, fmt.Errorf("open: zero context: %w", ErrUnsupported)
	}
	var stack closeStack
	stack.add(func() error {
		rc, err := rdma.IbvCloseDevice(context)
		return nativeProviderRC("ibv_close_device", int(rc), err)
	})

	var port rdma.IbvPortAttr
	if rc, err := rdma.IbvQueryPortAttr(context, portNumber, &port); err != nil || rc != 0 {
		_ = stack.close()
		return nil, nativeProviderRC("ibv_query_port", rc, err)
	}
	if port.State != rdma.IBV_PORT_ACTIVE {
		_ = stack.close()
		return nil, fmt.Errorf("port %d state %d: %w", portNumber, port.State, ErrUnsupported)
	}
	route, err := queryRouteGID(context, portNumber, port)
	if err != nil {
		_ = stack.close()
		return nil, err
	}
	pd, err := rdma.IbvAllocPd(context)
	if err != nil {
		_ = stack.close()
		return nil, fmt.Errorf("allocate protection domain: %w", err)
	}
	if pd == 0 {
		_ = stack.close()
		return nil, fmt.Errorf("allocate protection domain: %w", ErrUnsupported)
	}
	stack.add(func() error {
		rc, err := rdma.IbvDeallocPd(pd)
		return nativeProviderRC("ibv_dealloc_pd", int(rc), err)
	})
	return &nativeDevice{name: device.Name, context: context, pd: pd, portNum: portNumber, port: port, route: route, cleanup: stack}, nil
}

func queryRouteGID(context rdma.RDMAContext, portNumber uint8, port rdma.IbvPortAttr) (xrdma.RouteGID, error) {
	limit := xrdma.RouteGIDScanLimit(int32(port.GIDTblLen))
	gids := make([]xrdma.RouteGID, 0, limit)
	for index := 0; index < limit; index++ {
		var gid rdma.IbvGID
		rc, err := rdma.IbvQueryGidInto(context, portNumber, index, &gid)
		if err != nil || rc != 0 {
			return xrdma.RouteGID{}, nativeProviderRC(fmt.Sprintf("ibv_query_gid[%d]", index), rc, err)
		}
		gids = append(gids, xrdma.RouteGID{Index: index, GID: gid})
	}
	route, ok := xrdma.SelectRouteGID(gids, -1, port.LinkLayer)
	if !ok {
		return xrdma.RouteGID{}, fmt.Errorf("port %d has no safe route GID: %w", portNumber, ErrUnsupported)
	}
	return route, nil
}

// Close releases the shared device resources in the required reverse order.
func (d *nativeDevice) Close() error {
	if d == nil {
		return nil
	}
	err := d.cleanup.close()
	if d.list != nil {
		err = errors.Join(err, d.list.Close())
		d.list = nil
	}
	d.context = 0
	d.pd = 0
	d.portNum = 0
	d.port = rdma.IbvPortAttr{}
	d.route = xrdma.RouteGID{}
	return err
}

func nativeProviderRC(name string, rc int, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if rc != 0 {
		return fmt.Errorf("%s: rc=%d", name, rc)
	}
	return nil
}
