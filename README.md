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
BenchmarkXxhash-16                              300000000                4.17 ns/op            0 B/op          0 allocs/op
BenchmarkMurmur3-16                             100000000               16.9 ns/op             0 B/op          0 allocs/op
BenchmarkThomasWangUint64Hash-16                2000000000               1.46 ns/op            0 B/op          0 allocs/op
BenchmarkInverseThomasWangUint64Hash-16         300000000                5.06 ns/op            0 B/op          0 allocs/op
```