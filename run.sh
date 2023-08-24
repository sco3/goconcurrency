#!/usr/bin/env -S bash -x -e



function run {
rm -rf cnt
go build 
./cnt

rm -rf cnt
go build -compiler gccgo -gccgoflags " -O3 "  -o cnt main.go
./cnt

rm -rf cnt
}


run > readme.md 2>&1

sed -i 's/\+/\\+/g' readme.md



