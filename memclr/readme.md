
### About

This benchmark is to check how efficient
<a href="http://www.tapirgames.com/blog/golang-array-iterations">the memclr optimization</a> is.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.8.3 linux/amd64
$ go test -bench=.
N = 102400
Benchmark_ArrayPtr_N-4       	   20000	     89876 ns/op
Benchmark_Memclr_Array_N-4   	   30000	     49816 ns/op
Benchmark_Memclr_Slice_N-4   	   30000	     51804 ns/op
Benchmark_/ArrayPtr_16777216-4         	     100	  17251480 ns/op
Benchmark_/ArrayPtr_8388608-4          	     200	   8645346 ns/op
Benchmark_/ArrayPtr_4194304-4          	     300	   4315414 ns/op
Benchmark_/ArrayPtr_2097152-4          	    1000	   2155097 ns/op
Benchmark_/ArrayPtr_1048576-4          	    2000	   1080005 ns/op
Benchmark_/ArrayPtr_524288-4           	    3000	    494772 ns/op
Benchmark_/ArrayPtr_262144-4           	   10000	    231480 ns/op
Benchmark_/ArrayPtr_131072-4           	   10000	    114633 ns/op
Benchmark_/ArrayPtr_65536-4            	   30000	     57341 ns/op
Benchmark_/ArrayPtr_32768-4            	   50000	     28651 ns/op
Benchmark_/ArrayPtr_16384-4            	  100000	     14318 ns/op
Benchmark_/ArrayPtr_8192-4             	  200000	      7306 ns/op
Benchmark_/ArrayPtr_4096-4             	  500000	      3592 ns/op
Benchmark_/ArrayPtr_2048-4             	 1000000	      1804 ns/op
Benchmark_/ArrayPtr_1024-4             	 1000000	      1053 ns/op
Benchmark_/ArrayPtr_512-4              	 3000000	       460 ns/op
Benchmark_/ArrayPtr_256-4              	10000000	       241 ns/op
Benchmark_/ArrayPtr_128-4              	10000000	       122 ns/op
Benchmark_/ArrayPtr_64-4               	20000000	        66.7 ns/op
Benchmark_/ArrayPtr_32-4               	50000000	        29.2 ns/op
Benchmark_/ArrayPtr_16-4               	100000000	        15.3 ns/op
Benchmark_/ArrayPtr_8-4                	200000000	         8.33 ns/op
Benchmark_/ArrayPtr_4-4                	300000000	         4.38 ns/op
Benchmark_/ArrayPtr_2-4                	1000000000	         2.63 ns/op
Benchmark_/Memclr_Slice_16777216-4    	     100	  17478608 ns/op
Benchmark_/Memclr_Slice_8388608-4     	     200	   8730535 ns/op
Benchmark_/Memclr_Slice_4194304-4     	     300	   4363549 ns/op
Benchmark_/Memclr_Slice_2097152-4     	    1000	   2173571 ns/op
Benchmark_/Memclr_Slice_1048576-4     	    2000	   1084490 ns/op
Benchmark_/Memclr_Slice_524288-4      	    3000	    457028 ns/op
Benchmark_/Memclr_Slice_262144-4      	   10000	    139271 ns/op
Benchmark_/Memclr_Slice_131072-4      	   20000	     63758 ns/op
Benchmark_/Memclr_Slice_65536-4       	   50000	     31289 ns/op
Benchmark_/Memclr_Slice_32768-4       	  100000	     12749 ns/op
Benchmark_/Memclr_Slice_16384-4       	  200000	      6094 ns/op
Benchmark_/Memclr_Slice_8192-4        	  500000	      2982 ns/op
Benchmark_/Memclr_Slice_4096-4        	 1000000	      1055 ns/op
Benchmark_/Memclr_Slice_2048-4        	 3000000	       466 ns/op
Benchmark_/Memclr_Slice_1024-4        	10000000	       231 ns/op
Benchmark_/Memclr_Slice_512-4         	10000000	       118 ns/op
Benchmark_/Memclr_Slice_256-4         	20000000	        60.4 ns/op
Benchmark_/Memclr_Slice_128-4         	50000000	        33.0 ns/op
Benchmark_/Memclr_Slice_64-4          	100000000	        17.1 ns/op
Benchmark_/Memclr_Slice_32-4          	100000000	        10.9 ns/op
Benchmark_/Memclr_Slice_16-4          	200000000	         9.17 ns/op
Benchmark_/Memclr_Slice_8-4           	200000000	         8.73 ns/op
Benchmark_/Memclr_Slice_4-4           	200000000	         7.86 ns/op
Benchmark_/Memclr_Slice_2-4           	200000000	         7.41 ns/op
```

On the test machine, when the length of the container is at less 1M,
the memclr version is even slower than the general loop version.
However, some people report that the memclr version is constantly
more efficient on their machines.

