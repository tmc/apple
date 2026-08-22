// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// C struct types

// EsAuthorizationResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_authorization_result_t
type EsAuthorizationResult struct {
	Right_name EsStringToken
	Rule_class EsAuthorizationRuleClass
	Granted    bool
}

// Es_authorization_result_t is a type alias for EsAuthorizationResult for use in objc.Send[T] calls.
type Es_authorization_result_t = EsAuthorizationResult

// EsBtmLaunchItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_btm_launch_item_t
type EsBtmLaunchItem struct {
	Item_type EsBtmItemType
	Legacy    bool
	Managed   bool
	Uid       uint32
	Item_url  EsStringToken
	App_url   EsStringToken
}

// Es_btm_launch_item_t is a type alias for EsBtmLaunchItem for use in objc.Send[T] calls.
type Es_btm_launch_item_t = EsBtmLaunchItem

// EsEventAccess - A type for an event that indicates the checking of a file’s access permission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_access_t
type EsEventAccess struct {
	Mode     int32     // The file access permission to check.
	Target   *EsFile   // The file to check for access.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_access_t is a type alias for EsEventAccess for use in objc.Send[T] calls.
type Es_event_access_t = EsEventAccess

// EsEventAuthenticationAutoUnlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_auto_unlock_t
type EsEventAuthenticationAutoUnlock struct {
	Username EsStringToken
	Type     EsAutoUnlockType
}

// Es_event_authentication_auto_unlock_t is a type alias for EsEventAuthenticationAutoUnlock for use in objc.Send[T] calls.
type Es_event_authentication_auto_unlock_t = EsEventAuthenticationAutoUnlock

// EsEventAuthenticationOd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_od_t
type EsEventAuthenticationOd struct {
	Instigator       *EsProcess
	Record_type      EsStringToken
	Record_name      EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_authentication_od_t is a type alias for EsEventAuthenticationOd for use in objc.Send[T] calls.
type Es_event_authentication_od_t = EsEventAuthenticationOd

// EsEventAuthentication
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_t
type EsEventAuthentication struct {
	Success bool
	Type    EsAuthenticationType
	Data    [1]uint64
}

// Es_event_authentication_t is a type alias for EsEventAuthentication for use in objc.Send[T] calls.
type Es_event_authentication_t = EsEventAuthentication

// EsEventAuthenticationToken
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_token_t
type EsEventAuthenticationToken struct {
	Instigator         *EsProcess
	Pubkey_hash        EsStringToken
	Token_id           EsStringToken
	Kerberos_principal EsStringToken
	Instigator_token   [32]byte
}

// Es_event_authentication_token_t is a type alias for EsEventAuthenticationToken for use in objc.Send[T] calls.
type Es_event_authentication_token_t = EsEventAuthenticationToken

// EsEventAuthenticationTouchid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_touchid_t
type EsEventAuthenticationTouchid struct {
	Instigator       *EsProcess
	Touchid_mode     EsTouchidMode
	Has_uid          bool
	Uid              [1]uint32
	Instigator_token [32]byte
}

// Es_event_authentication_touchid_t is a type alias for EsEventAuthenticationTouchid for use in objc.Send[T] calls.
type Es_event_authentication_touchid_t = EsEventAuthenticationTouchid

// EsEventAuthorizationJudgement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authorization_judgement_t
type EsEventAuthorizationJudgement struct {
	Instigator       *EsProcess
	Petitioner       *EsProcess
	Return_code      int32
	Result_count     uintptr
	Results          *EsAuthorizationResult
	Instigator_token [32]byte
	Petitioner_token [32]byte
}

// Es_event_authorization_judgement_t is a type alias for EsEventAuthorizationJudgement for use in objc.Send[T] calls.
type Es_event_authorization_judgement_t = EsEventAuthorizationJudgement

// EsEventAuthorizationPetition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authorization_petition_t
type EsEventAuthorizationPetition struct {
	Instigator       *EsProcess
	Petitioner       *EsProcess
	Flags            uint32
	Right_count      uintptr
	Rights           *EsStringToken
	Instigator_token [32]byte
	Petitioner_token [32]byte
}

// Es_event_authorization_petition_t is a type alias for EsEventAuthorizationPetition for use in objc.Send[T] calls.
type Es_event_authorization_petition_t = EsEventAuthorizationPetition

// EsEventBtmLaunchItemAdd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_btm_launch_item_add_t
type EsEventBtmLaunchItemAdd struct {
	Instigator       *EsProcess
	App              *EsProcess
	Item             *EsBtmLaunchItem
	Executable_path  EsStringToken
	Instigator_token *[32]byte
	App_token        *[32]byte
}

// Es_event_btm_launch_item_add_t is a type alias for EsEventBtmLaunchItemAdd for use in objc.Send[T] calls.
type Es_event_btm_launch_item_add_t = EsEventBtmLaunchItemAdd

// EsEventBtmLaunchItemRemove
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_btm_launch_item_remove_t
type EsEventBtmLaunchItemRemove struct {
	Instigator       *EsProcess
	App              *EsProcess
	Item             *EsBtmLaunchItem
	Instigator_token *[32]byte
	App_token        *[32]byte
}

// Es_event_btm_launch_item_remove_t is a type alias for EsEventBtmLaunchItemRemove for use in objc.Send[T] calls.
type Es_event_btm_launch_item_remove_t = EsEventBtmLaunchItemRemove

// EsEventChdir - A type for an event that indicates a change to a process’s working directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_chdir_t
type EsEventChdir struct {
	Target   *EsFile   // The new current working directory.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_chdir_t is a type alias for EsEventChdir for use in objc.Send[T] calls.
type Es_event_chdir_t = EsEventChdir

// EsEventChroot - A type for an event that indicates a change to a process’s root directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_chroot_t
type EsEventChroot struct {
	Target   *EsFile   // The new root directory.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_chroot_t is a type alias for EsEventChroot for use in objc.Send[T] calls.
type Es_event_chroot_t = EsEventChroot

// EsEventClone - A type for an event that indicates the cloning of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_clone_t
type EsEventClone struct {
	Source      *EsFile       // The file to clone.
	Target_dir  *EsFile       // The directory that contains the cloned file.
	Target_name EsStringToken // The name of the newly cloned file.
	Reserved    [64]uint8     // An unused field reserved for future use.

}

// Es_event_clone_t is a type alias for EsEventClone for use in objc.Send[T] calls.
type Es_event_clone_t = EsEventClone

// EsEventClose - A type for an event that indicates the closing of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_close_t
type EsEventClose struct {
	Modified bool    // A Boolean value that indicates whether the file has modifications.
	Target   *EsFile // The file to close.
	Reserved [64]uint8
}

// Es_event_close_t is a type alias for EsEventClose for use in objc.Send[T] calls.
type Es_event_close_t = EsEventClose

// EsEventCopyfile - A type for an event that indicates the copying of a file by use of a system call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_copyfile_t
type EsEventCopyfile struct {
	Source      *EsFile       // The file to clone.
	Target_file *EsFile       // The file, if any, that exists at the target location.
	Target_dir  *EsFile       // The directory that contains the copied file.
	Target_name EsStringToken // The name of the newly copied file.
	Mode        uint16        // The mode argument of the system call.
	Flags       int32         // The flags argument of the system call.
	Reserved    [56]uint8     // An unused field reserved for future use.

}

// Es_event_copyfile_t is a type alias for EsEventCopyfile for use in objc.Send[T] calls.
type Es_event_copyfile_t = EsEventCopyfile

// EsEventCreate - A type for an event that indicates the creation of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_create_t
type EsEventCreate struct {
	Destination_type EsDestinationType // The type of destination for the event, which can be either an existing file or information that describes a new file’s pending location.
	Destination      [4]uint64         // The file system destination of the created file.
	Reserved2        [16]uint8         // An unused field reserved for future use.
	Reserved         [48]uint8
	New_path         unsafe.Pointer
}

// Es_event_create_t is a type alias for EsEventCreate for use in objc.Send[T] calls.
type Es_event_create_t = EsEventCreate

// EsEventCsInvalidated - A type for an event that indicates the invalidation of a process’ code signing status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_cs_invalidated_t
type EsEventCsInvalidated struct {
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_cs_invalidated_t is a type alias for EsEventCsInvalidated for use in objc.Send[T] calls.
type Es_event_cs_invalidated_t = EsEventCsInvalidated

// EsEventDeleteextattr - A type for an event that indicates the deletion of an extended attribute from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_deleteextattr_t
type EsEventDeleteextattr struct {
	Target   *EsFile       // The file containing extended attributes to delete.
	Extattr  EsStringToken // The extended attribute to delete.
	Reserved [64]uint8     // An unused field reserved for future use.

}

// Es_event_deleteextattr_t is a type alias for EsEventDeleteextattr for use in objc.Send[T] calls.
type Es_event_deleteextattr_t = EsEventDeleteextattr

// EsEventDup - A type for an event that indicates the duplication of a file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_dup_t
type EsEventDup struct {
	Target   *EsFile   // The file that the duplicated file descriptor points to.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_dup_t is a type alias for EsEventDup for use in objc.Send[T] calls.
type Es_event_dup_t = EsEventDup

// EsEventExchangedata - A type for an event that indicates the exchange of data between two files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exchangedata_t
type EsEventExchangedata struct {
	File1    *EsFile   // The first file involved in the data exchange.
	File2    *EsFile   // The second file involved in the data exchange.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_exchangedata_t is a type alias for EsEventExchangedata for use in objc.Send[T] calls.
type Es_event_exchangedata_t = EsEventExchangedata

// EsEventExec - A type for an event that indicates the execution of a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exec_t
type EsEventExec struct {
	Target         *EsProcess // The process to execute.
	Dyld_exec_path EsStringToken
	Reserved       [64]uint8
}

// Es_event_exec_t is a type alias for EsEventExec for use in objc.Send[T] calls.
type Es_event_exec_t = EsEventExec

// EsEventExit - A type for an event that indicates a process exiting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exit_t
type EsEventExit struct {
	Stat     int32     // The exit status of the process.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_exit_t is a type alias for EsEventExit for use in objc.Send[T] calls.
type Es_event_exit_t = EsEventExit

// EsEventFcntl - A type for an event that indicates the manipulation of a file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fcntl_t
type EsEventFcntl struct {
	Target   *EsFile   // The target file to modify.
	Cmd      int32     // The file descriptor modification command.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_fcntl_t is a type alias for EsEventFcntl for use in objc.Send[T] calls.
type Es_event_fcntl_t = EsEventFcntl

// EsEventFileProviderMaterialize - A type for an event that indicates the materialization of a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_file_provider_materialize_t
type EsEventFileProviderMaterialize struct {
	Instigator       *EsProcess // The process that instigated the event.
	Source           *EsFile    // The source file.
	Target           *EsFile    // The target fle.
	Instigator_token [32]byte
	Reserved         [32]uint8 // An unused field reserved for future use.

}

// Es_event_file_provider_materialize_t is a type alias for EsEventFileProviderMaterialize for use in objc.Send[T] calls.
type Es_event_file_provider_materialize_t = EsEventFileProviderMaterialize

// EsEventFileProviderUpdate - A type for an event that indicates an update to a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_file_provider_update_t
type EsEventFileProviderUpdate struct {
	Source      *EsFile       // The source file of the event.
	Target_path EsStringToken // The target path to update.
	Reserved    [64]uint8     // An unused field reserved for future use.

}

// Es_event_file_provider_update_t is a type alias for EsEventFileProviderUpdate for use in objc.Send[T] calls.
type Es_event_file_provider_update_t = EsEventFileProviderUpdate

// EsEventFork - A type for an event that indicates the forking of a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fork_t
type EsEventFork struct {
	Child    *EsProcess // The forked child process.
	Reserved [64]uint8  // An unused field reserved for future use.

}

// Es_event_fork_t is a type alias for EsEventFork for use in objc.Send[T] calls.
type Es_event_fork_t = EsEventFork

// EsEventFsgetpath - A type for an event that indicates the retrieval of a file-system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fsgetpath_t
type EsEventFsgetpath struct {
	Target   *EsFile   // The file-system path of the targeted file.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_fsgetpath_t is a type alias for EsEventFsgetpath for use in objc.Send[T] calls.
type Es_event_fsgetpath_t = EsEventFsgetpath

// EsEventGatekeeperUserOverride
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_gatekeeper_user_override_t
type EsEventGatekeeperUserOverride struct {
	File_type    EsGatekeeperUserOverrideFileType
	File         [2]uint64
	Sha256       *EsSha256
	Signing_info *EsSignedFileInfo
}

// Es_event_gatekeeper_user_override_t is a type alias for EsEventGatekeeperUserOverride for use in objc.Send[T] calls.
type Es_event_gatekeeper_user_override_t = EsEventGatekeeperUserOverride

// EsEventGetTaskInspect - A type for an event that indicates the retrieval of a task’s inspect port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_inspect_t
type EsEventGetTaskInspect struct {
	Target   *EsProcess // The process targeted by this event.
	Type     EsGetTaskType
	Reserved [60]uint8 // An unused field reserved for future use.

}

// Es_event_get_task_inspect_t is a type alias for EsEventGetTaskInspect for use in objc.Send[T] calls.
type Es_event_get_task_inspect_t = EsEventGetTaskInspect

// EsEventGetTaskName - A type for an event that indicates the retrieval of a task’s name port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_name_t
type EsEventGetTaskName struct {
	Target   *EsProcess // The process targeted by this event.
	Type     EsGetTaskType
	Reserved [60]uint8 // An unused field reserved for future use.

}

// Es_event_get_task_name_t is a type alias for EsEventGetTaskName for use in objc.Send[T] calls.
type Es_event_get_task_name_t = EsEventGetTaskName

// EsEventGetTaskRead - A type for an event that indicates the retrieval of a task’s read port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_read_t
type EsEventGetTaskRead struct {
	Target   *EsProcess // The process targeted by this event.
	Type     EsGetTaskType
	Reserved [60]uint8 // An unused field reserved for future use.

}

// Es_event_get_task_read_t is a type alias for EsEventGetTaskRead for use in objc.Send[T] calls.
type Es_event_get_task_read_t = EsEventGetTaskRead

// EsEventGetTask - A type for an event that indicates the retrieval of a task’s control port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_t
type EsEventGetTask struct {
	Target   *EsProcess // The process targeted by this event.
	Type     EsGetTaskType
	Reserved [60]uint8 // An unused field reserved for future use.

}

// Es_event_get_task_t is a type alias for EsEventGetTask for use in objc.Send[T] calls.
type Es_event_get_task_t = EsEventGetTask

// EsEventGetattrlist - A type for an event that indicates the retrieval of attributes from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_getattrlist_t
type EsEventGetattrlist struct {
	Attrlist kernel.Attrlist // The attributes to retrieve, such as volume, directory, file, and fork attributes.
	Target   *EsFile         // The file for which to retrieve attributes.
	Reserved [64]uint8       // An unused field reserved for future use.

}

// Es_event_getattrlist_t is a type alias for EsEventGetattrlist for use in objc.Send[T] calls.
type Es_event_getattrlist_t = EsEventGetattrlist

// EsEventGetextattr - A type for an event that indicates the retrieval of an extended attribute from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_getextattr_t
type EsEventGetextattr struct {
	Target   *EsFile       // The file containing extended attributes to retrieve.
	Extattr  EsStringToken // The extended attribute to retrieve.
	Reserved [64]uint8     // An unused field reserved for future use.

}

// Es_event_getextattr_t is a type alias for EsEventGetextattr for use in objc.Send[T] calls.
type Es_event_getextattr_t = EsEventGetextattr

// EsEventID - An opaque identifier for events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_id_t
type EsEventID struct {
	Reserved [32]uint8 // An opaque value.

}

// Es_event_id_t is a type alias for EsEventID for use in objc.Send[T] calls.
type Es_event_id_t = EsEventID

// EsEventIokitOpen - A type for an event that indicates the opening of an IOKit device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_iokit_open_t
type EsEventIokitOpen struct {
	User_client_type   uint32        // The type of the IOKit client.
	User_client_class  EsStringToken // The name of the IOKit service client.
	Parent_registry_id uint64
	Parent_path        EsStringToken
	Reserved           [40]uint8 // An unused field reserved for future use.

}

// Es_event_iokit_open_t is a type alias for EsEventIokitOpen for use in objc.Send[T] calls.
type Es_event_iokit_open_t = EsEventIokitOpen

// EsEventKextload - A type for an event that indicates the loading of a kernel extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_kextload_t
type EsEventKextload struct {
	Identifier EsStringToken // A string identifying the kernel extension.
	Reserved   [64]uint8     // An unused field reserved for future use.

}

// Es_event_kextload_t is a type alias for EsEventKextload for use in objc.Send[T] calls.
type Es_event_kextload_t = EsEventKextload

// EsEventKextunload - A type for an event that indicates the unloading of a Kernel Extension (KEXT).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_kextunload_t
type EsEventKextunload struct {
	Identifier EsStringToken // A string identifying the kernel extension.
	Reserved   [64]uint8     // An unused field reserved for future use.

}

// Es_event_kextunload_t is a type alias for EsEventKextunload for use in objc.Send[T] calls.
type Es_event_kextunload_t = EsEventKextunload

// EsEventLink - A type for an event that indicates the creation of a hard link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_link_t
type EsEventLink struct {
	Source          *EsFile       // The source file for the link.
	Target_dir      *EsFile       // The directory that contains the newly-created link.
	Target_filename EsStringToken // The file name of the symbolic link.
	Reserved        [64]uint8     // An unused field reserved for future use.

}

// Es_event_link_t is a type alias for EsEventLink for use in objc.Send[T] calls.
type Es_event_link_t = EsEventLink

// EsEventListextattr - A type for an event that indicates the retrieval of multiple extended attributes from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_listextattr_t
type EsEventListextattr struct {
	Target   *EsFile   // The file containing extended attributes to list.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_listextattr_t is a type alias for EsEventListextattr for use in objc.Send[T] calls.
type Es_event_listextattr_t = EsEventListextattr

// EsEventLoginLogin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_login_login_t
type EsEventLoginLogin struct {
	Success         bool
	Failure_message EsStringToken
	Username        EsStringToken
	Has_uid         bool
	Uid             [1]uint32
}

// Es_event_login_login_t is a type alias for EsEventLoginLogin for use in objc.Send[T] calls.
type Es_event_login_login_t = EsEventLoginLogin

// EsEventLoginLogout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_login_logout_t
type EsEventLoginLogout struct {
	Username EsStringToken
	Uid      uint32
}

// Es_event_login_logout_t is a type alias for EsEventLoginLogout for use in objc.Send[T] calls.
type Es_event_login_logout_t = EsEventLoginLogout

// EsEventLookup - A type for an event that indicates the lookup of a file’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lookup_t
type EsEventLookup struct {
	Source_dir      *EsFile       // The source directory to look up.
	Relative_target EsStringToken // The filename to look up.
	Reserved        [64]uint8     // An unused field reserved for future use.

}

// Es_event_lookup_t is a type alias for EsEventLookup for use in objc.Send[T] calls.
type Es_event_lookup_t = EsEventLookup

// EsEventLwSessionLock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_lock_t
type EsEventLwSessionLock struct {
	Username             EsStringToken
	Graphical_session_id EsGraphicalSessionID
}

// Es_event_lw_session_lock_t is a type alias for EsEventLwSessionLock for use in objc.Send[T] calls.
type Es_event_lw_session_lock_t = EsEventLwSessionLock

// EsEventLwSessionLogin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_login_t
type EsEventLwSessionLogin struct {
	Username             EsStringToken
	Graphical_session_id EsGraphicalSessionID
}

// Es_event_lw_session_login_t is a type alias for EsEventLwSessionLogin for use in objc.Send[T] calls.
type Es_event_lw_session_login_t = EsEventLwSessionLogin

// EsEventLwSessionLogout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_logout_t
type EsEventLwSessionLogout struct {
	Username             EsStringToken
	Graphical_session_id EsGraphicalSessionID
}

// Es_event_lw_session_logout_t is a type alias for EsEventLwSessionLogout for use in objc.Send[T] calls.
type Es_event_lw_session_logout_t = EsEventLwSessionLogout

// EsEventLwSessionUnlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_unlock_t
type EsEventLwSessionUnlock struct {
	Username             EsStringToken
	Graphical_session_id EsGraphicalSessionID
}

// Es_event_lw_session_unlock_t is a type alias for EsEventLwSessionUnlock for use in objc.Send[T] calls.
type Es_event_lw_session_unlock_t = EsEventLwSessionUnlock

// EsEventMmap - A type for an event that indicates the mapping of memory to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mmap_t
type EsEventMmap struct {
	Protection     int32     // Options that affect the protection of mapped memory pages.
	Max_protection int32     // The maximum value you can use for protection flags.
	Flags          int32     // Flags that affect the behavior of the memory mapping operation.
	File_pos       uint64    // The offset into the memory-map file.
	Source         *EsFile   // The file to map memory into.
	Reserved       [64]uint8 // An unused field reserved for future use.

}

// Es_event_mmap_t is a type alias for EsEventMmap for use in objc.Send[T] calls.
type Es_event_mmap_t = EsEventMmap

// EsEventMount - A type for an event that indicates the mounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mount_t
type EsEventMount struct {
	Statfs      kernel.Pointer // The statistics of the mounted file system.
	Disposition EsMountDisposition
	Reserved    [60]uint8 // An unused field reserved for future use.

}

// Es_event_mount_t is a type alias for EsEventMount for use in objc.Send[T] calls.
type Es_event_mount_t = EsEventMount

// EsEventMprotect - A type for an event that indicates a change to protection of memory-mapped pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mprotect_t
type EsEventMprotect struct {
	Protection int32     // The protection to apply to the memory-mapped range.
	Address    uint64    // The starting memory address to protect.
	Size       uint64    // The length of the address range to protect.
	Reserved   [64]uint8 // An unused field reserved for future use.

}

// Es_event_mprotect_t is a type alias for EsEventMprotect for use in objc.Send[T] calls.
type Es_event_mprotect_t = EsEventMprotect

// EsEventOdAttributeSet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_set_t
type EsEventOdAttributeSet struct {
	Instigator            *EsProcess
	Error_code            int32
	Record_type           EsOdRecordType
	Record_name           EsStringToken
	Attribute_name        EsStringToken
	Attribute_value_count uintptr
	Attribute_values      *EsStringToken
	Node_name             EsStringToken
	Db_path               EsStringToken
	Instigator_token      [32]byte
}

// Es_event_od_attribute_set_t is a type alias for EsEventOdAttributeSet for use in objc.Send[T] calls.
type Es_event_od_attribute_set_t = EsEventOdAttributeSet

// EsEventOdAttributeValueAdd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_value_add_t
type EsEventOdAttributeValueAdd struct {
	Instigator       *EsProcess
	Error_code       int32
	Record_type      EsOdRecordType
	Record_name      EsStringToken
	Attribute_name   EsStringToken
	Attribute_value  EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_attribute_value_add_t is a type alias for EsEventOdAttributeValueAdd for use in objc.Send[T] calls.
type Es_event_od_attribute_value_add_t = EsEventOdAttributeValueAdd

// EsEventOdAttributeValueRemove
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_value_remove_t
type EsEventOdAttributeValueRemove struct {
	Instigator       *EsProcess
	Error_code       int32
	Record_type      EsOdRecordType
	Record_name      EsStringToken
	Attribute_name   EsStringToken
	Attribute_value  EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_attribute_value_remove_t is a type alias for EsEventOdAttributeValueRemove for use in objc.Send[T] calls.
type Es_event_od_attribute_value_remove_t = EsEventOdAttributeValueRemove

// EsEventOdCreateGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_create_group_t
type EsEventOdCreateGroup struct {
	Instigator       *EsProcess
	Error_code       int32
	Group_name       EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_create_group_t is a type alias for EsEventOdCreateGroup for use in objc.Send[T] calls.
type Es_event_od_create_group_t = EsEventOdCreateGroup

// EsEventOdCreateUser
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_create_user_t
type EsEventOdCreateUser struct {
	Instigator       *EsProcess
	Error_code       int32
	User_name        EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_create_user_t is a type alias for EsEventOdCreateUser for use in objc.Send[T] calls.
type Es_event_od_create_user_t = EsEventOdCreateUser

// EsEventOdDeleteGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_delete_group_t
type EsEventOdDeleteGroup struct {
	Instigator       *EsProcess
	Error_code       int32
	Group_name       EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_delete_group_t is a type alias for EsEventOdDeleteGroup for use in objc.Send[T] calls.
type Es_event_od_delete_group_t = EsEventOdDeleteGroup

// EsEventOdDeleteUser
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_delete_user_t
type EsEventOdDeleteUser struct {
	Instigator       *EsProcess
	Error_code       int32
	User_name        EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_delete_user_t is a type alias for EsEventOdDeleteUser for use in objc.Send[T] calls.
type Es_event_od_delete_user_t = EsEventOdDeleteUser

// EsEventOdDisableUser
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_disable_user_t
type EsEventOdDisableUser struct {
	Instigator       *EsProcess
	Error_code       int32
	User_name        EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_disable_user_t is a type alias for EsEventOdDisableUser for use in objc.Send[T] calls.
type Es_event_od_disable_user_t = EsEventOdDisableUser

// EsEventOdEnableUser
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_enable_user_t
type EsEventOdEnableUser struct {
	Instigator       *EsProcess
	Error_code       int32
	User_name        EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_enable_user_t is a type alias for EsEventOdEnableUser for use in objc.Send[T] calls.
type Es_event_od_enable_user_t = EsEventOdEnableUser

// EsEventOdGroupAdd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_add_t
type EsEventOdGroupAdd struct {
	Instigator       *EsProcess
	Error_code       int32
	Group_name       EsStringToken
	Member           *EsOdMemberID
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_group_add_t is a type alias for EsEventOdGroupAdd for use in objc.Send[T] calls.
type Es_event_od_group_add_t = EsEventOdGroupAdd

// EsEventOdGroupRemove
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_remove_t
type EsEventOdGroupRemove struct {
	Instigator       *EsProcess
	Error_code       int32
	Group_name       EsStringToken
	Member           *EsOdMemberID
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_group_remove_t is a type alias for EsEventOdGroupRemove for use in objc.Send[T] calls.
type Es_event_od_group_remove_t = EsEventOdGroupRemove

// EsEventOdGroupSet
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_set_t
type EsEventOdGroupSet struct {
	Instigator       *EsProcess
	Error_code       int32
	Group_name       EsStringToken
	Members          *EsOdMemberIDArray
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_group_set_t is a type alias for EsEventOdGroupSet for use in objc.Send[T] calls.
type Es_event_od_group_set_t = EsEventOdGroupSet

// EsEventOdModifyPassword
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_modify_password_t
type EsEventOdModifyPassword struct {
	Instigator       *EsProcess
	Error_code       int32
	Account_type     EsOdAccountType
	Account_name     EsStringToken
	Node_name        EsStringToken
	Db_path          EsStringToken
	Instigator_token [32]byte
}

// Es_event_od_modify_password_t is a type alias for EsEventOdModifyPassword for use in objc.Send[T] calls.
type Es_event_od_modify_password_t = EsEventOdModifyPassword

// EsEventOpen - A type for an event that indicates the opening of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_open_t
type EsEventOpen struct {
	Fflag    int32     // The file-opening mask as applied by the kernel.
	File     *EsFile   // The file to open.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_open_t is a type alias for EsEventOpen for use in objc.Send[T] calls.
type Es_event_open_t = EsEventOpen

// EsEventOpensshLogin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_openssh_login_t
type EsEventOpensshLogin struct {
	Success             bool
	Result_type         EsOpensshLoginResultType
	Source_address_type EsAddressType
	Source_address      EsStringToken
	Username            EsStringToken
	Has_uid             bool
	Uid                 [1]uint32
}

// Es_event_openssh_login_t is a type alias for EsEventOpensshLogin for use in objc.Send[T] calls.
type Es_event_openssh_login_t = EsEventOpensshLogin

// EsEventOpensshLogout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_openssh_logout_t
type EsEventOpensshLogout struct {
	Source_address_type EsAddressType
	Source_address      EsStringToken
	Username            EsStringToken
	Uid                 uint32
}

// Es_event_openssh_logout_t is a type alias for EsEventOpensshLogout for use in objc.Send[T] calls.
type Es_event_openssh_logout_t = EsEventOpensshLogout

// EsEventProcCheck - A type that indicates the call used and the data returned when a process checks on the access of the target process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_proc_check_t
type EsEventProcCheck struct {
	Target   *EsProcess      // The process targeted by this event.
	Type     EsProcCheckType // The type of call number used to check the access on the target process.
	Flavor   int32           // A representation of the information sought by a process based on the type member of [es_event_proc_check_t](<https://developer.apple.com/documentation/EndpointSecurity/es_event_proc_check_t>).
	Reserved [64]uint8       // An unused field reserved for future use.

}

// Es_event_proc_check_t is a type alias for EsEventProcCheck for use in objc.Send[T] calls.
type Es_event_proc_check_t = EsEventProcCheck

// EsEventProcSuspendResume - A type for an event that indicates a call to suspend, resume, or shut down sockets for a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_proc_suspend_resume_t
type EsEventProcSuspendResume struct {
	Target   *EsProcess              // The process targeted by this event.
	Type     EsProcSuspendResumeType // The type of event: suspend, resume, or socket shutdown.
	Reserved [64]uint8               // An unused field reserved for future use.

}

// Es_event_proc_suspend_resume_t is a type alias for EsEventProcSuspendResume for use in objc.Send[T] calls.
type Es_event_proc_suspend_resume_t = EsEventProcSuspendResume

// EsEventProfileAdd
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_profile_add_t
type EsEventProfileAdd struct {
	Instigator       *EsProcess
	Is_update        bool
	Profile          *EsProfile
	Instigator_token [32]byte
}

// Es_event_profile_add_t is a type alias for EsEventProfileAdd for use in objc.Send[T] calls.
type Es_event_profile_add_t = EsEventProfileAdd

// EsEventProfileRemove
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_profile_remove_t
type EsEventProfileRemove struct {
	Instigator       *EsProcess
	Profile          *EsProfile
	Instigator_token [32]byte
}

// Es_event_profile_remove_t is a type alias for EsEventProfileRemove for use in objc.Send[T] calls.
type Es_event_profile_remove_t = EsEventProfileRemove

// EsEventPtyClose - A type for an event that indicates the closing of a pseudoterminal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_pty_close_t
type EsEventPtyClose struct {
	Dev      int32     // The major and minor numbers of the device.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_pty_close_t is a type alias for EsEventPtyClose for use in objc.Send[T] calls.
type Es_event_pty_close_t = EsEventPtyClose

// EsEventPtyGrant - A type for an event that indicates the granting of a pseudoterminal device to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_pty_grant_t
type EsEventPtyGrant struct {
	Dev      int32     // The major and minor numbers of the device.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_pty_grant_t is a type alias for EsEventPtyGrant for use in objc.Send[T] calls.
type Es_event_pty_grant_t = EsEventPtyGrant

// EsEventReaddir - A type for an event that indicates the reading of a file-system directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_readdir_t
type EsEventReaddir struct {
	Target   *EsFile   // The directory from which to read contents.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_readdir_t is a type alias for EsEventReaddir for use in objc.Send[T] calls.
type Es_event_readdir_t = EsEventReaddir

// EsEventReadlink - A type for an event that indicates the reading of a symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_readlink_t
type EsEventReadlink struct {
	Source   *EsFile   // The source file pointed to by the link.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_readlink_t is a type alias for EsEventReadlink for use in objc.Send[T] calls.
type Es_event_readlink_t = EsEventReadlink

// EsEventRemoteThreadCreate - A type for an event that indicates an attempt by one process to create a thread in another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_remote_thread_create_t
type EsEventRemoteThreadCreate struct {
	Target       *EsProcess     // The process targeted to spawn a new thread.
	Thread_state *EsThreadState // The new thread’s state.
	Reserved     [64]uint8      // An unused field reserved for future use.

}

// Es_event_remote_thread_create_t is a type alias for EsEventRemoteThreadCreate for use in objc.Send[T] calls.
type Es_event_remote_thread_create_t = EsEventRemoteThreadCreate

// EsEventRemount - A type for an event that indicates the unmounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_remount_t
type EsEventRemount struct {
	Statfs        kernel.Pointer // The statistics of the remounted file system.
	Remount_flags uint64
	Disposition   EsMountDisposition
	Reserved      [52]uint8 // An unused field reserved for future use.

}

// Es_event_remount_t is a type alias for EsEventRemount for use in objc.Send[T] calls.
type Es_event_remount_t = EsEventRemount

// EsEventRename - A type for an event that indicates the renaming of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_rename_t
type EsEventRename struct {
	Source           *EsFile           // The source file to rename.
	Destination_type EsDestinationType // A property that indicates whether the destination is a new path or an existing file.
	Destination      [3]uint64         // The destination of the rename operation.
	Reserved         [64]uint8         // An unused field reserved for future use.
	New_path         unsafe.Pointer
}

// Es_event_rename_t is a type alias for EsEventRename for use in objc.Send[T] calls.
type Es_event_rename_t = EsEventRename

// EsEventScreensharingAttach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_screensharing_attach_t
type EsEventScreensharingAttach struct {
	Success                 bool
	Source_address_type     EsAddressType
	Source_address          EsStringToken
	Viewer_appleid          EsStringToken
	Authentication_type     EsStringToken
	Authentication_username EsStringToken
	Session_username        EsStringToken
	Existing_session        bool
	Graphical_session_id    EsGraphicalSessionID
}

// Es_event_screensharing_attach_t is a type alias for EsEventScreensharingAttach for use in objc.Send[T] calls.
type Es_event_screensharing_attach_t = EsEventScreensharingAttach

// EsEventScreensharingDetach
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_screensharing_detach_t
type EsEventScreensharingDetach struct {
	Source_address_type  EsAddressType
	Source_address       EsStringToken
	Viewer_appleid       EsStringToken
	Graphical_session_id EsGraphicalSessionID
}

// Es_event_screensharing_detach_t is a type alias for EsEventScreensharingDetach for use in objc.Send[T] calls.
type Es_event_screensharing_detach_t = EsEventScreensharingDetach

// EsEventSearchfs - A type for an event that indicates searching a volume or mounted file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_searchfs_t
type EsEventSearchfs struct {
	Attrlist kernel.Attrlist // The attributes used to perform the file system search.
	Target   *EsFile         // The volume to search.
	Reserved [64]uint8       // An unused field reserved for future use.

}

// Es_event_searchfs_t is a type alias for EsEventSearchfs for use in objc.Send[T] calls.
type Es_event_searchfs_t = EsEventSearchfs

// EsEventSetacl - A type for an event that indicates the setting of a file’s access control list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setacl_t
type EsEventSetacl struct {
	Target       *EsFile      // The file containing the access control list to set or clear.
	Set_or_clear EsSetOrClear // The access control list action represented by the event, either setting or clearing values.
	Acl          [1]uint64    // A union containing a settable access control list structure.
	Reserved     [64]uint8    // An unused field reserved for future use.

}

// Es_event_setacl_t is a type alias for EsEventSetacl for use in objc.Send[T] calls.
type Es_event_setacl_t = EsEventSetacl

// EsEventSetattrlist - A type for an event that indicates the setting of a file attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setattrlist_t
type EsEventSetattrlist struct {
	Attrlist kernel.Attrlist // The attributes to set, such as volume, directory, file, and fork attributes.
	Target   *EsFile         // The source file of this event.
	Reserved [64]uint8       // An unused field reserved for future use.

}

// Es_event_setattrlist_t is a type alias for EsEventSetattrlist for use in objc.Send[T] calls.
type Es_event_setattrlist_t = EsEventSetattrlist

// EsEventSetegid - A type for an event that indicates the setting of a process’s effective group ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setegid_t
type EsEventSetegid struct {
	Egid     uint32    // The effective group ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setegid_t is a type alias for EsEventSetegid for use in objc.Send[T] calls.
type Es_event_setegid_t = EsEventSetegid

// EsEventSeteuid - A type for an event that indicates the setting of a process’s effective user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_seteuid_t
type EsEventSeteuid struct {
	Euid     uint32    // The effective user ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_seteuid_t is a type alias for EsEventSeteuid for use in objc.Send[T] calls.
type Es_event_seteuid_t = EsEventSeteuid

// EsEventSetextattr - A type for an event that indicates the setting of a file’s extended attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setextattr_t
type EsEventSetextattr struct {
	Target   *EsFile       // The file containing extended attributes to set.
	Extattr  EsStringToken // The extended attribute.
	Reserved [64]uint8     // An unused field reserved for future use.

}

// Es_event_setextattr_t is a type alias for EsEventSetextattr for use in objc.Send[T] calls.
type Es_event_setextattr_t = EsEventSetextattr

// EsEventSetflags - A type for an event that indicates the setting of a file’s flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setflags_t
type EsEventSetflags struct {
	Flags    uint32    // The flags to set on the file.
	Target   *EsFile   // The source file of this event.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setflags_t is a type alias for EsEventSetflags for use in objc.Send[T] calls.
type Es_event_setflags_t = EsEventSetflags

// EsEventSetgid - A type for an event that indicates the setting of a process’s group ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setgid_t
type EsEventSetgid struct {
	Gid      uint32    // The group ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setgid_t is a type alias for EsEventSetgid for use in objc.Send[T] calls.
type Es_event_setgid_t = EsEventSetgid

// EsEventSetmode - A type for an event that indicates the setting of a file’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setmode_t
type EsEventSetmode struct {
	Mode     uint16    // The mode to set on the file.
	Target   *EsFile   // The source file of the event.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setmode_t is a type alias for EsEventSetmode for use in objc.Send[T] calls.
type Es_event_setmode_t = EsEventSetmode

// EsEventSetowner - A type for an event that indicates the setting of a file’s owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setowner_t
type EsEventSetowner struct {
	Uid      uint32    // The user identifier to set.
	Gid      uint32    // The group identifier to set.
	Target   *EsFile   // The file with ownership metadata to set.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setowner_t is a type alias for EsEventSetowner for use in objc.Send[T] calls.
type Es_event_setowner_t = EsEventSetowner

// EsEventSetregid - A type for an event that indicates the setting of a process’s real and effective group IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setregid_t
type EsEventSetregid struct {
	Rgid     uint32    // The real group ID.
	Egid     uint32    // The effective group ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setregid_t is a type alias for EsEventSetregid for use in objc.Send[T] calls.
type Es_event_setregid_t = EsEventSetregid

// EsEventSetreuid - A type for an event that indicates the setting of a process’s real and effective user IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setreuid_t
type EsEventSetreuid struct {
	Ruid     uint32    // The real user ID.
	Euid     uint32    // The effective user ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setreuid_t is a type alias for EsEventSetreuid for use in objc.Send[T] calls.
type Es_event_setreuid_t = EsEventSetreuid

// EsEventSettime - A type for an event that indicates the modification of the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_settime_t
type EsEventSettime struct {
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_settime_t is a type alias for EsEventSettime for use in objc.Send[T] calls.
type Es_event_settime_t = EsEventSettime

// EsEventSetuid - A type for an event that indicates the setting of a process’s user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setuid_t
type EsEventSetuid struct {
	Uid      uint32    // The user ID.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_setuid_t is a type alias for EsEventSetuid for use in objc.Send[T] calls.
type Es_event_setuid_t = EsEventSetuid

// EsEventSignal - A type for an event that indicates the sending of a signal to a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_signal_t
type EsEventSignal struct {
	Sig        int32      // The signal number sent to the target process.
	Target     *EsProcess // The process that the signal targets.
	Instigator *EsProcess
	Reserved   [56]uint8 // An unused field reserved for future use.

}

// Es_event_signal_t is a type alias for EsEventSignal for use in objc.Send[T] calls.
type Es_event_signal_t = EsEventSignal

// EsEventStat - A type for an event that indicates the retrieval of a file’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_stat_t
type EsEventStat struct {
	Target   *EsFile   // The file with status to retrieve.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_stat_t is a type alias for EsEventStat for use in objc.Send[T] calls.
type Es_event_stat_t = EsEventStat

// EsEventSu
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_su_t
type EsEventSu struct {
	Success         bool
	Failure_message EsStringToken
	From_uid        uint32
	From_username   EsStringToken
	Has_to_uid      bool
	To_uid          [1]uint32
	To_username     EsStringToken
	Shell           EsStringToken
	Argc            uintptr
	Argv            *EsStringToken
	Env_count       uintptr
	Env             *EsStringToken
}

// Es_event_su_t is a type alias for EsEventSu for use in objc.Send[T] calls.
type Es_event_su_t = EsEventSu

// EsEventSudo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_sudo_t
type EsEventSudo struct {
	Success       bool
	Reject_info   *EsSudoRejectInfo
	Has_from_uid  bool
	From_uid      [1]uint32
	From_username EsStringToken
	Has_to_uid    bool
	To_uid        [1]uint32
	To_username   EsStringToken
	Command       EsStringToken
}

// Es_event_sudo_t is a type alias for EsEventSudo for use in objc.Send[T] calls.
type Es_event_sudo_t = EsEventSudo

// EsEventTccModify
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_tcc_modify_t
type EsEventTccModify struct {
	Service           EsStringToken
	Identity          EsStringToken
	Identity_type     EsTccIdentityType // es_tcc_identity_type_t
	Update_type       EsTccEventType
	Instigator_token  [32]byte
	Instigator        *EsProcess
	Responsible_token *[32]byte
	Responsible       *EsProcess
	Right             EsTccAuthorizationRight  // ess_tcc_authorization_right_t
	Reason            EsTccAuthorizationReason // ess_tcc_authorization_reason_t

}

// Es_event_tcc_modify_t is a type alias for EsEventTccModify for use in objc.Send[T] calls.
type Es_event_tcc_modify_t = EsEventTccModify

// EsEventTrace - A type for an event that indicates an attempt by one process to attach to another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_trace_t
type EsEventTrace struct {
	Target   *EsProcess // The process receiving the attach.
	Reserved [64]uint8  // An unused field reserved for future use.

}

// Es_event_trace_t is a type alias for EsEventTrace for use in objc.Send[T] calls.
type Es_event_trace_t = EsEventTrace

// EsEventTruncate - A type for an event that indicates the truncation of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_truncate_t
type EsEventTruncate struct {
	Target   *EsFile   // The source file of this event.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_truncate_t is a type alias for EsEventTruncate for use in objc.Send[T] calls.
type Es_event_truncate_t = EsEventTruncate

// EsEventUipcBind - A type for an event that indicates the binding of a socket to a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_uipc_bind_t
type EsEventUipcBind struct {
	Dir      *EsFile       // The directory containing the socket file.
	Filename EsStringToken // The name of the socket file.
	Mode     uint16        // The mode of the socket file.
	Reserved [64]uint8     // An unused field reserved for future use.

}

// Es_event_uipc_bind_t is a type alias for EsEventUipcBind for use in objc.Send[T] calls.
type Es_event_uipc_bind_t = EsEventUipcBind

// EsEventUipcConnect - A type for an event that indicates the connection of a socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_uipc_connect_t
type EsEventUipcConnect struct {
	File     *EsFile   // The socket file bound to the socket.
	Domain   int32     // The communications domain of the socket.
	Type     int32     // The type of the socket.
	Protocol int32     // The protocol of the socket.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_uipc_connect_t is a type alias for EsEventUipcConnect for use in objc.Send[T] calls.
type Es_event_uipc_connect_t = EsEventUipcConnect

// EsEventUnlink - A type for an event that indicates the deletion of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_unlink_t
type EsEventUnlink struct {
	Target     *EsFile   // The file to unlink.
	Parent_dir *EsFile   // The directory that contains the file to unlink.
	Reserved   [64]uint8 // An unused field reserved for future use.

}

// Es_event_unlink_t is a type alias for EsEventUnlink for use in objc.Send[T] calls.
type Es_event_unlink_t = EsEventUnlink

// EsEventUnmount - A type for an event that indicates the unmounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_unmount_t
type EsEventUnmount struct {
	Statfs   kernel.Pointer // The statistics of the unmounted file system.
	Reserved [64]uint8      // An unused field reserved for future use.

}

// Es_event_unmount_t is a type alias for EsEventUnmount for use in objc.Send[T] calls.
type Es_event_unmount_t = EsEventUnmount

// EsEventUtimes - A type for an event that indicates a change to a file’s access time or modification time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_utimes_t
type EsEventUtimes struct {
	Target   *EsFile         // The file with time metadata to modify.
	Atime    kernel.Timespec // The new last-accessed time.
	Mtime    kernel.Timespec // The new last-modified time.
	Reserved [64]uint8       // An unused field reserved for future use.

}

// Es_event_utimes_t is a type alias for EsEventUtimes for use in objc.Send[T] calls.
type Es_event_utimes_t = EsEventUtimes

// EsEventWrite - A type for an event that indicates the writing of data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_write_t
type EsEventWrite struct {
	Target   *EsFile   // The source file of the event.
	Reserved [64]uint8 // An unused field reserved for future use.

}

// Es_event_write_t is a type alias for EsEventWrite for use in objc.Send[T] calls.
type Es_event_write_t = EsEventWrite

// EsEventXpMalwareDetected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xp_malware_detected_t
type EsEventXpMalwareDetected struct {
	Signature_version   EsStringToken
	Malware_identifier  EsStringToken
	Incident_identifier EsStringToken
	Detected_path       EsStringToken
	Detected_executable EsStringToken
}

// Es_event_xp_malware_detected_t is a type alias for EsEventXpMalwareDetected for use in objc.Send[T] calls.
type Es_event_xp_malware_detected_t = EsEventXpMalwareDetected

// EsEventXpMalwareRemediated
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xp_malware_remediated_t
type EsEventXpMalwareRemediated struct {
	Signature_version              EsStringToken
	Malware_identifier             EsStringToken
	Incident_identifier            EsStringToken
	Action_type                    EsStringToken
	Success                        bool
	Result_description             EsStringToken
	Remediated_path                EsStringToken
	Remediated_process_audit_token *[32]byte
}

// Es_event_xp_malware_remediated_t is a type alias for EsEventXpMalwareRemediated for use in objc.Send[T] calls.
type Es_event_xp_malware_remediated_t = EsEventXpMalwareRemediated

// EsEventXPCConnect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xpc_connect_t
type EsEventXPCConnect struct {
	Service_name        EsStringToken
	Service_domain_type EsXPCDomainType
}

// Es_event_xpc_connect_t is a type alias for EsEventXPCConnect for use in objc.Send[T] calls.
type Es_event_xpc_connect_t = EsEventXPCConnect

// EsEvents is a C union type. A C union of event-specific types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_events_t
type EsEvents [13]uint64

// Access returns the union interpreted as *EsEventAccess.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Access() *EsEventAccess {
	return (*EsEventAccess)(unsafe.Pointer(u))
}

// Chdir returns the union interpreted as *EsEventChdir.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Chdir() *EsEventChdir {
	return (*EsEventChdir)(unsafe.Pointer(u))
}

// Chroot returns the union interpreted as *EsEventChroot.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Chroot() *EsEventChroot {
	return (*EsEventChroot)(unsafe.Pointer(u))
}

// Clone returns the union interpreted as *EsEventClone.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Clone() *EsEventClone {
	return (*EsEventClone)(unsafe.Pointer(u))
}

// Close returns the union interpreted as *EsEventClose.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Close() *EsEventClose {
	return (*EsEventClose)(unsafe.Pointer(u))
}

// Copyfile returns the union interpreted as *EsEventCopyfile.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Copyfile() *EsEventCopyfile {
	return (*EsEventCopyfile)(unsafe.Pointer(u))
}

// Create returns the union interpreted as *EsEventCreate.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Create() *EsEventCreate {
	return (*EsEventCreate)(unsafe.Pointer(u))
}

// Cs_invalidated returns the union interpreted as *EsEventCsInvalidated.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Cs_invalidated() *EsEventCsInvalidated {
	return (*EsEventCsInvalidated)(unsafe.Pointer(u))
}

// Deleteextattr returns the union interpreted as *EsEventDeleteextattr.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Deleteextattr() *EsEventDeleteextattr {
	return (*EsEventDeleteextattr)(unsafe.Pointer(u))
}

// Dup returns the union interpreted as *EsEventDup.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Dup() *EsEventDup {
	return (*EsEventDup)(unsafe.Pointer(u))
}

// Exchangedata returns the union interpreted as *EsEventExchangedata.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Exchangedata() *EsEventExchangedata {
	return (*EsEventExchangedata)(unsafe.Pointer(u))
}

// Exec returns the union interpreted as *EsEventExec.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Exec() *EsEventExec {
	return (*EsEventExec)(unsafe.Pointer(u))
}

// Exit returns the union interpreted as *EsEventExit.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Exit() *EsEventExit {
	return (*EsEventExit)(unsafe.Pointer(u))
}

// File_provider_materialize returns the union interpreted as *EsEventFileProviderMaterialize.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) File_provider_materialize() *EsEventFileProviderMaterialize {
	return (*EsEventFileProviderMaterialize)(unsafe.Pointer(u))
}

// File_provider_update returns the union interpreted as *EsEventFileProviderUpdate.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) File_provider_update() *EsEventFileProviderUpdate {
	return (*EsEventFileProviderUpdate)(unsafe.Pointer(u))
}

// Fcntl returns the union interpreted as *EsEventFcntl.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Fcntl() *EsEventFcntl {
	return (*EsEventFcntl)(unsafe.Pointer(u))
}

// Fork returns the union interpreted as *EsEventFork.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Fork() *EsEventFork {
	return (*EsEventFork)(unsafe.Pointer(u))
}

// Fsgetpath returns the union interpreted as *EsEventFsgetpath.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Fsgetpath() *EsEventFsgetpath {
	return (*EsEventFsgetpath)(unsafe.Pointer(u))
}

// Get_task returns the union interpreted as *EsEventGetTask.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Get_task() *EsEventGetTask {
	return (*EsEventGetTask)(unsafe.Pointer(u))
}

// Get_task_read returns the union interpreted as *EsEventGetTaskRead.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Get_task_read() *EsEventGetTaskRead {
	return (*EsEventGetTaskRead)(unsafe.Pointer(u))
}

// Get_task_inspect returns the union interpreted as *EsEventGetTaskInspect.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Get_task_inspect() *EsEventGetTaskInspect {
	return (*EsEventGetTaskInspect)(unsafe.Pointer(u))
}

// Get_task_name returns the union interpreted as *EsEventGetTaskName.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Get_task_name() *EsEventGetTaskName {
	return (*EsEventGetTaskName)(unsafe.Pointer(u))
}

// Getattrlist returns the union interpreted as *EsEventGetattrlist.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Getattrlist() *EsEventGetattrlist {
	return (*EsEventGetattrlist)(unsafe.Pointer(u))
}

// Getextattr returns the union interpreted as *EsEventGetextattr.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Getextattr() *EsEventGetextattr {
	return (*EsEventGetextattr)(unsafe.Pointer(u))
}

// Iokit_open returns the union interpreted as *EsEventIokitOpen.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Iokit_open() *EsEventIokitOpen {
	return (*EsEventIokitOpen)(unsafe.Pointer(u))
}

// Kextload returns the union interpreted as *EsEventKextload.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Kextload() *EsEventKextload {
	return (*EsEventKextload)(unsafe.Pointer(u))
}

// Kextunload returns the union interpreted as *EsEventKextunload.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Kextunload() *EsEventKextunload {
	return (*EsEventKextunload)(unsafe.Pointer(u))
}

// Link returns the union interpreted as *EsEventLink.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Link() *EsEventLink {
	return (*EsEventLink)(unsafe.Pointer(u))
}

// Listextattr returns the union interpreted as *EsEventListextattr.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Listextattr() *EsEventListextattr {
	return (*EsEventListextattr)(unsafe.Pointer(u))
}

// Lookup returns the union interpreted as *EsEventLookup.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Lookup() *EsEventLookup {
	return (*EsEventLookup)(unsafe.Pointer(u))
}

// Mmap returns the union interpreted as *EsEventMmap.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Mmap() *EsEventMmap {
	return (*EsEventMmap)(unsafe.Pointer(u))
}

// Mount returns the union interpreted as *EsEventMount.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Mount() *EsEventMount {
	return (*EsEventMount)(unsafe.Pointer(u))
}

// Mprotect returns the union interpreted as *EsEventMprotect.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Mprotect() *EsEventMprotect {
	return (*EsEventMprotect)(unsafe.Pointer(u))
}

// Open returns the union interpreted as *EsEventOpen.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Open() *EsEventOpen {
	return (*EsEventOpen)(unsafe.Pointer(u))
}

// Proc_check returns the union interpreted as *EsEventProcCheck.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Proc_check() *EsEventProcCheck {
	return (*EsEventProcCheck)(unsafe.Pointer(u))
}

// Proc_suspend_resume returns the union interpreted as *EsEventProcSuspendResume.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Proc_suspend_resume() *EsEventProcSuspendResume {
	return (*EsEventProcSuspendResume)(unsafe.Pointer(u))
}

// Pty_close returns the union interpreted as *EsEventPtyClose.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Pty_close() *EsEventPtyClose {
	return (*EsEventPtyClose)(unsafe.Pointer(u))
}

// Pty_grant returns the union interpreted as *EsEventPtyGrant.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Pty_grant() *EsEventPtyGrant {
	return (*EsEventPtyGrant)(unsafe.Pointer(u))
}

// Readdir returns the union interpreted as *EsEventReaddir.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Readdir() *EsEventReaddir {
	return (*EsEventReaddir)(unsafe.Pointer(u))
}

// Readlink returns the union interpreted as *EsEventReadlink.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Readlink() *EsEventReadlink {
	return (*EsEventReadlink)(unsafe.Pointer(u))
}

