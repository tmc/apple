package zerocopy_test

import (
	"fmt"

	"github.com/tmc/apple/x/zerocopy"
)

func Example() {
	buf := make([]byte, 1024)

	alias := buf[:512] // shares memory with buf
	fmt.Println("alias:", zerocopy.Check(alias, buf))

	copied := append([]byte(nil), buf...) // does not
	err := zerocopy.Check(buf, copied)
	fmt.Println("copy fails:", err != nil)
	// Output:
	// alias: <nil>
	// copy fails: true
}
