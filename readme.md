>>> rm -rf cnt
>>> go build
>>> ./cnt
Sequential Add, Sum: 450013341,  Time Taken: 78.233762ms
cpu: 4
Concurrent Add, Sum: 450013341,  Time Taken: 54.332917ms
>>> rm -rf cnt
>>> go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
>>> ./cnt
Sequential Add, Sum: 450003069,  Time Taken: 88.47569ms
cpu: 4
Concurrent Add, Sum: 450003069,  Time Taken: 133.225237ms
>>> rm -rf cnt
