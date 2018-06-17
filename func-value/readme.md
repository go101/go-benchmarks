
### About

This benchmark is to benchmark several ways of using methods as function values.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz

$ go version
go version go1.10.3 linux/amd64

$ go test -bench=. 
goos: linux
goarch: amd64
Benchmark_TypeMethodAsFunc-4     	  300000	      5276 ns/op
Benchmark_ValueMethodAsFunc-4    	   10000	    141505 ns/op
Benchmark_ValueMethodAsFunc2-4   	  200000	     10122 ns/op
Benchmark_DirectMethod-4         	  300000	      5282 ns/op

$ go build -gcflags -m func-value.go
# command-line-arguments
./func-value.go:46:6: can inline main
./func-value.go:5:10: (*myType).myMethod m does not escape
./func-value.go:23:19: myObj.myMethod escapes to heap
./func-value.go:25:20: myObj2.myMethod escapes to heap
./func-value.go:32:20: fz myObj.myMethod does not escape
./func-value.go:34:20: fz myObj2.myMethod does not escape
```

It looks, if the function value escapes to heap,
then the assigning a method to this function is costly.

A value method needs internal structure to hold,
but a type method just needs a pointer to a global internal structure value to hold.
A direct method call is equivalent to a type method call.


