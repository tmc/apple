// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"fmt"
)

type Es int32

const (
	// ES_CLEAR: A case that indicates the event represents a clearing of the access control list.
	ES_CLEAR Es = 1
	// ES_SET: A case that indicates the event represents a setting of access control list values.
	ES_SET Es = 0
)

func (e Es) String() string {
	switch e {
	case ES_CLEAR:
		return "ES_CLEAR"
	case ES_SET:
		return "ES_SET"
	default:
		return fmt.Sprintf("Es(%d)", e)
	}
}

type EsActionType int32

const (
	// ES_ACTION_TYPE_AUTH: The authentication action type.
	ES_ACTION_TYPE_AUTH EsActionType = 0
	// ES_ACTION_TYPE_NOTIFY: The notification action type.
	ES_ACTION_TYPE_NOTIFY EsActionType = 1
)

func (e EsActionType) String() string {
	switch e {
	case ES_ACTION_TYPE_AUTH:
		return "ES_ACTION_TYPE_AUTH"
	case ES_ACTION_TYPE_NOTIFY:
		return "ES_ACTION_TYPE_NOTIFY"
	default:
		return fmt.Sprintf("EsActionType(%d)", e)
	}
}

type EsAddressType int32

const (
	ES_ADDRESS_TYPE_IPV4         EsAddressType = 1
	ES_ADDRESS_TYPE_IPV6         EsAddressType = 2
	ES_ADDRESS_TYPE_NAMED_SOCKET EsAddressType = 3
	ES_ADDRESS_TYPE_NONE         EsAddressType = 0
)

func (e EsAddressType) String() string {
	switch e {
	case ES_ADDRESS_TYPE_IPV4:
		return "ES_ADDRESS_TYPE_IPV4"
	case ES_ADDRESS_TYPE_IPV6:
		return "ES_ADDRESS_TYPE_IPV6"
	case ES_ADDRESS_TYPE_NAMED_SOCKET:
		return "ES_ADDRESS_TYPE_NAMED_SOCKET"
	case ES_ADDRESS_TYPE_NONE:
		return "ES_ADDRESS_TYPE_NONE"
	default:
		return fmt.Sprintf("EsAddressType(%d)", e)
	}
}

type EsAuthResult int32

const (
	// ES_AUTH_RESULT_ALLOW: The caller authorizes the event and allows it to continue.
	ES_AUTH_RESULT_ALLOW EsAuthResult = 0
	// ES_AUTH_RESULT_DENY: The caller denies authorization to the event and prevents it from continuing.
	ES_AUTH_RESULT_DENY EsAuthResult = 1
)

func (e EsAuthResult) String() string {
	switch e {
	case ES_AUTH_RESULT_ALLOW:
		return "ES_AUTH_RESULT_ALLOW"
	case ES_AUTH_RESULT_DENY:
		return "ES_AUTH_RESULT_DENY"
	default:
		return fmt.Sprintf("EsAuthResult(%d)", e)
	}
}

type EsAuthenticationType int32

const (
	ES_AUTHENTICATION_TYPE_AUTO_UNLOCK EsAuthenticationType = 3
	ES_AUTHENTICATION_TYPE_LAST        EsAuthenticationType = 4
	ES_AUTHENTICATION_TYPE_OD          EsAuthenticationType = 0
	ES_AUTHENTICATION_TYPE_TOKEN       EsAuthenticationType = 2
	ES_AUTHENTICATION_TYPE_TOUCHID     EsAuthenticationType = 1
)

func (e EsAuthenticationType) String() string {
	switch e {
	case ES_AUTHENTICATION_TYPE_AUTO_UNLOCK:
		return "ES_AUTHENTICATION_TYPE_AUTO_UNLOCK"
	case ES_AUTHENTICATION_TYPE_LAST:
		return "ES_AUTHENTICATION_TYPE_LAST"
	case ES_AUTHENTICATION_TYPE_OD:
		return "ES_AUTHENTICATION_TYPE_OD"
	case ES_AUTHENTICATION_TYPE_TOKEN:
		return "ES_AUTHENTICATION_TYPE_TOKEN"
	case ES_AUTHENTICATION_TYPE_TOUCHID:
		return "ES_AUTHENTICATION_TYPE_TOUCHID"
	default:
		return fmt.Sprintf("EsAuthenticationType(%d)", e)
	}
}

type EsAuthorizationRuleClass int32

const (
	ES_AUTHORIZATION_RULE_CLASS_ALLOW     EsAuthorizationRuleClass = 3
	ES_AUTHORIZATION_RULE_CLASS_DENY      EsAuthorizationRuleClass = 4
	ES_AUTHORIZATION_RULE_CLASS_INVALID   EsAuthorizationRuleClass = 6
	ES_AUTHORIZATION_RULE_CLASS_MECHANISM EsAuthorizationRuleClass = 2
	ES_AUTHORIZATION_RULE_CLASS_RULE      EsAuthorizationRuleClass = 1
	ES_AUTHORIZATION_RULE_CLASS_UNKNOWN   EsAuthorizationRuleClass = 5
	ES_AUTHORIZATION_RULE_CLASS_USER      EsAuthorizationRuleClass = 0
)

func (e EsAuthorizationRuleClass) String() string {
	switch e {
	case ES_AUTHORIZATION_RULE_CLASS_ALLOW:
		return "ES_AUTHORIZATION_RULE_CLASS_ALLOW"
	case ES_AUTHORIZATION_RULE_CLASS_DENY:
		return "ES_AUTHORIZATION_RULE_CLASS_DENY"
	case ES_AUTHORIZATION_RULE_CLASS_INVALID:
		return "ES_AUTHORIZATION_RULE_CLASS_INVALID"
	case ES_AUTHORIZATION_RULE_CLASS_MECHANISM:
		return "ES_AUTHORIZATION_RULE_CLASS_MECHANISM"
	case ES_AUTHORIZATION_RULE_CLASS_RULE:
		return "ES_AUTHORIZATION_RULE_CLASS_RULE"
	case ES_AUTHORIZATION_RULE_CLASS_UNKNOWN:
		return "ES_AUTHORIZATION_RULE_CLASS_UNKNOWN"
	case ES_AUTHORIZATION_RULE_CLASS_USER:
		return "ES_AUTHORIZATION_RULE_CLASS_USER"
	default:
		return fmt.Sprintf("EsAuthorizationRuleClass(%d)", e)
	}
}

type EsAutoUnlock int32

const (
	ES_AUTO_UNLOCK_AUTH_PROMPT    EsAutoUnlock = 2
	ES_AUTO_UNLOCK_MACHINE_UNLOCK EsAutoUnlock = 1
)

func (e EsAutoUnlock) String() string {
	switch e {
	case ES_AUTO_UNLOCK_AUTH_PROMPT:
		return "ES_AUTO_UNLOCK_AUTH_PROMPT"
	case ES_AUTO_UNLOCK_MACHINE_UNLOCK:
		return "ES_AUTO_UNLOCK_MACHINE_UNLOCK"
	default:
		return fmt.Sprintf("EsAutoUnlock(%d)", e)
	}
}

type EsBtmItemType int32

const (
	ES_BTM_ITEM_TYPE_AGENT      EsBtmItemType = 3
	ES_BTM_ITEM_TYPE_APP        EsBtmItemType = 1
	ES_BTM_ITEM_TYPE_DAEMON     EsBtmItemType = 4
	ES_BTM_ITEM_TYPE_LOGIN_ITEM EsBtmItemType = 2
	ES_BTM_ITEM_TYPE_USER_ITEM  EsBtmItemType = 0
)

func (e EsBtmItemType) String() string {
	switch e {
	case ES_BTM_ITEM_TYPE_AGENT:
		return "ES_BTM_ITEM_TYPE_AGENT"
	case ES_BTM_ITEM_TYPE_APP:
		return "ES_BTM_ITEM_TYPE_APP"
	case ES_BTM_ITEM_TYPE_DAEMON:
		return "ES_BTM_ITEM_TYPE_DAEMON"
	case ES_BTM_ITEM_TYPE_LOGIN_ITEM:
		return "ES_BTM_ITEM_TYPE_LOGIN_ITEM"
	case ES_BTM_ITEM_TYPE_USER_ITEM:
		return "ES_BTM_ITEM_TYPE_USER_ITEM"
	default:
		return fmt.Sprintf("EsBtmItemType(%d)", e)
	}
}

type EsClearCacheResult int32

const (
	// ES_CLEAR_CACHE_RESULT_ERR_INTERNAL: Communication with the Endpoint Security system failed.
	ES_CLEAR_CACHE_RESULT_ERR_INTERNAL EsClearCacheResult = 1
	// ES_CLEAR_CACHE_RESULT_ERR_THROTTLE: Clearing the cache failed because the rate of calls was too high.
	ES_CLEAR_CACHE_RESULT_ERR_THROTTLE EsClearCacheResult = 2
	// ES_CLEAR_CACHE_RESULT_SUCCESS: Clearing the cache succeeded.
	ES_CLEAR_CACHE_RESULT_SUCCESS EsClearCacheResult = 0
)

func (e EsClearCacheResult) String() string {
	switch e {
	case ES_CLEAR_CACHE_RESULT_ERR_INTERNAL:
		return "ES_CLEAR_CACHE_RESULT_ERR_INTERNAL"
	case ES_CLEAR_CACHE_RESULT_ERR_THROTTLE:
		return "ES_CLEAR_CACHE_RESULT_ERR_THROTTLE"
	case ES_CLEAR_CACHE_RESULT_SUCCESS:
		return "ES_CLEAR_CACHE_RESULT_SUCCESS"
	default:
		return fmt.Sprintf("EsClearCacheResult(%d)", e)
	}
}

type EsCsValidationCategory int32

const (
	ES_CS_VALIDATION_CATEGORY_APP_STORE     EsCsValidationCategory = 4
	ES_CS_VALIDATION_CATEGORY_DEVELOPER_ID  EsCsValidationCategory = 6
	ES_CS_VALIDATION_CATEGORY_DEVELOPMENT   EsCsValidationCategory = 3
	ES_CS_VALIDATION_CATEGORY_ENTERPRISE    EsCsValidationCategory = 5
	ES_CS_VALIDATION_CATEGORY_INVALID       EsCsValidationCategory = 0
	ES_CS_VALIDATION_CATEGORY_LOCAL_SIGNING EsCsValidationCategory = 7
	ES_CS_VALIDATION_CATEGORY_NONE          EsCsValidationCategory = 10
	ES_CS_VALIDATION_CATEGORY_OOPJIT        EsCsValidationCategory = 9
	ES_CS_VALIDATION_CATEGORY_PLATFORM      EsCsValidationCategory = 1
	ES_CS_VALIDATION_CATEGORY_ROSETTA       EsCsValidationCategory = 8
	ES_CS_VALIDATION_CATEGORY_TESTFLIGHT    EsCsValidationCategory = 2
)

func (e EsCsValidationCategory) String() string {
	switch e {
	case ES_CS_VALIDATION_CATEGORY_APP_STORE:
		return "ES_CS_VALIDATION_CATEGORY_APP_STORE"
	case ES_CS_VALIDATION_CATEGORY_DEVELOPER_ID:
		return "ES_CS_VALIDATION_CATEGORY_DEVELOPER_ID"
	case ES_CS_VALIDATION_CATEGORY_DEVELOPMENT:
		return "ES_CS_VALIDATION_CATEGORY_DEVELOPMENT"
	case ES_CS_VALIDATION_CATEGORY_ENTERPRISE:
		return "ES_CS_VALIDATION_CATEGORY_ENTERPRISE"
	case ES_CS_VALIDATION_CATEGORY_INVALID:
		return "ES_CS_VALIDATION_CATEGORY_INVALID"
	case ES_CS_VALIDATION_CATEGORY_LOCAL_SIGNING:
		return "ES_CS_VALIDATION_CATEGORY_LOCAL_SIGNING"
	case ES_CS_VALIDATION_CATEGORY_NONE:
		return "ES_CS_VALIDATION_CATEGORY_NONE"
	case ES_CS_VALIDATION_CATEGORY_OOPJIT:
		return "ES_CS_VALIDATION_CATEGORY_OOPJIT"
	case ES_CS_VALIDATION_CATEGORY_PLATFORM:
		return "ES_CS_VALIDATION_CATEGORY_PLATFORM"
	case ES_CS_VALIDATION_CATEGORY_ROSETTA:
		return "ES_CS_VALIDATION_CATEGORY_ROSETTA"
	case ES_CS_VALIDATION_CATEGORY_TESTFLIGHT:
		return "ES_CS_VALIDATION_CATEGORY_TESTFLIGHT"
	default:
		return fmt.Sprintf("EsCsValidationCategory(%d)", e)
	}
}

