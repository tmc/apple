package main

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

// MNIST image geometry and label count.
const (
	imageSize  = 28
	numClasses = 10
)

// mirror serves the four idx files the original sample fetched from
// yann.lecun.com, which no longer serves them.
const mirror = "https://ossci-datasets.s3.amazonaws.com/mnist/"

// dataset is one split of MNIST: count images of imageSize*imageSize bytes and
// the matching labels.
type dataset struct {
	count  int
	pixels []byte // count*imageSize*imageSize, one byte per pixel
	labels []byte // count labels in [0, numClasses)
}

// loadMNIST reads the training and test splits from dir, downloading any file
// that is missing.
func loadMNIST(dir string) (train, test *dataset, err error) {
	if dir == "" {
		return nil, nil, fmt.Errorf("no data directory; pass -data")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, nil, err
	}
	train, err = loadSplit(dir, "train-images-idx3-ubyte", "train-labels-idx1-ubyte")
	if err != nil {
		return nil, nil, err
	}
	test, err = loadSplit(dir, "t10k-images-idx3-ubyte", "t10k-labels-idx1-ubyte")
	if err != nil {
		return nil, nil, err
	}
	return train, test, nil
}

func loadSplit(dir, imageFile, labelFile string) (*dataset, error) {
	pixels, images, err := readImages(dir, imageFile)
	if err != nil {
		return nil, err
	}
	labels, err := readLabels(dir, labelFile)
	if err != nil {
		return nil, err
	}
	if len(labels) != images {
		return nil, fmt.Errorf("%s has %d images but %s has %d labels", imageFile, images, labelFile, len(labels))
	}
	return &dataset{count: images, pixels: pixels, labels: labels}, nil
}

// readImages returns the pixel bytes and image count of an idx3 file.
func readImages(dir, name string) ([]byte, int, error) {
	b, err := readIDX(dir, name)
	if err != nil {
		return nil, 0, err
	}
	if len(b) < 16 {
		return nil, 0, fmt.Errorf("%s: truncated header", name)
	}
	magic := binary.BigEndian.Uint32(b)
	rows := int(binary.BigEndian.Uint32(b[8:]))
	cols := int(binary.BigEndian.Uint32(b[12:]))
	if magic != 0x00000803 || rows != imageSize || cols != imageSize {
		return nil, 0, fmt.Errorf("%s: want %dx%d images, got magic %#x and %dx%d", name, imageSize, imageSize, magic, rows, cols)
	}
	n := int(binary.BigEndian.Uint32(b[4:]))
	pixels := b[16:]
	if len(pixels) < n*rows*cols {
		return nil, 0, fmt.Errorf("%s: want %d images, file holds %d", name, n, len(pixels)/(rows*cols))
	}
	return pixels[:n*rows*cols], n, nil
}

// readLabels returns the label bytes of an idx1 file.
func readLabels(dir, name string) ([]byte, error) {
	b, err := readIDX(dir, name)
	if err != nil {
		return nil, err
	}
	if len(b) < 8 {
		return nil, fmt.Errorf("%s: truncated header", name)
	}
	if magic := binary.BigEndian.Uint32(b); magic != 0x00000801 {
		return nil, fmt.Errorf("%s: want a label file, got magic %#x", name, magic)
	}
	n := int(binary.BigEndian.Uint32(b[4:]))
	labels := b[8:]
	if len(labels) < n {
		return nil, fmt.Errorf("%s: want %d labels, file holds %d", name, n, len(labels))
	}
	return labels[:n], nil
}

// readIDX returns the contents of dir/name, fetching and decompressing
// name+".gz" from the mirror when the file is not there yet.
func readIDX(dir, name string) ([]byte, error) {
	path := filepath.Join(dir, name)
	b, err := os.ReadFile(path)
	if err == nil {
		return b, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	if err := download(mirror+name+".gz", path); err != nil {
		return nil, fmt.Errorf("%s is missing and could not be downloaded: %w", path, err)
	}
	return os.ReadFile(path)
}

// download fetches a gzipped idx file and writes the decompressed bytes to
// path, leaving nothing behind if any step fails.
func download(url, path string) error {
	fmt.Printf("downloading %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s: %s", url, resp.Status)
	}
	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("get %s: %w", url, err)
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := io.Copy(tmp, zr); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

// randomBatch fills images and oneHot with uniformly sampled training cases.
func (d *dataset) randomBatch(rng *rand.Rand, images, oneHot []float32) {
	n := len(oneHot) / numClasses
	for i := 0; i < n; i++ {
		d.writeCase(rng.Intn(d.count), i, images, oneHot)
	}
}

// batchAt fills images and oneHot with the cases starting at start.
func (d *dataset) batchAt(start int, images, oneHot []float32) {
	n := len(oneHot) / numClasses
	for i := 0; i < n; i++ {
		d.writeCase(start+i, i, images, oneHot)
	}
}

// writeCase writes case idx of the dataset into slot of the batch, scaling
// pixels to [0, 1] and expanding the label to one-hot.
func (d *dataset) writeCase(idx, slot int, images, oneHot []float32) {
	const pixels = imageSize * imageSize
	src := d.pixels[idx*pixels : (idx+1)*pixels]
	dst := images[slot*pixels : (slot+1)*pixels]
	for i, p := range src {
		dst[i] = float32(p) / 255
	}
	labels := oneHot[slot*numClasses : (slot+1)*numClasses]
	for i := range labels {
		labels[i] = 0
	}
	labels[d.labels[idx]] = 1
}

// defaultDataDir is where the idx files are cached when -data is not given.
func defaultDataDir() string {
	cache, err := os.UserCacheDir()
	if err != nil {
		return ""
	}
	return filepath.Join(cache, "mnist")
}
