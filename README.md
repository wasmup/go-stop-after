# Stop a gouroutine after a period of time
```sh
BenchmarkAfterFunc-8        	1000000000 0.4468 ns/op  0 B/op  0 allocs/op
BenchmarkDoneChannel-8      	121966824   9.855 ns/op  0 B/op  0 allocs/op
BenchmarkTimeSince-8        	89790115    12.95 ns/op  0 B/op  0 allocs/op
BenchmarkContextErr-8       	58508900    19.78 ns/op  0 B/op  0 allocs/op
BenchmarkAfterFuncMutex-8   	58323207    20.00 ns/op  0 B/op  0 allocs/op
BenchmarkContext-8          	48947625    27.43 ns/op  0 B/op  0 allocs/op
```