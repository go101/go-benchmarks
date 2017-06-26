
### About

This benchmark is to check the cost of all kinds of sync methods.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
Benchmark_NoSync-4                            	1000000000	         2.84 ns/op
Benchmark_Atomic-4                            	100000000	        10.1 ns/op
Benchmark_Mutex-4                             	50000000	        26.5 ns/op
Benchmark_Channel-4                           	20000000	        85.4 ns/op
Benchmark_Select_OneCase-4                    	20000000	        85.0 ns/op
Benchmark_Select_TwoCases-4                   	 5000000	       252 ns/op
Benchmark_Select_TwoCases_Plus_TrySent-4      	 5000000	       264 ns/op
Benchmark_Select_TwoCases_Plus_TryReceive-4   	 5000000	       261 ns/op
Benchmark_Select_ThreeCases-4                 	 5000000	       338 ns/op
```

From the result, we can get an impression that
* the cost of mutex lock+unlock is three times of atomic.
* the cost of channel send+receive operations are three times of mutex.
* one-case select will be translated to a simple channel send/receive op.
* for mult-case selects, one more case means 1.5 times more cost of channel send+receive operations.
* the cost of one-case-plus-default (try-send and try receive) select is about equal to
  (a little less than) the cost of mutex lock+unlock.


