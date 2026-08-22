// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

// rawSyms_<Entry> is the transitive set of raw XPC symbols reachable from
// <Entry>, including through callees. Guards use these instead of
// hand-listed literals so a newly added raw_ call cannot go unguarded.

var rawSyms_Listener_Activate = []string{"xpc_listener_activate", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_Listener_Cancel = []string{"xpc_listener_cancel"}
var rawSyms_PeerRequirement_Close = []string{"xpc_release"}
var rawSyms_Session_Activate = []string{"xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate"}
var rawSyms_Session_Call = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_Session_CallDictionary = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_Session_Cancel = []string{"xpc_session_cancel"}
var rawSyms_Session_Notify = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_NotifyDictionary = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_SetCancellationHandler = []string{"xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_set_cancel_handler"}
var rawSyms_Session_SetIncomingMessageHandler = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_SetPeerRequirement = []string{"xpc_session_set_peer_requirement"}
var rawSyms_Session_SetTargetQueue = []string{"xpc_session_set_target_queue"}
var rawSyms_ReceivedMessage_Decode = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_ReceivedMessage_Dictionary = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_ReceivedMessage_SenderSatisfies = []string{"xpc_peer_requirement_match_received_message"}
var rawSyms_DialMachService = []string{"xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_set_peer_requirement"}
var rawSyms_DialXPCService = []string{"xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_set_peer_requirement"}
var rawSyms_NewAnonymousListener = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_listener_activate", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_send_message", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewEntitlementExistsRequirement = []string{"xpc_peer_requirement_create_entitlement_exists", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewEntitlementMatchesRequirement = []string{"xpc_bool_create", "xpc_int64_create", "xpc_peer_requirement_create_entitlement_matches_value", "xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_string_create"}
var rawSyms_NewLightweightCodeRequirement = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_peer_requirement_create_lwcr", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewPlatformBinaryRequirement = []string{"xpc_peer_requirement_create_platform_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewPlatformBinarySignedAsRequirement = []string{"xpc_peer_requirement_create_platform_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewSameTeamRequirement = []string{"xpc_peer_requirement_create_team_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewSameTeamSignedAsRequirement = []string{"xpc_peer_requirement_create_team_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewServiceListener = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_listener_activate", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_send_message", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_PeerRequirementFromHandle = []string{"xpc_retain"}

// rawReachability indexes the above by entry point, for the tests.
var rawReachability = map[string][]string{
	"(*Listener).Activate":                 rawSyms_Listener_Activate,
	"(*Listener).Cancel":                   rawSyms_Listener_Cancel,
	"(*PeerRequirement).Close":             rawSyms_PeerRequirement_Close,
	"(*Session).Activate":                  rawSyms_Session_Activate,
	"(*Session).Call":                      rawSyms_Session_Call,
	"(*Session).CallDictionary":            rawSyms_Session_CallDictionary,
	"(*Session).Cancel":                    rawSyms_Session_Cancel,
	"(*Session).Notify":                    rawSyms_Session_Notify,
	"(*Session).NotifyDictionary":          rawSyms_Session_NotifyDictionary,
	"(*Session).SetCancellationHandler":    rawSyms_Session_SetCancellationHandler,
	"(*Session).SetIncomingMessageHandler": rawSyms_Session_SetIncomingMessageHandler,
	"(*Session).SetPeerRequirement":        rawSyms_Session_SetPeerRequirement,
	"(*Session).SetTargetQueue":            rawSyms_Session_SetTargetQueue,
	"(ReceivedMessage).Decode":             rawSyms_ReceivedMessage_Decode,
	"(ReceivedMessage).Dictionary":         rawSyms_ReceivedMessage_Dictionary,
	"(ReceivedMessage).SenderSatisfies":    rawSyms_ReceivedMessage_SenderSatisfies,
	"DialMachService":                      rawSyms_DialMachService,
	"DialXPCService":                       rawSyms_DialXPCService,
	"NewAnonymousListener":                 rawSyms_NewAnonymousListener,
	"NewEntitlementExistsRequirement":      rawSyms_NewEntitlementExistsRequirement,
	"NewEntitlementMatchesRequirement":     rawSyms_NewEntitlementMatchesRequirement,
	"NewLightweightCodeRequirement":        rawSyms_NewLightweightCodeRequirement,
	"NewPlatformBinaryRequirement":         rawSyms_NewPlatformBinaryRequirement,
	"NewPlatformBinarySignedAsRequirement": rawSyms_NewPlatformBinarySignedAsRequirement,
	"NewSameTeamRequirement":               rawSyms_NewSameTeamRequirement,
	"NewSameTeamSignedAsRequirement":       rawSyms_NewSameTeamSignedAsRequirement,
	"NewServiceListener":                   rawSyms_NewServiceListener,
	"PeerRequirementFromHandle":            rawSyms_PeerRequirementFromHandle,
}

// rawAllCalledSymbols is the union of raw symbols named by any function in
// the package, exported or not.
var rawAllCalledSymbols = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_listener_activate", "xpc_listener_cancel", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_peer_requirement_create_entitlement_exists", "xpc_peer_requirement_create_entitlement_matches_value", "xpc_peer_requirement_create_lwcr", "xpc_peer_requirement_create_platform_identity", "xpc_peer_requirement_create_team_identity", "xpc_peer_requirement_match_received_message", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_cancel", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_send_message", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_session_set_peer_requirement", "xpc_session_set_target_queue", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}

// rawReachEdge is one call edge the generator could not follow.
type rawReachEdge struct {
	Entry, File, Expr, Kind string
	Line                    int
}

// rawReachUnresolved lists call edges the generator could not follow. Each
// entry is a place where a reachable set above may be an UNDER-estimate.
// It is emitted, not dropped.
var rawReachUnresolved = []rawReachEdge{
	{Entry: "(RichError).Error", File: "xpc.highlevel.gen.go", Line: 95, Expr: "e.cause.Error(...)", Kind: "unresolvable receiver expression"},
	{Entry: "decodeJSONPayload", File: "xpc.highlevel.gen.go", Line: 580, Expr: "dec.Decode(...)", Kind: "unresolved receiver type"},
	{Entry: "newListener", File: "xpc.highlevel.gen.go", Line: 834, Expr: "incoming(...)", Kind: "func-value or unknown callee"},
	{Entry: "newRequirement", File: "xpc.highlevel.gen.go", Line: 437, Expr: "create(...)", Kind: "func-value or unknown callee"},
	{Entry: "targetQueuePointer", File: "xpc.highlevel.gen.go", Line: 268, Expr: "queue.Handle(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcCancelTrampoline", File: "xpc.highlevel.gen.go", Line: 2198, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2163, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2171, Expr: "err.Error(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2176, Expr: "encErr.Error(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcReplyTrampoline", File: "xpc.highlevel.gen.go", Line: 2273, Expr: "reply(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcReplyTrampoline", File: "xpc.highlevel.gen.go", Line: 2277, Expr: "reply(...)", Kind: "func-value or unknown callee"},
}
