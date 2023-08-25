>+ rm -rf cnt
>+ go build
>+ ./cnt
Sequential Add, Sum: 450032660,  Time Taken: 57.698831ms
cpu: 12
Concurrent Add, Sum: 450032644,  Time Taken: 43.534476ms
>+ rm -rf cnt
>+ go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
>+ ./cnt
Sequential Add, Sum: 449997688,  Time Taken: 71.077362ms
cpu: 12
Concurrent Add, Sum: 449997672,  Time Taken: 44.02399ms
>+ rm -rf cnt
