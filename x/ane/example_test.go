//go:build darwin

package ane_test

import (
	"fmt"
	"log"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// Compile and evaluate an identity model on the ANE.
func Example() {
	rt, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rt.Close()

	const channels = 4
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	k, err := rt.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	input := []float32{1, 2, 3, 4}
	if err := k.WriteInputF32(0, input); err != nil {
		log.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		log.Fatal(err)
	}

	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

// Open and close a Runtime, checking device information.
func ExampleRuntime() {
	rt, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rt.Close()

	info := rt.Info()
	fmt.Printf("HasANE=%v Cores=%d Arch=%s\n", info.HasANE, info.NumCores, info.Architecture)
}

// Write float32 input data and read float32 output data.
func ExampleKernel_WriteInputF32() {
	rt, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rt.Close()

	const channels = 2
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	k, err := rt.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	input := []float32{3.14, 2.72}
	if err := k.WriteInputF32(0, input); err != nil {
		log.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		log.Fatal(err)
	}

	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
