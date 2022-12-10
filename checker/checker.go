package checker

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/williamneok/urlcheckertoolcli/config"
	"github.com/williamneok/urlcheckertoolcli/notification"
)

var app *config.AppConfig

func NewCheckerApp(a *config.AppConfig) {
	app = a
}

func URLCheckStart(urlAddress []string, wg *sync.WaitGroup) {
	if app.Repeat {
		if app.RepeatTime > 0 {
			for {
				app.CheckCount = 0
				for _, url := range urlAddress {
					wg.Add(1)
					app.CheckCount += 1
					if app.IsGoRoutine {
						go UrlCheck(url, wg)
					} else {
						UrlCheck(url, wg)
					}

				}
				time.Sleep(time.Duration(app.RepeatTime) * time.Second)
			}
		} else {
			log.Fatal("please set repeat timer more than 1 second")
		}
	} else {
		app.CheckCount = 0
		for _, url := range urlAddress {
			wg.Add(1)
			app.CheckCount += 1
			if app.IsGoRoutine {
				go UrlCheck(url, wg)
			} else {
				UrlCheck(url, wg)
			}

		}
	}

}

func UrlCheck(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	ping := http.Client{
		Timeout: 15 * time.Second,
	}

	res, err := ping.Get(url)

	if err != nil {
		fmt.Printf("url: %v - unreachable\n", url)
		if app.IsNotifyBot {
			notification.TelegramBotMsg(fmt.Sprintf("url: %v", url))
		}

		return
	}
	fmt.Printf("url: %v - %v\n", url, res.Status)

}
