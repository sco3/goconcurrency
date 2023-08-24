\+ rm -rf cnt
\+ go build
\+ ./cnt
Sequential Add, Sum: 450010337,  Time Taken: 75.39752ms
cpu: 4
Concurrent Add, Sum: 450010337,  Time Taken: 57.072392ms
\+ rm -rf cnt
\+ go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
\+ ./cnt
Sequential Add, Sum: 449998772,  Time Taken: 88.360735ms
cpu: 4
Concurrent Add, Sum: 449998772,  Time Taken: 129.353267ms
\+ rm -rf cnt
