//go:build darwin

package pdfkit_test

import (
	"fmt"

	"github.com/tmc/apple/pdfkit"
)

func ExamplePDFDocument() {
	doc := pdfkit.NewPDFDocument()
	page := pdfkit.NewPDFPage()
	doc.InsertPageAtIndex(page, 0)
	fmt.Println("PageCount after insert:", doc.PageCount())
	fmt.Println("IsEncrypted:", doc.IsEncrypted())
	fmt.Println("AllowsPrinting:", doc.AllowsPrinting())
	doc.RemovePageAtIndex(0)
	fmt.Println("PageCount after remove:", doc.PageCount())

	// Output:
	// PageCount after insert: 1
	// IsEncrypted: false
	// AllowsPrinting: true
	// PageCount after remove: 0
}

func ExamplePDFView() {
	view := pdfkit.NewPDFView()
	view.SetAutoScales(true)
	view.SetDisplayMode(pdfkit.KPDFDisplaySinglePage)

	fmt.Printf("AutoScales: %t\n", view.AutoScales())
	fmt.Printf("DisplayMode: %s\n", view.DisplayMode())

	// Output:
	// AutoScales: true
	// DisplayMode: KPDFDisplaySinglePage
}
