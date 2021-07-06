
### About

This benchmark is to check the copy costs of struct/array values with different number of fields/elements.

### Result

```
$ go version
go version go1.16.3 linux/amd64
$ $ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/go101/go-benchmarks/copy-cost
cpu: Intel(R) Core(TM) i5-4210U CPU @ 1.70GHz
Benchmark_CopyArray_b1-4        	1000000000	         0.4144 ns/op
Benchmark_CopyArray_b2-4        	1000000000	         0.4339 ns/op
Benchmark_CopyArray_b3-4        	1000000000	         0.7557 ns/op
Benchmark_CopyArray_b4-4        	1000000000	         0.7602 ns/op
Benchmark_CopyArray_b5-4        	1000000000	         1.125 ns/op
Benchmark_CopyArray_b6-4        	1000000000	         1.130 ns/op
Benchmark_CopyArray_b7-4        	768449886	         1.503 ns/op
Benchmark_CopyArray_b8-4        	790409236	         1.506 ns/op
Benchmark_CopyArray_b9-4        	447599713	         2.627 ns/op
Benchmark_CopyArray1-4          	1000000000	         0.5254 ns/op
Benchmark_CopyArray2-4          	1000000000	         0.5255 ns/op
Benchmark_CopyArray3-4          	1000000000	         0.7509 ns/op
Benchmark_CopyArray4-4          	1000000000	         0.7500 ns/op
Benchmark_CopyArray5-4          	1000000000	         1.127 ns/op
Benchmark_CopyArray6-4          	1000000000	         1.124 ns/op
Benchmark_CopyArray7-4          	783194616	         1.502 ns/op
Benchmark_CopyArray8-4          	781279333	         1.505 ns/op
Benchmark_CopyArray9-4          	454654815	         2.631 ns/op
Benchmark_CopyStruct_b1-4       	1000000000	         0.4600 ns/op
Benchmark_CopyStruct_b2-4       	1000000000	         0.4383 ns/op
Benchmark_CopyStruct_b3-4       	1000000000	         0.7527 ns/op
Benchmark_CopyStruct_b4-4       	1000000000	         0.7518 ns/op
Benchmark_CopyStruct_b5-4       	1000000000	         1.126 ns/op
Benchmark_CopyStruct_b6-4       	1000000000	         1.130 ns/op
Benchmark_CopyStruct_b7-4       	780992246	         1.522 ns/op
Benchmark_CopyStruct_b8-4       	770636612	         1.518 ns/op
Benchmark_CopyStruct_b9-4       	449050116	         2.668 ns/op
Benchmark_CopyStruct1-4         	1000000000	         0.5255 ns/op
Benchmark_CopyStruct2-4         	1000000000	         0.7594 ns/op
Benchmark_CopyStruct3-4         	1000000000	         1.154 ns/op
Benchmark_CopyStruct4-4         	776528011	         1.578 ns/op
Benchmark_CopyStruct5-4         	1000000000	         1.150 ns/op
Benchmark_CopyStruct6-4         	1000000000	         1.188 ns/op
Benchmark_CopyStruct7-4         	747694483	         1.556 ns/op
Benchmark_CopyStruct8-4         	772873843	         1.530 ns/op
Benchmark_CopyStruct9-4         	446456098	         2.654 ns/op
Benchmark_RangeArraySlice1-4    	19898918	        57.23 ns/op
Benchmark_RangeArraySlice2-4    	14207414	        83.02 ns/op
Benchmark_RangeArraySlice3-4    	 9576225	       122.6 ns/op
Benchmark_RangeArraySlice4-4    	 6338328	       194.7 ns/op
Benchmark_RangeArraySlice5-4    	 7340664	       158.0 ns/op
Benchmark_RangeArraySlice6-4    	 5014474	       238.7 ns/op
Benchmark_RangeArraySlice7-4    	 5476196	       216.3 ns/op
Benchmark_RangeArraySlice8-4    	 4458610	       265.7 ns/op
Benchmark_RangeArraySlice9-4    	 3897891	       305.2 ns/op
Benchmark_RangeStructSlice1-4   	20163406	        59.66 ns/op
Benchmark_RangeStructSlice2-4   	18116866	        63.28 ns/op
Benchmark_RangeStructSlice3-4   	18559531	        63.11 ns/op
Benchmark_RangeStructSlice4-4   	18548559	        63.03 ns/op
Benchmark_RangeStructSlice5-4   	 7502433	       157.4 ns/op
Benchmark_RangeStructSlice6-4   	 5031660	       237.2 ns/op
Benchmark_RangeStructSlice7-4   	 5509942	       215.7 ns/op
Benchmark_RangeStructSlice8-4   	 4459640	       266.8 ns/op
Benchmark_RangeStructSlice9-4   	 3898930	       305.5 ns/op
```

It looks that the official Go runtime views struct values
with less than 5 native-word-size fields as small-sized values.

