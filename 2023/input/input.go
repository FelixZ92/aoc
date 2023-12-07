package input

import (
	"bufio"
	"os"
	"strings"
)

func FromStdin() []string {
	scanner := bufio.NewScanner((os.Stdin))
	input := make([]string, 0)
	for {
		scanner.Scan()
		txt := scanner.Text()
		if len(txt) != 0 {
			input = append(input, txt)
		} else {
			break
		}
	}
	return input
}

func FromString(s string) []string {
	return strings.Split(s, "\n")
}

func FromFile(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return input
}
