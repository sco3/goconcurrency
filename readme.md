```
+ rm -rf goconcurrency
+ go build
+ ./goconcurrency 5e7
Generated 50000000 numbers, Time Taken: 1.273327559s
Load avergate: 0.33 0.36 0.45 4/1251 9304

Sequential Add, Sum: 224990373,  Time Taken: 36.152895ms
Load avergate: 0.33 0.36 0.45 2/1252 9306


Max processors: 1
Concurrent Add, Sum: 224990373,  Time Taken: 36.924844ms
Load avergate: 0.33 0.36 0.45 5/1252 9307


Max processors: 2
Concurrent Add, Sum: 224990373,  Time Taken: 26.797675ms
Load avergate: 0.33 0.36 0.45 2/1252 9308


Max processors: 3
Concurrent Add, Sum: 224990373,  Time Taken: 27.049865ms
Load avergate: 0.33 0.36 0.45 4/1252 9309

+ rm -f goconcurrency
```
