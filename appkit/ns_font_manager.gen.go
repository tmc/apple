// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFontManager] class.
var (
	_NSFontManagerClass     NSFontManagerClass
	_NSFontManagerClassOnce sync.Once
)

func getNSFontManagerClass() NSFontManagerClass {
	_NSFontManagerClassOnce.Do(func() {
		_NSFontManagerClass = NSFontManagerClass{class: objc.GetClass("NSFontManager")}
	})
	return _NSFontManagerClass
}

// GetNSFontManagerClass returns the class object for NSFontManager.
func GetNSFontManagerClass() NSFontManagerClass {
	return getNSFontManagerClass()
}

type NSFontManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontManagerClass) Alloc() NSFontManager {
	rv := objc.Send[NSFontManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The center of activity for the font-conversion system.
//
// # Overview
// 
// The font manager records the currently selected font, updates the Font
// panel and Font menu to reflect the selected font, initiates font changes,
// and converts fonts in response to requests from text-bearing objects. In a
// more prosaic role, [NSFontManager] can be queried for the fonts available
// to the application and for the particular attributes of a font, such as
// whether it’s condensed or extended.
// 
// You typically set up a font manager and the Font menu using Interface
// Builder. However, you can also do so programmatically by getting the shared
// font manager instance and having it create the standard Font menu at
// runtime:
// 
// You can then add the Font menu to your app’s main menu. After the Font
// menu is installed, your app automatically gains the functionality of both
// the Font menu and the Font panel.
// 
// Font collections are managed by [NSFontManager].
//
// # Getting Available Fonts
//
//   - [NSFontManager.AvailableFonts]: The names of the fonts available in the system (not the [NSFont](<doc://com.apple.appkit/documentation/AppKit/NSFont>) objects themselves).
//   - [NSFontManager.AvailableFontFamilies]: The names of the font families available in the system.
//   - [NSFontManager.AvailableFontNamesWithTraits]: Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the [NSFont] objects themselves).
//   - [NSFontManager.AvailableMembersOfFontFamily]: Returns an array with one entry for each available member of a font family.
//
// # Setting and Examining the Selected Font
//
//   - [NSFontManager.SetSelectedFontIsMultiple]: Records the specified font as the currently selected font and updates the Font panel.
//   - [NSFontManager.SelectedFont]: The currently selected font object.
//   - [NSFontManager.Multiple]: A Boolean value that indicates whether the currently selected font has multiple fonts.
//   - [NSFontManager.SendAction]: A Boolean value that indicates whether a responder handled the font manager’s action message.
//   - [NSFontManager.LocalizedNameForFamilyFace]: Returns a localized string with the name of the specified font family and face, if one exists.
//
// # Sending Action Methods
//
//   - [NSFontManager.AddFontTrait]: Adds a trait to the font.
//   - [NSFontManager.RemoveFontTrait]: Removes a trait from the font.
//   - [NSFontManager.ModifyFont]: Modifies a trait of the font.
//   - [NSFontManager.ModifyFontViaPanel]: Modifies a font trait using input from the Font panel.
//   - [NSFontManager.OrderFrontStylesPanel]: Opens the Font Styles panel.
//   - [NSFontManager.OrderFrontFontPanel]: Opens the Font panel, creating it if necessary, and displays that panel in front of the app’s windows.
//
// # Converting Fonts Automatically
//
//   - [NSFontManager.ConvertFont]: Converts the given font according to the object that initiated a font change, typically the Font panel or Font menu.
//
// # Converting Fonts Manually
//
//   - [NSFontManager.ConvertFontToFace]: Returns a font whose traits are as similar as possible to those of the given font except for the typeface, which is changed to the given typeface.
//   - [NSFontManager.ConvertFontToFamily]: Returns a font whose traits are as similar as possible to those of the given font except for the font family, which is changed to the given family.
//   - [NSFontManager.ConvertFontToHaveTrait]: Returns a new version of the font object containing a single additional trait.
//   - [NSFontManager.ConvertFontToNotHaveTrait]: Returns a new version of a font object without the specified traits.
//   - [NSFontManager.ConvertFontToSize]: Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
//   - [NSFontManager.ConvertWeightOfFont]: Returns a font object whose weight is greater or lesser than that of the given font.
//   - [NSFontManager.CurrentFontAction]: The current font conversion action.
//   - [NSFontManager.ConvertFontTraits]: Converts font traits to a new traits mask value.
//
// # Getting a Particular Font
//
//   - [NSFontManager.FontWithFamilyTraitsWeightSize]: Attempts to load a font with the specified characteristics.
//
// # Examining Fonts
//
//   - [NSFontManager.TraitsOfFont]: Returns the traits of the given font.
//   - [NSFontManager.FontNamedHasTraits]: Indicates whether the given font has all the specified traits.
//   - [NSFontManager.WeightOfFont]: Returns an approximation of the specified font’s weight.
//
// # Managing the Font Panel and Font Menu
//
//   - [NSFontManager.Enabled]: A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//   - [NSFontManager.SetEnabled]
//   - [NSFontManager.FontPanel]: Returns the application’s shared Font panel object, creating it if necessary.
//   - [NSFontManager.SetFontMenu]: Records the given menu as the application’s Font menu.
//   - [NSFontManager.FontMenu]: Returns the menu that’s connected to the font conversion system, creating it if necessary.
//
// # Accessing the Action Property
//
//   - [NSFontManager.Action]: The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//   - [NSFontManager.SetAction]
//   - [NSFontManager.Target]: The object that receives action messages related to the font manager.
//   - [NSFontManager.SetTarget]
//
// # Setting Attributes
//
//   - [NSFontManager.SetSelectedAttributesIsMultiple]: Informs the Font panel that the specified font attributes changed for the selected text.
//   - [NSFontManager.ConvertAttributes]: Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager
type NSFontManager struct {
	objectivec.Object
}

// NSFontManagerFromID constructs a [NSFontManager] from an objc.ID.
//
// The center of activity for the font-conversion system.
func NSFontManagerFromID(id objc.ID) NSFontManager {
	return NSFontManager{objectivec.Object{id}}
}
// NOTE: NSFontManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSFontManager] class.
//
// # Getting Available Fonts
//
//   - [INSFontManager.AvailableFonts]: The names of the fonts available in the system (not the [NSFont](<doc://com.apple.appkit/documentation/AppKit/NSFont>) objects themselves).
//   - [INSFontManager.AvailableFontFamilies]: The names of the font families available in the system.
//   - [INSFontManager.AvailableFontNamesWithTraits]: Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the [NSFont] objects themselves).
//   - [INSFontManager.AvailableMembersOfFontFamily]: Returns an array with one entry for each available member of a font family.
//
// # Setting and Examining the Selected Font
//
//   - [INSFontManager.SetSelectedFontIsMultiple]: Records the specified font as the currently selected font and updates the Font panel.
//   - [INSFontManager.SelectedFont]: The currently selected font object.
//   - [INSFontManager.Multiple]: A Boolean value that indicates whether the currently selected font has multiple fonts.
//   - [INSFontManager.SendAction]: A Boolean value that indicates whether a responder handled the font manager’s action message.
//   - [INSFontManager.LocalizedNameForFamilyFace]: Returns a localized string with the name of the specified font family and face, if one exists.
//
// # Sending Action Methods
//
//   - [INSFontManager.AddFontTrait]: Adds a trait to the font.
//   - [INSFontManager.RemoveFontTrait]: Removes a trait from the font.
//   - [INSFontManager.ModifyFont]: Modifies a trait of the font.
//   - [INSFontManager.ModifyFontViaPanel]: Modifies a font trait using input from the Font panel.
//   - [INSFontManager.OrderFrontStylesPanel]: Opens the Font Styles panel.
//   - [INSFontManager.OrderFrontFontPanel]: Opens the Font panel, creating it if necessary, and displays that panel in front of the app’s windows.
//
// # Converting Fonts Automatically
//
//   - [INSFontManager.ConvertFont]: Converts the given font according to the object that initiated a font change, typically the Font panel or Font menu.
//
// # Converting Fonts Manually
//
//   - [INSFontManager.ConvertFontToFace]: Returns a font whose traits are as similar as possible to those of the given font except for the typeface, which is changed to the given typeface.
//   - [INSFontManager.ConvertFontToFamily]: Returns a font whose traits are as similar as possible to those of the given font except for the font family, which is changed to the given family.
//   - [INSFontManager.ConvertFontToHaveTrait]: Returns a new version of the font object containing a single additional trait.
//   - [INSFontManager.ConvertFontToNotHaveTrait]: Returns a new version of a font object without the specified traits.
//   - [INSFontManager.ConvertFontToSize]: Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
//   - [INSFontManager.ConvertWeightOfFont]: Returns a font object whose weight is greater or lesser than that of the given font.
//   - [INSFontManager.CurrentFontAction]: The current font conversion action.
//   - [INSFontManager.ConvertFontTraits]: Converts font traits to a new traits mask value.
//
// # Getting a Particular Font
//
//   - [INSFontManager.FontWithFamilyTraitsWeightSize]: Attempts to load a font with the specified characteristics.
//
// # Examining Fonts
//
//   - [INSFontManager.TraitsOfFont]: Returns the traits of the given font.
//   - [INSFontManager.FontNamedHasTraits]: Indicates whether the given font has all the specified traits.
//   - [INSFontManager.WeightOfFont]: Returns an approximation of the specified font’s weight.
//
// # Managing the Font Panel and Font Menu
//
//   - [INSFontManager.Enabled]: A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
//   - [INSFontManager.SetEnabled]
//   - [INSFontManager.FontPanel]: Returns the application’s shared Font panel object, creating it if necessary.
//   - [INSFontManager.SetFontMenu]: Records the given menu as the application’s Font menu.
//   - [INSFontManager.FontMenu]: Returns the menu that’s connected to the font conversion system, creating it if necessary.
//
// # Accessing the Action Property
//
//   - [INSFontManager.Action]: The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
//   - [INSFontManager.SetAction]
//   - [INSFontManager.Target]: The object that receives action messages related to the font manager.
//   - [INSFontManager.SetTarget]
//
// # Setting Attributes
//
//   - [INSFontManager.SetSelectedAttributesIsMultiple]: Informs the Font panel that the specified font attributes changed for the selected text.
//   - [INSFontManager.ConvertAttributes]: Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager
type INSFontManager interface {
	objectivec.IObject
	NSMenuItemValidation

	// Topic: Getting Available Fonts

	// The names of the fonts available in the system (not the [NSFont](<doc://com.apple.appkit/documentation/AppKit/NSFont>) objects themselves).
	AvailableFonts() []string
	// The names of the font families available in the system.
	AvailableFontFamilies() []string
	// Returns the names of the fonts available in the system whose traits are described exactly by the given font trait mask (not the [NSFont] objects themselves).
	AvailableFontNamesWithTraits(someTraits NSFontTraitMask) []string
	// Returns an array with one entry for each available member of a font family.
	AvailableMembersOfFontFamily(fam string) []foundation.NSArray

	// Topic: Setting and Examining the Selected Font

	// Records the specified font as the currently selected font and updates the Font panel.
	SetSelectedFontIsMultiple(fontObj NSFont, flag bool)
	// The currently selected font object.
	SelectedFont() NSFont
	// A Boolean value that indicates whether the currently selected font has multiple fonts.
	Multiple() bool
	// A Boolean value that indicates whether a responder handled the font manager’s action message.
	SendAction() bool
	// Returns a localized string with the name of the specified font family and face, if one exists.
	LocalizedNameForFamilyFace(family string, faceKey string) string

	// Topic: Sending Action Methods

	// Adds a trait to the font.
	AddFontTrait(sender objectivec.IObject)
	// Removes a trait from the font.
	RemoveFontTrait(sender objectivec.IObject)
	// Modifies a trait of the font.
	ModifyFont(sender objectivec.IObject)
	// Modifies a font trait using input from the Font panel.
	ModifyFontViaPanel(sender objectivec.IObject)
	// Opens the Font Styles panel.
	OrderFrontStylesPanel(sender objectivec.IObject)
	// Opens the Font panel, creating it if necessary, and displays that panel in front of the app’s windows.
	OrderFrontFontPanel(sender objectivec.IObject)

	// Topic: Converting Fonts Automatically

	// Converts the given font according to the object that initiated a font change, typically the Font panel or Font menu.
	ConvertFont(fontObj NSFont) NSFont

	// Topic: Converting Fonts Manually

	// Returns a font whose traits are as similar as possible to those of the given font except for the typeface, which is changed to the given typeface.
	ConvertFontToFace(fontObj NSFont, typeface string) NSFont
	// Returns a font whose traits are as similar as possible to those of the given font except for the font family, which is changed to the given family.
	ConvertFontToFamily(fontObj NSFont, family string) NSFont
	// Returns a new version of the font object containing a single additional trait.
	ConvertFontToHaveTrait(fontObj NSFont, trait NSFontTraitMask) NSFont
	// Returns a new version of a font object without the specified traits.
	ConvertFontToNotHaveTrait(fontObj NSFont, trait NSFontTraitMask) NSFont
	// Returns a font object whose traits are the same as those of the given font, except for the size, which is changed to the given size.
	ConvertFontToSize(fontObj NSFont, size float64) NSFont
	// Returns a font object whose weight is greater or lesser than that of the given font.
	ConvertWeightOfFont(upFlag bool, fontObj NSFont) NSFont
	// The current font conversion action.
	CurrentFontAction() NSFontAction
	// Converts font traits to a new traits mask value.
	ConvertFontTraits(traits NSFontTraitMask) NSFontTraitMask

	// Topic: Getting a Particular Font

	// Attempts to load a font with the specified characteristics.
	FontWithFamilyTraitsWeightSize(family string, traits NSFontTraitMask, weight int, size float64) NSFont

	// Topic: Examining Fonts

	// Returns the traits of the given font.
	TraitsOfFont(fontObj NSFont) NSFontTraitMask
	// Indicates whether the given font has all the specified traits.
	FontNamedHasTraits(fName string, someTraits NSFontTraitMask) bool
	// Returns an approximation of the specified font’s weight.
	WeightOfFont(fontObj NSFont) int

	// Topic: Managing the Font Panel and Font Menu

	// A Boolean value that indicates whether the font conversion system’s Font panel and Font menu items are enabled.
	Enabled() bool
	SetEnabled(value bool)
	// Returns the application’s shared Font panel object, creating it if necessary.
	FontPanel(create bool) NSFontPanel
	// Records the given menu as the application’s Font menu.
	SetFontMenu(newMenu INSMenu)
	// Returns the menu that’s connected to the font conversion system, creating it if necessary.
	FontMenu(create bool) INSMenu

	// Topic: Accessing the Action Property

	// The action sent to the first responder when the user selects a new font from the Font panel or chooses a command from the Font menu.
	Action() objc.SEL
	SetAction(value objc.SEL)
	// The object that receives action messages related to the font manager.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)

	// Topic: Setting Attributes

	// Informs the Font panel that the specified font attributes changed for the selected text.
	SetSelectedAttributesIsMultiple(attributes foundation.INSDictionary, flag bool)
	// Converts attributes in response to an object initiating an attribute change, typically the Font panel or Font menu.
	ConvertAttributes(attributes foundation.INSDictionary) foundation.INSDictionary

	// The names of the currently loaded font collections.
	CollectionNames() objectivec.IObject
	SetCollectionNames(value objectivec.IObject)
	// The font manager’s delegate.
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
}





// Init initializes the instance.
func (f NSFontManager) Init() NSFontManager {
	rv := objc.Send[NSFontManager](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFontManager) Autorelease() NSFontManager {
	rv := objc.Send[NSFontManager](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFontManager creates a new NSFontManager instance.
func NewNSFontManager() NSFontManager {
	class := getNSFontManagerClass()
	rv := objc.Send[NSFontManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the names of the fonts available in the system whose traits are
// described exactly by the given font trait mask (not the [NSFont] objects
// themselves).
//
// someTraits: The font traits for which to return font names. You specify the desired
// traits by combining the font trait mask values described in [Constants]
// using the C bitwise OR operator.
//
// # Return Value
// 
// The names of the corresponding fonts.
//
// # Discussion
// 
// These fonts are in various system font directories.
// 
// If `someTraits` is 0, this method returns all fonts that are neither italic
// nor bold. This result is the same one you’d get if `fontTraitMask` were
// [NSUnitalicFontMask] `|` [NSUnboldFontMask].
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontNames(with:)
func (f NSFontManager) AvailableFontNamesWithTraits(someTraits NSFontTraitMask) []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("availableFontNamesWithTraits:"), someTraits)
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array with one entry for each available member of a font family.
//
// fam: The name of a font family, like one specified by the value of
// [AvailableFontFamilies].
//
// # Return Value
// 
// The available members of `family`. See the following discussion for a
// specific description.
//
// # Discussion
// 
// Each entry of the returned [NSArray] is another [NSArray] with four
// members, as follows:
// 
// - - The PostScript font name, as an [NSString] object. - - The part of the
// font name used in the font panel that’s not the font name, as an
// [NSString] object. This value is not localized—for example, `"Roman"`,
// `"Italic"`, or `"Bold"`. - - The font’s weight, as an [NSNumber]. - - The
// font’s traits, as an [NSNumber].
// 
// The members of the family are arranged in the font panel order (narrowest
// to widest, lightest to boldest, plain to italic).
// 
// For example, if you call `@"Times"`, it might return an array like this:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/availableMembers(ofFontFamily:)
func (f NSFontManager) AvailableMembersOfFontFamily(fam string) []foundation.NSArray {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("availableMembersOfFontFamily:"), objc.String(fam))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSArray {
		return foundation.NSArrayFromID(id)
	})
}

// Records the specified font as the currently selected font and updates the
// Font panel.
//
// fontObj: The font to set as selected.
//
// flag: If [true], the Font panel indicates that more than one font is contained in
// the selection; if [false], it does not.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// An object that manipulates fonts should invoke this method whenever it
// becomes first responder and whenever its selection changes. It shouldn’t
// invoke this method in the process of handling a [changeFont:] message, as
// this causes the font manager to lose the information necessary to effect
// the change. After all fonts have been converted, the font manager itself
// records the new selected font.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedFont(_:isMultiple:)
func (f NSFontManager) SetSelectedFontIsMultiple(fontObj NSFont, flag bool) {
	objc.Send[objc.ID](f.ID, objc.Sel("setSelectedFont:isMultiple:"), fontObj, flag)
}

// A Boolean value that indicates whether a responder handled the font
// manager’s action message.
//
// # Discussion
// 
// By default, the font manager’s action message is [changeFont:]. The value
// of this property is [true] if some responder object handled the
// [changeFont:] message or [false] if the message went unheard.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/sendAction()
func (f NSFontManager) SendAction() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("sendAction"))
	return rv
}

// Returns a localized string with the name of the specified font family and
// face, if one exists.
//
// family: The font family, for example, `@"Times"`.
//
// faceKey: The font face, for example, `@"Roman"`.
//
// # Return Value
// 
// A localized string with the name of the specified font family and face, or,
// if `face` is `nil`, the font family only.
//
// # Discussion
// 
// The user’s locale is determined from the user’s [NSLanguages] default
// setting. The method also loads the localized strings for the font, if they
// aren’t already loaded.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/localizedName(forFamily:face:)
func (f NSFontManager) LocalizedNameForFamilyFace(family string, faceKey string) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("localizedNameForFamily:face:"), objc.String(family), objc.String(faceKey))
	return foundation.NSStringFromID(rv).String()
}

