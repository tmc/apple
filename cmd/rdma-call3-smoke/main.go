// Command rdma-call3-smoke checks that a production binary links the Call3
// purego trampoline.
package main

import (
	"fmt"

	_ "github.com/tmc/apple/rdma"
)

func main() {
	fmt.Println("rdma Call3 trampoline: available")
}
