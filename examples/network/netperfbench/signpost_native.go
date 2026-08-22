//go:build darwin && native_signpost

package main

/*
#include <os/signpost.h>
#include <stdint.h>

enum {
	npRoundTrip,
	npSend,
	npSendCallback,
	npSendSignal,
	npReceive,
	npReceiveCallback,
	npReceiveSignal,
	npReceiveWait,
	npSendWait,
};

static os_log_t npLog;

static void
npSignpostInit(void)
{
	npLog = os_log_create("github.com.tmc.apple.netperfbench", OS_LOG_CATEGORY_POINTS_OF_INTEREST);
}

static int
npSignpostEnabled(void)
{
	return npLog != NULL && os_signpost_enabled(npLog);
}

static uint64_t
npSignpostID(void)
{
	return os_signpost_id_generate(npLog);
}

// Each case is a C macro call site with a literal name. os_signpost records
// those literals in __oslogstring; passing a dynamically allocated Go string
// creates an interval but leaves Instruments' Name column blank.
static void
npSignpostEmit(uint8_t type, uint64_t id, uint32_t name)
{
	switch (name) {
	case npRoundTrip:
		os_signpost_emit_with_type(npLog, type, id, "round-trip");
		break;
	case npSend:
		os_signpost_emit_with_type(npLog, type, id, "send");
		break;
	case npSendCallback:
		os_signpost_emit_with_type(npLog, type, id, "send-callback");
		break;
	case npSendSignal:
		os_signpost_emit_with_type(npLog, type, id, "send-signal");
		break;
	case npReceive:
		os_signpost_emit_with_type(npLog, type, id, "receive");
		break;
	case npReceiveCallback:
		os_signpost_emit_with_type(npLog, type, id, "receive-callback");
		break;
	case npReceiveSignal:
		os_signpost_emit_with_type(npLog, type, id, "receive-signal");
		break;
	case npReceiveWait:
		os_signpost_emit_with_type(npLog, type, id, "receive-wait");
		break;
	case npSendWait:
		os_signpost_emit_with_type(npLog, type, id, "send-wait");
		break;
	}
}
*/
import "C"

import "flag"

var traceSignposts = flag.Bool("signpost", false, "emit send and receive signposts for xctrace")

const (
	signpostRoundTrip = iota
	signpostSend
	signpostSendCallback
	signpostSendSignal
	signpostReceive
	signpostReceiveCallback
	signpostReceiveSignal
	signpostReceiveWait
	signpostSendWait
)

const (
	signpostEvent = iota
	signpostBegin
	signpostEnd
)

func enableSignposts() {
	if *traceSignposts {
		C.npSignpostInit()
	}
}

type signpostInterval struct {
	id   uint64
	name uint32
}

func beginSignpost(name string) *signpostInterval {
	if !*traceSignposts || C.npSignpostEnabled() == 0 {
		return nil
	}
	id := uint64(C.npSignpostID())
	nameID := signpostName(name)
	C.npSignpostEmit(C.uint8_t(signpostBegin), C.uint64_t(id), C.uint32_t(nameID))
	return &signpostInterval{id: id, name: nameID}
}

func (s *signpostInterval) end() {
	if s != nil {
		C.npSignpostEmit(C.uint8_t(signpostEnd), C.uint64_t(s.id), C.uint32_t(s.name))
	}
}

func (s *signpostInterval) event(name string) {
	if s != nil {
		C.npSignpostEmit(C.uint8_t(signpostEvent), C.uint64_t(s.id), C.uint32_t(signpostName(name)))
	}
}

func signpostName(name string) uint32 {
	switch name {
	case "round-trip":
		return signpostRoundTrip
	case "send":
		return signpostSend
	case "send-callback":
		return signpostSendCallback
	case "send-signal":
		return signpostSendSignal
	case "receive":
		return signpostReceive
	case "receive-callback":
		return signpostReceiveCallback
	case "receive-signal":
		return signpostReceiveSignal
	case "receive-wait":
		return signpostReceiveWait
	case "send-wait":
		return signpostSendWait
	}
	return signpostRoundTrip
}
