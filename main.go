package main

import (
	"errors"
	"fmt"
	"log"
	"mail_sender/internal/aggregator"
	"mail_sender/internal/app"
	"mail_sender/internal/cfg"
	"mail_sender/internal/http_server"
	"mail_sender/internal/http_server/gin_router"
	"mail_sender/internal/sender"
	"mail_sender/internal/storage"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/run"
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

	sender := sender.NewSender(config.Sender.From, config.Sender.Password, config.Sender.SmtpHost, config.Sender.SmtpPort)

	application := app.NewApp(storage, aggregator, sender)

	handler := gin_router.NewHandler(application)

	router := gin_router.NewRouter(handler)

	server := http_server.NewServer(servConfig, router)

	var (
		serviceGroup        group.Group
		interruptionChannel = make(chan os.Signal, 1)
	)

	serviceGroup.Add(func() error {
		signal.Notify(interruptionChannel, syscall.SIGINT, syscall.SIGTERM)
		osSignal := <-interruptionChannel

		return fmt.Errorf("%w: %s", ErrOsSignal, osSignal)
	}, func(error) {
		interruptionChannel <- syscall.SIGINT
	})

	serviceGroup.Add(func() error {
		log.Println("HTTP API started")

		return server.Run()
	}, func(error) {
		err = server.Shutdown()
		log.Printf("shutdown Http Server error: %s", err)
	})

}
