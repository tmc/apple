// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"syscall"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// Es_authorization_result_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_authorization_result_t
type Es_authorization_result_t struct {
	Right_name Es_string_token_t
	Rule_class EsAuthorizationRuleClass
	Granted    bool
}

// Es_btm_launch_item_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_btm_launch_item_t
type Es_btm_launch_item_t struct {
	Item_type EsBtmItemType
	Legacy    bool
	Managed   bool
	Uid       uint32
	Item_url  Es_string_token_t
	App_url   Es_string_token_t
}

// Es_event_access_t - A type for an event that indicates the checking of a file’s access permission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_access_t
type Es_event_access_t struct {
	Mode     int32      // The file access permission to check.
	Target   *Es_file_t // The file to check for access.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_authentication_auto_unlock_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_auto_unlock_t
type Es_event_authentication_auto_unlock_t struct {
	Username Es_string_token_t
	Type     Es_auto_unlock_type_t
}

// Es_event_authentication_od_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_od_t
type Es_event_authentication_od_t struct {
	Instigator       *Es_process_t
	Record_type      Es_string_token_t
	Record_name      Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_authentication_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_t
type Es_event_authentication_t struct {
	Success     bool
	Type        EsAuthenticationType
	Data        [8]byte
	Auto_unlock *Es_event_authentication_auto_unlock_t
	Od          *Es_event_authentication_od_t
	Token       *Es_event_authentication_token_t
	Touchid     *Es_event_authentication_touchid_t
}

// Es_event_authentication_token_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_token_t
type Es_event_authentication_token_t struct {
	Instigator         *Es_process_t
	Pubkey_hash        Es_string_token_t
	Token_id           Es_string_token_t
	Kerberos_principal Es_string_token_t
	Instigator_token   [32]byte
}

// Es_event_authentication_touchid_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authentication_touchid_t
type Es_event_authentication_touchid_t struct {
	Instigator       *Es_process_t
	Touchid_mode     EsTouchidMode
	Has_uid          bool
	Instigator_token [32]byte
	Uid              [4]byte
}

// Es_event_authorization_judgement_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authorization_judgement_t
type Es_event_authorization_judgement_t struct {
	Instigator       *Es_process_t
	Petitioner       *Es_process_t
	Return_code      int
	Result_count     uintptr
	Results          *Es_authorization_result_t
	Instigator_token [32]byte
	Petitioner_token [32]byte
}

// Es_event_authorization_petition_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_authorization_petition_t
type Es_event_authorization_petition_t struct {
	Instigator       *Es_process_t
	Petitioner       *Es_process_t
	Flags            uint32
	Right_count      uintptr
	Rights           *Es_string_token_t
	Instigator_token [32]byte
	Petitioner_token [32]byte
}

// Es_event_btm_launch_item_add_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_btm_launch_item_add_t
type Es_event_btm_launch_item_add_t struct {
	Instigator       *Es_process_t
	App              *Es_process_t
	Item             *Es_btm_launch_item_t
	Executable_path  Es_string_token_t
	Instigator_token *[32]byte
	App_token        *[32]byte
}

// Es_event_btm_launch_item_remove_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_btm_launch_item_remove_t
type Es_event_btm_launch_item_remove_t struct {
	Instigator       *Es_process_t
	App              *Es_process_t
	Item             *Es_btm_launch_item_t
	Instigator_token *[32]byte
	App_token        *[32]byte
}

// Es_event_chdir_t - A type for an event that indicates a change to a process’s working directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_chdir_t
type Es_event_chdir_t struct {
	Target   *Es_file_t // The new current working directory.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_chroot_t - A type for an event that indicates a change to a process’s root directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_chroot_t
type Es_event_chroot_t struct {
	Target   *Es_file_t // The new root directory.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_clone_t - A type for an event that indicates the cloning of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_clone_t
type Es_event_clone_t struct {
	Source      *Es_file_t        // The file to clone.
	Target_dir  *Es_file_t        // The directory that contains the cloned file.
	Target_name Es_string_token_t // The name of the newly cloned file.
	Reserved    uint8             // An unused field reserved for future use.

}

// Es_event_close_t - A type for an event that indicates the closing of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_close_t
type Es_event_close_t struct {
	Modified            bool       // A Boolean value that indicates whether the file has modifications.
	Target              *Es_file_t // The file to close.
	Reserved            uint8
	Was_mapped_writable bool
}

// Es_event_copyfile_t - A type for an event that indicates the copying of a file by use of a system call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_copyfile_t
type Es_event_copyfile_t struct {
	Source      *Es_file_t        // The file to clone.
	Target_file *Es_file_t        // The file, if any, that exists at the target location.
	Target_dir  *Es_file_t        // The directory that contains the copied file.
	Target_name Es_string_token_t // The name of the newly copied file.
	Mode        uint16            // The mode argument of the system call.
	Flags       int32             // The flags argument of the system call.
	Reserved    uint8             // An unused field reserved for future use.

}

// Es_event_create_t - A type for an event that indicates the creation of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_create_t
type Es_event_create_t struct {
	Destination_type EsDestinationType // The type of destination for the event, which can be either an existing file or information that describes a new file’s pending location.
	Destination      [32]byte          // The file system destination of the created file.
	Reserved2        uint8             // An unused field reserved for future use.
	Acl              unsafe.Pointer
	Existing_file    *Es_file_t
	New_path         unsafe.Pointer
	Dir              *Es_file_t
	Filename         Es_string_token_t
	Mode             uint16
	Reserved         uint8
}

// Es_event_cs_invalidated_t - A type for an event that indicates the invalidation of a process’ code signing status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_cs_invalidated_t
type Es_event_cs_invalidated_t struct {
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_deleteextattr_t - A type for an event that indicates the deletion of an extended attribute from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_deleteextattr_t
type Es_event_deleteextattr_t struct {
	Target   *Es_file_t        // The file containing extended attributes to delete.
	Extattr  Es_string_token_t // The extended attribute to delete.
	Reserved uint8             // An unused field reserved for future use.

}

// Es_event_dup_t - A type for an event that indicates the duplication of a file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_dup_t
type Es_event_dup_t struct {
	Target   *Es_file_t // The file that the duplicated file descriptor points to.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_exchangedata_t - A type for an event that indicates the exchange of data between two files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exchangedata_t
type Es_event_exchangedata_t struct {
	File1    *Es_file_t // The first file involved in the data exchange.
	File2    *Es_file_t // The second file involved in the data exchange.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_exec_t - A type for an event that indicates the execution of a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exec_t
type Es_event_exec_t struct {
	Target           *Es_process_t // The process to execute.
	Dyld_exec_path   Es_string_token_t
	Cwd              *Es_file_t
	Image_cpusubtype int32
	Image_cputype    int32
	Last_fd          int
	Reserved         uint8
	Script           *Es_file_t
}

// Es_event_exit_t - A type for an event that indicates a process exiting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_exit_t
type Es_event_exit_t struct {
	Stat     int   // The exit status of the process.
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_fcntl_t - A type for an event that indicates the manipulation of a file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fcntl_t
type Es_event_fcntl_t struct {
	Target   *Es_file_t // The target file to modify.
	Cmd      int32      // The file descriptor modification command.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_file_provider_materialize_t - A type for an event that indicates the materialization of a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_file_provider_materialize_t
type Es_event_file_provider_materialize_t struct {
	Instigator       *Es_process_t // The process that instigated the event.
	Source           *Es_file_t    // The source file.
	Target           *Es_file_t    // The target fle.
	Instigator_token [32]byte
	Reserved         uint8 // An unused field reserved for future use.

}

// Es_event_file_provider_update_t - A type for an event that indicates an update to a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_file_provider_update_t
type Es_event_file_provider_update_t struct {
	Source      *Es_file_t        // The source file of the event.
	Target_path Es_string_token_t // The target path to update.
	Reserved    uint8             // An unused field reserved for future use.

}

// Es_event_fork_t - A type for an event that indicates the forking of a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fork_t
type Es_event_fork_t struct {
	Child    *Es_process_t // The forked child process.
	Reserved uint8         // An unused field reserved for future use.

}

// Es_event_fsgetpath_t - A type for an event that indicates the retrieval of a file-system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_fsgetpath_t
type Es_event_fsgetpath_t struct {
	Target   *Es_file_t // The file-system path of the targeted file.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_gatekeeper_user_override_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_gatekeeper_user_override_t
type Es_event_gatekeeper_user_override_t struct {
	File_type    EsGatekeeperUserOverrideFileType
	Sha256       *Es_sha256_t
	Signing_info *Es_signed_file_info_t
	File         [16]byte
	File_path    Es_string_token_t
}

// Es_event_get_task_inspect_t - A type for an event that indicates the retrieval of a task’s inspect port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_inspect_t
type Es_event_get_task_inspect_t struct {
	Target   *Es_process_t // The process targeted by this event.
	Type     EsGetTaskType
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_get_task_name_t - A type for an event that indicates the retrieval of a task’s name port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_name_t
type Es_event_get_task_name_t struct {
	Target   *Es_process_t // The process targeted by this event.
	Type     EsGetTaskType
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_get_task_read_t - A type for an event that indicates the retrieval of a task’s read port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_read_t
type Es_event_get_task_read_t struct {
	Target   *Es_process_t // The process targeted by this event.
	Type     EsGetTaskType
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_get_task_t - A type for an event that indicates the retrieval of a task’s control port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_get_task_t
type Es_event_get_task_t struct {
	Target   *Es_process_t // The process targeted by this event.
	Type     EsGetTaskType
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_getattrlist_t - A type for an event that indicates the retrieval of attributes from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_getattrlist_t
type Es_event_getattrlist_t struct {
	Attrlist kernel.Attrlist // The attributes to retrieve, such as volume, directory, file, and fork attributes.
	Target   *Es_file_t      // The file for which to retrieve attributes.
	Reserved uint8           // An unused field reserved for future use.

}

// Es_event_getextattr_t - A type for an event that indicates the retrieval of an extended attribute from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_getextattr_t
type Es_event_getextattr_t struct {
	Target   *Es_file_t        // The file containing extended attributes to retrieve.
	Extattr  Es_string_token_t // The extended attribute to retrieve.
	Reserved uint8             // An unused field reserved for future use.

}

// Es_event_id_t - An opaque identifier for events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_id_t
type Es_event_id_t struct {
	Reserved uint8 // An opaque value.

}

// Es_event_iokit_open_t - A type for an event that indicates the opening of an IOKit device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_iokit_open_t
type Es_event_iokit_open_t struct {
	User_client_type   uint32            // The type of the IOKit client.
	User_client_class  Es_string_token_t // The name of the IOKit service client.
	Parent_registry_id uint64
	Parent_path        Es_string_token_t
	Reserved           uint8 // An unused field reserved for future use.

}

// Es_event_kextload_t - A type for an event that indicates the loading of a kernel extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_kextload_t
type Es_event_kextload_t struct {
	Identifier Es_string_token_t // A string identifying the kernel extension.
	Reserved   uint8             // An unused field reserved for future use.

}

// Es_event_kextunload_t - A type for an event that indicates the unloading of a Kernel Extension (KEXT).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_kextunload_t
type Es_event_kextunload_t struct {
	Identifier Es_string_token_t // A string identifying the kernel extension.
	Reserved   uint8             // An unused field reserved for future use.

}

// Es_event_link_t - A type for an event that indicates the creation of a hard link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_link_t
type Es_event_link_t struct {
	Source          *Es_file_t        // The source file for the link.
	Target_dir      *Es_file_t        // The directory that contains the newly-created link.
	Target_filename Es_string_token_t // The file name of the symbolic link.
	Reserved        uint8             // An unused field reserved for future use.

}

// Es_event_listextattr_t - A type for an event that indicates the retrieval of multiple extended attributes from a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_listextattr_t
type Es_event_listextattr_t struct {
	Target   *Es_file_t // The file containing extended attributes to list.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_login_login_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_login_login_t
type Es_event_login_login_t struct {
	Success         bool
	Failure_message Es_string_token_t
	Username        Es_string_token_t
	Has_uid         bool
	Uid             [4]byte
}

// Es_event_login_logout_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_login_logout_t
type Es_event_login_logout_t struct {
	Username Es_string_token_t
	Uid      uint32
}

// Es_event_lookup_t - A type for an event that indicates the lookup of a file’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lookup_t
type Es_event_lookup_t struct {
	Source_dir      *Es_file_t        // The source directory to look up.
	Relative_target Es_string_token_t // The filename to look up.
	Reserved        uint8             // An unused field reserved for future use.

}

// Es_event_lw_session_lock_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_lock_t
type Es_event_lw_session_lock_t struct {
	Username             Es_string_token_t
	Graphical_session_id Es_graphical_session_id_t
}

// Es_event_lw_session_login_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_login_t
type Es_event_lw_session_login_t struct {
	Username             Es_string_token_t
	Graphical_session_id Es_graphical_session_id_t
}

// Es_event_lw_session_logout_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_logout_t
type Es_event_lw_session_logout_t struct {
	Username             Es_string_token_t
	Graphical_session_id Es_graphical_session_id_t
}

// Es_event_lw_session_unlock_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_lw_session_unlock_t
type Es_event_lw_session_unlock_t struct {
	Username             Es_string_token_t
	Graphical_session_id Es_graphical_session_id_t
}

// Es_event_mmap_t - A type for an event that indicates the mapping of memory to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mmap_t
type Es_event_mmap_t struct {
	Protection     int32      // Options that affect the protection of mapped memory pages.
	Max_protection int32      // The maximum value you can use for protection flags.
	Flags          int32      // Flags that affect the behavior of the memory mapping operation.
	File_pos       uint64     // The offset into the memory-map file.
	Source         *Es_file_t // The file to map memory into.
	Reserved       uint8      // An unused field reserved for future use.

}

// Es_event_mount_t - A type for an event that indicates the mounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mount_t
type Es_event_mount_t struct {
	Statfs      objectivec.IObject // The statistics of the mounted file system.
	Disposition EsMountDisposition
	Reserved    uint8 // An unused field reserved for future use.

}

// Es_event_mprotect_t - A type for an event that indicates a change to protection of memory-mapped pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_mprotect_t
type Es_event_mprotect_t struct {
	Protection int32              // The protection to apply to the memory-mapped range.
	Address    kernel.User_addr_t // The starting memory address to protect.
	Size       kernel.User_size_t // The length of the address range to protect.
	Reserved   uint8              // An unused field reserved for future use.

}

// Es_event_od_attribute_set_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_set_t
type Es_event_od_attribute_set_t struct {
	Instigator            *Es_process_t
	Error_code            int
	Record_type           EsOdRecordType
	Record_name           Es_string_token_t
	Attribute_name        Es_string_token_t
	Attribute_value_count uintptr
	Attribute_values      *Es_string_token_t
	Node_name             Es_string_token_t
	Db_path               Es_string_token_t
	Instigator_token      [32]byte
}

// Es_event_od_attribute_value_add_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_value_add_t
type Es_event_od_attribute_value_add_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Record_type      EsOdRecordType
	Record_name      Es_string_token_t
	Attribute_name   Es_string_token_t
	Attribute_value  Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_attribute_value_remove_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_attribute_value_remove_t
type Es_event_od_attribute_value_remove_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Record_type      EsOdRecordType
	Record_name      Es_string_token_t
	Attribute_name   Es_string_token_t
	Attribute_value  Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_create_group_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_create_group_t
type Es_event_od_create_group_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Group_name       Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_create_user_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_create_user_t
type Es_event_od_create_user_t struct {
	Instigator       *Es_process_t
	Error_code       int
	User_name        Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_delete_group_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_delete_group_t
type Es_event_od_delete_group_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Group_name       Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_delete_user_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_delete_user_t
type Es_event_od_delete_user_t struct {
	Instigator       *Es_process_t
	Error_code       int
	User_name        Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_disable_user_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_disable_user_t
type Es_event_od_disable_user_t struct {
	Instigator       *Es_process_t
	Error_code       int
	User_name        Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_enable_user_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_enable_user_t
type Es_event_od_enable_user_t struct {
	Instigator       *Es_process_t
	Error_code       int
	User_name        Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_group_add_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_add_t
type Es_event_od_group_add_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Group_name       Es_string_token_t
	Member           *Es_od_member_id_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_group_remove_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_remove_t
type Es_event_od_group_remove_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Group_name       Es_string_token_t
	Member           *Es_od_member_id_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_group_set_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_group_set_t
type Es_event_od_group_set_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Group_name       Es_string_token_t
	Members          *Es_od_member_id_array_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_od_modify_password_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_od_modify_password_t
type Es_event_od_modify_password_t struct {
	Instigator       *Es_process_t
	Error_code       int
	Account_type     EsOdAccountType
	Account_name     Es_string_token_t
	Node_name        Es_string_token_t
	Db_path          Es_string_token_t
	Instigator_token [32]byte
}

// Es_event_open_t - A type for an event that indicates the opening of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_open_t
type Es_event_open_t struct {
	Fflag    int32      // The file-opening mask as applied by the kernel.
	File     *Es_file_t // The file to open.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_openssh_login_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_openssh_login_t
type Es_event_openssh_login_t struct {
	Success             bool
	Result_type         Es_openssh_login_result_type_t
	Source_address_type EsAddressType
	Source_address      Es_string_token_t
	Username            Es_string_token_t
	Has_uid             bool
	Uid                 [4]byte
}

// Es_event_openssh_logout_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_openssh_logout_t
type Es_event_openssh_logout_t struct {
	Source_address_type EsAddressType
	Source_address      Es_string_token_t
	Username            Es_string_token_t
	Uid                 uint32
}

// Es_event_proc_check_t - A type that indicates the call used and the data returned when a process checks on the access of the target process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_proc_check_t
type Es_event_proc_check_t struct {
	Target   *Es_process_t   // The process targeted by this event.
	Type     EsProcCheckType // The type of call number used to check the access on the target process.
	Flavor   int             // A representation of the information sought by a process based on the type member of [es_event_proc_check_t](<doc://com.apple.endpointsecurity/documentation/EndpointSecurity/es_event_proc_check_t>).
	Reserved uint8           // An unused field reserved for future use.

}

// Es_event_proc_suspend_resume_t - A type for an event that indicates a call to suspend, resume, or shut down sockets for a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_proc_suspend_resume_t
type Es_event_proc_suspend_resume_t struct {
	Target   *Es_process_t           // The process targeted by this event.
	Type     EsProcSuspendResumeType // The type of event: suspend, resume, or socket shutdown.
	Reserved uint8                   // An unused field reserved for future use.

}

// Es_event_profile_add_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_profile_add_t
type Es_event_profile_add_t struct {
	Instigator       *Es_process_t
	Is_update        bool
	Profile          *Es_profile_t
	Instigator_token [32]byte
}

// Es_event_profile_remove_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_profile_remove_t
type Es_event_profile_remove_t struct {
	Instigator       *Es_process_t
	Profile          *Es_profile_t
	Instigator_token [32]byte
}

// Es_event_pty_close_t - A type for an event that indicates the closing of a pseudoterminal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_pty_close_t
type Es_event_pty_close_t struct {
	Dev      int32 // The major and minor numbers of the device.
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_pty_grant_t - A type for an event that indicates the granting of a pseudoterminal device to a user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_pty_grant_t
type Es_event_pty_grant_t struct {
	Dev      int32 // The major and minor numbers of the device.
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_readdir_t - A type for an event that indicates the reading of a file-system directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_readdir_t
type Es_event_readdir_t struct {
	Target   *Es_file_t // The directory from which to read contents.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_readlink_t - A type for an event that indicates the reading of a symbolic link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_readlink_t
type Es_event_readlink_t struct {
	Source   *Es_file_t // The source file pointed to by the link.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_remote_thread_create_t - A type for an event that indicates an attempt by one process to create a thread in another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_remote_thread_create_t
type Es_event_remote_thread_create_t struct {
	Target       *Es_process_t      // The process targeted to spawn a new thread.
	Thread_state *Es_thread_state_t // The new thread’s state.
	Reserved     uint8              // An unused field reserved for future use.

}

// Es_event_remount_t - A type for an event that indicates the unmounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_remount_t
type Es_event_remount_t struct {
	Statfs        objectivec.IObject // The statistics of the remounted file system.
	Remount_flags uint64
	Disposition   EsMountDisposition
	Reserved      uint8 // An unused field reserved for future use.

}

// Es_event_rename_t - A type for an event that indicates the renaming of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_rename_t
type Es_event_rename_t struct {
	Source           *Es_file_t        // The source file to rename.
	Destination_type EsDestinationType // A property that indicates whether the destination is a new path or an existing file.
	Destination      [24]byte          // The destination of the rename operation.
	Reserved         uint8             // An unused field reserved for future use.
	Existing_file    *Es_file_t
	New_path         unsafe.Pointer
	Dir              *Es_file_t
	Filename         Es_string_token_t
}

// Es_event_screensharing_attach_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_screensharing_attach_t
type Es_event_screensharing_attach_t struct {
	Success                 bool
	Source_address_type     EsAddressType
	Source_address          Es_string_token_t
	Viewer_appleid          Es_string_token_t
	Authentication_type     Es_string_token_t
	Authentication_username Es_string_token_t
	Session_username        Es_string_token_t
	Existing_session        bool
	Graphical_session_id    Es_graphical_session_id_t
}

// Es_event_screensharing_detach_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_screensharing_detach_t
type Es_event_screensharing_detach_t struct {
	Source_address_type  EsAddressType
	Source_address       Es_string_token_t
	Viewer_appleid       Es_string_token_t
	Graphical_session_id Es_graphical_session_id_t
}

// Es_event_searchfs_t - A type for an event that indicates searching a volume or mounted file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_searchfs_t
type Es_event_searchfs_t struct {
	Attrlist kernel.Attrlist // The attributes used to perform the file system search.
	Target   *Es_file_t      // The volume to search.
	Reserved uint8           // An unused field reserved for future use.

}

// Es_event_setacl_t - A type for an event that indicates the setting of a file’s access control list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setacl_t
type Es_event_setacl_t struct {
	Target       *Es_file_t        // The file containing the access control list to set or clear.
	Set_or_clear Es_set_or_clear_t // The access control list action represented by the event, either setting or clearing values.
	Acl          [8]byte           // A union containing a settable access control list structure.
	Reserved     uint8             // An unused field reserved for future use.
	Set          unsafe.Pointer
}

// Es_event_setattrlist_t - A type for an event that indicates the setting of a file attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setattrlist_t
type Es_event_setattrlist_t struct {
	Attrlist kernel.Attrlist // The attributes to set, such as volume, directory, file, and fork attributes.
	Target   *Es_file_t      // The source file of this event.
	Reserved uint8           // An unused field reserved for future use.

}

// Es_event_setegid_t - A type for an event that indicates the setting of a process’s effective group ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setegid_t
type Es_event_setegid_t struct {
	Egid     uint32 // The effective group ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_seteuid_t - A type for an event that indicates the setting of a process’s effective user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_seteuid_t
type Es_event_seteuid_t struct {
	Euid     uint32 // The effective user ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_setextattr_t - A type for an event that indicates the setting of a file’s extended attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setextattr_t
type Es_event_setextattr_t struct {
	Target   *Es_file_t        // The file containing extended attributes to set.
	Extattr  Es_string_token_t // The extended attribute.
	Reserved uint8             // An unused field reserved for future use.

}

// Es_event_setflags_t - A type for an event that indicates the setting of a file’s flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setflags_t
type Es_event_setflags_t struct {
	Flags    uint32     // The flags to set on the file.
	Target   *Es_file_t // The source file of this event.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_setgid_t - A type for an event that indicates the setting of a process’s group ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setgid_t
type Es_event_setgid_t struct {
	Gid      uint32 // The group ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_setmode_t - A type for an event that indicates the setting of a file’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setmode_t
type Es_event_setmode_t struct {
	Mode     uint16     // The mode to set on the file.
	Target   *Es_file_t // The source file of the event.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_setowner_t - A type for an event that indicates the setting of a file’s owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setowner_t
type Es_event_setowner_t struct {
	Uid      uint32     // The user identifier to set.
	Gid      uint32     // The group identifier to set.
	Target   *Es_file_t // The file with ownership metadata to set.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_setregid_t - A type for an event that indicates the setting of a process’s real and effective group IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setregid_t
type Es_event_setregid_t struct {
	Rgid     uint32 // The real group ID.
	Egid     uint32 // The effective group ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_setreuid_t - A type for an event that indicates the setting of a process’s real and effective user IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setreuid_t
type Es_event_setreuid_t struct {
	Ruid     uint32 // The real user ID.
	Euid     uint32 // The effective user ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_settime_t - A type for an event that indicates the modification of the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_settime_t
type Es_event_settime_t struct {
	Reserved uint8 // An unused field reserved for future use.

}

// Es_event_setuid_t - A type for an event that indicates the setting of a process’s user ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_setuid_t
type Es_event_setuid_t struct {
	Uid      uint32 // The user ID.
	Reserved uint8  // An unused field reserved for future use.

}

// Es_event_signal_t - A type for an event that indicates the sending of a signal to a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_signal_t
type Es_event_signal_t struct {
	Sig        int           // The signal number sent to the target process.
	Target     *Es_process_t // The process that the signal targets.
	Instigator *Es_process_t
	Reserved   uint8 // An unused field reserved for future use.

}

// Es_event_stat_t - A type for an event that indicates the retrieval of a file’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_stat_t
type Es_event_stat_t struct {
	Target   *Es_file_t // The file with status to retrieve.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_su_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_su_t
type Es_event_su_t struct {
	Success         bool
	Failure_message Es_string_token_t
	From_uid        uint32
	From_username   Es_string_token_t
	Has_to_uid      bool
	To_username     Es_string_token_t
	Shell           Es_string_token_t
	Argc            uintptr
	Argv            *Es_string_token_t
	Env_count       uintptr
	Env             *Es_string_token_t
	To_uid          [4]byte
	Uid             uint32
}

// Es_event_sudo_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_sudo_t
type Es_event_sudo_t struct {
	Success       bool
	Reject_info   *Es_sudo_reject_info_t
	Has_from_uid  bool
	From_username Es_string_token_t
	Has_to_uid    bool
	To_username   Es_string_token_t
	Command       Es_string_token_t
	From_uid      [4]byte
	To_uid        [4]byte
	Uid           uint32
}

// Es_event_tcc_modify_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_tcc_modify_t
type Es_event_tcc_modify_t struct {
	Service           Es_string_token_t
	Identity          Es_string_token_t
	Identity_type     EsTccIdentityType // es_tcc_identity_type_t
	Update_type       EsTccEventType
	Instigator_token  [32]byte
	Instigator        *Es_process_t
	Responsible_token *[32]byte
	Responsible       *Es_process_t
	Right             EsTccAuthorizationRight  // ess_tcc_authorization_right_t
	Reason            EsTccAuthorizationReason // ess_tcc_authorization_reason_t

}

// Es_event_trace_t - A type for an event that indicates an attempt by one process to attach to another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_trace_t
type Es_event_trace_t struct {
	Target   *Es_process_t // The process receiving the attach.
	Reserved uint8         // An unused field reserved for future use.

}

// Es_event_truncate_t - A type for an event that indicates the truncation of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_truncate_t
type Es_event_truncate_t struct {
	Target   *Es_file_t // The source file of this event.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_uipc_bind_t - A type for an event that indicates the binding of a socket to a path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_uipc_bind_t
type Es_event_uipc_bind_t struct {
	Dir      *Es_file_t        // The directory containing the socket file.
	Filename Es_string_token_t // The name of the socket file.
	Mode     uint16            // The mode of the socket file.
	Reserved uint8             // An unused field reserved for future use.

}

// Es_event_uipc_connect_t - A type for an event that indicates the connection of a socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_uipc_connect_t
type Es_event_uipc_connect_t struct {
	File     *Es_file_t // The socket file bound to the socket.
	Domain   int        // The communications domain of the socket.
	Type     int        // The type of the socket.
	Protocol int        // The protocol of the socket.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_unlink_t - A type for an event that indicates the deletion of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_unlink_t
type Es_event_unlink_t struct {
	Target     *Es_file_t // The file to unlink.
	Parent_dir *Es_file_t // The directory that contains the file to unlink.
	Reserved   uint8      // An unused field reserved for future use.

}

// Es_event_unmount_t - A type for an event that indicates the unmounting of a file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_unmount_t
type Es_event_unmount_t struct {
	Statfs   objectivec.IObject // The statistics of the unmounted file system.
	Reserved uint8              // An unused field reserved for future use.

}

// Es_event_utimes_t - A type for an event that indicates a change to a file’s access time or modification time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_utimes_t
type Es_event_utimes_t struct {
	Target   *Es_file_t       // The file with time metadata to modify.
	Atime    syscall.Timespec // The new last-accessed time.
	Mtime    syscall.Timespec // The new last-modified time.
	Reserved uint8            // An unused field reserved for future use.

}

// Es_event_write_t - A type for an event that indicates the writing of data to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_write_t
type Es_event_write_t struct {
	Target   *Es_file_t // The source file of the event.
	Reserved uint8      // An unused field reserved for future use.

}

// Es_event_xp_malware_detected_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xp_malware_detected_t
type Es_event_xp_malware_detected_t struct {
	Signature_version   Es_string_token_t
	Malware_identifier  Es_string_token_t
	Incident_identifier Es_string_token_t
	Detected_path       Es_string_token_t
	Detected_executable Es_string_token_t
}

// Es_event_xp_malware_remediated_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xp_malware_remediated_t
type Es_event_xp_malware_remediated_t struct {
	Signature_version              Es_string_token_t
	Malware_identifier             Es_string_token_t
	Incident_identifier            Es_string_token_t
	Action_type                    Es_string_token_t
	Success                        bool
	Result_description             Es_string_token_t
	Remediated_path                Es_string_token_t
	Remediated_process_audit_token *[32]byte
}

// Es_event_xpc_connect_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_event_xpc_connect_t
type Es_event_xpc_connect_t struct {
	Service_name        Es_string_token_t
	Service_domain_type Es_xpc_domain_type_t
}

// Es_events_t is a C union type. A C union of event-specific types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_events_t
type Es_events_t [104]byte

// Access returns the union interpreted as *Es_event_access_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Access() *Es_event_access_t {
	return (*Es_event_access_t)(unsafe.Pointer(u))
}

// Clone returns the union interpreted as *Es_event_clone_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Clone() *Es_event_clone_t {
	return (*Es_event_clone_t)(unsafe.Pointer(u))
}

// Copyfile returns the union interpreted as *Es_event_copyfile_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Copyfile() *Es_event_copyfile_t {
	return (*Es_event_copyfile_t)(unsafe.Pointer(u))
}

// Close returns the union interpreted as *Es_event_close_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Close() *Es_event_close_t {
	return (*Es_event_close_t)(unsafe.Pointer(u))
}

// Create returns the union interpreted as *Es_event_create_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Create() *Es_event_create_t {
	return (*Es_event_create_t)(unsafe.Pointer(u))
}

// Dup returns the union interpreted as *Es_event_dup_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Dup() *Es_event_dup_t {
	return (*Es_event_dup_t)(unsafe.Pointer(u))
}

// Exchangedata returns the union interpreted as *Es_event_exchangedata_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Exchangedata() *Es_event_exchangedata_t {
	return (*Es_event_exchangedata_t)(unsafe.Pointer(u))
}

// Fcntl returns the union interpreted as *Es_event_fcntl_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Fcntl() *Es_event_fcntl_t {
	return (*Es_event_fcntl_t)(unsafe.Pointer(u))
}

// Open returns the union interpreted as *Es_event_open_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Open() *Es_event_open_t {
	return (*Es_event_open_t)(unsafe.Pointer(u))
}

// Rename returns the union interpreted as *Es_event_rename_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Rename() *Es_event_rename_t {
	return (*Es_event_rename_t)(unsafe.Pointer(u))
}

// Write returns the union interpreted as *Es_event_write_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Write() *Es_event_write_t {
	return (*Es_event_write_t)(unsafe.Pointer(u))
}

// Truncate returns the union interpreted as *Es_event_truncate_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Truncate() *Es_event_truncate_t {
	return (*Es_event_truncate_t)(unsafe.Pointer(u))
}

// Lookup returns the union interpreted as *Es_event_lookup_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Lookup() *Es_event_lookup_t {
	return (*Es_event_lookup_t)(unsafe.Pointer(u))
}

// Searchfs returns the union interpreted as *Es_event_searchfs_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Searchfs() *Es_event_searchfs_t {
	return (*Es_event_searchfs_t)(unsafe.Pointer(u))
}

// Deleteextattr returns the union interpreted as *Es_event_deleteextattr_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Deleteextattr() *Es_event_deleteextattr_t {
	return (*Es_event_deleteextattr_t)(unsafe.Pointer(u))
}

// Fsgetpath returns the union interpreted as *Es_event_fsgetpath_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Fsgetpath() *Es_event_fsgetpath_t {
	return (*Es_event_fsgetpath_t)(unsafe.Pointer(u))
}

// Getattrlist returns the union interpreted as *Es_event_getattrlist_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Getattrlist() *Es_event_getattrlist_t {
	return (*Es_event_getattrlist_t)(unsafe.Pointer(u))
}

// Getextattr returns the union interpreted as *Es_event_getextattr_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Getextattr() *Es_event_getextattr_t {
	return (*Es_event_getextattr_t)(unsafe.Pointer(u))
}

// Listextattr returns the union interpreted as *Es_event_listextattr_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Listextattr() *Es_event_listextattr_t {
	return (*Es_event_listextattr_t)(unsafe.Pointer(u))
}

// Readdir returns the union interpreted as *Es_event_readdir_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Readdir() *Es_event_readdir_t {
	return (*Es_event_readdir_t)(unsafe.Pointer(u))
}

// Setacl returns the union interpreted as *Es_event_setacl_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setacl() *Es_event_setacl_t {
	return (*Es_event_setacl_t)(unsafe.Pointer(u))
}

// Setattrlist returns the union interpreted as *Es_event_setattrlist_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setattrlist() *Es_event_setattrlist_t {
	return (*Es_event_setattrlist_t)(unsafe.Pointer(u))
}

// Setextattr returns the union interpreted as *Es_event_setextattr_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setextattr() *Es_event_setextattr_t {
	return (*Es_event_setextattr_t)(unsafe.Pointer(u))
}

// Setflags returns the union interpreted as *Es_event_setflags_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setflags() *Es_event_setflags_t {
	return (*Es_event_setflags_t)(unsafe.Pointer(u))
}

// Setmode returns the union interpreted as *Es_event_setmode_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setmode() *Es_event_setmode_t {
	return (*Es_event_setmode_t)(unsafe.Pointer(u))
}

// Setowner returns the union interpreted as *Es_event_setowner_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setowner() *Es_event_setowner_t {
	return (*Es_event_setowner_t)(unsafe.Pointer(u))
}

// Stat returns the union interpreted as *Es_event_stat_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Stat() *Es_event_stat_t {
	return (*Es_event_stat_t)(unsafe.Pointer(u))
}

// Utimes returns the union interpreted as *Es_event_utimes_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Utimes() *Es_event_utimes_t {
	return (*Es_event_utimes_t)(unsafe.Pointer(u))
}

// File_provider_materialize returns the union interpreted as *Es_event_file_provider_materialize_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) File_provider_materialize() *Es_event_file_provider_materialize_t {
	return (*Es_event_file_provider_materialize_t)(unsafe.Pointer(u))
}

// File_provider_update returns the union interpreted as *Es_event_file_provider_update_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) File_provider_update() *Es_event_file_provider_update_t {
	return (*Es_event_file_provider_update_t)(unsafe.Pointer(u))
}

// Link returns the union interpreted as *Es_event_link_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Link() *Es_event_link_t {
	return (*Es_event_link_t)(unsafe.Pointer(u))
}

// Readlink returns the union interpreted as *Es_event_readlink_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Readlink() *Es_event_readlink_t {
	return (*Es_event_readlink_t)(unsafe.Pointer(u))
}

// Unlink returns the union interpreted as *Es_event_unlink_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Unlink() *Es_event_unlink_t {
	return (*Es_event_unlink_t)(unsafe.Pointer(u))
}

// Mount returns the union interpreted as *Es_event_mount_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Mount() *Es_event_mount_t {
	return (*Es_event_mount_t)(unsafe.Pointer(u))
}

// Unmount returns the union interpreted as *Es_event_unmount_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Unmount() *Es_event_unmount_t {
	return (*Es_event_unmount_t)(unsafe.Pointer(u))
}

