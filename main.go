package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func euler_1() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Printf("Result - %d", sum)
}

func euler_2() {
	sum, first, second := 0, 1, 2

	for second < 4000000 {
		if second%2 == 0 {
			sum += second
		}
		first, second = second, first+second
	}
	fmt.Printf("Result - %d", sum)
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

func euler_3() {
	num := 600851475143
	primes := getPrimes(int(math.Sqrt(float64(num))))
	fmt.Println(len(primes))
	for i := len(primes) - 1; i > 0; i-- {
		if num%primes[i] == 0 {
			fmt.Printf("Result - %d", primes[i])
			break
		}
	}
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

func euler_4() {
	max := 0
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			current := i * j
			if isPalindrom(strconv.Itoa(current)[:]) {
				if current > max {
					max = current
				}
			}
		}
	}
	fmt.Printf("Result - %d", max)
}

func main() {
	var taskno = flag.Int("taskno", 1, "select euler task number")
	flag.Parse()

	switch *taskno {
	case 1:
		euler_1()
	case 2:
		euler_2()
	case 3:
		euler_3()
	case 4:
		euler_4()
	}
}
