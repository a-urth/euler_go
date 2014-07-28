package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func resultPrinter(res int) {
	fmt.Printf("Result - %d\n", res)
}

func euler1() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func euler2() int {
	sum, first, second := 0, 1, 2

	for second < 4000000 {
		if second%2 == 0 {
			sum += second
		}
		first, second = second, first+second
	}
	return sum
}

func getPrimes(edge int) []int {
	limit := edge + 1
	notPrimes := make(map[int]bool)
	var primes []int
	for i := 2; i < limit; i++ {
		if _, ok := notPrimes[i]; !ok {
			primes = append(primes, i)
		}
		for n := i * i; n < limit; n += i {
			notPrimes[n] = true
		}
	}
	return primes
}

func euler3() int {
	num := 600851475143
	primes := getPrimes(int(math.Sqrt(float64(num))))
	for i := len(primes) - 1; i > 0; i-- {
		if num%primes[i] == 0 {
			return primes[i]
		}
	}
	return 0
}

func isPalindrom(slice string) bool {
	l := len(slice)
	for c := 0; c < l/2; c++ {
		if slice[c] != slice[l-1-c] {
			return false
		}
	}
	return true
}

func euler4() int {
	maxPalindrom := 0
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			current := i * j
			if isPalindrom(strconv.Itoa(current)[:]) {
				if current > maxPalindrom {
					maxPalindrom = current
				}
			}
		}
	}
	return maxPalindrom
}

func euler5() int {
	lcm := 1
	prevArr, currentArr := make([]int, 19), make([]int, 19)
	for i, _ := range prevArr {
		prevArr[i] = i + 2
	}

	for divider := 2; ; {
		needCurrentDividerAgain := false
		validDivider := false
		amIDone := true

		for i, val := range prevArr {
			if val != 1 {
				amIDone = false
			}
			if val%divider == 0 {
				validDivider = true
				currentArr[i] = val / divider
				if currentArr[i] != 1 {
					needCurrentDividerAgain = true
				}
			} else {
				currentArr[i] = val
			}
		}

		if amIDone {
			break
		}

		if validDivider {
			lcm *= divider
		}

		prevArr = currentArr
		if needCurrentDividerAgain {
			continue
		} else {
			divider += 1
		}
	}

	return lcm
}

func euler6() int {
	sum := 0
	for i := 1; i < 101; i++ {
		for j := i; j < 101; j++ {
			if i != j {
				sum += 2 * i * j
			}
		}
	}
	return sum
}

func euler7() int {
	nth := 10001
	apprxUpperBound := nth * 2 * int(math.Log(float64(nth)))
	primes := getPrimes(apprxUpperBound)
	fmt.Println(len(primes))
	return primes[nth-1]
}

func euler8() int {
	num := "316717653133062491922511967442657474235534919493496983520312774506326239578318016984" +
		"80186947885184385861560789112949495459501737958331952853208805511125406987471585238" +
		"63050715693290963295227443043557668966489504452445231617318564030987111217223831136" +
		"22298934233803081353362766142828064444866452387493035890729629049156044077239071381" +
		"05158593079608667017242712188399879790879227492190169972088809377665727333001053367" +
		"88122023542180975125454059475224352584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482839722413756570560574902614079729" +
		"68652414535100474821663704844031998900088952434506585412275886668811642717147992444" +
		"29282308634656748139191231628245861786645835912456652947654568284891288314260769004" +
		"22421902267105562632111110937054421750694165896040807198403850962455444362981230987" +
		"87992724428490918884580156166097919133875499200524063689912560717606058861164671094" +
		"05077541002256983155200055935729725716362695618826704282524836008232575304207529634" +
		"50"
	max := 0
	step := 13
	for i := 0; i <= len(num)-step; i += 1 {
		slice := num[i : i+step]
		if strings.Contains(slice, "0") {
			continue
		}
		multiple := 1
		for _, s := range slice {
			num, _ := strconv.Atoi(string(s))
			multiple *= num
		}
		if multiple > max {
			max = multiple
		}
	}
	return max
}

func euler9() int {
	for n := 1; ; n++ {
		for m := n; m < 100; m++ {
			a, b, c := m*m-n*n, 2*m*n, m*m+n*n
			if a+b+c == 1000 {
				return a * b * c
			}
		}
	}
	return 0
}

func euler10() int {
	primes := getPrimes(2000000)
	sum := 0
	for _, prime := range primes {
		sum += prime
	}
	return sum
}

func checkMult(a int, b int, c int, d int, currentMax int) int {
	t := a * b * c * d
	if t > currentMax {
		return t
	} else {
		return currentMax
	}
}

func euler11() int {
	square := [][]int{
		[]int{8, 02, 22, 97, 38, 15, 00, 40, 00, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8},
		[]int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 00},
		[]int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
		[]int{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91},
		[]int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		[]int{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		[]int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		[]int{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		[]int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		[]int{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
		[]int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92},
		[]int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
		[]int{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		[]int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40},
		[]int{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		[]int{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		[]int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		[]int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16},
		[]int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54},
		[]int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48},
	}
	size := len(square)
	tokenLen := 4
	max := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			canRight := j <= size-tokenLen
			canLeft := j >= tokenLen-1
			canDown := i <= size-tokenLen
			if canRight {
				max = checkMult(
					square[i][j],
					square[i][j+1],
					square[i][j+2],
					square[i][j+3],
					max,
				)
			}
			if canDown {
				max = checkMult(
					square[i][j],
					square[i+1][j],
					square[i+2][j],
					square[i+3][j],
					max,
				)
			}
			if canRight && canDown {
				max = checkMult(
					square[i][j],
					square[i+1][j+1],
					square[i+2][j+2],
					square[i+3][j+3],
					max,
				)
			}
			if canLeft && canDown {
				max = checkMult(
					square[i][j],
					square[i+1][j-1],
					square[i+2][j-2],
					square[i+3][j-3],
					max,
				)
			}
		}
	}
	return max
}

func getDivisors(num int) []int {
	factors := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func euler12() int {
	theNum := 1
	for i := 2; ; i++ {
		theNum += i
		divisors := len(getDivisors(theNum))
		fmt.Println(divisors)
		if divisors == 500 {
			return theNum
		}

	}
	return 0
}

func main() {
	var taskno = flag.Int("taskno", 1, "select euler task number")
	flag.Parse()

	switch *taskno {
	case 1:
		resultPrinter(euler1())
	case 2:
		resultPrinter(euler2())
	case 3:
		resultPrinter(euler3())
	case 4:
		resultPrinter(euler4())
	case 5:
		resultPrinter(euler5())
	case 6:
		resultPrinter(euler6())
	case 7:
		resultPrinter(euler7())
	case 8:
		resultPrinter(euler8())
	case 9:
		resultPrinter(euler9())
	case 10:
		resultPrinter(euler10())
	case 11:
		resultPrinter(euler11())
	case 12:
		resultPrinter(euler12())
	}
}
