package ext4

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"
)

// Build writes an ext4 image at imagePath populated from the directory tree at
// rootfsDir. The image is a classic indirect-mapped ext4 filesystem (see the
// package documentation) that mounts read-write on Linux.
func Build(ctx context.Context, rootfsDir, imagePath string) error {
	return (&Builder{}).Build(ctx, rootfsDir, imagePath)
}

// Builder builds an ext4 image. The zero value is ready to use; fields exist
// to override defaults, primarily for tests.
type Builder struct {
	// BlocksPerGroup overrides the blocks-per-group (default 32768). Tests set
	// a small value to force multi-group layouts.
	BlocksPerGroup int
	// now overrides the timestamp source (tests pin it for determinism).
	now func() time.Time
	// uuid overrides the random volume UUID (tests pin it for determinism).
	uuid *[16]byte
}

const (
	defaultBlockSize      = 4096
	defaultBlocksPerGroup = 32768
)

// Build writes the image.
func (b *Builder) Build(ctx context.Context, rootfsDir, imagePath string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	bs := defaultBlockSize
	bpg := b.BlocksPerGroup
	if bpg == 0 {
		bpg = defaultBlocksPerGroup
	}

	root, err := walk(rootfsDir)
	if err != nil {
		return fmt.Errorf("walk rootfs: %w", err)
	}

	lay, err := newLayout(root, bs, bpg)
	if err != nil {
		return fmt.Errorf("layout: %w", err)
	}

	now := time.Now
	if b.now != nil {
		now = b.now
	}
	lay.mkfsTime = now()

	if b.uuid != nil {
		lay.uuid = *b.uuid
	} else if _, err := rand.Read(lay.uuid[:]); err != nil {
		return fmt.Errorf("uuid: %w", err)
	}

	im, err := newImage(imagePath, bs, lay.blocksCount)
	if err != nil {
		return err
	}
	if err := lay.write(im); err != nil {
		im.close()
		return fmt.Errorf("write image: %w", err)
	}
	if err := im.close(); err != nil {
		return fmt.Errorf("close image: %w", err)
	}
	return nil
}
