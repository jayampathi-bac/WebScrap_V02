package main

import(
	"fmt"
	"github.com/gocolly/colly"
)
func main() {
	var location string;
	var category string;
	c := colly.NewCollector()
	fmt.Print("Enter Location : ")
	fmt.Scan(&location)
	fmt.Print("Enter Category : ")
	fmt.Scan(&category)

	c.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		fmt.Println(element.Text,"\n")
	})

	c.Visit("https://ikman.lk/en/ads/"+ location +"/" +category);


}
