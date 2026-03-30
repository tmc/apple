// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/objc"
)

// Initializes an index path with the indexes of a specific row and section in
// a table view.
//
// row: An index number identifying a row in a [UITableView] object in a section
// identified by `section`.
//
// section: An index number identifying a section in a [UITableView] object.
//
// # Return Value
//
// An [NSIndexPath] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/init(forRow:inSection:)
//
// [UITableView]: https://developer.apple.com/documentation/UIKit/UITableView
//
// [UITableView]: https://developer.apple.com/documentation/UIKit/UITableView
func (_NSIndexPathClass NSIndexPathClass) IndexPathForRowInSection(row int, section int) NSIndexPath {
	rv := objc.Send[objc.ID](objc.ID(_NSIndexPathClass.class), objc.Sel("indexPathForRow:inSection:"), row, section)
	return NSIndexPathFromID(rv)
}

// An index number identifying a row in a section of a table view.
//
// # Discussion
//
// The section the row is in is identified by the value of [Section].
//
// See: https://developer.apple.com/documentation/Foundation/NSIndexPath/row
func (i NSIndexPath) Row() int {
	rv := objc.Send[int](i.ID, objc.Sel("row"))
	return rv
}
