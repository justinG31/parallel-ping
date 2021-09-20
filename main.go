package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// create variable to wait for go-routines to finish
var waitToFinish sync.WaitGroup

// data structure to capture ping data from go routines
type pingData struct {
	url  string
	time time.Duration
}

func main() {

	// specifies number of processors to use
	runtime.GOMAXPROCS(4)

	// calls function to get user input
	var urlA, urlB, urlC = askInput()
	totalTime1 := time.Now()
	// specify how many routines needed to be waited for
	waitToFinish.Add(9)
	// make channel and start go routines
	c := make(chan pingData, 9)
	//make three routines for each website
	for i := 0; i < 3; i++ {
		go singlePing(urlA, c)
	}
	for j := 0; j < 3; j++ {

		go singlePing(urlB, c)
	}
	for k := 0; k < 3; k++ {
		go singlePing(urlC, c)
	}

	//let the go-routines finish running before the main function stops running
	waitToFinish.Wait()
	// receive data from channel and print out the data
	ping1, ping2, ping3, ping4, ping5, ping6, ping7, ping8, ping9 := <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c

	fmt.Println(ping1.url, "is: ", ping1.time)
	fmt.Println(ping2.url, "is: ", ping2.time)
	fmt.Println(ping3.url, "is: ", ping3.time)
	fmt.Println(ping4.url, "is: ", ping4.time)
	fmt.Println(ping5.url, "is: ", ping5.time)
	fmt.Println(ping6.url, "is: ", ping6.time)
	fmt.Println(ping7.url, "is: ", ping7.time)
	fmt.Println(ping8.url, "is: ", ping8.time)
	fmt.Println(ping9.url, "is: ", ping9.time)

	totalTime2 := time.Now()
	totalTime := totalTime2.Sub(totalTime1)
	fmt.Println("total main time:", totalTime)
}

// asks for user input of web addresses
func askInput() (string, string, string) {
	fmt.Println("We will ping 3 websites for you to compare.")
	fmt.Println("Enter the first web address that you want to ping:")
	var url1 string
	fmt.Scanln(&url1)

	fmt.Println("Enter the second web address that you want to ping:")
	var url2 string
	fmt.Scanln(&url2)

	fmt.Println("Enter the third web address that you want to ping:")
	var url3 string
	fmt.Scanln(&url3)

	fmt.Println("Thanks! We will try", 100, "pings to", url1, ",", url2, ", and", url3)
	return url1, url2, url3
}

func singlePing(url string, c chan pingData) {
	time1 := time.Now()
	// creating command using input of number of pings and web address
	cmd := exec.Command("ping", "-c 100", url)

	// prints if error occurs
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error! %s\n", err.Error()))
	}

	// captures time at end of routine
	time2 := time.Now()
	timeDiff := time2.Sub(time1)

	// sends data through channel
	c <- pingData{url, timeDiff}

	// utilize the sync variable to know when the function is done running
	defer waitToFinish.Done()
}
