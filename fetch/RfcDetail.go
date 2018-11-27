package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// RfcDetail will https://www.rfc-editor.orgから取得orCacheから取得する
func RfcDetail(rfcNo string) {
	var cacheName = getCacheName(rfcNo)

	if checkIfCached(cacheName) {
		readCache(cacheName)
	} else {
		res, _ := httpGetRfcDetail(rfcNo)
		b, err := ioutil.ReadAll(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", b)

		makeCache(cacheName, b)
	}
}

func httpGetRfcDetail(rfcNo string) (io.ReadCloser, error) {
	res, err := http.Get("https://www.rfc-editor.org/rfc/rfc" + rfcNo + ".txt")
	if err != nil {
		fmt.Println("find error." + err.Error())
		return nil, err
	}

	return res.Body, nil
}

func getCacheName(rfcNo string) string {
	return "RFC" + rfcNo + ".txt"
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
