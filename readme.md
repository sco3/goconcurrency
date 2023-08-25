> rm -rf cnt
> go build
> ./cnt
Sequential Add, Sum: 450001386,  Time Taken: 63.683241ms
cpu: 12
Concurrent Add, Sum: 450001369,  Time Taken: 43.54735ms
> rm -rf cnt
> go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
> ./cnt
Sequential Add, Sum: 449945728,  Time Taken: 70.85446ms
cpu: 12
Concurrent Add, Sum: 449945715,  Time Taken: 43.813169ms
> rm -rf cnt