// Remount returns the union interpreted as *Es_event_remount_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Remount() *Es_event_remount_t {
	return (*Es_event_remount_t)(unsafe.Pointer(u))
}

// Mmap returns the union interpreted as *Es_event_mmap_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Mmap() *Es_event_mmap_t {
	return (*Es_event_mmap_t)(unsafe.Pointer(u))
}

// Mprotect returns the union interpreted as *Es_event_mprotect_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Mprotect() *Es_event_mprotect_t {
	return (*Es_event_mprotect_t)(unsafe.Pointer(u))
}

// Chdir returns the union interpreted as *Es_event_chdir_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Chdir() *Es_event_chdir_t {
	return (*Es_event_chdir_t)(unsafe.Pointer(u))
}

// Chroot returns the union interpreted as *Es_event_chroot_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Chroot() *Es_event_chroot_t {
	return (*Es_event_chroot_t)(unsafe.Pointer(u))
}

// Exec returns the union interpreted as *Es_event_exec_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Exec() *Es_event_exec_t {
	return (*Es_event_exec_t)(unsafe.Pointer(u))
}

// Fork returns the union interpreted as *Es_event_fork_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Fork() *Es_event_fork_t {
	return (*Es_event_fork_t)(unsafe.Pointer(u))
}

// Proc_check returns the union interpreted as *Es_event_proc_check_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Proc_check() *Es_event_proc_check_t {
	return (*Es_event_proc_check_t)(unsafe.Pointer(u))
}

