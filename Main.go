package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 54
	var b int = 687
	var ab int = (a * b)
	var result float64 = (math.Pow(float64(ab), 0.5))
	fmt.Println(result)

}
