package main

import (
	"fmt"
	"github.com/sizer/gorfc/fetch"
)

func main() {
	fmt.Println("main started.")
	fetch.RfcDetail("8493")
	fmt.Println("main finished.")
}
