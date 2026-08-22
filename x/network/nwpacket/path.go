//go:build darwin

package nwpacket

import (
	"errors"
	"fmt"
	"net"
	"strings"

	applenetwork "github.com/tmc/apple/network"
	"github.com/tmc/apple/objc"
)

// PathReporter reports the Network.framework path for a peer connection.
type PathReporter interface {
	PeerPath(net.Addr) (Path, error)
}

// Path describes the Network.framework path observed for a peer connection.
type Path struct {
	Status     applenetwork.NWPathStatus
	Interfaces []PathInterface
}

// PathInterface describes one interface in a Network.framework path.
type PathInterface struct {
	Name  string
	Index uint32
	Type  applenetwork.NWInterfaceType
}

// UsesInterface reports whether p includes name.
func (p Path) UsesInterface(name string) bool {
	for _, iface := range p.Interfaces {
		if iface.Name == name {
			return true
		}
	}
	return false
}

// InterfaceNames returns the path interface names in Network.framework order.
func (p Path) InterfaceNames() []string {
	names := make([]string, 0, len(p.Interfaces))
	for _, iface := range p.Interfaces {
		names = append(names, iface.Name)
	}
	return names
}

func (p Path) String() string {
	if len(p.Interfaces) == 0 {
		return fmt.Sprintf("status=%s interfaces=none", p.Status)
	}
	parts := make([]string, 0, len(p.Interfaces))
	for _, iface := range p.Interfaces {
		parts = append(parts, iface.String())
	}
	return fmt.Sprintf("status=%s interfaces=%s", p.Status, strings.Join(parts, ","))
}

func (i PathInterface) String() string {
	if i.Index == 0 {
		return fmt.Sprintf("%s/%s", i.Name, i.Type)
	}
	return fmt.Sprintf("%s/%s(%d)", i.Name, i.Type, i.Index)
}

// PeerPath returns the current Network.framework path for addr.
//
// PeerPath reports only existing peer connections. Callers typically call it
// after WriteTo or ReadFrom has established a peer.
func (c *nwPacketConn) PeerPath(addr net.Addr) (Path, error) {
	peer, err := c.existingPeerConn(addr)
	if err != nil {
		return Path{}, err
	}
	path := applenetwork.NWConnectionCopyCurrentPath(peer.conn)
	return pathFromNWPath(path)
}

func (c *nwPacketConn) existingPeerConn(addr net.Addr) (*nwPeerConn, error) {
	if addr == nil {
		return nil, errors.New("missing peer address")
	}
	udpAddr, err := netAddrToUDP(addr)
	if err != nil {
		return nil, err
	}
	if udpAddr.Zone == "" && udpAddr.IP.To4() == nil && udpAddr.IP.IsLinkLocalUnicast() {
		udpAddr.Zone = c.config.InterfaceName
	}
	key := udpAddr.String()
	c.mu.Lock()
	defer c.mu.Unlock()
	peer := c.conns[key]
	if peer == nil {
		return nil, fmt.Errorf("no peer connection for %s", key)
	}
	return peer, nil
}

func pathFromNWPath(path applenetwork.NWPath) (Path, error) {
	if path.ID == 0 {
		return Path{}, errors.New("nw path is unavailable")
	}
	result := Path{Status: applenetwork.NWPathGetStatus(path)}
	applenetwork.NWPathEnumerateInterfaces(path, func(iface applenetwork.NWInterface) bool {
		if iface.ID == 0 {
			return true
		}
		result.Interfaces = append(result.Interfaces, PathInterface{
			Name:  objc.GoString(applenetwork.NWInterfaceGetName(iface)),
			Index: applenetwork.NWInterfaceGetIndex(iface),
			Type:  applenetwork.NWInterfaceGetType(iface),
		})
		return true
	})
	return result, nil
}