// Remote_thread_create returns the union interpreted as *EsEventRemoteThreadCreate.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Remote_thread_create() *EsEventRemoteThreadCreate {
	return (*EsEventRemoteThreadCreate)(unsafe.Pointer(u))
}

// Remount returns the union interpreted as *EsEventRemount.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Remount() *EsEventRemount {
	return (*EsEventRemount)(unsafe.Pointer(u))
}

// Rename returns the union interpreted as *EsEventRename.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Rename() *EsEventRename {
	return (*EsEventRename)(unsafe.Pointer(u))
}

// Searchfs returns the union interpreted as *EsEventSearchfs.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Searchfs() *EsEventSearchfs {
	return (*EsEventSearchfs)(unsafe.Pointer(u))
}

// Setacl returns the union interpreted as *EsEventSetacl.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setacl() *EsEventSetacl {
	return (*EsEventSetacl)(unsafe.Pointer(u))
}

// Setattrlist returns the union interpreted as *EsEventSetattrlist.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setattrlist() *EsEventSetattrlist {
	return (*EsEventSetattrlist)(unsafe.Pointer(u))
}

// Setextattr returns the union interpreted as *EsEventSetextattr.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setextattr() *EsEventSetextattr {
	return (*EsEventSetextattr)(unsafe.Pointer(u))
}

