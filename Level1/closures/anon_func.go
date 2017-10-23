package main

import (
	"fmt"
)

func sq(x float32) float32 {
	return x * x
}

func inv(x float32) float32 {
	return 1. / x
}

type FloatProc func(float32) float32

func wrapFloatProcessor(f FloatProc) FloatProc {
	numCalls := 0

	return func(x float32) float32 {
		fmt.Println("NumCalls", numCalls)
		numCalls++
		return f(x)
	}
}

func main() {
	f := wrapFloatProcessor(inv)
	fmt.Println(f(0.), f(1.), f(2))
}

// func makeSquareFunc() func(float32) float32 {
// 	numCalls := 0

// 	return func(x float32) float32 {
// 		fmt.Println("NumCalls", numCalls)
// 		numCalls++
// 		return x * x

// 	}
// }
