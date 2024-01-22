package main

// Factored import
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var c, python, java bool
var i, j int = 1, 2

var varKeywordRequired = true

func main() {
	fmt.Println("UTF-8 => Hello utf-8 test: İiŞşçCÖoÜuü")
	fmt.Println("Time:", time.Now())
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Printf("Formatted strings => Calculate: %g\n", math.Pow(10, 2))
	fmt.Printf("Exported names => PI number: %g\n", math.Pi)

	fmt.Println("Functions:", multiply(10, 10))
	fmt.Println("Functions same type parameters:", power(10, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y, total := sum(10, 10)
	fmt.Printf("x: %d, y: %d, sum: %d\n", x, y, total)

	fmt.Println(calculateSomeNumbers(100))

	var i int
	fmt.Println(i, c, python, java)

	initVariables()
	
	fmt.Println(varKeywordRequired)
	varKeywordNotRequired()
}

func multiply(x int, y int) int {
	return x * y
}

func power(x, y float64) float64 {
	return math.Pow(x, y)
}

func swap(x, y string) (string, string) {
	return y, x
}

func sum(x, y int) (int, int, int) {
	total := x + y
	return x, y, total
}

// Use only short functions
func calculateSomeNumbers(x int) (a, b int) {
	a = x + 10
	b = x - 10
	return
}

func initVariables() {
	fmt.Println(i, j)

	var c, python, java = false, true, "java"
	fmt.Println(c, python, java)
}

func varKeywordNotRequired() {
	a := 10
	c, python, java := true, false, "java"
	fmt.Println(a, c, python, java)
}
