package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// used https://stackoverflow.com/questions/42213996/trying-to-parse-a-stdout-on-command-line-with-golang as a reference
func main() {
	// calls function to get user input
	var url, count = askInput()
	sCount := strconv.Itoa(count)

	// creating command using input of number of pings and web address
	cmd := exec.Command("ping", "-c", sCount, url)
	output := &bytes.Buffer{}
	cmd.Stdout = output

	// prints if error occurs
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error! %s\n", err.Error()))
	}

	// prints the output
	if len(output.Bytes()) > 0 {
		fmt.Printf("Output: %s\n", string(output.Bytes()))
	}
}

// asks for user input of web address for ping and number of pings
func askInput() (string, int) {
	fmt.Println("Enter the web address that you want to ping:")
	var url string
	fmt.Scanln(&url)

	fmt.Println("Enter the number of pings would like to do:")
	var count int
	fmt.Scanln(&count)
	fmt.Println("Thanks! We will try", count, "pings to", url)
	return url, count
}
