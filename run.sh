#!/usr/bin/env -S bash -x



function run {
rm -rf cnt
go build 
./cnt && rm -f cnt

#rm -rf cnt
#go build -compiler gccgo -gccgoflags " -O3 "  -o cnt main.go
#./cnt
#cat /proc/loadavg

}


run > readme.md 2>&1

sed -i '1 i\```' readme.md
echo '```' >> readme.md 
cat readme.md


