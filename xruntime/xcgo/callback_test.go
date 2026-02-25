//go:build cgo

package xcgo_test

import (
	"testing"

	xcgotest "go.jcbhmr.com/xstd/xruntime/xcgo/internal/test"
)

func TestNewCallback_add(t *testing.T) {
	xcgotest.TestNewCallback_add(t)
}
