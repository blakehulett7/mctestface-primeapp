package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Intro()

	// Make a channel to listen for the quit signal
	done_chan := make(chan bool)

	go ReadUserInput(done_chan)
	<-done_chan
	close(done_chan)

	fmt.Println("Et Filii, et Spiritus Sancti!")
}

func Intro() {
	fmt.Println("In nomine Patris...")
	fmt.Println("-------------------")
	fmt.Println("Enter a whole number to see if it is prime. Enter q to quit.")
	Prompt()
}

func Prompt() {
	fmt.Print("-> ")
}

func ReadUserInput(done_chan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := CheckNumbers(scanner)

		if done {
			done_chan <- true
			return
		}

		fmt.Println(res)
		Prompt()
	}
}

func CheckNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := IsPrime(num)

	return msg, false
}

func IsPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("%d is a negative number. Therefore, it is not prime!", n)
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number, it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
