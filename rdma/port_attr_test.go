package rdma

import (
	"testing"
	"unsafe"
)

// TestIbvPortAttrActiveMTULayout guards the generated partial port-attribute
// binding.
func TestIbvPortAttrActiveMTULayout(t *testing.T) {
	if got, want := unsafe.Offsetof(IbvPortAttr{}.ActiveMTU), uintptr(8); got != want {
		t.Fatalf("IbvPortAttr.ActiveMTU offset = %d, want %d", got, want)
	}
}

// These declarations keep the native JACCL-facing verbs API in the build.
var (
	_ = IbvModifyQpToErr
	_ = IbvPortAttr{ActiveMTU: IBV_MTU_1024}
)
