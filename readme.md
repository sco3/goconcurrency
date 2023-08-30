```
+ rm -f goconcurrency
+ go build
+ ./goconcurrency 5e8
Generated 500000000 numbers, Time Taken: 12.860580014s
Load avergate: 1.52 2.04 1.39 2/1114 5774

Sequential Add, Sum: 2249908007,  Time Taken: 367.775041ms
Load avergate: 1.52 2.04 1.39 4/1115 5776


Max processors: 1
Concurrent Add, Sum: 2249908007,  Time Taken: 370.164899ms
Load avergate: 1.52 2.04 1.39 2/1115 5777


Max processors: 2
Concurrent Add, Sum: 2249908007,  Time Taken: 267.43456ms
Load avergate: 1.52 2.04 1.39 2/1115 5778


Max processors: 3
Concurrent Add, Sum: 2249908007,  Time Taken: 268.890986ms
Load avergate: 1.52 2.04 1.39 2/1115 5779

+ rm -f goconcurrency
```