type EsDestinationType int32

const (
	// ES_DESTINATION_TYPE_EXISTING_FILE: The destination is an existing file.
	ES_DESTINATION_TYPE_EXISTING_FILE EsDestinationType = 0
	// ES_DESTINATION_TYPE_NEW_PATH: The destination is a path to a new location.
	ES_DESTINATION_TYPE_NEW_PATH EsDestinationType = 1
)

func (e EsDestinationType) String() string {
	switch e {
	case ES_DESTINATION_TYPE_EXISTING_FILE:
		return "ES_DESTINATION_TYPE_EXISTING_FILE"
	case ES_DESTINATION_TYPE_NEW_PATH:
		return "ES_DESTINATION_TYPE_NEW_PATH"
	default:
		return fmt.Sprintf("EsDestinationType(%d)", e)
	}
}

type EsEventType int32

const (
	// ES_EVENT_TYPE_AUTH_CHDIR: An identifier for a process that requests permission from the operating system to change the working directory for the process.
	ES_EVENT_TYPE_AUTH_CHDIR EsEventType = 50
	// ES_EVENT_TYPE_AUTH_CHROOT: An identifier for a process that requests permission from the operating system to change the root directory for the process.
	ES_EVENT_TYPE_AUTH_CHROOT EsEventType = 56
	// ES_EVENT_TYPE_AUTH_CLONE: An identifier for a process that requests permission from the operating system to clone a file.
	ES_EVENT_TYPE_AUTH_CLONE EsEventType = 60
	// ES_EVENT_TYPE_AUTH_COPYFILE: An identifier for a process that requests permission from the operating system to copy a file.
	ES_EVENT_TYPE_AUTH_COPYFILE EsEventType = 109
	// ES_EVENT_TYPE_AUTH_CREATE: An identifier for a process that requests permission from the operating system to create a file.
	ES_EVENT_TYPE_AUTH_CREATE EsEventType = 44
	// ES_EVENT_TYPE_AUTH_DELETEEXTATTR: An identifier for a process that requests permission from the operating system to delete an extended attribute from a file.
	ES_EVENT_TYPE_AUTH_DELETEEXTATTR EsEventType = 69
	// ES_EVENT_TYPE_AUTH_EXCHANGEDATA: An identifier for a process that requests permission from the operating system to exchange data between two files.
	ES_EVENT_TYPE_AUTH_EXCHANGEDATA EsEventType = 80
	// ES_EVENT_TYPE_AUTH_EXEC: An identifier for a process that requests permission from the operating system to execute another image.
	ES_EVENT_TYPE_AUTH_EXEC EsEventType = 0
	// ES_EVENT_TYPE_AUTH_FCNTL: An identifier for a process that requests permission from the operating system to manipulate a file descriptor.
	ES_EVENT_TYPE_AUTH_FCNTL EsEventType = 90
	// ES_EVENT_TYPE_AUTH_FILE_PROVIDER_MATERIALIZE: An identifier for a process that requests permission for a file provider to return a reference to a file.
	ES_EVENT_TYPE_AUTH_FILE_PROVIDER_MATERIALIZE EsEventType = 34
	// ES_EVENT_TYPE_AUTH_FILE_PROVIDER_UPDATE: An identifier for a process that requests permission from the operating system to update a file.
	ES_EVENT_TYPE_AUTH_FILE_PROVIDER_UPDATE EsEventType = 36
	// ES_EVENT_TYPE_AUTH_FSGETPATH: An identifier for a process that requests permission from the operating system to retrieve a file system path.
	ES_EVENT_TYPE_AUTH_FSGETPATH EsEventType = 71
	// ES_EVENT_TYPE_AUTH_GETATTRLIST: An identifier for a process that requests permission from the operating system to retrieve attributes from a file.
	ES_EVENT_TYPE_AUTH_GETATTRLIST EsEventType = 52
	// ES_EVENT_TYPE_AUTH_GETEXTATTR: An identifier for a process that requests permission from the operating system to retrieve an extended attribute from a file.
	ES_EVENT_TYPE_AUTH_GETEXTATTR EsEventType = 63
	// ES_EVENT_TYPE_AUTH_GET_TASK: An identifier for a process that requests permission from the operating system to retrieve a process’s task control port.
	ES_EVENT_TYPE_AUTH_GET_TASK EsEventType = 87
	// ES_EVENT_TYPE_AUTH_GET_TASK_READ: An identifier for a process that requests permission from the operating system to retrieve a process’s task read port.
	ES_EVENT_TYPE_AUTH_GET_TASK_READ EsEventType = 100
	// ES_EVENT_TYPE_AUTH_IOKIT_OPEN: An identifier for a process that requests permission from the operating system to open an IOKit device.
	ES_EVENT_TYPE_AUTH_IOKIT_OPEN EsEventType = 91
	// ES_EVENT_TYPE_AUTH_KEXTLOAD: An identifier for a process that requests permission from the operating system to load a kernel extension (KEXT).
	ES_EVENT_TYPE_AUTH_KEXTLOAD EsEventType = 2
	// ES_EVENT_TYPE_AUTH_LINK: An identifier for a process that requests permission from the operating system to create a hard link.
	ES_EVENT_TYPE_AUTH_LINK EsEventType = 42
	// ES_EVENT_TYPE_AUTH_LISTEXTATTR: An identifier for a process that requests permission from the operating system to retrieve multiple extended attributes from a file.
	ES_EVENT_TYPE_AUTH_LISTEXTATTR EsEventType = 65
	// ES_EVENT_TYPE_AUTH_MMAP: An identifier for a process that requests permission from the operating system to map a file into memory.
	ES_EVENT_TYPE_AUTH_MMAP EsEventType = 3
	// ES_EVENT_TYPE_AUTH_MOUNT: An identifier for a process that requests permission from the operating system to mount a file system.
	ES_EVENT_TYPE_AUTH_MOUNT EsEventType = 5
	// ES_EVENT_TYPE_AUTH_MPROTECT: An identifier for a process that requests permission from the operating system to change the protection of memory-mapped pages.
	ES_EVENT_TYPE_AUTH_MPROTECT EsEventType = 4
	// ES_EVENT_TYPE_AUTH_OPEN: An identifier for a process that requests permission from the operating system to open a file.
	ES_EVENT_TYPE_AUTH_OPEN EsEventType = 1
	// ES_EVENT_TYPE_AUTH_PROC_CHECK: An identifier for a process that requests permission from the operating system to get information about a process.
	ES_EVENT_TYPE_AUTH_PROC_CHECK EsEventType = 85
	// ES_EVENT_TYPE_AUTH_PROC_SUSPEND_RESUME: An identifier for a process that requests permission from the operating system to suspend, resume, or shut down sockets for another process.
	ES_EVENT_TYPE_AUTH_PROC_SUSPEND_RESUME EsEventType = 92
	// ES_EVENT_TYPE_AUTH_READDIR: An identifier for a process that requests permission from the operating system to read a file system directory.
	ES_EVENT_TYPE_AUTH_READDIR EsEventType = 67
	// ES_EVENT_TYPE_AUTH_READLINK: An identifier for a process that requests permission from the operating system to read a symbolic link.
	ES_EVENT_TYPE_AUTH_READLINK EsEventType = 38
	// ES_EVENT_TYPE_AUTH_REMOUNT: An identifier for a process that requests permission from the operating system to mount a file system.
	ES_EVENT_TYPE_AUTH_REMOUNT EsEventType = 98
	// ES_EVENT_TYPE_AUTH_RENAME: An identifier for a process that requests permission from the operating system to rename a file.
	ES_EVENT_TYPE_AUTH_RENAME EsEventType = 6
	// ES_EVENT_TYPE_AUTH_SEARCHFS: An identifier for a process that requests permission from the operating system to search a volume or mounted file system.
	ES_EVENT_TYPE_AUTH_SEARCHFS EsEventType = 88
	// ES_EVENT_TYPE_AUTH_SETACL: An identifier for a process that requests permission from the operating system to set a file’s access control list.
	ES_EVENT_TYPE_AUTH_SETACL EsEventType = 81
	// ES_EVENT_TYPE_AUTH_SETATTRLIST: An identifier for a process that requests permission from the operating system to set attributes of a file.
	ES_EVENT_TYPE_AUTH_SETATTRLIST EsEventType = 45
	// ES_EVENT_TYPE_AUTH_SETEXTATTR: An identifier for a process that requests permission from the operating system to set an extended attribute of a file.
	ES_EVENT_TYPE_AUTH_SETEXTATTR EsEventType = 46
	// ES_EVENT_TYPE_AUTH_SETFLAGS: An identifier for a process that requests permission from the operating system to set a file’s flags.
	ES_EVENT_TYPE_AUTH_SETFLAGS EsEventType = 47
	// ES_EVENT_TYPE_AUTH_SETMODE: An identifier for a process that requests permission from the operating system to set a file’s mode.
	ES_EVENT_TYPE_AUTH_SETMODE EsEventType = 48
	// ES_EVENT_TYPE_AUTH_SETOWNER: An identifier for a process that requests permission from the operating system to set a file’s owner.
	ES_EVENT_TYPE_AUTH_SETOWNER EsEventType = 49
	// ES_EVENT_TYPE_AUTH_SETTIME: An identifier for a process that requests permission from the operating system to modify the system time.
	ES_EVENT_TYPE_AUTH_SETTIME EsEventType = 74
	// ES_EVENT_TYPE_AUTH_SIGNAL: An identifier for a process that requests permission from the operating system to send a signal to a process.
	ES_EVENT_TYPE_AUTH_SIGNAL EsEventType = 7
	// ES_EVENT_TYPE_AUTH_TRUNCATE: An identifier for a process that requests permission from the operating system to truncate a file.
	ES_EVENT_TYPE_AUTH_TRUNCATE EsEventType = 40
	// ES_EVENT_TYPE_AUTH_UIPC_BIND: An identifier for a process that requests permission from the operating system to bind a UNIX domain socket.
	ES_EVENT_TYPE_AUTH_UIPC_BIND EsEventType = 77
	// ES_EVENT_TYPE_AUTH_UIPC_CONNECT: An identifier for a process that requests permission from the operating system to connect a UNIX domain socket.
	ES_EVENT_TYPE_AUTH_UIPC_CONNECT EsEventType = 79
	// ES_EVENT_TYPE_AUTH_UNLINK: An identifier for a process that requests permission from the operating system to delete a file.
	ES_EVENT_TYPE_AUTH_UNLINK EsEventType = 8
	// ES_EVENT_TYPE_AUTH_UTIMES: An identifier for a process that requests permission from the operating system to change a file’s access or modification time.
	ES_EVENT_TYPE_AUTH_UTIMES EsEventType = 58
	// ES_EVENT_TYPE_LAST: A value that indicates the last member of the enumeration.
	ES_EVENT_TYPE_LAST EsEventType = 155
	// ES_EVENT_TYPE_NOTIFY_ACCESS: An identifier for a process that notifies endpoint security that it is checking a file’s access permission.
	ES_EVENT_TYPE_NOTIFY_ACCESS                  EsEventType = 55
	ES_EVENT_TYPE_NOTIFY_AUTHENTICATION          EsEventType = 111
	ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_JUDGEMENT EsEventType = 130
	ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_PETITION  EsEventType = 129
	ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_ADD     EsEventType = 124
	ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_REMOVE  EsEventType = 125
	// ES_EVENT_TYPE_NOTIFY_CHDIR: An identifier for a process that notifies endpoint security that it is changing the working directory for the process.
	ES_EVENT_TYPE_NOTIFY_CHDIR EsEventType = 51
	// ES_EVENT_TYPE_NOTIFY_CHROOT: An identifier for a process that notifies endpoint security that it is changing the root directory for the process.
	ES_EVENT_TYPE_NOTIFY_CHROOT EsEventType = 57
	// ES_EVENT_TYPE_NOTIFY_CLONE: An identifier for a process that notifies endpoint security that it is cloning a file.
	ES_EVENT_TYPE_NOTIFY_CLONE EsEventType = 61
	// ES_EVENT_TYPE_NOTIFY_CLOSE: An identifier for a process that notifies endpoint security that it is closing a file.
	ES_EVENT_TYPE_NOTIFY_CLOSE EsEventType = 12
	// ES_EVENT_TYPE_NOTIFY_COPYFILE: An identifier for a process that notifies endpoint security that it is copying a file.
	ES_EVENT_TYPE_NOTIFY_COPYFILE EsEventType = 110
	// ES_EVENT_TYPE_NOTIFY_CREATE: An identifier for a process that notifies endpoint security that it is creating a file.
	ES_EVENT_TYPE_NOTIFY_CREATE EsEventType = 13
	// ES_EVENT_TYPE_NOTIFY_CS_INVALIDATED: An identifier for a process that notifies endpoint security that its code signing status is now invalid.
	ES_EVENT_TYPE_NOTIFY_CS_INVALIDATED EsEventType = 94
	// ES_EVENT_TYPE_NOTIFY_DELETEEXTATTR: An identifier for a process that notifies endpoint security that it is deleting an extended attribute from a file.
	ES_EVENT_TYPE_NOTIFY_DELETEEXTATTR EsEventType = 70
	// ES_EVENT_TYPE_NOTIFY_DUP: An identifier for a process that notifies endpoint security that it is duplicating a file descriptor.
	ES_EVENT_TYPE_NOTIFY_DUP EsEventType = 73
	// ES_EVENT_TYPE_NOTIFY_EXCHANGEDATA: An identifier for a process that notifies endpoint security that it is exchanging data between two files.
	ES_EVENT_TYPE_NOTIFY_EXCHANGEDATA EsEventType = 14
	// ES_EVENT_TYPE_NOTIFY_EXEC: An identifier for a process that notifies endpoint security that it is executing an image.
	ES_EVENT_TYPE_NOTIFY_EXEC EsEventType = 9
	// ES_EVENT_TYPE_NOTIFY_EXIT: An identifier for a process that notifies endpoint security that it is exiting.
	ES_EVENT_TYPE_NOTIFY_EXIT EsEventType = 15
	// ES_EVENT_TYPE_NOTIFY_FCNTL: An identifier for a process that notifies endpoint security that it is manipulating a file descriptor.
	ES_EVENT_TYPE_NOTIFY_FCNTL EsEventType = 62
	// ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_MATERIALIZE: An identifier for a process that notifies endpoint security that a file provider returned a reference to a file.
	ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_MATERIALIZE EsEventType = 35
	// ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_UPDATE: An identifier for a process that notifies endpoint security that it is updating a file.
	ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_UPDATE EsEventType = 37
	// ES_EVENT_TYPE_NOTIFY_FORK: An identifier for a process that notifies endpoint security that it is forking another process.
	ES_EVENT_TYPE_NOTIFY_FORK EsEventType = 11
	// ES_EVENT_TYPE_NOTIFY_FSGETPATH: An identifier for a process that notifies endpoint security that it is retrieving a file system path.
	ES_EVENT_TYPE_NOTIFY_FSGETPATH                EsEventType = 72
	ES_EVENT_TYPE_NOTIFY_GATEKEEPER_USER_OVERRIDE EsEventType = 146
	// ES_EVENT_TYPE_NOTIFY_GETATTRLIST: An identifier for a process that notifies endpoint security that it is retrieving attributes from a file.
	ES_EVENT_TYPE_NOTIFY_GETATTRLIST EsEventType = 53
	// ES_EVENT_TYPE_NOTIFY_GETEXTATTR: An identifier for a process that notifies endpoint security that it is retrieving an extended attribute from a file.
	ES_EVENT_TYPE_NOTIFY_GETEXTATTR EsEventType = 64
	// ES_EVENT_TYPE_NOTIFY_GET_TASK: An identifier for a process that notifies endpoint security that it is retrieving the task control port for another process.
	ES_EVENT_TYPE_NOTIFY_GET_TASK EsEventType = 16
	// ES_EVENT_TYPE_NOTIFY_GET_TASK_INSPECT: An identifier for a process that notifies endpoint security that it is retrieving the task inspect port for another process.
	ES_EVENT_TYPE_NOTIFY_GET_TASK_INSPECT EsEventType = 102
	// ES_EVENT_TYPE_NOTIFY_GET_TASK_NAME: An identifier for a process that notifies endpoint security that it is retrieving the task name port for another process.
	ES_EVENT_TYPE_NOTIFY_GET_TASK_NAME EsEventType = 95
	// ES_EVENT_TYPE_NOTIFY_GET_TASK_READ: An identifier for a process that notifies endpoint security that it is retrieving the task read port for another process.
	ES_EVENT_TYPE_NOTIFY_GET_TASK_READ EsEventType = 101
	// ES_EVENT_TYPE_NOTIFY_IOKIT_OPEN: An identifier for a process that notifies endpoint security that it is opening an IOKit device.
	ES_EVENT_TYPE_NOTIFY_IOKIT_OPEN EsEventType = 24
	// ES_EVENT_TYPE_NOTIFY_KEXTLOAD: An identifier for a process that notifies endpoint security that it is loading a kernel extension (KEXT).
	ES_EVENT_TYPE_NOTIFY_KEXTLOAD EsEventType = 17
	// ES_EVENT_TYPE_NOTIFY_KEXTUNLOAD: An identifier for a process that notifies endpoint security that it is unloading a kernel extension (KEXT).
	ES_EVENT_TYPE_NOTIFY_KEXTUNLOAD EsEventType = 18
	// ES_EVENT_TYPE_NOTIFY_LINK: An identifier for a process that notifies endpoint security that it is creating a hard link.
	ES_EVENT_TYPE_NOTIFY_LINK EsEventType = 19
	// ES_EVENT_TYPE_NOTIFY_LISTEXTATTR: An identifier for a process that notifies endpoint security that it is retrieving multiple extended attributes from a file.
	ES_EVENT_TYPE_NOTIFY_LISTEXTATTR  EsEventType = 66
	ES_EVENT_TYPE_NOTIFY_LOGIN_LOGIN  EsEventType = 122
	ES_EVENT_TYPE_NOTIFY_LOGIN_LOGOUT EsEventType = 123
	// ES_EVENT_TYPE_NOTIFY_LOOKUP: An identifier for a process that notifies endpoint security that it is looking up a file’s path.
	ES_EVENT_TYPE_NOTIFY_LOOKUP            EsEventType = 43
	ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOCK   EsEventType = 116
	ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGIN  EsEventType = 114
	ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGOUT EsEventType = 115
	ES_EVENT_TYPE_NOTIFY_LW_SESSION_UNLOCK EsEventType = 117
	// ES_EVENT_TYPE_NOTIFY_MMAP: An identifier for a process that notifies endpoint security that it is mapping a file into memory.
	ES_EVENT_TYPE_NOTIFY_MMAP EsEventType = 20
	// ES_EVENT_TYPE_NOTIFY_MOUNT: An identifier for a process that notifies endpoint security that it is mounting a file system.
	ES_EVENT_TYPE_NOTIFY_MOUNT EsEventType = 22
	// ES_EVENT_TYPE_NOTIFY_MPROTECT: An identifier for a process that notifies endpoint security that it is changing the protection of memory-mapped pages.
	ES_EVENT_TYPE_NOTIFY_MPROTECT                  EsEventType = 21
	ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_SET          EsEventType = 140
	ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_ADD    EsEventType = 138
	ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_REMOVE EsEventType = 139
	ES_EVENT_TYPE_NOTIFY_OD_CREATE_GROUP           EsEventType = 142
	ES_EVENT_TYPE_NOTIFY_OD_CREATE_USER            EsEventType = 141
	ES_EVENT_TYPE_NOTIFY_OD_DELETE_GROUP           EsEventType = 144
	ES_EVENT_TYPE_NOTIFY_OD_DELETE_USER            EsEventType = 143
	ES_EVENT_TYPE_NOTIFY_OD_DISABLE_USER           EsEventType = 136
	ES_EVENT_TYPE_NOTIFY_OD_ENABLE_USER            EsEventType = 137
	ES_EVENT_TYPE_NOTIFY_OD_GROUP_ADD              EsEventType = 132
	ES_EVENT_TYPE_NOTIFY_OD_GROUP_REMOVE           EsEventType = 133
	ES_EVENT_TYPE_NOTIFY_OD_GROUP_SET              EsEventType = 134
	ES_EVENT_TYPE_NOTIFY_OD_MODIFY_PASSWORD        EsEventType = 135
	// ES_EVENT_TYPE_NOTIFY_OPEN: An identifier for a process that notifies endpoint security that it is opening a file.
	ES_EVENT_TYPE_NOTIFY_OPEN           EsEventType = 10
	ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGIN  EsEventType = 120
	ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGOUT EsEventType = 121
	// ES_EVENT_TYPE_NOTIFY_PROC_CHECK: An identifier for a process that notifies endpoint security that it is checking information about another process.
	ES_EVENT_TYPE_NOTIFY_PROC_CHECK EsEventType = 86
	// ES_EVENT_TYPE_NOTIFY_PROC_SUSPEND_RESUME: An identifier for a process that notifies endpoint security that it is suspending, resuming, or shutting down sockets for another process.
	ES_EVENT_TYPE_NOTIFY_PROC_SUSPEND_RESUME EsEventType = 93
	ES_EVENT_TYPE_NOTIFY_PROFILE_ADD         EsEventType = 126
	ES_EVENT_TYPE_NOTIFY_PROFILE_REMOVE      EsEventType = 127
	// ES_EVENT_TYPE_NOTIFY_PTY_CLOSE: An identifier for a process that notifies endpoint security that it is closing a pseudoterminal device.
	ES_EVENT_TYPE_NOTIFY_PTY_CLOSE EsEventType = 84
	// ES_EVENT_TYPE_NOTIFY_PTY_GRANT: An identifier for a process that notifies endpoint security that it is granting a pseudoterminal device to a user.
	ES_EVENT_TYPE_NOTIFY_PTY_GRANT EsEventType = 83
	// ES_EVENT_TYPE_NOTIFY_READDIR: An identifier for a process that notifies endpoint security that it is reading a file system directory.
	ES_EVENT_TYPE_NOTIFY_READDIR EsEventType = 68
	// ES_EVENT_TYPE_NOTIFY_READLINK: An identifier for a process that notifies endpoint security that it is reading a symbolic link.
	ES_EVENT_TYPE_NOTIFY_READLINK EsEventType = 39
	// ES_EVENT_TYPE_NOTIFY_REMOTE_THREAD_CREATE: An identifier for a process that notifies endpoint security that it is spawning a thread in another process.
	ES_EVENT_TYPE_NOTIFY_REMOTE_THREAD_CREATE EsEventType = 97
	// ES_EVENT_TYPE_NOTIFY_REMOUNT: An identifier for a process that notifies endpoint security that it is remounting a file system.
	ES_EVENT_TYPE_NOTIFY_REMOUNT EsEventType = 99
	// ES_EVENT_TYPE_NOTIFY_RENAME: An identifier for a process that notifies endpoint security that it is renaming a file.
	ES_EVENT_TYPE_NOTIFY_RENAME               EsEventType = 25
	ES_EVENT_TYPE_NOTIFY_SCREENSHARING_ATTACH EsEventType = 118
	ES_EVENT_TYPE_NOTIFY_SCREENSHARING_DETACH EsEventType = 119
	// ES_EVENT_TYPE_NOTIFY_SEARCHFS: An identifier for a process that notifies endpoint security that it is searching a volume or mounted file system.
	ES_EVENT_TYPE_NOTIFY_SEARCHFS EsEventType = 89
	// ES_EVENT_TYPE_NOTIFY_SETACL: An identifier for a process that notifies endpoint security that it is setting a file’s access control list.
	ES_EVENT_TYPE_NOTIFY_SETACL EsEventType = 82
	// ES_EVENT_TYPE_NOTIFY_SETATTRLIST: An identifier for a process that notifies endpoint security that it is setting attributes of a file.
	ES_EVENT_TYPE_NOTIFY_SETATTRLIST EsEventType = 26
	// ES_EVENT_TYPE_NOTIFY_SETEGID: An identifier for a process that notifies endpoint security that it is setting its effective group ID.
	ES_EVENT_TYPE_NOTIFY_SETEGID EsEventType = 106
	// ES_EVENT_TYPE_NOTIFY_SETEUID: An identifier for a process that notifies endpoint security that it is setting its effective user ID.
	ES_EVENT_TYPE_NOTIFY_SETEUID EsEventType = 105
	// ES_EVENT_TYPE_NOTIFY_SETEXTATTR: An identifier for a process that notifies endpoint security that it is setting an extended attribute of a file.
	ES_EVENT_TYPE_NOTIFY_SETEXTATTR EsEventType = 27
	// ES_EVENT_TYPE_NOTIFY_SETFLAGS: An identifier for a process that notifies endpoint security that it is setting a file’s flags.
	ES_EVENT_TYPE_NOTIFY_SETFLAGS EsEventType = 28
	// ES_EVENT_TYPE_NOTIFY_SETGID: An identifier for a process that notifies endpoint security that it is setting its group ID.
	ES_EVENT_TYPE_NOTIFY_SETGID EsEventType = 104
	// ES_EVENT_TYPE_NOTIFY_SETMODE: An identifier for a process that notifies endpoint security that it is setting a file’s mode.
	ES_EVENT_TYPE_NOTIFY_SETMODE EsEventType = 29
	// ES_EVENT_TYPE_NOTIFY_SETOWNER: An identifier for a process that notifies endpoint security that it is setting a file’s owner.
	ES_EVENT_TYPE_NOTIFY_SETOWNER EsEventType = 30
	// ES_EVENT_TYPE_NOTIFY_SETREGID: An identifier for a process that notifies endpoint security that it is setting its real and effective group IDs.
	ES_EVENT_TYPE_NOTIFY_SETREGID EsEventType = 108
	// ES_EVENT_TYPE_NOTIFY_SETREUID: An identifier for a process that notifies endpoint security that it is setting its real and effective user IDs.
	ES_EVENT_TYPE_NOTIFY_SETREUID EsEventType = 107
	// ES_EVENT_TYPE_NOTIFY_SETTIME: An identifier for a process that notifies endpoint security that it is modifying the system time.
	ES_EVENT_TYPE_NOTIFY_SETTIME EsEventType = 75
	// ES_EVENT_TYPE_NOTIFY_SETUID: An identifier for a process that notifies endpoint security that it is setting its user ID.
	ES_EVENT_TYPE_NOTIFY_SETUID EsEventType = 103
	// ES_EVENT_TYPE_NOTIFY_SIGNAL: An identifier for a process that notifies endpoint security that it is sending a signal to another process.
	ES_EVENT_TYPE_NOTIFY_SIGNAL EsEventType = 31
	// ES_EVENT_TYPE_NOTIFY_STAT: An identifier for a process that notifies endpoint security that it is retrieving a file’s status.
	ES_EVENT_TYPE_NOTIFY_STAT       EsEventType = 54
	ES_EVENT_TYPE_NOTIFY_SU         EsEventType = 128
	ES_EVENT_TYPE_NOTIFY_SUDO       EsEventType = 131
	ES_EVENT_TYPE_NOTIFY_TCC_MODIFY EsEventType = 147
	// ES_EVENT_TYPE_NOTIFY_TRACE: An identifier for a process that notifies endpoint security that it is attaching to another process.
	ES_EVENT_TYPE_NOTIFY_TRACE EsEventType = 96
	// ES_EVENT_TYPE_NOTIFY_TRUNCATE: An identifier for a process that notifies endpoint security that it is truncating a file.
	ES_EVENT_TYPE_NOTIFY_TRUNCATE EsEventType = 41
	// ES_EVENT_TYPE_NOTIFY_UIPC_BIND: An identifier for a process that notifies endpoint security that it is binding a UNIX domain socket.
	ES_EVENT_TYPE_NOTIFY_UIPC_BIND EsEventType = 76
	// ES_EVENT_TYPE_NOTIFY_UIPC_CONNECT: An identifier for a process that notifies endpoint security that it is connecting to a UNIX domain socket.
	ES_EVENT_TYPE_NOTIFY_UIPC_CONNECT EsEventType = 78
	// ES_EVENT_TYPE_NOTIFY_UNLINK: An identifier for a process that notifies endpoint security that it is deleting a file.
	ES_EVENT_TYPE_NOTIFY_UNLINK EsEventType = 32
	// ES_EVENT_TYPE_NOTIFY_UNMOUNT: An identifier for a process that notifies endpoint security that it is unmounting a file system.
	ES_EVENT_TYPE_NOTIFY_UNMOUNT EsEventType = 23
	// ES_EVENT_TYPE_NOTIFY_UTIMES: An identifier for a process that notifies endpoint security that it is changing a file’s access or modification time.
	ES_EVENT_TYPE_NOTIFY_UTIMES EsEventType = 59
	// ES_EVENT_TYPE_NOTIFY_WRITE: An identifier for a process that notifies endpoint security that it is writing data to a file.
	ES_EVENT_TYPE_NOTIFY_WRITE                 EsEventType = 33
	ES_EVENT_TYPE_NOTIFY_XPC_CONNECT           EsEventType = 145
	ES_EVENT_TYPE_NOTIFY_XP_MALWARE_DETECTED   EsEventType = 112
	ES_EVENT_TYPE_NOTIFY_XP_MALWARE_REMEDIATED EsEventType = 113
	ES_EVENT_TYPE_RESERVED_0                   EsEventType = 148
	ES_EVENT_TYPE_RESERVED_1                   EsEventType = 149
	ES_EVENT_TYPE_RESERVED_2                   EsEventType = 150
	ES_EVENT_TYPE_RESERVED_3                   EsEventType = 151
	ES_EVENT_TYPE_RESERVED_4                   EsEventType = 152
	ES_EVENT_TYPE_RESERVED_5                   EsEventType = 153
	ES_EVENT_TYPE_RESERVED_6                   EsEventType = 154
)

