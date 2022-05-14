package parser

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Problem struct {
	ProblemId    int
	ProblemImage string   // main problem image, might be null
	Parts        []string // problem description, formulas as images might follow text
	Answer       string
}

func (p *Problem) AddPart(el string) {
	el = strings.ReplaceAll(el, "\u202f", " ")
	el = strings.ReplaceAll(el, "\u2009", " ")
	el = strings.TrimSpace(el)

	if len(p.Parts) != 0 && !isLink(p.Parts[len(p.Parts)-1]) && !isLink(el) {
		extra := " "
		if el == "." || el == "," {
			extra = ""
		}

		p.Parts[len(p.Parts)-1] += extra + el
	} else {
		p.Parts = append(p.Parts, el)
	}
}

func (p *Problem) AddAnswer(el string) {
	el = strings.TrimSpace(el)
	el = strings.TrimPrefix(el, "Ответ:")
	el = strings.TrimSuffix(el, ".")
	el = strings.TrimSpace(el)

	p.Answer = el
}

const problemImageBaseUrl string = "https://ege.sdamgia.ru"

type ProblemWError struct {
	Problem *Problem
	Error   error
}

// https://ege.sdamgia.ru/test?filter=all&category_id=14&ttest=true

func ParseProblemsIds(categoryId int) ([]int, error) {
	url := fmt.Sprintf("https://ege.sdamgia.ru/test?filter=all&category_id=%d&ttest=true", categoryId)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("posos", err)
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
				return nil, err
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

func ParseProblem(problemId int) (*Problem, error) {
	url := fmt.Sprintf("https://ege.sdamgia.ru/problem?id=%d&print=true", problemId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("posos", err)
	}
	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)

	problem := Problem{ProblemId: problemId, ProblemImage: "", Parts: []string{}, Answer: ""}

	tokens := []html.Token{}

out:
	for {
		tt := tokenizer.Next()
		curToken := tokenizer.Token()
		// log.Printf("token: %#v", curToken)
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			if isDescription(tokens) {
				problem.AddPart(curToken.Data)
			}

			if isAnswer(tokens) {
				problem.AddAnswer(curToken.Data)
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
				problem.ProblemImage = problemImageBaseUrl + imgSrc
			}
		}
	}

	return &problem, nil
}

func isDescription(tokens []html.Token) bool {
	query := []string{"nobreak", "pbody", "left_margin"}

	return tokensContain(tokens, query)
}

func isAnswer(tokens []html.Token) bool {
	query := []string{"answer"}

	return tokensContain(tokens, query)
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
