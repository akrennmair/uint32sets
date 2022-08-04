# Benchmarking uint32 set implementations

This is a set of uint32 set implementations that I created for benchmarking purposes.

## Problem

I had a problem that boiled down to 2 sub-problems:

* cardinality of set E, where E is the set of all values (i.e. the count-distinct problem)
* is x ∈ E ?

In the specific program, the `add` operation to the set took up the vast 
majority of the total CPU time, i.e. `adding` elements to the set as fast as 
possible is the #1 point for optimization. Also, counting had to be precise,
so probabilistic data structure like HyperLogLog and Bloom filters were out of 
the question.

## Benchmark Results

This was on my machine (Mid-2015 MacBook Pro):

```text
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/akrennmair/uint32sets
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkAll/listSet-8         	      12	  95979117 ns/op	10577912 B/op	      35 allocs/op
BenchmarkAll/mapSet-8          	      68	  15157180 ns/op	  472802 B/op	     497 allocs/op
BenchmarkAll/slotSet-8         	     199	   5860525 ns/op	 3303391 B/op	   18618 allocs/op
BenchmarkAll/slotPreallocSet-8 	     180	   6370050 ns/op	 3747555 B/op	   18552 allocs/op
BenchmarkAll/slotBitmapSet-8   	      27	  43413609 ns/op	141828103 B/op	   17249 allocs/op
PASS
ok  	github.com/akrennmair/uint32sets	8.689s
```