// Signal returns the union interpreted as *Es_event_signal_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Signal() *Es_event_signal_t {
	return (*Es_event_signal_t)(unsafe.Pointer(u))
}

// Exit returns the union interpreted as *Es_event_exit_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Exit() *Es_event_exit_t {
	return (*Es_event_exit_t)(unsafe.Pointer(u))
}

// Proc_suspend_resume returns the union interpreted as *Es_event_proc_suspend_resume_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Proc_suspend_resume() *Es_event_proc_suspend_resume_t {
	return (*Es_event_proc_suspend_resume_t)(unsafe.Pointer(u))
}

// Trace returns the union interpreted as *Es_event_trace_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Trace() *Es_event_trace_t {
	return (*Es_event_trace_t)(unsafe.Pointer(u))
}

// Remote_thread_create returns the union interpreted as *Es_event_remote_thread_create_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Remote_thread_create() *Es_event_remote_thread_create_t {
	return (*Es_event_remote_thread_create_t)(unsafe.Pointer(u))
}

// Get_task returns the union interpreted as *Es_event_get_task_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Get_task() *Es_event_get_task_t {
	return (*Es_event_get_task_t)(unsafe.Pointer(u))
}

// Get_task_read returns the union interpreted as *Es_event_get_task_read_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Get_task_read() *Es_event_get_task_read_t {
	return (*Es_event_get_task_read_t)(unsafe.Pointer(u))
}

