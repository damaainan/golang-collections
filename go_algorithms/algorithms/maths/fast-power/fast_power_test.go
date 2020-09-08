package fast_power

import (
	"math"
	"testing"
)

func TestFastPower(t *testing.T) {
	a := uint32(2)

	for i := 0; i < 100000; i++ {
		n, ok := fastPower(a, i)

		if ok != nil || n != uint32(math.Pow(float64(a), float64(i))) {
			t.Error()
		}
	}
}

func TestSlowPower(t *testing.T) {
	a := uint32(2)

	for i := 0; i < 100000; i++ {
		n, ok := slowPower(a, i)

		if ok != nil || n != uint32(math.Pow(float64(a), float64(i))) {
			t.Error()
		}
	}
}

func BenchmarkFastPower(b *testing.B) {
	a := uint32(2)

	for i := 0; i < b.N; i++ {
		fastPower(a, 100000)
	}
}

func BenchmarkSlowPower(b *testing.B) {
	a := uint32(2)

	for i := 0; i < b.N; i++ {
		slowPower(a, 100000)
	}
}
