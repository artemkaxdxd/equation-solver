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
			fmt.Println("Error. Failed to read input. Please try again", err)
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

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occured while trying to open file.", err)
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	var line string
	for scanner.Scan() {
        line = scanner.Text()
    }
	if err := scanner.Err(); err != nil {
        fmt.Println("Error occured while trying to scan file.", err)
		os.Exit(1)
    }
	
	line = strings.Trim(line, " ")
	paramSplit := strings.Split(line, " ")

	if len(paramSplit) != 3 { 
		fmt.Printf("Error. Wrong number of parameters. Must be 3\n")
		os.Exit(1)
	}

	for i, val := range paramSplit {
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error code = %s\n", err)
			os.Exit(1)
		}

		if valFloat == 0 && i == 0 {
			fmt.Fprint(os.Stderr, "Error. a cannot be 0\n")
			os.Exit(1)
		}
		parameters = append(parameters, valFloat)
	}
	fmt.Println(paramSplit)

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
	fmt.Printf("Equation is: (%.2f) x^2 + (%.2f) x + (%.2f) = 0\n", 
		parameters[0], parameters[1], parameters[2])

	if len(results) == 1 {
		fmt.Printf("x1 = %f\n", results[0])
	} else if len(results) == 2 {
		fmt.Printf("x1 = %f\nx2 = %f\n", results[0], results[1])
	}
	fmt.Printf("There are %d root(s)", len(results))
}