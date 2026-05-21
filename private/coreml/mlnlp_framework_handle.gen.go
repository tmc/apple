// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNLPFrameworkHandle] class.
var (
	_MLNLPFrameworkHandleClass     MLNLPFrameworkHandleClass
	_MLNLPFrameworkHandleClassOnce sync.Once
)

func getMLNLPFrameworkHandleClass() MLNLPFrameworkHandleClass {
	_MLNLPFrameworkHandleClassOnce.Do(func() {
		_MLNLPFrameworkHandleClass = MLNLPFrameworkHandleClass{class: objc.GetClass("_MLNLPFrameworkHandle")}
	})
	return _MLNLPFrameworkHandleClass
}

// GetMLNLPFrameworkHandleClass returns the class object for _MLNLPFrameworkHandle.
func GetMLNLPFrameworkHandleClass() MLNLPFrameworkHandleClass {
	return getMLNLPFrameworkHandleClass()
}

type MLNLPFrameworkHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNLPFrameworkHandleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNLPFrameworkHandleClass) Alloc() MLNLPFrameworkHandle {
	rv := objc.Send[MLNLPFrameworkHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNLPFrameworkHandle.NLPClassifierModelCopyPredictedLabelForTextImpl]
//   - [MLNLPFrameworkHandle.NLPClassifierModelCreateWithDataImpl]
//   - [MLNLPFrameworkHandle.NLPClassifierModelGetCurrentRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPClassifierModelGetRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPClassifierModelIsRevisionSupportedImpl]
//   - [MLNLPFrameworkHandle.NLPEmbeddingModelCopyVectorForStringImpl]
//   - [MLNLPFrameworkHandle.NLPEmbeddingModelCreateWithDataImpl]
//   - [MLNLPFrameworkHandle.NLPEmbeddingModelGetCurrentRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPEmbeddingModelGetRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPEmbeddingModelIsRevisionSupportedImpl]
//   - [MLNLPFrameworkHandle.NLPGazetteerModelCopyLabelForStringImpl]
//   - [MLNLPFrameworkHandle.NLPGazetteerModelCreateWithDataImpl]
//   - [MLNLPFrameworkHandle.NLPGazetteerModelGetCurrentRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPGazetteerModelGetRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPGazetteerModelIsRevisionSupportedImpl]
//   - [MLNLPFrameworkHandle.NLPSequenceModelCopyPredictedTokensAndLabelsForTextImpl]
//   - [MLNLPFrameworkHandle.NLPSequenceModelCreateWithDataImpl]
//   - [MLNLPFrameworkHandle.NLPSequenceModelGetCurrentRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPSequenceModelGetRevisionImpl]
//   - [MLNLPFrameworkHandle.NLPSequenceModelIsRevisionSupportedImpl]
//   - [MLNLPFrameworkHandle.InitializeEmbeddingModelWithDataError]
//   - [MLNLPFrameworkHandle.InitializeGazetteerModelWithDataError]
//   - [MLNLPFrameworkHandle.InitializeSentenceClassifierModelWithDataError]
//   - [MLNLPFrameworkHandle.InitializeWordTaggingModelWithDataError]
//   - [MLNLPFrameworkHandle.IsValid]
//   - [MLNLPFrameworkHandle.PredictLabelForWordStringInputStringError]
//   - [MLNLPFrameworkHandle.PredictLabelsForSentenceStringInputStringError]
//   - [MLNLPFrameworkHandle.PredictTokensLabelsLocationsLengthsForStringInputStringError]
//   - [MLNLPFrameworkHandle.PredictVectorForStringInputStringError]
//   - [MLNLPFrameworkHandle.Valid]
type MLNLPFrameworkHandle struct {
	objectivec.Object
}

// MLNLPFrameworkHandleFromID constructs a [MLNLPFrameworkHandle] from an objc.ID.
func MLNLPFrameworkHandleFromID(id objc.ID) MLNLPFrameworkHandle {
	return MLNLPFrameworkHandle{objectivec.Object{ID: id}}
}

// Ensure MLNLPFrameworkHandle implements IMLNLPFrameworkHandle.
var _ IMLNLPFrameworkHandle = MLNLPFrameworkHandle{}

// An interface definition for the [MLNLPFrameworkHandle] class.
//
// # Methods
//
//   - [IMLNLPFrameworkHandle.NLPClassifierModelCopyPredictedLabelForTextImpl]
//   - [IMLNLPFrameworkHandle.NLPClassifierModelCreateWithDataImpl]
//   - [IMLNLPFrameworkHandle.NLPClassifierModelGetCurrentRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPClassifierModelGetRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPClassifierModelIsRevisionSupportedImpl]
//   - [IMLNLPFrameworkHandle.NLPEmbeddingModelCopyVectorForStringImpl]
//   - [IMLNLPFrameworkHandle.NLPEmbeddingModelCreateWithDataImpl]
//   - [IMLNLPFrameworkHandle.NLPEmbeddingModelGetCurrentRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPEmbeddingModelGetRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPEmbeddingModelIsRevisionSupportedImpl]
//   - [IMLNLPFrameworkHandle.NLPGazetteerModelCopyLabelForStringImpl]
//   - [IMLNLPFrameworkHandle.NLPGazetteerModelCreateWithDataImpl]
//   - [IMLNLPFrameworkHandle.NLPGazetteerModelGetCurrentRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPGazetteerModelGetRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPGazetteerModelIsRevisionSupportedImpl]
//   - [IMLNLPFrameworkHandle.NLPSequenceModelCopyPredictedTokensAndLabelsForTextImpl]
//   - [IMLNLPFrameworkHandle.NLPSequenceModelCreateWithDataImpl]
//   - [IMLNLPFrameworkHandle.NLPSequenceModelGetCurrentRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPSequenceModelGetRevisionImpl]
//   - [IMLNLPFrameworkHandle.NLPSequenceModelIsRevisionSupportedImpl]
//   - [IMLNLPFrameworkHandle.InitializeEmbeddingModelWithDataError]
//   - [IMLNLPFrameworkHandle.InitializeGazetteerModelWithDataError]
//   - [IMLNLPFrameworkHandle.InitializeSentenceClassifierModelWithDataError]
//   - [IMLNLPFrameworkHandle.InitializeWordTaggingModelWithDataError]
//   - [IMLNLPFrameworkHandle.IsValid]
//   - [IMLNLPFrameworkHandle.PredictLabelForWordStringInputStringError]
//   - [IMLNLPFrameworkHandle.PredictLabelsForSentenceStringInputStringError]
//   - [IMLNLPFrameworkHandle.PredictTokensLabelsLocationsLengthsForStringInputStringError]
//   - [IMLNLPFrameworkHandle.PredictVectorForStringInputStringError]
//   - [IMLNLPFrameworkHandle.Valid]
type IMLNLPFrameworkHandle interface {
	objectivec.IObject

	// Topic: Methods

	NLPClassifierModelCopyPredictedLabelForTextImpl() VoidHandler
	NLPClassifierModelCreateWithDataImpl() VoidHandler
	NLPClassifierModelGetCurrentRevisionImpl() VoidHandler
	NLPClassifierModelGetRevisionImpl() VoidHandler
	NLPClassifierModelIsRevisionSupportedImpl() VoidHandler
	NLPEmbeddingModelCopyVectorForStringImpl() VoidHandler
	NLPEmbeddingModelCreateWithDataImpl() VoidHandler
	NLPEmbeddingModelGetCurrentRevisionImpl() VoidHandler
	NLPEmbeddingModelGetRevisionImpl() VoidHandler
	NLPEmbeddingModelIsRevisionSupportedImpl() VoidHandler
	NLPGazetteerModelCopyLabelForStringImpl() VoidHandler
	NLPGazetteerModelCreateWithDataImpl() VoidHandler
	NLPGazetteerModelGetCurrentRevisionImpl() VoidHandler
	NLPGazetteerModelGetRevisionImpl() VoidHandler
	NLPGazetteerModelIsRevisionSupportedImpl() VoidHandler
	NLPSequenceModelCopyPredictedTokensAndLabelsForTextImpl() VoidHandler
	NLPSequenceModelCreateWithDataImpl() VoidHandler
	NLPSequenceModelGetCurrentRevisionImpl() VoidHandler
	NLPSequenceModelGetRevisionImpl() VoidHandler
	NLPSequenceModelIsRevisionSupportedImpl() VoidHandler
	InitializeEmbeddingModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error)
	InitializeGazetteerModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error)
	InitializeSentenceClassifierModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error)
	InitializeWordTaggingModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error)
	IsValid() bool
	PredictLabelForWordStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error)
	PredictLabelsForSentenceStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error)
	PredictTokensLabelsLocationsLengthsForStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error)
	PredictVectorForStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error)
	Valid() bool
}

