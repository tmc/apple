package signpost

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

// TestMessageRoundTrip emits a message-carrying interval and reads it back
// through log show, verifying the hand-built "%{public}s" argument buffer
// decodes to the message rather than crashing or dropping it.
func TestMessageRoundTrip(t *testing.T) {
	log := New("com.tmc.apple.signposttest", PointsOfInterest)
	if !log.Enabled() {
		t.Skip("signposts disabled for this log handle")
	}
	if namePool()[messageFormat] == nil {
		t.Skip("message format not pooled in __oslogstring (cgo-free test binary); message would render as compose failure")
	}
	msg := fmt.Sprintf("roundtrip-%d", time.Now().UnixNano())
	id := log.NewID()
	log.IntervalBeginMessage(id, "MsgTest", msg)
	log.IntervalEndMessage(id, "MsgTest", msg)
	log.EventMessage(log.NewID(), "MsgTestEvent", msg)

	deadline := time.Now().Add(15 * time.Second)
	for {
		out, err := exec.Command("log", "show", "--last", "1m", "--signpost",
			"--predicate", `subsystem == "com.tmc.apple.signposttest"`).Output()
		if err != nil {
			t.Skipf("log show unavailable: %v", err)
		}
		if bytes.Contains(out, []byte(msg)) {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("message %q not found in log show output:\n%s", msg, out)
		}
		time.Sleep(time.Second)
	}
}

// TestMessageEmptyFallsBack ensures an empty message takes the name-only
// path and emits without constructing an argument buffer.
func TestMessageEmptyFallsBack(t *testing.T) {
	log := New("com.tmc.apple.signposttest", PointsOfInterest)
	id := log.NewID()
	log.IntervalBeginMessage(id, "EmptyMsg", "")
	log.IntervalEndMessage(id, "EmptyMsg", "")
}
