// Command http-fetch downloads a URL with Foundation and prints a short preview.
//
// Usage:
//
//	http-fetch [url]
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/foundation"
)

func main() {
	urlString := "https://example.com"
	switch len(os.Args) {
	case 1:
	case 2:
		urlString = os.Args[1]
	default:
		fmt.Fprintln(os.Stderr, "usage: http-fetch [url]")
		os.Exit(2)
	}

	url := foundation.NewURLWithString(urlString)
	if url.AbsoluteString() == "" {
		fmt.Fprintf(os.Stderr, "http-fetch: invalid URL %q\n", urlString)
		os.Exit(1)
	}
	fmt.Printf("url=%s\n", url.AbsoluteString())

	// Fetch content synchronously via NSData.
	data := foundation.GetNSDataClass().DataWithContentsOfURL(url)
	fmt.Printf("bytes=%d\n", data.Length())

	// Decode a prefix of the response as a string.
	str := foundation.GetNSStringClass().Alloc().InitWithDataEncoding(data, uint(foundation.NSUTF8StringEncoding))
	full := str.UTF8String()
	preview := full
	if len(preview) > 80 {
		preview = preview[:80] + "..."
	}
	fmt.Printf("preview=%s\n", preview)
}
