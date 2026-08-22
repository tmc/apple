// Command mnistdraw is an interactive version of Apple's "Training a Neural
// Network using MPS Graph" sample: draw a digit with the mouse and watch the
// network classify it as you draw.
//
// It is the UI half of the sample, which the headless mnisttrain example does
// not cover. The same MPSGraph network is trained on MNIST at startup, then
// each stroke is normalized the way MNIST was and fed back through the graph
// for a live prediction.
//
// The window has three parts:
//
//   - a canvas to draw on, left
//   - one bar per digit showing the softmax probability, right
//   - Train, Clear, and a status line, bottom
//
// Training runs on a background goroutine so the window stays responsive, and
// the graph uses a dynamic batch dimension so the very same weights serve both
// the 40-image training steps and the single-image canvas predictions.
//
// Usage:
//
//	mnistdraw [-data dir] [-iterations n] [-batch n] [-lr rate]
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

var (
	dataDir    = flag.String("data", defaultDataDir(), "directory holding the MNIST idx files")
	iterations = flag.Int("iterations", 300, "training iterations to run when Train is pressed")
	batchSize  = flag.Int("batch", 40, "images per training batch")
	rate       = flag.Float64("lr", 0.01, "learning rate")
	seed       = flag.Int64("seed", 1, "random seed for the weights and batch sampling")
	selftest   = flag.Bool("selftest", false, "train and classify test digits through the canvas, without opening a window")
)

// Window geometry. The layout is fixed, so plain frames are clearer here than
// a lattice of layout constraints.
const (
	pad        = 20.0
	barWidth   = 220.0
	barHeight  = 22.0
	barGap     = 6.0
	panelWidth = canvasSize + barWidth + 3*pad
	footer     = 84.0
	winWidth   = panelWidth
	winHeight  = canvasSize + footer + 2*pad
)

// Go starts init and main on the OS main thread, but the main goroutine may be
// migrated to another thread after blocking work -- and this program loads ~55MB
// of MNIST before it opens a window. AppKit must run on the real main thread, so
// pin it here, before any of that. appkit.RunApp locks the thread too, but by
// then it can only pin whichever thread the goroutine has already landed on.
func init() {
	runtime.LockOSThread()
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "mnistdraw: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if *selftest {
		runSelfTest()
		return nil
	}

	train, _, err := loadMNIST(*dataDir)
	if err != nil {
		return err
	}
	ui := &app{canvas: newCanvas(), train: train}
	appkit.RunApp(func(nsapp appkit.NSApplication, _ appkit.NSApplicationDelegateObject) {
		ui.build(nsapp)
		go ui.trainLoop()
	})
	return nil
}

// app holds the window, the widgets it updates, and the network behind them.
type app struct {
	canvas *canvas
	train  *dataset

	imageView appkit.NSImageView
	bars      [numClasses]appkit.NSLevelIndicator
	barLabels [numClasses]appkit.NSTextField
	status    appkit.NSTextField
	trainBtn  appkit.NSButton

	// net is written by the training goroutine and read by the main thread
	// only after ready closes, so no lock is needed.
	net      *classifier
	trainer  *session
	infer    *session
	ready    chan struct{}
	training bool
}

