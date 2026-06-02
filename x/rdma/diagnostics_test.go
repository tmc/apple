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
			if !strings.Contains(hint, "can indicate per-boot AppleThunderboltRDMA resource exhaustion") {
				t.Fatalf("hint = %q, want symptom-inferred provider resource hint", hint)
			}
			if !strings.Contains(hint, "no provider resource budget was read") {
				t.Fatalf("hint = %q, want no-budget-read caveat", hint)
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
			if !strings.Contains(hint, "may indicate per-boot AppleThunderboltRDMA resource exhaustion") {
				t.Fatalf("hint = %q, want symptom-inferred provider resource hint", hint)
			}
			if !strings.Contains(hint, "no provider resource budget was read") {
				t.Fatalf("hint = %q, want no-budget-read caveat", hint)
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
	if hint := ResourceExhaustionHint(err); !strings.Contains(hint, "provider may be wedged for this boot") {
		t.Fatalf("hint = %q, want wedged-provider hint", hint)
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