func (e EsEventType) String() string {
	switch e {
	case ES_EVENT_TYPE_AUTH_CHDIR:
		return "ES_EVENT_TYPE_AUTH_CHDIR"
	case ES_EVENT_TYPE_AUTH_CHROOT:
		return "ES_EVENT_TYPE_AUTH_CHROOT"
	case ES_EVENT_TYPE_AUTH_CLONE:
		return "ES_EVENT_TYPE_AUTH_CLONE"
	case ES_EVENT_TYPE_AUTH_COPYFILE:
		return "ES_EVENT_TYPE_AUTH_COPYFILE"
	case ES_EVENT_TYPE_AUTH_CREATE:
		return "ES_EVENT_TYPE_AUTH_CREATE"
	case ES_EVENT_TYPE_AUTH_DELETEEXTATTR:
		return "ES_EVENT_TYPE_AUTH_DELETEEXTATTR"
	case ES_EVENT_TYPE_AUTH_EXCHANGEDATA:
		return "ES_EVENT_TYPE_AUTH_EXCHANGEDATA"
	case ES_EVENT_TYPE_AUTH_EXEC:
		return "ES_EVENT_TYPE_AUTH_EXEC"
	case ES_EVENT_TYPE_AUTH_FCNTL:
		return "ES_EVENT_TYPE_AUTH_FCNTL"
	case ES_EVENT_TYPE_AUTH_FILE_PROVIDER_MATERIALIZE:
		return "ES_EVENT_TYPE_AUTH_FILE_PROVIDER_MATERIALIZE"
	case ES_EVENT_TYPE_AUTH_FILE_PROVIDER_UPDATE:
		return "ES_EVENT_TYPE_AUTH_FILE_PROVIDER_UPDATE"
	case ES_EVENT_TYPE_AUTH_FSGETPATH:
		return "ES_EVENT_TYPE_AUTH_FSGETPATH"
	case ES_EVENT_TYPE_AUTH_GETATTRLIST:
		return "ES_EVENT_TYPE_AUTH_GETATTRLIST"
	case ES_EVENT_TYPE_AUTH_GETEXTATTR:
		return "ES_EVENT_TYPE_AUTH_GETEXTATTR"
	case ES_EVENT_TYPE_AUTH_GET_TASK:
		return "ES_EVENT_TYPE_AUTH_GET_TASK"
	case ES_EVENT_TYPE_AUTH_GET_TASK_READ:
		return "ES_EVENT_TYPE_AUTH_GET_TASK_READ"
	case ES_EVENT_TYPE_AUTH_IOKIT_OPEN:
		return "ES_EVENT_TYPE_AUTH_IOKIT_OPEN"
	case ES_EVENT_TYPE_AUTH_KEXTLOAD:
		return "ES_EVENT_TYPE_AUTH_KEXTLOAD"
	case ES_EVENT_TYPE_AUTH_LINK:
		return "ES_EVENT_TYPE_AUTH_LINK"
	case ES_EVENT_TYPE_AUTH_LISTEXTATTR:
		return "ES_EVENT_TYPE_AUTH_LISTEXTATTR"
	case ES_EVENT_TYPE_AUTH_MMAP:
		return "ES_EVENT_TYPE_AUTH_MMAP"
	case ES_EVENT_TYPE_AUTH_MOUNT:
		return "ES_EVENT_TYPE_AUTH_MOUNT"
	case ES_EVENT_TYPE_AUTH_MPROTECT:
		return "ES_EVENT_TYPE_AUTH_MPROTECT"
	case ES_EVENT_TYPE_AUTH_OPEN:
		return "ES_EVENT_TYPE_AUTH_OPEN"
	case ES_EVENT_TYPE_AUTH_PROC_CHECK:
		return "ES_EVENT_TYPE_AUTH_PROC_CHECK"
	case ES_EVENT_TYPE_AUTH_PROC_SUSPEND_RESUME:
		return "ES_EVENT_TYPE_AUTH_PROC_SUSPEND_RESUME"
	case ES_EVENT_TYPE_AUTH_READDIR:
		return "ES_EVENT_TYPE_AUTH_READDIR"
	case ES_EVENT_TYPE_AUTH_READLINK:
		return "ES_EVENT_TYPE_AUTH_READLINK"
	case ES_EVENT_TYPE_AUTH_REMOUNT:
		return "ES_EVENT_TYPE_AUTH_REMOUNT"
	case ES_EVENT_TYPE_AUTH_RENAME:
		return "ES_EVENT_TYPE_AUTH_RENAME"
	case ES_EVENT_TYPE_AUTH_SEARCHFS:
		return "ES_EVENT_TYPE_AUTH_SEARCHFS"
	case ES_EVENT_TYPE_AUTH_SETACL:
		return "ES_EVENT_TYPE_AUTH_SETACL"
	case ES_EVENT_TYPE_AUTH_SETATTRLIST:
		return "ES_EVENT_TYPE_AUTH_SETATTRLIST"
	case ES_EVENT_TYPE_AUTH_SETEXTATTR:
		return "ES_EVENT_TYPE_AUTH_SETEXTATTR"
	case ES_EVENT_TYPE_AUTH_SETFLAGS:
		return "ES_EVENT_TYPE_AUTH_SETFLAGS"
	case ES_EVENT_TYPE_AUTH_SETMODE:
		return "ES_EVENT_TYPE_AUTH_SETMODE"
	case ES_EVENT_TYPE_AUTH_SETOWNER:
		return "ES_EVENT_TYPE_AUTH_SETOWNER"
	case ES_EVENT_TYPE_AUTH_SETTIME:
		return "ES_EVENT_TYPE_AUTH_SETTIME"
	case ES_EVENT_TYPE_AUTH_SIGNAL:
		return "ES_EVENT_TYPE_AUTH_SIGNAL"
	case ES_EVENT_TYPE_AUTH_TRUNCATE:
		return "ES_EVENT_TYPE_AUTH_TRUNCATE"
	case ES_EVENT_TYPE_AUTH_UIPC_BIND:
		return "ES_EVENT_TYPE_AUTH_UIPC_BIND"
	case ES_EVENT_TYPE_AUTH_UIPC_CONNECT:
		return "ES_EVENT_TYPE_AUTH_UIPC_CONNECT"
	case ES_EVENT_TYPE_AUTH_UNLINK:
		return "ES_EVENT_TYPE_AUTH_UNLINK"
	case ES_EVENT_TYPE_AUTH_UTIMES:
		return "ES_EVENT_TYPE_AUTH_UTIMES"
	case ES_EVENT_TYPE_LAST:
		return "ES_EVENT_TYPE_LAST"
	case ES_EVENT_TYPE_NOTIFY_ACCESS:
		return "ES_EVENT_TYPE_NOTIFY_ACCESS"
	case ES_EVENT_TYPE_NOTIFY_AUTHENTICATION:
		return "ES_EVENT_TYPE_NOTIFY_AUTHENTICATION"
	case ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_JUDGEMENT:
		return "ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_JUDGEMENT"
	case ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_PETITION:
		return "ES_EVENT_TYPE_NOTIFY_AUTHORIZATION_PETITION"
	case ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_ADD:
		return "ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_ADD"
	case ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_REMOVE:
		return "ES_EVENT_TYPE_NOTIFY_BTM_LAUNCH_ITEM_REMOVE"
	case ES_EVENT_TYPE_NOTIFY_CHDIR:
		return "ES_EVENT_TYPE_NOTIFY_CHDIR"
	case ES_EVENT_TYPE_NOTIFY_CHROOT:
		return "ES_EVENT_TYPE_NOTIFY_CHROOT"
	case ES_EVENT_TYPE_NOTIFY_CLONE:
		return "ES_EVENT_TYPE_NOTIFY_CLONE"
	case ES_EVENT_TYPE_NOTIFY_CLOSE:
		return "ES_EVENT_TYPE_NOTIFY_CLOSE"
	case ES_EVENT_TYPE_NOTIFY_COPYFILE:
		return "ES_EVENT_TYPE_NOTIFY_COPYFILE"
	case ES_EVENT_TYPE_NOTIFY_CREATE:
		return "ES_EVENT_TYPE_NOTIFY_CREATE"
	case ES_EVENT_TYPE_NOTIFY_CS_INVALIDATED:
		return "ES_EVENT_TYPE_NOTIFY_CS_INVALIDATED"
	case ES_EVENT_TYPE_NOTIFY_DELETEEXTATTR:
		return "ES_EVENT_TYPE_NOTIFY_DELETEEXTATTR"
	case ES_EVENT_TYPE_NOTIFY_DUP:
		return "ES_EVENT_TYPE_NOTIFY_DUP"
	case ES_EVENT_TYPE_NOTIFY_EXCHANGEDATA:
		return "ES_EVENT_TYPE_NOTIFY_EXCHANGEDATA"
	case ES_EVENT_TYPE_NOTIFY_EXEC:
		return "ES_EVENT_TYPE_NOTIFY_EXEC"
	case ES_EVENT_TYPE_NOTIFY_EXIT:
		return "ES_EVENT_TYPE_NOTIFY_EXIT"
	case ES_EVENT_TYPE_NOTIFY_FCNTL:
		return "ES_EVENT_TYPE_NOTIFY_FCNTL"
	case ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_MATERIALIZE:
		return "ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_MATERIALIZE"
	case ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_UPDATE:
		return "ES_EVENT_TYPE_NOTIFY_FILE_PROVIDER_UPDATE"
	case ES_EVENT_TYPE_NOTIFY_FORK:
		return "ES_EVENT_TYPE_NOTIFY_FORK"
	case ES_EVENT_TYPE_NOTIFY_FSGETPATH:
		return "ES_EVENT_TYPE_NOTIFY_FSGETPATH"
	case ES_EVENT_TYPE_NOTIFY_GATEKEEPER_USER_OVERRIDE:
		return "ES_EVENT_TYPE_NOTIFY_GATEKEEPER_USER_OVERRIDE"
	case ES_EVENT_TYPE_NOTIFY_GETATTRLIST:
		return "ES_EVENT_TYPE_NOTIFY_GETATTRLIST"
	case ES_EVENT_TYPE_NOTIFY_GETEXTATTR:
		return "ES_EVENT_TYPE_NOTIFY_GETEXTATTR"
	case ES_EVENT_TYPE_NOTIFY_GET_TASK:
		return "ES_EVENT_TYPE_NOTIFY_GET_TASK"
	case ES_EVENT_TYPE_NOTIFY_GET_TASK_INSPECT:
		return "ES_EVENT_TYPE_NOTIFY_GET_TASK_INSPECT"
	case ES_EVENT_TYPE_NOTIFY_GET_TASK_NAME:
		return "ES_EVENT_TYPE_NOTIFY_GET_TASK_NAME"
	case ES_EVENT_TYPE_NOTIFY_GET_TASK_READ:
		return "ES_EVENT_TYPE_NOTIFY_GET_TASK_READ"
	case ES_EVENT_TYPE_NOTIFY_IOKIT_OPEN:
		return "ES_EVENT_TYPE_NOTIFY_IOKIT_OPEN"
	case ES_EVENT_TYPE_NOTIFY_KEXTLOAD:
		return "ES_EVENT_TYPE_NOTIFY_KEXTLOAD"
	case ES_EVENT_TYPE_NOTIFY_KEXTUNLOAD:
		return "ES_EVENT_TYPE_NOTIFY_KEXTUNLOAD"
	case ES_EVENT_TYPE_NOTIFY_LINK:
		return "ES_EVENT_TYPE_NOTIFY_LINK"
	case ES_EVENT_TYPE_NOTIFY_LISTEXTATTR:
		return "ES_EVENT_TYPE_NOTIFY_LISTEXTATTR"
	case ES_EVENT_TYPE_NOTIFY_LOGIN_LOGIN:
		return "ES_EVENT_TYPE_NOTIFY_LOGIN_LOGIN"
	case ES_EVENT_TYPE_NOTIFY_LOGIN_LOGOUT:
		return "ES_EVENT_TYPE_NOTIFY_LOGIN_LOGOUT"
	case ES_EVENT_TYPE_NOTIFY_LOOKUP:
		return "ES_EVENT_TYPE_NOTIFY_LOOKUP"
	case ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOCK:
		return "ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOCK"
	case ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGIN:
		return "ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGIN"
	case ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGOUT:
		return "ES_EVENT_TYPE_NOTIFY_LW_SESSION_LOGOUT"
	case ES_EVENT_TYPE_NOTIFY_LW_SESSION_UNLOCK:
		return "ES_EVENT_TYPE_NOTIFY_LW_SESSION_UNLOCK"
	case ES_EVENT_TYPE_NOTIFY_MMAP:
		return "ES_EVENT_TYPE_NOTIFY_MMAP"
	case ES_EVENT_TYPE_NOTIFY_MOUNT:
		return "ES_EVENT_TYPE_NOTIFY_MOUNT"
	case ES_EVENT_TYPE_NOTIFY_MPROTECT:
		return "ES_EVENT_TYPE_NOTIFY_MPROTECT"
	case ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_SET:
		return "ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_SET"
	case ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_ADD:
		return "ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_ADD"
	case ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_REMOVE:
		return "ES_EVENT_TYPE_NOTIFY_OD_ATTRIBUTE_VALUE_REMOVE"
	case ES_EVENT_TYPE_NOTIFY_OD_CREATE_GROUP:
		return "ES_EVENT_TYPE_NOTIFY_OD_CREATE_GROUP"
	case ES_EVENT_TYPE_NOTIFY_OD_CREATE_USER:
		return "ES_EVENT_TYPE_NOTIFY_OD_CREATE_USER"
	case ES_EVENT_TYPE_NOTIFY_OD_DELETE_GROUP:
		return "ES_EVENT_TYPE_NOTIFY_OD_DELETE_GROUP"
	case ES_EVENT_TYPE_NOTIFY_OD_DELETE_USER:
		return "ES_EVENT_TYPE_NOTIFY_OD_DELETE_USER"
	case ES_EVENT_TYPE_NOTIFY_OD_DISABLE_USER:
		return "ES_EVENT_TYPE_NOTIFY_OD_DISABLE_USER"
	case ES_EVENT_TYPE_NOTIFY_OD_ENABLE_USER:
		return "ES_EVENT_TYPE_NOTIFY_OD_ENABLE_USER"
	case ES_EVENT_TYPE_NOTIFY_OD_GROUP_ADD:
		return "ES_EVENT_TYPE_NOTIFY_OD_GROUP_ADD"
	case ES_EVENT_TYPE_NOTIFY_OD_GROUP_REMOVE:
		return "ES_EVENT_TYPE_NOTIFY_OD_GROUP_REMOVE"
	case ES_EVENT_TYPE_NOTIFY_OD_GROUP_SET:
		return "ES_EVENT_TYPE_NOTIFY_OD_GROUP_SET"
	case ES_EVENT_TYPE_NOTIFY_OD_MODIFY_PASSWORD:
		return "ES_EVENT_TYPE_NOTIFY_OD_MODIFY_PASSWORD"
	case ES_EVENT_TYPE_NOTIFY_OPEN:
		return "ES_EVENT_TYPE_NOTIFY_OPEN"
	case ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGIN:
		return "ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGIN"
	case ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGOUT:
		return "ES_EVENT_TYPE_NOTIFY_OPENSSH_LOGOUT"
	case ES_EVENT_TYPE_NOTIFY_PROC_CHECK:
		return "ES_EVENT_TYPE_NOTIFY_PROC_CHECK"
	case ES_EVENT_TYPE_NOTIFY_PROC_SUSPEND_RESUME:
		return "ES_EVENT_TYPE_NOTIFY_PROC_SUSPEND_RESUME"
	case ES_EVENT_TYPE_NOTIFY_PROFILE_ADD:
		return "ES_EVENT_TYPE_NOTIFY_PROFILE_ADD"
	case ES_EVENT_TYPE_NOTIFY_PROFILE_REMOVE:
		return "ES_EVENT_TYPE_NOTIFY_PROFILE_REMOVE"
	case ES_EVENT_TYPE_NOTIFY_PTY_CLOSE:
		return "ES_EVENT_TYPE_NOTIFY_PTY_CLOSE"
	case ES_EVENT_TYPE_NOTIFY_PTY_GRANT:
		return "ES_EVENT_TYPE_NOTIFY_PTY_GRANT"
	case ES_EVENT_TYPE_NOTIFY_READDIR:
		return "ES_EVENT_TYPE_NOTIFY_READDIR"
	case ES_EVENT_TYPE_NOTIFY_READLINK:
		return "ES_EVENT_TYPE_NOTIFY_READLINK"
	case ES_EVENT_TYPE_NOTIFY_REMOTE_THREAD_CREATE:
		return "ES_EVENT_TYPE_NOTIFY_REMOTE_THREAD_CREATE"
	case ES_EVENT_TYPE_NOTIFY_REMOUNT:
		return "ES_EVENT_TYPE_NOTIFY_REMOUNT"
	case ES_EVENT_TYPE_NOTIFY_RENAME:
		return "ES_EVENT_TYPE_NOTIFY_RENAME"
	case ES_EVENT_TYPE_NOTIFY_SCREENSHARING_ATTACH:
		return "ES_EVENT_TYPE_NOTIFY_SCREENSHARING_ATTACH"
	case ES_EVENT_TYPE_NOTIFY_SCREENSHARING_DETACH:
		return "ES_EVENT_TYPE_NOTIFY_SCREENSHARING_DETACH"
	case ES_EVENT_TYPE_NOTIFY_SEARCHFS:
		return "ES_EVENT_TYPE_NOTIFY_SEARCHFS"
	case ES_EVENT_TYPE_NOTIFY_SETACL:
		return "ES_EVENT_TYPE_NOTIFY_SETACL"
	case ES_EVENT_TYPE_NOTIFY_SETATTRLIST:
		return "ES_EVENT_TYPE_NOTIFY_SETATTRLIST"
	case ES_EVENT_TYPE_NOTIFY_SETEGID:
		return "ES_EVENT_TYPE_NOTIFY_SETEGID"
	case ES_EVENT_TYPE_NOTIFY_SETEUID:
		return "ES_EVENT_TYPE_NOTIFY_SETEUID"
	case ES_EVENT_TYPE_NOTIFY_SETEXTATTR:
		return "ES_EVENT_TYPE_NOTIFY_SETEXTATTR"
	case ES_EVENT_TYPE_NOTIFY_SETFLAGS:
		return "ES_EVENT_TYPE_NOTIFY_SETFLAGS"
	case ES_EVENT_TYPE_NOTIFY_SETGID:
		return "ES_EVENT_TYPE_NOTIFY_SETGID"
	case ES_EVENT_TYPE_NOTIFY_SETMODE:
		return "ES_EVENT_TYPE_NOTIFY_SETMODE"
	case ES_EVENT_TYPE_NOTIFY_SETOWNER:
		return "ES_EVENT_TYPE_NOTIFY_SETOWNER"
	case ES_EVENT_TYPE_NOTIFY_SETREGID:
		return "ES_EVENT_TYPE_NOTIFY_SETREGID"
	case ES_EVENT_TYPE_NOTIFY_SETREUID:
		return "ES_EVENT_TYPE_NOTIFY_SETREUID"
	case ES_EVENT_TYPE_NOTIFY_SETTIME:
		return "ES_EVENT_TYPE_NOTIFY_SETTIME"
	case ES_EVENT_TYPE_NOTIFY_SETUID:
		return "ES_EVENT_TYPE_NOTIFY_SETUID"
	case ES_EVENT_TYPE_NOTIFY_SIGNAL:
		return "ES_EVENT_TYPE_NOTIFY_SIGNAL"
	case ES_EVENT_TYPE_NOTIFY_STAT:
		return "ES_EVENT_TYPE_NOTIFY_STAT"
	case ES_EVENT_TYPE_NOTIFY_SU:
		return "ES_EVENT_TYPE_NOTIFY_SU"
	case ES_EVENT_TYPE_NOTIFY_SUDO:
		return "ES_EVENT_TYPE_NOTIFY_SUDO"
	case ES_EVENT_TYPE_NOTIFY_TCC_MODIFY:
		return "ES_EVENT_TYPE_NOTIFY_TCC_MODIFY"
	case ES_EVENT_TYPE_NOTIFY_TRACE:
		return "ES_EVENT_TYPE_NOTIFY_TRACE"
	case ES_EVENT_TYPE_NOTIFY_TRUNCATE:
		return "ES_EVENT_TYPE_NOTIFY_TRUNCATE"
	case ES_EVENT_TYPE_NOTIFY_UIPC_BIND:
		return "ES_EVENT_TYPE_NOTIFY_UIPC_BIND"
	case ES_EVENT_TYPE_NOTIFY_UIPC_CONNECT:
		return "ES_EVENT_TYPE_NOTIFY_UIPC_CONNECT"
	case ES_EVENT_TYPE_NOTIFY_UNLINK:
		return "ES_EVENT_TYPE_NOTIFY_UNLINK"
	case ES_EVENT_TYPE_NOTIFY_UNMOUNT:
		return "ES_EVENT_TYPE_NOTIFY_UNMOUNT"
	case ES_EVENT_TYPE_NOTIFY_UTIMES:
		return "ES_EVENT_TYPE_NOTIFY_UTIMES"
	case ES_EVENT_TYPE_NOTIFY_WRITE:
		return "ES_EVENT_TYPE_NOTIFY_WRITE"
	case ES_EVENT_TYPE_NOTIFY_XPC_CONNECT:
		return "ES_EVENT_TYPE_NOTIFY_XPC_CONNECT"
	case ES_EVENT_TYPE_NOTIFY_XP_MALWARE_DETECTED:
		return "ES_EVENT_TYPE_NOTIFY_XP_MALWARE_DETECTED"
	case ES_EVENT_TYPE_NOTIFY_XP_MALWARE_REMEDIATED:
		return "ES_EVENT_TYPE_NOTIFY_XP_MALWARE_REMEDIATED"
	case ES_EVENT_TYPE_RESERVED_0:
		return "ES_EVENT_TYPE_RESERVED_0"
	case ES_EVENT_TYPE_RESERVED_1:
		return "ES_EVENT_TYPE_RESERVED_1"
	case ES_EVENT_TYPE_RESERVED_2:
		return "ES_EVENT_TYPE_RESERVED_2"
	case ES_EVENT_TYPE_RESERVED_3:
		return "ES_EVENT_TYPE_RESERVED_3"
	case ES_EVENT_TYPE_RESERVED_4:
		return "ES_EVENT_TYPE_RESERVED_4"
	case ES_EVENT_TYPE_RESERVED_5:
		return "ES_EVENT_TYPE_RESERVED_5"
	case ES_EVENT_TYPE_RESERVED_6:
		return "ES_EVENT_TYPE_RESERVED_6"
	default:
		return fmt.Sprintf("EsEventType(%d)", e)
	}
}

