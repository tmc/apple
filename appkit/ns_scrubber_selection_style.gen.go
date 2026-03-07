// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScrubberSelectionStyle] class.
var (
	_NSScrubberSelectionStyleClass     NSScrubberSelectionStyleClass
	_NSScrubberSelectionStyleClassOnce sync.Once
)

func getNSScrubberSelectionStyleClass() NSScrubberSelectionStyleClass {
	_NSScrubberSelectionStyleClassOnce.Do(func() {
		_NSScrubberSelectionStyleClass = NSScrubberSelectionStyleClass{class: objc.GetClass("NSScrubberSelectionStyle")}
	})
	return _NSScrubberSelectionStyleClass
}

// GetNSScrubberSelectionStyleClass returns the class object for NSScrubberSelectionStyle.
func GetNSScrubberSelectionStyleClass() NSScrubberSelectionStyleClass {
	return getNSScrubberSelectionStyleClass()
}

type NSScrubberSelectionStyleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberSelectionStyleClass) Alloc() NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class that provides decorative accessory views for selected and
// highlighted items within a scrubber control.
//
// # Overview
// 
// Choose a selection style ([NSScrubberSelectionStyle.OutlineOverlayStyle] or
// [NSScrubberSelectionStyle.RoundedBackgroundStyle]), or create a custom selection style by
// subclassing [NSScrubberSelectionStyle] and overriding [SelectionView].
//
// # Creating a selection style
//
//   - [NSScrubberSelectionStyle.InitWithCoder]: Initializes a scrubber selection style when included from a nib or Storyboard.
//   - [NSScrubberSelectionStyle.MakeSelectionView]: Provides an opportunity to create a customized scrubber selection style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle
type NSScrubberSelectionStyle struct {
	objectivec.Object
}

// NSScrubberSelectionStyleFromID constructs a [NSScrubberSelectionStyle] from an objc.ID.
//
// An abstract class that provides decorative accessory views for selected and
// highlighted items within a scrubber control.
func NSScrubberSelectionStyleFromID(id objc.ID) NSScrubberSelectionStyle {
	return NSScrubberSelectionStyle{objectivec.Object{id}}
}
// NOTE: NSScrubberSelectionStyle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSScrubberSelectionStyle] class.
//
// # Creating a selection style
//
//   - [INSScrubberSelectionStyle.InitWithCoder]: Initializes a scrubber selection style when included from a nib or Storyboard.
//   - [INSScrubberSelectionStyle.MakeSelectionView]: Provides an opportunity to create a customized scrubber selection style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle
type INSScrubberSelectionStyle interface {
	objectivec.IObject

	// Topic: Creating a selection style

	// Initializes a scrubber selection style when included from a nib or Storyboard.
	InitWithCoder(coder foundation.INSCoder) NSScrubberSelectionStyle
	// Provides an opportunity to create a customized scrubber selection style.
	MakeSelectionView() INSScrubberSelectionView

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSScrubberSelectionStyle) Init() NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberSelectionStyle) Autorelease() NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberSelectionStyle creates a new NSScrubberSelectionStyle instance.
func NewNSScrubberSelectionStyle() NSScrubberSelectionStyle {
	class := getNSScrubberSelectionStyleClass()
	rv := objc.Send[NSScrubberSelectionStyle](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a scrubber selection style when included from a nib or
// Storyboard.
//
// # Return Value
// 
// A properly initialized scrubber selection style object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/init(coder:)
func NewScrubberSelectionStyleWithCoder(coder foundation.INSCoder) NSScrubberSelectionStyle {
	instance := getNSScrubberSelectionStyleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberSelectionStyleFromID(rv)
}







// Initializes a scrubber selection style when included from a nib or
// Storyboard.
//
// # Return Value
// 
// A properly initialized scrubber selection style object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/init(coder:)
func (s NSScrubberSelectionStyle) InitWithCoder(coder foundation.INSCoder) NSScrubberSelectionStyle {
	rv := objc.Send[NSScrubberSelectionStyle](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Provides an opportunity to create a customized scrubber selection style.
//
// # Return Value
// 
// A correctly configured scrubber selection view that represents the
// appearance of your custom selection style.
//
// # Discussion
// 
// In an [NSScrubberSelectionStyle] subclass that you create, override this
// method to create a custom selection style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/makeSelectionView()
func (s NSScrubberSelectionStyle) MakeSelectionView() INSScrubberSelectionView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeSelectionView"))
	return NSScrubberSelectionViewFromID(rv)
}
func (s NSScrubberSelectionStyle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
















// A built-in selection style that draws the outline of the scrubber item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (_NSScrubberSelectionStyleClass NSScrubberSelectionStyleClass) OutlineOverlayStyle() NSScrubberSelectionStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSScrubberSelectionStyleClass.class), objc.Sel("outlineOverlayStyle"))
	return NSScrubberSelectionStyleFromID(objc.ID(rv))
}



// A built-in selection style that draws a rounded rectangle as the background
// of the scrubber item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/roundedBackground
func (_NSScrubberSelectionStyleClass NSScrubberSelectionStyleClass) RoundedBackgroundStyle() NSScrubberSelectionStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSScrubberSelectionStyleClass.class), objc.Sel("roundedBackgroundStyle"))
	return NSScrubberSelectionStyleFromID(objc.ID(rv))
}





















