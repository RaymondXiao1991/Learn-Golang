package mapbench

import (
	"math/rand"
	"testing"
	"time"
)

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinMapStoreParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinMapStoreParallel
BenchmarkBuiltinMapStoreParallel-8       6708006               173.1 ns/op             5 B/op          0 allocs/op
PASS
ok      mapbench        1.736s
*/
func BenchmarkBuiltinMapStoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinMapStore(k, k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinRwMapStoreParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinRwMapStoreParallel
BenchmarkBuiltinRwMapStoreParallel-8     6035082               205.8 ns/op            12 B/op          0 allocs/op
PASS
ok      mapbench        2.042s
*/
func BenchmarkBuiltinRwMapStoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinRwMapStore(k, k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkSyncMapStoreParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkSyncMapStoreParallel
BenchmarkSyncMapStoreParallel-8           4241119               342.5 ns/op            46 B/op          3 allocs/op
PASS
ok      mapbench        2.137s
*/
func BenchmarkSyncMapStoreParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			syncMapStore(k, k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinMapLookupParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinMapLookupParallel
BenchmarkBuiltinMapLookupParallel-8     13381762                90.69 ns/op            0 B/op          0 allocs/op
PASS
ok      mapbench        2.857s
*/
func BenchmarkBuiltinMapLookupParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinMapLookup(k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinRwMapLookupParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinRwMapLookupParallel
BenchmarkBuiltinRwMapLookupParallel-8           13459736                83.46 ns/op            0 B/op          0 allocs/op
PASS
ok      mapbench        1.804s
*/
func BenchmarkBuiltinRwMapLookupParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinRwMapLookup(k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkSyncMapLookupParallel$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkSyncMapLookupParallel
BenchmarkSyncMapLookupParallel-8         300566817                4.021 ns/op           0 B/op          0 allocs/op
PASS
ok      mapbench        2.191s
*/
func BenchmarkSyncMapLookupParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			syncMapLookup(k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinMapDelete$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinMapDelete
BenchmarkBuiltinMapDelete-8     12903069                86.63 ns/op            0 B/op          0 allocs/op
PASS
ok      mapbench        1.326s
*/
func BenchmarkBuiltinMapDelete(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinMapDelete(k)
		}
	})
}

/*
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkBuiltinRwMapDelete$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkBuiltinRwMapDelete
BenchmarkBuiltinRwMapDelete-8            9161419               130.4 ns/op             0 B/op          0 allocs/op
PASS
ok      mapbench        2.918s
*/
func BenchmarkBuiltinRwMapDelete(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			builtinRwMapDelete(k)
		}
	})
}

/*Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkSyncMapDelete$ mapbench

goos: darwin
goarch: arm64
pkg: mapbench
BenchmarkSyncMapDelete
BenchmarkSyncMapDelete-8        302859402                3.799 ns/op           0 B/op          0 allocs/op
PASS
ok      mapbench        2.209s
*/
func BenchmarkSyncMapDelete(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for pb.Next() {
			k := r.Intn(100000000)
			syncMapDelete(k)
		}
	})
}