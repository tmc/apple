// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"github.com/tmc/apple/kernel"
)

type Es_action_type_t = EsActionType

type Es_address_type_t = EsAddressType

type Es_auth_result_t = EsAuthResult

type Es_authentication_type_t = EsAuthenticationType

type Es_authorization_rule_class_t = EsAuthorizationRuleClass

type Es_auto_unlock_type_t = EsAutoUnlock

type Es_btm_item_type_t = EsBtmItemType

// See: https://developer.apple.com/documentation/EndpointSecurity/es_cdhash_t
type Es_cdhash_t = kernel.Pointer

type Es_clear_cache_result_t = EsClearCacheResult

// Es_client_t is an opaque type that stores the Endpoint Security client state.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_client_t
type Es_client_t = kernel.Pointer

type Es_cs_validation_category_t = EsCsValidationCategory

type Es_destination_type_t = EsDestinationType

type Es_event_type_t = EsEventType

type Es_gatekeeper_user_override_file_type_t = EsGatekeeperUserOverrideFileType

type Es_get_task_type_t = EsGetTaskType

// See: https://developer.apple.com/documentation/EndpointSecurity/es_graphical_session_id_t
type Es_graphical_session_id_t = uint32

// Es_handler_block_t is a block that handles a message received from Endpoint Security.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_handler_block_t
type Es_handler_block_t = func(kernel.Pointer, *Es_message_t)

type Es_mount_disposition_t = EsMountDisposition

type Es_mute_inversion_type_t = EsMuteInversionType

type Es_mute_inverted_return_t = EsMute

type Es_mute_path_type_t = EsMutePathType

type Es_new_client_result_t = EsNewClientResult

type Es_od_account_type_t = EsOdAccountType

type Es_od_member_type_t = EsOdMemberType

type Es_od_record_type_t = EsOdRecordType

type Es_openssh_login_result_type_t = EsOpenssh

type Es_proc_check_type_t = EsProcCheckType

type Es_proc_suspend_resume_type_t = EsProcSuspendResumeType

type Es_profile_source_t = EsProfileSource

type Es_respond_result_t = EsRespondResult

type Es_result_type_t = EsResultType

type Es_return_t = EsReturn

type Es_set_or_clear_t = Es

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sha256_t
type Es_sha256_t = kernel.Pointer

// Es_statfs_t is this typedef is no longer used, but exists for API backwards compatibility.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_statfs_t
type Es_statfs_t = kernel.Pointer

type Es_sudo_plugin_type_t = EsSudoPluginType

type Es_tcc_authorization_reason_t = EsTccAuthorizationReason

type Es_tcc_authorization_right_t = EsTccAuthorizationRight

type Es_tcc_event_type_t = EsTccEventType

type Es_tcc_identity_type_t = EsTccIdentityType

type Es_touchid_mode_t = EsTouchidMode

type Es_xpc_domain_type_t = EsXPCDomainType

// EsAutoUnlockType is a Go-name alias for Es_auto_unlock_type_t.
type EsAutoUnlockType = Es_auto_unlock_type_t

// EsCdhash is a Go-name alias for Es_cdhash_t.
type EsCdhash = Es_cdhash_t

// EsClient is a Go-name alias for Es_client_t.
type EsClient = Es_client_t

// EsGraphicalSessionID is a Go-name alias for Es_graphical_session_id_t.
type EsGraphicalSessionID = Es_graphical_session_id_t

// EsHandlerBlock is a Go-name alias for Es_handler_block_t.
type EsHandlerBlock = Es_handler_block_t

// EsMuteInvertedReturn is a Go-name alias for Es_mute_inverted_return_t.
type EsMuteInvertedReturn = Es_mute_inverted_return_t

// EsOpensshLoginResultType is a Go-name alias for Es_openssh_login_result_type_t.
type EsOpensshLoginResultType = Es_openssh_login_result_type_t

// EsSetOrClear is a Go-name alias for Es_set_or_clear_t.
type EsSetOrClear = Es_set_or_clear_t

// EsSha256 is a Go-name alias for Es_sha256_t.
type EsSha256 = Es_sha256_t

// EsStatfs is a Go-name alias for Es_statfs_t.
type EsStatfs = Es_statfs_t
