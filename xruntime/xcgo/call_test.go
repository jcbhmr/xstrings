//go:build cgo

package xcgo_test

import (
	"testing"

	xcgotest "go.jcbhmr.com/xstd/xruntime/xcgo/internal/test"
)

func TestCall_add(t *testing.T) {
	xcgotest.TestCall_add(t)
}

func TestCall_getFortyTwo(t *testing.T) {
	xcgotest.TestCall_getFortyTwo(t)
}
