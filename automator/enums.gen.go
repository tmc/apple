// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Automator/AMError/Code
type AMErrorCode int

const (
	// AMActionApplicationResourceError: An error that indicates an app required by the action is not found.
	AMActionApplicationResourceError AMErrorCode = -206
	// AMActionApplicationVersionResourceError: An error that indicates an app required by the action is the wrong version.
	AMActionApplicationVersionResourceError AMErrorCode = -207
	// AMActionArchitectureMismatchError: An error that indicates the action’s binary is not compatible with the current processor.
	AMActionArchitectureMismatchError AMErrorCode = -202
	// AMActionExceptionError: An error that indicates an action encounters an exception while running.
	AMActionExceptionError AMErrorCode = -213
	// AMActionExecutionError: An error that indicates an action encounters an error while running (reason unknown).
	AMActionExecutionError AMErrorCode = -212
	// AMActionFailedGatekeeperError: An error that indicates the action doesn’t meet the Gatekeeper security policy.
	AMActionFailedGatekeeperError AMErrorCode = -217
	// AMActionFileResourceError: An error that indicates a file required by the action is not found.
	AMActionFileResourceError AMErrorCode = -208
	// AMActionInitializationError: An error that indicates Automator is unable to initialize an action (reason unknown).
	AMActionInitializationError AMErrorCode = -211
	// AMActionInsufficientDataError: An error that indicates the action requires input data to run, but none was supplied.
	AMActionInsufficientDataError AMErrorCode = -215
	// AMActionIsDeprecatedError: An error that indicates the action has been deprecated.
	AMActionIsDeprecatedError AMErrorCode = -216
	// AMActionLicenseResourceError: An error that indicates a license required by the action was not found.
	AMActionLicenseResourceError AMErrorCode = -209
	// AMActionLinkError: An error that indicates the action’s executable failed to load due to linking issues.
	AMActionLinkError AMErrorCode = -205
	// AMActionLoadError: An error that indicates the action’s executable failed to load.
	AMActionLoadError AMErrorCode = -204
	// AMActionMalwareError: An error that indicates the action has been identified as malware by XProtect.
	AMActionMalwareError AMErrorCode = -221
	// AMActionNotLoadableError: An error that indicates the action’s executable is of a type that is not loadable in the current process.
	AMActionNotLoadableError AMErrorCode = -201
	// AMActionPropertyListInvalidError: An error that indicates the property list for an action is invalid.
	AMActionPropertyListInvalidError AMErrorCode = -214
	// AMActionQuarantineError: An error that indicates action has been quarantined by XProtect, the antimalware system on the Mac.
	AMActionQuarantineError AMErrorCode = -219
	// AMActionRequiredActionResourceError: An error that indicates an action required by the action is not loaded.
	AMActionRequiredActionResourceError AMErrorCode = -210
	// AMActionRuntimeMismatchError: An error that indicates an attempt was made to load an action that is not compiled in a way that is compatible with the current app.
	AMActionRuntimeMismatchError AMErrorCode = -203
	// AMActionSignatureCorruptError: An error that indicates developer signature for this action is corrupted.
	AMActionSignatureCorruptError AMErrorCode = -218
	// AMActionThirdPartyActionsNotAllowedError: An error that indicates the action is a third party action, and loading it has not been allowed by the user.
	AMActionThirdPartyActionsNotAllowedError AMErrorCode = -222
	// AMActionXPCError: An error that indicates the remote process running the action has crashed.
	AMActionXPCError AMErrorCode = -223
	// AMActionXProtectError: An error that indicates XProtect is unable to successfully analyze the action.
	AMActionXProtectError AMErrorCode = -220
	// AMConversionFailedError: An error that occurs when, for example, the converter encounters an error converting data from one type to another.
	AMConversionFailedError AMErrorCode = -302
	// AMConversionNoDataError: An error that occurs when the converter determines that the conversion, though possible, would produce a nil result.
	AMConversionNoDataError AMErrorCode = -301
	// AMConversionNotPossibleError: An error that occurs when the converter determines that it is unable to convert from one data type to another.
	AMConversionNotPossibleError AMErrorCode = -300
	// AMNoSuchActionError: An error that indicates the action could not be located on the system.
	AMNoSuchActionError AMErrorCode = -200
	// AMUserCanceledError: An error that indicates the user cancelled.
	AMUserCanceledError AMErrorCode = -128
	// AMWorkflowActionsNotLoadedError: An error that indicates one of the actions of the workflow couldn’t be loaded.
	AMWorkflowActionsNotLoadedError AMErrorCode = -113
	// AMWorkflowNewerActionVersionError: An error that indicates an action in a workflow is newer than the installed action.
	AMWorkflowNewerActionVersionError AMErrorCode = -111
	// AMWorkflowNewerVersionError: An error that indicates an attempt to open a workflow document that was saved with a newer version of Automator.
	AMWorkflowNewerVersionError AMErrorCode = -100
	// AMWorkflowNoEnabledActionsError: An error that indicates there are no enabled actions in the workflow.
	AMWorkflowNoEnabledActionsError AMErrorCode = -114
	// AMWorkflowOlderActionVersionError: An error that indicates an action in a workflow is older than the installed action.
	AMWorkflowOlderActionVersionError AMErrorCode = -112
	// AMWorkflowPropertyListInvalidError: An error that indicates an attempt to open a workflow document whose property list couldn’t be read.
	AMWorkflowPropertyListInvalidError AMErrorCode = -101
)

