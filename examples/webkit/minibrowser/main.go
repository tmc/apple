// Command minibrowser opens a window containing a WKWebView, loads a URL, and
// runs JavaScript in the page once the load finishes.
//
// The JavaScript result is printed to stdout, so the program doubles as a small
// demonstration of the Go-to-page bridge:
//
//	minibrowser https://example.com "document.title"
//
// Usage: minibrowser [url] [javascript]
//
// The URL defaults to https://example.com and the JavaScript expression
// defaults to document.title. Closing the window quits the program.
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/webkit"
)

func main() {
	urlString := "https://example.com"
	script := "document.title"
	if len(os.Args) > 1 {
		urlString = os.Args[1]
	}
	if len(os.Args) > 2 {
		script = os.Args[2]
	}

	url := foundation.NewURLWithString(urlString)
	if url.GetID() == 0 {
		fmt.Fprintf(os.Stderr, "minibrowser: not a valid URL: %s\n", urlString)
		os.Exit(1)
	}

	appkit.RunApp(func(app appkit.NSApplication, _ appkit.NSApplicationDelegateObject) {
		frame := corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: 0},
			Size:   corefoundation.CGSize{Width: 1000, Height: 700},
		}

		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			frame,
			appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable|appkit.NSWindowStyleMaskResizable,
			appkit.NSBackingStoreBuffered,
			false,
		)
		window.SetTitle(urlString)

		config := webkit.NewWKWebViewConfiguration()
		webView := webkit.NewWebViewWithFrameConfiguration(frame, config)
		webView.SetAllowsBackForwardNavigationGestures(true)
		// Makes the page reachable from Safari's Develop menu.
		webView.SetInspectable(true)

		// The navigation delegate is where the JavaScript bridge hangs off: the
		// page is only ready to be scripted after didFinishNavigation.
		delegate := webkit.NewWKNavigationDelegate(webkit.WKNavigationDelegateConfig{
			WebViewDidFinishNavigation: func(w webkit.WKWebView, _ webkit.WKNavigation) {
				window.SetTitle(w.Title())
				w.EvaluateJavaScriptCompletionHandler(script, func(result objectivec.IObject, err error) {
					if err != nil {
						fmt.Fprintf(os.Stderr, "minibrowser: javascript failed: %v\n", err)
						return
					}
					if result == nil || result.GetID() == 0 {
						fmt.Println("<undefined>")
						return
					}
					fmt.Println(objectivec.ObjectFromID(result.GetID()).Description())
				})
			},
			WebViewDidFailNavigationWithError: func(_ webkit.WKWebView, _ webkit.WKNavigation, e foundation.NSError) {
				fmt.Fprintf(os.Stderr, "minibrowser: navigation failed: %s\n", e.LocalizedDescription())
			},
			WebViewDidFailProvisionalNavigationWithError: func(_ webkit.WKWebView, _ webkit.WKNavigation, e foundation.NSError) {
				fmt.Fprintf(os.Stderr, "minibrowser: could not load %s: %s\n", urlString, e.LocalizedDescription())
			},
		})
		webView.SetNavigationDelegate(delegate)

		window.SetContentView(webView)
		window.Center()
		window.MakeKeyAndOrderFront(nil)
		app.Activate()

		webView.LoadRequest(foundation.NewURLRequestWithURL(url))
	})
}
