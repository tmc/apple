//go:build darwin

package rdma_test

import (
	"fmt"

	"github.com/tmc/apple/rdma"
)

func ExampleClassifyCompletionStatus() {
	// Each class is reached by a different status code, so the example covers
	// one of each rather than only the successful path.
	fmt.Println(rdma.ClassifyCompletionStatus(rdma.IBV_WC_SUCCESS))
	fmt.Println(rdma.ClassifyCompletionStatus(rdma.IBV_WC_LOC_PROT_ERR))
	fmt.Println(rdma.ClassifyCompletionStatus(rdma.IBV_WC_REM_ACCESS_ERR))

	// Any other status is a provider failure rather than a protection error;
	// the distinction is the point of the classifier.
	fmt.Println(rdma.ClassifyCompletionStatus(rdma.IBV_WC_LOC_ACCESS_ERR + 1))

	// Output:
	// success
	// protection error
	// protection error
	// failure
}
