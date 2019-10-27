package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "min value of cube side")
	end = flag.Int("end", 6, "max value of cube side")
	n = flag.Int("n", 1, "number of random integers to be generated")
	norepeat = flag.Bool("norepeat", false, "if generated numbers must not repeat")
)

func randInterval(l int, r int) int {
	return rand.Intn(r) + l
}


func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	// Dice roll 1..6
	if *start > *end {
		fmt.Print("Произошла ошибка! Значение start не может быть больше end!")
	} else {
		for i := 0; i < *n; i++ {
			number := randInterval(*start, *end + 1)

			if i != *n - 1 {
				fmt.Print(number, ",")
			} else {
				fmt.Print(number)
			}
		}
	}
}
