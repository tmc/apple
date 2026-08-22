// Code generated from Apple documentation. DO NOT EDIT.

package corespotlight

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// CSIndexErrorDomain is the error domain for the index.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSIndexErrorDomain
	CSIndexErrorDomain string
	// CSMailboxArchive is an archive mailbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxArchive
	CSMailboxArchive string
	// CSMailboxDrafts is a drafts mailbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxDrafts
	CSMailboxDrafts string
	// CSMailboxInbox is an inbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxInbox
	CSMailboxInbox string
	// CSMailboxJunk is a junk mailbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxJunk
	CSMailboxJunk string
	// CSMailboxSent is a sent mailbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxSent
	CSMailboxSent string
	// CSMailboxTrash is a trash mailbox.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSMailboxTrash
	CSMailboxTrash string
	// CSQueryContinuationActionType is indicates that the activity type to continue is a search or query.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSQueryContinuationActionType
	CSQueryContinuationActionType string
	// CSSearchQueryString is provides the key for the current query in the info dictionary of the user activity object.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryString
	CSSearchQueryString string
	// CSSearchableItemActionType is indicates that the activity type to continue is related to a searchable item.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemActionType
	CSSearchableItemActionType string
	// CSSearchableItemActivityIdentifier is the key you use to access a searchable item in a user activity object.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemActivityIdentifier
	CSSearchableItemActivityIdentifier string
)

var (
	// CSSearchQueryErrorDomain is the error domain for search queries.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryErrorDomain
	CSSearchQueryErrorDomain foundation.NSErrorDomain
)

var (
	// See: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestionHighlightAttributeName
	CSSuggestionHighlightAttributeName foundation.NSAttributedStringKey
)

var (
	// CoreSpotlightVersionNumber is the project version number for Core Spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CoreSpotlightVersionNumber
	CoreSpotlightVersionNumber float64
)

var (
	// CoreSpotlightVersionString is the project version string for Core Spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreSpotlight/CoreSpotlightVersionString
	CoreSpotlightVersionString uint8
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSIndexErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSIndexErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxArchive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxArchive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxDrafts"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxDrafts = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxInbox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxInbox = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxJunk"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxJunk = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxSent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxSent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSMailboxTrash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSMailboxTrash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSQueryContinuationActionType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSQueryContinuationActionType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSSearchQueryErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSSearchQueryErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSSearchQueryString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSSearchQueryString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSSearchableItemActionType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSSearchableItemActionType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSSearchableItemActivityIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSSearchableItemActivityIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CSSuggestionHighlightAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CSSuggestionHighlightAttributeName = foundation.NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CoreSpotlightVersionNumber"); err == nil && ptr != 0 {
		CoreSpotlightVersionNumber = objc.ValueAt[float64](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CoreSpotlightVersionString"); err == nil && ptr != 0 {
		CoreSpotlightVersionString = objc.ValueAt[uint8](ptr)
	}

}
