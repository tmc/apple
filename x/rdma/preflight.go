package rdma

import "strings"

const (
	ReasonNoSafeRouteGID     = "active Thunderbolt RDMA device has no automatic-safe route GID; index 0 requires an explicit diagnostic override"
	ReasonRouteGIDIndexZero  = "active Thunderbolt RDMA route GID resolved to index 0"
	ReasonNoActiveTBDevice   = "no PORT_ACTIVE Thunderbolt RDMA device found"
	ReasonNoPeerInterface    = "no AppleThunderboltRDMAPeerInterface entries found"
	ReasonNoXDomainService   = "no IOThunderboltXDomainService entries found"
	ReasonNoRecentLog        = "no recent AppleThunderboltRDMA log lines captured"
	ReasonRecentRTRFailure   = "recent AppleThunderboltRDMA log contains Failed INIT->RTR"
	ReasonReadOnlyPreflight  = "read-only preflight passed; safe_to_attempt_rtr is necessary, not sufficient"
	IORegistryPeerInterface  = "AppleThunderboltRDMAPeerInterface"
	IORegistryXDomainService = "IOThunderboltXDomainService"

	RTRAttemptWarning = "rdma-pingpong drives QP INIT->RTR, which can wedge Apple Thunderbolt RDMA ports; run rdmainfo preflight, run rdma-probe, and read the README first, then pass -allow-rtr for one bounded attempt"
)

// PreflightReport is the read-only evidence needed to decide whether an RDMA
// RTR attempt is even eligible to run.
type PreflightReport struct {
	Devices    []PreflightDevice
	IORegistry map[string]int
	RecentLog  []string
}

// PreflightDevice is the read-only state for one RDMA device.
type PreflightDevice struct {
	Name          string
	State         int32
	LinkLayer     uint8
	RouteGIDIndex *int
}

// DerivePreflightSafety returns whether report permits one bounded RTR attempt.
//
// A true result is necessary, not sufficient. It is not a production or
// performance proof.
func DerivePreflightSafety(report PreflightReport) (bool, []string) {
	var reasons []string
	activeThunderbolt := false
	for _, dev := range report.Devices {
		if dev.State == PortActive && dev.LinkLayer == LinkLayerThunderbolt {
			activeThunderbolt = true
			if dev.RouteGIDIndex == nil {
				reasons = append(reasons, ReasonNoSafeRouteGID)
				continue
			}
			if *dev.RouteGIDIndex == 0 {
				reasons = append(reasons, ReasonRouteGIDIndexZero)
			}
		}
	}
	if !activeThunderbolt {
		reasons = append(reasons, ReasonNoActiveTBDevice)
	}
	if report.IORegistry[IORegistryPeerInterface] == 0 {
		reasons = append(reasons, ReasonNoPeerInterface)
	}
	if report.IORegistry[IORegistryXDomainService] == 0 {
		reasons = append(reasons, ReasonNoXDomainService)
	}
	if len(report.RecentLog) == 0 {
		reasons = append(reasons, ReasonNoRecentLog)
	}
	for _, line := range report.RecentLog {
		if FailedRTRLogLine(line) {
			reasons = append(reasons, ReasonRecentRTRFailure)
			break
		}
	}
	if len(reasons) == 0 {
		reasons = append(reasons, ReasonReadOnlyPreflight)
		return true, reasons
	}
	return false, reasons
}

// FailedRTRLogLine reports whether line contains Apple's INIT-to-RTR failure
// marker. Apple logs have used both ASCII "->" and Unicode "→".
func FailedRTRLogLine(line string) bool {
	line = strings.ReplaceAll(line, "→", "->")
	return strings.Contains(line, "Failed INIT->RTR")
}

type rtrUnsafeError struct{}

func (rtrUnsafeError) Error() string {
	return RTRAttemptWarning
}

func (rtrUnsafeError) Is(target error) bool {
	return target == ErrRTRUnsafe
}

// RequireRTRAttemptAllowed returns ErrRTRUnsafe unless allow is true.
func RequireRTRAttemptAllowed(allow bool) error {
	if allow {
		return nil
	}
	return rtrUnsafeError{}
}
