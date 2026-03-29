// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNetworkHeaderEncoding] class.
var (
	_MLNetworkHeaderEncodingClass     MLNetworkHeaderEncodingClass
	_MLNetworkHeaderEncodingClassOnce sync.Once
)

func getMLNetworkHeaderEncodingClass() MLNetworkHeaderEncodingClass {
	_MLNetworkHeaderEncodingClassOnce.Do(func() {
		_MLNetworkHeaderEncodingClass = MLNetworkHeaderEncodingClass{class: objc.GetClass("_MLNetworkHeaderEncoding")}
	})
	return _MLNetworkHeaderEncodingClass
}

// GetMLNetworkHeaderEncodingClass returns the class object for _MLNetworkHeaderEncoding.
func GetMLNetworkHeaderEncodingClass() MLNetworkHeaderEncodingClass {
	return getMLNetworkHeaderEncodingClass()
}

type MLNetworkHeaderEncodingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNetworkHeaderEncodingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNetworkHeaderEncodingClass) Alloc() MLNetworkHeaderEncoding {
	rv := objc.Send[MLNetworkHeaderEncoding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding
type MLNetworkHeaderEncoding struct {
	objectivec.Object
}

// MLNetworkHeaderEncodingFromID constructs a [MLNetworkHeaderEncoding] from an objc.ID.
func MLNetworkHeaderEncodingFromID(id objc.ID) MLNetworkHeaderEncoding {
	return MLNetworkHeaderEncoding{objectivec.Object{ID: id}}
}
// Ensure MLNetworkHeaderEncoding implements IMLNetworkHeaderEncoding.
var _ IMLNetworkHeaderEncoding = MLNetworkHeaderEncoding{}

// An interface definition for the [MLNetworkHeaderEncoding] class.
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding
type IMLNetworkHeaderEncoding interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLNetworkHeaderEncoding) Init() MLNetworkHeaderEncoding {
	rv := objc.Send[MLNetworkHeaderEncoding](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworkHeaderEncoding) Autorelease() MLNetworkHeaderEncoding {
	rv := objc.Send[MLNetworkHeaderEncoding](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworkHeaderEncoding creates a new MLNetworkHeaderEncoding instance.
func NewMLNetworkHeaderEncoding() MLNetworkHeaderEncoding {
	class := getMLNetworkHeaderEncodingClass()
	rv := objc.Send[MLNetworkHeaderEncoding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/acknowledgeFailData
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) AcknowledgeFailData() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("acknowledgeFailData"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/acknowledgeSucessData
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) AcknowledgeSucessData() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("acknowledgeSucessData"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/constructDataPacket:HeaderEncoding:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) ConstructDataPacketHeaderEncoding(packet objectivec.IObject, encoding uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("constructDataPacket:HeaderEncoding:"), packet, encoding)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/custom:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) Custom(custom objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("custom:"), custom)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/getHeaderDataSize:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderDataSize(size objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderDataSize:"), size)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/getHeaderDataStart:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderDataStart(start objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderDataStart:"), start)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/getHeaderEncoding:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderEncoding(encoding objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderEncoding:"), encoding)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/getHeaderEnd:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderEnd(end objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderEnd:"), end)
	return rv
}
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/getHeaderSize
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderSize() uint64 {
	rv := objc.Send[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderSize"))
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderAcknowledgeFailData:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderAcknowledgeFailData(data uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderAcknowledgeFailData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderAcknowledgeSucessData:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderAcknowledgeSucessData(data uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderAcknowledgeSucessData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderCustom:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderCustom(custom uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderCustom:"), custom)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderError:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderError(error_ uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderError:"), error_)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderIncomingData:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderIncomingData(data uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderIncomingData:"), data)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderLoad:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderLoad(load uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderLoad:"), load)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderPredictFeature:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderPredictFeature(feature uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderPredictFeature:"), feature)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderText:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderText(text uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderText:"), text)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/isHeaderUnLoad:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderUnLoad(load uint64) bool {
	rv := objc.Send[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderUnLoad:"), load)
	return rv
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/loadModel:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) LoadModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("loadModel:"), model)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/predictFeature:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) PredictFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("predictFeature:"), feature)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/textDebug:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) TextDebug(debug objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("textDebug:"), debug)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/RemoteCoreML/_MLNetworkHeaderEncoding/unLoadModel:
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) UnLoadModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("unLoadModel:"), model)
	return objectivec.Object{ID: rv}
}

