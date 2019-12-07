package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func compute(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func computeRec(mass float64) float64 {
	fuel := math.Floor(mass/3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + computeRec(fuel)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0.0
	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += computeRec(mass)
	}

	fmt.Println(sum)
}
