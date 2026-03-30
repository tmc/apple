// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLInternalNLPModelWriter] class.
var (
	_MLInternalNLPModelWriterClass     MLInternalNLPModelWriterClass
	_MLInternalNLPModelWriterClassOnce sync.Once
)

func getMLInternalNLPModelWriterClass() MLInternalNLPModelWriterClass {
	_MLInternalNLPModelWriterClassOnce.Do(func() {
		_MLInternalNLPModelWriterClass = MLInternalNLPModelWriterClass{class: objc.GetClass("_MLInternalNLPModelWriter")}
	})
	return _MLInternalNLPModelWriterClass
}

// GetMLInternalNLPModelWriterClass returns the class object for _MLInternalNLPModelWriter.
func GetMLInternalNLPModelWriterClass() MLInternalNLPModelWriterClass {
	return getMLInternalNLPModelWriterClass()
}

type MLInternalNLPModelWriterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLInternalNLPModelWriterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLInternalNLPModelWriterClass) Alloc() MLInternalNLPModelWriter {
	rv := objc.Send[MLInternalNLPModelWriter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter
type MLInternalNLPModelWriter struct {
	objectivec.Object
}

// MLInternalNLPModelWriterFromID constructs a [MLInternalNLPModelWriter] from an objc.ID.
func MLInternalNLPModelWriterFromID(id objc.ID) MLInternalNLPModelWriter {
	return MLInternalNLPModelWriter{objectivec.Object{ID: id}}
}

// Ensure MLInternalNLPModelWriter implements IMLInternalNLPModelWriter.
var _ IMLInternalNLPModelWriter = MLInternalNLPModelWriter{}

// An interface definition for the [MLInternalNLPModelWriter] class.
//
// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter
type IMLInternalNLPModelWriter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLInternalNLPModelWriter) Init() MLInternalNLPModelWriter {
	rv := objc.Send[MLInternalNLPModelWriter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLInternalNLPModelWriter) Autorelease() MLInternalNLPModelWriter {
	rv := objc.Send[MLInternalNLPModelWriter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLInternalNLPModelWriter creates a new MLInternalNLPModelWriter instance.
func NewMLInternalNLPModelWriter() MLInternalNLPModelWriter {
	class := getMLInternalNLPModelWriterClass()
	rv := objc.Send[MLInternalNLPModelWriter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter/saveCustomSentenceClassifierModelToURL:modelData:stringInputName:classname:NSError:
func (_MLInternalNLPModelWriterClass MLInternalNLPModelWriterClass) SaveCustomSentenceClassifierModelToURLModelDataStringInputNameClassnameNSError(url foundation.INSURL, data objectivec.IObject, name objectivec.IObject, classname objectivec.IObject, sError []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLInternalNLPModelWriterClass.class), objc.Sel("saveCustomSentenceClassifierModelToURL:modelData:stringInputName:classname:NSError:"), url, data, name, classname, objectivec.IObjectSliceToNSArray(sError))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter/saveCustomSentenceModelToURL:modelData:stringInputName:classname:NSError:
func (_MLInternalNLPModelWriterClass MLInternalNLPModelWriterClass) SaveCustomSentenceModelToURLModelDataStringInputNameClassnameNSError(url foundation.INSURL, data objectivec.IObject, name objectivec.IObject, classname objectivec.IObject, sError []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLInternalNLPModelWriterClass.class), objc.Sel("saveCustomSentenceModelToURL:modelData:stringInputName:classname:NSError:"), url, data, name, classname, objectivec.IObjectSliceToNSArray(sError))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter/saveCustomSequenceModelToURL:modelData:stringInputName:classname:NSError:
func (_MLInternalNLPModelWriterClass MLInternalNLPModelWriterClass) SaveCustomSequenceModelToURLModelDataStringInputNameClassnameNSError(url foundation.INSURL, data objectivec.IObject, name objectivec.IObject, classname objectivec.IObject, sError []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLInternalNLPModelWriterClass.class), objc.Sel("saveCustomSequenceModelToURL:modelData:stringInputName:classname:NSError:"), url, data, name, classname, objectivec.IObjectSliceToNSArray(sError))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLInternalNLPModelWriter/saveCustomWordTaggingModelToURL:modelData:stringInputName:classname:NSError:
func (_MLInternalNLPModelWriterClass MLInternalNLPModelWriterClass) SaveCustomWordTaggingModelToURLModelDataStringInputNameClassnameNSError(url foundation.INSURL, data objectivec.IObject, name objectivec.IObject, classname objectivec.IObject, sError []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLInternalNLPModelWriterClass.class), objc.Sel("saveCustomWordTaggingModelToURL:modelData:stringInputName:classname:NSError:"), url, data, name, classname, objectivec.IObjectSliceToNSArray(sError))
	return rv
}
