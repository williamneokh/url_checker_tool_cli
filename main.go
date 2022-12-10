package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/williamneok/urlcheckertoolcli/checker"
	"github.com/williamneok/urlcheckertoolcli/config"
	"github.com/williamneok/urlcheckertoolcli/envtool"
	"github.com/williamneok/urlcheckertoolcli/notification"
	"github.com/williamneok/urlcheckertoolcli/testdata"
)

func main() {

	var App config.AppConfig

	App.Repeat = false     // set bool to turn on/off repeat mode base on the second given below (App.RepeatTime)
	App.RepeatTime = 20    //number in seconds
	App.IsNotifyBot = true //set bool to turn on/off the telegram notification
	App.IsGoRoutine = true //set bool to turn on/off the go routine capability
	App.BotToken = envtool.GoDotEnvVariable("BOTTOKEN")
	App.TelegramChannel = envtool.GoDotEnvVariable("CHANNEL")
	checker.NewCheckerApp(&App)
	notification.NewNotificationConfig(&App)
	timeNow := time.Now()
	defer func() {
		fmt.Printf("time taken: %v for total: %v link checked", time.Since(timeNow), App.CheckCount)
	}()

	var wg sync.WaitGroup

	checker.URLCheckStart(testdata.CustomURL, &wg)

	wg.Wait()
}
