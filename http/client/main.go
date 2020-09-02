package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://0.0.0.0:8080/api/v1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(body)
	}
}
