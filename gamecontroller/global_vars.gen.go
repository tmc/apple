// Code generated from Apple documentation. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// GCControllerDidBecomeCurrentNotification is a notification that posts when a controller becomes the current controller.
	//
	// See: https://developer.apple.com/documentation/GameController/GCControllerDidBecomeCurrentNotification
	GCControllerDidBecomeCurrentNotification string
	// GCControllerDidConnectNotification is a notification that posts after a controller connects to the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCControllerDidConnectNotification
	GCControllerDidConnectNotification string
	// GCControllerDidDisconnectNotification is a notification that posts after a controller disconnects from the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCControllerDidDisconnectNotification
	GCControllerDidDisconnectNotification string
	// GCControllerDidStopBeingCurrentNotification is a notification that posts when a controller stops being the current controller.
	//
	// See: https://developer.apple.com/documentation/GameController/GCControllerDidStopBeingCurrentNotification
	GCControllerDidStopBeingCurrentNotification string
	// GCControllerUserCustomizationsDidChangeNotification is a notification that posts when the user customizes the button mappings or other settings of a controller.
	//
	// See: https://developer.apple.com/documentation/GameController/GCControllerUserCustomizationsDidChangeNotification
	GCControllerUserCustomizationsDidChangeNotification string
	// GCInputDirectionalCardinalDpad is the name of the controller’s optional secondary directional pad element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDirectionalCardinalDpad
	GCInputDirectionalCardinalDpad string
	// GCInputDirectionalCenterButton is the name of the controller’s optional button on the directional gamepad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDirectionalCenterButton
	GCInputDirectionalCenterButton string
	// GCInputDirectionalDpad is the name of the controller’s primary directional pad element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDirectionalDpad
	GCInputDirectionalDpad string
	// GCInputDirectionalTouchSurfaceButton is the name of the controller’s touch surface button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDirectionalTouchSurfaceButton
	GCInputDirectionalTouchSurfaceButton string
	// GCInputMicroGamepadButtonA is the name of the micro gamepad’s primary button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputMicroGamepadButtonA
	GCInputMicroGamepadButtonA string
	// GCInputMicroGamepadButtonMenu is the name of the micro gamepad’s menu button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputMicroGamepadButtonMenu
	GCInputMicroGamepadButtonMenu string
	// GCInputMicroGamepadButtonX is the name of the micro gamepad’s secondary button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputMicroGamepadButtonX
	GCInputMicroGamepadButtonX string
	// GCInputMicroGamepadDpad is the name of the micro gamepad’s primary directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputMicroGamepadDpad
	GCInputMicroGamepadDpad string
	// GCKeyA is the keyboard code for the a or A character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyA
	GCKeyA string
	// GCKeyApplication is the keyboard code for the Application key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyApplication
	GCKeyApplication string
	// GCKeyB is the keyboard code for the b or B character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyB
	GCKeyB string
	// GCKeyBackslash is the keyboard code for the \ or | key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyBackslash
	GCKeyBackslash string
	// GCKeyC is the keyboard code for the c or C character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyC
	GCKeyC string
	// GCKeyCapsLock is the keyboard code for the Caps Lock key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyCapsLock
	GCKeyCapsLock string
	// GCKeyCloseBracket is the keyboard code for the ] or } key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyCloseBracket
	GCKeyCloseBracket string
	// GCKeyComma is the keyboard code for the Comma or < key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyComma
	GCKeyComma string
	// GCKeyD is the keyboard code for the d or D character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyD
	GCKeyD string
	// GCKeyDeleteForward is the keyboard code for the Delete-Forward key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyDeleteForward
	GCKeyDeleteForward string
	// GCKeyDeleteOrBackspace is the keyboard code for the Delete or Backspace key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyDeleteOrBackspace
	GCKeyDeleteOrBackspace string
	// GCKeyDownArrow is the keyboard code for the Down Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyDownArrow
	GCKeyDownArrow string
	// GCKeyE is the keyboard code for the e or E character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyE
	GCKeyE string
	// GCKeyEight is the keyboard code for the 8 or * character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyEight
	GCKeyEight string
	// GCKeyEnd is the keyboard code for the End key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyEnd
	GCKeyEnd string
	// GCKeyEqualSign is the keyboard code for the = or + key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyEqualSign
	GCKeyEqualSign string
	// GCKeyEscape is the keyboard code for the Escape key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyEscape
	GCKeyEscape string
	// GCKeyF is the keyboard code for the f or F character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF
	GCKeyF string
	// GCKeyF1 is the keyboard code for the F1 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF1
	GCKeyF1 string
	// GCKeyF10 is the keyboard code for the F10 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF10
	GCKeyF10 string
	// GCKeyF11 is the keyboard code for the F11 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF11
	GCKeyF11 string
	// GCKeyF12 is the keyboard code for the F12 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF12
	GCKeyF12 string
	// GCKeyF13 is the keyboard code for the F13 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF13
	GCKeyF13 string
	// GCKeyF14 is the keyboard code for the F14 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF14
	GCKeyF14 string
	// GCKeyF15 is the keyboard code for the F15 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF15
	GCKeyF15 string
	// GCKeyF16 is the keyboard code for the F16 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF16
	GCKeyF16 string
	// GCKeyF17 is the keyboard code for the F17 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF17
	GCKeyF17 string
	// GCKeyF18 is the keyboard code for the F18 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF18
	GCKeyF18 string
	// GCKeyF19 is the keyboard code for the F19 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF19
	GCKeyF19 string
	// GCKeyF2 is the keyboard code for the F2 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF2
	GCKeyF2 string
	// GCKeyF20 is the keyboard code for the F20 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF20
	GCKeyF20 string
	// GCKeyF3 is the keyboard code for the F3 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF3
	GCKeyF3 string
	// GCKeyF4 is the keyboard code for the F4 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF4
	GCKeyF4 string
	// GCKeyF5 is the keyboard code for the F5 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF5
	GCKeyF5 string
	// GCKeyF6 is the keyboard code for the F6 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF6
	GCKeyF6 string
	// GCKeyF7 is the keyboard code for the F7 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF7
	GCKeyF7 string
	// GCKeyF8 is the keyboard code for the F8 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF8
	GCKeyF8 string
	// GCKeyF9 is the keyboard code for the F9 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyF9
	GCKeyF9 string
	// GCKeyFive is the keyboard code for the 5 or % character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyFive
	GCKeyFive string
	// GCKeyFour is the keyboard code for the 4 or $ character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyFour
	GCKeyFour string
	// GCKeyG is the keyboard code for the g or G character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyG
	GCKeyG string
	// GCKeyGraveAccentAndTilde is the keyboard code for the Grave Accent or Tilde key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyGraveAccentAndTilde
	GCKeyGraveAccentAndTilde string
	// GCKeyH is the keyboard code for the h or H character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyH
	GCKeyH string
	// GCKeyHome is the keyboard code for the Home key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyHome
	GCKeyHome string
	// GCKeyHyphen is the keyboard code for the - or _ key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyHyphen
	GCKeyHyphen string
	// GCKeyI is the keyboard code for the i or I character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyI
	GCKeyI string
	// GCKeyInsert is the keyboard code for the Insert key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInsert
	GCKeyInsert string
	// GCKeyInternational1 is the keyboard code for the first international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational1
	GCKeyInternational1 string
	// GCKeyInternational2 is the keyboard code for the second international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational2
	GCKeyInternational2 string
	// GCKeyInternational3 is the keyboard code for the third international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational3
	GCKeyInternational3 string
	// GCKeyInternational4 is the keyboard code for the fourth international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational4
	GCKeyInternational4 string
	// GCKeyInternational5 is the keyboard code for the fifth international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational5
	GCKeyInternational5 string
	// GCKeyInternational6 is the keyboard code for the sixth international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational6
	GCKeyInternational6 string
	// GCKeyInternational7 is the keyboard code for the seventh international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational7
	GCKeyInternational7 string
	// GCKeyInternational8 is the keyboard code for the eighth international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational8
	GCKeyInternational8 string
	// GCKeyInternational9 is the keyboard code for the ninth international key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyInternational9
	GCKeyInternational9 string
	// GCKeyJ is the keyboard code for the j or J character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyJ
	GCKeyJ string
	// GCKeyK is the keyboard code for the k or K character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyK
	GCKeyK string
	// GCKeyKeypad0 is the keyboard code for the keypad 0 or Insert key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad0
	GCKeyKeypad0 string
	// GCKeyKeypad1 is the keyboard code for the keypad 1 or End key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad1
	GCKeyKeypad1 string
	// GCKeyKeypad2 is the keyboard code for the keypad 2 or Down Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad2
	GCKeyKeypad2 string
	// GCKeyKeypad3 is the keyboard code for the keypad 3 or Page Down key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad3
	GCKeyKeypad3 string
	// GCKeyKeypad4 is the keyboard code for the keypad 4 or Left Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad4
	GCKeyKeypad4 string
	// GCKeyKeypad5 is the keyboard code for the keypad 5 key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad5
	GCKeyKeypad5 string
	// GCKeyKeypad6 is the keyboard code for the keypad 6 or Right Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad6
	GCKeyKeypad6 string
	// GCKeyKeypad7 is the keyboard code for the keypad 7 or Home key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad7
	GCKeyKeypad7 string
	// GCKeyKeypad8 is the keyboard code for the keypad 8 or Up Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad8
	GCKeyKeypad8 string
	// GCKeyKeypad9 is the keyboard code for the keypad 9 or Page Up key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypad9
	GCKeyKeypad9 string
	// GCKeyKeypadAsterisk is the keyboard code for the keypad * key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadAsterisk
	GCKeyKeypadAsterisk string
	// GCKeyKeypadEnter is the keyboard code for the keypad Enter key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadEnter
	GCKeyKeypadEnter string
	// GCKeyKeypadEqualSign is the keyboard code for the keypad = key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadEqualSign
	GCKeyKeypadEqualSign string
	// GCKeyKeypadHyphen is the keyboard code for the keypad - key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadHyphen
	GCKeyKeypadHyphen string
	// GCKeyKeypadNumLock is the keyboard code for the keypad Num Lock or Clear key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadNumLock
	GCKeyKeypadNumLock string
	// GCKeyKeypadPeriod is the keyboard code for the keypad Period or Delete key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadPeriod
	GCKeyKeypadPeriod string
	// GCKeyKeypadPlus is the keyboard code for the keypad + key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadPlus
	GCKeyKeypadPlus string
	// GCKeyKeypadSlash is the keyboard code for the keypad / key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyKeypadSlash
	GCKeyKeypadSlash string
	// GCKeyL is the keyboard code for the l or L character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyL
	GCKeyL string
	// GCKeyLANG1 is the keyboard code for the first language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG1
	GCKeyLANG1 string
	// GCKeyLANG2 is the keyboard code for the second language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG2
	GCKeyLANG2 string
	// GCKeyLANG3 is the keyboard code for the third language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG3
	GCKeyLANG3 string
	// GCKeyLANG4 is the keyboard code for the fourth language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG4
	GCKeyLANG4 string
	// GCKeyLANG5 is the keyboard code for the fifth language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG5
	GCKeyLANG5 string
	// GCKeyLANG6 is the keyboard code for the sixth language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG6
	GCKeyLANG6 string
	// GCKeyLANG7 is the keyboard code for the seventh language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG7
	GCKeyLANG7 string
	// GCKeyLANG8 is the keyboard code for the eighth language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG8
	GCKeyLANG8 string
	// GCKeyLANG9 is the keyboard code for the ninth language key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLANG9
	GCKeyLANG9 string
	// GCKeyLeftAlt is the keyboard code for the Option or Alt key on the left side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLeftAlt
	GCKeyLeftAlt string
	// GCKeyLeftArrow is the keyboard code for the Left Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLeftArrow
	GCKeyLeftArrow string
	// GCKeyLeftControl is the keyboard code for the Control key on the left side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLeftControl
	GCKeyLeftControl string
	// GCKeyLeftGUI is the keyboard code for the Command key on the left side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLeftGUI
	GCKeyLeftGUI string
	// GCKeyLeftShift is the keyboard code for the Shift key on the left side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyLeftShift
	GCKeyLeftShift string
	// GCKeyM is the keyboard code for the m or M character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyM
	GCKeyM string
	// GCKeyN is the keyboard code for the n or N character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyN
	GCKeyN string
	// GCKeyNine is the keyboard code for the 9 or ( character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyNine
	GCKeyNine string
	// GCKeyNonUSBackslash is the keyboard code for the non-US Slash or | key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyNonUSBackslash
	GCKeyNonUSBackslash string
	// GCKeyNonUSPound is the keyboard code for the non-US Pound or _ key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyNonUSPound
	GCKeyNonUSPound string
	// GCKeyO is the keyboard code for the o or O character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyO
	GCKeyO string
	// GCKeyOne is the keyboard code for the 1 or ! character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyOne
	GCKeyOne string
	// GCKeyOpenBracket is the keyboard code for the [ or { key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyOpenBracket
	GCKeyOpenBracket string
	// GCKeyP is the keyboard code for the p or P character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyP
	GCKeyP string
	// GCKeyPageDown is the keyboard code for the Page Down key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPageDown
	GCKeyPageDown string
	// GCKeyPageUp is the keyboard code for the Page Up key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPageUp
	GCKeyPageUp string
	// GCKeyPause is the keyboard code for the Pause key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPause
	GCKeyPause string
	// GCKeyPeriod is the keyboard code for the Period or > key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPeriod
	GCKeyPeriod string
	// GCKeyPower is the keyboard code for the Power key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPower
	GCKeyPower string
	// GCKeyPrintScreen is the keyboard code for the Print Screen key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyPrintScreen
	GCKeyPrintScreen string
	// GCKeyQ is the keyboard code for the q or Q character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyQ
	GCKeyQ string
	// GCKeyQuote is the keyboard code for the ’ or “ key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyQuote
	GCKeyQuote string
	// GCKeyR is the keyboard code for the r or R character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyR
	GCKeyR string
	// GCKeyReturnOrEnter is the keyboard code for the Return or Enter key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyReturnOrEnter
	GCKeyReturnOrEnter string
	// GCKeyRightAlt is the keyboard code for the Option or Alt key on the right side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyRightAlt
	GCKeyRightAlt string
	// GCKeyRightArrow is the keyboard code for the Right Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyRightArrow
	GCKeyRightArrow string
	// GCKeyRightControl is the keyboard code for the Control key on the right side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyRightControl
	GCKeyRightControl string
	// GCKeyRightGUI is the keyboard code for the Command key on the right side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyRightGUI
	GCKeyRightGUI string
	// GCKeyRightShift is the keyboard code for the Shift key on the right side of the keyboard.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyRightShift
	GCKeyRightShift string
	// GCKeyS is the keyboard code for the s or S character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyS
	GCKeyS string
	// GCKeyScrollLock is the keyboard code for the Scroll Lock key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyScrollLock
	GCKeyScrollLock string
	// GCKeySemicolon is the keyboard code for the ; or : key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeySemicolon
	GCKeySemicolon string
	// GCKeySeven is the keyboard code for the 7 or & character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeySeven
	GCKeySeven string
	// GCKeySix is the keyboard code for the 6 or ^ character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeySix
	GCKeySix string
	// GCKeySlash is the keyboard code for the / or ? key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeySlash
	GCKeySlash string
	// GCKeySpacebar is the keyboard code for the Space Bar key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeySpacebar
	GCKeySpacebar string
	// GCKeyT is the keyboard code for the t or T character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyT
	GCKeyT string
	// GCKeyTab is the keyboard code for the Tab key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyTab
	GCKeyTab string
	// GCKeyThree is the keyboard code for the 3 or # character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyThree
	GCKeyThree string
	// GCKeyTwo is the keyboard code for the 2 or @ character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyTwo
	GCKeyTwo string
	// GCKeyU is the keyboard code for the u or U character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyU
	GCKeyU string
	// GCKeyUpArrow is the keyboard code for the Up Arrow key.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyUpArrow
	GCKeyUpArrow string
	// GCKeyV is the keyboard code for the v or V character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyV
	GCKeyV string
	// GCKeyW is the keyboard code for the w or W character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyW
	GCKeyW string
	// GCKeyX is the keyboard code for the x or X character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyX
	GCKeyX string
	// GCKeyY is the keyboard code for the y or Y character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyY
	GCKeyY string
	// GCKeyZ is the keyboard code for the z or Z character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyZ
	GCKeyZ string
	// GCKeyZero is the keyboard code for the 0 or ) character.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyZero
	GCKeyZero string
	// GCKeyboardDidConnectNotification is a notification that posts after a keyboard connects to the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyboardDidConnectNotification
	GCKeyboardDidConnectNotification string
	// GCKeyboardDidDisconnectNotification is a notification that posts after a single keyboard, or the last of multiple keyboards, disconnects from the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCKeyboardDidDisconnectNotification
	GCKeyboardDidDisconnectNotification string
	// GCMouseDidBecomeCurrentNotification is a notification that posts when a mouse becomes the most recent mouse that the user connects.
	//
	// See: https://developer.apple.com/documentation/GameController/GCMouseDidBecomeCurrentNotification
	GCMouseDidBecomeCurrentNotification string
	// GCMouseDidConnectNotification is a notification that posts after a mouse connects to the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCMouseDidConnectNotification
	GCMouseDidConnectNotification string
	// GCMouseDidDisconnectNotification is a notification that posts after a mouse disconnects from the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCMouseDidDisconnectNotification
	GCMouseDidDisconnectNotification string
	// GCMouseDidStopBeingCurrentNotification is a notification that posts when a mouse stops being the most recent mouse that the user connects.
	//
	// See: https://developer.apple.com/documentation/GameController/GCMouseDidStopBeingCurrentNotification
	GCMouseDidStopBeingCurrentNotification string
	// GCProductCategoryArcadeStick is the arcade stick product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryArcadeStick
	GCProductCategoryArcadeStick string
	// GCProductCategoryCoalescedRemote is the Apple TV Remote product category for physical and virtual remotes that Game Center combines into a single device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryCoalescedRemote
	GCProductCategoryCoalescedRemote string
	// GCProductCategoryControlCenterRemote is the virtual remote in the Control Center on iOS and tvOS devices for controlling the Apple TV.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryControlCenterRemote
	GCProductCategoryControlCenterRemote string
	// GCProductCategoryDualSense is the DualSense product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryDualSense
	GCProductCategoryDualSense string
	// GCProductCategoryDualShock4 is the DualShock 4 product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryDualShock4
	GCProductCategoryDualShock4 string
	// GCProductCategoryHID is the category for products that support the human interface device (HID) protocol.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryHID
	GCProductCategoryHID string
	// GCProductCategoryKeyboard is the keyboard product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryKeyboard
	GCProductCategoryKeyboard string
	// GCProductCategoryMFi is the MFi product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryMFi
	GCProductCategoryMFi string
	// GCProductCategoryMouse is the mouse product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryMouse
	GCProductCategoryMouse string
	// GCProductCategorySiriRemote1stGen is the first-generation Siri Remote or first-generation Apple TV Remote product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategorySiriRemote1stGen
	GCProductCategorySiriRemote1stGen string
	// GCProductCategorySiriRemote2ndGen is the second-generation Siri Remote or second-generation Apple TV Remote product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategorySiriRemote2ndGen
	GCProductCategorySiriRemote2ndGen string
	// GCProductCategorySpatialController is the category for game controller products that support 6DoF tracking on visionOS.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategorySpatialController
	GCProductCategorySpatialController string
	// GCProductCategoryUniversalElectronicsRemote is the product category for a Universal Electronics remote that works with Apple TV.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryUniversalElectronicsRemote
	GCProductCategoryUniversalElectronicsRemote string
	// GCProductCategoryXboxOne is the Xbox product category.
	//
	// See: https://developer.apple.com/documentation/GameController/GCProductCategoryXboxOne
	GCProductCategoryXboxOne string
	// GCRacingWheelDidConnectNotification is a notification that posts after a racing wheel controller connects to the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRacingWheelDidConnectNotification
	GCRacingWheelDidConnectNotification string
	// GCRacingWheelDidDisconnectNotification is a notification that posts after a racing wheel controller disconnects from the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRacingWheelDidDisconnectNotification
	GCRacingWheelDidDisconnectNotification string
)

