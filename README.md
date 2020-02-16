# Benchmark of three uint64 hash functions

## Hash functions

- [maphsh](https://tip.golang.org/pkg/hash/maphash/), go1.14rc1
- [xxhash](https://github.com/cespare/xxhash), version [d7df741](https://github.com/cespare/xxhash/tree/d7df74196a9e781ede915320c11c378c1b2f3a1f)
- [murmur3](https://github.com/spaolacci/murmur3), version [539464a](https://github.com/spaolacci/murmur3/tree/539464a789e9b9f01bc857458ffe2c5c1a2ed382)
- [Thomas Wang’s 64-bit integer hash function](https://gist.github.com/badboy/6267743)
- [Inverse of Thomas Wang’s 64-bit integer hash function](https://naml.us/post/inverse-of-a-hash-function/)

## Result

```
$ go test . -bench=Benchmark* -benchmem
goos: linux
goarch: amd64
pkg: github.com/shenwei356/uint64-hash-bench
BenchmarkGoMapHash-16                           74581323                16.2 ns/op             0 B/op          0 allocs/op
BenchmarkXxhashMethod-16                        107216905               11.6 ns/op             0 B/op          0 allocs/op
BenchmarkXxhashFunction-16                      282878127                4.15 ns/op            0 B/op          0 allocs/op
BenchmarkMurmur3-16                             42436113                29.0 ns/op             0 B/op          0 allocs/op
BenchmarkThomasWangUint64Hash-16                777305191                1.52 ns/op            0 B/op          0 allocs/op
BenchmarkInverseThomasWangUint64Hash-16         231070869                5.14 ns/op            0 B/op          0 allocs/op

BenchmarkMutexLockUnlock-16                     88848709                13.6 ns/op             0 B/op          0 allocs/op
```