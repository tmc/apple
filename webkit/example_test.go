//go:build darwin

package webkit_test

import (
	"fmt"

	"github.com/tmc/apple/webkit"
)

func ExampleWKWebViewConfiguration() {
	config := webkit.NewWKWebViewConfiguration()
	fmt.Printf("Initial AirPlay: %t\n", config.AllowsAirPlayForMediaPlayback())
	config.SetAllowsAirPlayForMediaPlayback(false)
	fmt.Printf("Updated AirPlay: %t\n", config.AllowsAirPlayForMediaPlayback())

	// Output:
	// Initial AirPlay: true
	// Updated AirPlay: false
}

func ExampleWKPreferences() {
	prefs := webkit.NewWKPreferences()
	prefs.SetMinimumFontSize(14)
	prefs.SetJavaScriptCanOpenWindowsAutomatically(false)

	fmt.Println(prefs.MinimumFontSize())
	fmt.Println(prefs.JavaScriptCanOpenWindowsAutomatically())

	// Output:
	// 14
	// false
}
