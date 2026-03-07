// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of methods that can be implemented by the delegate of a path control object to support dragging to and from the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate
type NSPathControlDelegate interface {
	objectivec.IObject
}



// NSPathControlDelegateObject wraps an existing Objective-C object that conforms to the NSPathControlDelegate protocol.
type NSPathControlDelegateObject struct {
	objectivec.Object
}
func (o NSPathControlDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSPathControlDelegateObjectFromID constructs a [NSPathControlDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPathControlDelegateObjectFromID(id objc.ID) NSPathControlDelegateObject {
	return NSPathControlDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Implement this method to enable dragging from the control.
//
// pathControl: The path control that sent the message.
//
// pathComponentCell: The path component cell from which the drag is beginning.
//
// pasteboard: The pasteboard.
//
// # Discussion
// 
// This method is called when a drag is about to begin. You can refuse to
// allow the drag to happen by returning [false] and allow it by returning
// [true]. By default, the pasteboard automatically has the following types on
// it: [NSStringPboardType], [NSURLPboardType] (if there is a URL value for
// the cell being dragged), and [NSFilenamesPboardType] (if the URL value
// returns [true] from -[isFileURL]). You can customize the types placed on
// the pasteboard at this time, if desired. Implementation of this method is
// optional.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isFileURL]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:shouldDrag:with:)-35j1e

func (o NSPathControlDelegateObject) PathControlShouldDragPathComponentCellWithPasteboard(pathControl INSPathControl, pathComponentCell INSPathComponentCell, pasteboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("pathControl:shouldDragPathComponentCell:withPasteboard:"), pathControl, pathComponentCell, pasteboard)
	return rv
	}

// Implement this method to enable dragging onto the control.
//
// pathControl: The path control that sent the message.
//
// info: An object containing details about this dragging operation.
//
// # Discussion
// 
// This method is called when something is dragged over the control. Return
// [NSDragOperationNone] to refuse the drop, or return anything else to accept
// it.
// 
// If not implemented, and the control’s cell is editable, the drop is
// accepted if it contains an [NSURLPboardType] or [NSFilenamesPboardType]
// that conforms to the cell’s allowed types. Implementation of this method
// is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:validateDrop:)

func (o NSPathControlDelegateObject) PathControlValidateDrop(pathControl INSPathControl, info NSDraggingInfo) NSDragOperation {
	
	rv := objc.Send[NSDragOperation](o.ID, objc.Sel("pathControl:validateDrop:"), pathControl, info)
	return rv
	}

// Implement this method to accept previously validated contents dropped onto
// the control.
//
// pathControl: The path control that sent the message.
//
// info: An object containing details about this dragging operation.
//
// # Discussion
// 
// In order to accept the dropped contents previously accepted from
// [PathControlValidateDrop], you must implement this method. This method is
// called from performDragOperation:. You should change the URL value based on
// the dragged information.
// 
// If not implemented, and the control’s cell is editable, the drop is
// accepted if it contains an [NSURLPboardType] or [NSFilenamesPboardType]
// that conforms to the cell’s allowed types. The cell’s URL value is
// automatically changed, and the action is invoked. Implementation of this
// method is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:acceptDrop:)

func (o NSPathControlDelegateObject) PathControlAcceptDrop(pathControl INSPathControl, info NSDraggingInfo) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("pathControl:acceptDrop:"), pathControl, info)
	return rv
	}

// Implement this method to customize the Open panel shown by a pop-up–style
// path.
//
// pathControl: The path control displaying the Open panel.
//
// openPanel: The Open panel to be displayed.
//
// # Discussion
// 
// This method is called before the Open panel is shown but after its allowed
// file types are set to the cell’s allowed types. At this time, you can
// further customize the Open panel as required. This method is called only
// when the style is set to [NSPathStylePopUp]. Implementation of this method
// is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:willDisplay:)

func (o NSPathControlDelegateObject) PathControlWillDisplayOpenPanel(pathControl INSPathControl, openPanel INSOpenPanel) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pathControl:willDisplayOpenPanel:"), pathControl, openPanel)
	}

