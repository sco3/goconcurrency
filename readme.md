```
+ rm -f goconcurrency
+ go build
+ ./goconcurrency 5e8
Generated 500000000 numbers, Time Taken: 13.072440055s
Load avergate: 0.98 1.65 1.42 3/1563 12593

Sequential Add, Sum: 2249932838,  Time Taken: 372.157833ms
Load avergate: 0.98 1.65 1.42 2/1564 12595


Max processors: 1
Concurrent Add, Sum: 2249932838,  Time Taken: 373.192117ms
Load avergate: 0.98 1.65 1.42 2/1564 12596


Max processors: 2
Concurrent Add, Sum: 2249932838,  Time Taken: 272.291565ms
Load avergate: 0.98 1.65 1.42 11/1564 12597


Max processors: 3
Concurrent Add, Sum: 2249932838,  Time Taken: 271.767156ms
Load avergate: 0.98 1.65 1.42 2/1564 12598

+ rm -f goconcurrency
```
