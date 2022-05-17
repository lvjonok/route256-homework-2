package parser

import "strings"

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