// Get_task_inspect returns the union interpreted as *Es_event_get_task_inspect_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Get_task_inspect() *Es_event_get_task_inspect_t {
	return (*Es_event_get_task_inspect_t)(unsafe.Pointer(u))
}

// Get_task_name returns the union interpreted as *Es_event_get_task_name_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Get_task_name() *Es_event_get_task_name_t {
	return (*Es_event_get_task_name_t)(unsafe.Pointer(u))
}

// Setuid returns the union interpreted as *Es_event_setuid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setuid() *Es_event_setuid_t {
	return (*Es_event_setuid_t)(unsafe.Pointer(u))
}

// Setgid returns the union interpreted as *Es_event_setgid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setgid() *Es_event_setgid_t {
	return (*Es_event_setgid_t)(unsafe.Pointer(u))
}

// Seteuid returns the union interpreted as *Es_event_seteuid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Seteuid() *Es_event_seteuid_t {
	return (*Es_event_seteuid_t)(unsafe.Pointer(u))
}

// Setegid returns the union interpreted as *Es_event_setegid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setegid() *Es_event_setegid_t {
	return (*Es_event_setegid_t)(unsafe.Pointer(u))
}

// Setreuid returns the union interpreted as *Es_event_setreuid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setreuid() *Es_event_setreuid_t {
	return (*Es_event_setreuid_t)(unsafe.Pointer(u))
}

