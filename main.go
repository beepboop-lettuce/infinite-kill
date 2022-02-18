package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randCharList := "|\\-/"
	randChar := []rune(randCharList)

	for {
		randIndex := rand.Intn(len(randChar))
		pick := randChar[randIndex]

		fmt.Printf("\r%s Please wait. Processing....", string(pick))
		time.Sleep(250 * time.Millisecond)
	}
}
