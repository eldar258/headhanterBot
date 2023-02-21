package spam

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"main/consts"
	"strings"
	"time"
)

const (
	selectorCompanyName     = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div.bloko-column.bloko-column_container.bloko-column_xs-4.bloko-column_s-8.bloko-column_m-12.bloko-column_l-10 > div:nth-child(2) > div > div.bloko-column.bloko-column_xs-4.bloko-column_s-8.bloko-column_m-12.bloko-column_l-6 > div > div > div > div.vacancy-company-details > span > a > span > span"
	jsClickButton           = `document.querySelector("div > div > div > div > div > div > div > div > div > div > div > div > div> div > div > a").click()`
	selectorMessageBtn      = "#RESPONSE_MODAL_FORM_ID > div > div > div:nth-child(3) > button"
	selectorMessageTextArea = "#RESPONSE_MODAL_FORM_ID > div > div > textarea"
	selectorSendBtn         = "body > div.bloko-modal-overlay.bloko-modal-overlay_visible > div > div.bloko-modal > div.bloko-modal-footer > div > button.bloko-button.bloko-button_kind-primary"
)

func Spam(ctx context.Context, links [50]string) (context.Context, int) {
	successfulCount := 0
	for _, link := range links {
		err := spam(ctx, link)
		if err == nil {
			successfulCount++
		}
	}

	return ctx, successfulCount
}

func spam(ctx context.Context, link string) error {
	newTab, cancel := chromedp.NewContext(ctx)
	defer cancel()

	newTab, cancel = context.WithTimeout(newTab, 20*time.Second)
	defer cancel()

	var companyName string
	var messageBtn []*cdp.Node
	err := chromedp.Run(newTab,
		chromedp.Navigate(link),
		chromedp.TextContent(selectorCompanyName, &companyName, chromedp.NodeVisible),
		chromedp.Evaluate(jsClickButton, nil),
		chromedp.WaitVisible(selectorSendBtn),
		chromedp.Nodes(selectorMessageBtn, &messageBtn, chromedp.AtLeast(0)),
	)
	if err != nil {
		LoggerError(err, link, "WHEN CLICK SEND BTN")
		return err
	}
	log.Println("Start spam to ", companyName)

	message := strings.Replace(consts.MESSAGE_TAMPLATE, "*", companyName, -1)
	log.Println(len(messageBtn))
	if len(messageBtn) > 0 {
		err = chromedp.Run(newTab,
			chromedp.Click(selectorMessageBtn),
		)
		if err != nil {
			LoggerError(err, link, "WHEN CLICK MESSAGE BTN")
			return err
		}
	}

	err = chromedp.Run(newTab,
		chromedp.SendKeys(selectorMessageTextArea, message),
		chromedp.Click(selectorSendBtn),
		chromedp.WaitNotPresent(selectorSendBtn),
	)
	if err != nil {
		LoggerError(err, link, "WHEN SEND MESSAGE")
		return err
	} else {
		log.Println("Successful spam to ", companyName, " link:", link)
	}

	return err
}

func LoggerError(err error, strings ...string) {
	log.Println(err)
	for i := range strings {
		log.Println(strings[i])
	}

	log.Println()
}
