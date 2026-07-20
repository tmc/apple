package rdma

import (
	"testing"

	apple "github.com/tmc/apple/rdma"
)

func TestJACCLNativeAPI(t *testing.T) {
	_, _, err := RTRAttr(
		LocalQP{PortNum: 1, ActiveMTU: apple.IBV_MTU_1024},
		RemoteQP{QPN: 1, ActiveMTU: apple.IBV_MTU_1024},
		RTRPolicy{},
	)
	if err != nil {
		t.Fatal(err)
	}
}
