package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zyfdegh/boomer"
)

func main() {
	// uncomment this line in `boomer.go` will print seconds
	// fmt.Printf("\r%ds ", p.seconds)

	boomer, err := boomer.NewBoomer(5, f)
	if err != nil {
		log.Printf("new boomer error: %v", err)
		return
	}
	boomer.Arm()

	time.Sleep(time.Second * 3)
	err = boomer.Rewind()
	if err != nil {
		log.Printf("rewind timer error: %v", err)
		return
	}

	//	boomer.Unarm()

	time.Sleep(time.Second * 10)
}

func f() {
	fmt.Println("Boom!")
	return
}
