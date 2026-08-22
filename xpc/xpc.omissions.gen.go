// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

type swiftMemberOmission struct {
	Identifier string
	Title      string
	SymbolKind string
	Reason     string
}

var xpcSwiftOmissions = []swiftMemberOmission{
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions",
		Title:      "XPCListener.InitializationOptions",
		SymbolKind: "struct",
		Reason:     "not included in v1 generated surface",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions/inactive",
		Title:      "inactive",
		SymbolKind: "property",
		Reason:     "not included in v1 generated surface",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions/none",
		Title:      "none",
		SymbolKind: "property",
		Reason:     "not included in v1 generated surface",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCListener/endpoint",
		Title:      "endpoint",
		SymbolKind: "property",
		Reason:     "out of scope for v1: xpc/listener.h exports no endpoint accessor, and xpc_endpoint_create takes xpc_connection_t (xpc/endpoint.h:22; :12-14 documents other argument types as undefined behavior); endpoint export is available on the classic connection API, already bound, which v1 does not layer under Listener",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/expectsReply",
		Title:      "expectsReply",
		SymbolKind: "property",
		Reason:     "not included in v1 generated surface",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/handoffReply(to:_:)",
		Title:      "handoffReply(to:_:)",
		SymbolKind: "method",
		Reason:     "no public C equivalent: the string \"handoff\" appears in no public xpc header",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/isSync",
		Title:      "isSync",
		SymbolKind: "property",
		Reason:     "not included in v1 generated surface",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:cancellationHandler:)",
		Title:      "init(endpoint:targetQueue:options:cancellationHandler:)",
		SymbolKind: "init",
		Reason:     "out of scope for v1: xpc/session.h declares exactly two session constructors (xpc_session_create_xpc_service, xpc_session_create_mach_service), neither endpoint-taking; an endpoint peer is fully expressible on the classic connection API (xpc_connection_create_from_endpoint, already bound), which v1 does not layer under Session",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)-2jmkk",
		Title:      "init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)",
		SymbolKind: "init",
		Reason:     "out of scope for v1: xpc/session.h declares exactly two session constructors (xpc_session_create_xpc_service, xpc_session_create_mach_service), neither endpoint-taking; an endpoint peer is fully expressible on the classic connection API (xpc_connection_create_from_endpoint, already bound), which v1 does not layer under Session",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)-546jo",
		Title:      "init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)",
		SymbolKind: "init",
		Reason:     "out of scope for v1: xpc/session.h declares exactly two session constructors (xpc_session_create_xpc_service, xpc_session_create_mach_service), neither endpoint-taking; an endpoint peer is fully expressible on the classic connection API (xpc_connection_create_from_endpoint, already bound), which v1 does not layer under Session",
	},
	{
		Identifier: "doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)-6zd1x",
		Title:      "init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)",
		SymbolKind: "init",
		Reason:     "out of scope for v1: xpc/session.h declares exactly two session constructors (xpc_session_create_xpc_service, xpc_session_create_mach_service), neither endpoint-taking; an endpoint peer is fully expressible on the classic connection API (xpc_connection_create_from_endpoint, already bound), which v1 does not layer under Session",
	},
}

// xpcSwiftMemberPopulation records how many member documents the generating run
// classified, per Swift root type.
//
// The ledger above is a difference: population minus the members the generated
// surface covers. Its length therefore says nothing on its own, because the
// population is whatever the doc fetch had loaded — a member whose page was not
// in the set is absent from the ledger and from the covered surface alike, and
// nothing reports it. Recording the denominator is what lets a test tell "the
// ledger is complete" apart from "the ledger is short because the run saw
// less". See TestSwiftOmissionLedgerPopulation.
var xpcSwiftMemberPopulation = map[string]int{
	"Listener":        21,
	"Session":         40,
	"ReceivedMessage": 6,
}
