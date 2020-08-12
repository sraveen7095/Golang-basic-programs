package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "Result is: %s",r.URL.Path)
}

func main() {
	//To View the browser on defined port else by default it is 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//Prints within the terminal
	fmt.Println("hello world")


	//Addition

	fmt.Print("Enter First Number: ") //Print function is used to display output in same line
	var first int
	fmt.Scanln(&first) // Take input from user
	fmt.Print("Enter Second Number: ")
	var second int
	fmt.Scanln(&second)

	var res = first + second
	fmt.Println(res) // Addition of two string
	//To view in browser
	http.HandleFunc("/",Index)
	http.ListenAndServe(":"+port,nil)
}