// Setflags returns the union interpreted as *EsEventSetflags.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setflags() *EsEventSetflags {
	return (*EsEventSetflags)(unsafe.Pointer(u))
}

// Setmode returns the union interpreted as *EsEventSetmode.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setmode() *EsEventSetmode {
	return (*EsEventSetmode)(unsafe.Pointer(u))
}

// Setowner returns the union interpreted as *EsEventSetowner.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setowner() *EsEventSetowner {
	return (*EsEventSetowner)(unsafe.Pointer(u))
}

// Settime returns the union interpreted as *EsEventSettime.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Settime() *EsEventSettime {
	return (*EsEventSettime)(unsafe.Pointer(u))
}

// Setuid returns the union interpreted as *EsEventSetuid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setuid() *EsEventSetuid {
	return (*EsEventSetuid)(unsafe.Pointer(u))
}

// Setgid returns the union interpreted as *EsEventSetgid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setgid() *EsEventSetgid {
	return (*EsEventSetgid)(unsafe.Pointer(u))
}

// Seteuid returns the union interpreted as *EsEventSeteuid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Seteuid() *EsEventSeteuid {
	return (*EsEventSeteuid)(unsafe.Pointer(u))
}

// Setegid returns the union interpreted as *EsEventSetegid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setegid() *EsEventSetegid {
	return (*EsEventSetegid)(unsafe.Pointer(u))
}

