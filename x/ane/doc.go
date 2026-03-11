// Package ane provides high-level access to the Apple Neural Engine.
//
// Open a Runtime, compile MIL programs into Kernels, and evaluate them
// on the ANE hardware. IOSurface buffers are managed automatically.
//
//	rt, err := ane.Open()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rt.Close()
//
//	k, err := rt.Compile(ane.CompileOptions{
//		MILText:    milText,
//		WeightBlob: blob,
//		ModelType:  ane.ModelTypeMIL,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	k.WriteInputF32(0, input)
//	if err := k.Eval(); err != nil {
//		log.Fatal(err)
//	}
//	k.ReadOutputF32(0, output)
package ane
