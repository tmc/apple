// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

type swiftMemberOmission struct {
	Identifier string
	Title      string
	SymbolKind string
	Reason     string
}

// rawSymbolOmission records a deliberate decision not to expose one raw C
// surface. Symbols may name a family when one Go codec operation supersedes
// every member of that family.
type rawSymbolOmission struct {
	Symbols []string
	Reason  string
}

var xpcRawOmissions = []rawSymbolOmission{
	{
		Symbols: []string{"xpc_(array|dictionary)_(get|set)_(bool|int64|uint64|double|string|data|date|uuid|array|dictionary|value|count)"},
		Reason:  "the Dictionary codec exposes these dynamic values as ordinary Go map and slice operations; parallel typed accessors would duplicate and potentially disagree with the codec",
	},
	{Symbols: []string{"xpc_array_create", "xpc_dictionary_create", "xpc_string_get_length"}, Reason: "internal construction and inspection are already covered by the codec; no distinct Go operation is needed"},
	{Symbols: []string{"xpc_array_dup_fd", "xpc_dictionary_dup_fd", "xpc_array_set_fd", "xpc_dictionary_set_fd"}, Reason: "FileDescriptor is the single ownership-safe representation for descriptor transfer"},
	{Symbols: []string{"xpc_data_create_with_dispatch_data", "xpc_data_get_bytes"}, Reason: "the dispatch bridge and copyRawData already provide the safe Go ownership and copy semantics"},
	{Symbols: []string{"xpc_string_create_with_format", "xpc_string_create_with_format_and_arguments"}, Reason: "C varargs are not safely expressible through purego; use fmt.Sprintf before encoding"},
	{Symbols: []string{"xpc_connection_get_context", "xpc_connection_set_context", "xpc_connection_set_finalizer_f"}, Reason: "storing Go pointers in C context memory violates Go pointer and GC rules; no safe tokenized contract has been defined"},
	{Symbols: []string{"xpc_main"}, Reason: "never returns and conflicts with the Go runtime main goroutine; Listener plus select is the supported service shape"},
	{Symbols: []string{"xpc_(array|dictionary)_(create_connection|set_connection)", "xpc_dictionary_get_remote_connection"}, Reason: "connection-valued container entries need a retained peer and a server-side ownership model that the Dictionary codec does not yet define"},
	{Symbols: []string{"xpc_dictionary_(copy_mach_send|set_mach_send)"}, Reason: "Mach send-right ownership needs an explicit mach_port_deallocate lifecycle type before it can be exposed safely"},
	{Symbols: []string{"xpc_connection_send_barrier", "xpc_connection_send_message_with_reply"}, Reason: "asynchronous handlers require a cancellable token lifecycle distinct from the synchronous Connection call API"},
	{Symbols: []string{"xpc_connection_set_peer_entitlement_matches_value_requirement", "xpc_connection_set_peer_lightweight_code_requirement"}, Reason: "the value-bearing requirement setters need a dedicated ownership contract; PeerRequirement constructors are the supported API"},
	{Symbols: []string{"xpc_activity_get_state", "xpc_activity_set_state", "xpc_activity_set_criteria"}, Reason: "activity state sentinels and mutable criteria require named platform constants and lifecycle semantics beyond the initial retained Activity API"},
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
