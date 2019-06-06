[![GoDoc](https://godoc.org/github.com/dsincl12/wyrand?status.svg)](https://godoc.org/github.com/dsincl12/wyrand)
[![Go Report Card](https://goreportcard.com/badge/github.com/dsincl12/wyrand)](https://goreportcard.com/report/github.com/dsincl12/wyrand)
[![GitHub license](https://img.shields.io/github/license/dsincl12/wyrand.svg)](https://github.com/dsincl12/wyrand/blob/master/LICENSE)

# WyRand
WyRand is a thread-safe implementation of the WyRand pseudo-random generator written in Go.

## Installation

```
go get github.com/dsincl12/wyrand
```

## Usage example

```go
import "github.com/dsincl12/wyrand"

func main() {
    w := wyrand.New(uint64(time.Now().UnixNano())) // time seed
    
    fmt.Println(w.Next()) // returns pseudo-random uint64 value
}
```

### Benchmarks
```
BenchmarkNext-4          100000000          10.7 ns/op          0 B/op          0 allocs/op
```

## Credits
Written by [David Sinclair](https://github.com/dsincl12)  
Original by Wang Yi [WyRand](https://github.com/wangyi-fudan/wyhash)  

## License

Apache License 2.0

Copyright (c) 2019 David Sinclair

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.