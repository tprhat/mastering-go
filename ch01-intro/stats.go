package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func stats() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need one more argument!")
		return
	}
	var min, max float64
	var initialized = 0

	nValues := 0
	var sum float64
	// start from 1 because args[0] is file name
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		nValues += 1
		sum += n

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

	}
	fmt.Println("Number of values: ", nValues)
	fmt.Println("Min value: ", min)
	fmt.Println("Max value: ", max)

	if nValues == 0 {
		return
	}

	meanValue := sum / float64(nValues)
	fmt.Println("Mean: ", meanValue)

	//std dev
	var squared float64
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		squared += math.Pow((n - meanValue), 2)

	}
	stddev := math.Sqrt(squared / float64(nValues))
	fmt.Println("Standard deviation: ", stddev)
}
