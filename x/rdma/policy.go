package rdma

import bindings "github.com/tmc/apple/rdma"

const (
	// PortActive is the ibv_port_state value for PORT_ACTIVE.
	PortActive int32 = 4

	// LinkLayerThunderbolt is Apple's ibv_port_attr link_layer value for
	// Thunderbolt RDMA.
	LinkLayerThunderbolt uint8 = 100

	// MaxRouteGIDScan bounds automatic route-GID discovery. Apple Thunderbolt
	// devices can report large gid_tbl_len values while only exposing a small
	// useful prefix.
	MaxRouteGIDScan = 8
)

// GID is an RDMA global identifier.
type GID = bindings.IbvGID

// RouteGID is a nonzero GID candidate returned by ibv_query_gid.
type RouteGID struct {
	Index int
	GID   GID
}

// GIDInfo is the JSON-friendly form used by diagnostics.
type GIDInfo struct {
	Index      int
	Raw        string
	IPv4Mapped bool
}

// IsZeroGID reports whether gid is all zeros.
func IsZeroGID(gid GID) bool {
	for _, b := range gid {
		if b != 0 {
			return false
		}
	}
	return true
}

// IsIPv4MappedGID reports whether gid has the IPv4-mapped IPv6 prefix.
func IsIPv4MappedGID(gid GID) bool {
	for i := 0; i < 10; i++ {
		if gid[i] != 0 {
			return false
		}
	}
	return gid[10] == 0xff && gid[11] == 0xff
}

// SelectRouteGID selects a route GID for QP RTR setup.
//
// If preferred is non-negative, it is treated as an explicit diagnostic
// override and must match a nonzero candidate. Automatic selection rejects
// Thunderbolt index 0, prefers IPv4-mapped GIDs, then accepts index 1. Other
// link layers may fall back to the first nonzero candidate.
func SelectRouteGID(gids []RouteGID, preferred int, linkLayer uint8) (RouteGID, bool) {
	if preferred >= 0 {
		for _, entry := range gids {
			if entry.Index == preferred && !IsZeroGID(entry.GID) {
				return entry, true
			}
		}
		return RouteGID{Index: -1}, false
	}
	for _, entry := range gids {
		if linkLayer == LinkLayerThunderbolt && entry.Index == 0 {
			continue
		}
		if IsIPv4MappedGID(entry.GID) {
			return entry, true
		}
	}
	for _, entry := range gids {
		if entry.Index == 1 && !IsZeroGID(entry.GID) {
			return entry, true
		}
	}
	if linkLayer == LinkLayerThunderbolt {
		return RouteGID{Index: -1}, false
	}
	for _, entry := range gids {
		if !IsZeroGID(entry.GID) {
			return entry, true
		}
	}
	return RouteGID{Index: -1}, false
}

// SelectRouteGIDInfo applies the automatic route-GID policy to diagnostic
// records.
func SelectRouteGIDInfo(gids []GIDInfo, linkLayer uint8) (GIDInfo, bool) {
	for _, gid := range gids {
		if !nonzeroGIDInfo(gid) {
			continue
		}
		if linkLayer == LinkLayerThunderbolt && gid.Index == 0 {
			continue
		}
		if gid.IPv4Mapped {
			return gid, true
		}
	}
	for _, gid := range gids {
		if gid.Index == 1 && nonzeroGIDInfo(gid) {
			return gid, true
		}
	}
	if linkLayer == LinkLayerThunderbolt {
		return GIDInfo{Index: -1}, false
	}
	for _, gid := range gids {
		if nonzeroGIDInfo(gid) {
			return gid, true
		}
	}
	return GIDInfo{Index: -1}, false
}

func nonzeroGIDInfo(gid GIDInfo) bool {
	for _, r := range gid.Raw {
		if r != '0' {
			return true
		}
	}
	return false
}

// RouteGIDScanLimit returns the bounded automatic GID scan length.
func RouteGIDScanLimit(tableLen int32) int {
	limit := MaxRouteGIDScan
	if tableLen > 0 && int(tableLen) < limit {
		limit = int(tableLen)
	}
	return limit
}

// PreflightGIDScanLimit returns the bounded diagnostic GID scan length.
func PreflightGIDScanLimit(tableLen int32, requested int) int {
	if requested < 0 {
		requested = 0
	}
	if requested > MaxRouteGIDScan {
		requested = MaxRouteGIDScan
	}
	if tableLen > 0 && requested > int(tableLen) {
		requested = int(tableLen)
	}
	return requested
}
