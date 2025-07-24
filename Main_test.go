package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Alpha_IsPrime(t *testing.T) {
	prime_test := []struct {
		Name     string
		TestNum  int
		Expected bool
		Msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 4, false, "4 is not a prime number, it is divisible by 2!"},
		{"definition 0", 0, false, "0 is not prime, by definition!"},
		{"definition 1", 1, false, "1 is not prime, by definition!"},
		{"negative", -1, false, "-1 is a negative number. Therefore, it is not prime!"},
	}

	for _, test_case := range prime_test {
		result, msg := IsPrime(test_case.TestNum)

		if result != test_case.Expected {
			t.Errorf("%s: expected %v, but got %v", test_case.Name, test_case.Expected, result)
		}

		if msg != test_case.Msg {
			t.Errorf("%s: expected message %v, but got %v", test_case.Name, test_case.Msg, msg)
		}
	}

}

func Test_Alpha_Prompt(t *testing.T) {
	// This is a demonstration for how to test output printed to the terminal
	old_out := os.Stdout
	read_pipe, write_pipe, _ := os.Pipe()
	os.Stdout = write_pipe

	Prompt()

	write_pipe.Close()
	os.Stdout = old_out

	printed, _ := io.ReadAll(read_pipe)
	if string(printed) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(printed))
	}
}

func TestIntro(t *testing.T) {
	old_out := os.Stdout
	read_pipe, write_pipe, _ := os.Pipe()
	os.Stdout = write_pipe

	Intro()

	write_pipe.Close()
	os.Stdout = old_out

	printed, _ := io.ReadAll(read_pipe)

	if !strings.Contains(string(printed), "Enter a whole number") {
		t.Errorf("intro text not correct, got %s but expected it contain 'Enter a whole number'", string(printed))
	}

}

func TestCheckNumbers(t *testing.T) {
	// This is a demonstration for how to test user input
	tests := []struct {
		Name        string
		Input       string
		ExpectedMsg string
	}{
		{"Quit", "q", ""},
		{"Empty", "", "Please enter a whole number"},
		{"Word", "Seven", "Please enter a whole number"},
		{"Decimal", "1.7", "Please enter a whole number"},
		{"Negative", "-1", "-1 is a negative number. Therefore, it is not prime!"},
		{"Definition", "0", "0 is not prime, by definition!"},
		{"NotPrime", "4", "4 is not a prime number, it is divisible by 2!"},
		{"Prime", "7", "7 is a prime number!"},
	}

	for _, test_case := range tests {
		input := strings.NewReader(test_case.Input)
		reader := bufio.NewScanner(input)
		res, _ := CheckNumbers(reader)

		if !strings.EqualFold(test_case.ExpectedMsg, res) {
			t.Errorf("%s failed: expected %v but got %v", test_case.Name, test_case.ExpectedMsg, res)
		}
	}
}

func TestReadUserInput(t *testing.T) {
	done_chan := make(chan bool)
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go ReadUserInput(&stdin, done_chan)
	<-done_chan
	close(done_chan)
}
