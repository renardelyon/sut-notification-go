package main

import (
	"log"
	"sut-notification-go/application"
	"sut-notification-go/config"
	"sut-notification-go/domain/notification/service"
	notifpb "sut-notification-go/pb/notification"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err.Error())
	}

	app, err := application.Setup(&c)
	if err != nil {
		log.Fatalln("Failed at application setup: ", err.Error())
	}

	s := service.NewService(app.DbClients)

	notifpb.RegisterNotificationServiceServer(app.GrpcServer, s)

	err = app.Run(&c)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
