package newton_sqrt

import (
	"testing"
)

func TestNewtonSqrt(t *testing.T) {
	if newtonSqrt(333, 1e-7, 1e7) != 18.248287591617554 ||
		newtonSqrt(61009, 1e-7, 1e7) != 247 ||
		newtonSqrt(9, 1e-7, 1e7) != 3 {
		t.Error()
	}
}
