package open

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

const (
	selectorEmail  = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div:nth-child(1) > div.account-login-tile-content-wrapper > div.account-login-tile-content > div > div:nth-child(2) > div > form > div.bloko-form-item > fieldset > input"
	selectorNext   = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div:nth-child(1) > div.account-login-tile-content-wrapper > div.account-login-tile-content > div > div:nth-child(2) > div > form > div.account-login-actions > button.bloko-button.bloko-button_kind-primary"
	email          = "9eldik7@gmail.com"
	selectorCode   = "#HH-React-Root > div > div.HH-MainContent.HH-Supernova-MainContent > div.main-content > div > div > div > div > div > div > div > div > div > div > div > form > div > div.verification-content > div > fieldset > input"
	xpathNext2     = "/html/body/div[5]/div/div[3]/div[1]/div/div/div/div/div/div/div/div/div/div/div/form/div/div[7]/button[not(@disabled)]"
	selectorSearch = "//*[@id=\"HH-React-Root\"]/div/div[3]/div[1]/div[2]/div/div[1]/div[1]/div[1]/div/div[1]/div[2]/a[1]/span/span[1]"
)

func Open() context.Context {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	var buff []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://hh.ru/account/login?backurl=%2F&hhtmFrom=main&role=applicant"),
		chromedp.SendKeys(selectorEmail, email, chromedp.NodeVisible),
		chromedp.Click(selectorNext, chromedp.NodeVisible),

		chromedp.SendKeys(selectorCode, getCode(), chromedp.NodeVisible),
		chromedp.Click(xpathNext2, chromedp.NodeVisible),
		chromedp.WaitVisible(selectorSearch),

		chromedp.FullScreenshot(&buff, 100),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("fileName.png", buff, 0644); err != nil {
		log.Fatal(err)
	}

	return ctx
}

func getCode() string {
	fmt.Println("Enter code from email ", email)

	var result string
	if _, err := fmt.Scanln(&result); err != nil {
		log.Fatal(err)
	}

	return result
}
