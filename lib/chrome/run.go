package chrome

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

type Cookie struct {
	Name   string
	Value  string
	Domain string
}

func Run(email string) {
	url := "https://manypw.slack.com/client"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res []*cdp.Node
	cookies := []Cookie{}

	chromedp.Run(ctx, setcookies(
		url,
		cookies,
	))

	time.Sleep(time.Second)
	chromedp.Run(ctx,
		chromedp.Nodes("//*", &res),
	)

	for _, item := range res {

		var innerHTML string
		chromedp.Run(ctx,
			chromedp.InnerHTML(item.FullXPath(), &innerHTML),
		)
		fmt.Println(innerHTML)
	}

	/*
		fmt.Println("1")
		chromedp.Run(ctx,
			chromedp.Navigate(url),
		)
		fmt.Println("2")
		time.Sleep(time.Second)
		sel := `//input[@name='email']`
		chromedp.Run(ctx,
			chromedp.SetValue(sel, email, chromedp.BySearch),
			//chromedp.SendKeys(sel, email),
			chromedp.Click(`//button[@id='submit_btn']`),
		)
		fmt.Println("3")
		//chromedp.Run(ctx,
		//	chromedp.Click(`//button[@id='submit_btn']`),
		//)
		fmt.Println("4")
		//chromedp.Submit(sel).Do(ctx)
		//chromedp.WaitReady("body").Do(ctx)

	*/
}

func setcookies(host string, cookies []Cookie) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			//expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for _, cookie := range cookies {
				err := network.SetCookie(cookie.Name, cookie.Value).
					//WithExpires(&expr).
					WithDomain(cookie.Domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Navigate(host),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}
