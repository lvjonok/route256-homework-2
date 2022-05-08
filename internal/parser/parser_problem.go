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
	if len(p.Parts) != 0 && !isLink(p.Parts[len(p.Parts)-1]) {
		p.Parts[len(p.Parts)-1] += el
	} else {
		p.Parts = append(p.Parts, el)
	}
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

func ParseProblem(problemId int, output chan<- *ProblemWError) {
	url := fmt.Sprintf("https://ege.sdamgia.ru/problem?id=%d", problemId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("posos", err)
	}
	defer resp.Body.Close()
	tokenizer := html.NewTokenizer(resp.Body)

	problem := Problem{ProblemId: problemId, ProblemImage: "", Parts: []string{}, Answer: ""}

	seenAnswerLabel := false

	tokens := []html.Token{}

out:
	for {
		tt := tokenizer.Next()
		curToken := tokenizer.Token()
		switch {
		case tt == html.ErrorToken:
			break out
		case tt == html.TextToken:
			if tokens[len(tokens)-1].Data == "i" {
				problem.AddPart(curToken.Data)
				fmt.Printf("I tag %#v\n", curToken)
				continue
			}

			if tokens[len(tokens)-1].Data == "sup" {
				problem.AddPart("^" + curToken.Data)
				fmt.Printf("sup tag %#v\n", curToken)
				continue
			}

			if curToken.Data == "Решение" {
				seenAnswerLabel = true
				// fmt.Println("seen решение")
				// fmt.Println(problem)
			}

			// if curToken.Data == " На рисунке изображён график функции вида " {
			// 	fmt.Printf("found answer %v\n", curToken)
			// 	for _, tok := range tokens {
			// 		fmt.Printf("tok: %#v\n", tok)
			// 	}
			// }

			if problem.Answer == "" && !seenAnswerLabel {
				res, err := isProblemDescription(tokens)
				if err != nil {
					output <- &ProblemWError{Problem: &problem, Error: err}
					return
				}

				if res {
					curToken.Data = strings.ReplaceAll(curToken.Data, "\u202f", " ")
					curToken.Data = strings.ReplaceAll(curToken.Data, "\u2009", " ")
					problem.AddPart(curToken.Data)
				}
			}

			res, err := isProblemAnswer(tokens)
			if err != nil {
				output <- &ProblemWError{Problem: &problem, Error: err}
				return
			}
			// if res || strings.TrimSpace(curToken.Data) != "" && strings.Contains(curToken.Data, ": -4.") {
			// 	fmt.Printf("found answer %v\n", curToken)
			// 	for _, tok := range tokens {
			// 		fmt.Printf("tok: %#v\n", tok)
			// 	}
			// 	fmt.Printf("\n\n\t\tres: %v, text: '%v' %v\n", res, curToken.Data, strings.Contains(curToken.Data, ": -4."))
			// }

			if res && strings.TrimSpace(curToken.Data) != "" {
				// fmt.Printf("\twrote answer %v\n", curToken)
				// // fmt.Printf("found answer\n")
				// // for _, tok := range tokens {
				// // 	fmt.Printf("tok: %#v\n", tok)
				// // }

				rawAnswer := curToken.Data
				rawAnswer = strings.TrimSpace(rawAnswer)
				rawAnswer = strings.TrimSuffix(rawAnswer, ".")

				splitted := strings.Split(rawAnswer, ":")
				if len(splitted) != 2 {
					// output <- &ProblemWError{Problem: &problem, Error: fmt.Errorf("fail to parse raw answer %v %v", problemId, curToken.Data)}
					// return
					break
				}
				// rawAnswer = strings.TrimPrefix(rawAnswer, ":")
				rawAnswer = strings.TrimSpace(splitted[1])

				problem.Answer = rawAnswer
			}

		case tt == html.EndTagToken:
			tokens = tokens[:len(tokens)-1]
		case tt == html.SelfClosingTagToken:
			// image is a part of descriptions (formulas)
			if curToken.Data == "img" {
				res, err := isImageDescr(tokens)
				if err != nil {
					output <- &ProblemWError{Problem: &problem, Error: err}
					return
				}

				if !res {
					continue
				}

				imgSrc, err := getAttrToken(&curToken, "src")
				if !strings.HasPrefix(imgSrc, "https://ege.sdamgia.ru/formula/svg") {
					continue
				}
				// fmt.Printf("try descr %#v\n", curToken)
				if err != nil {
					output <- &ProblemWError{Problem: &problem, Error: fmt.Errorf("image has no src: %v err: %v", curToken, err)}
					return
				} else if !seenAnswerLabel {
					problem.Parts = append(problem.Parts, imgSrc)
					// fmt.Printf("added new image %v\n", imgSrc)
				}

				// if !res {
				// 	continue
				// }

			}
		case tt == html.StartTagToken:
			tokens = append(tokens, curToken)

			res, err := isImageProblem(tokens)
			if err != nil {
				output <- &ProblemWError{Problem: &problem, Error: err}
				return
			}

			if res {
				imgSrc, err := getAttrToken(&curToken, "src")
				if err != nil {
					output <- &ProblemWError{Problem: &problem, Error: fmt.Errorf("image has no src: %v err: %v", curToken, err)}
				} else {
					problem.ProblemImage = problemImageBaseUrl + imgSrc
				}
			}
		}
	}

	output <- &ProblemWError{Problem: &problem, Error: nil}
}

