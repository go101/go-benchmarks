
### About

This benchmark is to check the efficiency differences between
using `append` and `make` to create slices.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
Benchmark_AllocWithMake-4     	    1000	   1305125 ns/op
Benchmark_AllocWithAppend-4   	    2000	    796847 ns/op
```

It looks create a slice by using `append` is more efficient than
by using `copy`. This is some counter-intuitive.