// Init initializes the instance.
func (m MLNLPFrameworkHandle) Init() MLNLPFrameworkHandle {
	rv := objc.Send[MLNLPFrameworkHandle](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNLPFrameworkHandle) Autorelease() MLNLPFrameworkHandle {
	rv := objc.Send[MLNLPFrameworkHandle](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNLPFrameworkHandle creates a new MLNLPFrameworkHandle instance.
func NewMLNLPFrameworkHandle() MLNLPFrameworkHandle {
	class := getMLNLPFrameworkHandleClass()
	rv := objc.Send[MLNLPFrameworkHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLNLPFrameworkHandle) InitializeEmbeddingModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initializeEmbeddingModelWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(MLNLPFrameworkHandle), foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPFrameworkHandleFromID(rv), nil

}
func (m MLNLPFrameworkHandle) InitializeGazetteerModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initializeGazetteerModelWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(MLNLPFrameworkHandle), foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPFrameworkHandleFromID(rv), nil

}
func (m MLNLPFrameworkHandle) InitializeSentenceClassifierModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initializeSentenceClassifierModelWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(MLNLPFrameworkHandle), foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPFrameworkHandleFromID(rv), nil

}
func (m MLNLPFrameworkHandle) InitializeWordTaggingModelWithDataError(data objectivec.IObject) (MLNLPFrameworkHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initializeWordTaggingModelWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(MLNLPFrameworkHandle), foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPFrameworkHandleFromID(rv), nil

}
func (m MLNLPFrameworkHandle) IsValid() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isValid"))
	return rv
}
func (m MLNLPFrameworkHandle) PredictLabelForWordStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictLabelForWordString:inputString:error:"), string_, string_2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNLPFrameworkHandle) PredictLabelsForSentenceStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictLabelsForSentenceString:inputString:error:"), string_, string_2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNLPFrameworkHandle) PredictTokensLabelsLocationsLengthsForStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictTokensLabelsLocationsLengthsForString:inputString:error:"), string_, string_2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNLPFrameworkHandle) PredictVectorForStringInputStringError(string_ unsafe.Pointer, string_2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictVectorForString:inputString:error:"), string_, string_2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (_MLNLPFrameworkHandleClass MLNLPFrameworkHandleClass) SharedHandle() MLNLPFrameworkHandle {
	rv := objc.Send[objc.ID](objc.ID(_MLNLPFrameworkHandleClass.class), objc.Sel("sharedHandle"))
	return MLNLPFrameworkHandleFromID(rv)
}

