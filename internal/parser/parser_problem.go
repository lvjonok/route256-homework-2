package parser

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ParseProblem(problemId int) (*Problem, error) {
	url := fmt.Sprintf("%s/problem?id=%d&print=true", baseURL, problemId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get page for problem parsing, err: <%v>", err)
	}
	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)

	problem := Problem{ProblemId: problemId, ProblemImage: "", Parts: []string{}, Answer: ""}

	tokens := []html.Token{}

out:
	for {
		tt := tokenizer.Next()
		curToken := tokenizer.Token()
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			if isDescription(tokens) {
				problem.AddPart(curToken.Data)
			}

			if isAnswer(tokens) {
				problem.AddAnswer(curToken.Data)
				break out
			}
		case tt == html.EndTagToken:
			tokens = tokens[:len(tokens)-1]
		case tt == html.SelfClosingTagToken:
			// check that image in description
			if val, ok := getAttrToken(&curToken, "class"); ok == nil && val == "tex" && isDescription(tokens) {
				imgSrc, _ := getAttrToken(&curToken, "src")
				problem.AddPart(imgSrc)
			}
		case tt == html.StartTagToken:
			tokens = append(tokens, curToken)
			if curToken.Data == "img" && isDescription(tokens) {
				imgSrc, _ := getAttrToken(&curToken, "src")
				imgSrc = strings.TrimSpace(imgSrc)
				problem.ProblemImage = baseURL + imgSrc
			}
		}
	}

	return &problem, nil
}

func isDescription(tokens []html.Token) bool {
	query := []string{"nobreak", "pbody", "left_margin"}

	anySolution := false
	for _, token := range tokens {
		if tokenContainsClass(&token, "solution") {
			anySolution = true
		}
	}

	return tokensContain(tokens, query) && !anySolution
}

func isAnswer(tokens []html.Token) bool {
	query := []string{"answer"}

	return tokensContain(tokens, query)
}
