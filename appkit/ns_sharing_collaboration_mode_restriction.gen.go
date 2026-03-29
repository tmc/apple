// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSharingCollaborationModeRestriction] class.
var (
	_NSSharingCollaborationModeRestrictionClass     NSSharingCollaborationModeRestrictionClass
	_NSSharingCollaborationModeRestrictionClassOnce sync.Once
)

func getNSSharingCollaborationModeRestrictionClass() NSSharingCollaborationModeRestrictionClass {
	_NSSharingCollaborationModeRestrictionClassOnce.Do(func() {
		_NSSharingCollaborationModeRestrictionClass = NSSharingCollaborationModeRestrictionClass{class: objc.GetClass("NSSharingCollaborationModeRestriction")}
	})
	return _NSSharingCollaborationModeRestrictionClass
}

// GetNSSharingCollaborationModeRestrictionClass returns the class object for NSSharingCollaborationModeRestriction.
func GetNSSharingCollaborationModeRestrictionClass() NSSharingCollaborationModeRestrictionClass {
	return getNSSharingCollaborationModeRestrictionClass()
}

type NSSharingCollaborationModeRestrictionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSharingCollaborationModeRestrictionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSharingCollaborationModeRestrictionClass) Alloc() NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

//
// # Initializers
//
//   - [NSSharingCollaborationModeRestriction.InitWithDisabledMode]
//   - [NSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessage]
//   - [NSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle]
//   - [NSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL]
//
// # Instance Properties
//
//   - [NSSharingCollaborationModeRestriction.AlertDismissButtonTitle]: The label on the alert button which will simply confirm that the alert was viewed and dismiss it Defaults to “OK”
//   - [NSSharingCollaborationModeRestriction.AlertMessage]: The message of the alert if a reason for disabling is provided
//   - [NSSharingCollaborationModeRestriction.AlertRecoverySuggestionButtonLaunchURL]: The URL that is opened when the user selects the recovery suggestion, if any
//   - [NSSharingCollaborationModeRestriction.AlertRecoverySuggestionButtonTitle]: The label on the recovery suggestion button if it is provided
//   - [NSSharingCollaborationModeRestriction.AlertTitle]: The title of the alert if a reason for disabling is provided
//   - [NSSharingCollaborationModeRestriction.DisabledMode]: The type of sharing which should be disabled
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction
type NSSharingCollaborationModeRestriction struct {
	objectivec.Object
}

// NSSharingCollaborationModeRestrictionFromID constructs a [NSSharingCollaborationModeRestriction] from an objc.ID.
func NSSharingCollaborationModeRestrictionFromID(id objc.ID) NSSharingCollaborationModeRestriction {
	return NSSharingCollaborationModeRestriction{objectivec.Object{ID: id}}
}
// NOTE: NSSharingCollaborationModeRestriction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSharingCollaborationModeRestriction] class.
//
// # Initializers
//
//   - [INSSharingCollaborationModeRestriction.InitWithDisabledMode]
//   - [INSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessage]
//   - [INSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle]
//   - [INSSharingCollaborationModeRestriction.InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL]
//
// # Instance Properties
//
//   - [INSSharingCollaborationModeRestriction.AlertDismissButtonTitle]: The label on the alert button which will simply confirm that the alert was viewed and dismiss it Defaults to “OK”
//   - [INSSharingCollaborationModeRestriction.AlertMessage]: The message of the alert if a reason for disabling is provided
//   - [INSSharingCollaborationModeRestriction.AlertRecoverySuggestionButtonLaunchURL]: The URL that is opened when the user selects the recovery suggestion, if any
//   - [INSSharingCollaborationModeRestriction.AlertRecoverySuggestionButtonTitle]: The label on the recovery suggestion button if it is provided
//   - [INSSharingCollaborationModeRestriction.AlertTitle]: The title of the alert if a reason for disabling is provided
//   - [INSSharingCollaborationModeRestriction.DisabledMode]: The type of sharing which should be disabled
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction
type INSSharingCollaborationModeRestriction interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithDisabledMode(disabledMode NSSharingCollaborationMode) NSSharingCollaborationModeRestriction
	InitWithDisabledModeAlertTitleAlertMessage(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string) NSSharingCollaborationModeRestriction
	InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string) NSSharingCollaborationModeRestriction
	InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string, alertRecoverySuggestionButtonTitle string, alertRecoverySuggestionButtonLaunchURL foundation.INSURL) NSSharingCollaborationModeRestriction

	// Topic: Instance Properties

	// The label on the alert button which will simply confirm that the alert was viewed and dismiss it Defaults to “OK”
	AlertDismissButtonTitle() string
	// The message of the alert if a reason for disabling is provided
	AlertMessage() string
	// The URL that is opened when the user selects the recovery suggestion, if any
	AlertRecoverySuggestionButtonLaunchURL() foundation.INSURL
	// The label on the recovery suggestion button if it is provided
	AlertRecoverySuggestionButtonTitle() string
	// The title of the alert if a reason for disabling is provided
	AlertTitle() string
	// The type of sharing which should be disabled
	DisabledMode() NSSharingCollaborationMode

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSharingCollaborationModeRestriction) Init() NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSharingCollaborationModeRestriction) Autorelease() NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSharingCollaborationModeRestriction creates a new NSSharingCollaborationModeRestriction instance.
func NewNSSharingCollaborationModeRestriction() NSSharingCollaborationModeRestriction {
	class := getNSSharingCollaborationModeRestrictionClass()
	rv := objc.Send[NSSharingCollaborationModeRestriction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// disabledMode: The disabled type of sharing
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:)
func NewSharingCollaborationModeRestrictionWithDisabledMode(disabledMode NSSharingCollaborationMode) NSSharingCollaborationModeRestriction {
	instance := getNSSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisabledMode:"), disabledMode)
	return NSSharingCollaborationModeRestrictionFromID(rv)
}

//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessage(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string) NSSharingCollaborationModeRestriction {
	instance := getNSSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage))
	return NSSharingCollaborationModeRestrictionFromID(rv)
}

