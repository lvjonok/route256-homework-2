package parser

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

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

func getTokenType(tok *html.Token) string {
	return tok.Data
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

func isEndingTypes(tokens []html.Token, ends []string) (bool, error) {
	if len(tokens) < len(ends) {
		return false, fmt.Errorf("length of tokens is less than ends")
	}

	tl := len(tokens)

	for idx, val := range ends {
		if tokens[tl-idx-1].Data != val {
			return false, nil
		}
	}

	return true, nil
}

func isLink(el string) bool {
	// TODO: might change
	return strings.Contains(el, "http")
}
