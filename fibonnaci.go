package main
import (
	"math/big"
)


//Fibonacci Numbers calculation function
func FibonacciNumbers(n int) []string {
	if n == 1 { //if only the first value is request just return 1
		return []string{"1"}
	}

	fibonacciSeq := make([]*big.Int, n) //create a *bigInt array to hold all the fibonacci numbers
	var i int

	fibonacciSeq[0] = big.NewInt(0) //set the first value in seq
	fibonacciSeq[1] = big.NewInt(1) //set the second value in seq

	for i = 2; i < n; i++ { //Calculate the remaining values in the sequence based on the first two values
		fibonacciSeq[i] = big.NewInt(0)                           //Initialize the seq[i] value
		fibonacciSeq[i].Add(fibonacciSeq[i-1], fibonacciSeq[i-2]) //Add the last two values in the sequence to produce the current value
	}

	result := make([]string, n) //create an array for the result
	for i = 0; i < n; i++ {
		result[i] = fibonacciSeq[i].Text(10) //base 10 convert the *big.NewInt values and further convert them to a string
	}
	return result

}