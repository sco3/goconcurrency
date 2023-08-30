package main

import (
	"cnt/counting"
	"fmt"
	"log"
	"os/exec"
	"runtime"
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
	var n int = 5e8
	numbers := counting.GenerateNumbers(n)
	fmt.Printf("Generated %v numbers, Time Taken: %s\n", n, time.Since(t))
	getLoadAvg()

	t = time.Now()
	sum := counting.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
	getLoadAvg()

	for maxproc := 1; maxproc <= 4; maxproc++ {
		runtime.GOMAXPROCS(maxproc)
		fmt.Printf("\nMax processors: %v\n", maxproc)

		t = time.Now()
		sum = counting.AddConcurrent(numbers)
		fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
		getLoadAvg()

	}
}
