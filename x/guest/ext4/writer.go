package ext4

import (
	"fmt"
	"os"
)

// image is a block-addressable view over the output file.
type image struct {
	f         *os.File
	blockSize int
}

// newImage creates the image file and grows it to total blocks (zero-filled).
func newImage(path string, blockSize int, blocks uint32) (*image, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create image: %w", err)
	}
	size := int64(blocks) * int64(blockSize)
	if err := f.Truncate(size); err != nil {
		f.Close()
		return nil, fmt.Errorf("truncate image: %w", err)
	}
	return &image{f: f, blockSize: blockSize}, nil
}

// writeBlock writes data at block number n. data must not exceed one block.
func (im *image) writeBlock(n uint32, data []byte) error {
	if len(data) > im.blockSize {
		return fmt.Errorf("block %d: data %d exceeds block size %d", n, len(data), im.blockSize)
	}
	off := int64(n) * int64(im.blockSize)
	if _, err := im.f.WriteAt(data, off); err != nil {
		return fmt.Errorf("write block %d: %w", n, err)
	}
	return nil
}

// writeAt writes data at an absolute byte offset.
func (im *image) writeAt(off int64, data []byte) error {
	if _, err := im.f.WriteAt(data, off); err != nil {
		return fmt.Errorf("write at %d: %w", off, err)
	}
	return nil
}

// close flushes and closes the underlying file.
func (im *image) close() error {
	if err := im.f.Sync(); err != nil {
		im.f.Close()
		return err
	}
	return im.f.Close()
}
