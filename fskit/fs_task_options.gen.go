// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSTaskOptions] class.
var (
	_FSTaskOptionsClass     FSTaskOptionsClass
	_FSTaskOptionsClassOnce sync.Once
)

func getFSTaskOptionsClass() FSTaskOptionsClass {
	_FSTaskOptionsClassOnce.Do(func() {
		_FSTaskOptionsClass = FSTaskOptionsClass{class: objc.GetClass("FSTaskOptions")}
	})
	return _FSTaskOptionsClass
}

// GetFSTaskOptionsClass returns the class object for FSTaskOptions.
func GetFSTaskOptionsClass() FSTaskOptionsClass {
	return getFSTaskOptionsClass()
}

type FSTaskOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSTaskOptionsClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSTaskOptionsClass) Alloc() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A class that passes command options to a task, optionally providing
// security-scoped URLs.
//
// # Retrieving task options
//
//   - [FSTaskOptions.TaskOptions]: An array of strings that represent command-line options for the task.
//
// # Retrieving task option URLs
//
//   - [FSTaskOptions.UrlForOption]: Retrieves a URL for a given option.
//
// See: https://developer.apple.com/documentation/FSKit/FSTaskOptions
type FSTaskOptions struct {
	objectivec.Object
}

// FSTaskOptionsFromID constructs a [FSTaskOptions] from an objc.ID.
//
// A class that passes command options to a task, optionally providing
// security-scoped URLs.
func FSTaskOptionsFromID(id objc.ID) FSTaskOptions {
	return FSTaskOptions{objectivec.Object{ID: id}}
}

// NOTE: FSTaskOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSTaskOptions] class.
//
// # Retrieving task options
//
//   - [IFSTaskOptions.TaskOptions]: An array of strings that represent command-line options for the task.
//
// # Retrieving task option URLs
//
//   - [IFSTaskOptions.UrlForOption]: Retrieves a URL for a given option.
//
// See: https://developer.apple.com/documentation/FSKit/FSTaskOptions
type IFSTaskOptions interface {
	objectivec.IObject

	// Topic: Retrieving task options

	// An array of strings that represent command-line options for the task.
	TaskOptions() []string

	// Topic: Retrieving task option URLs

	// Retrieves a URL for a given option.
	UrlForOption(option string) foundation.NSURL
}

// Init initializes the instance.
func (t FSTaskOptions) Init() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t FSTaskOptions) Autorelease() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTaskOptions creates a new FSTaskOptions instance.
func NewFSTaskOptions() FSTaskOptions {
	class := getFSTaskOptionsClass()
	rv := objc.Send[FSTaskOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves a URL for a given option.
//
// option: The option for which to retrieve the URL. This value doesn’t include
// leading dashes.
//
// # Discussion
//
// Some command-line options refer to paths that indicate a location in which
// the module needs access to a file outside of its container. FSKit passes
// these paths as a URL tagged by the option name.
//
// For example, `"-B" "./someFile"` returns the URL for `./someFile` when
// passed an option `"B"`. To indicate that your module treats a given option
// as a path, include it in the `pathOptions` dictionary within a command
// options dictionary ([FSActivatOptionSyntax], [FSCheckOptionSyntax], or
// [FSFormatOptionSyntax]). This dictionary uses the command option name as a
// key, and each entry has a value indicating what kind of entry to create.
//
// See: https://developer.apple.com/documentation/FSKit/FSTaskOptions/url(forOption:)
func (t FSTaskOptions) UrlForOption(option string) foundation.NSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("urlForOption:"), objc.String(option))
	return foundation.NSURLFromID(rv)
}

// An array of strings that represent command-line options for the task.
//
// # Discussion
//
// This property is equivalent to the `argv` array of C strings passed to a
// command-line tool.
//
// See: https://developer.apple.com/documentation/FSKit/FSTaskOptions/taskOptions
func (t FSTaskOptions) TaskOptions() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("taskOptions"))
	return objc.ConvertSliceToStrings(rv)
}
