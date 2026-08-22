//go:build darwin

package telemetry_test

import (
	"testing"

	"github.com/tmc/apple/x/ane/internal/anetest"
)

func TestMain(m *testing.M) {
	anetest.Run(m)
}
