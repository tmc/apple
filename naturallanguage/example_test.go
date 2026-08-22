//go:build darwin

package naturallanguage_test

import (
	"fmt"

	"github.com/tmc/apple/naturallanguage"
)

func ExampleNLLanguageRecognizer() {
	recognizer := naturallanguage.NewNLLanguageRecognizer()
	recognizer.ProcessString("Hello world")
	lang1 := recognizer.DominantLanguage()

	recognizer.Reset()
	recognizer.ProcessString("Hola mundo")
	lang2 := recognizer.DominantLanguage()

	fmt.Println("Dominant 1:", lang1)
	fmt.Println("Dominant 2:", lang2)

	// Output:
	// Dominant 1: en
	// Dominant 2: es
}

func ExampleNLTokenizer() {
	tokenizer := naturallanguage.NewTokenizerWithUnit(naturallanguage.NLTokenUnitWord)
	tokenizer.SetString("Hello world")

	r0 := tokenizer.TokenRangeAtIndex(0)
	r1 := tokenizer.TokenRangeAtIndex(6)

	fmt.Println("Unit:", tokenizer.Unit())
	fmt.Println("String:", tokenizer.String())
	fmt.Printf("Token 0 range: [%d, %d]\n", r0.Location, r0.Length)
	fmt.Printf("Token 6 range: [%d, %d]\n", r1.Location, r1.Length)

	// Output:
	// Unit: NLTokenUnitWord
	// String: Hello world
	// Token 0 range: [0, 5]
	// Token 6 range: [6, 5]
}