// build lays out the window and wires up the controls.
func (a *app) build(nsapp appkit.NSApplication) {
	a.ready = make(chan struct{})

	window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 200, Y: 200},
			Size:   corefoundation.CGSize{Width: winWidth, Height: winHeight},
		},
		appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable|appkit.NSWindowStyleMaskMiniaturizable,
		appkit.NSBackingStoreBuffered,
		false,
	)
	window.SetTitle("MNIST — draw a digit")
	content := window.ContentView()

	// Canvas, top left. Its frame is also the hit region for the mouse.
	canvasFrame := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: pad, Y: footer},
		Size:   corefoundation.CGSize{Width: canvasSize, Height: canvasSize},
	}
	a.imageView = appkit.NewImageViewWithFrame(canvasFrame)
	a.imageView.SetImageScaling(appkit.NSImageScaleAxesIndependently)
	a.imageView.SetImage(a.canvas.image())
	content.AddSubview(a.imageView)

	// One labelled bar per digit, top right, counting down from 9 so that 0
	// ends up at the bottom.
	barX := pad + canvasSize + pad
	for d := 0; d < numClasses; d++ {
		y := footer + float64(numClasses-1-d)*(barHeight+barGap)

		label := appkit.NewTextFieldLabelWithString(fmt.Sprintf("%d", d))
		label.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: barX, Y: y},
			Size:   corefoundation.CGSize{Width: 16, Height: barHeight},
		})
		content.AddSubview(label)
		a.barLabels[d] = label

		bar := appkit.NewLevelIndicatorWithFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: barX + 20, Y: y},
			Size:   corefoundation.CGSize{Width: barWidth - 70, Height: barHeight},
		})
		bar.SetLevelIndicatorStyle(appkit.NSLevelIndicatorStyleContinuousCapacity)
		bar.SetMinValue(0)
		bar.SetMaxValue(1)
		bar.SetDoubleValue(0)
		content.AddSubview(bar)
		a.bars[d] = bar
	}

	// Footer: status line and buttons.
	a.status = appkit.NewTextFieldLabelWithString("loading…")
	a.status.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: pad, Y: 46},
		Size:   corefoundation.CGSize{Width: winWidth - 2*pad, Height: 20},
	})
	content.AddSubview(a.status)

	a.trainBtn = appkit.NewButtonWithTitleTargetAction("Train", nil, 0)
	a.trainBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: pad, Y: 12},
		Size:   corefoundation.CGSize{Width: 110, Height: 28},
	})
	a.trainBtn.SetActionHandler(func() {
		if a.training {
			return
		}
		go a.trainLoop()
	})
	content.AddSubview(a.trainBtn)

	clear := appkit.NewButtonWithTitleTargetAction("Clear", nil, 0)
	clear.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: pad + 120, Y: 12},
		Size:   corefoundation.CGSize{Width: 110, Height: 28},
	})
	clear.SetActionHandler(func() {
		a.canvas.clear()
		a.imageView.SetImage(a.canvas.image())
		for d := 0; d < numClasses; d++ {
			a.bars[d].SetDoubleValue(0)
		}
		a.setStatus("cleared — draw a digit")
	})
	content.AddSubview(clear)

	// Mouse handling. A local monitor avoids having to register an NSView
	// subclass just to receive three events.
	appkit.GetNSEventClass().AddLocalMonitorForEventsMatchingMaskHandler(
		appkit.NSEventMaskLeftMouseDown|appkit.NSEventMaskLeftMouseDragged|appkit.NSEventMaskLeftMouseUp,
		func(e *appkit.NSEvent) appkit.NSEvent {
			a.handleMouse(e, canvasFrame)
			return *e
		})

	// Center on the primary screen, which is screens[0] -- the one whose frame
	// starts at the origin and carries the menu bar.
	//
	// Neither NSWindow.center() nor NSScreen.mainScreen() is right here: both
	// follow the screen the window is already on, and a window born at a fixed
	// origin can land on an adjacent display. On a multi-monitor desk that puts
	// it somewhere the user never sees.
	if screens := appkit.GetNSScreenClass().Screens(); len(screens) > 0 {
		visible := screens[0].VisibleFrame()
		window.SetFrameOrigin(corefoundation.CGPoint{
			X: visible.Origin.X + (visible.Size.Width-winWidth)/2,
			Y: visible.Origin.Y + (visible.Size.Height-winHeight)/2,
		})
	}

	window.MakeKeyAndOrderFront(nil)
	// RunApp already set the regular activation policy, so the window can take
	// focus even though it was launched from a terminal.
	nsapp.Activate()
}

// handleMouse routes a mouse event to the canvas when it lands inside frame.
//
// Window coordinates have their origin at the bottom left, and so does the
// canvas frame, but the pixel buffer is top-row-first: hence the y flip.
func (a *app) handleMouse(e *appkit.NSEvent, frame corefoundation.CGRect) {
	p := e.LocationInWindow()
	x := p.X - frame.Origin.X
	y := frame.Size.Height - (p.Y - frame.Origin.Y)

	switch e.Type() {
	case appkit.NSEventTypeLeftMouseDown:
		if x < 0 || y < 0 || x >= canvasSize || y >= canvasSize {
			return
		}
		a.canvas.beginStroke(x, y)
	case appkit.NSEventTypeLeftMouseDragged:
		if !a.canvas.drawing {
			return
		}
		a.canvas.extendStroke(x, y)
	case appkit.NSEventTypeLeftMouseUp:
		if !a.canvas.drawing {
			return
		}
		a.canvas.endStroke()
	default:
		return
	}

	a.imageView.SetImage(a.canvas.image())
	a.predict()
}

