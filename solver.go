package main

import (
	"fmt"
	"math"
	"os"
	"bufio"
	"strconv"
	"strings"	
)

func inMode () []float64 {
	var parameters []float64
	paramNames := []string{"a", "b", "c"}

	reader := bufio.NewReader(os.Stdin)
	for len(parameters) != 3 {
		fmt.Printf("%s: ", paramNames[len(parameters)])
		param, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		param = strings.TrimSuffix(param, "\r\n")
		paramFloat, err := strconv.ParseFloat(param, 64)
		if err != nil {
			fmt.Printf("Error. Expected a valid real number, got %s instead\n", param)
			continue
		}

		if paramFloat == 0 && len(parameters) == 0 {
			fmt.Print("Error. a cannot be 0\n")
			continue
		}
		parameters = append(parameters, paramFloat)
	}

	return parameters
}

func nonInMode (fileName string) []float64 {
	var parameters []float64
	return parameters
}

func equation (parameters []float64) []float64 {
	var roots []float64

	a := parameters[0]
	b := parameters[1]
	c := parameters[2]

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
	var results []float64
	var parameters []float64
	var fileFrom string
	
	if len(os.Args) > 1 {
		fileFrom = os.Args[1]
	}

	if fileFrom != "" {
		parameters = nonInMode(fileFrom)
	} else {
		parameters = inMode()
	}
	
	results = equation(parameters)

	if len(results) == 0 {
		fmt.Printf("There are 0 roots");
	} else if len(results) == 1 {
		fmt.Printf("x1 = %f", results[0]);
	} else {
		fmt.Printf("x1 = %f\nx2 = %f", results[0], results[1]);
	}
}