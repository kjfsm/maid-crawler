package main

import (
	"fmt"

	"github.com/sclevine/agouti"
)

func main() {
	url := "https://www.cafe-athome.com/maids"

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		fmt.Printf("Start error: %v\n", err)
		return
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		fmt.Printf("NewPage error: %v\n", err)
		return
	}

	if err := page.Navigate(url); err != nil {
		fmt.Printf("Navigate error: %v\n", err)
		return
	}
	content, err := page.HTML()
	if err != nil {
		fmt.Printf("HTML error: %v\n", err)
		return
	}
	fmt.Println(content)
}