// Setreuid returns the union interpreted as *EsEventSetreuid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setreuid() *EsEventSetreuid {
	return (*EsEventSetreuid)(unsafe.Pointer(u))
}

// Setregid returns the union interpreted as *EsEventSetregid.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Setregid() *EsEventSetregid {
	return (*EsEventSetregid)(unsafe.Pointer(u))
}

// Signal returns the union interpreted as *EsEventSignal.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Signal() *EsEventSignal {
	return (*EsEventSignal)(unsafe.Pointer(u))
}

// Stat returns the union interpreted as *EsEventStat.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Stat() *EsEventStat {
	return (*EsEventStat)(unsafe.Pointer(u))
}

// Trace returns the union interpreted as *EsEventTrace.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Trace() *EsEventTrace {
	return (*EsEventTrace)(unsafe.Pointer(u))
}

// Truncate returns the union interpreted as *EsEventTruncate.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Truncate() *EsEventTruncate {
	return (*EsEventTruncate)(unsafe.Pointer(u))
}

// Uipc_bind returns the union interpreted as *EsEventUipcBind.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Uipc_bind() *EsEventUipcBind {
	return (*EsEventUipcBind)(unsafe.Pointer(u))
}

// Uipc_connect returns the union interpreted as *EsEventUipcConnect.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Uipc_connect() *EsEventUipcConnect {
	return (*EsEventUipcConnect)(unsafe.Pointer(u))
}

