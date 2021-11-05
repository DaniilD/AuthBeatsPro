package internal

import (
	"AuthBeatsPro/internal/configs"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Application struct {
	configLocator *configs.ConfigLocator
	server        *Server
	Router        *Router
}

func NewApplication(configLocator *configs.ConfigLocator, server *Server) *Application {
	return &Application{
		configLocator: configLocator,
		server:        server,
	}
}

func (application *Application) Configure() {
	application.configureServer()
	application.configureDB()
	application.configJWT()
	//application.configureSentry()
}

/*func (application *Application) configureSentry() {
	dsn := os.Getenv("SENTRY_DSN")
	release := os.Getenv("SENTRY_RELEASE")
	isDebug, err := strconv.Atoi(os.Getenv("SENTRY_DEBUG"))

	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.SentryConfigInstance().Dsn = dsn
	application.configLocator.SentryConfigInstance().Release = release
	application.configLocator.SentryConfigInstance().IsDebug = isDebug == 1
}*/

func (application *Application) configureServer() {
	host := os.Getenv("APPLICATION_HOST")
	port, err := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.ServerConfigInstance().Host = host
	application.configLocator.ServerConfigInstance().Port = port

}

func (application *Application) configureDB() {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	driver := os.Getenv("DB_DRIVER")
	sslMode := os.Getenv("DB_SSLMODE")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	application.configLocator.DBConfigInstance().Host = host
	application.configLocator.DBConfigInstance().Name = name
	application.configLocator.DBConfigInstance().User = user
	application.configLocator.DBConfigInstance().Password = password
	application.configLocator.DBConfigInstance().Port = port
	application.configLocator.DBConfigInstance().Driver = driver
	application.configLocator.DBConfigInstance().SslMode = sslMode
}

func (application *Application) configJWT() {
	signingKey := os.Getenv("SIGNING_KEY")

	application.configLocator.JWTConfig().SigningKey = signingKey
}

func (application *Application) Run() {
	port := application.configLocator.ServerConfigInstance().Port
	host := application.configLocator.ServerConfigInstance().Host

	log.Print("Server Start")
	if err := application.server.Run(host, fmt.Sprint(port), application.Router); err != nil {
		log.Fatal(err)
	}

}

func (application *Application) InitRouts() {
	application.Router.InitRouts()
}

/*
func (application *Application) InitSentry() {
	dsn := application.configLocator.SentryConfigInstance().Dsn
	release := application.configLocator.SentryConfigInstance().Release
	isDebug := application.configLocator.SentryConfigInstance().IsDebug

	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     release,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: isDebug,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
}*/