// Adds a trait to the font.
//
// sender: The control that sent the message.
//
// # Discussion
// 
// By default, the action message is [changeFont:]. This action method causes
// the receiver to send its action message up the responder chain.
// 
// When a responder replies by providing a font to convert in a [ConvertFont]
// message, the receiver converts the font by adding the trait specified by
// `sender`. This trait is determined by sending a `tag` message to `sender`
// and interpreting it as a font trait mask for a [ConvertFontToHaveTrait]
// message.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/addFontTrait(_:)
func (f NSFontManager) AddFontTrait(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("addFontTrait:"), sender)
}

// Removes a trait from the font.
//
// sender: The control that sent the message.
//
// # Discussion
// 
// By default, the action message is [changeFont:]. This action method causes
// the receiver to send its action message up the responder chain.
// 
// When a responder replies by providing a font to convert in a [ConvertFont]
// message, the receiver converts the font by removing the trait specified by
// `sender`. This trait is determined by sending a `tag` message to `sender`
// and interpreting it as a font trait mask for a [ConvertFontToNotHaveTrait]
// message.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/removeFontTrait(_:)
func (f NSFontManager) RemoveFontTrait(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("removeFontTrait:"), sender)
}

// Modifies a trait of the font.
//
// sender: The control that sent the message.
//
// # Discussion
// 
// By default, the action message is [changeFont:]. This action method causes
// the receiver to send its action message up the responder chain.
// 
// When a responder replies by providing a font to convert in a [ConvertFont]
// message, the receiver converts the font in the manner specified by
// `sender`. The conversion is determined by sending a `tag` message to
// `sender` and invoking a corresponding method:
// 
// [Table data omitted]
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFont(_:)
func (f NSFontManager) ModifyFont(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("modifyFont:"), sender)
}