type EsGatekeeperUserOverrideFileType int32

const (
	ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_FILE EsGatekeeperUserOverrideFileType = 1
	ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_PATH EsGatekeeperUserOverrideFileType = 0
)

func (e EsGatekeeperUserOverrideFileType) String() string {
	switch e {
	case ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_FILE:
		return "ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_FILE"
	case ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_PATH:
		return "ES_GATEKEEPER_USER_OVERRIDE_FILE_TYPE_PATH"
	default:
		return fmt.Sprintf("EsGatekeeperUserOverrideFileType(%d)", e)
	}
}

type EsGetTaskType int32

const (
	ES_GET_TASK_TYPE_EXPOSE_TASK    EsGetTaskType = 1
	ES_GET_TASK_TYPE_IDENTITY_TOKEN EsGetTaskType = 2
	ES_GET_TASK_TYPE_TASK_FOR_PID   EsGetTaskType = 0
)

func (e EsGetTaskType) String() string {
	switch e {
	case ES_GET_TASK_TYPE_EXPOSE_TASK:
		return "ES_GET_TASK_TYPE_EXPOSE_TASK"
	case ES_GET_TASK_TYPE_IDENTITY_TOKEN:
		return "ES_GET_TASK_TYPE_IDENTITY_TOKEN"
	case ES_GET_TASK_TYPE_TASK_FOR_PID:
		return "ES_GET_TASK_TYPE_TASK_FOR_PID"
	default:
		return fmt.Sprintf("EsGetTaskType(%d)", e)
	}
}

