package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// RfcDetail will https://www.rfc-editor.orgから取得orCacheから取得する
func RfcDetail(rfcNo string) {
	var cacheName = "RFC" + rfcNo + ".txt"

	if checkIfCached(cacheName) {
		readCache(cacheName)
	} else {
		httpGetRfcDetail(rfcNo)
	}
}

func httpGetRfcDetail(rfcNo string) {
	res, err := http.Get("https://www.rfc-editor.org/rfc/rfc" + rfcNo + ".txt")
	if err != nil {
		fmt.Println("find error." + err.Error())
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", b)
	
	makeCache("RFC" + rfcNo + ".txt", b)
}

func makeCache(filename string, data []byte) {
	ioutil.WriteFile(getCacheFilePath(filename), data, 0644)	
}

func readCache(cacheName string) {
	b, err := ioutil.ReadFile(getCacheFilePath(cacheName))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", b)
}

func checkIfCached(key string) bool {
	if _, err := os.Stat(getCacheFilePath(key)); os.IsNotExist(err) {
		return false
	}
	return true
}

func getCacheFilePath(filename string) string {
	return "tmp/cache/" + filename
}
