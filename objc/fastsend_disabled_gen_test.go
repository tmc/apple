//go:build darwin && objc_slowpath

package objc

import "testing"

// Run with: go test -tags objc_slowpath ./objc/
//
// This is the assertion that the tag does something. It failed for months
// without anyone noticing, because nothing ran it: the tag was inert, and the
// only test that would have caught it did not exist.
func TestFastSendDisabledBySlowpathTag(t *testing.T) {
	if objcMsgSendAddr != 0 {
		t.Errorf("objcMsgSendAddr is %#x under -tags objc_slowpath: the fast "+
			"path is armed anyway, so the tag is inert", objcMsgSendAddr)
	}
}
