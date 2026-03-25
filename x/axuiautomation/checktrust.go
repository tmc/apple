package axuiautomation

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/screencapturekit"
)

var checkTrustOnce sync.Once

// IsTrustedFresh performs a live TCC query for accessibility permission.
// Pass prompt=true to trigger the system permission dialog if not trusted.
// Unlike AXIsProcessTrusted(), this always queries the current TCC state.
// Must be called from the main thread (ObjC runtime requirement).
func IsTrustedFresh(prompt bool) bool {
	key := objc.Send[uintptr](objc.ID(objc.GetClass("NSString")), objc.Sel("stringWithUTF8String:"), "AXTrustedCheckOptionPrompt\x00")
	var boolVal bool
	if prompt {
		boolVal = true
	}
	val := objc.Send[uintptr](objc.ID(objc.GetClass("NSNumber")), objc.Sel("numberWithBool:"), boolVal)
	dict := objc.Send[uintptr](objc.ID(objc.GetClass("NSDictionary")), objc.Sel("dictionaryWithObject:forKey:"), val, key)
	return AXIsProcessTrustedWithOptions(dict)
}

// execName returns the base name of the running executable, stripping any
// .app bundle path components (e.g. "axmcp.app/Contents/MacOS/axmcp" → "axmcp").
func execName() string {
	exe, err := os.Executable()
	if err != nil {
		return "axmcp"
	}
	name := filepath.Base(exe)
	name = strings.TrimSuffix(name, ".app")
	return name
}

// CheckTrust checks if the process has Accessibility permission and, if not,
// shows a floating window with a spinner and buttons to open System Settings
// or reset TCC. It polls until permission is granted, then briefly shows a
// green checkmark before closing.
func CheckTrust() {
	checkTrustOnce.Do(func() {
		// Check without triggering the system prompt.
		if AXIsProcessTrusted() {
			return
		}
		showWaitingForPermissionWindow()
	})
}

// CheckTrustWithPrompt is like CheckTrust but also triggers the system
// universalAccessAuthWarn permission dialog. Prefer CheckTrust for a less
// disruptive experience.
func CheckTrustWithPrompt() {
	checkTrustOnce.Do(func() {
		if IsTrustedFresh(true) {
			return
		}
		showWaitingForPermissionWindow()
	})
}

func openAccessibilityPrefs() {
	exec.Command("open", "x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility").Start()
	// Re-trigger the system prompt as requested by the user.
	IsTrustedFresh(true)
}

func resetTCC() {
	// Derive bundle ID from the running app bundle's Info.plist if possible,
	// falling back to a pattern based on execName.
	bundleID := infoPlistBundleID()
	if bundleID == "" {
		bundleID = "dev.tmc." + execName()
	}
	exec.Command("tccutil", "reset", "Accessibility", bundleID).Run()
	// Re-query with prompt=true to re-register the entry in TCC and
	// triggering the system universalAccessAuthWarn popup as requested.
	IsTrustedFresh(true)
	openAccessibilityPrefs()
}

// infoPlistBundleID reads CFBundleIdentifier from the running app bundle.
func infoPlistBundleID() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	// Walk up from Contents/MacOS/<exe> to Contents/Info.plist.
	plist := filepath.Join(filepath.Dir(filepath.Dir(exe)), "Info.plist")
	out, err := exec.Command("defaults", "read", plist, "CFBundleIdentifier").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// bindButtonAction wires a Go callback to an NSButton.
func bindButtonAction(btn appkit.NSButton, fn func()) {
	btn.SetActionHandler(fn)
}

func makeButton(title string, frame corefoundation.CGRect, fn func()) appkit.NSButton {
	btn := appkit.NewButtonWithFrame(frame)
	btn.SetTitle(title)
	btn.SetBezelStyle(appkit.NSBezelStyleAccessoryBar)
	bindButtonAction(btn, fn)
	return btn
}

