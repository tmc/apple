// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
)

// A hint to the system of the relative order for purging tagged sets of
// resources in the bundle.
//
// priority: A number specifying the relative priority of preserving the resources in
// the group specified by `tag`.
// 
// Possible values are between `0.0` and `1.0`. The default is `0.0`. The
// system will attempt to purge resources with lower priorities first.
//
// tags: A set of tag names specifying resources stored in the bundle. Must not be
// `nil`. An exception is thrown if any of the tags in the set do not exist in
// your app.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/setPreservationPriority(_:forTags:)
func (b Bundle) SetPreservationPriorityForTags(priority float64, tags INSSet) {
objc.Send[objc.ID](b.ID, objc.Sel("setPreservationPriority:forTags:"), priority, tags)
}

// Returns the current preservation priority for the specified tag.
//
// tag: A string specifying the identifier for a group of related resources. An
// exception is thrown if `tag` does not exist in your app.
//
// # Return Value
// 
// The preservation priority for the specified `tag`. Possible values are
// between `0.0` and `1.0`
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/preservationPriority(forTag:)
func (b Bundle) PreservationPriorityForTag(tag string) float64 {
rv := objc.Send[float64](b.ID, objc.Sel("preservationPriorityForTag:"), objc.String(tag))
return rv
}

