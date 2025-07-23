package main

import (
	"io"
	"os"
	"testing"
)

func TestIsPrime(t *testing.T) {
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

func TestPrompt(t *testing.T) {
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

}
