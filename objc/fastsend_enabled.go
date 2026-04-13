//go:build !objc_slowpath

package objc

func init() {
	initFastSend()
}
