package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://http.cat")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	//	var responseObject http.Response
	//	json.Unmarshal(responseData, &responseObject)

	//	fmt.Println(responseObject.Name)
	//	fmt.Println(len(responseObject.Cat))
}
