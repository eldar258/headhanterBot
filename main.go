package main

import (
	"main/open"
)

func main() {
	ctx := open.Open()

	ctx, links := search.Search(ctx)

	ctx = spam.Spam(ctx, links)
}
