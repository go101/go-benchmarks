
### About

This benchmark is to benchmark several ways of using methods as function values.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.2 linux/amd64
$ go test -bench=.
M = 8, N = 8192
goos: linux
goarch: amd64
Benchmark_TypeMethodAsFunc-4              	  300000	      5311 ns/op
Benchmark_ValueMethodAsFunc-4             	   10000	    131127 ns/op
Benchmark_ValueMethodAsFunc2-4            	  200000	      6170 ns/op
Benchmark_ValueMethodAsFunc_GlobalVar-4   	   10000	    128795 ns/opp
```

It looks, if the function value escapes to heap,
then the assigning a method to this function is costly.

```
$ go build -gcflags -m func-value.go 
# command-line-arguments
./func-value.go:49:6: can inline (*myType).myMethod
./func-value.go:14:19: can inline (*myType).("".myMethod)-fm
./func-value.go:14:19: inlining call to (*myType).myMethod
./func-value.go:51:6: can inline main
./func-value.go:14:19: myObj.myMethod escapes to heap
./func-value.go:16:20: myObj2.myMethod escapes to heap
./func-value.go:24:20: fy2 myObj.myMethod does not escape
./func-value.go:26:20: fy2 myObj2.myMethod does not escape
./func-value.go:33:19: myObj.myMethod escapes to heap
./func-value.go:35:20: myObj2.myMethod escapes to heap
./func-value.go:49:10: (*myType).myMethod m does not escape
```

A value method needs internal structure to hold,
but a type method just needs a pointer to a global internal structure value to hold.