func IsScreenRecordingTrusted() bool {
	scClass := screencapturekit.GetSCShareableContentClass()
	done := make(chan bool)
	handler := func(content *screencapturekit.SCShareableContent, err error) {
		done <- content != nil && err == nil
	}
	scClass.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler(true, true, handler)
	select {
	case result := <-done:
		return result
	case <-time.After(2 * time.Second):
		return false
	}
}

func checkScreenCaptureProcessTrusted(prompt bool) bool {
	trusted := IsScreenRecordingTrusted()
	if !trusted && prompt {
		coregraphics.CGRequestScreenCaptureAccess()
	}
	return trusted
}

func openScreenRecordingPrefs() {
	exec.Command("open", "x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture").Start()
	checkScreenCaptureProcessTrusted(true)
}

func resetScreenRecordingTCC() {
	bundleID := infoPlistBundleID()
	if bundleID == "" {
		bundleID = "dev.tmc." + execName()
	}
	exec.Command("tccutil", "reset", "ScreenCapture", bundleID).Run()
	checkScreenCaptureProcessTrusted(true)
	openScreenRecordingPrefs()
}

var checkScreenCaptureOnce sync.Once

func CheckScreenCapture() {
	checkScreenCaptureOnce.Do(func() {
		// Check without triggering the system prompt.
		if IsScreenRecordingTrusted() {
			return
		}

		dispatch.MainQueue().Async(func() {
			app := appkit.GetNSApplicationClass().SharedApplication()
			name := execName()

			// Check again on main thread, maybe prompting this time.
			if checkScreenCaptureProcessTrusted(true) {
				return
			}

			const w = 480.0
			const h = 180.0
			const padding = 20.0
			const spinSz = 24.0
			const titleH = 22.0
			const subtitleH = 40.0
			const btnH = 24.0
			const btnGap = 8.0

			win := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
				corefoundation.CGRect{
					Origin: corefoundation.CGPoint{X: 0, Y: 0},
					Size:   corefoundation.CGSize{Width: w, Height: h},
				},
				appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskFullSizeContentView,
				appkit.NSBackingStoreBuffered,
				false,
			)
			win.SetTitle(name + " — Screen Recording Permission Required")
			win.SetLevel(appkit.FloatingWindowLevel)

			content := appkit.NSViewFromID(win.ContentView().GetID())

			spinner := appkit.NewProgressIndicatorWithFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: padding, Y: h - padding - spinSz - 4},
				Size:   corefoundation.CGSize{Width: spinSz, Height: spinSz},
			})
			spinner.SetStyle(appkit.NSProgressIndicatorStyleSpinning)
			spinner.SetIndeterminate(true)
			spinner.StartAnimation(nil)
			content.AddSubview(spinner)

			labelX := padding + spinSz + padding
			labelW := w - labelX - padding

			titleLabel := appkit.NewTextFieldLabelWithString(
				`"` + name + `.app" requires screen recording permissions to capture images.`,
			)
			titleLabel.SetFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: labelX, Y: h - padding - titleH},
				Size:   corefoundation.CGSize{Width: labelW, Height: titleH},
			})
			content.AddSubview(titleLabel)

			subtitleLabel := appkit.NewTextFieldLabelWithString(
				"Grant access in Privacy & Security settings, located in System Settings.",
			)
			subtitleLabel.SetFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: labelX, Y: h - padding - titleH - subtitleH - 2},
				Size:   corefoundation.CGSize{Width: labelW, Height: subtitleH},
			})
			content.AddSubview(subtitleLabel)

			openW := labelW * 0.50
			smallW := (labelW - openW - (2 * btnGap)) / 2

			openBtn := makeButton("Open System Settings…", corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: labelX, Y: padding},
				Size:   corefoundation.CGSize{Width: openW, Height: btnH},
			}, openScreenRecordingPrefs)
			content.AddSubview(openBtn)

			retryBtn := makeButton("Retry", corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: labelX + openW + btnGap, Y: padding},
				Size:   corefoundation.CGSize{Width: smallW, Height: btnH},
			}, func() { checkScreenCaptureProcessTrusted(true) })
			content.AddSubview(retryBtn)

			resetBtn := makeButton("Reset TCC…", corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: labelX + openW + smallW + (2 * btnGap), Y: padding},
				Size:   corefoundation.CGSize{Width: smallW, Height: btnH},
			}, resetScreenRecordingTCC)
			content.AddSubview(resetBtn)

			win.Center()
			win.MakeKeyAndOrderFront(nil)
			app.Activate()

			var pollTimer *time.Timer
			var poll func()
			poll = func() {
				if !checkScreenCaptureProcessTrusted(false) {
					pollTimer = time.AfterFunc(1000*time.Millisecond, func() {
						dispatch.MainQueue().Async(poll)
					})
					_ = pollTimer // suppress declared and not used
					return
				}

				spinner.StopAnimation(nil)
				spinner.SetIsHidden(true)
				openBtn.SetIsHidden(true)
				retryBtn.SetIsHidden(true)
				resetBtn.SetIsHidden(true)
				subtitleLabel.SetIsHidden(true)

				const checkSz = 36.0
				baseImg := appkit.NewImageWithSystemSymbolNameAccessibilityDescription(
					"checkmark.circle.fill", "Permission granted",
				)
				sizeCfg := appkit.NewImageSymbolConfigurationWithPointSizeWeight(checkSz, appkit.NSFontWeights.Medium)
				colorCfg := appkit.NewImageSymbolConfigurationWithHierarchicalColor(
					appkit.GetNSColorClass().SystemGreen(),
				)
				cfg := sizeCfg.ConfigurationByApplyingConfiguration(colorCfg)
				checkImg := appkit.NSImageFromID(baseImg.ImageWithSymbolConfiguration(cfg).GetID())
				checkView := appkit.NewImageViewWithFrame(corefoundation.CGRect{
					Origin: corefoundation.CGPoint{X: padding - 4, Y: (h - checkSz) / 2},
					Size:   corefoundation.CGSize{Width: checkSz, Height: checkSz},
				})
				checkView.SetImage(checkImg)
				content.AddSubview(checkView)
				titleLabel.SetStringValue("Screen Recording permission granted.")

				time.AfterFunc(1200*time.Millisecond, func() {
					dispatch.MainQueue().Async(func() {
						win.Close()
					})
				})
			}
			poll() // start polling
		})
	})
}

