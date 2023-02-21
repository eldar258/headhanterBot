package search

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"main/consts"
)

const (
	selectorLinks = "div > div > div > div > h3 > span > a"
)

func Search(ctx context.Context) (context.Context, [50]string) {
	var nodesLink []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(consts.SEARCH_LINK),
		chromedp.Nodes(selectorLinks, &nodesLink),
	)
	if err != nil {
		log.Panicln(err)
	}

	var result [50]string
	for i := range nodesLink {
		result[i] = nodesLink[i].AttributeValue("href")
	}

	return ctx, result
}
