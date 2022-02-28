package main

import (
	"fmt"
	"time"
)

func main() {
	defer TimeTrack(time.Now(), "Polyomino Count")
	// Create YxY matrix
	// Start from the top, fill 1st row count as 1, remember matrix (Maybe unique sum, not the matrix itself)
	// Element at 0,0 must be one. Removes translation
	// Matrix is gapless (use magic)
	// Transpose every matrix and compare duplicates

}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s : %s\n", name, elapsed)
}