// predict classifies whatever is on the canvas and moves the bars.
func (a *app) predict() {
	select {
	case <-a.ready:
	default:
		return // still training the first time through
	}
	if a.canvas.isEmpty() {
		return
	}

	probs := a.infer.run(a.canvas.normalized())
	best := 0
	for d := 1; d < numClasses; d++ {
		if probs[d] > probs[best] {
			best = d
		}
	}
	for d := 0; d < numClasses; d++ {
		a.bars[d].SetDoubleValue(float64(probs[d]))
	}
	a.setStatus(fmt.Sprintf("prediction: %d  (%.1f%% confident)", best, 100*probs[best]))
}

// trainLoop builds the network if needed and runs a round of training. It runs
// on its own goroutine and thread, so the GPU work never blocks the run loop.
func (a *app) trainLoop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	a.training = true
	a.onMain(func() { a.trainBtn.SetEnabled(false) })
	defer func() {
		a.training = false
		a.onMain(func() { a.trainBtn.SetEnabled(true) })
	}()

	if a.net == nil {
		if err := a.setup(); err != nil {
			a.setStatus("setup failed: " + err.Error())
			return
		}
	}

	rng := rand.New(rand.NewSource(*seed))
	images := make([]float32, *batchSize*imageSize*imageSize)
	labels := make([]float32, *batchSize*numClasses)

	for i := 1; i <= *iterations; i++ {
		a.train.randomBatch(rng, images, labels)
		loss := a.trainer.step(images, labels)
		if i%10 == 0 || i == *iterations {
			a.setStatus(fmt.Sprintf("training %d/%d — loss %.4f", i, *iterations, loss))
		}
	}

	select {
	case <-a.ready:
	default:
		close(a.ready)
	}
	a.setStatus("ready — draw a digit on the canvas")
}

// setup creates the device, the graph, and the two run configurations.
func (a *app) setup() error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a command queue")
	}

	net, err := newClassifier(rand.New(rand.NewSource(*seed)), *batchSize, *rate)
	if err != nil {
		return err
	}
	a.net = net

	if a.trainer, err = newSession(device, queue, net, *batchSize); err != nil {
		return err
	}
	if a.infer, err = newSession(device, queue, net, 1); err != nil {
		return err
	}
	return nil
}

// setStatus updates the status line from any goroutine.
func (a *app) setStatus(s string) {
	a.onMain(func() { a.status.SetStringValue(s) })
}

// onMain runs fn on the main thread, where every AppKit call belongs.
func (a *app) onMain(fn func()) {
	dispatch.MainQueue().Async(fn)
}

// session binds one batch size to the graph: the buffers, the feeds, and the
// result dictionaries needed to run it. Two sessions share one set of weights,
// which is what the graph's dynamic batch dimension buys.
type session struct {
	queue  metal.MTLCommandQueue
	net    *classifier
	batch  int
	images *feed
	labels *feed
	feeds  foundation.NSDictionary
	loss   *result
	probs  *result
}

func newSession(device metal.MTLDeviceObject, queue metal.MTLCommandQueue, net *classifier, batch int) (*session, error) {
	images, err := newFeed(device, net.images, batch*imageSize*imageSize, batch, imageSize*imageSize)
	if err != nil {
		return nil, err
	}
	labels, err := newFeed(device, net.labels, batch*numClasses, batch, numClasses)
	if err != nil {
		return nil, err
	}
	feeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{images.data, labels.data},
		[]objectivec.IObject{images.tensor, labels.tensor},
	)
	if feeds.GetID() == 0 {
		return nil, fmt.Errorf("could not create the feeds dictionary")
	}
	loss, err := newResult(device, net.loss, 1, 1)
	if err != nil {
		return nil, err
	}
	probs, err := newResult(device, net.probabilities, batch*numClasses, batch, numClasses)
	if err != nil {
		return nil, err
	}
	return &session{queue: queue, net: net, batch: batch, images: images, labels: labels, feeds: feeds, loss: loss, probs: probs}, nil
}

// step runs one training iteration and returns the mean loss.
func (s *session) step(images, labels []float32) float32 {
	s.images.write(images)
	s.labels.write(labels)
	s.net.graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(
		s.queue, s.feeds, s.net.updates, s.loss.dictionary)
	return s.loss.read(1)[0]
}

// run classifies one batch and returns the softmax row for the first image.
// Passing no assign operations is what makes this inference: the weights are
// read, never written.
func (s *session) run(images []float32) []float32 {
	s.images.write(images)
	s.net.graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(
		s.queue, s.feeds, nil, s.probs.dictionary)
	return s.probs.read(numClasses)
}
