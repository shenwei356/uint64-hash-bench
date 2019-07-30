package bench

import (
	"encoding/binary"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/spaolacci/murmur3"
)

var be = binary.BigEndian

var H uint64

func BenchmarkXxhash(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		h = xxhash.Sum64(buf)
	}
	H = h
}

func BenchmarkMurmur3(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		h = murmur3.Sum64(buf)
	}
	H = h
}

func BenchmarkThomasWangUint64Hash(b *testing.B) {
	var key uint64
	var h uint64
	for i := 0; i < b.N; i++ {
		key = uint64(i)

		h = ThomasWangUint64Hash(key)
	}
	H = h
}

func BenchmarkInverseThomasWangUint64Hash(b *testing.B) {
	var key uint64
	var h uint64
	for i := 0; i < b.N; i++ {
		key = uint64(i)

		h = InverseThomasWangUint64Hash(key)
	}
	H = h
}

func TestThomasWangUint64Hahs(t *testing.T) {
	var h, j uint64
	for i := uint64(0); i < 100000000; i++ {
		h = ThomasWangUint64Hash(i)
		j = InverseThomasWangUint64Hash(h)

		if j != i {
			t.Errorf("error when testing for %d", i)
		}
	}
	H = h
}