func (m MLNLPFrameworkHandle) NLPClassifierModelCopyPredictedLabelForTextImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPClassifierModelCopyPredictedLabelForTextImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPClassifierModelCreateWithDataImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPClassifierModelCreateWithDataImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPClassifierModelGetCurrentRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPClassifierModelGetCurrentRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPClassifierModelGetRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPClassifierModelGetRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPClassifierModelIsRevisionSupportedImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPClassifierModelIsRevisionSupportedImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPEmbeddingModelCopyVectorForStringImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPEmbeddingModelCopyVectorForStringImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPEmbeddingModelCreateWithDataImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPEmbeddingModelCreateWithDataImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPEmbeddingModelGetCurrentRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPEmbeddingModelGetCurrentRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPEmbeddingModelGetRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPEmbeddingModelGetRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPEmbeddingModelIsRevisionSupportedImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPEmbeddingModelIsRevisionSupportedImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPGazetteerModelCopyLabelForStringImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPGazetteerModelCopyLabelForStringImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPGazetteerModelCreateWithDataImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPGazetteerModelCreateWithDataImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPGazetteerModelGetCurrentRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPGazetteerModelGetCurrentRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPGazetteerModelGetRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPGazetteerModelGetRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPGazetteerModelIsRevisionSupportedImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPGazetteerModelIsRevisionSupportedImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPSequenceModelCopyPredictedTokensAndLabelsForTextImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPSequenceModelCopyPredictedTokensAndLabelsForTextImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPSequenceModelCreateWithDataImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPSequenceModelCreateWithDataImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPSequenceModelGetCurrentRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPSequenceModelGetCurrentRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPSequenceModelGetRevisionImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPSequenceModelGetRevisionImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) NLPSequenceModelIsRevisionSupportedImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NLPSequenceModelIsRevisionSupportedImpl"))
	_ = rv
	return nil
}
func (m MLNLPFrameworkHandle) Valid() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("valid"))
	return rv
}
