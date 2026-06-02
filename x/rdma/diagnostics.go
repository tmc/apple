package rdma

import (
	"errors"
	"syscall"

	"github.com/tmc/apple/rdma"
)

// ResourceExhaustionHint returns an operator hint for failures that can
// indicate the Apple Thunderbolt RDMA per-boot resource exhaustion pattern.
func ResourceExhaustionHint(err error) string {
	var e *rdma.ProviderError
	if !errors.As(err, &e) || e == nil || !e.ContextOpen {
		return ""
	}
	if e.ErrnoSet {
		switch e.Errno {
		case int(syscall.ENOMEM), int(syscall.EBUSY):
			return "provider status may indicate per-boot AppleThunderboltRDMA resource exhaustion or contaminated IOKit state; no provider resource budget was read; stop live RDMA probes and reboot the affected node before retrying"
		}
	}
	if e.Failure == rdma.FailureProviderTimeout {
		return "AppleThunderboltRDMA provider may be wedged for this boot; no provider resource budget was read; stop live RDMA probes and reboot the affected node before retrying"
	}
	if e.Failure == rdma.FailureNilProviderResult && providerResourceOperation(e.Operation) {
		return "provider returned nil after opening a context; this can indicate per-boot AppleThunderboltRDMA resource exhaustion or contaminated IOKit state; no provider resource budget was read; stop live RDMA probes and reboot before retrying"
	}
	return ""
}

func providerResourceOperation(op string) bool {
	switch op {
	case "ibv_alloc_pd", "ibv_create_cq", "ibv_create_qp", "ibv_reg_mr":
		return true
	default:
		return false
	}
}
