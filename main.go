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
	}
}
