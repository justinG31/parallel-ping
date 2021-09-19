package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// used https://stackoverflow.com/questions/42213996/trying-to-parse-a-stdout-on-command-line-with-golang as a reference
func main() {
	// calls function to get user input
<<<<<<< HEAD
	var url, count = askInput()
	sCount := strconv.Itoa(count)

	for i := 0; i < 3; i++ {
		go singlePing(sCount, url)
	}

	var input string
	fmt.Scanln(&input)

=======
	var url1, url2, url3 = askInput()

	fmt.Println("Pinging", url1)
	singlePing(url1)
	fmt.Println("Pinging", url2)
	singlePing(url2)
	fmt.Println("Pinging", url3)
	singlePing(url3)
>>>>>>> fd45475f9743bade93ce52c4eabe181ae3151eeb
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

	fmt.Println("runtime is", timeDiff)
}
