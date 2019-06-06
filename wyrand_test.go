package wyrand

import (
	"testing"
	"time"
)

func BenchmarkNext(b *testing.B) {
	b.ReportAllocs()

	w := New(uint64(time.Now().UnixNano()))

	for i := 0; i < b.N; i++ {
		w.Next()
	}
}