// Modifies a font trait using input from the Font panel.
//
// sender: The control that sent the message.
//
// # Discussion
// 
// By default, the action message is [changeFont:]. This action method causes
// the receiver to send its action message up the responder chain.
// 
// When a responder replies by providing a font to convert in a [ConvertFont]
// message, the receiver converts the font by sending a [PanelConvertFont]
// message to the Font panel. The panel in turn may send
// [ConvertFontToFamily], [ConvertFontToHaveTrait], and other specific
// conversion methods to make its change.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/modifyFontViaPanel(_:)
func (f NSFontManager) ModifyFontViaPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("modifyFontViaPanel:"), sender)
}

// Opens the Font Styles panel.
//
// sender: The control that sent the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontStylesPanel(_:)
func (f NSFontManager) OrderFrontStylesPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("orderFrontStylesPanel:"), sender)
}

// Opens the Font panel, creating it if necessary, and displays that panel in
// front of the app’s windows.
//
// sender: The control that sent the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/orderFrontFontPanel(_:)
func (f NSFontManager) OrderFrontFontPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("orderFrontFontPanel:"), sender)
}

// Converts the given font according to the object that initiated a font
// change, typically the Font panel or Font menu.
//
// fontObj: The font to convert.
//
// # Return Value
// 
// The converted font, or `aFont` itself if the conversion isn’t possible.
//
// # Discussion
// 
// This method is invoked in response to an action message such as
// [AddFontTrait] or [ModifyFontViaPanel]. These initiating methods cause the
// font manager to query the sender for the action to take and the traits to
// change. See Converting Fonts Manually for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:)
func (f NSFontManager) ConvertFont(fontObj NSFont) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:"), fontObj)
	return NSFontFromID(rv)
}

