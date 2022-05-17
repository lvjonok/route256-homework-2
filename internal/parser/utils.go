package parser

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

const baseURL string = "https://ege.sdamgia.ru"

func isTokenClass(tok *html.Token, className string) bool {
	for _, attr := range tok.Attr {
		if attr.Key == "class" && attr.Val == className {
			return true
		}
	}

	return false
}

func getAttrToken(tok *html.Token, key string) (string, error) {
	for _, attr := range tok.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}

	return "", fmt.Errorf("did not find key %v in token %v", key, tok)
}

func isEndingClassNames(tokens []html.Token, ends []string) (bool, error) {
	if len(tokens) < len(ends) {
		return false, fmt.Errorf("length of tokens is less than ends")
	}

	tl := len(tokens)

	for idx, className := range ends {
		if !isTokenClass(&tokens[tl-idx-1], className) {
			return false, nil
		}
	}

	return true, nil
}

func isLink(el string) bool {
	return strings.HasPrefix(el, "http")
}

// tokensContain returns true if query of class names is present in given order in tokens
func tokensContain(tokens []html.Token, query []string) bool {
	for i := 0; i < len(tokens)-len(query)+1; i++ {
		good := true
		for elidx, el := range query {
			if val, ok := getAttrToken(&tokens[i+elidx], "class"); ok != nil || val != el {
				good = false
			}
		}
		if good {
			return true
		}
	}

	return false
}

func tokenContainsClass(token *html.Token, query string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "class" && attr.Val == query {
			return true
		}
	}
	return false
}