// Implement this method to customize the menu of a pop-up–style path.
//
// pathControl: The path control displaying the pop-up menu.
//
// menu: The pop-up menu to be displayed.
//
// # Discussion
// 
// This method is called before the pop-up menu is shown. At this time, you
// can further customize the menu as required, adding and removing items. This
// method is called only when the style is set to [NSPathStylePopUp].
// Implementation of this method is optional.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:willPopUp:)

func (o NSPathControlDelegateObject) PathControlWillPopUpMenu(pathControl INSPathControl, menu INSMenu) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pathControl:willPopUpMenu:"), pathControl, menu)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlDelegate/pathControl(_:shouldDrag:with:)-5ciyd

func (o NSPathControlDelegateObject) PathControlShouldDragItemWithPasteboard(pathControl INSPathControl, pathItem INSPathControlItem, pasteboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("pathControl:shouldDragItem:withPasteboard:"), pathControl, pathItem, pasteboard)
	return rv
	}





// NSPathControlDelegateConfig holds optional typed callbacks for [NSPathControlDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate
type NSPathControlDelegateConfig struct {

	// Other Methods
	// ShouldDragPathComponentCellWithPasteboard — Implement this method to enable dragging from the control.
	ShouldDragPathComponentCellWithPasteboard func(pathControl NSPathControl, pathComponentCell NSPathComponentCell, pasteboard NSPasteboard) bool
	// WillDisplayOpenPanel — Implement this method to customize the Open panel shown by a pop-up–style path.
	WillDisplayOpenPanel func(pathControl NSPathControl, openPanel NSOpenPanel)
	// WillPopUpMenu — Implement this method to customize the menu of a pop-up–style path.
	WillPopUpMenu func(pathControl NSPathControl, menu NSMenu)
	ShouldDragItemWithPasteboard func(pathControl NSPathControl, pathItem NSPathControlItem, pasteboard NSPasteboard) bool
}

// NewNSPathControlDelegate creates an Objective-C object implementing the [NSPathControlDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSPathControlDelegateObject] satisfies the [NSPathControlDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate
func NewNSPathControlDelegate(config NSPathControlDelegateConfig) NSPathControlDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSPathControlDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldDragPathComponentCellWithPasteboard != nil {
		fn := config.ShouldDragPathComponentCellWithPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathControl:shouldDragPathComponentCell:withPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathControlID objc.ID, pathComponentCellID objc.ID, pasteboardID objc.ID) bool {
				pathControl := NSPathControlFromID(pathControlID)
				pathComponentCell := NSPathComponentCellFromID(pathComponentCellID)
				pasteboard := NSPasteboardFromID(pasteboardID)
				return fn(pathControl, pathComponentCell, pasteboard)
			},
		})
	}

	if config.WillDisplayOpenPanel != nil {
		fn := config.WillDisplayOpenPanel
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathControl:willDisplayOpenPanel:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathControlID objc.ID, openPanelID objc.ID) {
				pathControl := NSPathControlFromID(pathControlID)
				openPanel := NSOpenPanelFromID(openPanelID)
				fn(pathControl, openPanel)
			},
		})
	}

	if config.WillPopUpMenu != nil {
		fn := config.WillPopUpMenu
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathControl:willPopUpMenu:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathControlID objc.ID, menuID objc.ID) {
				pathControl := NSPathControlFromID(pathControlID)
				menu := NSMenuFromID(menuID)
				fn(pathControl, menu)
			},
		})
	}

	if config.ShouldDragItemWithPasteboard != nil {
		fn := config.ShouldDragItemWithPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pathControl:shouldDragItem:withPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pathControlID objc.ID, pathItemID objc.ID, pasteboardID objc.ID) bool {
				pathControl := NSPathControlFromID(pathControlID)
				pathItem := NSPathControlItemFromID(pathItemID)
				pasteboard := NSPasteboardFromID(pasteboardID)
				return fn(pathControl, pathItem, pasteboard)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSPathControlDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSPathControlDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSPathControlDelegateObjectFromID(instance)
}





