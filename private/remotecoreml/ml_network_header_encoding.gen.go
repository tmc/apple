// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

package remotecoreml

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[MLNetworkHeaderEncoding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
type IMLNetworkHeaderEncoding interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLNetworkHeaderEncoding) Init() MLNetworkHeaderEncoding {
	rv := objc.SendIfResponds[MLNetworkHeaderEncoding](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNetworkHeaderEncoding) Autorelease() MLNetworkHeaderEncoding {
	rv := objc.SendIfResponds[MLNetworkHeaderEncoding](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNetworkHeaderEncoding creates a new MLNetworkHeaderEncoding instance.
func NewMLNetworkHeaderEncoding() MLNetworkHeaderEncoding {
	class := getMLNetworkHeaderEncodingClass()
	rv := objc.SendIfResponds[MLNetworkHeaderEncoding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) AcknowledgeFailData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("acknowledgeFailData"))
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) AcknowledgeSucessData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("acknowledgeSucessData"))
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) ConstructDataPacketHeaderEncoding(packet objectivec.IObject, encoding uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("constructDataPacket:HeaderEncoding:"), packet, encoding)
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) Custom(custom objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("custom:"), custom)
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderDataSize(size objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderDataSize:"), size)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderDataStart(start objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderDataStart:"), start)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderEncoding(encoding objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderEncoding:"), encoding)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderEnd(end objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderEnd:"), end)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) GetHeaderSize() uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("getHeaderSize"))
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderAcknowledgeFailData(data uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderAcknowledgeFailData:"), data)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderAcknowledgeSucessData(data uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderAcknowledgeSucessData:"), data)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderCustom(custom uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderCustom:"), custom)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderError(error_ uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderError:"), error_)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderIncomingData(data uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderIncomingData:"), data)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderLoad(load uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderLoad:"), load)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderPredictFeature(feature uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderPredictFeature:"), feature)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderText(text uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderText:"), text)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) IsHeaderUnLoad(load uint64) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("isHeaderUnLoad:"), load)
	return rv
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) LoadModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("loadModel:"), model)
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) PredictFeature(feature objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("predictFeature:"), feature)
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) TextDebug(debug objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("textDebug:"), debug)
	return objectivec.Object{ID: rv}
}
func (_MLNetworkHeaderEncodingClass MLNetworkHeaderEncodingClass) UnLoadModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNetworkHeaderEncodingClass.class), objc.Sel("unLoadModel:"), model)
	return objectivec.Object{ID: rv}
}
