package main

import (
	"testing"
)

func TestFibonacciValues (t *testing.T)  {

	testResults := FibonacciNumbers(1001) //Calculate the fibonacci sequence

	hundredth := "218922995834555169026"
	validResultMiddle := "86168291600238450732788312165664788095941068326060883324529903470149056115823592713458328176574447204501"
	validResultEnd := "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"



	if testResults[0] != "0" { //Check the first value is valid
		t.Error("The first value in the Fibonnaci Sequence should be 0")
	}

	if testResults[99] != hundredth {
		t.Error("Incorrect value for 100:th Fibonnaci Number")
	}

	if testResults[499] != validResultMiddle{ //check the last value in the sequence is valid
		t.Error("Incorrect value for 500:th Fibonnaci Number")
	}


	if testResults[1000] != validResultEnd{ //check the last value in the sequence is valid
		t.Error("Incorrect value for 1001:th Fibonnaci Number")
	}
}




