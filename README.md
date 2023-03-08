# bench-golang-validators

`go test -bench .`

```
goos: windows
goarch: amd64
pkg: benchGolangValidators
BenchmarkAsaskevich-10    	  130369	      8267 ns/op	    5321 B/op	      69 allocs/op
BenchmarkPlayground-10    	 1000000	      1096 ns/op	     114 B/op	       7 allocs/op
BenchmarkSaddam-10        	  139081	      8671 ns/op	    1744 B/op	      60 allocs/op
BenchmarkBuffalo-10       	  284940	      3952 ns/op	    1474 B/op	      33 allocs/op
BenchmarkFender-10        	 1330290	       909 ns/op	     740 B/op	      18 allocs/op
BenchmarkOzzo-10          	  256746	      4633 ns/op	    5029 B/op	      99 allocs/op

PASS
ok      benchGolangValidators   7.300s
```
