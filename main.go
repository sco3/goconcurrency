package main

import (
	"fmt"
	"goconcurrency/counting"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func getLoadAvg() {
	out, err := exec.Command("cat", "/proc/loadavg").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Load avergate: %s\n", out)

}

func main() {
	t := time.Now()

	var n int64 = 1000

	if len(os.Args) > 1 {
		//fmt.Printf("args: %v\n", os.Args[1])
		f, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("Cannot parse %v\n", os.Args[1])
			return
		}
		n = int64(f)

	}
	numbers := counting.GenerateNumbers(n)
	fmt.Printf("Generated %v numbers, Time Taken: %s\n", n, time.Since(t))
	getLoadAvg()

	t = time.Now()
	sum := counting.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
	getLoadAvg()

	for maxproc := 1; maxproc < 4; maxproc++ {
		runtime.GOMAXPROCS(maxproc)
		fmt.Printf("\nMax processors: %v\n", maxproc)

		t = time.Now()
		sum = counting.AddConcurrent(numbers, maxproc)
		fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
		getLoadAvg()

	}
}