func (e AMErrorCode) String() string {
	switch e {
	case AMActionApplicationResourceError:
		return "AMActionApplicationResourceError"
	case AMActionApplicationVersionResourceError:
		return "AMActionApplicationVersionResourceError"
	case AMActionArchitectureMismatchError:
		return "AMActionArchitectureMismatchError"
	case AMActionExceptionError:
		return "AMActionExceptionError"
	case AMActionExecutionError:
		return "AMActionExecutionError"
	case AMActionFailedGatekeeperError:
		return "AMActionFailedGatekeeperError"
	case AMActionFileResourceError:
		return "AMActionFileResourceError"
	case AMActionInitializationError:
		return "AMActionInitializationError"
	case AMActionInsufficientDataError:
		return "AMActionInsufficientDataError"
	case AMActionIsDeprecatedError:
		return "AMActionIsDeprecatedError"
	case AMActionLicenseResourceError:
		return "AMActionLicenseResourceError"
	case AMActionLinkError:
		return "AMActionLinkError"
	case AMActionLoadError:
		return "AMActionLoadError"
	case AMActionMalwareError:
		return "AMActionMalwareError"
	case AMActionNotLoadableError:
		return "AMActionNotLoadableError"
	case AMActionPropertyListInvalidError:
		return "AMActionPropertyListInvalidError"
	case AMActionQuarantineError:
		return "AMActionQuarantineError"
	case AMActionRequiredActionResourceError:
		return "AMActionRequiredActionResourceError"
	case AMActionRuntimeMismatchError:
		return "AMActionRuntimeMismatchError"
	case AMActionSignatureCorruptError:
		return "AMActionSignatureCorruptError"
	case AMActionThirdPartyActionsNotAllowedError:
		return "AMActionThirdPartyActionsNotAllowedError"
	case AMActionXPCError:
		return "AMActionXPCError"
	case AMActionXProtectError:
		return "AMActionXProtectError"
	case AMConversionFailedError:
		return "AMConversionFailedError"
	case AMConversionNoDataError:
		return "AMConversionNoDataError"
	case AMConversionNotPossibleError:
		return "AMConversionNotPossibleError"
	case AMNoSuchActionError:
		return "AMNoSuchActionError"
	case AMUserCanceledError:
		return "AMUserCanceledError"
	case AMWorkflowActionsNotLoadedError:
		return "AMWorkflowActionsNotLoadedError"
	case AMWorkflowNewerActionVersionError:
		return "AMWorkflowNewerActionVersionError"
	case AMWorkflowNewerVersionError:
		return "AMWorkflowNewerVersionError"
	case AMWorkflowNoEnabledActionsError:
		return "AMWorkflowNoEnabledActionsError"
	case AMWorkflowOlderActionVersionError:
		return "AMWorkflowOlderActionVersionError"
	case AMWorkflowPropertyListInvalidError:
		return "AMWorkflowPropertyListInvalidError"
	default:
		return fmt.Sprintf("AMErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Automator/AMLogLevel
type AMLogLevel uint

const (
	// AMLogLevelDebug: The debug log level.
	AMLogLevelDebug AMLogLevel = 0
	// AMLogLevelError: The error log level.
	AMLogLevelError AMLogLevel = 3
	// AMLogLevelInfo: The informational log level.
	AMLogLevelInfo AMLogLevel = 1
	// AMLogLevelWarn: The warning log level.
	AMLogLevelWarn AMLogLevel = 2
)

func (e AMLogLevel) String() string {
	switch e {
	case AMLogLevelDebug:
		return "AMLogLevelDebug"
	case AMLogLevelError:
		return "AMLogLevelError"
	case AMLogLevelInfo:
		return "AMLogLevelInfo"
	case AMLogLevelWarn:
		return "AMLogLevelWarn"
	default:
		return fmt.Sprintf("AMLogLevel(%d)", e)
	}
}
