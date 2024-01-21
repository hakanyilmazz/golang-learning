package main

// Factored import
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("UTF-8 => Hello utf-8 test: İiŞşçCÖoÜuü")
	fmt.Println("Time:", time.Now())
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Printf("Formatted strings => Calculate: %g\n", math.Pow(10, 2))
	fmt.Printf("Exported names => PI number: %g\n", math.Pi)

	fmt.Println("Functions:", multiply(10, 10))
	fmt.Println("Functions same type parameters:", power(10, 2))
}

func multiply(x int, y int) int {
	return x * y
}

func power(x, y float64) float64 {
	return math.Pow(x, y)
}
