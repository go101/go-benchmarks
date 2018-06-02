
### About

This benchmark is to check the copy costs of struct values with different number of fields.

### Result

```
$ cat /proc/cpuinfo | grep 'model name' | uniq
model name	: Intel(R) Core(TM) i3-2350M CPU @ 2.30GHz
$ go version
go version go1.10.2 linux/amd64
$ go test -bench=.
Benchmark_Loop_2-4               	  500000	      3150 ns/op
Benchmark_Range_OneIterVar_2-4   	  500000	      3206 ns/op
Benchmark_Range_TwoIterVar_2-4   	  500000	      3201 ns/op
Benchmark_Loop_3-4               	  500000	      3059 ns/op
Benchmark_Range_OneIterVar_3-4   	  500000	      3149 ns/op
Benchmark_Range_TwoIterVar_3-4   	  500000	      3143 ns/op
Benchmark_Loop_4-4               	  500000	      3145 ns/op
Benchmark_Range_OneIterVar_4-4   	  500000	      3087 ns/op
Benchmark_Range_TwoIterVar_4-4   	  500000	      3062 ns/op
Benchmark_Loop_5-4               	  500000	      3238 ns/op
Benchmark_Range_OneIterVar_5-4   	  500000	      3183 ns/op
Benchmark_Range_TwoIterVar_5-4   	  200000	      7069 ns/op
Benchmark_Loop_6-4               	  500000	      3230 ns/op
Benchmark_Range_OneIterVar_6-4   	  500000	      3137 ns/op
Benchmark_Range_TwoIterVar_6-4   	  200000	      7107 ns/op
```

It looks that the official Go runtime views struct values
with less than 5 native-word-size fields as small-sized values.
The three ways to interate a slice with small-sized element values
almost perform the same. However, the for-range loop with
two loop variables is much less efficient than the other ways
for iterating a slice with non-small-sized element values.

