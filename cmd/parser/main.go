package main

import (
	"fmt"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/parser"
)

func main() {
	_, err := parser.ParseCategories()
	fmt.Println(err)

	// if err == nil {
	// 	for _, val := range res {
	// 		fmt.Println(val)
	// 	}
	// }
	// out := make(chan *ProblemWError)

	// res, err := ParseProblemsIds(81)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("res %v", res)
	// fmt.Printf("err %v", err)

	// res := []int{541821} //, 510312, 26766, 26765, 26761, 503310}

	// var wg sync.WaitGroup

	// for _, id := range res {
	// 	wg.Add(1)
	// 	go func(problemId int) {
	// 		defer wg.Done()

	// 		fmt.Printf("start %d\n", problemId)
	// 		ParseProblem(problemId, out)
	// 	}(id)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(out)
	// }()

	// // wg.Wait()
	// // go ParseProblem(508951, out)
	// for x := range out {
	// 	// if x.Error != nil || x.Problem.Answer == "" {
	// 	fmt.Printf("output chan: %#v\n", x.Problem)
	// 	fmt.Printf("output chan: %#v\n\n", x.Error)
	// 	// }
	// }

}
