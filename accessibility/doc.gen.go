// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

// Package accessibility provides Go bindings for the Accessibility framework.
//
// Make your apps accessible to everyone who uses Apple devices.
//
// Accessibility features help a wide range of people interact with their
// devices. By creating your app with accessibility in mind, you make it
// possible for everyone to enjoy your app. Whether you’re developing a new
// app, or updating an existing one, consider the needs of all the people who
// might use your app.
//
// # Essentials
//
//   - [Accessibility updates]: Learn about important changes to Accessibility.
//   - [Performing accessibility testing for your app]: Test your app with accessibility settings and assistive technologies to discover and address accessibility issues.
//
// # Sample code
//
//   - [Creating accessible views]: Make your app accessible to everyone by applying accessibility modifiers to your SwiftUI views.
//   - [Delivering an exceptional accessibility experience]: Make improvements to your app’s interaction model to support assistive technologies such as VoiceOver.
//   - [Integrating accessibility into your app]: Make your app more accessible to users with disabilities by adding accessibility features.
//   - [Accessibility design for Mac Catalyst]: Improve navigation in your app by using keyboard shortcuts and accessibility containers.
//
// # Developer tools
//
//   - [Accessibility Inspector]: Reveal how your app represents itself to people using accessibility features.
//
// # Assistive technologies
//
//   - [Assistive technologies]: Make sure your app provides a great experience for people who use assistive technologies.
//
// # Accessibility framework
//
//   - [Accessibility API]: Browse API in the Accessibility framework. ([AXBrailleTable], [AXBrailleTranslator], [AXBrailleTranslationResult], [AXMathExpressionNumber], [AXMathExpressionIdentifier])
//
// # Platforms
//
//   - [Accessibility fundamentals]: Make your SwiftUI apps accessible to everyone, including people with disabilities.
//   - [Accessibility for UIKit]: Make your UIKit apps accessible to everyone who uses iOS and tvOS.
//   - [Accessibility for AppKit]: Make your AppKit apps accessible to everyone who uses macOS.
//   - [Accessibility for visionOS]: Make your apps accessible to everyone who uses visionOS.
//
// # WWDC Challenges
//
//   - [WWDC22 Challenge: Learn Switch Control through gaming]: Play a card-matching game using Switch Control.
//   - [WWDC21 Challenge: Large Text Challenge]: Design for large text sizes by modifying the user interface.
//   - [WWDC21 Challenge: Speech Synthesizer Simulator]: Simulate a conversation using speech synthesis.
//   - [WWDC21 Challenge: VoiceOver Maze]: Navigate to the end of a dark maze using VoiceOver as your guide.
//
// # Variables
//
//   - [AXPrefersActionSliderAlternativeDidChangeNotification]
//   - [AXReduceHighlightingEffectsEnabledDidChangeNotification]
//   - [AXShowBordersEnabledStatusDidChangeNotification]
//
// # Functions
//
//   - [AXOpenSettingsFeatureIsSupported]
//   - [AXPrefersActionSliderAlternative]
//   - [AXShowBordersEnabled]
//   - [AXReduceHighlightingEffectsEnabled]//
//
// # Key Types
//
//   - [AXBrailleTable] - A rule for translating print text to Braille, and back-translating Braille to print text.
//   - [AXChartDescriptor] - An object that contains all the semantic information about an accessible chart.
//   - [AXNumericDataAxisDescriptor] - An object that represents an axis of numerical data.
//   - [AXCustomContent] - Objects that define custom content and the timing of its output.
//   - [AXDataPoint] - An object that represents a single data point in a chart.
//   - [AXCategoricalDataAxisDescriptor] - An object that represents an axis of categorical data.
//   - [AXDataSeriesDescriptor] - An object that represents a series of data points.
//   - [AXBrailleMap] - A representation of a two-dimensional braille display.
//   - [AXDataPointValue] - A single data value.
//   - [AXMathExpressionFenced]
//
// [Accessibility API]: https://developer.apple.com/documentation/accessibility/accessibility-api
// [Accessibility Inspector]: https://developer.apple.com/documentation/accessibility/accessibility-inspector
// [Accessibility design for Mac Catalyst]: https://developer.apple.com/documentation/accessibility/accessibility_design_for_mac_catalyst
// [Accessibility for AppKit]: https://developer.apple.com/documentation/AppKit/accessibility-for-appkit
// [Accessibility for UIKit]: https://developer.apple.com/documentation/UIKit/accessibility-for-uikit
// [Accessibility for visionOS]: https://developer.apple.com/documentation/accessibility/accessibility-for-visionos
// [Accessibility fundamentals]: https://developer.apple.com/documentation/SwiftUI/Accessibility-fundamentals
// [Accessibility updates]: https://developer.apple.com/documentation/Updates/Accessibility
// [Assistive technologies]: https://developer.apple.com/documentation/accessibility/assistive-technologies
// [Creating accessible views]: https://developer.apple.com/documentation/SwiftUI/creating-accessible-views
// [Delivering an exceptional accessibility experience]: https://developer.apple.com/documentation/accessibility/delivering_an_exceptional_accessibility_experience
// [Integrating accessibility into your app]: https://developer.apple.com/documentation/accessibility/integrating-accessibility-into-your-app
// [Performing accessibility testing for your app]: https://developer.apple.com/documentation/accessibility/performing-accessibility-testing-for-your-app
// [WWDC21 Challenge: Large Text Challenge]: https://developer.apple.com/documentation/accessibility/wwdc21_challenge_large_text_challenge
// [WWDC21 Challenge: Speech Synthesizer Simulator]: https://developer.apple.com/documentation/accessibility/wwdc21_challenge_speech_synthesizer_simulator
// [WWDC21 Challenge: VoiceOver Maze]: https://developer.apple.com/documentation/accessibility/wwdc21_challenge_voiceover_maze
// [WWDC22 Challenge: Learn Switch Control through gaming]: https://developer.apple.com/documentation/accessibility/wwdc22_challenge_learn_switch_control_through_gaming
package accessibility

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Accessibility library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Accessibility.framework/Accessibility",
	"/usr/lib/libAccessibility.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: Accessibility: failed to load framework from any known path\n")
	}
}