type EsMountDisposition int32

const (
	ES_MOUNT_DISPOSITION_EXTERNAL EsMountDisposition = 0
	ES_MOUNT_DISPOSITION_INTERNAL EsMountDisposition = 1
	ES_MOUNT_DISPOSITION_NETWORK  EsMountDisposition = 2
	ES_MOUNT_DISPOSITION_NULLFS   EsMountDisposition = 4
	ES_MOUNT_DISPOSITION_UNKNOWN  EsMountDisposition = 5
	ES_MOUNT_DISPOSITION_VIRTUAL  EsMountDisposition = 3
)

func (e EsMountDisposition) String() string {
	switch e {
	case ES_MOUNT_DISPOSITION_EXTERNAL:
		return "ES_MOUNT_DISPOSITION_EXTERNAL"
	case ES_MOUNT_DISPOSITION_INTERNAL:
		return "ES_MOUNT_DISPOSITION_INTERNAL"
	case ES_MOUNT_DISPOSITION_NETWORK:
		return "ES_MOUNT_DISPOSITION_NETWORK"
	case ES_MOUNT_DISPOSITION_NULLFS:
		return "ES_MOUNT_DISPOSITION_NULLFS"
	case ES_MOUNT_DISPOSITION_UNKNOWN:
		return "ES_MOUNT_DISPOSITION_UNKNOWN"
	case ES_MOUNT_DISPOSITION_VIRTUAL:
		return "ES_MOUNT_DISPOSITION_VIRTUAL"
	default:
		return fmt.Sprintf("EsMountDisposition(%d)", e)
	}
}

