
### About

This benchmark is to compare different methods to stat the rune count in a string.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.2 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_CountStringRunes_Var_ForRange-4         	 5000000	       370 ns/op
Benchmark_CountStringRunes_Var_Utf8Lib-4          	 5000000	       360 ns/op
Benchmark_CountStringRunes_Var_LenRuneSlice-4     	 1000000	      1268 ns/op
Benchmark_CountStringRunes_Const_ForRange-4       	 5000000	       376 ns/op
Benchmark_CountStringRunes_Const_Utf8Lib-4        	 5000000	       356 ns/op
Benchmark_CountStringRunes_Const_LenRuneSlice-4   	100000000	        14.5 ns/op
Benchmark_CountStringRunes_Var_LenByteSlice-4     	10000000	       156 ns/op
Benchmark_CountStringRunes_Var_LenString-4        	2000000000	         0.88 ns/op
```

```
$ go version
go version devel +4fe688c6a4 Mon May 28 03:35:36 2018 +0000 linux/amd64
$ go test -bench=.
goos: linux
goarch: amd64
Benchmark_CountStringRunes_Var_ForRange-4         	 5000000	       375 ns/op
Benchmark_CountStringRunes_Var_Utf8Lib-4          	 5000000	       365 ns/op
Benchmark_CountStringRunes_Var_LenRuneSlice-4     	 5000000	       378 ns/op
Benchmark_CountStringRunes_Const_ForRange-4       	 5000000	       367 ns/op
Benchmark_CountStringRunes_Const_Utf8Lib-4        	 5000000	       364 ns/op
Benchmark_CountStringRunes_Const_LenRuneSlice-4   	2000000000	         0.88 ns/op
Benchmark_CountStringRunes_Var_LenByteSlice-4     	10000000	       160 ns/op
Benchmark_CountStringRunes_Var_LenString-4        	2000000000	         0.88 ns/op
```

Since 1.11, `len([]rune(aString)` is optimized.
It is strange that `len([]byte(aString))` is not optimized.


