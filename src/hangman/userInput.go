package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, _ := reader.ReadString('\n')

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}

func StringToHex(s string) string {
	hex := ""
	for _, c := range s {
		hex += fmt.Sprintf("%02x", c)
	}
	return hex
}