// Returns a font whose traits are as similar as possible to those of the
// given font except for the typeface, which is changed to the given typeface.
//
// fontObj: The font whose traits are matched.
//
// typeface: The new typeface; a fully specified family-face name, such as
// Helvetica-BoldOblique or Times-Roman.
//
// # Return Value
// 
// A font with matching traits and the given typeface, or `aFont` if it
// can’t be converted.
//
// # Discussion
// 
// This method attempts to match the weight and posture of `aFont` as closely
// as possible. Italic is mapped to Oblique, for example. Weights are mapped
// based on an approximate numeric scale of 0 to 15.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFace:)
func (f NSFontManager) ConvertFontToFace(fontObj NSFont, typeface string) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:toFace:"), fontObj, objc.String(typeface))
	return NSFontFromID(rv)
}

// Returns a font whose traits are as similar as possible to those of the
// given font except for the font family, which is changed to the given
// family.
//
// fontObj: The font whose traits are matched.
//
// family: The new font family; a generic font name, such as Helvetica or Times.
//
// # Return Value
// 
// A font with matching traits and the given family, or `aFont` if it can’t
// be converted.
//
// # Discussion
// 
// This method attempts to match the weight and posture of `aFont` as closely
// as possible. Italic is mapped to Oblique, for example. Weights are mapped
// based on an approximate numeric scale of 0 to 15.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toFamily:)
func (f NSFontManager) ConvertFontToFamily(fontObj NSFont, family string) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:toFamily:"), fontObj, objc.String(family))
	return NSFontFromID(rv)
}

