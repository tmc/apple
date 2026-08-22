// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeechThread] class.
var (
	_TTSSpeechThreadClass     TTSSpeechThreadClass
	_TTSSpeechThreadClassOnce sync.Once
)

func getTTSSpeechThreadClass() TTSSpeechThreadClass {
	_TTSSpeechThreadClassOnce.Do(func() {
		_TTSSpeechThreadClass = TTSSpeechThreadClass{class: objc.GetClass("TTSSpeechThread")}
	})
	return _TTSSpeechThreadClass
}

// GetTTSSpeechThreadClass returns the class object for TTSSpeechThread.
func GetTTSSpeechThreadClass() TTSSpeechThreadClass {
	return getTTSSpeechThreadClass()
}

type TTSSpeechThreadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechThreadClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechThreadClass) Alloc() TTSSpeechThread {
	rv := objc.SendIfResponds[TTSSpeechThread](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSSpeechThread.Stop]
//   - [TTSSpeechThread.Voucher]
//   - [TTSSpeechThread.SetVoucher]
type TTSSpeechThread struct {
	foundation.NSThread
}

// TTSSpeechThreadFromID constructs a [TTSSpeechThread] from an objc.ID.
func TTSSpeechThreadFromID(id objc.ID) TTSSpeechThread {
	return TTSSpeechThread{NSThread: foundation.NSThreadFromID(id)}
}

// Ensure TTSSpeechThread implements ITTSSpeechThread.
var _ ITTSSpeechThread = TTSSpeechThread{}

// An interface definition for the [TTSSpeechThread] class.
//
// # Methods
//
//   - [ITTSSpeechThread.Stop]
//   - [ITTSSpeechThread.Voucher]
//   - [ITTSSpeechThread.SetVoucher]
type ITTSSpeechThread interface {
	foundation.IThread

	// Topic: Methods

	Stop()
	Voucher() objectivec.Object
	SetVoucher(value objectivec.Object)
}

// Init initializes the instance.
func (t TTSSpeechThread) Init() TTSSpeechThread {
	rv := objc.SendIfResponds[TTSSpeechThread](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechThread) Autorelease() TTSSpeechThread {
	rv := objc.SendIfResponds[TTSSpeechThread](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechThread creates a new TTSSpeechThread instance.
func NewTTSSpeechThread() TTSSpeechThread {
	class := getTTSSpeechThreadClass()
	rv := objc.SendIfResponds[TTSSpeechThread](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSSpeechThread) Cancel() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("cancel"))
}
func (t TTSSpeechThread) Main() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("main"))
}
func (t TTSSpeechThread) Stop() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("stop"))
}

func (t TTSSpeechThread) Voucher() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voucher"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechThread) SetVoucher(value objectivec.Object) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoucher:"), value)
}
