// Command uuid-date prints the current NSDate and a freshly generated NSUUID.
package main

import (
	"fmt"

	"github.com/tmc/apple/foundation"
)

func main() {
	now := foundation.GetNSDateClass().Date()
	uuid := foundation.NewNSUUID()

	fmt.Printf("now=%s\n", now.Description())
	fmt.Printf("uuid=%s\n", uuid.UUIDString())
}
