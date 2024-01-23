package main

// Factored import
import (
	"fmt"
	"math"
	"math/cmplx"
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

	basicTypesOfGoLang()
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

func basicTypesOfGoLang() {
	// var golang bool

	// var name string

	// var number1 int
	// var number2 int8
	// var number3 int16
	// var number4 int32
	// var number5 int64

	// var number6 uint
	// var number7 uint8
	// var number8 uint16
	// var number9 uint32
	// var number10 uint64
	// var number11 uintptr

	// var byteValue byte // alias for int8

	// var runeValue rune // alias for int32

	// var floatValue1 float32
	// var floatValue2 float64

	// var complexValue1 complex64
	// var complexValue2 complex128

	// The int, uint, and uintptr types
	// are usually 32 bits wide on 32-bit systems
	// and 64 bits wide on 64-bit systems

	var (
		isGoLang      bool       = true
		maxIntValue   uint64     = 1<<64 - 1
		floatValue    float64    = 1.23
		complexNumber complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", isGoLang, isGoLang)
	fmt.Printf("Type: %T Value: %v\n", maxIntValue, maxIntValue)
	fmt.Printf("Type: %T Value: %v\n", complexNumber, complexNumber)
	fmt.Printf("Type: %T Value: %v\n", floatValue, floatValue)
}
