package tinyid

import (
	"reflect"
	"unsafe"
)

const (
	AlphabetStd       = "JedR8LNFY2j6MrhkBSADUyfP5amuH9xQCX4VqbgpsGtnW7vc3TwKE"
	AlphabetCanonical = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"

	defaultBlockSize = 24
	defaultLength    = 4
)

type Encoder struct {
	mask      int64
	BlockSize int
	Alphabet  string
	indices   [75]int
}

func New() Encoder {
	e := Encoder{
		Alphabet:  AlphabetStd,
		BlockSize: defaultBlockSize,
		mask:      (1 << defaultBlockSize) - 1,
	}
	e.init()
	return e
}

// Encode encodes x to an n-characters string
func (e Encoder) Encode(x int64, n int) string {
	b := make([]byte, n)
	e.enbase(shuffle(x, e.mask, e.BlockSize, reverse), &b)
	return bstos(b)
}

// Decode decodes an encoded s to its original 64bit integer
func (e Encoder) Decode(s string) int64 {
	return shuffle(e.debase(s), e.mask, e.BlockSize, recover)
}

func shuffle(x, mask int64, bs int, shuffler func(int64, int) int64) int64 {
	return (x & ^mask) | shuffler(x&mask, bs)
}

// x = 44 (0010 1100), bs = 5 (0 <= bs <= 64)
// => r = (0000 0110)
func reverse(x int64, bs int) (r int64) {
	b := bs - 1
	for i := 0; i < bs; i++ {
		if x&(1<<i) != 0 {
			r |= (1 << (b - i))
		}
	}
	return r
}

func recover(x int64, bs int) (r int64) {
	b := bs - 1
	for i := 0; i < bs; i++ {
		if x&(1<<(b-i)) != 0 {
			r |= (1 << i)
		}
	}
	return r
}

func (e Encoder) enbase(x int64, buf *[]byte) {
	n := len(*buf) - 1
	m := int64(len(e.Alphabet))

	for x > m {
		(*buf)[n] = e.Alphabet[x%m]
		x /= m
		n--
	}

	(*buf)[n] = e.Alphabet[x]

	// add padding
	for n > 0 {
		n--
		(*buf)[n] = e.Alphabet[0]
	}
}

func (e Encoder) debase(s string) (r int64) {
	n, m := len(s), len(e.Alphabet)
	for i := 0; i < n; i++ {
		r += int64(e.indices[s[n-i-1]-0x30] * powInt(m, i))
	}
	return
}

func powInt(x, y int) int {
	if y == 0 {
		return 1
	}

	r := x
	for y > 1 {
		r *= x
		y--
	}
	return r
}

func bstos(b []byte) (s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh.Data = uintptr(unsafe.Pointer(&b[0]))
	sh.Len = len(b)
	return
}

func (e *Encoder) init() {
	for i := 0; i < len(e.Alphabet); i++ {
		e.indices[e.Alphabet[i]-0x30] = i
	}
}

var DefaultEncoder = New()

func Encode(x int64, n int) string {
	return DefaultEncoder.Encode(x, n)
}

func Decode(s string) int64 {
	return DefaultEncoder.Decode(s)
}
