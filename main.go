package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

func main() {
	var err error
	digits := 20
	if len(os.Args) == 2 {
		digits, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Not valid integer")
			os.Exit(1)
		}
	}
	fmt.Println(digits)

	var total int64 = 0
	var coprimes int64 = 0

	total_chan := make(chan int64, runtime.NumCPU())
	coprimes_chan := make(chan int64, runtime.NumCPU())

	for thr := 0; thr < runtime.NumCPU(); thr++ {
		go CheckInts(digits, total_chan, coprimes_chan)
	}

	for {

		total += <-total_chan
		coprimes += <-coprimes_chan
		if total%100 == 0 {
			pi := math.Sqrt(6 / (float64(coprimes) / float64(total)))
			err := math.Abs(math.Pi-pi) / math.Pi * 100

			fmt.Printf("%d %d %40.38f %E%%\r", total, coprimes, pi, err)
		}
	}
}

func CheckInts(digits int, total_chan chan int64, coprimes_chan chan int64) {
	for {

		var abuffer bytes.Buffer
		var bbuffer bytes.Buffer
		for i := 0; i < digits; i++ {
			abuffer.WriteString(strconv.Itoa(rand.Intn(10)))
			bbuffer.WriteString(strconv.Itoa(rand.Intn(10)))
		}
		a := new(big.Int)
		b := new(big.Int)
		a.SetString(abuffer.String(), 10)
		b.SetString(bbuffer.String(), 10)

		total_chan <- 1
		if IsCoprime(a, b) {
			coprimes_chan <- 1
		} else {
			coprimes_chan <- 0
		}
	}
}

func IsCoprime(a, b *big.Int) bool {
	one := big.NewInt(1)
	gcd := new(big.Int)
	gcd.GCD(nil, nil, a, b)

	return gcd.Cmp(one) == 0
}
