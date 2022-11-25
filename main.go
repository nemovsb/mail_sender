package main

import (
	"errors"
	"log"
	"mail_sender/internal/aggregator"
	"mail_sender/internal/app"
	"mail_sender/internal/cfg"
	"mail_sender/internal/http_server"
	"mail_sender/internal/http_server/gin_router"
	"mail_sender/internal/sender"
	"mail_sender/internal/storage"
	"mail_sender/internal/tracker"
	"os"
)

var ErrOsSignal = errors.New("got os signal")

func main() {

	config, err := cfg.ViperConfigurationProvider(os.Getenv("GOLANG_ENVIRONMENT"), false)
	if err != nil {
		log.Fatal("Read config error: ", err)
	}

	servConfig := http_server.NewServerConfig(config.HttpServer.Port)

	aggregator := aggregator.Aggregator{}

	storage := storage.NewMockStorage()
	log.Println("storage ready")

	sender := sender.NewSender(config.Sender.From, config.Sender.Password, config.Sender.SmtpHost, config.Sender.SmtpPort)
	log.Println("sender ready")

	tracker := tracker.NewTracker()

	application := app.NewApp(storage, aggregator, sender, tracker)
	log.Println("application ready")

	handler := gin_router.NewHandler(application)
	log.Println("handler ready")

	router := gin_router.NewRouter(handler)
	log.Println("router ready")

	server := http_server.NewServer(servConfig, router)
	log.Println("server ready")

	checker := app.GetChecker(&application)

	go checker()

	server.Run()

}
