package wasm

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Madd(a int, b int) int {
	return a + b
}

func Mtext(i int) string {

	return fmt.Sprintf("This is a string: %d\n", i)
}

// MakeRequest from http://polyglot.ninja/golang-making-http-requests/
func MakeRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	return string(body)
}
