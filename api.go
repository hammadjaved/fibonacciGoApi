package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"strconv"
	"encoding/json"
)

type fibonacciResponse struct{
	Count int
	Numbers []string
}




func FibonacciApi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	count , errorIntConversion:= strconv.Atoi(ps.ByName("count")) //Parse the value passed in as a parameter through the call
	if errorIntConversion != nil || count <= 0 { //check if value entered is valid
		fmt.Fprint(w,"Invalid count value entered. Please enter a valid integer as the count value between: 1 to 1001.")
		return
		}

	fib :=&fibonacciResponse{ //create the response
		Count:count,
		Numbers: FibonacciNumbers(count)}

	fibJson , err := json.Marshal(fib) //Convert the response to Json
	if err != nil { //If conversion to Json fails give an error
		panic(err)
		return
	}
	w.Header().Set("Content-Type","application/json") //set the response body header to
	w.Write(fibJson) //add the converted Json
}



func main() {
	router := httprouter.New()
	router.GET("/fibonacci/:count", FibonacciApi)
	log.Fatal(http.ListenAndServe(":8080", router))
}