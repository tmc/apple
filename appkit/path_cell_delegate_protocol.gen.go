// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of methods that enable the delegate of a path cell object to customize the Open panel or pop-up menu of a path whose style is set to [NSPathControl.Style.popUp](<doc://com.apple.appkit/documentation/AppKit/NSPathControl/Style/popUp>).
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCellDelegate
type NSPathCellDelegate interface {
	objectivec.IObject
}

// NSPathCellDelegateObject wraps an existing Objective-C object that conforms to the NSPathCellDelegate protocol.
type NSPathCellDelegateObject struct {
	objectivec.Object
}
func (o NSPathCellDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPathCellDelegateObjectFromID constructs a [NSPathCellDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPathCellDelegateObjectFromID(id objc.ID) NSPathCellDelegateObject {
	return NSPathCellDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Implement this method to customize the Open panel shown by a pop-up–style
// path.
//
// pathCell: The path cell that sent the message.
//
// openPanel: The Open panel to be displayed.
//
// # Discussion
// 
// This method is called before the Open panel is shown but after its allowed
// file types are set to the cell’s allowed types. At this time, you can
// further customize the Open panel as required. This method is called only
// when the style is set to [NSPathStylePopUp].
// 
// Implementation of this method is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCellDelegate/pathCell(_:willDisplay:)
func (o NSPathCellDelegateObject) PathCellWillDisplayOpenPanel(pathCell INSPathCell, openPanel INSOpenPanel) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pathCell:willDisplayOpenPanel:"), pathCell, openPanel)
	}
// Implement this method to customize the menu of a pop-up–style path.
//
// pathCell: The path cell that sent the message.
//
// menu: The pop-up menu to be displayed.
//
// # Discussion
// 
// This method is called before the pop-up menu is shown. At this time, you
// can further customize the menu as required, adding and removing items. This
// method is called only when the style is set to [NSPathStylePopUp].
// 
// Implementation of this method is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCellDelegate/pathCell(_:willPopUp:)
func (o NSPathCellDelegateObject) PathCellWillPopUpMenu(pathCell INSPathCell, menu INSMenu) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pathCell:willPopUpMenu:"), pathCell, menu)
	}

// NSPathCellDelegateConfig holds optional typed callbacks for [NSPathCellDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspathcelldelegate
type NSPathCellDelegateConfig struct {

	// Other Methods
	// PathCellWillDisplayOpenPanel — Implement this method to customize the Open panel shown by a pop-up–style path.
	PathCellWillDisplayOpenPanel func(pathCell NSPathCell, openPanel NSOpenPanel)
	// PathCellWillPopUpMenu — Implement this method to customize the menu of a pop-up–style path.
	PathCellWillPopUpMenu func(pathCell NSPathCell, menu NSMenu)
}

// NewNSPathCellDelegate creates an Objective-C object implementing the [NSPathCellDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSPathCellDelegateObject] satisfies the [NSPathCellDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspathcelldelegate
func NewNSPathCellDelegate(config NSPathCellDelegateConfig) NSPathCellDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSPathCellDelegate_%d", n)

	var methods []objc.MethodDef

	if config.PathCellWillDisplayOpenPanel != nil {
		fn := config.PathCellWillDisplayOpenPanel
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathCell:willDisplayOpenPanel:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathCellID objc.ID, openPanelID objc.ID) {
				pathCell := NSPathCellFromID(pathCellID)
				openPanel := NSOpenPanelFromID(openPanelID)
				fn(pathCell, openPanel)
			},
		})
	}

	if config.PathCellWillPopUpMenu != nil {
		fn := config.PathCellWillPopUpMenu
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathCell:willPopUpMenu:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathCellID objc.ID, menuID objc.ID) {
				pathCell := NSPathCellFromID(pathCellID)
				menu := NSMenuFromID(menuID)
				fn(pathCell, menu)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSPathCellDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSPathCellDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSPathCellDelegateObjectFromID(instance)
}

