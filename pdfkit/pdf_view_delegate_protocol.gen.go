// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The delegate for the [PDFView] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate
type PDFViewDelegate interface {
	objectivec.IObject
}

// PDFViewDelegateObject wraps an existing Objective-C object that conforms to the PDFViewDelegate protocol.
type PDFViewDelegateObject struct {
	objectivec.Object
}

func (o PDFViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDFViewDelegateObjectFromID constructs a [PDFViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDFViewDelegateObjectFromID(id objc.ID) PDFViewDelegateObject {
	return PDFViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Performs a find operation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewPerformFind(_:)
func (o PDFViewDelegateObject) PDFViewPerformFind(sender IPDFView) {
	objc.Send[struct{}](o.ID, objc.Sel("PDFViewPerformFind:"), sender)
}

// Performs a go-to operation.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewPerformGo(toPage:)
func (o PDFViewDelegateObject) PDFViewPerformGoToPage(sender IPDFView) {
	objc.Send[struct{}](o.ID, objc.Sel("PDFViewPerformGoToPage:"), sender)
}

// Prints the current document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewPerformPrint(_:)
func (o PDFViewDelegateObject) PDFViewPerformPrint(sender IPDFView) {
	objc.Send[struct{}](o.ID, objc.Sel("PDFViewPerformPrint:"), sender)
}

// Opens a specified page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewOpenPDF(_:forRemoteGoToAction:)
func (o PDFViewDelegateObject) PDFViewOpenPDFForRemoteGoToAction(sender IPDFView, action IPDFActionRemoteGoTo) {
	objc.Send[struct{}](o.ID, objc.Sel("PDFViewOpenPDF:forRemoteGoToAction:"), sender, action)
}

// Overrides changes to the scale factor.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewWillChangeScaleFactor(_:toScale:)
func (o PDFViewDelegateObject) PDFViewWillChangeScaleFactorToScale(sender IPDFView, scaler float64) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("PDFViewWillChangeScaleFactor:toScale:"), sender, scaler)
	return rv
}

// Handle clicks on URL links in a view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewWillClick(onLink:with:)
func (o PDFViewDelegateObject) PDFViewWillClickOnLinkWithURL(sender IPDFView, url foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("PDFViewWillClickOnLink:withURL:"), sender, url)
}

// Overrides the job title used when the [PDFView] is printed.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewPrintJobTitle(_:)
func (o PDFViewDelegateObject) PDFViewPrintJobTitle(sender IPDFView) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("PDFViewPrintJobTitle:"), sender)
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/PDFKit/PDFViewDelegate/pdfViewParentViewController()
func (o PDFViewDelegateObject) PDFViewParentViewController() objc.ID {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("PDFViewParentViewController"))
	return rv
}

// PDFViewDelegateConfig holds optional typed callbacks for [PDFViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/pdfkit/pdfviewdelegate
type PDFViewDelegateConfig struct {

	// Other Methods
	// PDFViewPerformFind — Performs a find operation.
	PDFViewPerformFind func(sender PDFView)
	// PDFViewPerformGoToPage — Performs a go-to operation.
	PDFViewPerformGoToPage func(sender PDFView)
	// PDFViewPerformPrint — Prints the current document.
	PDFViewPerformPrint func(sender PDFView)
	// PDFViewOpenPDFForRemoteGoToAction — Opens a specified page.
	PDFViewOpenPDFForRemoteGoToAction func(sender PDFView, action PDFActionRemoteGoTo)
	// PDFViewWillChangeScaleFactorToScale — Overrides changes to the scale factor.
	PDFViewWillChangeScaleFactorToScale func(sender PDFView, scaler float64) float64
	// PDFViewWillClickOnLinkWithURL — Handle clicks on URL links in a view.
	PDFViewWillClickOnLinkWithURL func(sender PDFView, url foundation.NSURL)
}

// NewPDFViewDelegate creates an Objective-C object implementing the [PDFViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [PDFViewDelegateObject] satisfies the [PDFViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/pdfkit/pdfviewdelegate
func NewPDFViewDelegate(config PDFViewDelegateConfig) PDFViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoPDFViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.PDFViewPerformFind != nil {
		fn := config.PDFViewPerformFind
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewPerformFind:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := PDFViewFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.PDFViewPerformGoToPage != nil {
		fn := config.PDFViewPerformGoToPage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewPerformGoToPage:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := PDFViewFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.PDFViewPerformPrint != nil {
		fn := config.PDFViewPerformPrint
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewPerformPrint:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := PDFViewFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.PDFViewOpenPDFForRemoteGoToAction != nil {
		fn := config.PDFViewOpenPDFForRemoteGoToAction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewOpenPDF:forRemoteGoToAction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, actionID objc.ID) {
				sender := PDFViewFromID(senderID)
				action := PDFActionRemoteGoToFromID(actionID)
				fn(sender, action)
			},
		})
	}

	if config.PDFViewWillChangeScaleFactorToScale != nil {
		fn := config.PDFViewWillChangeScaleFactorToScale
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewWillChangeScaleFactor:toScale:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, scaler float64) float64 {
				sender := PDFViewFromID(senderID)
				return fn(sender, scaler)
			},
		})
	}

	if config.PDFViewWillClickOnLinkWithURL != nil {
		fn := config.PDFViewWillClickOnLinkWithURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("PDFViewWillClickOnLink:withURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, urlID objc.ID) {
				sender := PDFViewFromID(senderID)
				url := foundation.NSURLFromID(urlID)
				fn(sender, url)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("PDFViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewPDFViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return PDFViewDelegateObjectFromID(instance)
}
