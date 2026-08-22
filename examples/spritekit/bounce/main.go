// Command bounce shows a SpriteKit physics simulation in an AppKit window:
// colored sprites fall under gravity and bounce off the walls of the scene.
//
// Usage: bounce [-n count]
//
// A window server session is required. Close the window to quit.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/spritekit"
)

const (
	width  = 640
	height = 480
)

func main() {
	n := flag.Int("n", 12, "number of falling sprites")
	flag.Parse()
	if *n < 1 {
		fmt.Fprintf(os.Stderr, "bounce: -n must be at least 1\n")
		os.Exit(1)
	}

	// SpriteKit renders through the window server; without a session
	// (ssh, launchd daemon, CI) there is nothing to draw into.
	if coregraphics.CGSessionCopyCurrentDictionary() == 0 {
		fmt.Fprintf(os.Stderr, "bounce: no window server session; run from a graphical login\n")
		os.Exit(1)
	}

	appkit.RunApp(func(app appkit.NSApplication, _ appkit.NSApplicationDelegateObject) {
		frame := corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 200, Y: 200},
			Size:   corefoundation.CGSize{Width: width, Height: height},
		}
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			frame,
			appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable,
			appkit.NSBackingStoreBuffered,
			false,
		)
		window.SetTitle("Bounce")

		view := spritekit.NewSKView()
		view.SetFrame(window.ContentView().Bounds())
		view.SetAutoresizingMask(appkit.NSViewWidthSizable | appkit.NSViewHeightSizable)
		view.SetIgnoresSiblingOrder(true)
		view.SetShowsFPS(true)
		window.ContentView().AddSubview(view.NSView)

		view.PresentScene(newScene(*n))

		window.Center()
		window.MakeKeyAndOrderFront(nil)
		app.Activate()
	})
}

// newScene builds a scene whose edges are a static physics boundary,
// holding n dynamic sprites dropped from the top.
func newScene(n int) spritekit.SKScene {
	bounds := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 0, Y: 0},
		Size:   corefoundation.CGSize{Width: width, Height: height},
	}

	scene := spritekit.NewSceneWithSize(bounds.Size)
	scene.SetScaleMode(spritekit.SKSceneScaleModeAspectFit)
	scene.SetBackgroundColor(appkit.NewColorWithCalibratedRedGreenBlueAlpha(0.06, 0.07, 0.10, 1))

	// The scene's own body is the walls: an edge loop is always static.
	scene.SetPhysicsBody(spritekit.NewPhysicsBodyWithEdgeLoopFromRect(bounds))
	scene.PhysicsWorld().SetGravity(corefoundation.CGVector{Dx: 0, Dy: -6})

	for i := 0; i < n; i++ {
		size := corefoundation.CGSize{Width: 36, Height: 36}
		color := appkit.NewColorWithCalibratedRedGreenBlueAlpha(
			0.4+rand.Float64()*0.6,
			0.4+rand.Float64()*0.6,
			0.4+rand.Float64()*0.6,
			1,
		)
		sprite := spritekit.NewSpriteNodeWithColorSize(color, size)
		sprite.SetName(fmt.Sprintf("ball%d", i))
		sprite.SetPosition(corefoundation.CGPoint{
			X: 40 + rand.Float64()*(width-80),
			Y: height*0.5 + rand.Float64()*(height*0.4),
		})

		body := spritekit.NewPhysicsBodyWithRectangleOfSize(size)
		body.SetDynamic(true)
		body.SetRestitution(0.75) // bouncy
		body.SetFriction(0.2)
		body.SetLinearDamping(0.05)
		sprite.SetPhysicsBody(body)

		scene.AddChild(sprite.SKNode)
	}
	return scene
}
