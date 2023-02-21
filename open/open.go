package open

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"main/consts"
)

const (
	selectorEmail = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div:nth-child(1) > div.account-login-tile-content-wrapper > div.account-login-tile-content > div > div:nth-child(2) > div > form > div.bloko-form-item > fieldset > input"
	selectorNext  = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div:nth-child(1) > div.account-login-tile-content-wrapper > div.account-login-tile-content > div > div:nth-child(2) > div > form > div.account-login-actions > button.bloko-button.bloko-button_kind-primary"

	selectorCode   = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div > div > div > div > div > div > form > div > div.verification-content > div > fieldset > input"
	xpathNext2     = "/html/body/div[5]/div/div[3]/div[1]/div/div/div/div/div/div/div/div/div/div/div/form/div/div[7]/button[not(@disabled)]"
	selectorSearch = "//*[@id=\"HH-React-Root\"]/div/div[3]/div[1]/div[2]/div/div[1]/div[1]/div[1]/div/div[1]/div[2]/a[1]/span/span[1]"
)

func Open() (context.Context, context.CancelFunc) {
	ctx, _ := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...)
	ctx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://hh.ru/account/login?backurl=%2F&hhtmFrom=main&role=applicant"),

		chromedp.SendKeys(selectorEmail, consts.EMAIL),
		chromedp.Click(selectorNext),

		chromedp.WaitVisible(selectorCode),
	)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	err = chromedp.Run(ctx,
		chromedp.SendKeys(selectorCode, getCode()),
		chromedp.Click(xpathNext2),

		chromedp.WaitVisible(selectorSearch),
	)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	return ctx, cancel
}

func getCode() string {
	fmt.Println("Enter code from email ", consts.EMAIL)

	var result string
	if _, err := fmt.Scanln(&result); err != nil {
		log.Println(err)
	}

	return result
}