// Setregid returns the union interpreted as *Es_event_setregid_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Setregid() *Es_event_setregid_t {
	return (*Es_event_setregid_t)(unsafe.Pointer(u))
}

// Cs_invalidated returns the union interpreted as *Es_event_cs_invalidated_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Cs_invalidated() *Es_event_cs_invalidated_t {
	return (*Es_event_cs_invalidated_t)(unsafe.Pointer(u))
}

// Uipc_bind returns the union interpreted as *Es_event_uipc_bind_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Uipc_bind() *Es_event_uipc_bind_t {
	return (*Es_event_uipc_bind_t)(unsafe.Pointer(u))
}

// Uipc_connect returns the union interpreted as *Es_event_uipc_connect_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Uipc_connect() *Es_event_uipc_connect_t {
	return (*Es_event_uipc_connect_t)(unsafe.Pointer(u))
}

// Settime returns the union interpreted as *Es_event_settime_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Settime() *Es_event_settime_t {
	return (*Es_event_settime_t)(unsafe.Pointer(u))
}

// Iokit_open returns the union interpreted as *Es_event_iokit_open_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Iokit_open() *Es_event_iokit_open_t {
	return (*Es_event_iokit_open_t)(unsafe.Pointer(u))
}

// Kextload returns the union interpreted as *Es_event_kextload_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Kextload() *Es_event_kextload_t {
	return (*Es_event_kextload_t)(unsafe.Pointer(u))
}

