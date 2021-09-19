package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

<<<<<<< Updated upstream
// used https://stackoverflow.com/questions/42213996/trying-to-parse-a-stdout-on-command-line-with-golang as a reference
func main() {
	// calls function to get user input
	var url1, url2, url3 = askInput()
=======
//create variable to wait for go-routines to finish
var waitToFinish sync.WaitGroup

// data structure to capture ping data from go routines
type pingData struct {
	url  string
	time time.Duration
}

// used https://stackoverflow.com/questions/42213996/trying-to-parse-a-stdout-on-command-line-with-golang as a reference
func main() {
	// calls function to get user input
	var urlA, urlB, urlC = askInput()

	// specifiy how many routines needed to be waited for
	waitToFinish.Add(3)

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
>>>>>>> Stashed changes

	fmt.Println("Pinging", url1)
	go singlePing(url1)
	fmt.Println("Pinging", url2)
	go singlePing(url2)
	fmt.Println("Pinging", url3)
	go singlePing(url3)

	//let the go-routines finish running before the main function stops running
	waitToFinish.Wait()
	fmt.Println("All go-routines have finished")
}

// asks for user input of web address for ping and number of pings
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

func singlePing(url string) {
	time1 := time.Now()
	// creating command using input of number of pings and web address
	cmd := exec.Command("ping", "-c 100", url)

	// prints if error occurs
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error! %s\n", err.Error()))
	}

	time2 := time.Now()
	timeDiff := time2.Sub(time1)

<<<<<<< Updated upstream
	fmt.Println("runtime is", timeDiff)
=======
	// sends data through channel
	c <- pingData{url, timeDiff}

	// utilize the sync variable to know when the function is done running
	defer waitToFinish.Done()
>>>>>>> Stashed changes
}
