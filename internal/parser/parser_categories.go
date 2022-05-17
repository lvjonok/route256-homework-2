package parser

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

type ProblemCategory struct {
	Problem    int    // task number
	CategotyId int    // id of category from reshuege
	Title      string // title of category
}

func ParseCategories() ([]*ProblemCategory, error) {
	url := fmt.Sprintf("%s/prob_catalog", baseURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	categories := []*ProblemCategory{}
	tokens := []html.Token{}
	lastProblem := 0 // task number - [1, 11]

out:
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			if ok, _ := isProblemSpan(tokens); ok {
				pb, err := strconv.Atoi(tokenizer.Token().Data)
				if err != nil {
					fmt.Println("errored", err)
					return nil, err
				}
				lastProblem = int(pb)

				// we need to parse only 1-11 tasks
				if lastProblem == 12 {
					break out
				}
			}
			ok, err := isCategoryTitle(tokens)
			if err != nil {
				return nil, err
			}
			if ok {
				title := tokenizer.Token().Data

				catId, err := parseCategoryId(tokens)
				if err != nil {
					return nil, err
				}

				categories = append(categories, &ProblemCategory{
					Problem:    lastProblem,
					CategotyId: catId,
					Title:      title,
				})
			}

		case tt == html.EndTagToken:
			tokens = tokens[:len(tokens)-1] // remove last
		case tt == html.StartTagToken:
			t := tokenizer.Token()
			// add new token for our path
			tokens = append(tokens, t)
		}
	}

	return categories, nil
}

func isProblemSpan(tokens []html.Token) (bool, error) {
	if len(tokens) < 1 {
		return false, nil
	}
	return isEndingClassNames(tokens, []string{"pcat_num"})
}

func isCategoryTitle(tokens []html.Token) (bool, error) {
	if len(tokens) < 4 {
		return false, nil
	}
	return isEndingClassNames(tokens, []string{"cat_name", "cat_category", "cat_children", "cat_category"})
}

func parseCategoryId(tokens []html.Token) (int, error) {
	categoryIdRaw, err := getAttrToken(&tokens[len(tokens)-2], "data-id")
	if err != nil {
		return 0, fmt.Errorf("failed to parse category, err: %v", err)
	}
	categoryId, err := strconv.Atoi(categoryIdRaw)
	if err != nil {
		return 0, fmt.Errorf("failed to convert to int category id: %v, err: %v", categoryIdRaw, err)
	}

	return categoryId, nil
}
