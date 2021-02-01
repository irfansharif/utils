// Fetches metadata for a given book.
//
// Usage:
//  $ gbooks -t 'ready player one'
//    {
//      "Title": "A Clash of Kings",
//      "Details": "Book Two of A Song of Ice and Fire",
//      "Author": "George R. R. Martin"
//    }
//
// The search string could actually be anything, cause all it does is send it
// out to the google books API. Specifying the author, year, or other relevant
// search terms, will help refine the search.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

var title = flag.String("t", "", "name of the book")

type book struct {
	Title   string
	Details string
	Author  string
}

func main() {
	flag.Parse()
	if title == nil || *title == "" {
		log.Fatal("no title provided")
	}
	book, err := fetch(*title)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.MarshalIndent(book, "", "\t")
	fmt.Printf("%s\n", json)
}
