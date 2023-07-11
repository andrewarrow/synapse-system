package chrome

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Run(email string) {
	url := "https://manypw.slack.com/"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res []*cdp.Node
	err := chromedp.Run(ctx,
		submit(url, `//input[@name="email"]`, email, res))
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range res {

		var innerHTML string
		chromedp.Run(ctx,
			chromedp.InnerHTML(item.FullXPath(), &innerHTML),
		)
		fmt.Println(innerHTML)
	}

}

func submit(urlstr, sel, q string, res []*cdp.Node) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitReady("body"),
		chromedp.Nodes("//*", &res),
	}
}
