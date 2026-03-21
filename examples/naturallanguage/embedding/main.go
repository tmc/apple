// Command embedding computes word distances and nearest neighbors using
// Apple's NaturalLanguage word embeddings.
//
// Usage:
//
//	embedding distance <word1> <word2>
//	embedding neighbors [-n count] <word>
//	embedding info
//
// Output is JSON.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/naturallanguage"
	"github.com/tmc/apple/objc"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "distance":
		err = runDistance(args)
	case "neighbors":
		err = runNeighbors(args)
	case "info":
		err = runInfo(args)
	case "-h", "-help", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "embedding: unknown command %q\n", cmd)
		usage()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "embedding %s: %v\n", cmd, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: embedding <command> [flags]

Commands:
  distance <word1> <word2>   Compute cosine distance between two words
  neighbors [-n count] <word> Find nearest neighbor words
  info                        Show embedding metadata

Output is JSON.
`)
	os.Exit(2)
}

func loadEmbedding() (naturallanguage.NLEmbedding, error) {
	cls := naturallanguage.GetNLEmbeddingClass()
	emb := cls.WordEmbeddingForLanguage("en")
	if emb.ID == 0 {
		return emb, fmt.Errorf("no word embedding available for English")
	}
	return emb, nil
}

type distanceResult struct {
	Word1    string  `json:"word1"`
	Word2    string  `json:"word2"`
	Distance float64 `json:"distance"`
}

func runDistance(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: embedding distance <word1> <word2>")
	}
	word1, word2 := args[0], args[1]

	emb, err := loadEmbedding()
	if err != nil {
		return err
	}

	// distanceBetweenString:andString:distanceType:
	dist := objc.Send[float64](emb.ID, objc.Sel("distanceBetweenString:andString:distanceType:"),
		objc.String(word1),
		objc.String(word2),
		naturallanguage.NLDistanceTypeCosine,
	)

	return writeJSON(distanceResult{
		Word1:    word1,
		Word2:    word2,
		Distance: dist,
	})
}

type neighborEntry struct {
	Word     string  `json:"word"`
	Distance float64 `json:"distance"`
}

type neighborsResult struct {
	Query     string          `json:"query"`
	Neighbors []neighborEntry `json:"neighbors"`
}

func runNeighbors(args []string) error {
	fs := flag.NewFlagSet("neighbors", flag.ExitOnError)
	n := fs.Int("n", 10, "number of neighbors")
	fs.Parse(args)

	if fs.NArg() != 1 {
		return fmt.Errorf("usage: embedding neighbors [-n count] <word>")
	}
	word := fs.Arg(0)

	emb, err := loadEmbedding()
	if err != nil {
		return err
	}

	var neighbors []neighborEntry

	// Block signature: void(^)(NSString *neighbor, NLDistance distance, BOOL *stop)
	block := objc.NewBlock(func(_ objc.Block, neighborID objc.ID, distance float64, stop unsafe.Pointer) {
		neighbor := foundation.NSStringFromID(neighborID).String()
		neighbors = append(neighbors, neighborEntry{
			Word:     neighbor,
			Distance: distance,
		})
	})
	defer block.Release()

	// enumerateNeighborsForString:maximumCount:distanceType:usingBlock:
	objc.Send[objc.ID](emb.ID, objc.Sel("enumerateNeighborsForString:maximumCount:distanceType:usingBlock:"),
		objc.String(word),
		uintptr(*n),
		naturallanguage.NLDistanceTypeCosine,
		block,
	)

	if neighbors == nil {
		neighbors = []neighborEntry{}
	}
	return writeJSON(neighborsResult{
		Query:     word,
		Neighbors: neighbors,
	})
}

type infoResult struct {
	Language       string `json:"language"`
	Dimension      uint   `json:"dimension"`
	VocabularySize uint   `json:"vocabulary_size"`
	Revision       uint   `json:"revision"`
}

func runInfo(args []string) error {
	emb, err := loadEmbedding()
	if err != nil {
		return err
	}
	return writeJSON(infoResult{
		Language:       string(emb.Language()),
		Dimension:      emb.Dimension(),
		VocabularySize: emb.VocabularySize(),
		Revision:       emb.Revision(),
	})
}

func writeJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
