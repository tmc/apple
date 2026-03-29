// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLVNFrameworkHandle] class.
var (
	_MLVNFrameworkHandleClass     MLVNFrameworkHandleClass
	_MLVNFrameworkHandleClassOnce sync.Once
)

func getMLVNFrameworkHandleClass() MLVNFrameworkHandleClass {
	_MLVNFrameworkHandleClassOnce.Do(func() {
		_MLVNFrameworkHandleClass = MLVNFrameworkHandleClass{class: objc.GetClass("_MLVNFrameworkHandle")}
	})
	return _MLVNFrameworkHandleClass
}

// GetMLVNFrameworkHandleClass returns the class object for _MLVNFrameworkHandle.
func GetMLVNFrameworkHandleClass() MLVNFrameworkHandleClass {
	return getMLVNFrameworkHandleClass()
}

type MLVNFrameworkHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLVNFrameworkHandleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLVNFrameworkHandleClass) Alloc() MLVNFrameworkHandle {
	rv := objc.Send[MLVNFrameworkHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLVNFrameworkHandle.VNImageBufferClass]
//   - [MLVNFrameworkHandle.CreatePixelBufferFromCGImageConstraintCropRectCropAndScaleOptionOptionsError]
//   - [MLVNFrameworkHandle.CreatePixelBufferFromImageAtURLConstraintCropRectCropAndScaleOptionOptionsError]
//   - [MLVNFrameworkHandle.CreatePixelBufferFromVNImageBufferConstraintCropRectCropAndScaleOptionOptionsError]
//   - [MLVNFrameworkHandle.DetectionPrintShapes]
//   - [MLVNFrameworkHandle.DetectionPrintShapesImpl]
//   - [MLVNFrameworkHandle.DetectionPrintSupportedRevisions]
//   - [MLVNFrameworkHandle.DetectionPrintSupportedRevisionsImpl]
//   - [MLVNFrameworkHandle.DetectionPrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError]
//   - [MLVNFrameworkHandle.DetectionPrintsFromPixelBuffersImpl]
//   - [MLVNFrameworkHandle.DetectionPrintsFromPixelBuffersUsesCPUOnlyImpl]
//   - [MLVNFrameworkHandle.ElementCountForScenePrintRequestRevision]
//   - [MLVNFrameworkHandle.IsValid]
//   - [MLVNFrameworkHandle.LengthInBytesForScenePrintRequestRevision]
//   - [MLVNFrameworkHandle.ScenePrintElementCountImpl]
//   - [MLVNFrameworkHandle.ScenePrintLengthImpl]
//   - [MLVNFrameworkHandle.ScenePrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError]
//   - [MLVNFrameworkHandle.ScenePrintsFromPixelBuffersImpl]
//   - [MLVNFrameworkHandle.ScenePrintsFromPixelBuffersUsesCPUOnlyImpl]
//   - [MLVNFrameworkHandle.ValidForObjectprint]
//   - [MLVNFrameworkHandle.ValidForSceneprint]
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle
type MLVNFrameworkHandle struct {
	objectivec.Object
}

// MLVNFrameworkHandleFromID constructs a [MLVNFrameworkHandle] from an objc.ID.
func MLVNFrameworkHandleFromID(id objc.ID) MLVNFrameworkHandle {
	return MLVNFrameworkHandle{objectivec.Object{ID: id}}
}
// Ensure MLVNFrameworkHandle implements IMLVNFrameworkHandle.
var _ IMLVNFrameworkHandle = MLVNFrameworkHandle{}

// An interface definition for the [MLVNFrameworkHandle] class.
//
// # Methods
//
//   - [IMLVNFrameworkHandle.VNImageBufferClass]
//   - [IMLVNFrameworkHandle.CreatePixelBufferFromCGImageConstraintCropRectCropAndScaleOptionOptionsError]
//   - [IMLVNFrameworkHandle.CreatePixelBufferFromImageAtURLConstraintCropRectCropAndScaleOptionOptionsError]
//   - [IMLVNFrameworkHandle.CreatePixelBufferFromVNImageBufferConstraintCropRectCropAndScaleOptionOptionsError]
//   - [IMLVNFrameworkHandle.DetectionPrintShapes]
//   - [IMLVNFrameworkHandle.DetectionPrintShapesImpl]
//   - [IMLVNFrameworkHandle.DetectionPrintSupportedRevisions]
//   - [IMLVNFrameworkHandle.DetectionPrintSupportedRevisionsImpl]
//   - [IMLVNFrameworkHandle.DetectionPrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError]
//   - [IMLVNFrameworkHandle.DetectionPrintsFromPixelBuffersImpl]
//   - [IMLVNFrameworkHandle.DetectionPrintsFromPixelBuffersUsesCPUOnlyImpl]
//   - [IMLVNFrameworkHandle.ElementCountForScenePrintRequestRevision]
//   - [IMLVNFrameworkHandle.IsValid]
//   - [IMLVNFrameworkHandle.LengthInBytesForScenePrintRequestRevision]
//   - [IMLVNFrameworkHandle.ScenePrintElementCountImpl]
//   - [IMLVNFrameworkHandle.ScenePrintLengthImpl]
//   - [IMLVNFrameworkHandle.ScenePrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError]
//   - [IMLVNFrameworkHandle.ScenePrintsFromPixelBuffersImpl]
//   - [IMLVNFrameworkHandle.ScenePrintsFromPixelBuffersUsesCPUOnlyImpl]
//   - [IMLVNFrameworkHandle.ValidForObjectprint]
//   - [IMLVNFrameworkHandle.ValidForSceneprint]
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle
type IMLVNFrameworkHandle interface {
	objectivec.IObject

	// Topic: Methods

	VNImageBufferClass() objc.Class
	CreatePixelBufferFromCGImageConstraintCropRectCropAndScaleOptionOptionsError(cGImage *coregraphics.CGImageRef, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error)
	CreatePixelBufferFromImageAtURLConstraintCropRectCropAndScaleOptionOptionsError(url foundation.INSURL, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error)
	CreatePixelBufferFromVNImageBufferConstraintCropRectCropAndScaleOptionOptionsError(buffer objectivec.IObject, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error)
	DetectionPrintShapes(shapes uint64) objectivec.IObject
	DetectionPrintShapesImpl() VoidHandler
	DetectionPrintSupportedRevisions() objectivec.IObject
	DetectionPrintSupportedRevisionsImpl() VoidHandler
	DetectionPrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError(buffers corevideo.CVImageBufferRef, version uint64, options objectivec.IObject, cPUOnly bool) (objectivec.IObject, error)
	DetectionPrintsFromPixelBuffersImpl() VoidHandler
	DetectionPrintsFromPixelBuffersUsesCPUOnlyImpl() VoidHandler
	ElementCountForScenePrintRequestRevision(revision uint64) uint64
	IsValid() bool
	LengthInBytesForScenePrintRequestRevision(revision uint64) uint64
	ScenePrintElementCountImpl() VoidHandler
	ScenePrintLengthImpl() VoidHandler
	ScenePrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError(buffers corevideo.CVImageBufferRef, version uint64, options objectivec.IObject, cPUOnly bool) (objectivec.IObject, error)
	ScenePrintsFromPixelBuffersImpl() VoidHandler
	ScenePrintsFromPixelBuffersUsesCPUOnlyImpl() VoidHandler
	ValidForObjectprint() bool
	ValidForSceneprint() bool
}

// Init initializes the instance.
func (m MLVNFrameworkHandle) Init() MLVNFrameworkHandle {
	rv := objc.Send[MLVNFrameworkHandle](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLVNFrameworkHandle) Autorelease() MLVNFrameworkHandle {
	rv := objc.Send[MLVNFrameworkHandle](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLVNFrameworkHandle creates a new MLVNFrameworkHandle instance.
func NewMLVNFrameworkHandle() MLVNFrameworkHandle {
	class := getMLVNFrameworkHandleClass()
	rv := objc.Send[MLVNFrameworkHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/createPixelBufferFromCGImage:constraint:cropRect:cropAndScaleOption:options:error:
func (m MLVNFrameworkHandle) CreatePixelBufferFromCGImageConstraintCropRectCropAndScaleOptionOptionsError(cGImage *coregraphics.CGImageRef, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("createPixelBufferFromCGImage:constraint:cropRect:cropAndScaleOption:options:error:"), cGImage, constraint, rect, option, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/createPixelBufferFromImageAtURL:constraint:cropRect:cropAndScaleOption:options:error:
func (m MLVNFrameworkHandle) CreatePixelBufferFromImageAtURLConstraintCropRectCropAndScaleOptionOptionsError(url foundation.INSURL, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("createPixelBufferFromImageAtURL:constraint:cropRect:cropAndScaleOption:options:error:"), url, constraint, rect, option, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/createPixelBufferFromVNImageBuffer:constraint:cropRect:cropAndScaleOption:options:error:
func (m MLVNFrameworkHandle) CreatePixelBufferFromVNImageBufferConstraintCropRectCropAndScaleOptionOptionsError(buffer objectivec.IObject, constraint objectivec.IObject, rect corefoundation.CGRect, option uint64, options objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("createPixelBufferFromVNImageBuffer:constraint:cropRect:cropAndScaleOption:options:error:"), buffer, constraint, rect, option, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintShapes:
func (m MLVNFrameworkHandle) DetectionPrintShapes(shapes uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintShapes:"), shapes)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintSupportedRevisions
func (m MLVNFrameworkHandle) DetectionPrintSupportedRevisions() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintSupportedRevisions"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintsFromPixelBuffers:version:augmentationOptions:useCPUOnly:error:
func (m MLVNFrameworkHandle) DetectionPrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError(buffers corevideo.CVImageBufferRef, version uint64, options objectivec.IObject, cPUOnly bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintsFromPixelBuffers:version:augmentationOptions:useCPUOnly:error:"), buffers, version, options, cPUOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/elementCountForScenePrintRequestRevision:
func (m MLVNFrameworkHandle) ElementCountForScenePrintRequestRevision(revision uint64) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("elementCountForScenePrintRequestRevision:"), revision)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/isValid
func (m MLVNFrameworkHandle) IsValid() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isValid"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/lengthInBytesForScenePrintRequestRevision:
func (m MLVNFrameworkHandle) LengthInBytesForScenePrintRequestRevision(revision uint64) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("lengthInBytesForScenePrintRequestRevision:"), revision)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/scenePrintsFromPixelBuffers:version:augmentationOptions:useCPUOnly:error:
func (m MLVNFrameworkHandle) ScenePrintsFromPixelBuffersVersionAugmentationOptionsUseCPUOnlyError(buffers corevideo.CVImageBufferRef, version uint64, options objectivec.IObject, cPUOnly bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scenePrintsFromPixelBuffers:version:augmentationOptions:useCPUOnly:error:"), buffers, version, options, cPUOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/addOrientation:toOptions:
func (_MLVNFrameworkHandleClass MLVNFrameworkHandleClass) AddOrientationToOptions(orientation uint32, options objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLVNFrameworkHandleClass.class), objc.Sel("addOrientation:toOptions:"), orientation, options)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/sharedHandle
func (_MLVNFrameworkHandleClass MLVNFrameworkHandleClass) SharedHandle() *MLVNFrameworkHandle {
	rv := objc.Send[objc.ID](objc.ID(_MLVNFrameworkHandleClass.class), objc.Sel("sharedHandle"))
	if rv == 0 {
		return nil
	}
	val := MLVNFrameworkHandleFromID(rv)
	return &val
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/VNImageBufferClass
func (m MLVNFrameworkHandle) VNImageBufferClass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("VNImageBufferClass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintShapesImpl
func (m MLVNFrameworkHandle) DetectionPrintShapesImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintShapesImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintSupportedRevisionsImpl
func (m MLVNFrameworkHandle) DetectionPrintSupportedRevisionsImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintSupportedRevisionsImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintsFromPixelBuffersImpl
func (m MLVNFrameworkHandle) DetectionPrintsFromPixelBuffersImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintsFromPixelBuffersImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/detectionPrintsFromPixelBuffersUsesCPUOnlyImpl
func (m MLVNFrameworkHandle) DetectionPrintsFromPixelBuffersUsesCPUOnlyImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("detectionPrintsFromPixelBuffersUsesCPUOnlyImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/scenePrintElementCountImpl
func (m MLVNFrameworkHandle) ScenePrintElementCountImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scenePrintElementCountImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/scenePrintLengthImpl
func (m MLVNFrameworkHandle) ScenePrintLengthImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scenePrintLengthImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/scenePrintsFromPixelBuffersImpl
func (m MLVNFrameworkHandle) ScenePrintsFromPixelBuffersImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scenePrintsFromPixelBuffersImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/scenePrintsFromPixelBuffersUsesCPUOnlyImpl
func (m MLVNFrameworkHandle) ScenePrintsFromPixelBuffersUsesCPUOnlyImpl() VoidHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("scenePrintsFromPixelBuffersUsesCPUOnlyImpl"))
	_ = rv
	return nil
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/validForObjectprint
func (m MLVNFrameworkHandle) ValidForObjectprint() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("validForObjectprint"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/_MLVNFrameworkHandle/validForSceneprint
func (m MLVNFrameworkHandle) ValidForSceneprint() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("validForSceneprint"))
	return rv
}

