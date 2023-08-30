```
+ rm -rf cnt
+ go build
+ ./cnt
Generated 500000000 numbers, Time Taken: 13.516958966s
Load avergate: 0.94 4.52 3.58 4/1502 15432

Sequential Add, Sum: 2250104173,  Time Taken: 377.34533ms
Load avergate: 0.94 4.52 3.58 4/1503 15434


Max processors: 1
Concurrent Add, Sum: 2250104173,  Time Taken: 369.997009ms
Load avergate: 0.94 4.52 3.58 2/1503 15435


Max processors: 2
Concurrent Add, Sum: 2250104173,  Time Taken: 272.278887ms
Load avergate: 0.94 4.52 3.58 4/1503 15436


Max processors: 3
Concurrent Add, Sum: 2250104173,  Time Taken: 271.21523ms
Load avergate: 0.94 4.52 3.58 2/1503 15437


Max processors: 4
Concurrent Add, Sum: 2250104173,  Time Taken: 266.86153ms
Load avergate: 0.94 4.52 3.58 2/1503 15438

+ rm -f cnt
```