// Kextunload returns the union interpreted as *Es_event_kextunload_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Kextunload() *Es_event_kextunload_t {
	return (*Es_event_kextunload_t)(unsafe.Pointer(u))
}

// Pty_close returns the union interpreted as *Es_event_pty_close_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Pty_close() *Es_event_pty_close_t {
	return (*Es_event_pty_close_t)(unsafe.Pointer(u))
}

// Pty_grant returns the union interpreted as *Es_event_pty_grant_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Pty_grant() *Es_event_pty_grant_t {
	return (*Es_event_pty_grant_t)(unsafe.Pointer(u))
}

// Authentication returns the union interpreted as *Es_event_authentication_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Authentication() *Es_event_authentication_t {
	return (*Es_event_authentication_t)(unsafe.Pointer(u))
}

// Authorization_judgement returns the union interpreted as *Es_event_authorization_judgement_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Authorization_judgement() *Es_event_authorization_judgement_t {
	return (*Es_event_authorization_judgement_t)(unsafe.Pointer(u))
}

// Authorization_petition returns the union interpreted as *Es_event_authorization_petition_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Authorization_petition() *Es_event_authorization_petition_t {
	return (*Es_event_authorization_petition_t)(unsafe.Pointer(u))
}

// Btm_launch_item_add returns the union interpreted as *Es_event_btm_launch_item_add_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Btm_launch_item_add() *Es_event_btm_launch_item_add_t {
	return (*Es_event_btm_launch_item_add_t)(unsafe.Pointer(u))
}