var (
	// GCHapticDurationInfinite is an infinite duration for a haptics event.
	//
	// See: https://developer.apple.com/documentation/GameController/GCHapticDurationInfinite
	GCHapticDurationInfinite float32
)

var ()

var (
	// GCInputButtonA is the name for the controller’s A button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonA-a2w2
	GCInputButtonA GCInputButtonName
	// GCInputButtonB is the name for the controller’s B button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonB-a2w1
	GCInputButtonB GCInputButtonName
	// GCInputButtonHome is the name of the home or logo button element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonHome-8qxo4
	GCInputButtonHome GCInputButtonName
	// GCInputButtonMenu is the name of the primary menu button element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonMenu-83i52
	GCInputButtonMenu GCInputButtonName
	// GCInputButtonOptions is the name of the secondary menu button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonOptions-52c9a
	GCInputButtonOptions GCInputButtonName
	// GCInputButtonShare is the name of the share button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonShare-2xxjg
	GCInputButtonShare GCInputButtonName
	// GCInputButtonX is the name for the controller’s X button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonX-a2wr
	GCInputButtonX GCInputButtonName
	// GCInputButtonY is the name for the controller’s Y button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputButtonY-a2wq
	GCInputButtonY GCInputButtonName
	// GCInputDualShockTouchpadButton is the name of the button functionality of the DualShock 4 touchpad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDualShockTouchpadButton-47b5w
	GCInputDualShockTouchpadButton GCInputButtonName
	// See: https://developer.apple.com/documentation/GameController/GCInputGripButton
	GCInputGripButton GCInputButtonName
	// GCInputLeftBumper is the name of the additional left shoulder button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftBumper
	GCInputLeftBumper GCInputButtonName
	// GCInputLeftPaddle is the name for the left paddle input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftPaddle-818ad
	GCInputLeftPaddle GCInputButtonName
	// GCInputLeftShoulder is the name of the left shoulder button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftShoulder-3i8m0
	GCInputLeftShoulder GCInputButtonName
	// GCInputLeftThumbstickButton is the name of the left thumbstick button element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftThumbstickButton-4oxvn
	GCInputLeftThumbstickButton GCInputButtonName
	// GCInputLeftTrigger is the name of the left trigger.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftTrigger-4gjmx
	GCInputLeftTrigger GCInputButtonName
	// GCInputPedalAccelerator is the name of the accelerator element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputPedalAccelerator-6kjfe
	GCInputPedalAccelerator GCInputButtonName
	// GCInputPedalBrake is the name of the brake element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputPedalBrake-7jcp
	GCInputPedalBrake GCInputButtonName
	// GCInputPedalClutch is the name of the clutch element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputPedalClutch-3cn3p
	GCInputPedalClutch GCInputButtonName
	// GCInputRightBumper is the name of the additional right shoulder button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightBumper
	GCInputRightBumper GCInputButtonName
	// GCInputRightPaddle is the name for the right paddle input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightPaddle-1p8hd
	GCInputRightPaddle GCInputButtonName
	// GCInputRightShoulder is the name of the right shoulder button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightShoulder-513b5
	GCInputRightShoulder GCInputButtonName
	// GCInputRightThumbstickButton is the name of the right thumbstick button element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightThumbstickButton-3k0ld
	GCInputRightThumbstickButton GCInputButtonName
	// GCInputRightTrigger is the name of the right trigger.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightTrigger-6hvns
	GCInputRightTrigger GCInputButtonName
	// See: https://developer.apple.com/documentation/GameController/GCInputThumbstickButton
	GCInputThumbstickButton GCInputButtonName
	// See: https://developer.apple.com/documentation/GameController/GCInputTrigger
	GCInputTrigger GCInputButtonName
	// GCInputXboxPaddleFour is the name for the controller’s P4 paddle button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputXboxPaddleFour-4n8ze
	GCInputXboxPaddleFour GCInputButtonName
	// GCInputXboxPaddleOne is the name for the controller’s P1 paddle button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputXboxPaddleOne-795e3
	GCInputXboxPaddleOne GCInputButtonName
	// GCInputXboxPaddleThree is the name for the controller’s P3 paddle button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputXboxPaddleThree-7u2xy
	GCInputXboxPaddleThree GCInputButtonName
	// GCInputXboxPaddleTwo is the name for the controller’s P2 paddle button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputXboxPaddleTwo-4vrw4
	GCInputXboxPaddleTwo GCInputButtonName
)

