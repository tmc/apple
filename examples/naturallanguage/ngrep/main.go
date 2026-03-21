// Command ngrep searches lines for semantic similarity to a query using Apple's
// NaturalLanguage embeddings — a neural alternative to grep.
//
// Unlike grep which matches exact text patterns, ngrep finds lines whose meaning
// is close to the query.
//
// Usage:
//
//	ngrep <query> [file]
//	cat lines.txt | ngrep "fast car"
//
// Flags:
//
//	-n int         Maximum number of results (default: 10)
//	-threshold     Maximum distance to include (0-2, default: 1.0)
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
	"github.com/tmc/apple/objc"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	maxResults := flag.Int("n", 10, "maximum number of results")
	threshold := flag.Float64("threshold", 1.0, "maximum cosine distance (0-2)")
	showDist := flag.Bool("distance", false, "print distance alongside each line")
	lang := flag.String("lang", "en", "language for embeddings")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ngrep [flags] <query> [file]\n\n")
		fmt.Fprintf(os.Stderr, "Search lines by semantic similarity using NaturalLanguage embeddings.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}

	query := flag.Arg(0)
	fileArgs := flag.Args()[1:]

	lines, err := readLines(fileArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ngrep: %v\n", err)
		os.Exit(1)
	}
	if len(lines) == 0 {
		os.Exit(1)
	}

	emb := naturallanguage.GetNLEmbeddingClass().SentenceEmbeddingForLanguage(naturallanguage.NLLanguage(*lang))
	if emb.ID == 0 {
		emb = naturallanguage.GetNLEmbeddingClass().WordEmbeddingForLanguage(naturallanguage.NLLanguage(*lang))
	}
	if emb.ID == 0 {
		fmt.Fprintf(os.Stderr, "ngrep: no embedding available for language %q\n", *lang)
		os.Exit(1)
	}

	type scored struct {
		line string
		dist float64
	}

	var matches []scored
	for _, line := range lines {
		dist := objc.Send[float64](emb.ID, objc.Sel("distanceBetweenString:andString:distanceType:"),
			objc.String(query),
			objc.String(line),
			naturallanguage.NLDistanceTypeCosine,
		)
		if dist <= *threshold {
			matches = append(matches, scored{line, dist})
		}
	}

	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].dist < matches[j].dist
	})

	if len(matches) > *maxResults {
		matches = matches[:*maxResults]
	}

	if len(matches) == 0 {
		os.Exit(1)
	}

	for _, m := range matches {
		if *showDist {
			fmt.Printf("%.4f\t%s\n", m.dist, m.line)
		} else {
			fmt.Println(m.line)
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
