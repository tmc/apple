// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerStreamDataForMetadata] class.
var (
	_GTShaderProfilerStreamDataForMetadataClass     GTShaderProfilerStreamDataForMetadataClass
	_GTShaderProfilerStreamDataForMetadataClassOnce sync.Once
)

func getGTShaderProfilerStreamDataForMetadataClass() GTShaderProfilerStreamDataForMetadataClass {
	_GTShaderProfilerStreamDataForMetadataClassOnce.Do(func() {
		_GTShaderProfilerStreamDataForMetadataClass = GTShaderProfilerStreamDataForMetadataClass{class: objc.GetClass("GTShaderProfilerStreamDataForMetadata")}
	})
	return _GTShaderProfilerStreamDataForMetadataClass
}

// GetGTShaderProfilerStreamDataForMetadataClass returns the class object for GTShaderProfilerStreamDataForMetadata.
func GetGTShaderProfilerStreamDataForMetadataClass() GTShaderProfilerStreamDataForMetadataClass {
	return getGTShaderProfilerStreamDataForMetadataClass()
}

type GTShaderProfilerStreamDataForMetadataClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerStreamDataForMetadataClass) Alloc() GTShaderProfilerStreamDataForMetadata {
	rv := objc.Send[GTShaderProfilerStreamDataForMetadata](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataForMetadata
type GTShaderProfilerStreamDataForMetadata struct {
	GTShaderProfilerStreamData
}

// GTShaderProfilerStreamDataForMetadataFromID constructs a [GTShaderProfilerStreamDataForMetadata] from an objc.ID.
func GTShaderProfilerStreamDataForMetadataFromID(id objc.ID) GTShaderProfilerStreamDataForMetadata {
	return GTShaderProfilerStreamDataForMetadata{GTShaderProfilerStreamData: GTShaderProfilerStreamDataFromID(id)}
}
// Ensure GTShaderProfilerStreamDataForMetadata implements IGTShaderProfilerStreamDataForMetadata.
var _ IGTShaderProfilerStreamDataForMetadata = GTShaderProfilerStreamDataForMetadata{}

// An interface definition for the [GTShaderProfilerStreamDataForMetadata] class.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataForMetadata
type IGTShaderProfilerStreamDataForMetadata interface {
	IGTShaderProfilerStreamData
}

// Init initializes the instance.
func (g GTShaderProfilerStreamDataForMetadata) Init() GTShaderProfilerStreamDataForMetadata {
	rv := objc.Send[GTShaderProfilerStreamDataForMetadata](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStreamDataForMetadata) Autorelease() GTShaderProfilerStreamDataForMetadata {
	rv := objc.Send[GTShaderProfilerStreamDataForMetadata](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStreamDataForMetadata creates a new GTShaderProfilerStreamDataForMetadata instance.
func NewGTShaderProfilerStreamDataForMetadata() GTShaderProfilerStreamDataForMetadata {
	class := getGTShaderProfilerStreamDataForMetadataClass()
	rv := objc.Send[GTShaderProfilerStreamDataForMetadata](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamDataForMetadata/initWithCoder:
func NewGTShaderProfilerStreamDataForMetadataWithCoder(coder objectivec.IObject) GTShaderProfilerStreamDataForMetadata {
	instance := getGTShaderProfilerStreamDataForMetadataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerStreamDataForMetadataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamData/initWithNewFileFormatV2Support:
func NewGTShaderProfilerStreamDataForMetadataWithNewFileFormatV2Support(v2Support bool) GTShaderProfilerStreamDataForMetadata {
	instance := getGTShaderProfilerStreamDataForMetadataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNewFileFormatV2Support:"), v2Support)
	return GTShaderProfilerStreamDataForMetadataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerStreamData/initWithPreSiBundle:
func NewGTShaderProfilerStreamDataForMetadataWithPreSiBundle(bundle objectivec.IObject) GTShaderProfilerStreamDataForMetadata {
	instance := getGTShaderProfilerStreamDataForMetadataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreSiBundle:"), bundle)
	return GTShaderProfilerStreamDataForMetadataFromID(rv)
}

