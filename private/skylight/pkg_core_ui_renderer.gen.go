// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGCoreUIRenderer] class.
var (
	_PKGCoreUIRendererClass     PKGCoreUIRendererClass
	_PKGCoreUIRendererClassOnce sync.Once
)

func getPKGCoreUIRendererClass() PKGCoreUIRendererClass {
	_PKGCoreUIRendererClassOnce.Do(func() {
		_PKGCoreUIRendererClass = PKGCoreUIRendererClass{class: objc.GetClass("PKGCoreUIRenderer")}
	})
	return _PKGCoreUIRendererClass
}

// GetPKGCoreUIRendererClass returns the class object for PKGCoreUIRenderer.
func GetPKGCoreUIRendererClass() PKGCoreUIRendererClass {
	return getPKGCoreUIRendererClass()
}

type PKGCoreUIRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCoreUIRendererClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCoreUIRendererClass) Alloc() PKGCoreUIRenderer {
	rv := objc.Send[PKGCoreUIRenderer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGCoreUIRenderer.Renderer]
//   - [PKGCoreUIRenderer.RendererName]
//   - [PKGCoreUIRenderer.InitWithRendererName]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer
type PKGCoreUIRenderer struct {
	objectivec.Object
}

// PKGCoreUIRendererFromID constructs a [PKGCoreUIRenderer] from an objc.ID.
func PKGCoreUIRendererFromID(id objc.ID) PKGCoreUIRenderer {
	return PKGCoreUIRenderer{objectivec.Object{ID: id}}
}

// Ensure PKGCoreUIRenderer implements IPKGCoreUIRenderer.
var _ IPKGCoreUIRenderer = PKGCoreUIRenderer{}

// An interface definition for the [PKGCoreUIRenderer] class.
//
// # Methods
//
//   - [IPKGCoreUIRenderer.Renderer]
//   - [IPKGCoreUIRenderer.RendererName]
//   - [IPKGCoreUIRenderer.InitWithRendererName]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer
type IPKGCoreUIRenderer interface {
	objectivec.IObject

	// Topic: Methods

	Renderer() OpaqueCUIRendererRefRef
	RendererName() objectivec.IObject
	InitWithRendererName(name objectivec.IObject) PKGCoreUIRenderer
}

// Init initializes the instance.
func (g PKGCoreUIRenderer) Init() PKGCoreUIRenderer {
	rv := objc.Send[PKGCoreUIRenderer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g PKGCoreUIRenderer) Autorelease() PKGCoreUIRenderer {
	rv := objc.Send[PKGCoreUIRenderer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCoreUIRenderer creates a new PKGCoreUIRenderer instance.
func NewPKGCoreUIRenderer() PKGCoreUIRenderer {
	class := getPKGCoreUIRendererClass()
	rv := objc.Send[PKGCoreUIRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer/initWithRendererName:
func NewGCoreUIRendererWithRendererName(name objectivec.IObject) PKGCoreUIRenderer {
	instance := getPKGCoreUIRendererClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRendererName:"), name)
	return PKGCoreUIRendererFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer/renderer
func (g PKGCoreUIRenderer) Renderer() OpaqueCUIRendererRefRef {
	rv := objc.Send[OpaqueCUIRendererRefRef](g.ID, objc.Sel("renderer"))
	return OpaqueCUIRendererRefRef(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer/rendererName
func (g PKGCoreUIRenderer) RendererName() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rendererName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer/initWithRendererName:
func (g PKGCoreUIRenderer) InitWithRendererName(name objectivec.IObject) PKGCoreUIRenderer {
	rv := objc.Send[PKGCoreUIRenderer](g.ID, objc.Sel("initWithRendererName:"), name)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIRenderer/rendererForTheme:useAX:
func (_PKGCoreUIRendererClass PKGCoreUIRendererClass) RendererForThemeUseAX(theme uint32, ax bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_PKGCoreUIRendererClass.class), objc.Sel("rendererForTheme:useAX:"), theme, ax)
	return objectivec.Object{ID: rv}
}
