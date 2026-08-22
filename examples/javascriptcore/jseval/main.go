// Command jseval evaluates a JavaScript expression with JavaScriptCore and
// prints the result, reporting any uncaught JavaScript exception.
//
// The script comes from the command line arguments, or from standard input if
// no arguments are given.
//
// Usage: jseval [script]
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tmc/apple/javascriptcore"
)

func main() {
	script, err := readScript()
	if err != nil {
		fmt.Fprintf(os.Stderr, "jseval: %v\n", err)
		os.Exit(1)
	}
	if strings.TrimSpace(script) == "" {
		fmt.Fprintf(os.Stderr, "usage: jseval [script]\n")
		os.Exit(1)
	}

	ctx := javascriptcore.NewJSContext()
	ctx.SetName("jseval")

	value := ctx.EvaluateScript(script)
	if exc := ctx.Exception(); exc.GetID() != 0 {
		fmt.Fprintf(os.Stderr, "jseval: %s\n", exc.ToString())
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", typeName(value), value.ToString())
}

// readScript returns the script from the command line, or from standard input
// if no arguments were given.
func readScript() (string, error) {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " "), nil
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("reading stdin: %w", err)
	}
	return string(b), nil
}

// typeName reports the JavaScript type of value.
func typeName(value javascriptcore.IJSValue) string {
	switch {
	case value.IsUndefined():
		return "undefined"
	case value.IsNull():
		return "null"
	case value.IsBoolean():
		return "boolean"
	case value.IsNumber():
		return "number"
	case value.IsString():
		return "string"
	case value.IsSymbol():
		return "symbol"
	case value.IsArray():
		return "array"
	case value.IsDate():
		return "date"
	case value.IsObject():
		return "object"
	}
	return "unknown"
}
