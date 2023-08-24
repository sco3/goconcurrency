> rm -rf cnt
> go build
> ./cnt
Sequential Add, Sum: 449993345,  Time Taken: 76.258825ms
cpu: 4
Concurrent Add, Sum: 449993345,  Time Taken: 54.742231ms
> rm -rf cnt
> go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
> ./cnt
Sequential Add, Sum: 449987838,  Time Taken: 89.333084ms
cpu: 4
Concurrent Add, Sum: 449987838,  Time Taken: 129.348793ms
> rm -rf cnt
