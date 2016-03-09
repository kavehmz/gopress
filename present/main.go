package main

import (
	"fmt"
	"time"

	"github.com/kavehmz/prime"
)

func main() {
	t := time.Now().UnixNano()

	fmt.Println("number of primes (non-segmented-method):", len(prime.SieveOfEratosthenes(1000000000)))
	fmt.Println("number of primes (segmented-method):", len(prime.Primes(1000000000)))
	fmt.Println("seconds it took:", float64(time.Now().UnixNano()-t)/1000000000)

}
