package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args[1]
	n, _ := strconv.Atoi(arg)

	start := time.Now()

	// fmt.Println(fibPd(n))
	fibPd(n)

	end := time.Now()
	totalTime := end.Sub(start)

	fmt.Printf("\nIn GO - Fibonnaci series in position %d took: %s\n", n, totalTime)
}

func fibPd(pos int) *big.Int {
	mem := make([]*big.Int, pos+1)
	for i := 0; i <= pos; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(mem[i-1], mem[i-2])
		}
		mem[i] = f
	}
	return mem[pos]
}
