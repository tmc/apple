package main

import (
	"fmt"
	"log"

	"github.com/tmc/apple/foundation"
)

func main() {
	url := foundation.NewURLWithString("https://example.com")
	if url.AbsoluteString() == "" {
		log.Fatal("failed to create URL")
	}
	fmt.Printf("url=%s\n", url.AbsoluteString())

	// Fetch content synchronously via NSData.
	data := foundation.GetNSDataClass().DataWithContentsOfURL(url)
	if data.Length() == 0 {
		log.Fatal("failed to fetch URL content")
	}
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
