package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  int
}

func printCheapBooks(books []Book, limit int) {
	for _, b := range books {
		if b.Price < limit {
			fmt.Printf("Title: %s , Author: %s, Price: %d\n", b.Title, b.Author, b.Price)
		}
	}
}

func main() {

	mybooks := []Book{
		{Title: "A", Author: "a", Price: 25000},
		{Title: "B", Author: "b", Price: 14000},
		{Title: "C", Author: "c", Price: 14000},
		{Title: "D", Author: "d", Price: 13000},
		{Title: "E", Author: "e", Price: 11000},
	}

	printCheapBooks(mybooks, 15000)

	//array

}
