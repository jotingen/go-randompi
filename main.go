package main

import (
    "fmt"
	"math"
    "math/rand"
)

func main() {
	var total int64 = 0
	var coprimes int64 = 0
	
	for ;; {
		a := rand.Int63()
		b := rand.Int63()
		
		total++
		if GCD(a,b) == 1 {
			coprimes++
		}
		
		if total%100000==0 {
			pi := math.Sqrt(6/(float64(coprimes)/float64(total)))
			err := math.Abs(math.Pi-pi)/math.Pi * 100
		
			fmt.Printf("\r%d %19d %19d %40.38f %E%%",total,a,b,pi,err)
		}
	}
}

func GCD(a, b int64) int64 {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}