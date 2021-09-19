package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// data structure to capture ping data from go routines
type pingData struct {
	url  string
	time time.Duration
}

// used https://stackoverflow.com/questions/42213996/trying-to-parse-a-stdout-on-command-line-with-golang as a reference
func main() {
	// calls function to get user input
	var urlA, urlB, urlC = askInput()

	// make channel and start go routines
	c := make(chan pingData)
	fmt.Println("Pinging", urlA)
	go singlePing(urlA, c)
	fmt.Println("Pinging", urlB)
	go singlePing(urlB, c)
	fmt.Println("Pinging", urlC)
	go singlePing(urlC, c)

	// receive data from channel
	ping1, ping2, ping3 := <-c, <-c, <-c

	// display data from go routines
	fmt.Println(ping1.url, "is:", ping1.time)
	fmt.Println(ping2.url, "is:", ping2.time)
	fmt.Println(ping3.url, "is:", ping3.time)

	var input string
	fmt.Scanln(&input)
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

// function for go routines to ping a single web address 100 times
func singlePing(url string, c chan pingData) {
	// captures time at start of routine
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
}
