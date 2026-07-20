// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BaseSystemSpeechRenderer] class.
var (
	_BaseSystemSpeechRendererClass     BaseSystemSpeechRendererClass
	_BaseSystemSpeechRendererClassOnce sync.Once
)

func getBaseSystemSpeechRendererClass() BaseSystemSpeechRendererClass {
	_BaseSystemSpeechRendererClassOnce.Do(func() {
		_BaseSystemSpeechRendererClass = BaseSystemSpeechRendererClass{class: objc.GetClass("_TtC12TextToSpeechP33_CFE865E2CA63263E0B31A28183D6B2F124BaseSystemSpeechRenderer")}
	})
	return _BaseSystemSpeechRendererClass
}

// GetBaseSystemSpeechRendererClass returns the class object for _TtC12TextToSpeechP33_CFE865E2CA63263E0B31A28183D6B2F124BaseSystemSpeechRenderer.
func GetBaseSystemSpeechRendererClass() BaseSystemSpeechRendererClass {
	return getBaseSystemSpeechRendererClass()
}

type BaseSystemSpeechRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BaseSystemSpeechRendererClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BaseSystemSpeechRendererClass) Alloc() BaseSystemSpeechRenderer {
	rv := objc.Send[BaseSystemSpeechRenderer](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

type BaseSystemSpeechRenderer struct {
	objectivec.Object
}

// BaseSystemSpeechRendererFromID constructs a [BaseSystemSpeechRenderer] from an objc.ID.
func BaseSystemSpeechRendererFromID(id objc.ID) BaseSystemSpeechRenderer {
	return BaseSystemSpeechRenderer{objectivec.Object{ID: id}}
}

// NOTE: BaseSystemSpeechRenderer struct embeds objectivec.Object (parent type unavailable) but
// IBaseSystemSpeechRenderer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [BaseSystemSpeechRenderer] class.
type IBaseSystemSpeechRenderer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BaseSystemSpeechRenderer) Init() BaseSystemSpeechRenderer {
	rv := objc.Send[BaseSystemSpeechRenderer](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BaseSystemSpeechRenderer) Autorelease() BaseSystemSpeechRenderer {
	rv := objc.Send[BaseSystemSpeechRenderer](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBaseSystemSpeechRenderer creates a new BaseSystemSpeechRenderer instance.
func NewBaseSystemSpeechRenderer() BaseSystemSpeechRenderer {
	class := getBaseSystemSpeechRendererClass()
	rv := objc.Send[BaseSystemSpeechRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