// Unlink returns the union interpreted as *EsEventUnlink.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Unlink() *EsEventUnlink {
	return (*EsEventUnlink)(unsafe.Pointer(u))
}

// Unmount returns the union interpreted as *EsEventUnmount.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Unmount() *EsEventUnmount {
	return (*EsEventUnmount)(unsafe.Pointer(u))
}

// Utimes returns the union interpreted as *EsEventUtimes.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Utimes() *EsEventUtimes {
	return (*EsEventUtimes)(unsafe.Pointer(u))
}

// Write returns the union interpreted as *EsEventWrite.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Write() *EsEventWrite {
	return (*EsEventWrite)(unsafe.Pointer(u))
}

// Authentication returns the union interpreted as *EsEventAuthentication.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Authentication() *EsEventAuthentication {
	return (*EsEventAuthentication)(unsafe.Pointer(u))
}

// Xp_malware_detected returns the union interpreted as *EsEventXpMalwareDetected.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Xp_malware_detected() *EsEventXpMalwareDetected {
	return (*EsEventXpMalwareDetected)(unsafe.Pointer(u))
}

// Xp_malware_remediated returns the union interpreted as *EsEventXpMalwareRemediated.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Xp_malware_remediated() *EsEventXpMalwareRemediated {
	return (*EsEventXpMalwareRemediated)(unsafe.Pointer(u))
}

