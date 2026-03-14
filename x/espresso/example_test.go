package espresso_test

import (
	"fmt"

	"github.com/tmc/apple/x/espresso"
)

func Example_types() {
	fmt.Println(espresso.PlatformAuto)
	fmt.Println(espresso.PlatformCPU)
	fmt.Println(espresso.PlatformGPU)
	fmt.Println(espresso.PlatformANE)

	s := espresso.Shape{Batch: 1, Channels: 3, Height: 224, Width: 224}
	fmt.Println(s.Elements())

	fmt.Println(espresso.EngineCPU)
	fmt.Println(espresso.EngineGPU)
	fmt.Println(espresso.EngineANE)

	// Output:
	// auto
	// cpu
	// gpu
	// ane
	// 150528
	// cpu
	// gpu
	// ane
}

func Example_argmax() {
	data := []float32{0.1, 0.7, 0.2}
	fmt.Println(espresso.Argmax(data))
	// Output:
	// 1
}

func Example_softmax() {
	data := []float32{1.0, 2.0, 3.0}
	result := espresso.Softmax(data)
	fmt.Printf("%.4f %.4f %.4f\n", result[0], result[1], result[2])
	// Output:
	// 0.0900 0.2447 0.6652
}

func Example_topK() {
	data := []float32{0.1, 0.9, 0.3, 0.7, 0.5}
	top3 := espresso.TopK(data, 3)
	fmt.Println(top3)
	// Output:
	// [1 3 4]
}