// Returns a new version of the font object containing a single additional
// trait.
//
// fontObj: The font whose traits are matched.
//
// trait: The new trait; may be any one of the traits described in [Constants]. Using
// [NSUnboldFontMask] or [NSUnitalicFontMask] removes the bold or italic
// trait, respectively.
//
// # Return Value
// 
// A font with matching traits including the given trait, or `aFont` if it
// can’t be converted.
//
// # Discussion
// 
// Using [NSUnboldFontMask] or [NSUnitalicFontMask] removes the bold or italic
// trait, respectively.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toHaveTrait:)
func (f NSFontManager) ConvertFontToHaveTrait(fontObj NSFont, trait NSFontTraitMask) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:toHaveTrait:"), fontObj, trait)
	return NSFontFromID(rv)
}

// Returns a new version of a font object without the specified traits.
//
// fontObj: The font whose traits are matched.
//
// trait: The mask for the traits to remove, created using the C bitwise OR operator
// to combine the traits described in [Constants]. Using [BoldFontMask] or
// [ItalicFontMask] removes the bold or italic trait, respectively.
//
// # Return Value
// 
// A font with matching traits minus the given traits, or `aFont` if it
// can’t be converted.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toNotHaveTrait:)
func (f NSFontManager) ConvertFontToNotHaveTrait(fontObj NSFont, trait NSFontTraitMask) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:toNotHaveTrait:"), fontObj, trait)
	return NSFontFromID(rv)
}

