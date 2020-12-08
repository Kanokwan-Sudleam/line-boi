package main

import (
	"log"
	"os"

	"github.com/Kanokwan-Sudleam/line-boi/servicemanagement"
	"github.com/Kanokwan-Sudleam/line-boi/servicemanagement/delivery/http"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	startService()
}

func connectLineBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func startService() {
	e := echo.New()
	bankCoreInfo := servicemanagement.NewBankCoreServiceInfo()
	http.NewServiceHTTPHandler(e, connectLineBot(), bankCoreInfo)
	e.Logger.Fatal(e.Start(":6000"))
}