// Lw_session_login returns the union interpreted as *EsEventLwSessionLogin.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Lw_session_login() *EsEventLwSessionLogin {
	return (*EsEventLwSessionLogin)(unsafe.Pointer(u))
}

// Lw_session_logout returns the union interpreted as *EsEventLwSessionLogout.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Lw_session_logout() *EsEventLwSessionLogout {
	return (*EsEventLwSessionLogout)(unsafe.Pointer(u))
}

// Lw_session_lock returns the union interpreted as *EsEventLwSessionLock.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Lw_session_lock() *EsEventLwSessionLock {
	return (*EsEventLwSessionLock)(unsafe.Pointer(u))
}

// Lw_session_unlock returns the union interpreted as *EsEventLwSessionUnlock.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Lw_session_unlock() *EsEventLwSessionUnlock {
	return (*EsEventLwSessionUnlock)(unsafe.Pointer(u))
}

// Screensharing_attach returns the union interpreted as *EsEventScreensharingAttach.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Screensharing_attach() *EsEventScreensharingAttach {
	return (*EsEventScreensharingAttach)(unsafe.Pointer(u))
}

// Screensharing_detach returns the union interpreted as *EsEventScreensharingDetach.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Screensharing_detach() *EsEventScreensharingDetach {
	return (*EsEventScreensharingDetach)(unsafe.Pointer(u))
}