// Returns a font object whose traits are the same as those of the given font,
// except for the size, which is changed to the given size.
//
// fontObj: The font whose traits are matched.
//
// size: The new font size.
//
// # Return Value
// 
// A font with matching traits except in the new size, or `aFont` if it
// can’t be converted.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convert(_:toSize:)
func (f NSFontManager) ConvertFontToSize(fontObj NSFont, size float64) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertFont:toSize:"), fontObj, size)
	return NSFontFromID(rv)
}

// Returns a font object whose weight is greater or lesser than that of the
// given font.
//
// upFlag: If [true], a heavier font is returned; if it’s [false], a lighter font is
// returned.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// fontObj: The font whose weight is increased or decreased.
//
// # Return Value
// 
// A font with matching traits except for the new weight, or `aFont` if it
// can’t be converted.
//
// # Discussion
// 
// Weights are graded along the following scale. The list on the left gives
// Apple’s terminology, and the list on the right gives the ISO equivalents.
// Names on the same line are treated as identical:
// 
// [Table data omitted]
// 
// The [NSFontManager] implementation of this method refuses to convert a
// font’s weight if it can’t maintain all other traits, such as italic and
// condensed. You might wish to override this method to allow a looser
// interpretation of weight conversion.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convertWeight(_:of:)
func (f NSFontManager) ConvertWeightOfFont(upFlag bool, fontObj NSFont) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertWeight:ofFont:"), upFlag, fontObj)
	return NSFontFromID(rv)
}

// Converts font traits to a new traits mask value.
//
// traits: The current font traits.
//
// # Return Value
// 
// The new traits mask value to be used by [ConvertFont].
//
// # Discussion
// 
// This method is intended to be invoked to query the font traits while the
// action message (usually [changeFont:]) is being invoked when the current
// font action is either [NSFontAction.addTraitFontAction] or
// [NSFontAction.removeTraitFontAction].
//
// [NSFontAction.addTraitFontAction]: https://developer.apple.com/documentation/AppKit/NSFontAction/addTraitFontAction
// [NSFontAction.removeTraitFontAction]: https://developer.apple.com/documentation/AppKit/NSFontAction/removeTraitFontAction
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convertFontTraits(_:)
func (f NSFontManager) ConvertFontTraits(traits NSFontTraitMask) NSFontTraitMask {
	rv := objc.Send[NSFontTraitMask](f.ID, objc.Sel("convertFontTraits:"), traits)
	return NSFontTraitMask(rv)
}

// Attempts to load a font with the specified characteristics.
//
// family: The generic name of the desired font, such as Times or Helvetica.
//
// traits: The font traits, specified by combining the font trait mask values
// described in [Constants] using the C bitwise OR operator. Using
// [NSUnboldFontMask] or [NSUnitalicFontMask] loads a font that doesn’t have
// either the bold or italic trait, respectively.
//
// weight: A hint for the weight desired, on a scale of 0 to 15: a value of 5
// indicates a normal or book weight, and 9 or more a bold or heavier weight.
// The weight is ignored if `fontTraitMask` includes [NSBoldFontMask].
//
// size: The point size of the desired font.
//
// # Return Value
// 
// A font with the specified characteristics if successful, or `nil` if not.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/font(withFamily:traits:weight:size:)
func (f NSFontManager) FontWithFamilyTraitsWeightSize(family string, traits NSFontTraitMask, weight int, size float64) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontWithFamily:traits:weight:size:"), objc.String(family), traits, weight, size)
	return NSFontFromID(rv)
}

