package main

import (
	"fmt"

	"github.com/liangx8/gl/rutils"
)

func main() {
	for ix := 0; ix < 10; ix++ {
		fmt.Println(rutils.RandomString(18))
	}
}
