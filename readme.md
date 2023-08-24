+ rm -rf cnt
+ go build
+ ./cnt
Sequential Add, Sum: 449985516,  Time Taken: 76.760566ms
cpu: 4
Concurrent Add, Sum: 449985516,  Time Taken: 56.259691ms
+ rm -rf cnt
+ go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
+ ./cnt
Sequential Add, Sum: 450036169,  Time Taken: 88.451501ms
cpu: 4
Concurrent Add, Sum: 450036169,  Time Taken: 140.150352ms
+ rm -rf cnt
