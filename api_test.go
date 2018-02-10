package main

import (
	"testing"
	"time"
	"net/http"
	"encoding/json"
	"log"
	"bytes"
	"io/ioutil"
	"strconv"
)


func callApiWithCount(count int, client http.Client) []string {

	var buffer bytes.Buffer //create the get request url
	buffer.WriteString("http://localhost:8080/fibonacci/")
	buffer.WriteString(strconv.Itoa(count))
	url :=  buffer.String()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil { // if get request creation fails then stop the test
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "fibonacciapi-test")

	res, getError := client.Do(req) //Execute request
	if getError != nil { // if get request fails then stop the test
		log.Fatal(getError)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil { // if get reading response body fails stop the test
		log.Fatal(readErr)
	}

	response := fibonacciResponse{}
	jsonError := json.Unmarshal(body, &response)
	if jsonError != nil { // if json conversion of response body fails stop the test
		log.Fatal(jsonError)
	}
	return response.Numbers
}

/*A count value of 11000 gave an average response time of 500ms.
Disclaimer: This value will differ depending on where and how the webservice is running so a just to be on the safe side a max limit of 10000 was chosen for this api.
Also if the service is called through a web browser the, it takes the browser to interpret the data being passed to it as a response which significantly reduces load times for the data. Hence, a max
value of 1001 if the service is being called through the browser.
 */
func TestApiSpeed (t *testing.T) {
	client := http.Client{
		Timeout: time.Second * 2,
	} //Initialize the client

	go main() //run the service asynchronously

	for i := 10990; i < 11010 ; i++{
		start:= time.Now()
		callApiWithCount(i,client)
		elapsed := time.Since(start)
		t.Logf("%v took %s\n", i, elapsed)
	}
}

func TestApiResponse (t *testing.T){
	client := http.Client{
		Timeout: time.Second * 2,
	} //Initialize the client
	go main()
	testResults := callApiWithCount(1001,client)

	validResultMiddle := "86168291600238450732788312165664788095941068326060883324529903470149056115823592713458328176574447204501"
	validResultEnd := "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"

	if testResults[0] != "0" { //Check the first value is valid
		t.Error("The first value in the Fibonnaci Sequence should be 0")
	}

	if testResults[499] != validResultMiddle{ //check the last value in the sequence is valid
		t.Error("Incorrect value for 500:th Fibonnaci Number")
	}


	if testResults[1000] != validResultEnd{ //check the last value in the sequence is valid
		t.Error("Incorrect value for 1001:th Fibonnaci Number")
	}

}
