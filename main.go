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

func alreadyExists(a []int, b int) bool {
	flag_ := false

	for i := 0; i < len(a); i++ {
		if a[i] == b {
			flag_ = true
			break
		}
	}

	return flag_
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
	} else if *n > *end - *start + 1 {
		fmt.Printf("Невозможно сгененировать %d случайных чисел в интервале [%d; %d]", *n, *start, *end)
	} else {
		if *norepeat {
			a := make([]int, *end-*start+1) // Все сгенерированные числа

			for j := 0; j < len(a); j++ {
				a[j] = -10000
			}

			t := 0 // Указатель на свободную позицию в массиве а

			for i := 0; i < *n; i++ {
				var number int

				number = randInterval(*start, *end+1)

				for ; alreadyExists(a, number) == true; {
					number = randInterval(*start, *end+1)

				}

				a[t] = number
				t += 1
			}

			for j := 0; j < *n; j++ {
				if j != *n-1 {
					fmt.Print(a[j], ",")
				} else {
					fmt.Print(a[j])
				}
			}
		} else {
			for i := 0; i < *n; i++ {
				number := randInterval(*start, *end+1)

				if i != *n - 1 {
					fmt.Print(number, ",")
				} else {
					fmt.Print(number)
				}
			}
		}
	}
}