func isProblemAnswer(tokens []html.Token) (bool, error) {
	if len(tokens) < 5 {
		return false, nil
	}

	res, err := isEndingTypes(tokens, []string{"p", "p", "center", "p"})
	if err != nil {
		return false, fmt.Errorf("failed to get problem answer, err: %v", err)
	}
	val, _ := getAttrToken(&tokens[len(tokens)-5], "class")
	// val12, _ := getAttrToken(&tokens[len(tokens)-4], "class")

	res2, err := isEndingTypes(tokens, []string{"p", "p", "div", "div"})
	if err != nil {
		return false, fmt.Errorf("failed to get problem answer, err: %v", err)
	}
	val2, _ := getAttrToken(&tokens[len(tokens)-2], "class")
	divWidth, _ := getAttrToken(&tokens[len(tokens)-4], "width")

	res3, err := isEndingTypes(tokens, []string{"span", "div"})
	if err != nil {
		return false, fmt.Errorf("failed to get problem answer, err: %v", err)
	}
	val3, _ := getAttrToken(&tokens[len(tokens)-2], "class")

	res4, err := isEndingTypes(tokens, []string{"p", "p", "p"})
	if err != nil {
		return false, fmt.Errorf("failed to get problem answer, err: %v", err)
	}
	val41, _ := getAttrToken(&tokens[len(tokens)-2], "class")
	val42, _ := getAttrToken(&tokens[len(tokens)-3], "class")

	subr := (res3 && val3 == "answer") || (res2 && val2 == "left_margin" && divWidth != "100%")
	//  || (res && val12 == "left_margin")
	return (subr) || (res && val == "solution") || (res4 && val41 == "left_margin" && val42 == "left_margin"), nil
}

func isImageProblem(tokens []html.Token) (bool, error) {
	if len(tokens) < 2 {
		return false, nil
	}

	// last one should be <img> <- left_margin
	if tokens[len(tokens)-1].Data != "img" {
		return false, nil
	}

	// fmt.Printf("try image %v\n", tokens[len(tokens)-1])
	// for _, tok := range tokens[len(tokens)-4 : len(tokens)-1] {
	// 	fmt.Printf("tok: %#v\n", tok)
	// }

	return isProblemDescription(tokens[:len(tokens)-1])
}

func isImageDescr(tokens []html.Token) (bool, error) {
	if len(tokens) < 1 {
		return false, nil
	}

	return isProblemDescription(tokens)
}

func isProblemDescription(tokens []html.Token) (bool, error) {
	// pattern for description is left_margin<-img

	if len(tokens) < 3 {
		return false, nil
	}

	isDescr, err1 := isEndingClassNames(tokens, []string{"left_margin"})

	isDescr2, err1 := isEndingClassNames(tokens[:len(tokens)-1], []string{"left_margin"})

	isSol, err2 := isEndingClassNames(tokens, []string{"left_margin", "solution"})

	// check that not comments
	isComm, err3 := isEndingTypes(tokens, []string{"p", "div", "div"})
	if err3 != nil {
		return false, err3
	}

	if width, err := getAttrToken(&tokens[len(tokens)-3], "width"); err != nil && width != "100%" {
		isComm = false
	}

	// isComm := false

	if err1 != nil || err2 != nil {
		return false, fmt.Errorf("problem description: one of error %v or %v", err1, err2)
	}

	return (isDescr || isDescr2) && !isSol && !isComm, nil
}