// Openssh_login returns the union interpreted as *EsEventOpensshLogin.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Openssh_login() *EsEventOpensshLogin {
	return (*EsEventOpensshLogin)(unsafe.Pointer(u))
}

// Openssh_logout returns the union interpreted as *EsEventOpensshLogout.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Openssh_logout() *EsEventOpensshLogout {
	return (*EsEventOpensshLogout)(unsafe.Pointer(u))
}

// Login_login returns the union interpreted as *EsEventLoginLogin.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Login_login() *EsEventLoginLogin {
	return (*EsEventLoginLogin)(unsafe.Pointer(u))
}

// Login_logout returns the union interpreted as *EsEventLoginLogout.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Login_logout() *EsEventLoginLogout {
	return (*EsEventLoginLogout)(unsafe.Pointer(u))
}

// Btm_launch_item_add returns the union interpreted as *EsEventBtmLaunchItemAdd.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Btm_launch_item_add() *EsEventBtmLaunchItemAdd {
	return (*EsEventBtmLaunchItemAdd)(unsafe.Pointer(u))
}

// Btm_launch_item_remove returns the union interpreted as *EsEventBtmLaunchItemRemove.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Btm_launch_item_remove() *EsEventBtmLaunchItemRemove {
	return (*EsEventBtmLaunchItemRemove)(unsafe.Pointer(u))
}

// Profile_add returns the union interpreted as *EsEventProfileAdd.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Profile_add() *EsEventProfileAdd {
	return (*EsEventProfileAdd)(unsafe.Pointer(u))
}

// Profile_remove returns the union interpreted as *EsEventProfileRemove.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Profile_remove() *EsEventProfileRemove {
	return (*EsEventProfileRemove)(unsafe.Pointer(u))
}

// Su returns the union interpreted as *EsEventSu.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Su() *EsEventSu {
	return (*EsEventSu)(unsafe.Pointer(u))
}

// Authorization_petition returns the union interpreted as *EsEventAuthorizationPetition.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Authorization_petition() *EsEventAuthorizationPetition {
	return (*EsEventAuthorizationPetition)(unsafe.Pointer(u))
}

// Authorization_judgement returns the union interpreted as *EsEventAuthorizationJudgement.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Authorization_judgement() *EsEventAuthorizationJudgement {
	return (*EsEventAuthorizationJudgement)(unsafe.Pointer(u))
}

// Sudo returns the union interpreted as *EsEventSudo.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Sudo() *EsEventSudo {
	return (*EsEventSudo)(unsafe.Pointer(u))
}

// Od_group_add returns the union interpreted as *EsEventOdGroupAdd.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_group_add() *EsEventOdGroupAdd {
	return (*EsEventOdGroupAdd)(unsafe.Pointer(u))
}

// Od_group_remove returns the union interpreted as *EsEventOdGroupRemove.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_group_remove() *EsEventOdGroupRemove {
	return (*EsEventOdGroupRemove)(unsafe.Pointer(u))
}

// Od_group_set returns the union interpreted as *EsEventOdGroupSet.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_group_set() *EsEventOdGroupSet {
	return (*EsEventOdGroupSet)(unsafe.Pointer(u))
}

// Od_modify_password returns the union interpreted as *EsEventOdModifyPassword.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_modify_password() *EsEventOdModifyPassword {
	return (*EsEventOdModifyPassword)(unsafe.Pointer(u))
}

// Od_disable_user returns the union interpreted as *EsEventOdDisableUser.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_disable_user() *EsEventOdDisableUser {
	return (*EsEventOdDisableUser)(unsafe.Pointer(u))
}

// Od_enable_user returns the union interpreted as *EsEventOdEnableUser.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_enable_user() *EsEventOdEnableUser {
	return (*EsEventOdEnableUser)(unsafe.Pointer(u))
}

// Od_attribute_value_add returns the union interpreted as *EsEventOdAttributeValueAdd.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_attribute_value_add() *EsEventOdAttributeValueAdd {
	return (*EsEventOdAttributeValueAdd)(unsafe.Pointer(u))
}

// Od_attribute_value_remove returns the union interpreted as *EsEventOdAttributeValueRemove.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_attribute_value_remove() *EsEventOdAttributeValueRemove {
	return (*EsEventOdAttributeValueRemove)(unsafe.Pointer(u))
}

// Od_attribute_set returns the union interpreted as *EsEventOdAttributeSet.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_attribute_set() *EsEventOdAttributeSet {
	return (*EsEventOdAttributeSet)(unsafe.Pointer(u))
}

// Od_create_user returns the union interpreted as *EsEventOdCreateUser.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_create_user() *EsEventOdCreateUser {
	return (*EsEventOdCreateUser)(unsafe.Pointer(u))
}

// Od_create_group returns the union interpreted as *EsEventOdCreateGroup.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_create_group() *EsEventOdCreateGroup {
	return (*EsEventOdCreateGroup)(unsafe.Pointer(u))
}

