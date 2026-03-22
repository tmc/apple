// Command nsort sorts lines by semantic similarity using Apple's NaturalLanguage
// sentence embeddings — a neural alternative to Unix sort.
//
// Lines are ordered by their cosine distance to a seed. By default the seed is
// the first line; use -seed to override.
//
// Usage:
//
//	nsort [flags] [file]
//	cat lines.txt | nsort -seed "animals"
//
// Flags:
//
//	-seed string   Seed word/phrase to sort by proximity to (default: first line)
//	-reverse       Sort farthest first
//	-distance      Print distance alongside each line
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"

	"github.com/tmc/apple/naturallanguage"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	seed := flag.String("seed", "", "seed word/phrase to sort by proximity to (default: first line)")
	reverse := flag.Bool("reverse", false, "sort farthest first")
	showDist := flag.Bool("distance", false, "print distance alongside each line")
	lang := flag.String("lang", "en", "language for embeddings")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nsort [flags] [file]\n\n")
		fmt.Fprintf(os.Stderr, "Sort lines by semantic similarity using NaturalLanguage embeddings.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	lines, err := readLines(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "nsort: %v\n", err)
		os.Exit(1)
	}
	if len(lines) == 0 {
		return
	}

	emb := naturallanguage.GetNLEmbeddingClass().SentenceEmbeddingForLanguage(naturallanguage.NLLanguage(*lang))
	if emb.ID == 0 {
		// Fall back to word embedding.
		emb = naturallanguage.GetNLEmbeddingClass().WordEmbeddingForLanguage(naturallanguage.NLLanguage(*lang))
	}
	if emb.ID == 0 {
		fmt.Fprintf(os.Stderr, "nsort: no embedding available for language %q\n", *lang)
		os.Exit(1)
	}

	seedStr := *seed
	if seedStr == "" {
		seedStr = lines[0]
	}

	type scored struct {
		line string
		dist float64
	}

	var items []scored
	for _, line := range lines {
		dist := emb.DistanceBetweenStringAndStringDistanceType(seedStr, line, naturallanguage.NLDistanceTypeCosine)
		items = append(items, scored{line, dist})
	}

	sort.SliceStable(items, func(i, j int) bool {
		if *reverse {
			return items[i].dist > items[j].dist
		}
		return items[i].dist < items[j].dist
	})

	for _, item := range items {
		if *showDist {
			fmt.Printf("%.4f\t%s\n", item.dist, item.line)
		} else {
			fmt.Println(item.line)
		}
	}
}

func readLines(args []string) ([]string, error) {
	var r io.Reader
	if len(args) == 0 || args[0] == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}

	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
