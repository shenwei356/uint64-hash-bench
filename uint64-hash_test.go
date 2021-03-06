package bench

import (
	"encoding/binary"
	"hash/maphash"
	"sync"
	"testing"
	"math/rand"

	"github.com/cespare/xxhash"
	"github.com/spaolacci/murmur3"
)

var be = binary.BigEndian

var H uint64

func BenchmarkGoMapHash(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64

	var hash maphash.Hash
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		hash.Reset()
		hash.Write(buf)
		h = hash.Sum64()
	}
	H = h
}

func BenchmarkXxhashMethod(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64

	hash := xxhash.New()
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		hash.Reset()
		hash.Write(buf)
		h = hash.Sum64()
	}
	H = h
}

func BenchmarkXxhashFunction(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64

	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		xxhash.Sum64(buf)
	}
	H = h
}

func BenchmarkMurmur3(b *testing.B) {
	buf := make([]byte, 8)
	var h uint64

	hash := murmur3.New64()
	for i := 0; i < b.N; i++ {
		be.PutUint64(buf, uint64(i))

		hash.Reset()
		hash.Write(buf)
		h = hash.Sum64()
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

func BenchmarkMutexLockUnlock(b *testing.B) {
	var h uint64
	var mux sync.Mutex
	for i := 0; i < b.N; i++ {
		mux.Lock()

		h = uint64(i)

		mux.Unlock()
	}
	H = h
}

var F float64
func BenchmarkRandFloat64(b *testing.B) {
	var f float64
	for i := 0; i < b.N; i++ {
		f = rand.Float64()
	}
	F = f
}

func BenchmarkRandFloat64Chanel(b *testing.B) {
	var f float64
	ch := make(chan float64, 5)
	done := make(chan int)
	go func() {
		for i := 0; i < b.N;i++ {
			ch <- rand.Float64()
		}
		done <-1
	}()
	for i := 0; i < b.N; i++ {
		f = <-ch
	}
	F = f
	<-done
}
