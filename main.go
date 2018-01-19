package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func handler(response http.ResponseWriter, request *http.Request) {
	responseString := ""
	if body, err  := ioutil.ReadAll(request.Body); err != nil {
		responseString = err.Error()
		http.Error(response, responseString, http.StatusInternalServerError)
	} else {
		responseString = string(body)
		response.Write(body)
	}
	log.Println(responseString)
}

func main() {

	log.Println("Starting echo service...")
	if err := http.ListenAndServe(":5001", http.HandlerFunc(handler));err != nil {
		log.Fatalln(err)
	}
}