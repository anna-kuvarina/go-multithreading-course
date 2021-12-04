package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

var windDist [8]int

func main() {
	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	startTime := time.Now()

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(data)

		// Change to array, each metar report is a separate item in the array
		metarReports := parseToArray(text)
		windsDirerions := extractWindDirection(metarReports)
		// Assign to N. NE, E, SE, S, SW, W, NW
		mineWindDistribution(windsDirerions)
	}
	elapsed := time.Since(startTime)
	fmt.Printf("%v\n", windDist)
	fmt.Printf("Processing took %s", elapsed)
}
