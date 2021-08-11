package tinyid

import (
	"fmt"
	"testing"
)

func i64tos(x int64) string {
	var b [64]byte
	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			b[63-i] = '1'
		} else {
			b[63-i] = '0'
		}
	}
	return string(b[:])
}

func ExampleEncoder() {
	var bs int = 64
	var mask int64 = (1 << bs) - 1
	var x int64 = 1

	fmt.Println("AA")
	fmt.Println("   mask", i64tos(mask))
	fmt.Println("  ^mask", i64tos(^mask))
	fmt.Println(" origin", i64tos(x))
	fmt.Println(" x&mask", i64tos(x&mask))
	fmt.Println("x&^mask", i64tos(x&^mask))
	fmt.Println("reverse", i64tos(reverse(x&mask, bs)))
	fmt.Println("encoded", i64tos(expand(x, mask, bs)))
}

func ExampleEncoder_Encode() {
	e := New()
	fmt.Println(e.Encode(10000, 4))

	// Output:
	// R3fu
}

func TestEncoder(t *testing.T) {
	e := New()
	x := int64(10000)
	encoded := e.Encode(x, 4)
	decoded := e.Decode(encoded)
	if decoded != x {
		t.Errorf("want %d, but got %d", x, decoded)
		t.Fail()
	}
}

func BenchmarkEncoder(b *testing.B) {
	e := New()
	b.ReportAllocs()
	b.Run("Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.Encode(10000, 4)
		}
	})
	b.Run("Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e.Decode("R3fu")
		}
	})
}
