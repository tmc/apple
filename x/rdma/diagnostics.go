package rdma

import (
	"errors"
	"syscall"

	"github.com/tmc/apple/rdma"
)

// ResourceExhaustionHint returns an operator hint for a provider failure after
// opening a context. The hint distinguishes a resource refusal from a timeout;
// neither result alone proves that the provider is wedged.
func ResourceExhaustionHint(err error) string {
	var e *rdma.ProviderError
	if !errors.As(err, &e) || e == nil || !e.ContextOpen {
		return ""
	}
	if e.ErrnoSet {
		switch e.Errno {
		case int(syscall.ENOMEM), int(syscall.EBUSY):
			return "provider refused a resource after opening a context; stop the current run and inspect both hosts with read-only rdma-probe before any new attempt"
		}
	}
	if e.Failure == rdma.FailureProviderTimeout {
		return "provider call timed out after opening a context; watchdog containment alone does not prove a wedge, so inspect both hosts with read-only rdma-probe before any new attempt"
	}
	if e.Failure == rdma.FailureNilProviderResult && providerResourceOperation(e.Operation) {
		return "provider returned nil for a resource after opening a context; stop the current run and inspect both hosts with read-only rdma-probe before any new attempt"
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
