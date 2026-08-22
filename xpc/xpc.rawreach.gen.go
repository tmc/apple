// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

// rawSyms_<Entry> is the transitive set of raw XPC symbols reachable from
// <Entry>, including through callees. Guards use these instead of
// hand-listed literals so a newly added raw_ call cannot go unguarded.

var rawSyms_Activity_Close = []string{"xpc_release"}
var rawSyms_Activity_Criteria = []string{"xpc_activity_copy_criteria", "xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_release", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_Activity_ShouldDefer = []string{"xpc_activity_should_defer"}
var rawSyms_Connection_Activate = []string{"xpc_connection_activate"}
var rawSyms_Connection_AuditSessionID = []string{"xpc_connection_get_asid"}
var rawSyms_Connection_CallDictionary = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_connection_send_message_with_reply_sync", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_Connection_Cancel = []string{"xpc_connection_cancel"}
var rawSyms_Connection_EGID = []string{"xpc_connection_get_egid"}
var rawSyms_Connection_EUID = []string{"xpc_connection_get_euid"}
var rawSyms_Connection_Endpoint = []string{"xpc_endpoint_create"}
var rawSyms_Connection_InvalidationReason = []string{"xpc_connection_copy_invalidation_reason"}
var rawSyms_Connection_Name = []string{"xpc_connection_get_name"}
var rawSyms_Connection_PID = []string{"xpc_connection_get_pid"}
var rawSyms_Connection_Resume = []string{"xpc_connection_resume"}
var rawSyms_Connection_SendDictionary = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_connection_send_message", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Connection_SetPeerCodeSigningRequirement = []string{"xpc_connection_set_peer_code_signing_requirement"}
var rawSyms_Connection_SetPeerEntitlementExistsRequirement = []string{"xpc_connection_set_peer_entitlement_exists_requirement"}
var rawSyms_Connection_SetPeerPlatformIdentityRequirement = []string{"xpc_connection_set_peer_platform_identity_requirement"}
var rawSyms_Connection_SetPeerRequirement = []string{"xpc_connection_set_peer_requirement"}
var rawSyms_Connection_SetPeerTeamIdentityRequirement = []string{"xpc_connection_set_peer_team_identity_requirement"}
var rawSyms_Connection_SetTargetQueue = []string{"xpc_connection_set_target_queue"}
var rawSyms_Connection_Suspend = []string{"xpc_connection_suspend"}
var rawSyms_FileDescriptor_Close = []string{"xpc_release"}
var rawSyms_FileDescriptor_Dup = []string{"xpc_fd_dup"}
var rawSyms_Listener_Activate = []string{"xpc_listener_activate", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_Listener_Cancel = []string{"xpc_listener_cancel"}
var rawSyms_Listener_SetPeerCodeSigningRequirement = []string{"xpc_listener_set_peer_code_signing_requirement"}
var rawSyms_Listener_String = []string{"xpc_listener_copy_description"}
var rawSyms_PeerRequirement_Close = []string{"xpc_release"}
var rawSyms_Session_Activate = []string{"xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate"}
var rawSyms_Session_Call = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_Session_CallDictionary = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_Session_Cancel = []string{"xpc_session_cancel"}
var rawSyms_Session_Notify = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_NotifyDictionary = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_SetCancellationHandler = []string{"xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_set_cancel_handler"}
var rawSyms_Session_SetIncomingMessageHandler = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_send_message", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Session_SetPeerCodeSigningRequirement = []string{"xpc_session_set_peer_code_signing_requirement"}
var rawSyms_Session_SetPeerRequirement = []string{"xpc_session_set_peer_requirement"}
var rawSyms_Session_SetTargetQueue = []string{"xpc_session_set_target_queue"}
var rawSyms_Session_String = []string{"xpc_session_copy_description"}
var rawSyms_SharedMemory_Close = []string{"xpc_release"}
var rawSyms_SharedMemory_Map = []string{"xpc_shmem_map"}
var rawSyms_ReceivedMessage_Decode = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_ReceivedMessage_Dictionary = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_ReceivedMessage_SenderSatisfies = []string{"xpc_peer_requirement_match_received_message"}
var rawSyms_ActivateSocket = []string{"launch_activate_socket"}
var rawSyms_CopyValue = []string{"xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_copy", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_get_type", "xpc_int64_create", "xpc_int64_get_value", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}
var rawSyms_CurrentDate = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create_from_current", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_release", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_DialMachService = []string{"xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_set_peer_requirement"}
var rawSyms_DialXPCService = []string{"xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_set_peer_requirement"}
var rawSyms_Equal = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_equal", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_Hash = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_hash", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewAnonymousConnection = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_connection_create", "xpc_connection_set_event_handler", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_release", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_NewAnonymousListener = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_listener_activate", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_send_message", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewConnectionFromEndpoint = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_connection_create_from_endpoint", "xpc_connection_set_event_handler", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_release", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_NewEntitlementExistsRequirement = []string{"xpc_peer_requirement_create_entitlement_exists", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewEntitlementMatchesRequirement = []string{"xpc_bool_create", "xpc_int64_create", "xpc_peer_requirement_create_entitlement_matches_value", "xpc_release", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_string_create"}
var rawSyms_NewFileDescriptor = []string{"xpc_fd_create"}
var rawSyms_NewLightweightCodeRequirement = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_peer_requirement_create_lwcr", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewMachServiceConnection = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_connection_create_mach_service", "xpc_connection_set_event_handler", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_release", "xpc_retain", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_NewPlatformBinaryRequirement = []string{"xpc_peer_requirement_create_platform_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewPlatformBinarySignedAsRequirement = []string{"xpc_peer_requirement_create_platform_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewSameTeamRequirement = []string{"xpc_peer_requirement_create_team_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewSameTeamSignedAsRequirement = []string{"xpc_peer_requirement_create_team_identity", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description"}
var rawSyms_NewServiceListener = []string{"xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_listener_activate", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_send_message", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_NewSharedMemory = []string{"xpc_shmem_create"}
var rawSyms_PeerRequirementFromHandle = []string{"xpc_retain"}
var rawSyms_RegisterActivity = []string{"xpc_activity_register", "xpc_array_append_value", "xpc_array_create_empty", "xpc_bool_create", "xpc_data_create", "xpc_date_create", "xpc_dictionary_create_empty", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_int64_create", "xpc_null_create", "xpc_release", "xpc_retain", "xpc_string_create", "xpc_uint64_create", "xpc_uuid_create"}
var rawSyms_SetEventStreamHandler = []string{"xpc_array_apply", "xpc_bool_get_value", "xpc_copy_description", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_double_get_value", "xpc_get_type", "xpc_int64_get_value", "xpc_retain", "xpc_set_event_stream_handler", "xpc_string_get_string_ptr", "xpc_type_get_name", "xpc_uint64_get_value", "xpc_uuid_get_bytes"}
var rawSyms_Transaction = []string{"xpc_transaction_begin", "xpc_transaction_end"}
var rawSyms_UnregisterActivity = []string{"xpc_activity_unregister"}

// rawReachability indexes the above by entry point, for the tests.
var rawReachability = map[string][]string{
	"(*Activity).Close":                                 rawSyms_Activity_Close,
	"(*Activity).Criteria":                              rawSyms_Activity_Criteria,
	"(*Activity).ShouldDefer":                           rawSyms_Activity_ShouldDefer,
	"(*Connection).Activate":                            rawSyms_Connection_Activate,
	"(*Connection).AuditSessionID":                      rawSyms_Connection_AuditSessionID,
	"(*Connection).CallDictionary":                      rawSyms_Connection_CallDictionary,
	"(*Connection).Cancel":                              rawSyms_Connection_Cancel,
	"(*Connection).EGID":                                rawSyms_Connection_EGID,
	"(*Connection).EUID":                                rawSyms_Connection_EUID,
	"(*Connection).Endpoint":                            rawSyms_Connection_Endpoint,
	"(*Connection).InvalidationReason":                  rawSyms_Connection_InvalidationReason,
	"(*Connection).Name":                                rawSyms_Connection_Name,
	"(*Connection).PID":                                 rawSyms_Connection_PID,
	"(*Connection).Resume":                              rawSyms_Connection_Resume,
	"(*Connection).SendDictionary":                      rawSyms_Connection_SendDictionary,
	"(*Connection).SetPeerCodeSigningRequirement":       rawSyms_Connection_SetPeerCodeSigningRequirement,
	"(*Connection).SetPeerEntitlementExistsRequirement": rawSyms_Connection_SetPeerEntitlementExistsRequirement,
	"(*Connection).SetPeerPlatformIdentityRequirement":  rawSyms_Connection_SetPeerPlatformIdentityRequirement,
	"(*Connection).SetPeerRequirement":                  rawSyms_Connection_SetPeerRequirement,
	"(*Connection).SetPeerTeamIdentityRequirement":      rawSyms_Connection_SetPeerTeamIdentityRequirement,
	"(*Connection).SetTargetQueue":                      rawSyms_Connection_SetTargetQueue,
	"(*Connection).Suspend":                             rawSyms_Connection_Suspend,
	"(*FileDescriptor).Close":                           rawSyms_FileDescriptor_Close,
	"(*FileDescriptor).Dup":                             rawSyms_FileDescriptor_Dup,
	"(*Listener).Activate":                              rawSyms_Listener_Activate,
	"(*Listener).Cancel":                                rawSyms_Listener_Cancel,
	"(*Listener).SetPeerCodeSigningRequirement":         rawSyms_Listener_SetPeerCodeSigningRequirement,
	"(*Listener).String":                                rawSyms_Listener_String,
	"(*PeerRequirement).Close":                          rawSyms_PeerRequirement_Close,
	"(*Session).Activate":                               rawSyms_Session_Activate,
	"(*Session).Call":                                   rawSyms_Session_Call,
	"(*Session).CallDictionary":                         rawSyms_Session_CallDictionary,
	"(*Session).Cancel":                                 rawSyms_Session_Cancel,
	"(*Session).Notify":                                 rawSyms_Session_Notify,
	"(*Session).NotifyDictionary":                       rawSyms_Session_NotifyDictionary,
	"(*Session).SetCancellationHandler":                 rawSyms_Session_SetCancellationHandler,
	"(*Session).SetIncomingMessageHandler":              rawSyms_Session_SetIncomingMessageHandler,
	"(*Session).SetPeerCodeSigningRequirement":          rawSyms_Session_SetPeerCodeSigningRequirement,
	"(*Session).SetPeerRequirement":                     rawSyms_Session_SetPeerRequirement,
	"(*Session).SetTargetQueue":                         rawSyms_Session_SetTargetQueue,
	"(*Session).String":                                 rawSyms_Session_String,
	"(*SharedMemory).Close":                             rawSyms_SharedMemory_Close,
	"(*SharedMemory).Map":                               rawSyms_SharedMemory_Map,
	"(ReceivedMessage).Decode":                          rawSyms_ReceivedMessage_Decode,
	"(ReceivedMessage).Dictionary":                      rawSyms_ReceivedMessage_Dictionary,
	"(ReceivedMessage).SenderSatisfies":                 rawSyms_ReceivedMessage_SenderSatisfies,
	"ActivateSocket":                                    rawSyms_ActivateSocket,
	"CopyValue":                                         rawSyms_CopyValue,
	"CurrentDate":                                       rawSyms_CurrentDate,
	"DialMachService":                                   rawSyms_DialMachService,
	"DialXPCService":                                    rawSyms_DialXPCService,
	"Equal":                                             rawSyms_Equal,
	"Hash":                                              rawSyms_Hash,
	"NewAnonymousConnection":                            rawSyms_NewAnonymousConnection,
	"NewAnonymousListener":                              rawSyms_NewAnonymousListener,
	"NewConnectionFromEndpoint":                         rawSyms_NewConnectionFromEndpoint,
	"NewEntitlementExistsRequirement":                   rawSyms_NewEntitlementExistsRequirement,
	"NewEntitlementMatchesRequirement":                  rawSyms_NewEntitlementMatchesRequirement,
	"NewFileDescriptor":                                 rawSyms_NewFileDescriptor,
	"NewLightweightCodeRequirement":                     rawSyms_NewLightweightCodeRequirement,
	"NewMachServiceConnection":                          rawSyms_NewMachServiceConnection,
	"NewPlatformBinaryRequirement":                      rawSyms_NewPlatformBinaryRequirement,
	"NewPlatformBinarySignedAsRequirement":              rawSyms_NewPlatformBinarySignedAsRequirement,
	"NewSameTeamRequirement":                            rawSyms_NewSameTeamRequirement,
	"NewSameTeamSignedAsRequirement":                    rawSyms_NewSameTeamSignedAsRequirement,
	"NewServiceListener":                                rawSyms_NewServiceListener,
	"NewSharedMemory":                                   rawSyms_NewSharedMemory,
	"PeerRequirementFromHandle":                         rawSyms_PeerRequirementFromHandle,
	"RegisterActivity":                                  rawSyms_RegisterActivity,
	"SetEventStreamHandler":                             rawSyms_SetEventStreamHandler,
	"Transaction":                                       rawSyms_Transaction,
	"UnregisterActivity":                                rawSyms_UnregisterActivity,
}

// rawAllCalledSymbols is the union of raw symbols named by any function in
// the package, exported or not.
var rawAllCalledSymbols = []string{"launch_activate_socket", "xpc_activity_copy_criteria", "xpc_activity_register", "xpc_activity_should_defer", "xpc_activity_unregister", "xpc_array_append_value", "xpc_array_apply", "xpc_array_create_empty", "xpc_bool_create", "xpc_bool_get_value", "xpc_connection_activate", "xpc_connection_cancel", "xpc_connection_copy_invalidation_reason", "xpc_connection_create", "xpc_connection_create_from_endpoint", "xpc_connection_create_mach_service", "xpc_connection_get_asid", "xpc_connection_get_egid", "xpc_connection_get_euid", "xpc_connection_get_name", "xpc_connection_get_pid", "xpc_connection_resume", "xpc_connection_send_message", "xpc_connection_send_message_with_reply_sync", "xpc_connection_set_event_handler", "xpc_connection_set_peer_code_signing_requirement", "xpc_connection_set_peer_entitlement_exists_requirement", "xpc_connection_set_peer_platform_identity_requirement", "xpc_connection_set_peer_requirement", "xpc_connection_set_peer_team_identity_requirement", "xpc_connection_set_target_queue", "xpc_connection_suspend", "xpc_copy", "xpc_copy_description", "xpc_data_create", "xpc_data_get_bytes_ptr", "xpc_data_get_length", "xpc_date_create", "xpc_date_create_from_current", "xpc_date_get_value", "xpc_dictionary_apply", "xpc_dictionary_create_empty", "xpc_dictionary_create_reply", "xpc_dictionary_set_bool", "xpc_dictionary_set_data", "xpc_dictionary_set_double", "xpc_dictionary_set_int64", "xpc_dictionary_set_string", "xpc_dictionary_set_uint64", "xpc_dictionary_set_value", "xpc_double_create", "xpc_double_get_value", "xpc_endpoint_create", "xpc_equal", "xpc_fd_create", "xpc_fd_dup", "xpc_get_type", "xpc_hash", "xpc_int64_create", "xpc_int64_get_value", "xpc_listener_activate", "xpc_listener_cancel", "xpc_listener_copy_description", "xpc_listener_create", "xpc_listener_reject_peer", "xpc_listener_set_peer_code_signing_requirement", "xpc_listener_set_peer_requirement", "xpc_null_create", "xpc_peer_requirement_create_entitlement_exists", "xpc_peer_requirement_create_entitlement_matches_value", "xpc_peer_requirement_create_lwcr", "xpc_peer_requirement_create_platform_identity", "xpc_peer_requirement_create_team_identity", "xpc_peer_requirement_match_received_message", "xpc_release", "xpc_retain", "xpc_rich_error_can_retry", "xpc_rich_error_copy_description", "xpc_session_activate", "xpc_session_cancel", "xpc_session_copy_description", "xpc_session_create_mach_service", "xpc_session_create_xpc_service", "xpc_session_send_message", "xpc_session_send_message_with_reply_async", "xpc_session_send_message_with_reply_sync", "xpc_session_set_cancel_handler", "xpc_session_set_incoming_message_handler", "xpc_session_set_peer_code_signing_requirement", "xpc_session_set_peer_requirement", "xpc_session_set_target_queue", "xpc_set_event_stream_handler", "xpc_shmem_create", "xpc_shmem_map", "xpc_string_create", "xpc_string_get_string_ptr", "xpc_transaction_begin", "xpc_transaction_end", "xpc_type_get_name", "xpc_uint64_create", "xpc_uint64_get_value", "xpc_uuid_create", "xpc_uuid_get_bytes"}

// rawReachEdge is one call edge the generator could not follow.
type rawReachEdge struct {
	Entry, File, Expr, Kind string
	Line                    int
}

// rawReachUnresolved lists call edges the generator could not follow. Each
// entry is a place where a reachable set above may be an UNDER-estimate.
// It is emitted, not dropped.
var rawReachUnresolved = []rawReachEdge{
	{Entry: "(RichError).Error", File: "xpc.highlevel.gen.go", Line: 267, Expr: "e.cause.Error(...)", Kind: "unresolvable receiver expression"},
	{Entry: "NewSharedMemory", File: "xpc.highlevel.gen.go", Line: 186, Expr: "m.Close(...)", Kind: "unresolved receiver type"},
	{Entry: "RegisterActivity", File: "xpc.highlevel.gen.go", Line: 348, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "SetEventStreamHandler", File: "xpc.highlevel.gen.go", Line: 755, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "decodeJSONPayload", File: "xpc.highlevel.gen.go", Line: 1074, Expr: "dec.Decode(...)", Kind: "unresolved receiver type"},
	{Entry: "encodeMessage", File: "xpc.highlevel.gen.go", Line: 1185, Expr: "iter.Key(...).String(...)", Kind: "unresolvable receiver expression"},
	{Entry: "entitlementValueToRawObject", File: "xpc.highlevel.gen.go", Line: 964, Expr: "rv.String(...)", Kind: "unresolved receiver type"},
	{Entry: "freeMemory", File: "xpc.highlevel.gen.go", Line: 2513, Expr: "libcfn_free(...)", Kind: "func-value or unknown callee"},
	{Entry: "jsonNumbersToWire", File: "xpc.highlevel.gen.go", Line: 1089, Expr: "x.String(...)", Kind: "unresolved receiver type"},
	{Entry: "mapShared", File: "xpc.highlevel.gen.go", Line: 2490, Expr: "libcfn_mmap(...)", Kind: "func-value or unknown callee"},
	{Entry: "munmapRegion", File: "xpc.highlevel.gen.go", Line: 2505, Expr: "libcfn_munmap(...)", Kind: "func-value or unknown callee"},
	{Entry: "newConnection", File: "xpc.highlevel.gen.go", Line: 532, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "newListener", File: "xpc.highlevel.gen.go", Line: 1328, Expr: "incoming(...)", Kind: "func-value or unknown callee"},
	{Entry: "newRequirement", File: "xpc.highlevel.gen.go", Line: 931, Expr: "create(...)", Kind: "func-value or unknown callee"},
	{Entry: "targetQueuePointer", File: "xpc.highlevel.gen.go", Line: 487, Expr: "queue.Handle(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcCancelTrampoline", File: "xpc.highlevel.gen.go", Line: 2882, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2847, Expr: "handler(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2855, Expr: "err.Error(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcIncomingTrampoline", File: "xpc.highlevel.gen.go", Line: 2860, Expr: "encErr.Error(...)", Kind: "unresolved receiver type"},
	{Entry: "xpcReplyTrampoline", File: "xpc.highlevel.gen.go", Line: 2957, Expr: "reply(...)", Kind: "func-value or unknown callee"},
	{Entry: "xpcReplyTrampoline", File: "xpc.highlevel.gen.go", Line: 2961, Expr: "reply(...)", Kind: "func-value or unknown callee"},
}
