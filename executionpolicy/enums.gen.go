// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus
type EPDeveloperToolStatus int

const (
	EPDeveloperToolStatusAuthorized    EPDeveloperToolStatus = 3
	EPDeveloperToolStatusDenied        EPDeveloperToolStatus = 2
	EPDeveloperToolStatusNotDetermined EPDeveloperToolStatus = 0
	EPDeveloperToolStatusRestricted    EPDeveloperToolStatus = 1
)

func (e EPDeveloperToolStatus) String() string {
	switch e {
	case EPDeveloperToolStatusAuthorized:
		return "EPDeveloperToolStatusAuthorized"
	case EPDeveloperToolStatusDenied:
		return "EPDeveloperToolStatusDenied"
	case EPDeveloperToolStatusNotDetermined:
		return "EPDeveloperToolStatusNotDetermined"
	case EPDeveloperToolStatusRestricted:
		return "EPDeveloperToolStatusRestricted"
	default:
		return fmt.Sprintf("EPDeveloperToolStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ExecutionPolicy/EPError-swift.struct/Code
type EPError int

const (
	EPErrorGeneric           EPError = 1
	EPErrorNotADeveloperTool EPError = 2
)

func (e EPError) String() string {
	switch e {
	case EPErrorGeneric:
		return "EPErrorGeneric"
	case EPErrorNotADeveloperTool:
		return "EPErrorNotADeveloperTool"
	default:
		return fmt.Sprintf("EPError(%d)", e)
	}
}
