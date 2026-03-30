// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLog] class.
var (
	_MLLogClass     MLLogClass
	_MLLogClassOnce sync.Once
)

func getMLLogClass() MLLogClass {
	_MLLogClassOnce.Do(func() {
		_MLLogClass = MLLogClass{class: objc.GetClass("_MLLog")}
	})
	return _MLLogClass
}

// GetMLLogClass returns the class object for _MLLog.
func GetMLLogClass() MLLogClass {
	return getMLLogClass()
}

type MLLogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLogClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLogClass) Alloc() MLLog {
	rv := objc.Send[MLLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog
type MLLog struct {
	objectivec.Object
}

// MLLogFromID constructs a [MLLog] from an objc.ID.
func MLLogFromID(id objc.ID) MLLog {
	return MLLog{objectivec.Object{ID: id}}
}

// Ensure MLLog implements IMLLog.
var _ IMLLog = MLLog{}

// An interface definition for the [MLLog] class.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog
type IMLLog interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLLog) Init() MLLog {
	rv := objc.Send[MLLog](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLLog) Autorelease() MLLog {
	rv := objc.Send[MLLog](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLog creates a new MLLog instance.
func NewMLLog() MLLog {
	class := getMLLogClass()
	rv := objc.Send[MLLog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog/clientFramework
func (_MLLogClass MLLogClass) ClientFramework() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLogClass.class), objc.Sel("clientFramework"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog/common
func (_MLLogClass MLLogClass) Common() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLogClass.class), objc.Sel("common"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog/daemon
func (_MLLogClass MLLogClass) Daemon() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLogClass.class), objc.Sel("daemon"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog/serverFramework
func (_MLLogClass MLLogClass) ServerFramework() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLogClass.class), objc.Sel("serverFramework"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLLog/tool
func (_MLLogClass MLLogClass) Tool() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLogClass.class), objc.Sel("tool"))
	return objectivec.Object{ID: rv}
}
