package main

import (
	"fmt"
	"html"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {

		log.Fatal("Woah there, nelson. You need to input a file")
	}
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	_, err = parseJson(string(contents))
	if err != nil {
		log.Fatal(err)
	}

}

func parseJson(contents string) (any, error) {

	tokens, err := lex(contents)
	if err != nil {
		return "", err
	}

	json, err := parseTokens(tokens)
	if err != nil {

		return "", err
	}

	hundred := "\\U0001F4AF"
	hex := strings.ReplaceAll(hundred, "\\U", "0x")
	i, _ := strconv.ParseInt(hex, 0, 64)
	str := html.UnescapeString(string(i))

	fmt.Println(`Ayo this is valid `, str)

	return json, nil
}
