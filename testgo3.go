package testgo3

import (
	"fmt"

	"github.com/ningzero/testgo4"
)

func CallTestgo3() {
	fmt.Println("CallTestgo3 v1.1.1")
	testgo4.CallTestgo4()
}
