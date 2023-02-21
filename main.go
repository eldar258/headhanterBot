package main

import (
	"log"
	"main/open"
	"main/search"
	"main/spam"
)

func main() {
	ctx, cancel := open.Open()
	defer cancel()

	ctx, links := search.Search(ctx)
	log.Println("Searched ", len(links), " link(s)")

	ctx, successfulCount := spam.Spam(ctx, links)
	log.Println("Response ", successfulCount)
}