type EsMute int32

const (
	ES_MUTE_INVERTED       EsMute = 0
	ES_MUTE_INVERTED_ERROR EsMute = 2
	ES_MUTE_NOT_INVERTED   EsMute = 1
)

func (e EsMute) String() string {
	switch e {
	case ES_MUTE_INVERTED:
		return "ES_MUTE_INVERTED"
	case ES_MUTE_INVERTED_ERROR:
		return "ES_MUTE_INVERTED_ERROR"
	case ES_MUTE_NOT_INVERTED:
		return "ES_MUTE_NOT_INVERTED"
	default:
		return fmt.Sprintf("EsMute(%d)", e)
	}
}

type EsMuteInversionType int32

const (
	ES_MUTE_INVERSION_TYPE_LAST        EsMuteInversionType = 3
	ES_MUTE_INVERSION_TYPE_PATH        EsMuteInversionType = 1
	ES_MUTE_INVERSION_TYPE_PROCESS     EsMuteInversionType = 0
	ES_MUTE_INVERSION_TYPE_TARGET_PATH EsMuteInversionType = 2
)

func (e EsMuteInversionType) String() string {
	switch e {
	case ES_MUTE_INVERSION_TYPE_LAST:
		return "ES_MUTE_INVERSION_TYPE_LAST"
	case ES_MUTE_INVERSION_TYPE_PATH:
		return "ES_MUTE_INVERSION_TYPE_PATH"
	case ES_MUTE_INVERSION_TYPE_PROCESS:
		return "ES_MUTE_INVERSION_TYPE_PROCESS"
	case ES_MUTE_INVERSION_TYPE_TARGET_PATH:
		return "ES_MUTE_INVERSION_TYPE_TARGET_PATH"
	default:
		return fmt.Sprintf("EsMuteInversionType(%d)", e)
	}
}

type EsMutePathType int32

const (
	// ES_MUTE_PATH_TYPE_LITERAL: A type for a path string used as a path literal.
	ES_MUTE_PATH_TYPE_LITERAL EsMutePathType = 1
	// ES_MUTE_PATH_TYPE_PREFIX: A type for a path string used as a prefix.
	ES_MUTE_PATH_TYPE_PREFIX         EsMutePathType = 0
	ES_MUTE_PATH_TYPE_TARGET_LITERAL EsMutePathType = 3
	ES_MUTE_PATH_TYPE_TARGET_PREFIX  EsMutePathType = 2
)

func (e EsMutePathType) String() string {
	switch e {
	case ES_MUTE_PATH_TYPE_LITERAL:
		return "ES_MUTE_PATH_TYPE_LITERAL"
	case ES_MUTE_PATH_TYPE_PREFIX:
		return "ES_MUTE_PATH_TYPE_PREFIX"
	case ES_MUTE_PATH_TYPE_TARGET_LITERAL:
		return "ES_MUTE_PATH_TYPE_TARGET_LITERAL"
	case ES_MUTE_PATH_TYPE_TARGET_PREFIX:
		return "ES_MUTE_PATH_TYPE_TARGET_PREFIX"
	default:
		return fmt.Sprintf("EsMutePathType(%d)", e)
	}
}

type EsNewClientResult int32

const (
	// ES_NEW_CLIENT_RESULT_ERR_INTERNAL: Communication with the Endpoint Security subsystem failed.
	ES_NEW_CLIENT_RESULT_ERR_INTERNAL EsNewClientResult = 2
	// ES_NEW_CLIENT_RESULT_ERR_INVALID_ARGUMENT: The attempt to create a new client contained one or more invalid arguments.
	ES_NEW_CLIENT_RESULT_ERR_INVALID_ARGUMENT EsNewClientResult = 1
	// ES_NEW_CLIENT_RESULT_ERR_NOT_ENTITLED: The caller isn’t properly entitled to connect to Endpoint Security.
	ES_NEW_CLIENT_RESULT_ERR_NOT_ENTITLED EsNewClientResult = 3
	// ES_NEW_CLIENT_RESULT_ERR_NOT_PERMITTED: The caller isn’t permitted to connect to Endpoint Security.
	ES_NEW_CLIENT_RESULT_ERR_NOT_PERMITTED EsNewClientResult = 4
	// ES_NEW_CLIENT_RESULT_ERR_NOT_PRIVILEGED: The caller isn’t running as root.
	ES_NEW_CLIENT_RESULT_ERR_NOT_PRIVILEGED EsNewClientResult = 5
	// ES_NEW_CLIENT_RESULT_ERR_TOO_MANY_CLIENTS: The caller has reached the maximum allowed number of simultaneously connected clients.
	ES_NEW_CLIENT_RESULT_ERR_TOO_MANY_CLIENTS EsNewClientResult = 6
	// ES_NEW_CLIENT_RESULT_SUCCESS: Endpoint Security successfully created the new client.
	ES_NEW_CLIENT_RESULT_SUCCESS EsNewClientResult = 0
)

func (e EsNewClientResult) String() string {
	switch e {
	case ES_NEW_CLIENT_RESULT_ERR_INTERNAL:
		return "ES_NEW_CLIENT_RESULT_ERR_INTERNAL"
	case ES_NEW_CLIENT_RESULT_ERR_INVALID_ARGUMENT:
		return "ES_NEW_CLIENT_RESULT_ERR_INVALID_ARGUMENT"
	case ES_NEW_CLIENT_RESULT_ERR_NOT_ENTITLED:
		return "ES_NEW_CLIENT_RESULT_ERR_NOT_ENTITLED"
	case ES_NEW_CLIENT_RESULT_ERR_NOT_PERMITTED:
		return "ES_NEW_CLIENT_RESULT_ERR_NOT_PERMITTED"
	case ES_NEW_CLIENT_RESULT_ERR_NOT_PRIVILEGED:
		return "ES_NEW_CLIENT_RESULT_ERR_NOT_PRIVILEGED"
	case ES_NEW_CLIENT_RESULT_ERR_TOO_MANY_CLIENTS:
		return "ES_NEW_CLIENT_RESULT_ERR_TOO_MANY_CLIENTS"
	case ES_NEW_CLIENT_RESULT_SUCCESS:
		return "ES_NEW_CLIENT_RESULT_SUCCESS"
	default:
		return fmt.Sprintf("EsNewClientResult(%d)", e)
	}
}

type EsOdAccountType int32

const (
	ES_OD_ACCOUNT_TYPE_COMPUTER EsOdAccountType = 1
	ES_OD_ACCOUNT_TYPE_USER     EsOdAccountType = 0
)

func (e EsOdAccountType) String() string {
	switch e {
	case ES_OD_ACCOUNT_TYPE_COMPUTER:
		return "ES_OD_ACCOUNT_TYPE_COMPUTER"
	case ES_OD_ACCOUNT_TYPE_USER:
		return "ES_OD_ACCOUNT_TYPE_USER"
	default:
		return fmt.Sprintf("EsOdAccountType(%d)", e)
	}
}

