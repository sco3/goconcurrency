```
+ rm -rf cnt
+ go build
+ ./cnt
Sequential Add, Sum: 450010025,  Time Taken: 58.707852ms
cpu: 12
Concurrent Add, Sum: 450009999,  Time Taken: 43.741512ms
+ rm -rf cnt
+ go build -compiler gccgo -gccgoflags ' -O3 ' -o cnt main.go
+ ./cnt
Sequential Add, Sum: 450044223,  Time Taken: 70.810421ms
cpu: 12
Concurrent Add, Sum: 450044191,  Time Taken: 43.955251ms
+ rm -rf cnt
```
