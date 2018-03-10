
### About

This benchmark is to compare the performances of different string concatenation methods.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.9.2 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_AddStrings_KnownCount-4          	20000000	       119 ns/op
Benchmark_AddStrings_UnnownCount-4         	 5000000	       301 ns/op
Benchmark_FmtSprintf_KnownCount-4          	 3000000	       608 ns/op
Benchmark_StringsJoin_NoSliceAllocated-4   	10000000	       186 ns/op
Benchmark_StringsJoin_SliceAllocated-4     	 3000000	       415 ns/op
Benchmark_BytesBuffer_New-4                	 5000000	       332 ns/op
Benchmark_AddStrings_Buffer_Reuse-4        	10000000	       126 ns/op
```

Generally, the plain `a+b+...` method is the most performant.
For non-constant number of strings to be concatenated,
the `bytes.Buffer` is the most performant if buffer is allowed to be reused.
If the strings to be concatenated are already stored in a `[]string` slice
and the number of strings is not constant,
then the `strings.Join` is the most performant.


