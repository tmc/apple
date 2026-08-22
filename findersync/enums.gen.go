// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/FinderSync/FIMenuKind
type FIMenuKind uint

const (
	// FIMenuKindContextualMenuForContainer: A shortcut menu created when the user control-clicks on the Finder window’s background.
	FIMenuKindContextualMenuForContainer FIMenuKind = 1
	// FIMenuKindContextualMenuForItems: A shortcut menu created when the user control-clicks on an item or a group of selected items inside the Finder window.
	FIMenuKindContextualMenuForItems FIMenuKind = 0
	// FIMenuKindContextualMenuForSidebar: A shortcut menu created when the user control-clicks on an item in the sidebar.
	FIMenuKindContextualMenuForSidebar FIMenuKind = 2
	// FIMenuKindToolbarItemMenu: A menu created when the user clicks on the extension’s toolbar button.
	FIMenuKindToolbarItemMenu FIMenuKind = 3
)

func (e FIMenuKind) String() string {
	switch e {
	case FIMenuKindContextualMenuForContainer:
		return "FIMenuKindContextualMenuForContainer"
	case FIMenuKindContextualMenuForItems:
		return "FIMenuKindContextualMenuForItems"
	case FIMenuKindContextualMenuForSidebar:
		return "FIMenuKindContextualMenuForSidebar"
	case FIMenuKindToolbarItemMenu:
		return "FIMenuKindToolbarItemMenu"
	default:
		return fmt.Sprintf("FIMenuKind(%d)", e)
	}
}
