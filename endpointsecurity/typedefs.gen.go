// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"github.com/tmc/apple/kernel"
)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_cdhash_t
type EsCdhash = kernel.Pointer

// EsClient is an opaque type that stores the Endpoint Security client state.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_client_t
type EsClient = kernel.Pointer

// See: https://developer.apple.com/documentation/EndpointSecurity/es_graphical_session_id_t
type EsGraphicalSessionID = uint32

// EsHandlerBlock is a block that handles a message received from Endpoint Security.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_handler_block_t
type EsHandlerBlock = func(*Es_client_t, *Es_message_t)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sha256_t
type EsSha256 = kernel.Pointer

// EsStatfs is this typedef is no longer used, but exists for API backwards compatibility.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_statfs_t
type EsStatfs = kernel.Pointer

// Es_action_type_t is a C-name alias for EsActionType.
type Es_action_type_t = EsActionType

// Es_address_type_t is a C-name alias for EsAddressType.
type Es_address_type_t = EsAddressType

// Es_auth_result_t is a C-name alias for EsAuthResult.
type Es_auth_result_t = EsAuthResult

// Es_authentication_type_t is a C-name alias for EsAuthenticationType.
type Es_authentication_type_t = EsAuthenticationType

// Es_authorization_rule_class_t is a C-name alias for EsAuthorizationRuleClass.
type Es_authorization_rule_class_t = EsAuthorizationRuleClass

// Es_auto_unlock_type_t is a C-name alias for EsAutoUnlockType.
type Es_auto_unlock_type_t = EsAutoUnlockType

// Es_btm_item_type_t is a C-name alias for EsBtmItemType.
type Es_btm_item_type_t = EsBtmItemType

// Es_cdhash_t is a C-name alias for EsCdhash.
type Es_cdhash_t = EsCdhash

// Es_clear_cache_result_t is a C-name alias for EsClearCacheResult.
type Es_clear_cache_result_t = EsClearCacheResult

// Es_client_t is a C-name alias for EsClient.
type Es_client_t = EsClient

// Es_cs_validation_category_t is a C-name alias for EsCsValidationCategory.
type Es_cs_validation_category_t = EsCsValidationCategory

// Es_destination_type_t is a C-name alias for EsDestinationType.
type Es_destination_type_t = EsDestinationType

// Es_event_type_t is a C-name alias for EsEventType.
type Es_event_type_t = EsEventType

// Es_gatekeeper_user_override_file_type_t is a C-name alias for EsGatekeeperUserOverrideFileType.
type Es_gatekeeper_user_override_file_type_t = EsGatekeeperUserOverrideFileType

// Es_get_task_type_t is a C-name alias for EsGetTaskType.
type Es_get_task_type_t = EsGetTaskType

// Es_graphical_session_id_t is a C-name alias for EsGraphicalSessionID.
type Es_graphical_session_id_t = EsGraphicalSessionID

// Es_handler_block_t is a C-name alias for EsHandlerBlock.
type Es_handler_block_t = EsHandlerBlock

// Es_mount_disposition_t is a C-name alias for EsMountDisposition.
type Es_mount_disposition_t = EsMountDisposition

// Es_mute_inversion_type_t is a C-name alias for EsMuteInversionType.
type Es_mute_inversion_type_t = EsMuteInversionType

// Es_mute_inverted_return_t is a C-name alias for EsMuteInvertedReturn.
type Es_mute_inverted_return_t = EsMuteInvertedReturn

// Es_mute_path_type_t is a C-name alias for EsMutePathType.
type Es_mute_path_type_t = EsMutePathType

// Es_new_client_result_t is a C-name alias for EsNewClientResult.
type Es_new_client_result_t = EsNewClientResult

// Es_od_account_type_t is a C-name alias for EsOdAccountType.
type Es_od_account_type_t = EsOdAccountType

// Es_od_member_type_t is a C-name alias for EsOdMemberType.
type Es_od_member_type_t = EsOdMemberType

// Es_od_record_type_t is a C-name alias for EsOdRecordType.
type Es_od_record_type_t = EsOdRecordType

// Es_openssh_login_result_type_t is a C-name alias for EsOpensshLoginResultType.
type Es_openssh_login_result_type_t = EsOpensshLoginResultType

// Es_proc_check_type_t is a C-name alias for EsProcCheckType.
type Es_proc_check_type_t = EsProcCheckType

// Es_proc_suspend_resume_type_t is a C-name alias for EsProcSuspendResumeType.
type Es_proc_suspend_resume_type_t = EsProcSuspendResumeType

// Es_profile_source_t is a C-name alias for EsProfileSource.
type Es_profile_source_t = EsProfileSource

// Es_respond_result_t is a C-name alias for EsRespondResult.
type Es_respond_result_t = EsRespondResult

// Es_result_type_t is a C-name alias for EsResultType.
type Es_result_type_t = EsResultType

// Es_return_t is a C-name alias for EsReturn.
type Es_return_t = EsReturn

// Es_set_or_clear_t is a C-name alias for EsSetOrClear.
type Es_set_or_clear_t = EsSetOrClear

// Es_sha256_t is a C-name alias for EsSha256.
type Es_sha256_t = EsSha256

// Es_statfs_t is a C-name alias for EsStatfs.
type Es_statfs_t = EsStatfs

// Es_sudo_plugin_type_t is a C-name alias for EsSudoPluginType.
type Es_sudo_plugin_type_t = EsSudoPluginType

// Es_tcc_authorization_reason_t is a C-name alias for EsTccAuthorizationReason.
type Es_tcc_authorization_reason_t = EsTccAuthorizationReason

// Es_tcc_authorization_right_t is a C-name alias for EsTccAuthorizationRight.
type Es_tcc_authorization_right_t = EsTccAuthorizationRight

// Es_tcc_event_type_t is a C-name alias for EsTccEventType.
type Es_tcc_event_type_t = EsTccEventType

// Es_tcc_identity_type_t is a C-name alias for EsTccIdentityType.
type Es_tcc_identity_type_t = EsTccIdentityType

// Es_touchid_mode_t is a C-name alias for EsTouchidMode.
type Es_touchid_mode_t = EsTouchidMode

// Es_xpc_domain_type_t is a C-name alias for EsXPCDomainType.
type Es_xpc_domain_type_t = EsXPCDomainType