type EsOdMemberType int32

const (
	ES_OD_MEMBER_TYPE_GROUP_UUID EsOdMemberType = 2
	ES_OD_MEMBER_TYPE_USER_NAME  EsOdMemberType = 0
	ES_OD_MEMBER_TYPE_USER_UUID  EsOdMemberType = 1
)

func (e EsOdMemberType) String() string {
	switch e {
	case ES_OD_MEMBER_TYPE_GROUP_UUID:
		return "ES_OD_MEMBER_TYPE_GROUP_UUID"
	case ES_OD_MEMBER_TYPE_USER_NAME:
		return "ES_OD_MEMBER_TYPE_USER_NAME"
	case ES_OD_MEMBER_TYPE_USER_UUID:
		return "ES_OD_MEMBER_TYPE_USER_UUID"
	default:
		return fmt.Sprintf("EsOdMemberType(%d)", e)
	}
}

type EsOdRecordType int32

const (
	ES_OD_RECORD_TYPE_GROUP EsOdRecordType = 1
	ES_OD_RECORD_TYPE_USER  EsOdRecordType = 0
)

func (e EsOdRecordType) String() string {
	switch e {
	case ES_OD_RECORD_TYPE_GROUP:
		return "ES_OD_RECORD_TYPE_GROUP"
	case ES_OD_RECORD_TYPE_USER:
		return "ES_OD_RECORD_TYPE_USER"
	default:
		return fmt.Sprintf("EsOdRecordType(%d)", e)
	}
}

type EsOpenssh int32

const (
	ES_OPENSSH_AUTH_FAIL_GSSAPI      EsOpenssh = 8
	ES_OPENSSH_AUTH_FAIL_HOSTBASED   EsOpenssh = 7
	ES_OPENSSH_AUTH_FAIL_KBDINT      EsOpenssh = 5
	ES_OPENSSH_AUTH_FAIL_NONE        EsOpenssh = 3
	ES_OPENSSH_AUTH_FAIL_PASSWD      EsOpenssh = 4
	ES_OPENSSH_AUTH_FAIL_PUBKEY      EsOpenssh = 6
	ES_OPENSSH_AUTH_SUCCESS          EsOpenssh = 2
	ES_OPENSSH_INVALID_USER          EsOpenssh = 9
	ES_OPENSSH_LOGIN_EXCEED_MAXTRIES EsOpenssh = 0
	ES_OPENSSH_LOGIN_ROOT_DENIED     EsOpenssh = 1
)

func (e EsOpenssh) String() string {
	switch e {
	case ES_OPENSSH_AUTH_FAIL_GSSAPI:
		return "ES_OPENSSH_AUTH_FAIL_GSSAPI"
	case ES_OPENSSH_AUTH_FAIL_HOSTBASED:
		return "ES_OPENSSH_AUTH_FAIL_HOSTBASED"
	case ES_OPENSSH_AUTH_FAIL_KBDINT:
		return "ES_OPENSSH_AUTH_FAIL_KBDINT"
	case ES_OPENSSH_AUTH_FAIL_NONE:
		return "ES_OPENSSH_AUTH_FAIL_NONE"
	case ES_OPENSSH_AUTH_FAIL_PASSWD:
		return "ES_OPENSSH_AUTH_FAIL_PASSWD"
	case ES_OPENSSH_AUTH_FAIL_PUBKEY:
		return "ES_OPENSSH_AUTH_FAIL_PUBKEY"
	case ES_OPENSSH_AUTH_SUCCESS:
		return "ES_OPENSSH_AUTH_SUCCESS"
	case ES_OPENSSH_INVALID_USER:
		return "ES_OPENSSH_INVALID_USER"
	case ES_OPENSSH_LOGIN_EXCEED_MAXTRIES:
		return "ES_OPENSSH_LOGIN_EXCEED_MAXTRIES"
	case ES_OPENSSH_LOGIN_ROOT_DENIED:
		return "ES_OPENSSH_LOGIN_ROOT_DENIED"
	default:
		return fmt.Sprintf("EsOpenssh(%d)", e)
	}
}

type EsProcCheckType int32

const (
	// ES_PROC_CHECK_TYPE_DIRTYCONTROL: A type of process check that uses the process’s dirty state.
	ES_PROC_CHECK_TYPE_DIRTYCONTROL EsProcCheckType = 8
	// ES_PROC_CHECK_TYPE_KERNMSGBUF: A type of process check that checks the message buffer.
	ES_PROC_CHECK_TYPE_KERNMSGBUF EsProcCheckType = 4
	// ES_PROC_CHECK_TYPE_LISTPIDS: A type of process check that lists related process identifiers.
	ES_PROC_CHECK_TYPE_LISTPIDS EsProcCheckType = 1
	// ES_PROC_CHECK_TYPE_PIDFDINFO: A type of process check that gets file descriptor information.
	ES_PROC_CHECK_TYPE_PIDFDINFO EsProcCheckType = 3
	// ES_PROC_CHECK_TYPE_PIDFILEPORTINFO: A type of process check that gets port information.
	ES_PROC_CHECK_TYPE_PIDFILEPORTINFO EsProcCheckType = 6
	// ES_PROC_CHECK_TYPE_PIDINFO: A type of process check that gets basic process information.
	ES_PROC_CHECK_TYPE_PIDINFO EsProcCheckType = 2
	// ES_PROC_CHECK_TYPE_PIDRUSAGE: A type of process check that gets a process’s resource usage information.
	ES_PROC_CHECK_TYPE_PIDRUSAGE EsProcCheckType = 9
	// ES_PROC_CHECK_TYPE_SETCONTROL: A type of process check that sets the process control state.
	ES_PROC_CHECK_TYPE_SETCONTROL EsProcCheckType = 5
	// ES_PROC_CHECK_TYPE_TERMINATE: A type of process check that terninates a process.
	ES_PROC_CHECK_TYPE_TERMINATE EsProcCheckType = 7
	// ES_PROC_CHECK_TYPE_UDATA_INFO: A type of process check that involves a user data token.
	ES_PROC_CHECK_TYPE_UDATA_INFO EsProcCheckType = 14
)

func (e EsProcCheckType) String() string {
	switch e {
	case ES_PROC_CHECK_TYPE_DIRTYCONTROL:
		return "ES_PROC_CHECK_TYPE_DIRTYCONTROL"
	case ES_PROC_CHECK_TYPE_KERNMSGBUF:
		return "ES_PROC_CHECK_TYPE_KERNMSGBUF"
	case ES_PROC_CHECK_TYPE_LISTPIDS:
		return "ES_PROC_CHECK_TYPE_LISTPIDS"
	case ES_PROC_CHECK_TYPE_PIDFDINFO:
		return "ES_PROC_CHECK_TYPE_PIDFDINFO"
	case ES_PROC_CHECK_TYPE_PIDFILEPORTINFO:
		return "ES_PROC_CHECK_TYPE_PIDFILEPORTINFO"
	case ES_PROC_CHECK_TYPE_PIDINFO:
		return "ES_PROC_CHECK_TYPE_PIDINFO"
	case ES_PROC_CHECK_TYPE_PIDRUSAGE:
		return "ES_PROC_CHECK_TYPE_PIDRUSAGE"
	case ES_PROC_CHECK_TYPE_SETCONTROL:
		return "ES_PROC_CHECK_TYPE_SETCONTROL"
	case ES_PROC_CHECK_TYPE_TERMINATE:
		return "ES_PROC_CHECK_TYPE_TERMINATE"
	case ES_PROC_CHECK_TYPE_UDATA_INFO:
		return "ES_PROC_CHECK_TYPE_UDATA_INFO"
	default:
		return fmt.Sprintf("EsProcCheckType(%d)", e)
	}
}

type EsProcSuspendResumeType int32

const (
	// ES_PROC_SUSPEND_RESUME_TYPE_RESUME: An event type for process resumption events.
	ES_PROC_SUSPEND_RESUME_TYPE_RESUME EsProcSuspendResumeType = 1
	// ES_PROC_SUSPEND_RESUME_TYPE_SHUTDOWN_SOCKETS: An event type for process socket shutdown events.
	ES_PROC_SUSPEND_RESUME_TYPE_SHUTDOWN_SOCKETS EsProcSuspendResumeType = 3
	// ES_PROC_SUSPEND_RESUME_TYPE_SUSPEND: An event type for process suspension events.
	ES_PROC_SUSPEND_RESUME_TYPE_SUSPEND EsProcSuspendResumeType = 0
)

func (e EsProcSuspendResumeType) String() string {
	switch e {
	case ES_PROC_SUSPEND_RESUME_TYPE_RESUME:
		return "ES_PROC_SUSPEND_RESUME_TYPE_RESUME"
	case ES_PROC_SUSPEND_RESUME_TYPE_SHUTDOWN_SOCKETS:
		return "ES_PROC_SUSPEND_RESUME_TYPE_SHUTDOWN_SOCKETS"
	case ES_PROC_SUSPEND_RESUME_TYPE_SUSPEND:
		return "ES_PROC_SUSPEND_RESUME_TYPE_SUSPEND"
	default:
		return fmt.Sprintf("EsProcSuspendResumeType(%d)", e)
	}
}

type EsProfileSource int32

const (
	ES_PROFILE_SOURCE_INSTALL EsProfileSource = 1
	ES_PROFILE_SOURCE_MANAGED EsProfileSource = 0
)

func (e EsProfileSource) String() string {
	switch e {
	case ES_PROFILE_SOURCE_INSTALL:
		return "ES_PROFILE_SOURCE_INSTALL"
	case ES_PROFILE_SOURCE_MANAGED:
		return "ES_PROFILE_SOURCE_MANAGED"
	default:
		return fmt.Sprintf("EsProfileSource(%d)", e)
	}
}

type EsRespondResult int32

const (
	// ES_RESPOND_RESULT_ERR_DUPLICATE_RESPONSE: The caller responded to a message that already received a response.
	ES_RESPOND_RESULT_ERR_DUPLICATE_RESPONSE EsRespondResult = 4
	// ES_RESPOND_RESULT_ERR_EVENT_TYPE: The caller performed an inappropriate response to the event.
	ES_RESPOND_RESULT_ERR_EVENT_TYPE EsRespondResult = 5
	// ES_RESPOND_RESULT_ERR_INTERNAL: Communication with the Endpoint Security system failed.
	ES_RESPOND_RESULT_ERR_INTERNAL EsRespondResult = 2
	// ES_RESPOND_RESULT_ERR_INVALID_ARGUMENT: The caller provided one or more invalid arguments.
	ES_RESPOND_RESULT_ERR_INVALID_ARGUMENT EsRespondResult = 1
	// ES_RESPOND_RESULT_NOT_FOUND: The system couldn’t find the message that the caller sent this response to.
	ES_RESPOND_RESULT_NOT_FOUND EsRespondResult = 3
	// ES_RESPOND_RESULT_SUCCESS: Endpoint Security successfully delivered the response.
	ES_RESPOND_RESULT_SUCCESS EsRespondResult = 0
)

func (e EsRespondResult) String() string {
	switch e {
	case ES_RESPOND_RESULT_ERR_DUPLICATE_RESPONSE:
		return "ES_RESPOND_RESULT_ERR_DUPLICATE_RESPONSE"
	case ES_RESPOND_RESULT_ERR_EVENT_TYPE:
		return "ES_RESPOND_RESULT_ERR_EVENT_TYPE"
	case ES_RESPOND_RESULT_ERR_INTERNAL:
		return "ES_RESPOND_RESULT_ERR_INTERNAL"
	case ES_RESPOND_RESULT_ERR_INVALID_ARGUMENT:
		return "ES_RESPOND_RESULT_ERR_INVALID_ARGUMENT"
	case ES_RESPOND_RESULT_NOT_FOUND:
		return "ES_RESPOND_RESULT_NOT_FOUND"
	case ES_RESPOND_RESULT_SUCCESS:
		return "ES_RESPOND_RESULT_SUCCESS"
	default:
		return fmt.Sprintf("EsRespondResult(%d)", e)
	}
}

type EsResultType int32

const (
	// ES_RESULT_TYPE_AUTH: The authorization result type.
	ES_RESULT_TYPE_AUTH EsResultType = 0
	// ES_RESULT_TYPE_FLAGS: The flags result type.
	ES_RESULT_TYPE_FLAGS EsResultType = 1
)

