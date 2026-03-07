package vzkit

import (
	"fmt"
	"strings"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ExtractNSErrorMessage extracts the localized error message from an NSError.
// Returns an empty string if the error ID is zero.
func ExtractNSErrorMessage(nsError objc.ID) string {
	if nsError == 0 {
		return ""
	}
	nsErr := foundation.NSErrorFromID(nsError)
	if desc := nsErr.LocalizedDescription(); desc != "" {
		return desc
	}
	domain := nsErr.Domain()
	code := nsErr.Code()
	if domain != "" {
		return fmt.Sprintf("%s (code %d)", domain, code)
	}
	return fmt.Sprintf("error code %d", code)
}

// FormatNSErrorDetailed returns a multi-line string with full NSError details
// including domain, code, failure reason, recovery suggestion, user info, and
// underlying errors.
func FormatNSErrorDetailed(nsError objc.ID) string {
	if nsError == 0 {
		return ""
	}
	nsErr := foundation.NSErrorFromID(nsError)
	var b strings.Builder

	b.WriteString("=== Detailed NSError Analysis ===\n")
	fmt.Fprintf(&b, "Domain: %s\n", nsErr.Domain())
	fmt.Fprintf(&b, "Code: %d\n", nsErr.Code())

	if desc := nsErr.LocalizedDescription(); desc != "" {
		fmt.Fprintf(&b, "LocalizedDescription: %s\n", desc)
	}
	if reason := nsErr.LocalizedFailureReason(); reason != "" {
		fmt.Fprintf(&b, "LocalizedFailureReason: %s\n", reason)
	}
	if suggestion := nsErr.LocalizedRecoverySuggestion(); suggestion != "" {
		fmt.Fprintf(&b, "LocalizedRecoverySuggestion: %s\n", suggestion)
	}

	if userInfo := nsErr.UserInfo(); userInfo != nil && userInfo.GetID() != 0 {
		keys := userInfo.AllKeys()
		fmt.Fprintf(&b, "\nUserInfo contents (%d keys):\n", len(keys))
		for _, key := range keys {
			value := userInfo.ObjectForKey(key)
			if value == nil {
				continue
			}
			valueID := value.GetID()
			valueObj := objectivec.NSObjectObjectFromID(valueID)
			isNSError := valueObj.IsKindOfClass(objc.GetClass("NSError"))
			if isNSError {
				underErr := foundation.NSErrorFromID(valueID)
				keyObj := objectivec.ObjectFromID(key.GetID())
				fmt.Fprintf(&b, "  %s: [NSError] %s (code %d) %s\n",
					keyObj.Description(), underErr.Domain(), underErr.Code(), underErr.LocalizedDescription())
			} else {
				keyObj := objectivec.ObjectFromID(key.GetID())
				valObj := objectivec.ObjectFromID(valueID)
				fmt.Fprintf(&b, "  %s: %s\n", keyObj.Description(), valObj.Description())
			}
		}
	}

	if underlying := nsErr.UnderlyingErrors(); len(underlying) > 0 {
		fmt.Fprintf(&b, "\nUnderlying errors (%d):\n", len(underlying))
		for i, ue := range underlying {
			fmt.Fprintf(&b, "  [%d] %s (code %d): %s\n", i, ue.Domain(), ue.Code(), ue.LocalizedDescription())
		}
	}

	b.WriteString("=== End NSError Analysis ===")
	return b.String()
}

// PrintNSErrorDetailed prints detailed NSError information to stdout.
func PrintNSErrorDetailed(nsError objc.ID) {
	if s := FormatNSErrorDetailed(nsError); s != "" {
		fmt.Println(s)
	}
}
