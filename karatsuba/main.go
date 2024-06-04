package main

import (
	"fmt"
	"math"
    "strconv"
)

func main() {
    fmt.Println(karatsuba(1222, 2022))
    fmt.Println(multiply(1222, 2022))
}

// function perkalian
func multiply(x int64, y int64) string {
    // ubah x dan y menjadi string agar mempermudah pengambilan digit
    a := strconv.FormatInt(x, 10) 
    b := strconv.FormatInt(y, 10) 

    // inisialisasi array 2d untuk menampung hasil perkalian
	columns := make([][]uint8, len(a)+len(b))
	for i := 0; i < len(columns); i++ {
		columns[i] = make([]uint8, len(a)+len(b))
	}

    // loop yang mengalikan digit per digit
	for i := len(b) - 1; i > -1; i-- {
		var tens uint8 = 0

		for j := len(a) - 1; j > -1; j-- {
			n1Int, _ := strconv.Atoi(string(a[j]))
			n2Int, _ := strconv.Atoi(string(b[i]))

			n1 := uint8(n1Int)
			n2 := uint8(n2Int)

			t1 := n1*n2 + tens

			tens = uint8(t1 / 10)

			newDigit := uint8(t1 % 10)

			columns[len(a)-j-1+(len(b)-i-1)][(len(b) - i - 1)] = newDigit
		}

		columns[len(a)+(len(b)-i-1)][len(b)-i-1] = tens
	}

    // inisialisasi result 
	resultStr := ""

	var leftoverFromLastDigit uint = 0
    
    // membangun final string
	for i := 0; i < len(columns); i++ {
		var total uint = 0

		for j := 0; j < len(columns[i]); j++ {

			total += uint(columns[i][j])
		}

		total += leftoverFromLastDigit

		leftoverFromLastDigit = uint(total / 10)

		newDigit := strconv.Itoa(int(total) % 10)

		resultStr = newDigit + resultStr
	}

	return resultStr
}

// mengambil berapa banyak digit yang ada pada sebuah integer
func getDecimalDigits(num int64) int64 {
	var result int64

	if num == 0 {
		return 1
	}

	if num < 0 {
		num = -num
	}
	for num > 0 {
		result++
		num = num / 10
	}

	return result
}

// membagi integer menjadi dua bagian atas dan bawah
func getHighAndLowDigits(num int64, digits int64) (int64, int64) {
	divisor := int64(math.Pow(10, float64(digits)))

	if num >= divisor {
		return num / divisor, num % divisor
	} else {
		return 0, num
	}
}

func karatsuba(x int64, y int64) int64 {
	var max_digits int64
	positive := true
    
    // apabila salah satu dari integer = 0 maka function langsung me-return 0
	if x == 0 || y == 0 {
		return 0
	}

    // apabila salah satu bilangan adalah bilangan negative maka flag positive akan berubah menjadi false
	if (x > 0 && y < 0) || (x < 0 && y > 0) {
		positive = false
	}
    
    // apabila salah satu bilangan lebih kecil dari 10 maka function langsung me-return x * y
	if x < 10 || y < 10 {
		return x * y
	}

	x_digits := getDecimalDigits(x)
	y_digits := getDecimalDigits(y)

	if x_digits >= y_digits {
		max_digits = x_digits / 2
	} else {
		max_digits = y_digits / 2
	}

	x_high, x_low := getHighAndLowDigits(x, max_digits)
	y_high, y_low := getHighAndLowDigits(y, max_digits)
    
    // proses rekursif antara bilangan x bawah dan y bawah
	z0 := karatsuba(x_low, y_low)
    // proses rekursif antara bilangan mid 
	z1 := karatsuba((x_low + x_high), (y_low + y_high))
    // proses rekursif antara bilangan x atas dan y atas
	z2 := karatsuba(x_high, y_high)

	if positive {
		return (z2 * int64(math.Pow(10, float64(2*max_digits)))) + (z1-z2-z0)*int64(math.Pow(10, float64(max_digits))) + z0
	} else {
		return -((z2 * int64(math.Pow(10, float64(2*max_digits)))) + (z1-z2-z0)*int64(math.Pow(10, float64(max_digits))) + z0)
	}

}
