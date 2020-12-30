package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://http.cat/404")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(responseData))
	fmt.Println(response.Header.Get("Content-Type"))

	ioutil.WriteFile("404.jpg", responseData, 0444)
	//fmt.Println(string(responseData))

	//	var responseObject http.Response
	//	json.Unmarshal(responseData, &responseObject)

	//	fmt.Println(responseObject.Name)
	//	fmt.Println(len(responseObject.Cat))
}
