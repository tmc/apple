//go:build darwin

// Package linear provides a cached linear (matrix multiply) operator backed
// by the Apple Neural Engine.
//
//	exec := linear.New(rt)
//	defer exec.Close()
//
//	y, err := exec.Linear(x, w, batch, inDim, outDim)
package linear
