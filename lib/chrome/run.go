package chrome

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Run() {
	url := "https://manypw.slack.com/client"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		//		chromedp.Nodes("//p[text()] | //li[text()]", &res),
		chromedp.Nodes("//*", &res),
	)
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
