package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Game struct {
	Name         string
	Combinations string
	DrawDate     string
	JackpotPrice string
	Winners      int
}

func main() {
	var result []Game

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tbody > tr", func(i int, tr *colly.HTMLElement) {
			if tr.Index == 0 {
				return
			}

			winners, _ := strconv.Atoi(tr.ChildText("td:nth-child(5)"))

			result = append(result, Game{
				Name:         tr.ChildText("td:nth-child(1)"),
				Combinations: tr.ChildText("td:nth-child(2)"),
				DrawDate:     tr.ChildText("td:nth-child(3)"),
				JackpotPrice: tr.ChildText("td:nth-child(4)"),
				Winners:      winners,
			})
		})
	})

	c.Visit("https://www.pcso.gov.ph/SearchLottoResult.aspx")

	content, _ := json.MarshalIndent(result, "", "  ")

	os.WriteFile("results.json", content, os.ModePerm)
	fmt.Println("Done scraping. Check 'results.json'")
}
