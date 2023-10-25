package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	crawler := colly.NewCollector()

	crawler.OnHTML("table.infobox-recipe", func(e *colly.HTMLElement) {
		table := e.DOM.Nodes[0]
		tbody := table.FirstChild
		row := tbody.FirstChild
		inMaterialSection := false
		for row != nil {
			firstColumn := row.FirstChild
			firstColumnChild := firstColumn.FirstChild
			if firstColumn.Data == "th" && firstColumnChild.Data == "Material" {
				inMaterialSection = true
			} else if firstColumn.Data == "th" && firstColumnChild.Data == "Total cost" {
				inMaterialSection = false
			} else if inMaterialSection {
				column := firstColumn.NextSibling
				materialName := column.FirstChild.FirstChild.Data
				column = column.NextSibling
				materialQuantity := column.FirstChild.Data
				fmt.Printf("Material: %s\n", materialName)
				fmt.Printf("Quantity: %s\n", materialQuantity)
			}
			row = row.NextSibling
		}
	})

	crawler.Visit("https://runescape.wiki/w/Kal-i-kra_mace")
}
