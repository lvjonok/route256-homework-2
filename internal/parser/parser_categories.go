package parser

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

type ProblemCategory struct {
	Problem    int
	CategotyId int
	Title      string
}

func ParseCategories() ([]*ProblemCategory, error) {
	url := "https://ege.sdamgia.ru/prob_catalog"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	categories := []*ProblemCategory{}

	// in parsed website each category follows

	tokens := []html.Token{}

	lastProblem := 0

out:
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			if ok, _ := isProblemSpan(tokens); ok {
				pb, err := strconv.ParseInt(tokenizer.Token().Data, 10, 32)
				if err != nil {
					fmt.Println("errored", err)
					return nil, err
				}
				lastProblem = int(pb)

				// we need to parse only 1-11 tasks
				if lastProblem == 12 {
					break out
				}
				// fmt.Printf("span last %v\n", pb)
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

			// res, _ := isLastCategory(tokens)
			// if res {
			// 	// fmt.Printf("last is category %#v\n", tokens[len(tokens)-1])
			// 	link, err := getAttrToken(&t, "href")
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	// fmt.Printf("link for category %v\n", link)
			// }
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
	// last token should be cat_name<-cat_category<-cat_children
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

func isLastCategory(tokens []html.Token) (bool, error) {
	// general pattern of class attr is: cat_category -> cat_children -> cat_category
	return isEndingClassNames(tokens, []string{"cat_show", "cat_category", "cat_children", "cat_category"})
}
