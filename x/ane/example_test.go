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
	c, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	const channels = 4
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	m, err := c.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	input := []float32{1, 2, 3, 4}
	if err := m.WriteInputF32(0, input); err != nil {
		log.Fatal(err)
	}
	if err := m.Eval(); err != nil {
		log.Fatal(err)
	}

	output := make([]float32, channels)
	if err := m.ReadOutputF32(0, output); err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

// Open and close a Client, checking device information.
func ExampleClient() {
	c, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	info := c.Info()
	fmt.Printf("HasANE=%v Cores=%d Arch=%s\n", info.HasANE, info.NumCores, info.Architecture)
}

// Write float32 input data and read float32 output data.
func ExampleModel_WriteInputF32() {
	c, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	const channels = 2
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	m, err := c.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	input := []float32{3.14, 2.72}
	if err := m.WriteInputF32(0, input); err != nil {
		log.Fatal(err)
	}
	if err := m.Eval(); err != nil {
		log.Fatal(err)
	}

	output := make([]float32, channels)
	if err := m.ReadOutputF32(0, output); err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
