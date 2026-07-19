package rdma

import (
	"strings"
	"syscall"
	"testing"

	"github.com/tmc/apple/rdma"
)

func TestResourceExhaustionHintForNilResourceResult(t *testing.T) {
	for _, op := range []string{"ibv_alloc_pd", "ibv_create_cq", "ibv_create_qp", "ibv_reg_mr"} {
		t.Run(op, func(t *testing.T) {
			err := &rdma.ProviderError{
				Operation:   op,
				ContextOpen: true,
				Failure:     rdma.FailureNilProviderResult,
				Result:      "resource",
				Cause:       rdma.ErrNilProviderResult,
			}
			hint := ResourceExhaustionHint(err)
			if !strings.Contains(hint, "provider returned nil for a resource") {
				t.Fatalf("hint = %q, want resource-refusal hint", hint)
			}
			if !strings.Contains(hint, "read-only rdma-probe") {
				t.Fatalf("hint = %q, want read-only follow-up", hint)
			}
		})
	}
}

func TestResourceExhaustionHintForResourceErrnos(t *testing.T) {
	for _, errno := range []syscall.Errno{syscall.ENOMEM, syscall.EBUSY} {
		t.Run(errno.Error(), func(t *testing.T) {
			err := &rdma.ProviderError{
				Operation:   "ibv_modify_qp",
				ContextOpen: true,
				Failure:     rdma.FailureProviderStatus,
				Errno:       int(errno),
				ErrnoSet:    true,
				Cause:       rdma.ErrProviderStatus,
			}
			hint := ResourceExhaustionHint(err)
			if !strings.Contains(hint, "provider refused a resource") {
				t.Fatalf("hint = %q, want resource-refusal hint", hint)
			}
			if !strings.Contains(hint, "read-only rdma-probe") {
				t.Fatalf("hint = %q, want read-only follow-up", hint)
			}
		})
	}
}

func TestResourceExhaustionHintForProviderTimeout(t *testing.T) {
	err := &rdma.ProviderError{
		Operation:   "ibv_alloc_pd",
		ContextOpen: true,
		Failure:     rdma.FailureProviderTimeout,
		Cause:       rdma.ErrProviderTimeout,
	}
	if hint := ResourceExhaustionHint(err); !strings.Contains(hint, "watchdog containment alone does not prove a wedge") {
		t.Fatalf("hint = %q, want containment hint", hint)
	}
}

func TestResourceExhaustionHintDoesNotFireBeforeOpen(t *testing.T) {
	err := &rdma.ProviderError{
		Operation:   "ibv_open_device",
		ContextOpen: false,
		Failure:     rdma.FailureNilProviderResult,
		Errno:       int(syscall.EBUSY),
		ErrnoSet:    true,
		Cause:       rdma.ErrNilProviderResult,
	}
	if hint := ResourceExhaustionHint(err); hint != "" {
		t.Fatalf("hint = %q, want empty before open", hint)
	}
}

func TestResourceExhaustionHintDoesNotFireForNonResourceNilResult(t *testing.T) {
	err := &rdma.ProviderError{
		Operation:   "ibv_query_port",
		ContextOpen: true,
		Failure:     rdma.FailureNilProviderResult,
		Result:      "port attributes",
		Cause:       rdma.ErrNilProviderResult,
	}
	if hint := ResourceExhaustionHint(err); hint != "" {
		t.Fatalf("hint = %q, want empty for non-resource nil result", hint)
	}
}