func (e EsResultType) String() string {
	switch e {
	case ES_RESULT_TYPE_AUTH:
		return "ES_RESULT_TYPE_AUTH"
	case ES_RESULT_TYPE_FLAGS:
		return "ES_RESULT_TYPE_FLAGS"
	default:
		return fmt.Sprintf("EsResultType(%d)", e)
	}
}

type EsReturn int32

const (
	// ES_RETURN_ERROR: The action failed with an error.
	ES_RETURN_ERROR EsReturn = 1
	// ES_RETURN_SUCCESS: The action succeeded.
	ES_RETURN_SUCCESS EsReturn = 0
)

func (e EsReturn) String() string {
	switch e {
	case ES_RETURN_ERROR:
		return "ES_RETURN_ERROR"
	case ES_RETURN_SUCCESS:
		return "ES_RETURN_SUCCESS"
	default:
		return fmt.Sprintf("EsReturn(%d)", e)
	}
}

type EsSudoPluginType int32

const (
	ES_SUDO_PLUGIN_TYPE_APPROVAL  EsSudoPluginType = 5
	ES_SUDO_PLUGIN_TYPE_AUDIT     EsSudoPluginType = 4
	ES_SUDO_PLUGIN_TYPE_FRONT_END EsSudoPluginType = 1
	ES_SUDO_PLUGIN_TYPE_IO        EsSudoPluginType = 3
	ES_SUDO_PLUGIN_TYPE_POLICY    EsSudoPluginType = 2
	ES_SUDO_PLUGIN_TYPE_UNKNOWN   EsSudoPluginType = 0
)

func (e EsSudoPluginType) String() string {
	switch e {
	case ES_SUDO_PLUGIN_TYPE_APPROVAL:
		return "ES_SUDO_PLUGIN_TYPE_APPROVAL"
	case ES_SUDO_PLUGIN_TYPE_AUDIT:
		return "ES_SUDO_PLUGIN_TYPE_AUDIT"
	case ES_SUDO_PLUGIN_TYPE_FRONT_END:
		return "ES_SUDO_PLUGIN_TYPE_FRONT_END"
	case ES_SUDO_PLUGIN_TYPE_IO:
		return "ES_SUDO_PLUGIN_TYPE_IO"
	case ES_SUDO_PLUGIN_TYPE_POLICY:
		return "ES_SUDO_PLUGIN_TYPE_POLICY"
	case ES_SUDO_PLUGIN_TYPE_UNKNOWN:
		return "ES_SUDO_PLUGIN_TYPE_UNKNOWN"
	default:
		return fmt.Sprintf("EsSudoPluginType(%d)", e)
	}
}

type EsTccAuthorizationReason int32

const (
	// ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY EsTccAuthorizationReason = 12
	// ES_TCC_AUTHORIZATION_REASON_ENTITLED: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_ENTITLED EsTccAuthorizationReason = 11
	ES_TCC_AUTHORIZATION_REASON_ERROR    EsTccAuthorizationReason = 1
	// ES_TCC_AUTHORIZATION_REASON_MDM_POLICY: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_MDM_POLICY EsTccAuthorizationReason = 6
	// ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING EsTccAuthorizationReason = 8
	ES_TCC_AUTHORIZATION_REASON_NONE                 EsTccAuthorizationReason = 0
	// ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN EsTccAuthorizationReason = 10
	// ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL EsTccAuthorizationReason = 13
	// ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT EsTccAuthorizationReason = 9
	// ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY EsTccAuthorizationReason = 7
	// ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY: A system process changed the authorization right
	ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY EsTccAuthorizationReason = 5
	// ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET: User changed the authorization right via Preferences
	ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET   EsTccAuthorizationReason = 4
	ES_TCC_AUTHORIZATION_REASON_USER_CONSENT EsTccAuthorizationReason = 2
	// ES_TCC_AUTHORIZATION_REASON_USER_SET: User answered a prompt
	ES_TCC_AUTHORIZATION_REASON_USER_SET EsTccAuthorizationReason = 3
)

func (e EsTccAuthorizationReason) String() string {
	switch e {
	case ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY:
		return "ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY"
	case ES_TCC_AUTHORIZATION_REASON_ENTITLED:
		return "ES_TCC_AUTHORIZATION_REASON_ENTITLED"
	case ES_TCC_AUTHORIZATION_REASON_ERROR:
		return "ES_TCC_AUTHORIZATION_REASON_ERROR"
	case ES_TCC_AUTHORIZATION_REASON_MDM_POLICY:
		return "ES_TCC_AUTHORIZATION_REASON_MDM_POLICY"
	case ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING:
		return "ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING"
	case ES_TCC_AUTHORIZATION_REASON_NONE:
		return "ES_TCC_AUTHORIZATION_REASON_NONE"
	case ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN:
		return "ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN"
	case ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL:
		return "ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL"
	case ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT:
		return "ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT"
	case ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY:
		return "ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY"
	case ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY:
		return "ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY"
	case ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET:
		return "ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET"
	case ES_TCC_AUTHORIZATION_REASON_USER_CONSENT:
		return "ES_TCC_AUTHORIZATION_REASON_USER_CONSENT"
	case ES_TCC_AUTHORIZATION_REASON_USER_SET:
		return "ES_TCC_AUTHORIZATION_REASON_USER_SET"
	default:
		return fmt.Sprintf("EsTccAuthorizationReason(%d)", e)
	}
}

type EsTccAuthorizationRight int32

const (
	ES_TCC_AUTHORIZATION_RIGHT_ADD_MODIFY_ADDED EsTccAuthorizationRight = 4
	ES_TCC_AUTHORIZATION_RIGHT_ALLOWED          EsTccAuthorizationRight = 2
	ES_TCC_AUTHORIZATION_RIGHT_DENIED           EsTccAuthorizationRight = 0
	ES_TCC_AUTHORIZATION_RIGHT_LEARN_MORE       EsTccAuthorizationRight = 6
	ES_TCC_AUTHORIZATION_RIGHT_LIMITED          EsTccAuthorizationRight = 3
	ES_TCC_AUTHORIZATION_RIGHT_SESSION_PID      EsTccAuthorizationRight = 5
	ES_TCC_AUTHORIZATION_RIGHT_UNKNOWN          EsTccAuthorizationRight = 1
)

func (e EsTccAuthorizationRight) String() string {
	switch e {
	case ES_TCC_AUTHORIZATION_RIGHT_ADD_MODIFY_ADDED:
		return "ES_TCC_AUTHORIZATION_RIGHT_ADD_MODIFY_ADDED"
	case ES_TCC_AUTHORIZATION_RIGHT_ALLOWED:
		return "ES_TCC_AUTHORIZATION_RIGHT_ALLOWED"
	case ES_TCC_AUTHORIZATION_RIGHT_DENIED:
		return "ES_TCC_AUTHORIZATION_RIGHT_DENIED"
	case ES_TCC_AUTHORIZATION_RIGHT_LEARN_MORE:
		return "ES_TCC_AUTHORIZATION_RIGHT_LEARN_MORE"
	case ES_TCC_AUTHORIZATION_RIGHT_LIMITED:
		return "ES_TCC_AUTHORIZATION_RIGHT_LIMITED"
	case ES_TCC_AUTHORIZATION_RIGHT_SESSION_PID:
		return "ES_TCC_AUTHORIZATION_RIGHT_SESSION_PID"
	case ES_TCC_AUTHORIZATION_RIGHT_UNKNOWN:
		return "ES_TCC_AUTHORIZATION_RIGHT_UNKNOWN"
	default:
		return fmt.Sprintf("EsTccAuthorizationRight(%d)", e)
	}
}

type EsTccEventType int32

const (
	ES_TCC_EVENT_TYPE_CREATE  EsTccEventType = 1
	ES_TCC_EVENT_TYPE_DELETE  EsTccEventType = 3
	ES_TCC_EVENT_TYPE_MODIFY  EsTccEventType = 2
	ES_TCC_EVENT_TYPE_UNKNOWN EsTccEventType = 0
)

func (e EsTccEventType) String() string {
	switch e {
	case ES_TCC_EVENT_TYPE_CREATE:
		return "ES_TCC_EVENT_TYPE_CREATE"
	case ES_TCC_EVENT_TYPE_DELETE:
		return "ES_TCC_EVENT_TYPE_DELETE"
	case ES_TCC_EVENT_TYPE_MODIFY:
		return "ES_TCC_EVENT_TYPE_MODIFY"
	case ES_TCC_EVENT_TYPE_UNKNOWN:
		return "ES_TCC_EVENT_TYPE_UNKNOWN"
	default:
		return fmt.Sprintf("EsTccEventType(%d)", e)
	}
}

type EsTccIdentityType int32

const (
	ES_TCC_IDENTITY_TYPE_BUNDLE_ID               EsTccIdentityType = 0
	ES_TCC_IDENTITY_TYPE_EXECUTABLE_PATH         EsTccIdentityType = 1
	ES_TCC_IDENTITY_TYPE_FILE_PROVIDER_DOMAIN_ID EsTccIdentityType = 3
	ES_TCC_IDENTITY_TYPE_POLICY_ID               EsTccIdentityType = 2
)

func (e EsTccIdentityType) String() string {
	switch e {
	case ES_TCC_IDENTITY_TYPE_BUNDLE_ID:
		return "ES_TCC_IDENTITY_TYPE_BUNDLE_ID"
	case ES_TCC_IDENTITY_TYPE_EXECUTABLE_PATH:
		return "ES_TCC_IDENTITY_TYPE_EXECUTABLE_PATH"
	case ES_TCC_IDENTITY_TYPE_FILE_PROVIDER_DOMAIN_ID:
		return "ES_TCC_IDENTITY_TYPE_FILE_PROVIDER_DOMAIN_ID"
	case ES_TCC_IDENTITY_TYPE_POLICY_ID:
		return "ES_TCC_IDENTITY_TYPE_POLICY_ID"
	default:
		return fmt.Sprintf("EsTccIdentityType(%d)", e)
	}
}

type EsTouchidMode int32

const (
	ES_TOUCHID_MODE_IDENTIFICATION EsTouchidMode = 1
	ES_TOUCHID_MODE_VERIFICATION   EsTouchidMode = 0
)

func (e EsTouchidMode) String() string {
	switch e {
	case ES_TOUCHID_MODE_IDENTIFICATION:
		return "ES_TOUCHID_MODE_IDENTIFICATION"
	case ES_TOUCHID_MODE_VERIFICATION:
		return "ES_TOUCHID_MODE_VERIFICATION"
	default:
		return fmt.Sprintf("EsTouchidMode(%d)", e)
	}
}

type EsXPCDomainType int32

const (
	ES_XPC_DOMAIN_TYPE_GUI        EsXPCDomainType = 8
	ES_XPC_DOMAIN_TYPE_MANAGER    EsXPCDomainType = 6
	ES_XPC_DOMAIN_TYPE_PID        EsXPCDomainType = 5
	ES_XPC_DOMAIN_TYPE_PORT       EsXPCDomainType = 7
	ES_XPC_DOMAIN_TYPE_SESSION    EsXPCDomainType = 4
	ES_XPC_DOMAIN_TYPE_SYSTEM     EsXPCDomainType = 1
	ES_XPC_DOMAIN_TYPE_USER       EsXPCDomainType = 2
	ES_XPC_DOMAIN_TYPE_USER_LOGIN EsXPCDomainType = 3
)

func (e EsXPCDomainType) String() string {
	switch e {
	case ES_XPC_DOMAIN_TYPE_GUI:
		return "ES_XPC_DOMAIN_TYPE_GUI"
	case ES_XPC_DOMAIN_TYPE_MANAGER:
		return "ES_XPC_DOMAIN_TYPE_MANAGER"
	case ES_XPC_DOMAIN_TYPE_PID:
		return "ES_XPC_DOMAIN_TYPE_PID"
	case ES_XPC_DOMAIN_TYPE_PORT:
		return "ES_XPC_DOMAIN_TYPE_PORT"
	case ES_XPC_DOMAIN_TYPE_SESSION:
		return "ES_XPC_DOMAIN_TYPE_SESSION"
	case ES_XPC_DOMAIN_TYPE_SYSTEM:
		return "ES_XPC_DOMAIN_TYPE_SYSTEM"
	case ES_XPC_DOMAIN_TYPE_USER:
		return "ES_XPC_DOMAIN_TYPE_USER"
	case ES_XPC_DOMAIN_TYPE_USER_LOGIN:
		return "ES_XPC_DOMAIN_TYPE_USER_LOGIN"
	default:
		return fmt.Sprintf("EsXPCDomainType(%d)", e)
	}
}
