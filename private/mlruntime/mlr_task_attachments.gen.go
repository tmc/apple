// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTaskAttachments] class.
var (
	_MLRTaskAttachmentsClass     MLRTaskAttachmentsClass
	_MLRTaskAttachmentsClassOnce sync.Once
)

func getMLRTaskAttachmentsClass() MLRTaskAttachmentsClass {
	_MLRTaskAttachmentsClassOnce.Do(func() {
		_MLRTaskAttachmentsClass = MLRTaskAttachmentsClass{class: objc.GetClass("MLRTaskAttachments")}
	})
	return _MLRTaskAttachmentsClass
}

// GetMLRTaskAttachmentsClass returns the class object for MLRTaskAttachments.
func GetMLRTaskAttachmentsClass() MLRTaskAttachmentsClass {
	return getMLRTaskAttachmentsClass()
}

type MLRTaskAttachmentsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTaskAttachmentsClass) Alloc() MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTaskAttachments.AttachmentURLs]
//   - [MLRTaskAttachments.AttachmentURLsForBasename]
//   - [MLRTaskAttachments.Count]
//   - [MLRTaskAttachments.EncodeWithCoder]
//   - [MLRTaskAttachments.InitWithCoder]
//   - [MLRTaskAttachments.InitWithContentsOfURLError]
//   - [MLRTaskAttachments.InitWithDESRecipe]
//   - [MLRTaskAttachments.InitWithURLs]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments
type MLRTaskAttachments struct {
	objectivec.Object
}

// MLRTaskAttachmentsFromID constructs a [MLRTaskAttachments] from an objc.ID.
func MLRTaskAttachmentsFromID(id objc.ID) MLRTaskAttachments {
	return MLRTaskAttachments{objectivec.Object{ID: id}}
}
// Ensure MLRTaskAttachments implements IMLRTaskAttachments.
var _ IMLRTaskAttachments = MLRTaskAttachments{}

// An interface definition for the [MLRTaskAttachments] class.
//
// # Methods
//
//   - [IMLRTaskAttachments.AttachmentURLs]
//   - [IMLRTaskAttachments.AttachmentURLsForBasename]
//   - [IMLRTaskAttachments.Count]
//   - [IMLRTaskAttachments.EncodeWithCoder]
//   - [IMLRTaskAttachments.InitWithCoder]
//   - [IMLRTaskAttachments.InitWithContentsOfURLError]
//   - [IMLRTaskAttachments.InitWithDESRecipe]
//   - [IMLRTaskAttachments.InitWithURLs]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments
type IMLRTaskAttachments interface {
	objectivec.IObject

	// Topic: Methods

	AttachmentURLs() foundation.INSArray
	AttachmentURLsForBasename(basename objectivec.IObject) objectivec.IObject
	Count() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	InitWithCoder(coder foundation.INSCoder) MLRTaskAttachments
	InitWithContentsOfURLError(url foundation.INSURL) (MLRTaskAttachments, error)
	InitWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskAttachments
	InitWithURLs(uRLs objectivec.IObject) MLRTaskAttachments
}

// Init initializes the instance.
func (r MLRTaskAttachments) Init() MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTaskAttachments) Autorelease() MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTaskAttachments creates a new MLRTaskAttachments instance.
func NewMLRTaskAttachments() MLRTaskAttachments {
	class := getMLRTaskAttachmentsClass()
	rv := objc.Send[MLRTaskAttachments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithCoder:
func NewRTaskAttachmentsWithCoder(coder objectivec.IObject) MLRTaskAttachments {
	instance := getMLRTaskAttachmentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLRTaskAttachmentsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithContentsOfURL:error:
func NewRTaskAttachmentsWithContentsOfURLError(url foundation.INSURL) (MLRTaskAttachments, error) {
	var errorPtr objc.ID
	instance := getMLRTaskAttachmentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTaskAttachmentsFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLRTaskAttachmentsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithDESRecipe:
func NewRTaskAttachmentsWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskAttachments {
	instance := getMLRTaskAttachmentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDESRecipe:"), dESRecipe)
	return MLRTaskAttachmentsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithURLs:
func NewRTaskAttachmentsWithURLs(uRLs objectivec.IObject) MLRTaskAttachments {
	instance := getMLRTaskAttachmentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURLs:"), uRLs)
	return MLRTaskAttachmentsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/attachmentURLsForBasename:
func (r MLRTaskAttachments) AttachmentURLsForBasename(basename objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("attachmentURLsForBasename:"), basename)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/encodeWithCoder:
func (r MLRTaskAttachments) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithCoder:
func (r MLRTaskAttachments) InitWithCoder(coder foundation.INSCoder) MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithContentsOfURL:error:
func (r MLRTaskAttachments) InitWithContentsOfURLError(url foundation.INSURL) (MLRTaskAttachments, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTaskAttachments{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLRTaskAttachmentsFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithDESRecipe:
func (r MLRTaskAttachments) InitWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](r.ID, objc.Sel("initWithDESRecipe:"), dESRecipe)
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/initWithURLs:
func (r MLRTaskAttachments) InitWithURLs(uRLs objectivec.IObject) MLRTaskAttachments {
	rv := objc.Send[MLRTaskAttachments](r.ID, objc.Sel("initWithURLs:"), uRLs)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/supportsSecureCoding
func (_MLRTaskAttachmentsClass MLRTaskAttachmentsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLRTaskAttachmentsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/attachmentURLs
func (r MLRTaskAttachments) AttachmentURLs() foundation.INSArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("attachmentURLs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskAttachments/count
func (r MLRTaskAttachments) Count() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("count"))
	return rv
}

