package main

import (
	"fmt"
	"regexp"
	"unicode"
)

var (
	jsonTrue  = []rune("true")
	jsonFalse = []rune("false")
	jsonNull  = []rune("null")
)

func lex(s string) ([]Token, error) {

	tokens := []Token{}

	for runes := []rune(s); len(runes) > 0; {

		char := runes[0]
		if unicode.IsSpace(char) {

			runes = runes[1:]

			continue
		}

		token := Token{}

		var err error

		//lexString
		token, runes, err = lexString(runes)

		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			continue
		}

		//lexBoolean
		token, runes, err = lexBoolean(runes)
		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			continue
		}

		//lexNull
		token, runes, err = lexNull(runes)
		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			continue
		}

		//lexNumbs
		token, runes, err = lexNumbers(runes)
		if err != nil {
			return []Token{}, err
		} else if token != (Token{}) {
			tokens = append(tokens, token)
			continue
		}

		_, ok := JsonSpecialChars[char]

		if ok {
			tokens = append(tokens, Token{JsonSyntax, string(char)})
			runes = runes[1:]
		} else {
			return tokens, fmt.Errorf("unexpected character %s", string(char))
		}
	}

	return tokens, nil
}

func lexString(runes []rune) (Token, []rune, error) {
	if runes[0] != '"' {
		return Token{}, runes, nil
	}

	runes = runes[1:]
	escaped := false
	for i, char := range runes {

		if escaped {
			switch char {

			case 'b', 'f', 'n', 'r', 't', '\\', '"':
				escaped = false
			default:
				return Token{}, runes, fmt.Errorf("Invalid escape token")
			}

		} else if char == '\\' {
			escaped = true
		} else if char == '"' {
			return Token{JsonString, string(runes[:i])}, runes[i+1:], nil

		}

	}
	return Token{}, runes, fmt.Errorf("Missing end of string quote \"")
}

func lexNumbers(runes []rune) (Token, []rune, error) {

	if !unicode.IsDigit(runes[0]) && runes[0] != '-' {
		return Token{}, runes, nil
	}

	var length = len(runes) - 1
	for i, char := range runes {

		if !unicode.IsDigit(char) && char != 'e' && char != 'E' && char != '.' && char != '-' && char != '+' {
			length = i - 1
			break
		}
	}

	tokenVal := string(runes[:length+1])

	if !regexp.MustCompile(`^-?\d+(?:\.\d+)?(?:[eE][\-+]?\d+)?$`).MatchString(tokenVal) {
		return Token{}, runes, fmt.Errorf("Invalid number %s", tokenVal)
	}

	return Token{JsonNumber, tokenVal}, runes[length+1:], nil
}

func lexBoolean(runes []rune) (Token, []rune, error) {

	if CompareRuneSlices(runes, jsonTrue, len(jsonTrue)) {
		return Token{JsonBoolean, string(runes[:len(jsonTrue)])}, runes[len(jsonTrue):], nil
	}
	if CompareRuneSlices(runes, jsonFalse, len(jsonFalse)) {
		return Token{JsonBoolean, string(runes[:len(jsonFalse)])}, runes[len(jsonFalse):], nil
	}

	return Token{}, runes, nil
}

func lexNull(runes []rune) (Token, []rune, error) {
	if CompareRuneSlices(runes, jsonNull, len(jsonNull)) {
		return Token{JsonNull, string(runes[:len(jsonNull)])}, runes[len(jsonNull):], nil
	}

	return Token{}, runes, nil
}

func CompareRuneSlices(r1 []rune, r2 []rune, n int) bool {

	if n > len(r1) || n > len(r2) {
		return false
	}

	for i := 0; i < n; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
