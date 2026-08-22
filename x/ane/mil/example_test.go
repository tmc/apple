//go:build darwin

package mil_test

import (
	"fmt"

	"github.com/tmc/apple/x/ane/mil"
)

// Pack several weight tensors into one blob file and reference each by its
// offset. Offset reports the value a BLOBFILE reference takes, which addresses
// the entry's metadata block rather than its raw data, so it must be read
// after every tensor has been added.
func ExampleBlobWriter() {
	w := mil.NewBlobWriter()
	first := w.AddFloat16([]float32{1, 2, 3, 4})
	second := w.AddFloat16(make([]float32, 64))

	blob, err := w.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d tensors in %d bytes\n", w.Count(), len(blob))
	fmt.Printf("BLOBFILE(path = string(\"@model_path/weights.bin\"), offset = uint64(%d))\n", w.Offset(first))
	fmt.Printf("BLOBFILE(path = string(\"@model_path/weights.bin\"), offset = uint64(%d))\n", w.Offset(second))

	// Output:
	// 2 tensors in 384 bytes
	// BLOBFILE(path = string("@model_path/weights.bin"), offset = uint64(64))
	// BLOBFILE(path = string("@model_path/weights.bin"), offset = uint64(192))
}

// A single weight tensor in its own file is the common case, and the offset is
// always 64: the file header takes the first 64 bytes and the one metadata
// block follows it.
func ExampleBuildWeightBlob() {
	blob, err := mil.BuildWeightBlob([]float32{1, 0, 0, 1}, 2, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(blob))

	// Output:
	// 136
}
