package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"log"
)
func main() {
	var location string;
	var category string;
	c := colly.NewCollector()
	fmt.Print("Enter Location : ")
	fmt.Scan(&location)
	fmt.Print("Enter Category : ")
	fmt.Scan(&category)

	//c.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
	//	fmt.Println(element.Text,"\n")
	//})
	////////////////////////////////////////////////////////////////////////////////////
	db, err := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/ikmandb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	///////////////////////////////////////////////////////////////////////////////////
	c.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		title:= element.ChildText(".heading--2eONR")
		descr:= element.ChildText(".description--2-ez3")
		price :=element.ChildText(".price--3SnqI")
		fmt.Println("Item============================================")
		fmt.Println("\tTitle: ",title)
		fmt.Println("\tDescription: ",descr )
		fmt.Println("\tPrice: ",price )
		sql,err := db.Query("INSERT INTO item(title, descr, price,location, category) VALUES (?,?,?,?,?)",title, descr, price,location, category)
		if err != nil {
			log.Fatal(err)
		}
		defer sql.Close()

	})

	c.Visit("https://ikman.lk/en/ads/"+ location +"/" +category);


}