// Btm_launch_item_remove returns the union interpreted as *Es_event_btm_launch_item_remove_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Btm_launch_item_remove() *Es_event_btm_launch_item_remove_t {
	return (*Es_event_btm_launch_item_remove_t)(unsafe.Pointer(u))
}

// Gatekeeper_user_override returns the union interpreted as *Es_event_gatekeeper_user_override_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Gatekeeper_user_override() *Es_event_gatekeeper_user_override_t {
	return (*Es_event_gatekeeper_user_override_t)(unsafe.Pointer(u))
}

// Login_login returns the union interpreted as *Es_event_login_login_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Login_login() *Es_event_login_login_t {
	return (*Es_event_login_login_t)(unsafe.Pointer(u))
}

// Login_logout returns the union interpreted as *Es_event_login_logout_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Login_logout() *Es_event_login_logout_t {
	return (*Es_event_login_logout_t)(unsafe.Pointer(u))
}

// Lw_session_lock returns the union interpreted as *Es_event_lw_session_lock_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Lw_session_lock() *Es_event_lw_session_lock_t {
	return (*Es_event_lw_session_lock_t)(unsafe.Pointer(u))
}

// Lw_session_login returns the union interpreted as *Es_event_lw_session_login_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Lw_session_login() *Es_event_lw_session_login_t {
	return (*Es_event_lw_session_login_t)(unsafe.Pointer(u))
}

// Lw_session_logout returns the union interpreted as *Es_event_lw_session_logout_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Lw_session_logout() *Es_event_lw_session_logout_t {
	return (*Es_event_lw_session_logout_t)(unsafe.Pointer(u))
}

// Lw_session_unlock returns the union interpreted as *Es_event_lw_session_unlock_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Lw_session_unlock() *Es_event_lw_session_unlock_t {
	return (*Es_event_lw_session_unlock_t)(unsafe.Pointer(u))
}

// Od_attribute_set returns the union interpreted as *Es_event_od_attribute_set_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_attribute_set() *Es_event_od_attribute_set_t {
	return (*Es_event_od_attribute_set_t)(unsafe.Pointer(u))
}

// Od_attribute_value_add returns the union interpreted as *Es_event_od_attribute_value_add_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_attribute_value_add() *Es_event_od_attribute_value_add_t {
	return (*Es_event_od_attribute_value_add_t)(unsafe.Pointer(u))
}

// Od_attribute_value_remove returns the union interpreted as *Es_event_od_attribute_value_remove_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_attribute_value_remove() *Es_event_od_attribute_value_remove_t {
	return (*Es_event_od_attribute_value_remove_t)(unsafe.Pointer(u))
}

// Od_create_group returns the union interpreted as *Es_event_od_create_group_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_create_group() *Es_event_od_create_group_t {
	return (*Es_event_od_create_group_t)(unsafe.Pointer(u))
}

// Od_create_user returns the union interpreted as *Es_event_od_create_user_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_create_user() *Es_event_od_create_user_t {
	return (*Es_event_od_create_user_t)(unsafe.Pointer(u))
}

// Od_delete_group returns the union interpreted as *Es_event_od_delete_group_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_delete_group() *Es_event_od_delete_group_t {
	return (*Es_event_od_delete_group_t)(unsafe.Pointer(u))
}

// Od_delete_user returns the union interpreted as *Es_event_od_delete_user_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_delete_user() *Es_event_od_delete_user_t {
	return (*Es_event_od_delete_user_t)(unsafe.Pointer(u))
}

// Od_disable_user returns the union interpreted as *Es_event_od_disable_user_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_disable_user() *Es_event_od_disable_user_t {
	return (*Es_event_od_disable_user_t)(unsafe.Pointer(u))
}

// Od_enable_user returns the union interpreted as *Es_event_od_enable_user_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_enable_user() *Es_event_od_enable_user_t {
	return (*Es_event_od_enable_user_t)(unsafe.Pointer(u))
}

// Od_group_add returns the union interpreted as *Es_event_od_group_add_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_group_add() *Es_event_od_group_add_t {
	return (*Es_event_od_group_add_t)(unsafe.Pointer(u))
}

// Od_group_remove returns the union interpreted as *Es_event_od_group_remove_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_group_remove() *Es_event_od_group_remove_t {
	return (*Es_event_od_group_remove_t)(unsafe.Pointer(u))
}

// Od_group_set returns the union interpreted as *Es_event_od_group_set_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_group_set() *Es_event_od_group_set_t {
	return (*Es_event_od_group_set_t)(unsafe.Pointer(u))
}

// Od_modify_password returns the union interpreted as *Es_event_od_modify_password_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Od_modify_password() *Es_event_od_modify_password_t {
	return (*Es_event_od_modify_password_t)(unsafe.Pointer(u))
}

// Openssh_login returns the union interpreted as *Es_event_openssh_login_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Openssh_login() *Es_event_openssh_login_t {
	return (*Es_event_openssh_login_t)(unsafe.Pointer(u))
}

// Openssh_logout returns the union interpreted as *Es_event_openssh_logout_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Openssh_logout() *Es_event_openssh_logout_t {
	return (*Es_event_openssh_logout_t)(unsafe.Pointer(u))
}

// Profile_add returns the union interpreted as *Es_event_profile_add_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Profile_add() *Es_event_profile_add_t {
	return (*Es_event_profile_add_t)(unsafe.Pointer(u))
}

// Profile_remove returns the union interpreted as *Es_event_profile_remove_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Profile_remove() *Es_event_profile_remove_t {
	return (*Es_event_profile_remove_t)(unsafe.Pointer(u))
}

// Screensharing_attach returns the union interpreted as *Es_event_screensharing_attach_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Screensharing_attach() *Es_event_screensharing_attach_t {
	return (*Es_event_screensharing_attach_t)(unsafe.Pointer(u))
}

// Screensharing_detach returns the union interpreted as *Es_event_screensharing_detach_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Screensharing_detach() *Es_event_screensharing_detach_t {
	return (*Es_event_screensharing_detach_t)(unsafe.Pointer(u))
}

// Su returns the union interpreted as *Es_event_su_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Su() *Es_event_su_t {
	return (*Es_event_su_t)(unsafe.Pointer(u))
}

// Sudo returns the union interpreted as *Es_event_sudo_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Sudo() *Es_event_sudo_t {
	return (*Es_event_sudo_t)(unsafe.Pointer(u))
}

// Tcc_modify returns the union interpreted as *Es_event_tcc_modify_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Tcc_modify() *Es_event_tcc_modify_t {
	return (*Es_event_tcc_modify_t)(unsafe.Pointer(u))
}

