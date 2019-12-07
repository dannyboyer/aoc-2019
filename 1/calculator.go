package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"fmt"
)

func calcFuelReqWithoutMass(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func calcFuelReq(mass float64) float64 {
	fuel := math.Floor(mass/3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calcFuelReq(fuel)
}

func getAnswer2(massList []float64) float64 {
	sum := 0.0
	for _, mass := range massList {
		sum += calcFuelReq(mass)
	}
	return sum
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)	
	var massList []float64
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		massList = append(massList, mass)
	}

	fmt.Println(getAnswer2(massList))	
}