// Od_delete_user returns the union interpreted as *EsEventOdDeleteUser.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_delete_user() *EsEventOdDeleteUser {
	return (*EsEventOdDeleteUser)(unsafe.Pointer(u))
}

// Od_delete_group returns the union interpreted as *EsEventOdDeleteGroup.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Od_delete_group() *EsEventOdDeleteGroup {
	return (*EsEventOdDeleteGroup)(unsafe.Pointer(u))
}

// Xpc_connect returns the union interpreted as *EsEventXPCConnect.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Xpc_connect() *EsEventXPCConnect {
	return (*EsEventXPCConnect)(unsafe.Pointer(u))
}

// Gatekeeper_user_override returns the union interpreted as *EsEventGatekeeperUserOverride.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Gatekeeper_user_override() *EsEventGatekeeperUserOverride {
	return (*EsEventGatekeeperUserOverride)(unsafe.Pointer(u))
}

// Tcc_modify returns the union interpreted as *EsEventTccModify.
// The returned pointer aliases the receiver's memory.
func (u *EsEvents) Tcc_modify() *EsEventTccModify {
	return (*EsEventTccModify)(unsafe.Pointer(u))
}

// Es_events_t is a type alias for EsEvents for use in objc.Send[T] calls.
type Es_events_t = EsEvents

// EsFd - A structure that describes an open file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_fd_t
type EsFd struct {
	Fd      int32  // The file descriptor number.
	Fdtype  uint32 // The file descriptor type, as a libproc type.
	Pipe_id uint64
	Pipe    unsafe.Pointer
}

// Es_fd_t is a type alias for EsFd for use in objc.Send[T] calls.
type Es_fd_t = EsFd

// EsFile - A type that represents a file related to an Endpoint Security event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_file_t
type EsFile struct {
	Path           EsStringToken // The file’s path.
	Path_truncated bool          // A Boolean value that indicates whether Endpoint Security truncated the path string.
	Stat           kernel.Stat   // The file’s metadata, such as file size, user and group identifiers, and access and modification dates.

}

// Es_file_t is a type alias for EsFile for use in objc.Send[T] calls.
type Es_file_t = EsFile

// EsMessage - A message from the Endpoint Security subsystem that describes a security event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_message_t
type EsMessage struct {
	Version        uint32          // The version of the Endpoint Security message.
	Time           kernel.Timespec // The time the event occurred, expressed as a Darwin time value.
	Mach_time      uint64          // The time the event occurred, as a Mach time value.
	Deadline       uint64          // The deadline by which your app must respond to the event.
	Process        *EsProcess      // The process that performed the action defined in a message.
	Seq_num        uint64          // The sequence number of the message.
	Action_type    EsActionType    // The type of action: authentication or notification.
	Action         [9]uint32       // The action monitored by Endpoint Security.
	Event_type     EsEventType     // The type of the message’s event.
	Event          EsEvents        // The event that triggered this message.
	Thread         *EsThread       // The thread that took the action defined in a message.
	Global_seq_num uint64          // The global sequence number of the message.

}

// Es_message_t is a type alias for EsMessage for use in objc.Send[T] calls.
type Es_message_t = EsMessage

// EsMutedPath - A structure that describes a path’s muted events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_path_t
type EsMutedPath struct {
	Type        EsMutePathType // The path type: prefix or literal.
	Event_count uintptr        // The number of elements in the muted events array.
	Events      *EsEventType   // An array containing the muted event types.
	Path        EsStringToken  // The muted path.

}

// Es_muted_path_t is a type alias for EsMutedPath for use in objc.Send[T] calls.
type Es_muted_path_t = EsMutedPath

// EsMutedPaths - A structure for a set of muted paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_t
type EsMutedPaths struct {
	Count uintptr      // The number of elements in the paths array.
	Paths *EsMutedPath // An array containing the muted paths.

}

// Es_muted_paths_t is a type alias for EsMutedPaths for use in objc.Send[T] calls.
type Es_muted_paths_t = EsMutedPaths

// EsMutedProcess - A structure that describes a process’s muted events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_process_t
type EsMutedProcess struct {
	Audit_token [32]byte     // The audit token associated with a muted process.
	Event_count uintptr      // The number of elements in the muted events array.
	Events      *EsEventType // An array containing the muted event types.

}

// Es_muted_process_t is a type alias for EsMutedProcess for use in objc.Send[T] calls.
type Es_muted_process_t = EsMutedProcess

// EsMutedProcesses - A structure for a set of muted processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_t
type EsMutedProcesses struct {
	Count     uintptr         // The number of elements in the processes array.
	Processes *EsMutedProcess // An array containing the muted processes.

}

// Es_muted_processes_t is a type alias for EsMutedProcesses for use in objc.Send[T] calls.
type Es_muted_processes_t = EsMutedProcesses

// EsOdMemberIDArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_od_member_id_array_t
type EsOdMemberIDArray struct {
	Member_type  EsOdMemberType
	Member_count uintptr
	Member_array [1]uint64
}

// Es_od_member_id_array_t is a type alias for EsOdMemberIDArray for use in objc.Send[T] calls.
type Es_od_member_id_array_t = EsOdMemberIDArray

// EsOdMemberID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_od_member_id_t
type EsOdMemberID struct {
	Member_type  EsOdMemberType
	Member_value [2]uint64
}

// Es_od_member_id_t is a type alias for EsOdMemberID for use in objc.Send[T] calls.
type Es_od_member_id_t = EsOdMemberID

// EsProcess - A type that describes a process, as delivered by an Endpoint Security message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_process_t
type EsProcess struct {
	Audit_token             [32]byte               // A token for use with Basic Security Module auditing functions.
	Ppid                    int32                  // The parent process identifier.
	Original_ppid           int32                  // The original parent process ID.
	Group_id                int32                  // The process group identifier.
	Session_id              int32                  // The identifier of the session that contains the process group.
	Codesigning_flags       uint32                 // The flags used to sign the process.
	Is_platform_binary      bool                   // A Boolean value that indicates whether the process is a platform binary.
	Is_es_client            bool                   // A Boolean value that indicates whether the process connects to the Endpoint Security subsystem.
	Cdhash                  [20]uint8              // The code directory hash value.
	Signing_id              EsStringToken          // The identifier used to sign the process.
	Team_id                 EsStringToken          // The team identifier used to sign the process.
	Executable              *EsFile                // The file containing the executed process.
	Tty                     *EsFile                // The TTY associated with the process sending the message.
	Start_time              kernel.Timeval         // The time the process started.
	Responsible_audit_token [32]byte               // The audit token of the process responsible for this process.
	Parent_audit_token      [32]byte               // The audit token of the parent process.
	Cs_validation_category  EsCsValidationCategory // es_cs_validation_category

}

// Es_process_t is a type alias for EsProcess for use in objc.Send[T] calls.
type Es_process_t = EsProcess

// EsProfile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_profile_t
type EsProfile struct {
	Identifier     EsStringToken
	Uuid           EsStringToken
	Install_source EsProfileSource
	Organization   EsStringToken
	Display_name   EsStringToken
	Scope          EsStringToken
}

// Es_profile_t is a type alias for EsProfile for use in objc.Send[T] calls.
type Es_profile_t = EsProfile

// EsResult - The result of the Endpoint Security subsystem authorization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_result_t
type EsResult struct {
	Result_type EsResultType // The type of the message’s result.
	Result      [8]uint32    // The message’s result, as either an authorization result or flags.

}

// Es_result_t is a type alias for EsResult for use in objc.Send[T] calls.
type Es_result_t = EsResult

// EsSignedFileInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_signed_file_info_t
type EsSignedFileInfo struct {
	Cdhash     [20]uint8
	Signing_id EsStringToken
	Team_id    EsStringToken
}

// Es_signed_file_info_t is a type alias for EsSignedFileInfo for use in objc.Send[T] calls.
type Es_signed_file_info_t = EsSignedFileInfo

// EsStringToken - A pointer to a null-terminated string, and the length in bytes of that string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_string_token_t
type EsStringToken struct {
	Length uintptr // The size of the data buffer, in bytes.
	Data   *byte   // The string data.

}

// Es_string_token_t is a type alias for EsStringToken for use in objc.Send[T] calls.
type Es_string_token_t = EsStringToken

// EsSudoRejectInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_sudo_reject_info_t
type EsSudoRejectInfo struct {
	Plugin_name     EsStringToken
	Plugin_type     EsSudoPluginType
	Failure_message EsStringToken
}

// Es_sudo_reject_info_t is a type alias for EsSudoRejectInfo for use in objc.Send[T] calls.
type Es_sudo_reject_info_t = EsSudoRejectInfo

// EsThreadState - A description of a thread’s machine-specfiic state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_thread_state_t
type EsThreadState struct {
	Flavor int32   // An indication of the representation of the machine-specific thread state.
	State  EsToken // The machine-specific thread state.

}

// Es_thread_state_t is a type alias for EsThreadState for use in objc.Send[T] calls.
type Es_thread_state_t = EsThreadState

// EsThread - A structure that represents a thread in a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_thread_t
type EsThread struct {
	Thread_id uint64 // The unique identifier of the thread.

}

// Es_thread_t is a type alias for EsThread for use in objc.Send[T] calls.
type Es_thread_t = EsThread

// EsToken - An arbitrary buffer of data with its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_token_t
type EsToken struct {
	Size uintptr // The size of the data buffer, in bytes.
	Data *uint8  // A data buffer.

}

// Es_token_t is a type alias for EsToken for use in objc.Send[T] calls.
type Es_token_t = EsToken