var (
	// GCInputDirectionPad is the name of the directional pad element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDirectionPad-8igy2
	GCInputDirectionPad GCInputDirectionPadName
	// GCInputDualShockTouchpadOne is the name of the first finger element to touch down on the DualShock 4 touchpad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDualShockTouchpadOne-5u54s
	GCInputDualShockTouchpadOne GCInputDirectionPadName
	// GCInputDualShockTouchpadTwo is the name of the second finger element to touch down on the DualShock 4 touchpad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputDualShockTouchpadTwo-38d95
	GCInputDualShockTouchpadTwo GCInputDirectionPadName
	// GCInputLeftThumbstick is the name of the left thumbstick element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputLeftThumbstick-5tkci
	GCInputLeftThumbstick GCInputDirectionPadName
	// GCInputRightThumbstick is the name of the right thumbstick element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputRightThumbstick-5o085
	GCInputRightThumbstick GCInputDirectionPadName
	// See: https://developer.apple.com/documentation/GameController/GCInputThumbstick
	GCInputThumbstick GCInputDirectionPadName
)

var (
	// GCInputShifter is the name of the shifter element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputShifter-2p2b7
	GCInputShifter GCInputElementName
)

var (
	// GCInputSteeringWheel is the name of the steering wheel element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCInputSteeringWheel-6eb58
	GCInputSteeringWheel GCInputAxisName
)

