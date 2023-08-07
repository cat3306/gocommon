package bytespool

import (
	"math/rand"
	"testing"
)

func BenchmarkName(b *testing.B) {

	for i := 1; i < b.N; i++ {
		buff := BUFFERPOOL.Get(uint32(i))

		newB := job(*buff)
		BUFFERPOOL.Put(newB)
	}
}

func BenchmarkMake(b *testing.B) {
	for i := 1; i < b.N; i++ {
		buf := make([]byte, i)
		job(buf)
	}
}
func job(b []byte) []byte {
	i := rand.Intn(cap(b)) + 1
	return b[:i-1]
}
