# Benchmark of three uint64 hash functions

## Hash functions

- [xxhash](https://github.com/cespare/xxhash), version [32f3a11](https://github.com/cespare/xxhash/tree/32f3a100fff237d14029d34ec8579ef02aaaf87c)
- [murmur3](https://github.com/spaolacci/murmur3), version [539464a](https://github.com/spaolacci/murmur3/tree/539464a789e9b9f01bc857458ffe2c5c1a2ed382)
- [Thomas Wang’s 64-bit integer hash function](https://gist.github.com/badboy/6267743)
- [Inverse of Thomas Wang’s 64-bit integer hash function](https://naml.us/post/inverse-of-a-hash-function/)

## Result

```
$ go test . -bench=Benchmark* -benchmem
goos: linux
goarch: amd64
pkg: github.com/shenwei356/uint64-hash-bench
BenchmarkXxhash-4                               200000000                6.73 ns/op            0 B/op          0 allocs/op
BenchmarkMurmur3-4                              50000000                 24.2 ns/op            0 B/op          0 allocs/op
BenchmarkThomasWangUint64Hash-4                 2000000000               0.32 ns/op            0 B/op          0 allocs/op
BenchmarkInverseThomasWangUint64Hash-4          200000000                7.27 ns/op            0 B/op          0 allocs/op
```