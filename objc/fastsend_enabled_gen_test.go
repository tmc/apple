//go:build darwin && !objc_slowpath

package objc

import "testing"

// The other half of the source check: a fastsend_enabled.gen.go that stopped
// calling initFastSend would pass it while disabling the fast path everywhere.
func TestFastSendArmedByDefault(t *testing.T) {
	if objcMsgSendAddr == 0 {
		t.Error("objcMsgSendAddr is zero in a default build: the fast path is " +
			"not armed, so every Send takes the reflect path")
	}
}
