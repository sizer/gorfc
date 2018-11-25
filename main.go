// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("main started.")
	fetchRfcDetail("8493")
	fmt.Println("main finished.")
}

func fetchRfcDetail(rfcNo string) {
	res, err := http.Get("https://www.rfc-editor.org/rfc/rfc" + rfcNo + ".txt")
	if err != nil {
		fmt.Println("find error." + err.Error())
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", b)
}
