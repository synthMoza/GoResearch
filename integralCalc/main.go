package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

// Input function for integrating on the given length
func function(x float64) float64 {
	return x * math.Sin(x)
}

// Type of function to pass to calculateIntegral function
type integrateFunction func(float64) float64

// Accuracy of integral calculation. Bigger this value is, more accurate result will be
const calcAccuracy = 2e4

// Check is the input integer is even
func isEven(input int) bool {
	return input%2 == 0
}

// Goroutine of integral calculation, calcualtes the integral of the given function on the given section
func calculateIntegralRoutine(f integrateFunction, a, b float64, routineNumber int, routineResults []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var tmpResult float64 = f(a) + f(b)
	iterationNumber := calcAccuracy * int(b-a) // accuracy of calculations

	step := (b - a) / float64(iterationNumber)
	for j := 0; j < iterationNumber; j++ {
		x := a + step*float64(j)

		if isEven(j) {
			tmpResult += 2 * f(x)
		} else {
			tmpResult += 4 * f(x)
		}
	}

	tmpResult *= step / 3
	routineResults[routineNumber] = tmpResult
}

// Launch all routines for calcualteIntegral function and return slice of each one result
func calculateIntegralDoRoutines(f integrateFunction, goroutinesNumber int, a, b float64) []float64 {
	wg := &sync.WaitGroup{}

	// Launch goroutines with corresponding input values
	routineResults := make([]float64, goroutinesNumber)
	step := (b - a) / float64(goroutinesNumber)
	for i := 0; i < goroutinesNumber; i++ {
		startValue := a + step*float64(i)
		endValue := a + step*float64(i+1)

		wg.Add(1)
		go calculateIntegralRoutine(f, startValue, endValue, i, routineResults, wg)
	}

	// Collect all results
	wg.Wait()
	return routineResults
}

// Calculate the integral of the given fucntion on the section [a;b]
func calculateIntegral(f integrateFunction, a, b float64) float64 {
	var result float64 = 0

	// Divide the section by goroutinesNumber sections
	goroutinesNumber := runtime.NumCPU()
	routineResults := calculateIntegralDoRoutines(f, goroutinesNumber, a, b)

	for i := 0; i < goroutinesNumber; i++ {
		result += routineResults[i]
	}

	calculateIntegralDoRoutines(f, goroutinesNumber, a, b)
	return result
}

func main() {
	fmt.Print("Integral calculation using goroutines, enter the length: ")

	// Take the right length to integrate on
	var a, b float64
	for {
		fmt.Scan(&a, &b)
		if a > b {
			fmt.Print("Enter the valid section (a <= b): ")
		} else {
			break
		}
	}

	result := calculateIntegral(function, a, b)
	fmt.Println("Integral of f(x) from", a, "to", b, "equals", result)
}
