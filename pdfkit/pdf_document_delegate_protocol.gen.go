// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The delegate for the [PDFDocument] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate
type PDFDocumentDelegate interface {
	objectivec.IObject
}

// PDFDocumentDelegateObject wraps an existing Objective-C object that conforms to the PDFDocumentDelegate protocol.
type PDFDocumentDelegateObject struct {
	objectivec.Object
}

func (o PDFDocumentDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDFDocumentDelegateObjectFromID constructs a [PDFDocumentDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDFDocumentDelegateObjectFromID(id objc.ID) PDFDocumentDelegateObject {
	return PDFDocumentDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when the [PDFDocumentDidUnlockNotification] notification is posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidUnlock(_:)
func (o PDFDocumentDelegateObject) DocumentDidUnlock(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidUnlock:"), notification)
}

// Called for every match found during a find operation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/didMatchString(_:)
func (o PDFDocumentDelegateObject) DidMatchString(instance IPDFSelection) {
	objc.Send[struct{}](o.ID, objc.Sel("didMatchString:"), instance)
}

// Called when the [PDFDocumentDidBeginFindNotification] notification is
// posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidBeginDocumentFind(_:)
func (o PDFDocumentDelegateObject) DocumentDidBeginDocumentFind(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidBeginDocumentFind:"), notification)
}

// Called when the [PDFDocumentDidBeginPageFindNotification] notification is
// posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidBeginPageFind(_:)
func (o PDFDocumentDelegateObject) DocumentDidBeginPageFind(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidBeginPageFind:"), notification)
}

// Called when the [PDFDocumentDidEndFindNotification] notification is posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidEndDocumentFind(_:)
func (o PDFDocumentDelegateObject) DocumentDidEndDocumentFind(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidEndDocumentFind:"), notification)
}

// Called when the [PDFDocumentDidEndPageFindNotification] notification is
// posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidEndPageFind(_:)
func (o PDFDocumentDelegateObject) DocumentDidEndPageFind(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidEndPageFind:"), notification)
}

// Called when the [PDFDocumentDidFindMatchNotification] notification is
// posted.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/documentDidFindMatch(_:)
func (o PDFDocumentDelegateObject) DocumentDidFindMatch(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("documentDidFindMatch:"), notification)
}

// Returns a [PDFPage] subclass for a page object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/classForPage()
func (o PDFDocumentDelegateObject) ClassForPage() objc.Class {
	rv := objc.Send[objc.Class](o.ID, objc.Sel("classForPage"))
	return rv
}

// Returns a [PDFAnnotation] subclass for an annotation type.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDelegate/class(forAnnotationType:)
func (o PDFDocumentDelegateObject) ClassForAnnotationType(annotationType string) objc.Class {
	rv := objc.Send[objc.Class](o.ID, objc.Sel("classForAnnotationType:"), objc.String(annotationType))
	return rv
}

// PDFDocumentDelegateConfig holds optional typed callbacks for [PDFDocumentDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/pdfkit/pdfdocumentdelegate
type PDFDocumentDelegateConfig struct {

	// Managing Document Security
	// DocumentDidUnlock — Called when the [PDFDocumentDidUnlockNotification] notification is posted.
	DocumentDidUnlock func(notification foundation.NSNotification)

	// Getting Document Search Notifications
	// DidMatchString — Called for every match found during a find operation.
	DidMatchString func(instance PDFSelection)
	// DocumentDidBeginDocumentFind — Called when the [PDFDocumentDidBeginFindNotification] notification is posted.
	DocumentDidBeginDocumentFind func(notification foundation.NSNotification)
	// DocumentDidBeginPageFind — Called when the [PDFDocumentDidBeginPageFindNotification] notification is posted.
	DocumentDidBeginPageFind func(notification foundation.NSNotification)
	// DocumentDidEndDocumentFind — Called when the [PDFDocumentDidEndFindNotification] notification is posted.
	DocumentDidEndDocumentFind func(notification foundation.NSNotification)
	// DocumentDidEndPageFind — Called when the [PDFDocumentDidEndPageFindNotification] notification is posted.
	DocumentDidEndPageFind func(notification foundation.NSNotification)
	// DocumentDidFindMatch — Called when the [PDFDocumentDidFindMatchNotification] notification is posted.
	DocumentDidFindMatch func(notification foundation.NSNotification)

	// Wrapping Document Elements
	// ClassForPage — Returns a [PDFPage] subclass for a page object.
	ClassForPage func() objc.Class
}

// NewPDFDocumentDelegate creates an Objective-C object implementing the [PDFDocumentDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [PDFDocumentDelegateObject] satisfies the [PDFDocumentDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/pdfkit/pdfdocumentdelegate
func NewPDFDocumentDelegate(config PDFDocumentDelegateConfig) PDFDocumentDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoPDFDocumentDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DocumentDidUnlock != nil {
		fn := config.DocumentDidUnlock
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidUnlock:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidMatchString != nil {
		fn := config.DidMatchString
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didMatchString:"),
			Fn: func(self objc.ID, _cmd objc.SEL, instanceID objc.ID) {
				instance := PDFSelectionFromID(instanceID)
				fn(instance)
			},
		})
	}

	if config.DocumentDidBeginDocumentFind != nil {
		fn := config.DocumentDidBeginDocumentFind
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidBeginDocumentFind:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DocumentDidBeginPageFind != nil {
		fn := config.DocumentDidBeginPageFind
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidBeginPageFind:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DocumentDidEndDocumentFind != nil {
		fn := config.DocumentDidEndDocumentFind
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidEndDocumentFind:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DocumentDidEndPageFind != nil {
		fn := config.DocumentDidEndPageFind
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidEndPageFind:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DocumentDidFindMatch != nil {
		fn := config.DocumentDidFindMatch
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("documentDidFindMatch:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ClassForPage != nil {
		fn := config.ClassForPage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("classForPage"),
			Fn: func(self objc.ID, _cmd objc.SEL) objc.Class {
				return fn()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("PDFDocumentDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewPDFDocumentDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return PDFDocumentDelegateObjectFromID(instance)
}
