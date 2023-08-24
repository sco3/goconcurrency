package main

import (
	"fmt"
	"time"

	"cnt/counting"
)

func main() {
	numbers := counting.GenerateNumbers(1e8)

	t := time.Now()
	sum := counting.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))

	t = time.Now()
	sum = counting.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
}
