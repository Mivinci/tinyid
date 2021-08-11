# Tinyid

A bit.ly-like tiny id generator porting from [short_url.py](https://github.com/mozillazg/ShortURL/blob/master/shorturl/libs/short_url.py)

A bit-shuffling approach is used to avoid generating consecutive, predictable URLs. However, 
the algorithm is deterministic and will guarantee that no collisions will occur.

## Example

In this example a 64bit number 10000 is encoded to a 4-characters string "R3fu" that's decoded to its original value.

```go
import (
    "fmt"
    "github.com/mivinci/tinyid"
)

func main() {
    tinyid.Encode(10000, 4) // R3fu
    tinyid.Decode("R3fu")   // 10000
}
```

or use your own set of characters or blocksize.

```go
e := tinyid.New()
e.Alphabet = "your own set of characters"
e.BlockSize = 32

e.Encode(10000, 4)
```

## Benckmarks

Test case: Encode(10000, 4) and Decode("R3fu")

```bash
goos: darwin
goarch: amd64
pkg: github.com/mivinci/tinyid
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkEncoder/Encode-4         	 8662382	       125.4 ns/op	       4 B/op	       1 allocs/op
BenchmarkEncoder/Decode-4         	13147922	        90.22 ns/op	       0 B/op	       0 allocs/op
```

## Other Portings

Here's a [C implementation](https://github.com/mivinci/tinyid.c) of tinyid.