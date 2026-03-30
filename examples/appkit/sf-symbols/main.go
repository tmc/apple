// Command sf-symbols shows a small grid of SF Symbols in an AppKit window.
//
// It builds on hello-window and counter by adding image views, nested stack
// views, and per-item styling without introducing custom drawing.
package main

import (
	"runtime"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// symbols is a curated set of SF Symbol names grouped by category.
var symbols = []struct {
	name  string
	label string
}{
	{"star.fill", "Favorites"},
	{"heart.fill", "Health"},
	{"bell.fill", "Alerts"},
	{"gear", "Settings"},
	{"magnifyingglass", "Search"},
	{"folder.fill", "Files"},
	{"doc.text.fill", "Documents"},
	{"person.fill", "Profile"},
	{"lock.fill", "Security"},
	{"wifi", "Network"},
	{"battery.100", "Battery"},
	{"cpu", "Processor"},
	{"memorychip", "Memory"},
	{"externaldrive.fill", "Storage"},
	{"desktopcomputer", "Desktop"},
	{"macbook", "Laptop"},
}

func init() { runtime.LockOSThread() }

func main() {
	app := appkit.GetNSApplicationClass().SharedApplication()
	app.SetActivationPolicy(appkit.NSApplicationActivationPolicyRegular)

	delegate := appkit.NewNSApplicationDelegate(appkit.NSApplicationDelegateConfig{
		ShouldTerminateAfterLastWindowClosed: func(_ appkit.NSApplication) bool { return true },
		DidFinishLaunching: func(_ foundation.NSNotification) {
			buildWindow(app)
		},
	})
	app.SetDelegate(delegate)
	app.Run()
}

func buildWindow(app appkit.NSApplication) {
	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 200, Y: 200},
			Size:   corefoundation.CGSize{Width: 520, Height: 460},
		},
		appkit.NSWindowStyleMaskTitled|
			appkit.NSWindowStyleMaskClosable|
			appkit.NSWindowStyleMaskMiniaturizable|
			appkit.NSWindowStyleMaskResizable,
		appkit.NSBackingStoreBuffered, false,
	)
	window.SetTitle("SF Symbols")
	window.SetMinSize(corefoundation.CGSize{Width: 360, Height: 320})

	contentView := window.ContentView().(appkit.NSView)
	font := appkit.GetNSFontClass()
	colors := appkit.GetNSColorClass()

	// Title label.
	title := appkit.NewTextFieldLabelWithString("SF Symbols from Go")
	title.SetFont(font.SystemFontOfSizeWeight(18, appkit.NSFontWeights.Semibold))

	// Subtitle label.
	subtitle := appkit.NewTextFieldLabelWithString("NSImage + NSImageSymbolConfiguration + Symbols framework")
	subtitle.SetFont(font.SystemFontOfSize(12))
	subtitle.SetTextColor(colors.SecondaryLabelColor())

	// Three symbol configurations at different scales.
	configSmall := appkit.NewImageSymbolConfigurationWithPointSizeWeight(16, appkit.NSFontWeights.Regular)
	configMedium := appkit.NewImageSymbolConfigurationWithPointSizeWeight(24, appkit.NSFontWeights.Medium)
	configLarge := appkit.NewImageSymbolConfigurationWithPointSizeWeight(32, appkit.NSFontWeights.Light)

	// Tint colors for visual variety.
	tintColors := []appkit.INSColor{
		colors.SystemBlueColor(),
		colors.SystemGreenColor(),
		colors.SystemOrangeColor(),
		colors.SystemPurpleColor(),
		colors.SystemRedColor(),
		colors.SystemTealColor(),
		colors.SystemIndigoColor(),
		colors.SystemPinkColor(),
	}

	layout := appkit.GetNSStackViewClass().Alloc().Init()
	layout.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
	layout.SetSpacing(16)
	layout.SetEdgeInsets(foundation.NSEdgeInsets{Top: 20, Left: 20, Bottom: 20, Right: 20})
	layout.SetTranslatesAutoresizingMaskIntoConstraints(false)
	layout.AddArrangedSubview(title)
	layout.AddArrangedSubview(subtitle)

	grid := appkit.GetNSStackViewClass().Alloc().Init()
	grid.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
	grid.SetSpacing(8)
	layout.AddArrangedSubview(grid)

	cols := 4
	for row := 0; row < len(symbols); row += cols {
		rowStack := appkit.GetNSStackViewClass().Alloc().Init()
		rowStack.SetOrientation(appkit.NSUserInterfaceLayoutOrientationHorizontal)
		rowStack.SetDistribution(appkit.NSStackViewDistributionFillEqually)
		rowStack.SetSpacing(8)

		for col := 0; col < cols; col++ {
			i := row + col
			sym := symbols[i]
			rowIndex := row / cols
			var config appkit.NSImageSymbolConfiguration
			size := 28.0
			switch rowIndex {
			case 0:
				config = configLarge
				size = 40
			case 1:
				config = configMedium
				size = 32
			default:
				config = configSmall
			}

			img := appkit.NewImageWithSystemSymbolNameAccessibilityDescription(sym.name, sym.label)
			configured := img.ImageWithSymbolConfiguration(config)

			imageView := appkit.NewNSImageView()
			imageView.SetImage(configured)
			imageView.SetContentTintColor(tintColors[i%len(tintColors)])
			imageView.SetImageAlignment(appkit.NSImageAlignCenter)
			imageView.SetTranslatesAutoresizingMaskIntoConstraints(false)
			imageView.WidthAnchor().ConstraintEqualToConstant(size).SetActive(true)
			imageView.HeightAnchor().ConstraintEqualToConstant(size).SetActive(true)

			label := appkit.NewTextFieldLabelWithString(sym.label)
			label.SetFont(font.SystemFontOfSize(10))
			label.SetTextColor(colors.SecondaryLabelColor())
			label.SetAlignment(appkit.NSTextAlignmentCenter)

			cell := appkit.GetNSStackViewClass().Alloc().Init()
			cell.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
			cell.SetAlignment(appkit.NSLayoutAttributeCenterX)
			cell.SetSpacing(6)
			cell.AddArrangedSubview(imageView)
			cell.AddArrangedSubview(label)

			rowStack.AddArrangedSubview(cell)
		}
		grid.AddArrangedSubview(rowStack)
	}

	contentView.AddSubview(layout)
	layout.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()).SetActive(true)
	layout.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()).SetActive(true)
	layout.TopAnchor().ConstraintEqualToAnchor(contentView.TopAnchor()).SetActive(true)
	layout.BottomAnchor().ConstraintEqualToAnchor(contentView.BottomAnchor()).SetActive(true)

	window.Center()
	window.MakeKeyAndOrderFront(nil)
	app.Activate()
}
