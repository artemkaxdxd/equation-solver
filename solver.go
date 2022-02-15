package main

import (
	"fmt"
	"math"
	// "os"
	// "strconv"
	// "strings"
	// "bufio"
)

func equation (a, b, c float64) []float64 {
	var roots []float64
	discriminant := b*b - 4*a*c
	
	if discriminant == 0 {
		root := -b / (2 * a)
		roots = append(roots, root)
	} else if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2*a)
		root2 := (-b - math.Sqrt(discriminant)) / (2*a)
		roots = append(roots, root1, root2)
	}
	return roots
}

func main () {
	var a, b, c float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)
	
	var results []float64

	results = append(results, equation(a, b, c)[0], equation(a, b, c)[1])
	fmt.Printf("root 1 = %f\nroot 2 = %f", results[0], results[1]);
}