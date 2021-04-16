package lexer

import (
	"fmt"
	"strings"

	"log"
)

var Keywords = []string{"YANCARLOS", "PERNILLO", "CASTILLO", "09051619260"}
var Signs = []string{"\"", "'", "(", ")"}
var Operators = []string{"+", "-", "*", "/", "%", "="}
var Numbers = "1234567890"

//LexElements := map[string]string{"uno":"1"}

func LexElements(s string) {
	elements := map[string]string{
		"=":  "EQUALS",
		"+":  "PLUS",
		"-":  "MINUS",
		"*":  "TIMES",
		"/":  "DIVIDE",
		"%":  "MOD",
		"\"": "ATOM",
		"'":  "APOSTROPHE",
		"(":  "LPAREN",
		")":  "RPAREN",
	}

	appearances := make(map[string]int, len(s))
	kw := foundKeywords(s)

	for _, v := range s {
		character := elements[string(v)]
		if containsNumbers(string(v)) {
			appearances["NUM"] = appearances["NUM"] + 1
		} else {
			appearances[character] = appearances[character] + 1
		}
	}

	fmt.Println(s)
	fmt.Println(kw)
	fmt.Println(appearances)

	checkErrors(kw, appearances)
	print("\n\n")
}

func containsNumbers(s string) bool {
	for _, v := range s {
		if !strings.Contains(Numbers, string(v)) {
			return false
		}
	}

	return true
}

func foundKeywords(s string) map[string]int {
	words := strings.Split(s, " ")
	appearances := make(map[string]int, len(words))
	for _, word := range words {
		for _, v := range Keywords {
			if v == word {
				appearances[v] = appearances[v] + 1
			}
		}
	}

	return appearances
}

func checkErrors(x, y map[string]int) {
	keys := strings.Join(Keywords, " ")
	if len(x) == 0 {
		log.Println("No Keywords")
	}
	for k, v := range x {
		if !strings.Contains(keys, k) {
			log.Println("Error ", k, " Value", v)
		}
	}

	keys = strings.Join(Signs, " ") + strings.Join(Operators, " ") + " NUM"
	for k, v := range y {
		if !strings.Contains(keys, k) {
			log.Println("Elementos que son parte del lenguajes ", k, v)
		}
	}
}
