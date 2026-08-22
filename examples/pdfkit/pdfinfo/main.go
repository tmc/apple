// Command pdfinfo reports document metadata and per-page information for a
// PDF file using the PDFKit framework. With -text it also prints the text of
// each page.
//
// Usage: pdfinfo [-text] <pdf-file>
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/pdfkit"
)

func main() {
	text := flag.Bool("text", false, "print the text of each page")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: pdfinfo [-text] <pdf-file>\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	path := flag.Arg(0)

	url := foundation.NewURLFileURLWithPath(path)
	doc := pdfkit.NewPDFDocumentWithURL(url)
	if doc.ID == 0 {
		fmt.Fprintf(os.Stderr, "pdfinfo: cannot open %s\n", path)
		os.Exit(1)
	}
	if doc.IsLocked() {
		fmt.Fprintf(os.Stderr, "pdfinfo: %s is locked\n", path)
		os.Exit(1)
	}

	fmt.Printf("File:      %s\n", path)
	fmt.Printf("Version:   %d.%d\n", doc.MajorVersion(), doc.MinorVersion())
	fmt.Printf("Pages:     %d\n", doc.PageCount())
	fmt.Printf("Encrypted: %t\n", doc.IsEncrypted())

	attrs := doc.DocumentAttributes()
	for _, a := range []struct {
		label string
		key   pdfkit.PDFDocumentAttribute
	}{
		{"Title", pdfkit.PDFDocumentTitleAttribute},
		{"Author", pdfkit.PDFDocumentAuthorAttribute},
		{"Subject", pdfkit.PDFDocumentSubjectAttribute},
		{"Creator", pdfkit.PDFDocumentCreatorAttribute},
		{"Producer", pdfkit.PDFDocumentProducerAttribute},
		{"Created", pdfkit.PDFDocumentCreationDateAttribute},
		{"Modified", pdfkit.PDFDocumentModificationDateAttribute},
	} {
		if v := attribute(attrs, a.key); v != "" {
			fmt.Printf("%-10s %s\n", a.label+":", v)
		}
	}

	for i := uint(0); i < doc.PageCount(); i++ {
		page := doc.PageAtIndex(i)
		if page == nil || page.GetID() == 0 {
			fmt.Fprintf(os.Stderr, "pdfinfo: page %d unavailable\n", i)
			continue
		}
		bounds := page.BoundsForBox(pdfkit.KPDFDisplayBoxMediaBox)
		fmt.Printf("\npage %d: label=%q chars=%d rotation=%d size=%.0fx%.0f\n",
			i+1, page.Label(), page.NumberOfCharacters(), page.Rotation(),
			bounds.Size.Width, bounds.Size.Height)
		if *text {
			fmt.Println(page.String())
		}
	}
}

// attribute returns the string form of the document attribute named by key,
// or "" if the document has no such attribute.
func attribute(attrs foundation.INSDictionary, key pdfkit.PDFDocumentAttribute) string {
	if attrs == nil || attrs.GetID() == 0 {
		return ""
	}
	v := attrs.ObjectForKey(foundation.NewStringWithCString(string(key)))
	if v == nil || v.GetID() == 0 {
		return ""
	}
	return objectivec.Object{ID: v.GetID()}.Description()
}