// Returns the traits of the given font.
//
// fontObj: The font whose traits are returned.
//
// # Return Value
// 
// The font traits, returned as a mask created by combining values listed in
// [Constants] with the C bitwise OR operator.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/traits(of:)
func (f NSFontManager) TraitsOfFont(fontObj NSFont) NSFontTraitMask {
	rv := objc.Send[NSFontTraitMask](f.ID, objc.Sel("traitsOfFont:"), fontObj)
	return NSFontTraitMask(rv)
}

// Indicates whether the given font has all the specified traits.
//
// fName: The name of the font.
//
// someTraits: The font traits to test, specified by combining the font trait mask values
// described in [Constants] using the C bitwise OR operator.
//
// # Return Value
// 
// [true] if the font named `typeface` has all the traits specified in
// `fontTraitMask`; [false] if it doesn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Using [NSUnboldFontMask] returns [true] if the font is not bold, [false]
// otherwise. Using [NSUnitalicFontMask] returns [true] if the font is not
// italic, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/fontNamed(_:hasTraits:)
func (f NSFontManager) FontNamedHasTraits(fName string, someTraits NSFontTraitMask) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("fontNamed:hasTraits:"), objc.String(fName), someTraits)
	return rv
}

// Returns an approximation of the specified font’s weight.
//
// fontObj: The font whose approximate weight is returned.
//
// # Return Value
// 
// An approximation of the weight of the given font, where 0 indicates the
// lightest possible weight, 5 indicates a normal or book weight, and 9 or
// more indicates a bold or heavier weight. Because this method returns only
// an approximation of a font’s weight, it is not guaranteed to return the
// exact weight with which `fontObj` was initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/weight(of:)
func (f NSFontManager) WeightOfFont(fontObj NSFont) int {
	rv := objc.Send[int](f.ID, objc.Sel("weightOfFont:"), fontObj)
	return rv
}

// Returns the application’s shared Font panel object, creating it if
// necessary.
//
// create: If [true], the Font panel object is created if necessary; if [false], it is
// not.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The application’s shared Font panel object.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/fontPanel(_:)
func (f NSFontManager) FontPanel(create bool) NSFontPanel {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontPanel:"), create)
	return NSFontPanelFromID(rv)
}

// Records the given menu as the application’s Font menu.
//
// newMenu: The new Font menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontMenu(_:)
func (f NSFontManager) SetFontMenu(newMenu INSMenu) {
	objc.Send[objc.ID](f.ID, objc.Sel("setFontMenu:"), newMenu)
}

// Returns the menu that’s connected to the font conversion system, creating
// it if necessary.
//
// create: If [true], the menu object is created if necessary; if [false], it is not.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The font conversion system menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/fontMenu(_:)
func (f NSFontManager) FontMenu(create bool) INSMenu {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fontMenu:"), create)
	return NSMenuFromID(rv)
}

// Informs the Font panel that the specified font attributes changed for the
// selected text.
//
// attributes: The new attributes.
//
// flag: If [true], informs the panel that multiple fonts or attributes are enclosed
// within the selection.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is used primarily by [NSTextView].
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/setSelectedAttributes(_:isMultiple:)
func (f NSFontManager) SetSelectedAttributesIsMultiple(attributes foundation.INSDictionary, flag bool) {
	objc.Send[objc.ID](f.ID, objc.Sel("setSelectedAttributes:isMultiple:"), attributes, flag)
}

// Converts attributes in response to an object initiating an attribute
// change, typically the Font panel or Font menu.
//
// attributes: The current attributes.
//
// # Return Value
// 
// The converted attributes, or `attributes` itself if the conversion isn’t
// possible.
//
// # Discussion
// 
// Attributes unused by the sender should not be changed or removed.
// 
// This method is usually invoked on the sender of [ChangeAttributes]. See
// [Working with the Font Manager] for more information.
//
// [Working with the Font Manager]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/FontHandling/FontHandling.html#//apple_ref/doc/uid/TP40009459-CH5-SW9
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/convertAttributes(_:)
func (f NSFontManager) ConvertAttributes(attributes foundation.INSDictionary) foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("convertAttributes:"), attributes)
	return foundation.NSDictionaryFromID(rv)
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
// 
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (f NSFontManager) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}





