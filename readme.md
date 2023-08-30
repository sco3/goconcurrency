```
+ rm -f goconcurrency
+ go build
+ ./goconcurrency 5e8
Generated 500000000 numbers, Time Taken: 13.062703705s
Load avergate: 1.46 1.85 1.47 2/1523 12383

Sequential Add, Sum: 2249974319,  Time Taken: 382.815859ms
Load avergate: 1.46 1.85 1.47 12/1523 12384


Max processors: 1
Concurrent Add, Sum: 2249974319,  Time Taken: 505.685869ms
Load avergate: 1.46 1.85 1.47 5/1522 12385


Max processors: 2
Concurrent Add, Sum: 2249974319,  Time Taken: 313.387707ms
Load avergate: 1.46 1.85 1.47 2/1522 12386


Max processors: 3
Concurrent Add, Sum: 2249974319,  Time Taken: 272.832454ms
Load avergate: 1.46 1.85 1.47 2/1522 12387

+ rm -f goconcurrency
```
