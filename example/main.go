package example

import (
	"fmt"
	"time"

	"github.com/zyfdegh/boomer"
)

func main() {
	boomer := boomer.NewBoomer(30, f)
	boomer.Arm()

	time.Sleep(time.Second * 5)
	boomer.Rewind()

	boomer.Unarm()

	time.Sleep(time.Second * 60)
}

func f() {
	fmt.Println("Boom!")
	return
}
