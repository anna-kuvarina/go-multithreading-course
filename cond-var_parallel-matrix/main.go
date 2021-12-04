package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	matrixSize = 200
)

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	resultMatrix = [matrixSize][matrixSize]int{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	for col := 0; col < matrixSize; col++ {
		for i := 0; i < matrixSize; i++ {
			resultMatrix[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}

func main() {
	fmt.Println("Working ...")
	start := time.Now()
	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		for row := 0; row < matrixSize; row++ {
			workOutRow(row)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s", elapsed)

}
