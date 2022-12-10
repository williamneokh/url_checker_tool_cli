# How to use

## Pre-prerequisite
1. Get your telegram bot api token
2. Create telegram channel and get the ID. (need to invite your bot into the channel)
3. Create an .env file with the following variable:
    BOTTOKEN=<your bot api token here e.g 1234567890:AAHYlkjsd8slokfo9sidfedsdf_dfrld>
    CHANNEL=<your telegram channel id here e.g -837423794>

## Configure your application

Inside main.go these are the setting that you can change:
1. App.Repeat = false     // set bool to turn on/off repeat mode base on the second given below (App.RepeatTime)
2. App.RepeatTime = 20    //number in seconds
3. App.IsNotifyBot = true //set bool to turn on/off the telegram notification
4. App.IsGoRoutine = true //set bool to turn on/off the go routine capability
5. checker.URLCheckStart(testdata.<option>, &wg) // option: URLshort, URL100, URL500 and URL1000

## To run the application in the terminal

1. type: go run .