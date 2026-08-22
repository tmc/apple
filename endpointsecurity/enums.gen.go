// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_action_type_t
type EsActionType int32

const (
	// EsActionTypeAuth: The authentication action type.
	EsActionTypeAuth EsActionType = 0
	// EsActionTypeNotify: The notification action type.
	EsActionTypeNotify EsActionType = 1
)

func (e EsActionType) String() string {
	switch e {
	case EsActionTypeAuth:
		return "EsActionTypeAuth"
	case EsActionTypeNotify:
		return "EsActionTypeNotify"
	default:
		return fmt.Sprintf("EsActionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_address_type_t
type EsAddressType int32

const (
	EsAddressTypeIpv4        EsAddressType = 1
	EsAddressTypeIpv6        EsAddressType = 2
	EsAddressTypeNamedSocket EsAddressType = 3
	EsAddressTypeNone        EsAddressType = 0
)

func (e EsAddressType) String() string {
	switch e {
	case EsAddressTypeIpv4:
		return "EsAddressTypeIpv4"
	case EsAddressTypeIpv6:
		return "EsAddressTypeIpv6"
	case EsAddressTypeNamedSocket:
		return "EsAddressTypeNamedSocket"
	case EsAddressTypeNone:
		return "EsAddressTypeNone"
	default:
		return fmt.Sprintf("EsAddressType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_auth_result_t
type EsAuthResult int32

const (
	// EsAuthResultAllow: The caller authorizes the event and allows it to continue.
	EsAuthResultAllow EsAuthResult = 0
	// EsAuthResultDeny: The caller denies authorization to the event and prevents it from continuing.
	EsAuthResultDeny EsAuthResult = 1
)

func (e EsAuthResult) String() string {
	switch e {
	case EsAuthResultAllow:
		return "EsAuthResultAllow"
	case EsAuthResultDeny:
		return "EsAuthResultDeny"
	default:
		return fmt.Sprintf("EsAuthResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_authentication_type_t
type EsAuthenticationType int32

const (
	EsAuthenticationTypeAutoUnlock EsAuthenticationType = 3
	EsAuthenticationTypeLast       EsAuthenticationType = 4
	EsAuthenticationTypeOd         EsAuthenticationType = 0
	EsAuthenticationTypeToken      EsAuthenticationType = 2
	EsAuthenticationTypeTouchid    EsAuthenticationType = 1
)

func (e EsAuthenticationType) String() string {
	switch e {
	case EsAuthenticationTypeAutoUnlock:
		return "EsAuthenticationTypeAutoUnlock"
	case EsAuthenticationTypeLast:
		return "EsAuthenticationTypeLast"
	case EsAuthenticationTypeOd:
		return "EsAuthenticationTypeOd"
	case EsAuthenticationTypeToken:
		return "EsAuthenticationTypeToken"
	case EsAuthenticationTypeTouchid:
		return "EsAuthenticationTypeTouchid"
	default:
		return fmt.Sprintf("EsAuthenticationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_authorization_rule_class_t
type EsAuthorizationRuleClass int32

const (
	EsAuthorizationRuleClassAllow     EsAuthorizationRuleClass = 3
	EsAuthorizationRuleClassDeny      EsAuthorizationRuleClass = 4
	EsAuthorizationRuleClassInvalid   EsAuthorizationRuleClass = 6
	EsAuthorizationRuleClassMechanism EsAuthorizationRuleClass = 2
	EsAuthorizationRuleClassRule      EsAuthorizationRuleClass = 1
	EsAuthorizationRuleClassUnknown   EsAuthorizationRuleClass = 5
	EsAuthorizationRuleClassUser      EsAuthorizationRuleClass = 0
)

func (e EsAuthorizationRuleClass) String() string {
	switch e {
	case EsAuthorizationRuleClassAllow:
		return "EsAuthorizationRuleClassAllow"
	case EsAuthorizationRuleClassDeny:
		return "EsAuthorizationRuleClassDeny"
	case EsAuthorizationRuleClassInvalid:
		return "EsAuthorizationRuleClassInvalid"
	case EsAuthorizationRuleClassMechanism:
		return "EsAuthorizationRuleClassMechanism"
	case EsAuthorizationRuleClassRule:
		return "EsAuthorizationRuleClassRule"
	case EsAuthorizationRuleClassUnknown:
		return "EsAuthorizationRuleClassUnknown"
	case EsAuthorizationRuleClassUser:
		return "EsAuthorizationRuleClassUser"
	default:
		return fmt.Sprintf("EsAuthorizationRuleClass(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_auto_unlock_type_t
type EsAutoUnlockType int32

const (
	EsAutoUnlockAuthPrompt    EsAutoUnlockType = 2
	EsAutoUnlockMachineUnlock EsAutoUnlockType = 1
)

func (e EsAutoUnlockType) String() string {
	switch e {
	case EsAutoUnlockAuthPrompt:
		return "EsAutoUnlockAuthPrompt"
	case EsAutoUnlockMachineUnlock:
		return "EsAutoUnlockMachineUnlock"
	default:
		return fmt.Sprintf("EsAutoUnlockType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_btm_item_type_t
type EsBtmItemType int32

const (
	EsBtmItemTypeAgent     EsBtmItemType = 3
	EsBtmItemTypeApp       EsBtmItemType = 1
	EsBtmItemTypeDaemon    EsBtmItemType = 4
	EsBtmItemTypeLoginItem EsBtmItemType = 2
	EsBtmItemTypeUserItem  EsBtmItemType = 0
)

func (e EsBtmItemType) String() string {
	switch e {
	case EsBtmItemTypeAgent:
		return "EsBtmItemTypeAgent"
	case EsBtmItemTypeApp:
		return "EsBtmItemTypeApp"
	case EsBtmItemTypeDaemon:
		return "EsBtmItemTypeDaemon"
	case EsBtmItemTypeLoginItem:
		return "EsBtmItemTypeLoginItem"
	case EsBtmItemTypeUserItem:
		return "EsBtmItemTypeUserItem"
	default:
		return fmt.Sprintf("EsBtmItemType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_clear_cache_result_t
type EsClearCacheResult int32

const (
	// EsClearCacheResultErrInternal: Communication with the Endpoint Security system failed.
	EsClearCacheResultErrInternal EsClearCacheResult = 1
	// EsClearCacheResultErrThrottle: Clearing the cache failed because the rate of calls was too high.
	EsClearCacheResultErrThrottle EsClearCacheResult = 2
	// EsClearCacheResultSuccess: Clearing the cache succeeded.
	EsClearCacheResultSuccess EsClearCacheResult = 0
)

func (e EsClearCacheResult) String() string {
	switch e {
	case EsClearCacheResultErrInternal:
		return "EsClearCacheResultErrInternal"
	case EsClearCacheResultErrThrottle:
		return "EsClearCacheResultErrThrottle"
	case EsClearCacheResultSuccess:
		return "EsClearCacheResultSuccess"
	default:
		return fmt.Sprintf("EsClearCacheResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_cs_validation_category_t
type EsCsValidationCategory int32

const (
	EsCsValidationCategoryAppStore     EsCsValidationCategory = 4
	EsCsValidationCategoryDeveloperID  EsCsValidationCategory = 6
	EsCsValidationCategoryDevelopment  EsCsValidationCategory = 3
	EsCsValidationCategoryEnterprise   EsCsValidationCategory = 5
	EsCsValidationCategoryInvalid      EsCsValidationCategory = 0
	EsCsValidationCategoryLocalSigning EsCsValidationCategory = 7
	EsCsValidationCategoryNone         EsCsValidationCategory = 10
	EsCsValidationCategoryOopjit       EsCsValidationCategory = 9
	EsCsValidationCategoryPlatform     EsCsValidationCategory = 1
	EsCsValidationCategoryRosetta      EsCsValidationCategory = 8
	EsCsValidationCategoryTestflight   EsCsValidationCategory = 2
)

func (e EsCsValidationCategory) String() string {
	switch e {
	case EsCsValidationCategoryAppStore:
		return "EsCsValidationCategoryAppStore"
	case EsCsValidationCategoryDeveloperID:
		return "EsCsValidationCategoryDeveloperID"
	case EsCsValidationCategoryDevelopment:
		return "EsCsValidationCategoryDevelopment"
	case EsCsValidationCategoryEnterprise:
		return "EsCsValidationCategoryEnterprise"
	case EsCsValidationCategoryInvalid:
		return "EsCsValidationCategoryInvalid"
	case EsCsValidationCategoryLocalSigning:
		return "EsCsValidationCategoryLocalSigning"
	case EsCsValidationCategoryNone:
		return "EsCsValidationCategoryNone"
	case EsCsValidationCategoryOopjit:
		return "EsCsValidationCategoryOopjit"
	case EsCsValidationCategoryPlatform:
		return "EsCsValidationCategoryPlatform"
	case EsCsValidationCategoryRosetta:
		return "EsCsValidationCategoryRosetta"
	case EsCsValidationCategoryTestflight:
		return "EsCsValidationCategoryTestflight"
	default:
		return fmt.Sprintf("EsCsValidationCategory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_destination_type_t
type EsDestinationType int32

const (
	// EsDestinationTypeExistingFile: The destination is an existing file.
	EsDestinationTypeExistingFile EsDestinationType = 0
	// EsDestinationTypeNewPath: The destination is a path to a new location.
	EsDestinationTypeNewPath EsDestinationType = 1
)

func (e EsDestinationType) String() string {
	switch e {
	case EsDestinationTypeExistingFile:
		return "EsDestinationTypeExistingFile"
	case EsDestinationTypeNewPath:
		return "EsDestinationTypeNewPath"
	default:
		return fmt.Sprintf("EsDestinationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_event_type_t
type EsEventType int32

const (
	// EsEventTypeAuthChdir: An identifier for a process that requests permission from the operating system to change the working directory for the process.
	EsEventTypeAuthChdir EsEventType = 50
	// EsEventTypeAuthChroot: An identifier for a process that requests permission from the operating system to change the root directory for the process.
	EsEventTypeAuthChroot EsEventType = 56
	// EsEventTypeAuthClone: An identifier for a process that requests permission from the operating system to clone a file.
	EsEventTypeAuthClone EsEventType = 60
	// EsEventTypeAuthCopyfile: An identifier for a process that requests permission from the operating system to copy a file.
	EsEventTypeAuthCopyfile EsEventType = 109
	// EsEventTypeAuthCreate: An identifier for a process that requests permission from the operating system to create a file.
	EsEventTypeAuthCreate EsEventType = 44
	// EsEventTypeAuthDeleteextattr: An identifier for a process that requests permission from the operating system to delete an extended attribute from a file.
	EsEventTypeAuthDeleteextattr EsEventType = 69
	// EsEventTypeAuthExchangedata: An identifier for a process that requests permission from the operating system to exchange data between two files.
	EsEventTypeAuthExchangedata EsEventType = 80
	// EsEventTypeAuthExec: An identifier for a process that requests permission from the operating system to execute another image.
	EsEventTypeAuthExec EsEventType = 0
	// EsEventTypeAuthFcntl: An identifier for a process that requests permission from the operating system to manipulate a file descriptor.
	EsEventTypeAuthFcntl EsEventType = 90
	// EsEventTypeAuthFileProviderMaterialize: An identifier for a process that requests permission for a file provider to return a reference to a file.
	EsEventTypeAuthFileProviderMaterialize EsEventType = 34
	// EsEventTypeAuthFileProviderUpdate: An identifier for a process that requests permission from the operating system to update a file.
	EsEventTypeAuthFileProviderUpdate EsEventType = 36
	// EsEventTypeAuthFsgetpath: An identifier for a process that requests permission from the operating system to retrieve a file system path.
	EsEventTypeAuthFsgetpath EsEventType = 71
	// EsEventTypeAuthGetTask: An identifier for a process that requests permission from the operating system to retrieve a process’s task control port.
	EsEventTypeAuthGetTask EsEventType = 87
	// EsEventTypeAuthGetTaskRead: An identifier for a process that requests permission from the operating system to retrieve a process’s task read port.
	EsEventTypeAuthGetTaskRead EsEventType = 100
	// EsEventTypeAuthGetattrlist: An identifier for a process that requests permission from the operating system to retrieve attributes from a file.
	EsEventTypeAuthGetattrlist EsEventType = 52
	// EsEventTypeAuthGetextattr: An identifier for a process that requests permission from the operating system to retrieve an extended attribute from a file.
	EsEventTypeAuthGetextattr EsEventType = 63
	// EsEventTypeAuthIokitOpen: An identifier for a process that requests permission from the operating system to open an IOKit device.
	EsEventTypeAuthIokitOpen EsEventType = 91
	// EsEventTypeAuthKextload: An identifier for a process that requests permission from the operating system to load a kernel extension (KEXT).
	EsEventTypeAuthKextload EsEventType = 2
	// EsEventTypeAuthLink: An identifier for a process that requests permission from the operating system to create a hard link.
	EsEventTypeAuthLink EsEventType = 42
	// EsEventTypeAuthListextattr: An identifier for a process that requests permission from the operating system to retrieve multiple extended attributes from a file.
	EsEventTypeAuthListextattr EsEventType = 65
	// EsEventTypeAuthMmap: An identifier for a process that requests permission from the operating system to map a file into memory.
	EsEventTypeAuthMmap EsEventType = 3
	// EsEventTypeAuthMount: An identifier for a process that requests permission from the operating system to mount a file system.
	EsEventTypeAuthMount EsEventType = 5
	// EsEventTypeAuthMprotect: An identifier for a process that requests permission from the operating system to change the protection of memory-mapped pages.
	EsEventTypeAuthMprotect EsEventType = 4
	// EsEventTypeAuthOpen: An identifier for a process that requests permission from the operating system to open a file.
	EsEventTypeAuthOpen EsEventType = 1
	// EsEventTypeAuthProcCheck: An identifier for a process that requests permission from the operating system to get information about a process.
	EsEventTypeAuthProcCheck EsEventType = 85
	// EsEventTypeAuthProcSuspendResume: An identifier for a process that requests permission from the operating system to suspend, resume, or shut down sockets for another process.
	EsEventTypeAuthProcSuspendResume EsEventType = 92
	// EsEventTypeAuthReaddir: An identifier for a process that requests permission from the operating system to read a file system directory.
	EsEventTypeAuthReaddir EsEventType = 67
	// EsEventTypeAuthReadlink: An identifier for a process that requests permission from the operating system to read a symbolic link.
	EsEventTypeAuthReadlink EsEventType = 38
	// EsEventTypeAuthRemount: An identifier for a process that requests permission from the operating system to mount a file system.
	EsEventTypeAuthRemount EsEventType = 98
	// EsEventTypeAuthRename: An identifier for a process that requests permission from the operating system to rename a file.
	EsEventTypeAuthRename EsEventType = 6
	// EsEventTypeAuthSearchfs: An identifier for a process that requests permission from the operating system to search a volume or mounted file system.
	EsEventTypeAuthSearchfs EsEventType = 88
	// EsEventTypeAuthSetacl: An identifier for a process that requests permission from the operating system to set a file’s access control list.
	EsEventTypeAuthSetacl EsEventType = 81
	// EsEventTypeAuthSetattrlist: An identifier for a process that requests permission from the operating system to set attributes of a file.
	EsEventTypeAuthSetattrlist EsEventType = 45
	// EsEventTypeAuthSetextattr: An identifier for a process that requests permission from the operating system to set an extended attribute of a file.
	EsEventTypeAuthSetextattr EsEventType = 46
	// EsEventTypeAuthSetflags: An identifier for a process that requests permission from the operating system to set a file’s flags.
	EsEventTypeAuthSetflags EsEventType = 47
	// EsEventTypeAuthSetmode: An identifier for a process that requests permission from the operating system to set a file’s mode.
	EsEventTypeAuthSetmode EsEventType = 48
	// EsEventTypeAuthSetowner: An identifier for a process that requests permission from the operating system to set a file’s owner.
	EsEventTypeAuthSetowner EsEventType = 49
	// EsEventTypeAuthSettime: An identifier for a process that requests permission from the operating system to modify the system time.
	EsEventTypeAuthSettime EsEventType = 74
	// EsEventTypeAuthSignal: An identifier for a process that requests permission from the operating system to send a signal to a process.
	EsEventTypeAuthSignal EsEventType = 7
	// EsEventTypeAuthTruncate: An identifier for a process that requests permission from the operating system to truncate a file.
	EsEventTypeAuthTruncate EsEventType = 40
	// EsEventTypeAuthUIPCBind: An identifier for a process that requests permission from the operating system to bind a UNIX domain socket.
	EsEventTypeAuthUIPCBind EsEventType = 77
	// EsEventTypeAuthUIPCConnect: An identifier for a process that requests permission from the operating system to connect a UNIX domain socket.
	EsEventTypeAuthUIPCConnect EsEventType = 79
	// EsEventTypeAuthUnlink: An identifier for a process that requests permission from the operating system to delete a file.
	EsEventTypeAuthUnlink EsEventType = 8
	// EsEventTypeAuthUtimes: An identifier for a process that requests permission from the operating system to change a file’s access or modification time.
	EsEventTypeAuthUtimes EsEventType = 58
	// EsEventTypeLast: A value that indicates the last member of the enumeration.
	EsEventTypeLast EsEventType = 157
	// EsEventTypeNotifyAccess: An identifier for a process that notifies endpoint security that it is checking a file’s access permission.
	EsEventTypeNotifyAccess                 EsEventType = 55
	EsEventTypeNotifyAuthentication         EsEventType = 111
	EsEventTypeNotifyAuthorizationJudgement EsEventType = 130
	EsEventTypeNotifyAuthorizationPetition  EsEventType = 129
	EsEventTypeNotifyBtmLaunchItemAdd       EsEventType = 124
	EsEventTypeNotifyBtmLaunchItemRemove    EsEventType = 125
	// EsEventTypeNotifyChdir: An identifier for a process that notifies endpoint security that it is changing the working directory for the process.
	EsEventTypeNotifyChdir EsEventType = 51
	// EsEventTypeNotifyChroot: An identifier for a process that notifies endpoint security that it is changing the root directory for the process.
	EsEventTypeNotifyChroot EsEventType = 57
	// EsEventTypeNotifyClone: An identifier for a process that notifies endpoint security that it is cloning a file.
	EsEventTypeNotifyClone EsEventType = 61
	// EsEventTypeNotifyClose: An identifier for a process that notifies endpoint security that it is closing a file.
	EsEventTypeNotifyClose EsEventType = 12
	// EsEventTypeNotifyCopyfile: An identifier for a process that notifies endpoint security that it is copying a file.
	EsEventTypeNotifyCopyfile EsEventType = 110
	// EsEventTypeNotifyCreate: An identifier for a process that notifies endpoint security that it is creating a file.
	EsEventTypeNotifyCreate EsEventType = 13
	// EsEventTypeNotifyCsInvalidated: An identifier for a process that notifies endpoint security that its code signing status is now invalid.
	EsEventTypeNotifyCsInvalidated EsEventType = 94
	// EsEventTypeNotifyDeleteextattr: An identifier for a process that notifies endpoint security that it is deleting an extended attribute from a file.
	EsEventTypeNotifyDeleteextattr EsEventType = 70
	// EsEventTypeNotifyDup: An identifier for a process that notifies endpoint security that it is duplicating a file descriptor.
	EsEventTypeNotifyDup EsEventType = 73
	// EsEventTypeNotifyExchangedata: An identifier for a process that notifies endpoint security that it is exchanging data between two files.
	EsEventTypeNotifyExchangedata EsEventType = 14
	// EsEventTypeNotifyExec: An identifier for a process that notifies endpoint security that it is executing an image.
	EsEventTypeNotifyExec EsEventType = 9
	// EsEventTypeNotifyExit: An identifier for a process that notifies endpoint security that it is exiting.
	EsEventTypeNotifyExit EsEventType = 15
	// EsEventTypeNotifyFcntl: An identifier for a process that notifies endpoint security that it is manipulating a file descriptor.
	EsEventTypeNotifyFcntl EsEventType = 62
	// EsEventTypeNotifyFileProviderMaterialize: An identifier for a process that notifies endpoint security that a file provider returned a reference to a file.
	EsEventTypeNotifyFileProviderMaterialize EsEventType = 35
	// EsEventTypeNotifyFileProviderUpdate: An identifier for a process that notifies endpoint security that it is updating a file.
	EsEventTypeNotifyFileProviderUpdate EsEventType = 37
	// EsEventTypeNotifyFork: An identifier for a process that notifies endpoint security that it is forking another process.
	EsEventTypeNotifyFork EsEventType = 11
	// EsEventTypeNotifyFsgetpath: An identifier for a process that notifies endpoint security that it is retrieving a file system path.
	EsEventTypeNotifyFsgetpath              EsEventType = 72
	EsEventTypeNotifyGatekeeperUserOverride EsEventType = 146
	// EsEventTypeNotifyGetTask: An identifier for a process that notifies endpoint security that it is retrieving the task control port for another process.
	EsEventTypeNotifyGetTask EsEventType = 16
	// EsEventTypeNotifyGetTaskInspect: An identifier for a process that notifies endpoint security that it is retrieving the task inspect port for another process.
	EsEventTypeNotifyGetTaskInspect EsEventType = 102
	// EsEventTypeNotifyGetTaskName: An identifier for a process that notifies endpoint security that it is retrieving the task name port for another process.
	EsEventTypeNotifyGetTaskName EsEventType = 95
	// EsEventTypeNotifyGetTaskRead: An identifier for a process that notifies endpoint security that it is retrieving the task read port for another process.
	EsEventTypeNotifyGetTaskRead EsEventType = 101
	// EsEventTypeNotifyGetattrlist: An identifier for a process that notifies endpoint security that it is retrieving attributes from a file.
	EsEventTypeNotifyGetattrlist EsEventType = 53
	// EsEventTypeNotifyGetextattr: An identifier for a process that notifies endpoint security that it is retrieving an extended attribute from a file.
	EsEventTypeNotifyGetextattr EsEventType = 64
	// EsEventTypeNotifyIokitOpen: An identifier for a process that notifies endpoint security that it is opening an IOKit device.
	EsEventTypeNotifyIokitOpen EsEventType = 24
	// EsEventTypeNotifyKextload: An identifier for a process that notifies endpoint security that it is loading a kernel extension (KEXT).
	EsEventTypeNotifyKextload EsEventType = 17
	// EsEventTypeNotifyKextunload: An identifier for a process that notifies endpoint security that it is unloading a kernel extension (KEXT).
	EsEventTypeNotifyKextunload EsEventType = 18
	// EsEventTypeNotifyLink: An identifier for a process that notifies endpoint security that it is creating a hard link.
	EsEventTypeNotifyLink EsEventType = 19
	// EsEventTypeNotifyListextattr: An identifier for a process that notifies endpoint security that it is retrieving multiple extended attributes from a file.
	EsEventTypeNotifyListextattr EsEventType = 66
	EsEventTypeNotifyLoginLogin  EsEventType = 122
	EsEventTypeNotifyLoginLogout EsEventType = 123
	// EsEventTypeNotifyLookup: An identifier for a process that notifies endpoint security that it is looking up a file’s path.
	EsEventTypeNotifyLookup          EsEventType = 43
	EsEventTypeNotifyLwSessionLock   EsEventType = 116
	EsEventTypeNotifyLwSessionLogin  EsEventType = 114
	EsEventTypeNotifyLwSessionLogout EsEventType = 115
	EsEventTypeNotifyLwSessionUnlock EsEventType = 117
	// EsEventTypeNotifyMmap: An identifier for a process that notifies endpoint security that it is mapping a file into memory.
	EsEventTypeNotifyMmap EsEventType = 20
	// EsEventTypeNotifyMount: An identifier for a process that notifies endpoint security that it is mounting a file system.
	EsEventTypeNotifyMount EsEventType = 22
	// EsEventTypeNotifyMprotect: An identifier for a process that notifies endpoint security that it is changing the protection of memory-mapped pages.
	EsEventTypeNotifyMprotect               EsEventType = 21
	EsEventTypeNotifyOdAttributeSet         EsEventType = 140
	EsEventTypeNotifyOdAttributeValueAdd    EsEventType = 138
	EsEventTypeNotifyOdAttributeValueRemove EsEventType = 139
	EsEventTypeNotifyOdCreateGroup          EsEventType = 142
	EsEventTypeNotifyOdCreateUser           EsEventType = 141
	EsEventTypeNotifyOdDeleteGroup          EsEventType = 144
	EsEventTypeNotifyOdDeleteUser           EsEventType = 143
	EsEventTypeNotifyOdDisableUser          EsEventType = 136
	EsEventTypeNotifyOdEnableUser           EsEventType = 137
	EsEventTypeNotifyOdGroupAdd             EsEventType = 132
	EsEventTypeNotifyOdGroupRemove          EsEventType = 133
	EsEventTypeNotifyOdGroupSet             EsEventType = 134
	EsEventTypeNotifyOdModifyPassword       EsEventType = 135
	// EsEventTypeNotifyOpen: An identifier for a process that notifies endpoint security that it is opening a file.
	EsEventTypeNotifyOpen          EsEventType = 10
	EsEventTypeNotifyOpensshLogin  EsEventType = 120
	EsEventTypeNotifyOpensshLogout EsEventType = 121
	// EsEventTypeNotifyProcCheck: An identifier for a process that notifies endpoint security that it is checking information about another process.
	EsEventTypeNotifyProcCheck EsEventType = 86
	// EsEventTypeNotifyProcSuspendResume: An identifier for a process that notifies endpoint security that it is suspending, resuming, or shutting down sockets for another process.
	EsEventTypeNotifyProcSuspendResume EsEventType = 93
	EsEventTypeNotifyProfileAdd        EsEventType = 126
	EsEventTypeNotifyProfileRemove     EsEventType = 127
	// EsEventTypeNotifyPtyClose: An identifier for a process that notifies endpoint security that it is closing a pseudoterminal device.
	EsEventTypeNotifyPtyClose EsEventType = 84
	// EsEventTypeNotifyPtyGrant: An identifier for a process that notifies endpoint security that it is granting a pseudoterminal device to a user.
	EsEventTypeNotifyPtyGrant EsEventType = 83
	// EsEventTypeNotifyReaddir: An identifier for a process that notifies endpoint security that it is reading a file system directory.
	EsEventTypeNotifyReaddir EsEventType = 68
	// EsEventTypeNotifyReadlink: An identifier for a process that notifies endpoint security that it is reading a symbolic link.
	EsEventTypeNotifyReadlink EsEventType = 39
	// EsEventTypeNotifyRemoteThreadCreate: An identifier for a process that notifies endpoint security that it is spawning a thread in another process.
	EsEventTypeNotifyRemoteThreadCreate EsEventType = 97
	// EsEventTypeNotifyRemount: An identifier for a process that notifies endpoint security that it is remounting a file system.
	EsEventTypeNotifyRemount EsEventType = 99
	// EsEventTypeNotifyRename: An identifier for a process that notifies endpoint security that it is renaming a file.
	EsEventTypeNotifyRename              EsEventType = 25
	EsEventTypeNotifyScreensharingAttach EsEventType = 118
	EsEventTypeNotifyScreensharingDetach EsEventType = 119
	// EsEventTypeNotifySearchfs: An identifier for a process that notifies endpoint security that it is searching a volume or mounted file system.
	EsEventTypeNotifySearchfs EsEventType = 89
	// EsEventTypeNotifySetacl: An identifier for a process that notifies endpoint security that it is setting a file’s access control list.
	EsEventTypeNotifySetacl EsEventType = 82
	// EsEventTypeNotifySetattrlist: An identifier for a process that notifies endpoint security that it is setting attributes of a file.
	EsEventTypeNotifySetattrlist EsEventType = 26
	// EsEventTypeNotifySetegid: An identifier for a process that notifies endpoint security that it is setting its effective group ID.
	EsEventTypeNotifySetegid EsEventType = 106
	// EsEventTypeNotifySeteuid: An identifier for a process that notifies endpoint security that it is setting its effective user ID.
	EsEventTypeNotifySeteuid EsEventType = 105
	// EsEventTypeNotifySetextattr: An identifier for a process that notifies endpoint security that it is setting an extended attribute of a file.
	EsEventTypeNotifySetextattr EsEventType = 27
	// EsEventTypeNotifySetflags: An identifier for a process that notifies endpoint security that it is setting a file’s flags.
	EsEventTypeNotifySetflags EsEventType = 28
	// EsEventTypeNotifySetgid: An identifier for a process that notifies endpoint security that it is setting its group ID.
	EsEventTypeNotifySetgid EsEventType = 104
	// EsEventTypeNotifySetmode: An identifier for a process that notifies endpoint security that it is setting a file’s mode.
	EsEventTypeNotifySetmode EsEventType = 29
	// EsEventTypeNotifySetowner: An identifier for a process that notifies endpoint security that it is setting a file’s owner.
	EsEventTypeNotifySetowner EsEventType = 30
	// EsEventTypeNotifySetregid: An identifier for a process that notifies endpoint security that it is setting its real and effective group IDs.
	EsEventTypeNotifySetregid EsEventType = 108
	// EsEventTypeNotifySetreuid: An identifier for a process that notifies endpoint security that it is setting its real and effective user IDs.
	EsEventTypeNotifySetreuid EsEventType = 107
	// EsEventTypeNotifySettime: An identifier for a process that notifies endpoint security that it is modifying the system time.
	EsEventTypeNotifySettime EsEventType = 75
	// EsEventTypeNotifySetuid: An identifier for a process that notifies endpoint security that it is setting its user ID.
	EsEventTypeNotifySetuid EsEventType = 103
	// EsEventTypeNotifySignal: An identifier for a process that notifies endpoint security that it is sending a signal to another process.
	EsEventTypeNotifySignal EsEventType = 31
	// EsEventTypeNotifyStat: An identifier for a process that notifies endpoint security that it is retrieving a file’s status.
	EsEventTypeNotifyStat      EsEventType = 54
	EsEventTypeNotifySu        EsEventType = 128
	EsEventTypeNotifySudo      EsEventType = 131
	EsEventTypeNotifyTccModify EsEventType = 147
	// EsEventTypeNotifyTrace: An identifier for a process that notifies endpoint security that it is attaching to another process.
	EsEventTypeNotifyTrace EsEventType = 96
	// EsEventTypeNotifyTruncate: An identifier for a process that notifies endpoint security that it is truncating a file.
	EsEventTypeNotifyTruncate EsEventType = 41
	// EsEventTypeNotifyUIPCBind: An identifier for a process that notifies endpoint security that it is binding a UNIX domain socket.
	EsEventTypeNotifyUIPCBind EsEventType = 76
	// EsEventTypeNotifyUIPCConnect: An identifier for a process that notifies endpoint security that it is connecting to a UNIX domain socket.
	EsEventTypeNotifyUIPCConnect EsEventType = 78
	// EsEventTypeNotifyUnlink: An identifier for a process that notifies endpoint security that it is deleting a file.
	EsEventTypeNotifyUnlink EsEventType = 32
	// EsEventTypeNotifyUnmount: An identifier for a process that notifies endpoint security that it is unmounting a file system.
	EsEventTypeNotifyUnmount EsEventType = 23
	// EsEventTypeNotifyUtimes: An identifier for a process that notifies endpoint security that it is changing a file’s access or modification time.
	EsEventTypeNotifyUtimes EsEventType = 59
	// EsEventTypeNotifyWrite: An identifier for a process that notifies endpoint security that it is writing data to a file.
	EsEventTypeNotifyWrite               EsEventType = 33
	EsEventTypeNotifyXPCConnect          EsEventType = 145
	EsEventTypeNotifyXpMalwareDetected   EsEventType = 112
	EsEventTypeNotifyXpMalwareRemediated EsEventType = 113
	EsEventTypeReserved0                 EsEventType = 148
	EsEventTypeReserved1                 EsEventType = 149
	EsEventTypeReserved2                 EsEventType = 150
	EsEventTypeReserved3                 EsEventType = 151
	EsEventTypeReserved4                 EsEventType = 152
	EsEventTypeReserved5                 EsEventType = 153
	EsEventTypeReserved6                 EsEventType = 154
	EsEventTypeReserved7                 EsEventType = 155
	EsEventTypeReserved8                 EsEventType = 156
)

func (e EsEventType) String() string {
	switch e {
	case EsEventTypeAuthChdir:
		return "EsEventTypeAuthChdir"
	case EsEventTypeAuthChroot:
		return "EsEventTypeAuthChroot"
	case EsEventTypeAuthClone:
		return "EsEventTypeAuthClone"
	case EsEventTypeAuthCopyfile:
		return "EsEventTypeAuthCopyfile"
	case EsEventTypeAuthCreate:
		return "EsEventTypeAuthCreate"
	case EsEventTypeAuthDeleteextattr:
		return "EsEventTypeAuthDeleteextattr"
	case EsEventTypeAuthExchangedata:
		return "EsEventTypeAuthExchangedata"
	case EsEventTypeAuthExec:
		return "EsEventTypeAuthExec"
	case EsEventTypeAuthFcntl:
		return "EsEventTypeAuthFcntl"
	case EsEventTypeAuthFileProviderMaterialize:
		return "EsEventTypeAuthFileProviderMaterialize"
	case EsEventTypeAuthFileProviderUpdate:
		return "EsEventTypeAuthFileProviderUpdate"
	case EsEventTypeAuthFsgetpath:
		return "EsEventTypeAuthFsgetpath"
	case EsEventTypeAuthGetTask:
		return "EsEventTypeAuthGetTask"
	case EsEventTypeAuthGetTaskRead:
		return "EsEventTypeAuthGetTaskRead"
	case EsEventTypeAuthGetattrlist:
		return "EsEventTypeAuthGetattrlist"
	case EsEventTypeAuthGetextattr:
		return "EsEventTypeAuthGetextattr"
	case EsEventTypeAuthIokitOpen:
		return "EsEventTypeAuthIokitOpen"
	case EsEventTypeAuthKextload:
		return "EsEventTypeAuthKextload"
	case EsEventTypeAuthLink:
		return "EsEventTypeAuthLink"
	case EsEventTypeAuthListextattr:
		return "EsEventTypeAuthListextattr"
	case EsEventTypeAuthMmap:
		return "EsEventTypeAuthMmap"
	case EsEventTypeAuthMount:
		return "EsEventTypeAuthMount"
	case EsEventTypeAuthMprotect:
		return "EsEventTypeAuthMprotect"
	case EsEventTypeAuthOpen:
		return "EsEventTypeAuthOpen"
	case EsEventTypeAuthProcCheck:
		return "EsEventTypeAuthProcCheck"
	case EsEventTypeAuthProcSuspendResume:
		return "EsEventTypeAuthProcSuspendResume"
	case EsEventTypeAuthReaddir:
		return "EsEventTypeAuthReaddir"
	case EsEventTypeAuthReadlink:
		return "EsEventTypeAuthReadlink"
	case EsEventTypeAuthRemount:
		return "EsEventTypeAuthRemount"
	case EsEventTypeAuthRename:
		return "EsEventTypeAuthRename"
	case EsEventTypeAuthSearchfs:
		return "EsEventTypeAuthSearchfs"
	case EsEventTypeAuthSetacl:
		return "EsEventTypeAuthSetacl"
	case EsEventTypeAuthSetattrlist:
		return "EsEventTypeAuthSetattrlist"
	case EsEventTypeAuthSetextattr:
		return "EsEventTypeAuthSetextattr"
	case EsEventTypeAuthSetflags:
		return "EsEventTypeAuthSetflags"
	case EsEventTypeAuthSetmode:
		return "EsEventTypeAuthSetmode"
	case EsEventTypeAuthSetowner:
		return "EsEventTypeAuthSetowner"
	case EsEventTypeAuthSettime:
		return "EsEventTypeAuthSettime"
	case EsEventTypeAuthSignal:
		return "EsEventTypeAuthSignal"
	case EsEventTypeAuthTruncate:
		return "EsEventTypeAuthTruncate"
	case EsEventTypeAuthUIPCBind:
		return "EsEventTypeAuthUIPCBind"
	case EsEventTypeAuthUIPCConnect:
		return "EsEventTypeAuthUIPCConnect"
	case EsEventTypeAuthUnlink:
		return "EsEventTypeAuthUnlink"
	case EsEventTypeAuthUtimes:
		return "EsEventTypeAuthUtimes"
	case EsEventTypeLast:
		return "EsEventTypeLast"
	case EsEventTypeNotifyAccess:
		return "EsEventTypeNotifyAccess"
	case EsEventTypeNotifyAuthentication:
		return "EsEventTypeNotifyAuthentication"
	case EsEventTypeNotifyAuthorizationJudgement:
		return "EsEventTypeNotifyAuthorizationJudgement"
	case EsEventTypeNotifyAuthorizationPetition:
		return "EsEventTypeNotifyAuthorizationPetition"
	case EsEventTypeNotifyBtmLaunchItemAdd:
		return "EsEventTypeNotifyBtmLaunchItemAdd"
	case EsEventTypeNotifyBtmLaunchItemRemove:
		return "EsEventTypeNotifyBtmLaunchItemRemove"
	case EsEventTypeNotifyChdir:
		return "EsEventTypeNotifyChdir"
	case EsEventTypeNotifyChroot:
		return "EsEventTypeNotifyChroot"
	case EsEventTypeNotifyClone:
		return "EsEventTypeNotifyClone"
	case EsEventTypeNotifyClose:
		return "EsEventTypeNotifyClose"
	case EsEventTypeNotifyCopyfile:
		return "EsEventTypeNotifyCopyfile"
	case EsEventTypeNotifyCreate:
		return "EsEventTypeNotifyCreate"
	case EsEventTypeNotifyCsInvalidated:
		return "EsEventTypeNotifyCsInvalidated"
	case EsEventTypeNotifyDeleteextattr:
		return "EsEventTypeNotifyDeleteextattr"
	case EsEventTypeNotifyDup:
		return "EsEventTypeNotifyDup"
	case EsEventTypeNotifyExchangedata:
		return "EsEventTypeNotifyExchangedata"
	case EsEventTypeNotifyExec:
		return "EsEventTypeNotifyExec"
	case EsEventTypeNotifyExit:
		return "EsEventTypeNotifyExit"
	case EsEventTypeNotifyFcntl:
		return "EsEventTypeNotifyFcntl"
	case EsEventTypeNotifyFileProviderMaterialize:
		return "EsEventTypeNotifyFileProviderMaterialize"
	case EsEventTypeNotifyFileProviderUpdate:
		return "EsEventTypeNotifyFileProviderUpdate"
	case EsEventTypeNotifyFork:
		return "EsEventTypeNotifyFork"
	case EsEventTypeNotifyFsgetpath:
		return "EsEventTypeNotifyFsgetpath"
	case EsEventTypeNotifyGatekeeperUserOverride:
		return "EsEventTypeNotifyGatekeeperUserOverride"
	case EsEventTypeNotifyGetTask:
		return "EsEventTypeNotifyGetTask"
	case EsEventTypeNotifyGetTaskInspect:
		return "EsEventTypeNotifyGetTaskInspect"
	case EsEventTypeNotifyGetTaskName:
		return "EsEventTypeNotifyGetTaskName"
	case EsEventTypeNotifyGetTaskRead:
		return "EsEventTypeNotifyGetTaskRead"
	case EsEventTypeNotifyGetattrlist:
		return "EsEventTypeNotifyGetattrlist"
	case EsEventTypeNotifyGetextattr:
		return "EsEventTypeNotifyGetextattr"
	case EsEventTypeNotifyIokitOpen:
		return "EsEventTypeNotifyIokitOpen"
	case EsEventTypeNotifyKextload:
		return "EsEventTypeNotifyKextload"
	case EsEventTypeNotifyKextunload:
		return "EsEventTypeNotifyKextunload"
	case EsEventTypeNotifyLink:
		return "EsEventTypeNotifyLink"
	case EsEventTypeNotifyListextattr:
		return "EsEventTypeNotifyListextattr"
	case EsEventTypeNotifyLoginLogin:
		return "EsEventTypeNotifyLoginLogin"
	case EsEventTypeNotifyLoginLogout:
		return "EsEventTypeNotifyLoginLogout"
	case EsEventTypeNotifyLookup:
		return "EsEventTypeNotifyLookup"
	case EsEventTypeNotifyLwSessionLock:
		return "EsEventTypeNotifyLwSessionLock"
	case EsEventTypeNotifyLwSessionLogin:
		return "EsEventTypeNotifyLwSessionLogin"
	case EsEventTypeNotifyLwSessionLogout:
		return "EsEventTypeNotifyLwSessionLogout"
	case EsEventTypeNotifyLwSessionUnlock:
		return "EsEventTypeNotifyLwSessionUnlock"
	case EsEventTypeNotifyMmap:
		return "EsEventTypeNotifyMmap"
	case EsEventTypeNotifyMount:
		return "EsEventTypeNotifyMount"
	case EsEventTypeNotifyMprotect:
		return "EsEventTypeNotifyMprotect"
	case EsEventTypeNotifyOdAttributeSet:
		return "EsEventTypeNotifyOdAttributeSet"
	case EsEventTypeNotifyOdAttributeValueAdd:
		return "EsEventTypeNotifyOdAttributeValueAdd"
	case EsEventTypeNotifyOdAttributeValueRemove:
		return "EsEventTypeNotifyOdAttributeValueRemove"
	case EsEventTypeNotifyOdCreateGroup:
		return "EsEventTypeNotifyOdCreateGroup"
	case EsEventTypeNotifyOdCreateUser:
		return "EsEventTypeNotifyOdCreateUser"
	case EsEventTypeNotifyOdDeleteGroup:
		return "EsEventTypeNotifyOdDeleteGroup"
	case EsEventTypeNotifyOdDeleteUser:
		return "EsEventTypeNotifyOdDeleteUser"
	case EsEventTypeNotifyOdDisableUser:
		return "EsEventTypeNotifyOdDisableUser"
	case EsEventTypeNotifyOdEnableUser:
		return "EsEventTypeNotifyOdEnableUser"
	case EsEventTypeNotifyOdGroupAdd:
		return "EsEventTypeNotifyOdGroupAdd"
	case EsEventTypeNotifyOdGroupRemove:
		return "EsEventTypeNotifyOdGroupRemove"
	case EsEventTypeNotifyOdGroupSet:
		return "EsEventTypeNotifyOdGroupSet"
	case EsEventTypeNotifyOdModifyPassword:
		return "EsEventTypeNotifyOdModifyPassword"
	case EsEventTypeNotifyOpen:
		return "EsEventTypeNotifyOpen"
	case EsEventTypeNotifyOpensshLogin:
		return "EsEventTypeNotifyOpensshLogin"
	case EsEventTypeNotifyOpensshLogout:
		return "EsEventTypeNotifyOpensshLogout"
	case EsEventTypeNotifyProcCheck:
		return "EsEventTypeNotifyProcCheck"
	case EsEventTypeNotifyProcSuspendResume:
		return "EsEventTypeNotifyProcSuspendResume"
	case EsEventTypeNotifyProfileAdd:
		return "EsEventTypeNotifyProfileAdd"
	case EsEventTypeNotifyProfileRemove:
		return "EsEventTypeNotifyProfileRemove"
	case EsEventTypeNotifyPtyClose:
		return "EsEventTypeNotifyPtyClose"
	case EsEventTypeNotifyPtyGrant:
		return "EsEventTypeNotifyPtyGrant"
	case EsEventTypeNotifyReaddir:
		return "EsEventTypeNotifyReaddir"
	case EsEventTypeNotifyReadlink:
		return "EsEventTypeNotifyReadlink"
	case EsEventTypeNotifyRemoteThreadCreate:
		return "EsEventTypeNotifyRemoteThreadCreate"
	case EsEventTypeNotifyRemount:
		return "EsEventTypeNotifyRemount"
	case EsEventTypeNotifyRename:
		return "EsEventTypeNotifyRename"
	case EsEventTypeNotifyScreensharingAttach:
		return "EsEventTypeNotifyScreensharingAttach"
	case EsEventTypeNotifyScreensharingDetach:
		return "EsEventTypeNotifyScreensharingDetach"
	case EsEventTypeNotifySearchfs:
		return "EsEventTypeNotifySearchfs"
	case EsEventTypeNotifySetacl:
		return "EsEventTypeNotifySetacl"
	case EsEventTypeNotifySetattrlist:
		return "EsEventTypeNotifySetattrlist"
	case EsEventTypeNotifySetegid:
		return "EsEventTypeNotifySetegid"
	case EsEventTypeNotifySeteuid:
		return "EsEventTypeNotifySeteuid"
	case EsEventTypeNotifySetextattr:
		return "EsEventTypeNotifySetextattr"
	case EsEventTypeNotifySetflags:
		return "EsEventTypeNotifySetflags"
	case EsEventTypeNotifySetgid:
		return "EsEventTypeNotifySetgid"
	case EsEventTypeNotifySetmode:
		return "EsEventTypeNotifySetmode"
	case EsEventTypeNotifySetowner:
		return "EsEventTypeNotifySetowner"
	case EsEventTypeNotifySetregid:
		return "EsEventTypeNotifySetregid"
	case EsEventTypeNotifySetreuid:
		return "EsEventTypeNotifySetreuid"
	case EsEventTypeNotifySettime:
		return "EsEventTypeNotifySettime"
	case EsEventTypeNotifySetuid:
		return "EsEventTypeNotifySetuid"
	case EsEventTypeNotifySignal:
		return "EsEventTypeNotifySignal"
	case EsEventTypeNotifyStat:
		return "EsEventTypeNotifyStat"
	case EsEventTypeNotifySu:
		return "EsEventTypeNotifySu"
	case EsEventTypeNotifySudo:
		return "EsEventTypeNotifySudo"
	case EsEventTypeNotifyTccModify:
		return "EsEventTypeNotifyTccModify"
	case EsEventTypeNotifyTrace:
		return "EsEventTypeNotifyTrace"
	case EsEventTypeNotifyTruncate:
		return "EsEventTypeNotifyTruncate"
	case EsEventTypeNotifyUIPCBind:
		return "EsEventTypeNotifyUIPCBind"
	case EsEventTypeNotifyUIPCConnect:
		return "EsEventTypeNotifyUIPCConnect"
	case EsEventTypeNotifyUnlink:
		return "EsEventTypeNotifyUnlink"
	case EsEventTypeNotifyUnmount:
		return "EsEventTypeNotifyUnmount"
	case EsEventTypeNotifyUtimes:
		return "EsEventTypeNotifyUtimes"
	case EsEventTypeNotifyWrite:
		return "EsEventTypeNotifyWrite"
	case EsEventTypeNotifyXPCConnect:
		return "EsEventTypeNotifyXPCConnect"
	case EsEventTypeNotifyXpMalwareDetected:
		return "EsEventTypeNotifyXpMalwareDetected"
	case EsEventTypeNotifyXpMalwareRemediated:
		return "EsEventTypeNotifyXpMalwareRemediated"
	case EsEventTypeReserved0:
		return "EsEventTypeReserved0"
	case EsEventTypeReserved1:
		return "EsEventTypeReserved1"
	case EsEventTypeReserved2:
		return "EsEventTypeReserved2"
	case EsEventTypeReserved3:
		return "EsEventTypeReserved3"
	case EsEventTypeReserved4:
		return "EsEventTypeReserved4"
	case EsEventTypeReserved5:
		return "EsEventTypeReserved5"
	case EsEventTypeReserved6:
		return "EsEventTypeReserved6"
	case EsEventTypeReserved7:
		return "EsEventTypeReserved7"
	case EsEventTypeReserved8:
		return "EsEventTypeReserved8"
	default:
		return fmt.Sprintf("EsEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_gatekeeper_user_override_file_type_t
type EsGatekeeperUserOverrideFileType int32

const (
	EsGatekeeperUserOverrideFileTypeFile EsGatekeeperUserOverrideFileType = 1
	EsGatekeeperUserOverrideFileTypePath EsGatekeeperUserOverrideFileType = 0
)

func (e EsGatekeeperUserOverrideFileType) String() string {
	switch e {
	case EsGatekeeperUserOverrideFileTypeFile:
		return "EsGatekeeperUserOverrideFileTypeFile"
	case EsGatekeeperUserOverrideFileTypePath:
		return "EsGatekeeperUserOverrideFileTypePath"
	default:
		return fmt.Sprintf("EsGatekeeperUserOverrideFileType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_get_task_type_t
type EsGetTaskType int32

const (
	EsGetTaskTypeExposeTask    EsGetTaskType = 1
	EsGetTaskTypeIdentityToken EsGetTaskType = 2
	EsGetTaskTypeTaskForPid    EsGetTaskType = 0
)

func (e EsGetTaskType) String() string {
	switch e {
	case EsGetTaskTypeExposeTask:
		return "EsGetTaskTypeExposeTask"
	case EsGetTaskTypeIdentityToken:
		return "EsGetTaskTypeIdentityToken"
	case EsGetTaskTypeTaskForPid:
		return "EsGetTaskTypeTaskForPid"
	default:
		return fmt.Sprintf("EsGetTaskType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_mount_disposition_t
type EsMountDisposition int32

const (
	EsMountDispositionExternal EsMountDisposition = 0
	EsMountDispositionInternal EsMountDisposition = 1
	EsMountDispositionNetwork  EsMountDisposition = 2
	EsMountDispositionNullfs   EsMountDisposition = 4
	EsMountDispositionUnknown  EsMountDisposition = 5
	EsMountDispositionVirtual  EsMountDisposition = 3
)

func (e EsMountDisposition) String() string {
	switch e {
	case EsMountDispositionExternal:
		return "EsMountDispositionExternal"
	case EsMountDispositionInternal:
		return "EsMountDispositionInternal"
	case EsMountDispositionNetwork:
		return "EsMountDispositionNetwork"
	case EsMountDispositionNullfs:
		return "EsMountDispositionNullfs"
	case EsMountDispositionUnknown:
		return "EsMountDispositionUnknown"
	case EsMountDispositionVirtual:
		return "EsMountDispositionVirtual"
	default:
		return fmt.Sprintf("EsMountDisposition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_inversion_type_t
type EsMuteInversionType int32

const (
	EsMuteInversionTypeLast       EsMuteInversionType = 3
	EsMuteInversionTypePath       EsMuteInversionType = 1
	EsMuteInversionTypeProcess    EsMuteInversionType = 0
	EsMuteInversionTypeTargetPath EsMuteInversionType = 2
)

func (e EsMuteInversionType) String() string {
	switch e {
	case EsMuteInversionTypeLast:
		return "EsMuteInversionTypeLast"
	case EsMuteInversionTypePath:
		return "EsMuteInversionTypePath"
	case EsMuteInversionTypeProcess:
		return "EsMuteInversionTypeProcess"
	case EsMuteInversionTypeTargetPath:
		return "EsMuteInversionTypeTargetPath"
	default:
		return fmt.Sprintf("EsMuteInversionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_inverted_return_t
type EsMuteInvertedReturn int32

const (
	EsMuteInverted      EsMuteInvertedReturn = 0
	EsMuteInvertedError EsMuteInvertedReturn = 2
	EsMuteNotInverted   EsMuteInvertedReturn = 1
)

func (e EsMuteInvertedReturn) String() string {
	switch e {
	case EsMuteInverted:
		return "EsMuteInverted"
	case EsMuteInvertedError:
		return "EsMuteInvertedError"
	case EsMuteNotInverted:
		return "EsMuteNotInverted"
	default:
		return fmt.Sprintf("EsMuteInvertedReturn(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_type_t
type EsMutePathType int32

const (
	// EsMutePathTypeLiteral: A type for a path string used as a path literal.
	EsMutePathTypeLiteral EsMutePathType = 1
	// EsMutePathTypePrefix: A type for a path string used as a prefix.
	EsMutePathTypePrefix        EsMutePathType = 0
	EsMutePathTypeTargetLiteral EsMutePathType = 3
	EsMutePathTypeTargetPrefix  EsMutePathType = 2
)

func (e EsMutePathType) String() string {
	switch e {
	case EsMutePathTypeLiteral:
		return "EsMutePathTypeLiteral"
	case EsMutePathTypePrefix:
		return "EsMutePathTypePrefix"
	case EsMutePathTypeTargetLiteral:
		return "EsMutePathTypeTargetLiteral"
	case EsMutePathTypeTargetPrefix:
		return "EsMutePathTypeTargetPrefix"
	default:
		return fmt.Sprintf("EsMutePathType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_new_client_result_t
type EsNewClientResult int32

const (
	// EsNewClientResultErrInternal: Communication with the Endpoint Security subsystem failed.
	EsNewClientResultErrInternal EsNewClientResult = 2
	// EsNewClientResultErrInvalidArgument: The attempt to create a new client contained one or more invalid arguments.
	EsNewClientResultErrInvalidArgument EsNewClientResult = 1
	// EsNewClientResultErrNotEntitled: The caller isn’t properly entitled to connect to Endpoint Security.
	EsNewClientResultErrNotEntitled EsNewClientResult = 3
	// EsNewClientResultErrNotPermitted: The caller isn’t permitted to connect to Endpoint Security.
	EsNewClientResultErrNotPermitted EsNewClientResult = 4
	// EsNewClientResultErrNotPrivileged: The caller isn’t running as root.
	EsNewClientResultErrNotPrivileged EsNewClientResult = 5
	// EsNewClientResultErrTooManyClients: The caller has reached the maximum allowed number of simultaneously connected clients.
	EsNewClientResultErrTooManyClients EsNewClientResult = 6
	// EsNewClientResultSuccess: Endpoint Security successfully created the new client.
	EsNewClientResultSuccess EsNewClientResult = 0
)

func (e EsNewClientResult) String() string {
	switch e {
	case EsNewClientResultErrInternal:
		return "EsNewClientResultErrInternal"
	case EsNewClientResultErrInvalidArgument:
		return "EsNewClientResultErrInvalidArgument"
	case EsNewClientResultErrNotEntitled:
		return "EsNewClientResultErrNotEntitled"
	case EsNewClientResultErrNotPermitted:
		return "EsNewClientResultErrNotPermitted"
	case EsNewClientResultErrNotPrivileged:
		return "EsNewClientResultErrNotPrivileged"
	case EsNewClientResultErrTooManyClients:
		return "EsNewClientResultErrTooManyClients"
	case EsNewClientResultSuccess:
		return "EsNewClientResultSuccess"
	default:
		return fmt.Sprintf("EsNewClientResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_od_account_type_t
type EsOdAccountType int32

const (
	EsOdAccountTypeComputer EsOdAccountType = 1
	EsOdAccountTypeUser     EsOdAccountType = 0
)

func (e EsOdAccountType) String() string {
	switch e {
	case EsOdAccountTypeComputer:
		return "EsOdAccountTypeComputer"
	case EsOdAccountTypeUser:
		return "EsOdAccountTypeUser"
	default:
		return fmt.Sprintf("EsOdAccountType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_od_member_type_t
type EsOdMemberType int32

const (
	EsOdMemberTypeGroupUuid EsOdMemberType = 2
	EsOdMemberTypeUserName  EsOdMemberType = 0
	EsOdMemberTypeUserUuid  EsOdMemberType = 1
)

func (e EsOdMemberType) String() string {
	switch e {
	case EsOdMemberTypeGroupUuid:
		return "EsOdMemberTypeGroupUuid"
	case EsOdMemberTypeUserName:
		return "EsOdMemberTypeUserName"
	case EsOdMemberTypeUserUuid:
		return "EsOdMemberTypeUserUuid"
	default:
		return fmt.Sprintf("EsOdMemberType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_od_record_type_t
type EsOdRecordType int32

const (
	EsOdRecordTypeGroup EsOdRecordType = 1
	EsOdRecordTypeUser  EsOdRecordType = 0
)

func (e EsOdRecordType) String() string {
	switch e {
	case EsOdRecordTypeGroup:
		return "EsOdRecordTypeGroup"
	case EsOdRecordTypeUser:
		return "EsOdRecordTypeUser"
	default:
		return fmt.Sprintf("EsOdRecordType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_openssh_login_result_type_t
type EsOpensshLoginResultType int32

const (
	EsOpensshAuthFailGssapi      EsOpensshLoginResultType = 8
	EsOpensshAuthFailHostbased   EsOpensshLoginResultType = 7
	EsOpensshAuthFailKbdint      EsOpensshLoginResultType = 5
	EsOpensshAuthFailNone        EsOpensshLoginResultType = 3
	EsOpensshAuthFailPasswd      EsOpensshLoginResultType = 4
	EsOpensshAuthFailPubkey      EsOpensshLoginResultType = 6
	EsOpensshAuthSuccess         EsOpensshLoginResultType = 2
	EsOpensshInvalidUser         EsOpensshLoginResultType = 9
	EsOpensshLoginExceedMaxtries EsOpensshLoginResultType = 0
	EsOpensshLoginRootDenied     EsOpensshLoginResultType = 1
)

func (e EsOpensshLoginResultType) String() string {
	switch e {
	case EsOpensshAuthFailGssapi:
		return "EsOpensshAuthFailGssapi"
	case EsOpensshAuthFailHostbased:
		return "EsOpensshAuthFailHostbased"
	case EsOpensshAuthFailKbdint:
		return "EsOpensshAuthFailKbdint"
	case EsOpensshAuthFailNone:
		return "EsOpensshAuthFailNone"
	case EsOpensshAuthFailPasswd:
		return "EsOpensshAuthFailPasswd"
	case EsOpensshAuthFailPubkey:
		return "EsOpensshAuthFailPubkey"
	case EsOpensshAuthSuccess:
		return "EsOpensshAuthSuccess"
	case EsOpensshInvalidUser:
		return "EsOpensshInvalidUser"
	case EsOpensshLoginExceedMaxtries:
		return "EsOpensshLoginExceedMaxtries"
	case EsOpensshLoginRootDenied:
		return "EsOpensshLoginRootDenied"
	default:
		return fmt.Sprintf("EsOpensshLoginResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_proc_check_type_t
type EsProcCheckType int32

const (
	// EsProcCheckTypeDirtycontrol: A type of process check that uses the process’s dirty state.
	EsProcCheckTypeDirtycontrol EsProcCheckType = 0x8
	// EsProcCheckTypeKernmsgbuf: A type of process check that checks the message buffer.
	EsProcCheckTypeKernmsgbuf EsProcCheckType = 0x4
	// EsProcCheckTypeListpids: A type of process check that lists related process identifiers.
	EsProcCheckTypeListpids EsProcCheckType = 0x1
	// EsProcCheckTypePidfdinfo: A type of process check that gets file descriptor information.
	EsProcCheckTypePidfdinfo EsProcCheckType = 0x3
	// EsProcCheckTypePidfileportinfo: A type of process check that gets port information.
	EsProcCheckTypePidfileportinfo EsProcCheckType = 0x6
	// EsProcCheckTypePidinfo: A type of process check that gets basic process information.
	EsProcCheckTypePidinfo EsProcCheckType = 0x2
	// EsProcCheckTypePidrusage: A type of process check that gets a process’s resource usage information.
	EsProcCheckTypePidrusage EsProcCheckType = 0x9
	// EsProcCheckTypeSetcontrol: A type of process check that sets the process control state.
	EsProcCheckTypeSetcontrol EsProcCheckType = 0x5
	// EsProcCheckTypeTerminate: A type of process check that terninates a process.
	EsProcCheckTypeTerminate EsProcCheckType = 0x7
	// EsProcCheckTypeUdataInfo: A type of process check that involves a user data token.
	EsProcCheckTypeUdataInfo EsProcCheckType = 0xe
)

func (e EsProcCheckType) String() string {
	switch e {
	case EsProcCheckTypeDirtycontrol:
		return "EsProcCheckTypeDirtycontrol"
	case EsProcCheckTypeKernmsgbuf:
		return "EsProcCheckTypeKernmsgbuf"
	case EsProcCheckTypeListpids:
		return "EsProcCheckTypeListpids"
	case EsProcCheckTypePidfdinfo:
		return "EsProcCheckTypePidfdinfo"
	case EsProcCheckTypePidfileportinfo:
		return "EsProcCheckTypePidfileportinfo"
	case EsProcCheckTypePidinfo:
		return "EsProcCheckTypePidinfo"
	case EsProcCheckTypePidrusage:
		return "EsProcCheckTypePidrusage"
	case EsProcCheckTypeSetcontrol:
		return "EsProcCheckTypeSetcontrol"
	case EsProcCheckTypeTerminate:
		return "EsProcCheckTypeTerminate"
	case EsProcCheckTypeUdataInfo:
		return "EsProcCheckTypeUdataInfo"
	default:
		return fmt.Sprintf("EsProcCheckType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_proc_suspend_resume_type_t
type EsProcSuspendResumeType int32

const (
	// EsProcSuspendResumeTypeResume: An event type for process resumption events.
	EsProcSuspendResumeTypeResume EsProcSuspendResumeType = 1
	// EsProcSuspendResumeTypeShutdownSockets: An event type for process socket shutdown events.
	EsProcSuspendResumeTypeShutdownSockets EsProcSuspendResumeType = 3
	// EsProcSuspendResumeTypeSuspend: An event type for process suspension events.
	EsProcSuspendResumeTypeSuspend EsProcSuspendResumeType = 0
)

func (e EsProcSuspendResumeType) String() string {
	switch e {
	case EsProcSuspendResumeTypeResume:
		return "EsProcSuspendResumeTypeResume"
	case EsProcSuspendResumeTypeShutdownSockets:
		return "EsProcSuspendResumeTypeShutdownSockets"
	case EsProcSuspendResumeTypeSuspend:
		return "EsProcSuspendResumeTypeSuspend"
	default:
		return fmt.Sprintf("EsProcSuspendResumeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_profile_source_t
type EsProfileSource int32

const (
	EsProfileSourceInstall EsProfileSource = 1
	EsProfileSourceManaged EsProfileSource = 0
)

func (e EsProfileSource) String() string {
	switch e {
	case EsProfileSourceInstall:
		return "EsProfileSourceInstall"
	case EsProfileSourceManaged:
		return "EsProfileSourceManaged"
	default:
		return fmt.Sprintf("EsProfileSource(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_respond_result_t
type EsRespondResult int32

const (
	// EsRespondResultErrDuplicateResponse: The caller responded to a message that already received a response.
	EsRespondResultErrDuplicateResponse EsRespondResult = 4
	// EsRespondResultErrEventType: The caller performed an inappropriate response to the event.
	EsRespondResultErrEventType EsRespondResult = 5
	// EsRespondResultErrInternal: Communication with the Endpoint Security system failed.
	EsRespondResultErrInternal EsRespondResult = 2
	// EsRespondResultErrInvalidArgument: The caller provided one or more invalid arguments.
	EsRespondResultErrInvalidArgument EsRespondResult = 1
	// EsRespondResultNotFound: The system couldn’t find the message that the caller sent this response to.
	EsRespondResultNotFound EsRespondResult = 3
	// EsRespondResultSuccess: Endpoint Security successfully delivered the response.
	EsRespondResultSuccess EsRespondResult = 0
)

func (e EsRespondResult) String() string {
	switch e {
	case EsRespondResultErrDuplicateResponse:
		return "EsRespondResultErrDuplicateResponse"
	case EsRespondResultErrEventType:
		return "EsRespondResultErrEventType"
	case EsRespondResultErrInternal:
		return "EsRespondResultErrInternal"
	case EsRespondResultErrInvalidArgument:
		return "EsRespondResultErrInvalidArgument"
	case EsRespondResultNotFound:
		return "EsRespondResultNotFound"
	case EsRespondResultSuccess:
		return "EsRespondResultSuccess"
	default:
		return fmt.Sprintf("EsRespondResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_result_type_t
type EsResultType int32

const (
	// EsResultTypeAuth: The authorization result type.
	EsResultTypeAuth EsResultType = 0
	// EsResultTypeFlags: The flags result type.
	EsResultTypeFlags EsResultType = 1
)

func (e EsResultType) String() string {
	switch e {
	case EsResultTypeAuth:
		return "EsResultTypeAuth"
	case EsResultTypeFlags:
		return "EsResultTypeFlags"
	default:
		return fmt.Sprintf("EsResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_return_t
type EsReturn int32

const (
	// EsReturnError: The action failed with an error.
	EsReturnError EsReturn = 1
	// EsReturnSuccess: The action succeeded.
	EsReturnSuccess EsReturn = 0
)

func (e EsReturn) String() string {
	switch e {
	case EsReturnError:
		return "EsReturnError"
	case EsReturnSuccess:
		return "EsReturnSuccess"
	default:
		return fmt.Sprintf("EsReturn(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_set_or_clear_t
type EsSetOrClear int32

const (
	// EsClear: A case that indicates the event represents a clearing of the access control list.
	EsClear EsSetOrClear = 1
	// EsSet: A case that indicates the event represents a setting of access control list values.
	EsSet EsSetOrClear = 0
)

func (e EsSetOrClear) String() string {
	switch e {
	case EsClear:
		return "EsClear"
	case EsSet:
		return "EsSet"
	default:
		return fmt.Sprintf("EsSetOrClear(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sudo_plugin_type_t
type EsSudoPluginType int32

const (
	EsSudoPluginTypeApproval EsSudoPluginType = 5
	EsSudoPluginTypeAudit    EsSudoPluginType = 4
	EsSudoPluginTypeFrontEnd EsSudoPluginType = 1
	EsSudoPluginTypeIO       EsSudoPluginType = 3
	EsSudoPluginTypePolicy   EsSudoPluginType = 2
	EsSudoPluginTypeUnknown  EsSudoPluginType = 0
)

func (e EsSudoPluginType) String() string {
	switch e {
	case EsSudoPluginTypeApproval:
		return "EsSudoPluginTypeApproval"
	case EsSudoPluginTypeAudit:
		return "EsSudoPluginTypeAudit"
	case EsSudoPluginTypeFrontEnd:
		return "EsSudoPluginTypeFrontEnd"
	case EsSudoPluginTypeIO:
		return "EsSudoPluginTypeIO"
	case EsSudoPluginTypePolicy:
		return "EsSudoPluginTypePolicy"
	case EsSudoPluginTypeUnknown:
		return "EsSudoPluginTypeUnknown"
	default:
		return fmt.Sprintf("EsSudoPluginType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_tcc_authorization_reason_t
type EsTccAuthorizationReason int32

const (
	// EsTccAuthorizationReasonAppTypePolicy: A system process changed the authorization right
	EsTccAuthorizationReasonAppTypePolicy EsTccAuthorizationReason = 12
	// EsTccAuthorizationReasonEntitled: A system process changed the authorization right
	EsTccAuthorizationReasonEntitled EsTccAuthorizationReason = 11
	EsTccAuthorizationReasonError    EsTccAuthorizationReason = 1
	// EsTccAuthorizationReasonMdmPolicy: A system process changed the authorization right
	EsTccAuthorizationReasonMdmPolicy EsTccAuthorizationReason = 6
	// EsTccAuthorizationReasonMissingUsageString: A system process changed the authorization right
	EsTccAuthorizationReasonMissingUsageString EsTccAuthorizationReason = 8
	EsTccAuthorizationReasonNone               EsTccAuthorizationReason = 0
	// EsTccAuthorizationReasonPreflightUnknown: A system process changed the authorization right
	EsTccAuthorizationReasonPreflightUnknown EsTccAuthorizationReason = 10
	// EsTccAuthorizationReasonPromptCancel: A system process changed the authorization right
	EsTccAuthorizationReasonPromptCancel EsTccAuthorizationReason = 13
	// EsTccAuthorizationReasonPromptTimeout: A system process changed the authorization right
	EsTccAuthorizationReasonPromptTimeout EsTccAuthorizationReason = 9
	// EsTccAuthorizationReasonServiceOverridePolicy: A system process changed the authorization right
	EsTccAuthorizationReasonServiceOverridePolicy EsTccAuthorizationReason = 7
	// EsTccAuthorizationReasonServicePolicy: A system process changed the authorization right
	EsTccAuthorizationReasonServicePolicy EsTccAuthorizationReason = 5
	// EsTccAuthorizationReasonSystemSet: User changed the authorization right via Preferences
	EsTccAuthorizationReasonSystemSet   EsTccAuthorizationReason = 4
	EsTccAuthorizationReasonUserConsent EsTccAuthorizationReason = 2
	// EsTccAuthorizationReasonUserSet: User answered a prompt
	EsTccAuthorizationReasonUserSet EsTccAuthorizationReason = 3
)

func (e EsTccAuthorizationReason) String() string {
	switch e {
	case EsTccAuthorizationReasonAppTypePolicy:
		return "EsTccAuthorizationReasonAppTypePolicy"
	case EsTccAuthorizationReasonEntitled:
		return "EsTccAuthorizationReasonEntitled"
	case EsTccAuthorizationReasonError:
		return "EsTccAuthorizationReasonError"
	case EsTccAuthorizationReasonMdmPolicy:
		return "EsTccAuthorizationReasonMdmPolicy"
	case EsTccAuthorizationReasonMissingUsageString:
		return "EsTccAuthorizationReasonMissingUsageString"
	case EsTccAuthorizationReasonNone:
		return "EsTccAuthorizationReasonNone"
	case EsTccAuthorizationReasonPreflightUnknown:
		return "EsTccAuthorizationReasonPreflightUnknown"
	case EsTccAuthorizationReasonPromptCancel:
		return "EsTccAuthorizationReasonPromptCancel"
	case EsTccAuthorizationReasonPromptTimeout:
		return "EsTccAuthorizationReasonPromptTimeout"
	case EsTccAuthorizationReasonServiceOverridePolicy:
		return "EsTccAuthorizationReasonServiceOverridePolicy"
	case EsTccAuthorizationReasonServicePolicy:
		return "EsTccAuthorizationReasonServicePolicy"
	case EsTccAuthorizationReasonSystemSet:
		return "EsTccAuthorizationReasonSystemSet"
	case EsTccAuthorizationReasonUserConsent:
		return "EsTccAuthorizationReasonUserConsent"
	case EsTccAuthorizationReasonUserSet:
		return "EsTccAuthorizationReasonUserSet"
	default:
		return fmt.Sprintf("EsTccAuthorizationReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_tcc_authorization_right_t
type EsTccAuthorizationRight int32

const (
	EsTccAuthorizationRightAddModifyAdded EsTccAuthorizationRight = 4
	EsTccAuthorizationRightAllowed        EsTccAuthorizationRight = 2
	EsTccAuthorizationRightDenied         EsTccAuthorizationRight = 0
	EsTccAuthorizationRightLearnMore      EsTccAuthorizationRight = 6
	EsTccAuthorizationRightLimited        EsTccAuthorizationRight = 3
	EsTccAuthorizationRightSessionPid     EsTccAuthorizationRight = 5
	EsTccAuthorizationRightUnknown        EsTccAuthorizationRight = 1
)

func (e EsTccAuthorizationRight) String() string {
	switch e {
	case EsTccAuthorizationRightAddModifyAdded:
		return "EsTccAuthorizationRightAddModifyAdded"
	case EsTccAuthorizationRightAllowed:
		return "EsTccAuthorizationRightAllowed"
	case EsTccAuthorizationRightDenied:
		return "EsTccAuthorizationRightDenied"
	case EsTccAuthorizationRightLearnMore:
		return "EsTccAuthorizationRightLearnMore"
	case EsTccAuthorizationRightLimited:
		return "EsTccAuthorizationRightLimited"
	case EsTccAuthorizationRightSessionPid:
		return "EsTccAuthorizationRightSessionPid"
	case EsTccAuthorizationRightUnknown:
		return "EsTccAuthorizationRightUnknown"
	default:
		return fmt.Sprintf("EsTccAuthorizationRight(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_tcc_event_type_t
type EsTccEventType int32

const (
	EsTccEventTypeCreate  EsTccEventType = 1
	EsTccEventTypeDelete  EsTccEventType = 3
	EsTccEventTypeModify  EsTccEventType = 2
	EsTccEventTypeUnknown EsTccEventType = 0
)

func (e EsTccEventType) String() string {
	switch e {
	case EsTccEventTypeCreate:
		return "EsTccEventTypeCreate"
	case EsTccEventTypeDelete:
		return "EsTccEventTypeDelete"
	case EsTccEventTypeModify:
		return "EsTccEventTypeModify"
	case EsTccEventTypeUnknown:
		return "EsTccEventTypeUnknown"
	default:
		return fmt.Sprintf("EsTccEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_tcc_identity_type_t
type EsTccIdentityType int32

const (
	EsTccIdentityTypeBundleID             EsTccIdentityType = 0
	EsTccIdentityTypeExecutablePath       EsTccIdentityType = 1
	EsTccIdentityTypeFileProviderDomainID EsTccIdentityType = 3
	EsTccIdentityTypePolicyID             EsTccIdentityType = 2
)

func (e EsTccIdentityType) String() string {
	switch e {
	case EsTccIdentityTypeBundleID:
		return "EsTccIdentityTypeBundleID"
	case EsTccIdentityTypeExecutablePath:
		return "EsTccIdentityTypeExecutablePath"
	case EsTccIdentityTypeFileProviderDomainID:
		return "EsTccIdentityTypeFileProviderDomainID"
	case EsTccIdentityTypePolicyID:
		return "EsTccIdentityTypePolicyID"
	default:
		return fmt.Sprintf("EsTccIdentityType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_touchid_mode_t
type EsTouchidMode int32

const (
	EsTouchidModeIdentification EsTouchidMode = 1
	EsTouchidModeVerification   EsTouchidMode = 0
)

func (e EsTouchidMode) String() string {
	switch e {
	case EsTouchidModeIdentification:
		return "EsTouchidModeIdentification"
	case EsTouchidModeVerification:
		return "EsTouchidModeVerification"
	default:
		return fmt.Sprintf("EsTouchidMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/EndpointSecurity/es_xpc_domain_type_t
type EsXPCDomainType int32

const (
	EsXPCDomainTypeGui       EsXPCDomainType = 8
	EsXPCDomainTypeManager   EsXPCDomainType = 6
	EsXPCDomainTypePid       EsXPCDomainType = 5
	EsXPCDomainTypePort      EsXPCDomainType = 7
	EsXPCDomainTypeSession   EsXPCDomainType = 4
	EsXPCDomainTypeSystem    EsXPCDomainType = 1
	EsXPCDomainTypeUser      EsXPCDomainType = 2
	EsXPCDomainTypeUserLogin EsXPCDomainType = 3
)

func (e EsXPCDomainType) String() string {
	switch e {
	case EsXPCDomainTypeGui:
		return "EsXPCDomainTypeGui"
	case EsXPCDomainTypeManager:
		return "EsXPCDomainTypeManager"
	case EsXPCDomainTypePid:
		return "EsXPCDomainTypePid"
	case EsXPCDomainTypePort:
		return "EsXPCDomainTypePort"
	case EsXPCDomainTypeSession:
		return "EsXPCDomainTypeSession"
	case EsXPCDomainTypeSystem:
		return "EsXPCDomainTypeSystem"
	case EsXPCDomainTypeUser:
		return "EsXPCDomainTypeUser"
	case EsXPCDomainTypeUserLogin:
		return "EsXPCDomainTypeUserLogin"
	default:
		return fmt.Sprintf("EsXPCDomainType(%d)", e)
	}
}

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

// Es_clear_cache_result_t is a C-name alias for EsClearCacheResult.
type Es_clear_cache_result_t = EsClearCacheResult

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