// Sets the class that creates the shared font manager object.
//
// factoryId: The new font manager factory class, which must be a subclass of
// [NSFontManager].
//
// # Discussion
// 
// When you call the [SharedFontManager] method of [NSFontManager], it creates
// an instance of `aClass`, if no instance already exists. The class in
// `aClass` must implement `init` as its designated initializer. The default
// font manager factory is [NSFontManager].
// 
// Call this method before AppKit loads your application’s main nib file,
// such as in your app delegate’s [ApplicationWillFinishLaunching] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontManagerFactory(_:)
func (_NSFontManagerClass NSFontManagerClass) SetFontManagerFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSFontManagerClass.class), objc.Sel("setFontManagerFactory:"), factoryId)
}

// Sets the class that creates the shared Font panel object.
//
// factoryId: The new font panel factory class, which should be a subclass of
// [NSFontPanel].
//
// # Discussion
// 
// Call this method before accessing the Font panel in any way, such as in
// your app delegate’s [ApplicationWillFinishLaunching] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/setFontPanelFactory(_:)
func (_NSFontManagerClass NSFontManagerClass) SetFontPanelFactory(factoryId objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSFontManagerClass.class), objc.Sel("setFontPanelFactory:"), factoryId)
}








// The names of the fonts available in the system (not the [NSFont] objects
// themselves).
//
// # Discussion
// 
// Note that these fonts are in various system font directories.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFonts
func (f NSFontManager) AvailableFonts() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("availableFonts"))
	return objc.ConvertSliceToStrings(rv)
}



// The names of the font families available in the system.
//
// # Discussion
// 
// Note that these fonts are in various system font directories.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/availableFontFamilies
func (f NSFontManager) AvailableFontFamilies() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("availableFontFamilies"))
	return objc.ConvertSliceToStrings(rv)
}



// The currently selected font object.
//
// # Discussion
// 
// The value of this property is the last font recorded with a
// [SetSelectedFontIsMultiple] message.
// 
// While fonts are being converted in response to a [ConvertFont] message, you
// can determine the font selected in the Font panel like this:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/selectedFont
func (f NSFontManager) SelectedFont() NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("selectedFont"))
	return NSFontFromID(objc.ID(rv))
}



// A Boolean value that indicates whether the currently selected font has
// multiple fonts.
//
// # Discussion
// 
// When the value of this property is [true], the last font selection recorded
// has multiple fonts; if the last font selection recorded is a single font,
// the value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/isMultiple
func (f NSFontManager) Multiple() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isMultiple"))
	return rv
}



// The current font conversion action.
//
// # Discussion
// 
// The value of this property represents the current font action used by the
// [ConvertFont] method. This property is intended to be used to query the
// font conversion action while the action message (usually [changeFont:]) is
// being invoked.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/currentFontAction
func (f NSFontManager) CurrentFontAction() NSFontAction {
	rv := objc.Send[NSFontAction](f.ID, objc.Sel("currentFontAction"))
	return NSFontAction(rv)
}



// A Boolean value that indicates whether the font conversion system’s Font
// panel and Font menu items are enabled.
//
// # Discussion
// 
// When the value of this property is [true], the font conversion system’s
// user interface items (the Font panel and Font menu items) are enabled; when
// the value is [false], these items are not enabled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/isEnabled
func (f NSFontManager) Enabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEnabled"))
	return rv
}
func (f NSFontManager) SetEnabled(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setEnabled:"), value)
}



// The action sent to the first responder when the user selects a new font
// from the Font panel or chooses a command from the Font menu.
//
// # Discussion
// 
// The default action is [changeFont:]. You should rarely need to change this
// setting.
//
// [changeFont:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/changeFont:
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/action
func (f NSFontManager) Action() objc.SEL {
	rv := objc.Send[objc.SEL](f.ID, objc.Sel("action"))
	return rv
}
func (f NSFontManager) SetAction(value objc.SEL) {
	objc.Send[struct{}](f.ID, objc.Sel("setAction:"), value)
}



// The object that receives action messages related to the font manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/target
func (f NSFontManager) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (f NSFontManager) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setTarget:"), value)
}



// The names of the currently loaded font collections.
//
// See: https://developer.apple.com/documentation/appkit/nsfontmanager/collectionnames
func (f NSFontManager) CollectionNames() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("collectionNames"))
	return objectivec.Object{ID: rv}
}
func (f NSFontManager) SetCollectionNames(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setCollectionNames:"), value)
}



// The font manager’s delegate.
//
// See: https://developer.apple.com/documentation/appkit/nsfontmanager/delegate
func (f NSFontManager) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (f NSFontManager) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelegate:"), value)
}







// Returns the shared instance of the font manager for the application,
// creating it if necessary.
//
// # Return Value
// 
// The shared font manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontManager/shared
func (_NSFontManagerClass NSFontManagerClass) SharedFontManager() NSFontManager {
	rv := objc.Send[objc.ID](objc.ID(_NSFontManagerClass.class), objc.Sel("sharedFontManager"))
	return NSFontManagerFromID(objc.ID(rv))
}









			// Protocol methods for NSMenuItemValidation
			