// Xp_malware_detected returns the union interpreted as *Es_event_xp_malware_detected_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Xp_malware_detected() *Es_event_xp_malware_detected_t {
	return (*Es_event_xp_malware_detected_t)(unsafe.Pointer(u))
}

// Xp_malware_remediated returns the union interpreted as *Es_event_xp_malware_remediated_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Xp_malware_remediated() *Es_event_xp_malware_remediated_t {
	return (*Es_event_xp_malware_remediated_t)(unsafe.Pointer(u))
}

// Xpc_connect returns the union interpreted as *Es_event_xpc_connect_t.
// The returned pointer aliases the receiver's memory.
func (u *Es_events_t) Xpc_connect() *Es_event_xpc_connect_t {
	return (*Es_event_xpc_connect_t)(unsafe.Pointer(u))
}

// Es_fd_t - A structure that describes an open file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_fd_t
type Es_fd_t struct {
	Fd      int32  // The file descriptor number.
	Fdtype  uint32 // The file descriptor type, as a libproc type.
	Pipe    unsafe.Pointer
	Pipe_id uint64
}

// Es_file_t - A type that represents a file related to an Endpoint Security event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_file_t
type Es_file_t struct {
	Path           Es_string_token_t // The file’s path.
	Path_truncated bool              // A Boolean value that indicates whether Endpoint Security truncated the path string.
	Stat           kernel.Stat       // The file’s metadata, such as file size, user and group identifiers, and access and modification dates.

}

// Es_message_t - A message from the Endpoint Security subsystem that describes a security event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_message_t
type Es_message_t struct {
	Version        uint32           // The version of the Endpoint Security message.
	Time           syscall.Timespec // The time the event occurred, expressed as a Darwin time value.
	Mach_time      uint64           // The time the event occurred, as a Mach time value.
	Deadline       uint64           // The deadline by which your app must respond to the event.
	Process        *Es_process_t    // The process that performed the action defined in a message.
	Seq_num        uint64           // The sequence number of the message.
	Action_type    EsActionType     // The type of action: authentication or notification.
	Event_type     EsEventType      // The type of the message’s event.
	Event          Es_events_t      // The event that triggered this message.
	Thread         *Es_thread_t     // The thread that took the action defined in a message.
	Global_seq_num uint64           // The global sequence number of the message.
	Action         [36]byte         // The action monitored by Endpoint Security.
	Auth           Es_event_id_t
	Notify         Es_result_t
	Opaque         uint64 // An opaque storage field.

}

// Es_muted_path_t - A structure that describes a path’s muted events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_path_t
type Es_muted_path_t struct {
	Type        EsMutePathType    // The path type: prefix or literal.
	Event_count uintptr           // The number of elements in the muted events array.
	Path        Es_string_token_t // The muted path.
	Events      *EsEventType      // An array containing the muted event types.

}

// Es_muted_paths_t - A structure for a set of muted paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_t
type Es_muted_paths_t struct {
	Count uintptr          // The number of elements in the paths array.
	Paths *Es_muted_path_t // An array containing the muted paths.

}

// Es_muted_process_t - A structure that describes a process’s muted events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_process_t
type Es_muted_process_t struct {
	Audit_token [32]byte     // The audit token associated with a muted process.
	Event_count uintptr      // The number of elements in the muted events array.
	Events      *EsEventType // An array containing the muted event types.

}

// Es_muted_processes_t - A structure for a set of muted processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_t
type Es_muted_processes_t struct {
	Count     uintptr             // The number of elements in the processes array.
	Processes *Es_muted_process_t // An array containing the muted processes.

}

// Es_od_member_id_array_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_od_member_id_array_t
type Es_od_member_id_array_t struct {
	Member_type  EsOdMemberType
	Member_count uintptr
	Member_array [8]byte
	Names        *Es_string_token_t
	Uuids        unsafe.Pointer
}

// Es_od_member_id_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_od_member_id_t
type Es_od_member_id_t struct {
	Member_type  EsOdMemberType
	Member_value [16]byte
	Name         Es_string_token_t
	Uuid         [16]byte
}

// Es_process_t - A type that describes a process, as delivered by an Endpoint Security message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_process_t
type Es_process_t struct {
	Audit_token             [32]byte               // A token for use with Basic Security Module auditing functions.
	Ppid                    int32                  // The parent process identifier.
	Original_ppid           int32                  // The original parent process ID.
	Group_id                int32                  // The process group identifier.
	Session_id              int32                  // The identifier of the session that contains the process group.
	Codesigning_flags       uint32                 // The flags used to sign the process.
	Is_platform_binary      bool                   // A Boolean value that indicates whether the process is a platform binary.
	Is_es_client            bool                   // A Boolean value that indicates whether the process connects to the Endpoint Security subsystem.
	Cdhash                  Es_cdhash_t            // The code directory hash value.
	Signing_id              Es_string_token_t      // The identifier used to sign the process.
	Team_id                 Es_string_token_t      // The team identifier used to sign the process.
	Executable              *Es_file_t             // The file containing the executed process.
	Tty                     *Es_file_t             // The TTY associated with the process sending the message.
	Start_time              kernel.Timeval         // The time the process started.
	Responsible_audit_token [32]byte               // The audit token of the process responsible for this process.
	Parent_audit_token      [32]byte               // The audit token of the parent process.
	Cs_validation_category  EsCsValidationCategory // es_cs_validation_category

}

// Es_profile_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_profile_t
type Es_profile_t struct {
	Identifier     Es_string_token_t
	Uuid           Es_string_token_t
	Install_source EsProfileSource
	Organization   Es_string_token_t
	Display_name   Es_string_token_t
	Scope          Es_string_token_t
}

// Es_result_t - The result of the Endpoint Security subsystem authorization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_result_t
type Es_result_t struct {
	Result_type EsResultType // The type of the message’s result.
	Result      [32]byte     // The message’s result, as either an authorization result or flags.
	Auth        EsAuthResult
	Flags       uint32
	Reserved    uint8
}

// Es_signed_file_info_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_signed_file_info_t
type Es_signed_file_info_t struct {
	Cdhash     Es_cdhash_t
	Signing_id Es_string_token_t
	Team_id    Es_string_token_t
}

// Es_string_token_t - A pointer to a null-terminated string, and the length in bytes of that string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_string_token_t
type Es_string_token_t struct {
	Length uintptr // The size of the data buffer, in bytes.
	Data   *byte   // The string data.

}

// Es_sudo_reject_info_t
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_sudo_reject_info_t
type Es_sudo_reject_info_t struct {
	Plugin_name     Es_string_token_t
	Plugin_type     EsSudoPluginType
	Failure_message Es_string_token_t
}

// Es_thread_state_t - A description of a thread’s machine-specfiic state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_thread_state_t
type Es_thread_state_t struct {
	Flavor int        // An indication of the representation of the machine-specific thread state.
	State  Es_token_t // The machine-specific thread state.

}

// Es_thread_t - A structure that represents a thread in a process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_thread_t
type Es_thread_t struct {
	Thread_id uint64 // The unique identifier of the thread.

}

// Es_token_t - An arbitrary buffer of data with its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_token_t
type Es_token_t struct {
	Size uintptr // The size of the data buffer, in bytes.
	Data *uint8  // A data buffer.

}
