package bench

import (
	"encoding/binary"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/spaolacci/murmur3"
)

var be = binary.BigEndian

func BenchmarkXxhash(b *testing.B) {
	buf := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		xxhash.Sum64(buf)
	}
}

func BenchmarkMurmur3(b *testing.B) {
	buf := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		murmur3.Sum64(buf)
	}
}

func BenchmarkThomasWangUint64Hash(b *testing.B) {
	var key uint64
	for i := 0; i < b.N; i++ {
		key = uint64(i)

		ThomasWangUint64Hash(key)
	}
}

func BenchmarkInverseThomasWangUint64Hash(b *testing.B) {
	var key uint64
	for i := 0; i < b.N; i++ {
		key = uint64(i)

		InverseThomasWangUint64Hash(key)
	}
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
}
