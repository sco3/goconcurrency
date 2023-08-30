#!/usr/bin/env -S bash -x



function run {
rm -rf goconcurrency
go build 
./goconcurrency 5e7 && rm -f goconcurrency

#rm -rf goconcurrency
#go build -compiler gccgo -gccgoflags " -O3 "  -o goconcurrency main.go
#./goconcurrency

}


run > readme.md 2>&1

sed -i '1 i\```' readme.md
echo '```' >> readme.md 
cat readme.md


