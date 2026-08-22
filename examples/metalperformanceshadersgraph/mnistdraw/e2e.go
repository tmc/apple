package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
)

// selfTest exercises the whole interactive path without a window: it trains,
// paints a real MNIST test image onto the canvas at canvas resolution, then
// runs the same normalize-and-classify step a mouse stroke would trigger.
//
// It is what makes this example checkable in CI, where nobody can draw a 3.
func selfTest() error {
	runtime.LockOSThread()

	train, test, err := loadMNIST(*dataDir)
	if err != nil {
		return err
	}
	a := &app{canvas: newCanvas(), train: train}
	if err := a.setup(); err != nil {
		return err
	}

	rng := rand.New(rand.NewSource(*seed))
	images := make([]float32, *batchSize*imageSize*imageSize)
	labels := make([]float32, *batchSize*numClasses)
	for i := 1; i <= *iterations; i++ {
		a.train.randomBatch(rng, images, labels)
		if loss := a.trainer.step(images, labels); i%50 == 0 || i == *iterations {
			fmt.Printf("training %d/%d — loss %.4f\n", i, *iterations, loss)
		}
	}

	// Classify the first few test digits through the canvas, the way a drawn
	// stroke would arrive.
	const cases = 20
	correct := 0
	one := make([]float32, imageSize*imageSize)
	oneLabel := make([]float32, numClasses)
	for i := 0; i < cases; i++ {
		test.batchAt(i, one, oneLabel)
		a.canvas.paint(one)

		probs := a.infer.run(a.canvas.normalized())
		best := 0
		for d := 1; d < numClasses; d++ {
			if probs[d] > probs[best] {
				best = d
			}
		}
		want := 0
		for d := range oneLabel {
			if oneLabel[d] == 1 {
				want = d
			}
		}
		if best == want {
			correct++
		}
		fmt.Printf("case %2d: predicted %d, actual %d, confidence %.2f\n", i, best, want, probs[best])
	}

	fmt.Printf("canvas round trip: %d/%d correct\n", correct, cases)
	if correct*4 < cases*3 { // the UI is broken well before accuracy drops this far
		return fmt.Errorf("only %d of %d canvas round trips classified correctly", correct, cases)
	}
	return nil
}

// runSelfTest is the -selftest entry point.
func runSelfTest() {
	if err := selfTest(); err != nil {
		fmt.Fprintf(os.Stderr, "mnistdraw: %v\n", err)
		os.Exit(1)
	}
}
