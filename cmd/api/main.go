package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"notif/internal/api"
	"notif/internal/db"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/TomBowyerResearchProject/common/middlewares"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"github.com/TomBowyerResearchProject/common/verification"
)

const timeBeforeTimeout = 15

func main() {
	logger.InitLogger("notif", logger.EmailConfig{
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Level:    os.Getenv("EMAIL_LEVEL"),
	})

	verification.Init(verification.VerificationConfig{
		VerificationURL:     os.Getenv("VERIFICATION_URL"),
		AuthorizationSecret: os.Getenv("AUTHORIZATION_SECRET"),
	})

	middlewares.Init(middlewares.Config{
		AllowedOrigin:  "*",
		AllowedMethods: "GET,POST,OPTIONS",
		AllowedHeaders: "*",
	})

	err := commonMongo.Connect(commonMongo.Config{
		URI:    os.Getenv("DATABASE_URL"),
		DBName: db.DBName,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	router := api.CreateRouter()

	srv := http.Server{
		Handler:      router,
		Addr:         os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		WriteTimeout: timeBeforeTimeout * time.Second,
		ReadTimeout:  timeBeforeTimeout * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		logger.Infof("Shutting down server")

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Infof("HTTP server Shutdown: %v", err)
		}

		monogodb := commonMongo.GetDatabase()
		if monogodb != nil {
			_ = monogodb.Client().Disconnect(context.Background())
		}

		logger.Infof("mongo disconnected")

		close(idleConnsClosed)
	}()

	logger.Info("Starting Server")

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Infof("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