//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// alertDismissButtonTitle: The label on the default alert button
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string) NSSharingCollaborationModeRestriction {
	instance := getNSSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage), objc.String(alertDismissButtonTitle))
	return NSSharingCollaborationModeRestrictionFromID(rv)
}

//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// alertDismissButtonTitle: The label on the default alert button
//
// alertRecoverySuggestionButtonTitle: The label on the optional recovery suggestion button on the alert
//
// alertRecoverySuggestionButtonLaunchURL: The URL that is opened when the optional recovery suggestion button is
// selected
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunch:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string, alertRecoverySuggestionButtonTitle string, alertRecoverySuggestionButtonLaunchURL foundation.INSURL) NSSharingCollaborationModeRestriction {
	instance := getNSSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunchURL:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage), objc.String(alertDismissButtonTitle), objc.String(alertRecoverySuggestionButtonTitle), alertRecoverySuggestionButtonLaunchURL)
	return NSSharingCollaborationModeRestrictionFromID(rv)
}

//
// disabledMode: The disabled type of sharing
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:)
func (s NSSharingCollaborationModeRestriction) InitWithDisabledMode(disabledMode NSSharingCollaborationMode) NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("initWithDisabledMode:"), disabledMode)
	return rv
}
//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:)
func (s NSSharingCollaborationModeRestriction) InitWithDisabledModeAlertTitleAlertMessage(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string) NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage))
	return rv
}
//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// alertDismissButtonTitle: The label on the default alert button
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:)
func (s NSSharingCollaborationModeRestriction) InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string) NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage), objc.String(alertDismissButtonTitle))
	return rv
}
//
// disabledMode: The disabled type of sharing
//
// alertTitle: The alert title
//
// alertMessage: The alert message
//
// alertDismissButtonTitle: The label on the default alert button
//
// alertRecoverySuggestionButtonTitle: The label on the optional recovery suggestion button on the alert
//
// alertRecoverySuggestionButtonLaunchURL: The URL that is opened when the optional recovery suggestion button is
// selected
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunch:)
func (s NSSharingCollaborationModeRestriction) InitWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL(disabledMode NSSharingCollaborationMode, alertTitle string, alertMessage string, alertDismissButtonTitle string, alertRecoverySuggestionButtonTitle string, alertRecoverySuggestionButtonLaunchURL foundation.INSURL) NSSharingCollaborationModeRestriction {
	rv := objc.Send[NSSharingCollaborationModeRestriction](s.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunchURL:"), disabledMode, objc.String(alertTitle), objc.String(alertMessage), objc.String(alertDismissButtonTitle), objc.String(alertRecoverySuggestionButtonTitle), alertRecoverySuggestionButtonLaunchURL)
	return rv
}
func (s NSSharingCollaborationModeRestriction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The label on the alert button which will simply confirm that the alert was
// viewed and dismiss it Defaults to “OK”
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertDismissButtonTitle
func (s NSSharingCollaborationModeRestriction) AlertDismissButtonTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alertDismissButtonTitle"))
	return foundation.NSStringFromID(rv).String()
}
// The message of the alert if a reason for disabling is provided
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertMessage
func (s NSSharingCollaborationModeRestriction) AlertMessage() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alertMessage"))
	return foundation.NSStringFromID(rv).String()
}
// The URL that is opened when the user selects the recovery suggestion, if
// any
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertRecoverySuggestionButtonLaunchURL
func (s NSSharingCollaborationModeRestriction) AlertRecoverySuggestionButtonLaunchURL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alertRecoverySuggestionButtonLaunchURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The label on the recovery suggestion button if it is provided
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertRecoverySuggestionButtonTitle
func (s NSSharingCollaborationModeRestriction) AlertRecoverySuggestionButtonTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alertRecoverySuggestionButtonTitle"))
	return foundation.NSStringFromID(rv).String()
}
// The title of the alert if a reason for disabling is provided
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertTitle
func (s NSSharingCollaborationModeRestriction) AlertTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alertTitle"))
	return foundation.NSStringFromID(rv).String()
}
// The type of sharing which should be disabled
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/disabledMode
func (s NSSharingCollaborationModeRestriction) DisabledMode() NSSharingCollaborationMode {
	rv := objc.Send[NSSharingCollaborationMode](s.ID, objc.Sel("disabledMode"))
	return NSSharingCollaborationMode(rv)
}

