package parser

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

// https://ege.sdamgia.ru/test?filter=all&category_id=14&ttest=true
func ParseProblemsIds(categoryId int) ([]int, error) {
	url := fmt.Sprintf("%s/test?filter=all&category_id=%d&ttest=true", baseURL, categoryId)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get problem for problems ids parsing, err: <%v>", err)
	}
	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)

	problemIds := []int{}
	tokens := []html.Token{}

out:
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			res, err := isProblemId(tokens)
			if err != nil {
				return nil, fmt.Errorf("failed to check if problem id, err: <%v>", err)
			}
			if res {
				idInt, err := strconv.Atoi(tokenizer.Token().Data)
				if err != nil {
					return nil, fmt.Errorf("failed to atoi problem id %v, err: %v", tokenizer.Token(), err)
				}

				problemIds = append(problemIds, idInt)
			}
		case tt == html.EndTagToken:
			tokens = tokens[:len(tokens)-1]
		case tt == html.StartTagToken:
			tokens = append(tokens, tokenizer.Token())
		}
	}

	return problemIds, nil
}

func isProblemId(tokens []html.Token) (bool, error) {
	// pattern for problem id is a(href="...")<-prob_nums
	if len(tokens) < 2 {
		return false, nil
	}

	// last token is without className, just text
	ok, err := isEndingClassNames(tokens[:len(tokens)-1], []string{"prob_nums"})
	if err != nil {
		return false, err
	}

	_, err = getAttrToken(&tokens[len(tokens)-1], "href")
	if ok && err == nil {
		return true, nil
	}

	return false, nil
}
