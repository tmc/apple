//go:build darwin

package ext4_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/x/guest/ext4"
)

func ExampleBuild() {
	rootfs, err := os.MkdirTemp("", "ext4-rootfs-*")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(rootfs)
	if err := os.WriteFile(filepath.Join(rootfs, "etc-release"), []byte("ID=example\n"), 0644); err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.MkdirTemp("", "ext4-image-*")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(out)
	imagePath := filepath.Join(out, "rootfs.ext4")

	if err := ext4.Build(context.Background(), rootfs, imagePath); err != nil {
		fmt.Println(err)
		return
	}
	info, err := os.Stat(imagePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Ext4 image created: size=%d bytes\n", info.Size())

	// Output:
	// Ext4 image created: size=134217728 bytes
}
