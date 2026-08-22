// Copyright 2026 The tmc/apple Authors. All rights reserved.

// Package buildgate reports, per framework, whether that framework compiles.
//
// It exists because a test that imports a framework cannot report on a
// framework that does not compile: the test binary is never built, so no
// assertion in it runs and no summary is printed. One unbuildable package
// therefore erases the result of every other package in the same test binary,
// and the reader sees an undifferentiated FAIL.
//
// That happened here. The sibling smoketest package measures eight frameworks
// at runtime; when vision stopped compiling, smoketest reported
// "[build failed]" and its eight-way summary vanished. The board did not go
// from 5 green to 4 green, it went from 8 measured to 0 measured, and said so
// nowhere.
//
// This package imports nothing from the module. It shells out to the go tool,
// so it compiles and reports whatever else is broken. UNBUILDABLE is a distinct
// outcome from a runtime crash and is counted apart from it, in the same spirit
// as smoketest counting SKIPPED apart from passing: a framework that could not
// be compiled has not been found clean, it has not been measured at all.
package buildgate
