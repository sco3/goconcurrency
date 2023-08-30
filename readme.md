```
+ rm -rf goconcurrency
+ go build
+ ./goconcurrency 5e8
Generated 500000000 numbers, Time Taken: 12.642730176s
Load avergate: 0.32 0.33 0.43 7/1248 9450

Sequential Add, Sum: 2249956274,  Time Taken: 368.478841ms
Load avergate: 0.32 0.33 0.43 4/1248 9451


Max processors: 1
Concurrent Add, Sum: 2249956274,  Time Taken: 369.402629ms
Load avergate: 0.32 0.33 0.43 5/1248 9452


Max processors: 2
Concurrent Add, Sum: 2249956274,  Time Taken: 269.479378ms
Load avergate: 0.32 0.33 0.43 4/1248 9453


Max processors: 3
Concurrent Add, Sum: 2249956274,  Time Taken: 272.388705ms
Load avergate: 0.32 0.33 0.43 2/1248 9454

+ rm -f goconcurrency
```
