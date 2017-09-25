package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var re, _ = regexp.Compile(`\r*\n*\s*[.,*¿?"{}'!¡]*`)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	for input == "" {
		fmt.Print("Texto: ")
		input, _ = reader.ReadString('\n')
		input = re.ReplaceAllString(input, "")
		if input == "" {
			fmt.Println("Texto invalido, intenta de nuevo.")
		}
	}

	if IsPalindrome(input) {
		color.Yellow("Es un palindromo.")
	} else {
		color.Red("No es un palindromo.")
	}
}

//IsPalindrome check if a string a palindrome
func IsPalindrome(input string) bool {
	input = re.ReplaceAllString(input, "")
	input = strings.ToUpper(input)

	var specialCharacters = map[string]string{
		"A": "áà",
		"E": "éè",
		"I": "íì",
		"O": "óò",
		"U": "úù",
	}

	for key, characters := range specialCharacters {
		var pattern = "(?i)[" + characters + "]"
		var re2, _ = regexp.Compile(pattern)
		input = re2.ReplaceAllString(input, key)
	}

	var j = 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != input[j] {
			return false
		}
		j++
	}

	return true
}