var ()

var (
	// GCPoint2Zero is the origin for a two dimensional point.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPoint2Zero
	GCPoint2Zero GCPoint2
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCControllerDidBecomeCurrentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCControllerDidBecomeCurrentNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCControllerDidConnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCControllerDidConnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCControllerDidDisconnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCControllerDidDisconnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCControllerDidStopBeingCurrentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCControllerDidStopBeingCurrentNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCControllerUserCustomizationsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCControllerUserCustomizationsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticDurationInfinite"); err == nil && ptr != 0 {
		GCHapticDurationInfinite = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityAll"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.All = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.Default = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityHandles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.Handles = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityLeftHandle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.LeftHandle = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityLeftTrigger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.LeftTrigger = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityRightHandle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.RightHandle = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityRightTrigger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.RightTrigger = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCHapticsLocalityTriggers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCHapticsLocalitys.Triggers = GCHapticsLocality(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonA"); err == nil && ptr != 0 {
		GCInputButtonA = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonB"); err == nil && ptr != 0 {
		GCInputButtonB = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonHome"); err == nil && ptr != 0 {
		GCInputButtonHome = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonMenu"); err == nil && ptr != 0 {
		GCInputButtonMenu = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonOptions"); err == nil && ptr != 0 {
		GCInputButtonOptions = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonShare"); err == nil && ptr != 0 {
		GCInputButtonShare = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonX"); err == nil && ptr != 0 {
		GCInputButtonX = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputButtonY"); err == nil && ptr != 0 {
		GCInputButtonY = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDirectionPad"); err == nil && ptr != 0 {
		GCInputDirectionPad = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDirectionalCardinalDpad"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputDirectionalCardinalDpad = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDirectionalCenterButton"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputDirectionalCenterButton = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDirectionalDpad"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputDirectionalDpad = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDirectionalTouchSurfaceButton"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputDirectionalTouchSurfaceButton = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDualShockTouchpadButton"); err == nil && ptr != 0 {
		GCInputDualShockTouchpadButton = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDualShockTouchpadOne"); err == nil && ptr != 0 {
		GCInputDualShockTouchpadOne = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputDualShockTouchpadTwo"); err == nil && ptr != 0 {
		GCInputDualShockTouchpadTwo = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputGripButton"); err == nil && ptr != 0 {
		GCInputGripButton = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftBumper"); err == nil && ptr != 0 {
		GCInputLeftBumper = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftPaddle"); err == nil && ptr != 0 {
		GCInputLeftPaddle = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftShoulder"); err == nil && ptr != 0 {
		GCInputLeftShoulder = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftThumbstick"); err == nil && ptr != 0 {
		GCInputLeftThumbstick = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftThumbstickButton"); err == nil && ptr != 0 {
		GCInputLeftThumbstickButton = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputLeftTrigger"); err == nil && ptr != 0 {
		GCInputLeftTrigger = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputMicroGamepadButtonA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputMicroGamepadButtonA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputMicroGamepadButtonMenu"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputMicroGamepadButtonMenu = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputMicroGamepadButtonX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputMicroGamepadButtonX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputMicroGamepadDpad"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCInputMicroGamepadDpad = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputPedalAccelerator"); err == nil && ptr != 0 {
		GCInputPedalAccelerator = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputPedalBrake"); err == nil && ptr != 0 {
		GCInputPedalBrake = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputPedalClutch"); err == nil && ptr != 0 {
		GCInputPedalClutch = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightBumper"); err == nil && ptr != 0 {
		GCInputRightBumper = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightPaddle"); err == nil && ptr != 0 {
		GCInputRightPaddle = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightShoulder"); err == nil && ptr != 0 {
		GCInputRightShoulder = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightThumbstick"); err == nil && ptr != 0 {
		GCInputRightThumbstick = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightThumbstickButton"); err == nil && ptr != 0 {
		GCInputRightThumbstickButton = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputRightTrigger"); err == nil && ptr != 0 {
		GCInputRightTrigger = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputShifter"); err == nil && ptr != 0 {
		GCInputShifter = *(*GCInputElementName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputSteeringWheel"); err == nil && ptr != 0 {
		GCInputSteeringWheel = *(*GCInputAxisName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputThumbstick"); err == nil && ptr != 0 {
		GCInputThumbstick = *(*GCInputDirectionPadName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputThumbstickButton"); err == nil && ptr != 0 {
		GCInputThumbstickButton = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputTrigger"); err == nil && ptr != 0 {
		GCInputTrigger = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputXboxPaddleFour"); err == nil && ptr != 0 {
		GCInputXboxPaddleFour = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputXboxPaddleOne"); err == nil && ptr != 0 {
		GCInputXboxPaddleOne = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputXboxPaddleThree"); err == nil && ptr != 0 {
		GCInputXboxPaddleThree = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCInputXboxPaddleTwo"); err == nil && ptr != 0 {
		GCInputXboxPaddleTwo = *(*GCInputButtonName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyApplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyApplication = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyBackslash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyBackslash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCapsLock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyCapsLock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCloseBracket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyCloseBracket = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeApplication"); err == nil && ptr != 0 {
		GCKeyCodes.Application = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeBackslash"); err == nil && ptr != 0 {
		GCKeyCodes.Backslash = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeCapsLock"); err == nil && ptr != 0 {
		GCKeyCodes.CapsLock = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeCloseBracket"); err == nil && ptr != 0 {
		GCKeyCodes.CloseBracket = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeComma"); err == nil && ptr != 0 {
		GCKeyCodes.Comma = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeDeleteForward"); err == nil && ptr != 0 {
		GCKeyCodes.DeleteForward = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeDeleteOrBackspace"); err == nil && ptr != 0 {
		GCKeyCodes.DeleteOrBackspace = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeDownArrow"); err == nil && ptr != 0 {
		GCKeyCodes.DownArrow = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeEight"); err == nil && ptr != 0 {
		GCKeyCodes.Eight = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeEnd"); err == nil && ptr != 0 {
		GCKeyCodes.End = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeEqualSign"); err == nil && ptr != 0 {
		GCKeyCodes.EqualSign = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeEscape"); err == nil && ptr != 0 {
		GCKeyCodes.Escape = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF1"); err == nil && ptr != 0 {
		GCKeyCodes.F1 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF10"); err == nil && ptr != 0 {
		GCKeyCodes.F10 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF11"); err == nil && ptr != 0 {
		GCKeyCodes.F11 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF12"); err == nil && ptr != 0 {
		GCKeyCodes.F12 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF13"); err == nil && ptr != 0 {
		GCKeyCodes.F13 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF14"); err == nil && ptr != 0 {
		GCKeyCodes.F14 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF15"); err == nil && ptr != 0 {
		GCKeyCodes.F15 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF16"); err == nil && ptr != 0 {
		GCKeyCodes.F16 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF17"); err == nil && ptr != 0 {
		GCKeyCodes.F17 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF18"); err == nil && ptr != 0 {
		GCKeyCodes.F18 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF19"); err == nil && ptr != 0 {
		GCKeyCodes.F19 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF2"); err == nil && ptr != 0 {
		GCKeyCodes.F2 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF20"); err == nil && ptr != 0 {
		GCKeyCodes.F20 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF3"); err == nil && ptr != 0 {
		GCKeyCodes.F3 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF4"); err == nil && ptr != 0 {
		GCKeyCodes.F4 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF5"); err == nil && ptr != 0 {
		GCKeyCodes.F5 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF6"); err == nil && ptr != 0 {
		GCKeyCodes.F6 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF7"); err == nil && ptr != 0 {
		GCKeyCodes.F7 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF8"); err == nil && ptr != 0 {
		GCKeyCodes.F8 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeF9"); err == nil && ptr != 0 {
		GCKeyCodes.F9 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeFive"); err == nil && ptr != 0 {
		GCKeyCodes.Five = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeFour"); err == nil && ptr != 0 {
		GCKeyCodes.Four = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeGraveAccentAndTilde"); err == nil && ptr != 0 {
		GCKeyCodes.GraveAccentAndTilde = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeHome"); err == nil && ptr != 0 {
		GCKeyCodes.Home = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeHyphen"); err == nil && ptr != 0 {
		GCKeyCodes.Hyphen = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInsert"); err == nil && ptr != 0 {
		GCKeyCodes.Insert = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational1"); err == nil && ptr != 0 {
		GCKeyCodes.International1 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational2"); err == nil && ptr != 0 {
		GCKeyCodes.International2 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational3"); err == nil && ptr != 0 {
		GCKeyCodes.International3 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational4"); err == nil && ptr != 0 {
		GCKeyCodes.International4 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational5"); err == nil && ptr != 0 {
		GCKeyCodes.International5 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational6"); err == nil && ptr != 0 {
		GCKeyCodes.International6 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational7"); err == nil && ptr != 0 {
		GCKeyCodes.International7 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational8"); err == nil && ptr != 0 {
		GCKeyCodes.International8 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeInternational9"); err == nil && ptr != 0 {
		GCKeyCodes.International9 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyA"); err == nil && ptr != 0 {
		GCKeyCodes.KeyA = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyB"); err == nil && ptr != 0 {
		GCKeyCodes.KeyB = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyC"); err == nil && ptr != 0 {
		GCKeyCodes.KeyC = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyD"); err == nil && ptr != 0 {
		GCKeyCodes.KeyD = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyE"); err == nil && ptr != 0 {
		GCKeyCodes.KeyE = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyF"); err == nil && ptr != 0 {
		GCKeyCodes.KeyF = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyG"); err == nil && ptr != 0 {
		GCKeyCodes.KeyG = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyH"); err == nil && ptr != 0 {
		GCKeyCodes.KeyH = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyI"); err == nil && ptr != 0 {
		GCKeyCodes.KeyI = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyJ"); err == nil && ptr != 0 {
		GCKeyCodes.KeyJ = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyK"); err == nil && ptr != 0 {
		GCKeyCodes.KeyK = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyL"); err == nil && ptr != 0 {
		GCKeyCodes.KeyL = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyM"); err == nil && ptr != 0 {
		GCKeyCodes.KeyM = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyN"); err == nil && ptr != 0 {
		GCKeyCodes.KeyN = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyO"); err == nil && ptr != 0 {
		GCKeyCodes.KeyO = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyP"); err == nil && ptr != 0 {
		GCKeyCodes.KeyP = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyQ"); err == nil && ptr != 0 {
		GCKeyCodes.KeyQ = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyR"); err == nil && ptr != 0 {
		GCKeyCodes.KeyR = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyS"); err == nil && ptr != 0 {
		GCKeyCodes.KeyS = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyT"); err == nil && ptr != 0 {
		GCKeyCodes.KeyT = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyU"); err == nil && ptr != 0 {
		GCKeyCodes.KeyU = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyV"); err == nil && ptr != 0 {
		GCKeyCodes.KeyV = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyW"); err == nil && ptr != 0 {
		GCKeyCodes.KeyW = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyX"); err == nil && ptr != 0 {
		GCKeyCodes.KeyX = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyY"); err == nil && ptr != 0 {
		GCKeyCodes.KeyY = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeyZ"); err == nil && ptr != 0 {
		GCKeyCodes.KeyZ = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad0"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad0 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad1"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad1 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad2"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad2 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad3"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad3 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad4"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad4 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad5"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad5 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad6"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad6 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad7"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad7 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad8"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad8 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypad9"); err == nil && ptr != 0 {
		GCKeyCodes.Keypad9 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadAsterisk"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadAsterisk = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadEnter"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadEnter = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadEqualSign"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadEqualSign = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadHyphen"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadHyphen = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadNumLock"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadNumLock = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadPeriod"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadPeriod = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadPlus"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadPlus = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeKeypadSlash"); err == nil && ptr != 0 {
		GCKeyCodes.KeypadSlash = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG1"); err == nil && ptr != 0 {
		GCKeyCodes.LANG1 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG2"); err == nil && ptr != 0 {
		GCKeyCodes.LANG2 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG3"); err == nil && ptr != 0 {
		GCKeyCodes.LANG3 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG4"); err == nil && ptr != 0 {
		GCKeyCodes.LANG4 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG5"); err == nil && ptr != 0 {
		GCKeyCodes.LANG5 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG6"); err == nil && ptr != 0 {
		GCKeyCodes.LANG6 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG7"); err == nil && ptr != 0 {
		GCKeyCodes.LANG7 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG8"); err == nil && ptr != 0 {
		GCKeyCodes.LANG8 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLANG9"); err == nil && ptr != 0 {
		GCKeyCodes.LANG9 = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLeftAlt"); err == nil && ptr != 0 {
		GCKeyCodes.LeftAlt = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLeftArrow"); err == nil && ptr != 0 {
		GCKeyCodes.LeftArrow = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLeftControl"); err == nil && ptr != 0 {
		GCKeyCodes.LeftControl = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLeftGUI"); err == nil && ptr != 0 {
		GCKeyCodes.LeftGUI = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeLeftShift"); err == nil && ptr != 0 {
		GCKeyCodes.LeftShift = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeNine"); err == nil && ptr != 0 {
		GCKeyCodes.Nine = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeNonUSBackslash"); err == nil && ptr != 0 {
		GCKeyCodes.NonUSBackslash = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeNonUSPound"); err == nil && ptr != 0 {
		GCKeyCodes.NonUSPound = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeOne"); err == nil && ptr != 0 {
		GCKeyCodes.One = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeOpenBracket"); err == nil && ptr != 0 {
		GCKeyCodes.OpenBracket = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePageDown"); err == nil && ptr != 0 {
		GCKeyCodes.PageDown = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePageUp"); err == nil && ptr != 0 {
		GCKeyCodes.PageUp = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePause"); err == nil && ptr != 0 {
		GCKeyCodes.Pause = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePeriod"); err == nil && ptr != 0 {
		GCKeyCodes.Period = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePower"); err == nil && ptr != 0 {
		GCKeyCodes.Power = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodePrintScreen"); err == nil && ptr != 0 {
		GCKeyCodes.PrintScreen = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeQuote"); err == nil && ptr != 0 {
		GCKeyCodes.Quote = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeReturnOrEnter"); err == nil && ptr != 0 {
		GCKeyCodes.ReturnOrEnter = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeRightAlt"); err == nil && ptr != 0 {
		GCKeyCodes.RightAlt = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeRightArrow"); err == nil && ptr != 0 {
		GCKeyCodes.RightArrow = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeRightControl"); err == nil && ptr != 0 {
		GCKeyCodes.RightControl = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeRightGUI"); err == nil && ptr != 0 {
		GCKeyCodes.RightGUI = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeRightShift"); err == nil && ptr != 0 {
		GCKeyCodes.RightShift = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeScrollLock"); err == nil && ptr != 0 {
		GCKeyCodes.ScrollLock = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeSemicolon"); err == nil && ptr != 0 {
		GCKeyCodes.Semicolon = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeSeven"); err == nil && ptr != 0 {
		GCKeyCodes.Seven = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeSix"); err == nil && ptr != 0 {
		GCKeyCodes.Six = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeSlash"); err == nil && ptr != 0 {
		GCKeyCodes.Slash = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeSpacebar"); err == nil && ptr != 0 {
		GCKeyCodes.Spacebar = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeTab"); err == nil && ptr != 0 {
		GCKeyCodes.Tab = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeThree"); err == nil && ptr != 0 {
		GCKeyCodes.Three = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeTwo"); err == nil && ptr != 0 {
		GCKeyCodes.Two = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeUpArrow"); err == nil && ptr != 0 {
		GCKeyCodes.UpArrow = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyCodeZero"); err == nil && ptr != 0 {
		GCKeyCodes.Zero = *(*GCKeyCode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyComma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyComma = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyD = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyDeleteForward"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyDeleteForward = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyDeleteOrBackspace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyDeleteOrBackspace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyDownArrow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyDownArrow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyEight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyEight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyEnd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyEnd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyEqualSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyEqualSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyEscape"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyEscape = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF10"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF10 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF11"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF11 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF12"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF12 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF13"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF13 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF14"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF14 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF15"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF15 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF16"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF16 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF17"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF17 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF18"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF18 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF19"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF19 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF20"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF20 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF7"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF7 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyF9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyF9 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyFive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyFive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyFour"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyFour = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyGraveAccentAndTilde"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyGraveAccentAndTilde = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyH"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyH = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyHome"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyHome = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyHyphen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyHyphen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInsert"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInsert = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational7"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational7 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyInternational9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyInternational9 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyJ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyJ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyK"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyK = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad0"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad0 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad7"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad7 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypad9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypad9 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadAsterisk"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadAsterisk = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadEnter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadEnter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadEqualSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadEqualSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadHyphen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadHyphen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadNumLock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadNumLock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadPeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadPeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadPlus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadPlus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyKeypadSlash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyKeypadSlash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG6"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG6 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG7"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG7 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG8"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG8 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLANG9"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLANG9 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLeftAlt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLeftAlt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLeftArrow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLeftArrow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLeftControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLeftControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLeftGUI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLeftGUI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyLeftShift"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyLeftShift = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyNine"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyNine = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyNonUSBackslash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyNonUSBackslash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyNonUSPound"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyNonUSPound = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyO"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyO = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyOne"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyOne = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyOpenBracket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyOpenBracket = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPageDown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPageDown = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPageUp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPageUp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPause"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPause = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPower"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPower = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyPrintScreen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyPrintScreen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyQuote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyQuote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyR = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyReturnOrEnter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyReturnOrEnter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyRightAlt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyRightAlt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyRightArrow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyRightArrow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyRightControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyRightControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyRightGUI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyRightGUI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyRightShift"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyRightShift = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyScrollLock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyScrollLock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeySemicolon"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeySemicolon = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeySeven"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeySeven = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeySix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeySix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeySlash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeySlash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeySpacebar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeySpacebar = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyT"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyT = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyTab"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyTab = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyThree"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyThree = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyTwo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyTwo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyU = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyUpArrow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyUpArrow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyW"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyW = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyX = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyZ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyZ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyZero"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyZero = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyboardDidConnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyboardDidConnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCKeyboardDidDisconnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCKeyboardDidDisconnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCMouseDidBecomeCurrentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCMouseDidBecomeCurrentNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCMouseDidConnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCMouseDidConnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCMouseDidDisconnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCMouseDidDisconnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCMouseDidStopBeingCurrentNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCMouseDidStopBeingCurrentNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCPoint2Zero"); err == nil && ptr != 0 {
		GCPoint2Zero = *(*GCPoint2)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryArcadeStick"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryArcadeStick = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryCoalescedRemote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryCoalescedRemote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryControlCenterRemote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryControlCenterRemote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryDualSense"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryDualSense = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryDualShock4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryDualShock4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryHID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryHID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryKeyboard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryKeyboard = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryMFi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryMFi = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryMouse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryMouse = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategorySiriRemote1stGen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategorySiriRemote1stGen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategorySiriRemote2ndGen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategorySiriRemote2ndGen = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategorySpatialController"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategorySpatialController = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryUniversalElectronicsRemote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryUniversalElectronicsRemote = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCProductCategoryXboxOne"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCProductCategoryXboxOne = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCRacingWheelDidConnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCRacingWheelDidConnectNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GCRacingWheelDidDisconnectNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GCRacingWheelDidDisconnectNotification = objc.GoString(cstr)
			}
		}
	}

}

// GCHapticsLocalitys provides typed accessors for [GCHapticsLocality] constants.
var GCHapticsLocalitys struct {
	// All: All locations of haptics actuators on a game controller.
	All GCHapticsLocality
	// Default: The default location of a haptics actuator on a game controller.
	Default GCHapticsLocality
	// Handles: All handles on a game controller.
	Handles GCHapticsLocality
	// LeftHandle: The left handle on a game controller.
	LeftHandle GCHapticsLocality
	// LeftTrigger: The left trigger on a game controller.
	LeftTrigger GCHapticsLocality
	// RightHandle: The right handle on a game controller.
	RightHandle GCHapticsLocality
	// RightTrigger: The right trigger on a game controller.
	RightTrigger GCHapticsLocality
	// Triggers: All triggers on a game controller.
	Triggers GCHapticsLocality
}

// GCKeyCodes provides typed accessors for [GCKeyCode] constants.
var GCKeyCodes struct {
	// Application: The keyboard code for the Application key.
	Application GCKeyCode
	// Backslash: The keyboard code for the \ or | key.
	Backslash GCKeyCode
	// CapsLock: The keyboard code for the Caps Lock key.
	CapsLock GCKeyCode
	// CloseBracket: The keyboard code for the ] or } key.
	CloseBracket GCKeyCode
	// Comma: The keyboard code for the Comma or < key.
	Comma GCKeyCode
	// DeleteForward: The keyboard code for the Delete-Forward key.
	DeleteForward GCKeyCode
	// DeleteOrBackspace: The keyboard code for the Delete or Backspace key.
	DeleteOrBackspace GCKeyCode
	// DownArrow: The keyboard code for the Down Arrow key.
	DownArrow GCKeyCode
	// Eight: The keyboard code for the 8 or * character.
	Eight GCKeyCode
	// End: The keyboard code for the End key.
	End GCKeyCode
	// EqualSign: The keyboard code for the = or + key.
	EqualSign GCKeyCode
	// Escape: The keyboard code for the Escape key.
	Escape GCKeyCode
	// F1: The keyboard code for the F1 key.
	F1 GCKeyCode
	// F10: The keyboard code for the F10 key.
	F10 GCKeyCode
	// F11: The keyboard code for the F11 key.
	F11 GCKeyCode
	// F12: The keyboard code for the F12 key.
	F12 GCKeyCode
	// F13: The keyboard code for the F13 key.
	F13 GCKeyCode
	// F14: The keyboard code for the F14 key.
	F14 GCKeyCode
	// F15: The keyboard code for the F15 key.
	F15 GCKeyCode
	// F16: The keyboard code for the F16 key.
	F16 GCKeyCode
	// F17: The keyboard code for the F17 key.
	F17 GCKeyCode
	// F18: The keyboard code for the F18 key.
	F18 GCKeyCode
	// F19: The keyboard code for the F19 key.
	F19 GCKeyCode
	// F2: The keyboard code for the F2 key.
	F2 GCKeyCode
	// F20: The keyboard code for the F20 key.
	F20 GCKeyCode
	// F3: The keyboard code for the F3 key.
	F3 GCKeyCode
	// F4: The keyboard code for the F4 key.
	F4 GCKeyCode
	// F5: The keyboard code for the F5 key.
	F5 GCKeyCode
	// F6: The keyboard code for the F6 key.
	F6 GCKeyCode
	// F7: The keyboard code for the F7 key.
	F7 GCKeyCode
	// F8: The keyboard code for the F8 key.
	F8 GCKeyCode
	// F9: The keyboard code for the F9 key.
	F9 GCKeyCode
	// Five: The keyboard code for the 5 or % character.
	Five GCKeyCode
	// Four: The keyboard code for the 4 or $ character.
	Four GCKeyCode
	// GraveAccentAndTilde: The keyboard code for the Grave Accent or Tilde key.
	GraveAccentAndTilde GCKeyCode
	// Home: The keyboard code for the Home key.
	Home GCKeyCode
	// Hyphen: The keyboard code for the - or _ key.
	Hyphen GCKeyCode
	// Insert: The keyboard code for the Insert key.
	Insert GCKeyCode
	// International1: The keyboard code for the first international key.
	International1 GCKeyCode
	// International2: The keyboard code for the second international key.
	International2 GCKeyCode
	// International3: The keyboard code for the third international key.
	International3 GCKeyCode
	// International4: The keyboard code for the fourth international key.
	International4 GCKeyCode
	// International5: The keyboard code for the fifth international key.
	International5 GCKeyCode
	// International6: The keyboard code for the sixth international key.
	International6 GCKeyCode
	// International7: The keyboard code for the seventh international key.
	International7 GCKeyCode
	// International8: The keyboard code for the eighth international key.
	International8 GCKeyCode
	// International9: The keyboard code for the ninth international key.
	International9 GCKeyCode
	// KeyA: The keyboard code for the a or A character.
	KeyA GCKeyCode
	// KeyB: The keyboard code for the b or B character.
	KeyB GCKeyCode
	// KeyC: The keyboard code for the c or C character.
	KeyC GCKeyCode
	// KeyD: The keyboard code for the d or D character.
	KeyD GCKeyCode
	// KeyE: The keyboard code for the e or E character.
	KeyE GCKeyCode
	// KeyF: The keyboard code for the f or F character.
	KeyF GCKeyCode
	// KeyG: The keyboard code for the g or G character.
	KeyG GCKeyCode
	// KeyH: The keyboard code for the h or H character.
	KeyH GCKeyCode
	// KeyI: The keyboard code for the i or I character.
	KeyI GCKeyCode
	// KeyJ: The keyboard code for the j or J character.
	KeyJ GCKeyCode
	// KeyK: The keyboard code for the k or K character.
	KeyK GCKeyCode
	// KeyL: The keyboard code for the l or L character.
	KeyL GCKeyCode
	// KeyM: The keyboard code for the m or M character.
	KeyM GCKeyCode
	// KeyN: The keyboard code for the n or N character.
	KeyN GCKeyCode
	// KeyO: The keyboard code for the o or O character.
	KeyO GCKeyCode
	// KeyP: The keyboard code for the p or P character.
	KeyP GCKeyCode
	// KeyQ: The keyboard code for the q or Q character.
	KeyQ GCKeyCode
	// KeyR: The keyboard code for the r or R character.
	KeyR GCKeyCode
	// KeyS: The keyboard code for the s or S character.
	KeyS GCKeyCode
	// KeyT: The keyboard code for the t or T character.
	KeyT GCKeyCode
	// KeyU: The keyboard code for the u or U character.
	KeyU GCKeyCode
	// KeyV: The keyboard code for the v or V character.
	KeyV GCKeyCode
	// KeyW: The keyboard code for the w or W character.
	KeyW GCKeyCode
	// KeyX: The keyboard code for the x or X character.
	KeyX GCKeyCode
	// KeyY: The keyboard code for the y or Y character.
	KeyY GCKeyCode
	// KeyZ: The keyboard code for the z or Z character.
	KeyZ GCKeyCode
	// Keypad0: The keyboard code for the keypad 0 or Insert key.
	Keypad0 GCKeyCode
	// Keypad1: The keyboard code for the keypad 1 or End key.
	Keypad1 GCKeyCode
	// Keypad2: The keyboard code for the keypad 2 or Down Arrow key.
	Keypad2 GCKeyCode
	// Keypad3: The keyboard code for the keypad 3 or Page Down key.
	Keypad3 GCKeyCode
	// Keypad4: The keyboard code for the keypad 4 or Left Arrow key.
	Keypad4 GCKeyCode
	// Keypad5: The keyboard code for the keypad 5 key.
	Keypad5 GCKeyCode
	// Keypad6: The keyboard code for the keypad 6 or Right Arrow key.
	Keypad6 GCKeyCode
	// Keypad7: The keyboard code for the keypad 7 or Home key.
	Keypad7 GCKeyCode
	// Keypad8: The keyboard code for the keypad 8 or Up Arrow key.
	Keypad8 GCKeyCode
	// Keypad9: The keyboard code for the keypad 9 or Page Up key.
	Keypad9 GCKeyCode
	// KeypadAsterisk: The keyboard code for the keypad * key.
	KeypadAsterisk GCKeyCode
	// KeypadEnter: The keyboard code for the keypad Enter key.
	KeypadEnter GCKeyCode
	// KeypadEqualSign: The keyboard code for the keypad = key.
	KeypadEqualSign GCKeyCode
	// KeypadHyphen: The keyboard code for the keypad - key.
	KeypadHyphen GCKeyCode
	// KeypadNumLock: The keyboard code for the keypad Num Lock or Clear key.
	KeypadNumLock GCKeyCode
	// KeypadPeriod: The keyboard code for the keypad Period or Delete key.
	KeypadPeriod GCKeyCode
	// KeypadPlus: The keyboard code for the keypad + key.
	KeypadPlus GCKeyCode
	// KeypadSlash: The keyboard code for the keypad / key.
	KeypadSlash GCKeyCode
	// LANG1: The keyboard code for the first language key.
	LANG1 GCKeyCode
	// LANG2: The keyboard code for the second language key.
	LANG2 GCKeyCode
	// LANG3: The keyboard code for the third language key.
	LANG3 GCKeyCode
	// LANG4: The keyboard code for the fourth language key.
	LANG4 GCKeyCode
	// LANG5: The keyboard code for the fifth language key.
	LANG5 GCKeyCode
	// LANG6: The keyboard code for the sixth language key.
	LANG6 GCKeyCode
	// LANG7: The keyboard code for the seventh language key.
	LANG7 GCKeyCode
	// LANG8: The keyboard code for the eighth language key.
	LANG8 GCKeyCode
	// LANG9: The keyboard code for the ninth language key.
	LANG9 GCKeyCode
	// LeftAlt: The keyboard code for the Option or Alt key on the left side of the keyboard.
	LeftAlt GCKeyCode
	// LeftArrow: The keyboard code for the Left Arrow key.
	LeftArrow GCKeyCode
	// LeftControl: The keyboard code for the Control key on the left side of the keyboard.
	LeftControl GCKeyCode
	// LeftGUI: The keyboard code for the Command key on the left side of the keyboard.
	LeftGUI GCKeyCode
	// LeftShift: The keyboard code for the Shift key on the left side of the keyboard.
	LeftShift GCKeyCode
	// Nine: The keyboard code for the 9 or ( character.
	Nine GCKeyCode
	// NonUSBackslash: The keyboard code for the non-US Slash or | key.
	NonUSBackslash GCKeyCode
	// NonUSPound: The keyboard code for the non-US Pound or _ key.
	NonUSPound GCKeyCode
	// One: The keyboard code for the 1 or ! character.
	One GCKeyCode
	// OpenBracket: The keyboard code for the [ or { key.
	OpenBracket GCKeyCode
	// PageDown: The keyboard code for the Page Down key.
	PageDown GCKeyCode
	// PageUp: The keyboard code for the Page Up key.
	PageUp GCKeyCode
	// Pause: The keyboard code for the Pause key.
	Pause GCKeyCode
	// Period: The keyboard code for the Period or > key.
	Period GCKeyCode
	// Power: The keyboard code for the Power key.
	Power GCKeyCode
	// PrintScreen: The keyboard code for the Print Screen key.
	PrintScreen GCKeyCode
	// Quote: The keyboard code for the ’ or “ key.
	Quote GCKeyCode
	// ReturnOrEnter: The keyboard code for the Return or Enter key.
	ReturnOrEnter GCKeyCode
	// RightAlt: The keyboard code for the Option or Alt key on the right side of the keyboard.
	RightAlt GCKeyCode
	// RightArrow: The keyboard code for the Right Arrow key.
	RightArrow GCKeyCode
	// RightControl: The keyboard code for the Control key on the right side of the keyboard.
	RightControl GCKeyCode
	// RightGUI: The keyboard code for the Command key on the right side of the keyboard.
	RightGUI GCKeyCode
	// RightShift: The keyboard code for the Shift key on the right side of the keyboard.
	RightShift GCKeyCode
	// ScrollLock: The keyboard code for the Scroll Lock key.
	ScrollLock GCKeyCode
	// Semicolon: The keyboard code for the ; or : key.
	Semicolon GCKeyCode
	// Seven: The keyboard code for the 7 or & character.
	Seven GCKeyCode
	// Six: The keyboard code for the 6 or ^ character.
	Six GCKeyCode
	// Slash: The keyboard code for the / or ? key.
	Slash GCKeyCode
	// Spacebar: The keyboard code for the Space Bar key.
	Spacebar GCKeyCode
	// Tab: The keyboard code for the Tab key.
	Tab GCKeyCode
	// Three: The keyboard code for the 3 or # character.
	Three GCKeyCode
	// Two: The keyboard code for the 2 or @ character.
	Two GCKeyCode
	// UpArrow: The keyboard code for the Up Arrow key.
	UpArrow GCKeyCode
	// Zero: The keyboard code for the 0 or ) character.
	Zero GCKeyCode
}
