// SF-symbols displays a grid of SF Symbols with various configurations.
//
// It demonstrates NSImage system symbol loading, NSImageView display,
// NSImageSymbolConfiguration for sizing and colors, and the symbols
// package for effect types.
package main

import (
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
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

func main() {
	appkit.RunApp(func(app appkit.NSApplication, delegate appkit.NSApplicationDelegateObject) {
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: 200, Y: 200},
				Size:   corefoundation.CGSize{Width: 520, Height: 460},
			},
			appkit.NSWindowStyleMaskTitled|
				appkit.NSWindowStyleMaskClosable|
				appkit.NSWindowStyleMaskMiniaturizable,
			appkit.NSBackingStoreBuffered, false,
		)
		window.SetTitle("SF Symbols")

		contentView := window.ContentView().(appkit.NSView)
		font := appkit.GetNSFontClass()
		colors := appkit.GetNSColorClass()

		// Title label.
		title := appkit.NewTextFieldLabelWithString("SF Symbols from Go")
		title.SetFont(font.SystemFontOfSizeWeight(18, appkit.NSFontWeights.Semibold))
		title.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 20, Y: 410},
			Size:   corefoundation.CGSize{Width: 480, Height: 28},
		})
		contentView.AddSubview(title)

		subtitle := appkit.NewTextFieldLabelWithString("NSImage + NSImageSymbolConfiguration + Symbols framework")
		subtitle.SetFont(font.SystemFontOfSize(12))
		subtitle.SetTextColor(colors.SecondaryLabelColor())
		subtitle.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 20, Y: 390},
			Size:   corefoundation.CGSize{Width: 480, Height: 18},
		})
		contentView.AddSubview(subtitle)

		// Symbol grid: 4 columns x 4 rows.
		cols := 4
		cellW := 120.0
		cellH := 80.0
		gridX := 20.0
		gridTop := 370.0

		// Three symbol configurations at different scales.
		configSmall := appkit.NewImageSymbolConfigurationWithPointSizeWeight(16, appkit.NSFontWeights.Regular)
		configMedium := appkit.NewImageSymbolConfigurationWithPointSizeWeight(24, appkit.NSFontWeights.Medium)
		configLarge := appkit.NewImageSymbolConfigurationWithPointSizeWeight(32, appkit.NSFontWeights.Light)

		// Tint colors for visual variety.
		tintColors := []appkit.INSColor{
			colors.SystemBlue(),
			colors.SystemGreen(),
			colors.SystemOrange(),
			colors.SystemPurple(),
			colors.SystemRed(),
			colors.SystemTeal(),
			colors.SystemIndigo(),
			colors.SystemPink(),
		}

		for i, sym := range symbols {
			col := i % cols
			row := i / cols
			x := gridX + float64(col)*cellW
			y := gridTop - float64(row+1)*cellH

			// Pick config based on row for visual variety.
			var config appkit.NSImageSymbolConfiguration
			switch row {
			case 0:
				config = configLarge
			case 1:
				config = configMedium
			default:
				config = configSmall
			}

			// Load system symbol and apply configuration.
			img := appkit.NewImageWithSystemSymbolNameAccessibilityDescription(sym.name, sym.label)
			configured := img.ImageWithSymbolConfiguration(config)

			imageView := appkit.NewNSImageView()
			imageView.SetImage(configured)
			imageView.SetContentTintColor(tintColors[i%len(tintColors)])
			imageView.SetFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: x + cellW/2 - 20, Y: y + 24},
				Size:   corefoundation.CGSize{Width: 40, Height: 40},
			})
			imageView.SetImageAlignment(appkit.NSImageAlignCenter)
			contentView.AddSubview(imageView)

			// Label below.
			label := appkit.NewTextFieldLabelWithString(sym.label)
			label.SetFont(font.SystemFontOfSize(10))
			label.SetTextColor(colors.SecondaryLabelColor())
			label.SetAlignment(appkit.NSTextAlignmentCenter)
			label.SetFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: x, Y: y + 4},
				Size:   corefoundation.CGSize{Width: cellW, Height: 16},
			})
			contentView.AddSubview(label)
		}

		window.Center()
		window.MakeKeyAndOrderFront(nil)
		app.Activate()
	})
}
