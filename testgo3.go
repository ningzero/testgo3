package testgo3

import (
	"fmt"

	"github.com/ningzero/testgo4"
)

func CallTestgo3() {
	fmt.Println("CallTestgo4")
	testgo4.CallTestgo2()
}
