//go:build darwin

package javascriptcore_test

import (
	"fmt"

	"github.com/tmc/apple/javascriptcore"
)

func ExampleJSContext_EvaluateScript() {
	ctx := javascriptcore.NewJSContext()
	val := ctx.EvaluateScript("6 * 7")
	fmt.Println(val.ToInt32())

	// Output:
	// 42
}

func ExampleJSValue_ToString() {
	ctx := javascriptcore.NewJSContext()
	val := ctx.EvaluateScript("'Hello, ' + 'World!'")
	fmt.Println(val.ToString())

	// Output:
	// Hello, World!
}