// showWaitingForPermissionWindow shows a floating panel with a spinner,
// the app name, and buttons to open System Settings or reset TCC.
// It polls until permission is granted, then briefly shows a green
// checkmark before closing.
func showWaitingForPermissionWindow() {
	dispatch.MainQueue().Async(func() {
		app := appkit.GetNSApplicationClass().SharedApplication()
		app.SetActivationPolicy(appkit.NSApplicationActivationPolicyAccessory)

		const (
			w         = 400.0
			h         = 148.0
			padding   = 16.0
			spinSz    = 24.0
			btnH      = 22.0
			btnGap    = 6.0
			titleH    = 34.0
			subtitleH = 38.0
		)

		name := execName()

		win := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			corefoundation.CGRect{Size: corefoundation.CGSize{Width: w, Height: h}},
			appkit.NSWindowStyleMaskTitled,
			appkit.NSBackingStoreBuffered,
			false,
		)
		win.SetTitle(name + " — Accessibility Permission Required")
		win.SetLevel(appkit.FloatingWindowLevel)

		content := appkit.NSViewFromID(win.ContentView().GetID())

		// Spinner — left side, vertically centred in text area.
		spinner := appkit.NewProgressIndicatorWithFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: padding, Y: h - padding - spinSz - 4},
			Size:   corefoundation.CGSize{Width: spinSz, Height: spinSz},
		})
		spinner.SetStyle(appkit.NSProgressIndicatorStyleSpinning)
		spinner.SetIndeterminate(true)
		spinner.StartAnimation(nil)
		content.AddSubview(spinner)

		labelX := padding + spinSz + padding
		labelW := w - labelX - padding

		// Bold title line.
		titleLabel := appkit.NewTextFieldLabelWithString(
			`"` + name + `.app" would like to control this computer using accessibility features.`,
		)
		titleLabel.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: labelX, Y: h - padding - titleH},
			Size:   corefoundation.CGSize{Width: labelW, Height: titleH},
		})
		content.AddSubview(titleLabel)

		// Secondary description.
		subtitleLabel := appkit.NewTextFieldLabelWithString(
			"Grant access in Privacy & Security settings, located in System Settings.",
		)
		subtitleLabel.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: labelX, Y: h - padding - titleH - subtitleH - 2},
			Size:   corefoundation.CGSize{Width: labelW, Height: subtitleH},
		})
		content.AddSubview(subtitleLabel)

		// Three buttons in the lower area: Open (large), Retry (small), Reset (small).
		// Distribute labelW across: openW, gap, retryW, gap, resetW
		// Make openW take ~50%
		openW := labelW * 0.50
		smallW := (labelW - openW - (2 * btnGap)) / 2

		openBtn := makeButton("Open System Settings…", corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: labelX, Y: padding},
			Size:   corefoundation.CGSize{Width: openW, Height: btnH},
		}, openAccessibilityPrefs)
		content.AddSubview(openBtn)

		retryBtn := makeButton("Retry", corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: labelX + openW + btnGap, Y: padding},
			Size:   corefoundation.CGSize{Width: smallW, Height: btnH},
		}, func() { IsTrustedFresh(true) })
		content.AddSubview(retryBtn)

		resetBtn := makeButton("Reset TCC…", corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: labelX + openW + smallW + (2 * btnGap), Y: padding},
			Size:   corefoundation.CGSize{Width: smallW, Height: btnH},
		}, resetTCC)
		content.AddSubview(resetBtn)

		win.Center()
		win.MakeKeyAndOrderFront(nil)
		app.Activate()

		// Poll on the main thread via time.AfterFunc + DispatchMainSafe so all
		// ObjC calls (including IsTrustedFresh and UI mutations) stay on the
		// main thread.
		var pollTimer *time.Timer
		var poll func()
		poll = func() {
			// Use IsTrustedFresh (live TCC query) — AXIsProcessTrusted() may cache.
			if !IsTrustedFresh(false) {
				pollTimer = time.AfterFunc(500*time.Millisecond, func() {
					dispatch.MainQueue().Async(poll)
				})
				return
			}
			// Permission granted — transition to success state.
			spinner.StopAnimation(nil)
			spinner.SetIsHidden(true)
			openBtn.SetIsHidden(true)
			retryBtn.SetIsHidden(true)
			resetBtn.SetIsHidden(true)
			subtitleLabel.SetIsHidden(true)

			const checkSz = 36.0
			baseImg := appkit.NewImageWithSystemSymbolNameAccessibilityDescription(
				"checkmark.circle.fill", "Permission granted",
			)
			sizeCfg := appkit.NewImageSymbolConfigurationWithPointSizeWeight(checkSz, appkit.NSFontWeights.Medium)
			colorCfg := appkit.NewImageSymbolConfigurationWithHierarchicalColor(
				appkit.GetNSColorClass().SystemGreen(),
			)
			cfg := sizeCfg.ConfigurationByApplyingConfiguration(colorCfg)
			checkImg := appkit.NSImageFromID(baseImg.ImageWithSymbolConfiguration(cfg).GetID())
			checkView := appkit.NewImageViewWithFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: padding - 4, Y: (h - checkSz) / 2},
				Size:   corefoundation.CGSize{Width: checkSz, Height: checkSz},
			})
			checkView.SetImage(checkImg)
			content.AddSubview(checkView)
			titleLabel.SetStringValue("Accessibility permission granted.")

			time.AfterFunc(1200*time.Millisecond, func() {
				dispatch.MainQueue().Async(func() {
					win.Close()
				})
			})
		}

		// Kick off first poll after a short delay.
		pollTimer = time.AfterFunc(500*time.Millisecond, func() {
			dispatch.MainQueue().Async(poll)
		})
		_ = pollTimer
	})
}
