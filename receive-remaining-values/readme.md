
### About

This benchmark is to compare the costs between two ways to receive
the remaining values stored in a buffered channel.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_CheckLen-4   	    2000	   1034230 ns/op
Benchmark_Select-4     	    2000	   1048183 ns/op
```

It looks the efficiencies of the two way are most the same.


