package templates

const DiFile = `package providers

import (
	"{{.module}}/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(router.New)

	return Container
}`

const MainFile = `package main

import (
	"fmt"
	"log"

	"{{.module}}/cmd/providers"
	"{{.module}}/config"
	"{{.module}}/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
`

const ConfigFile = `package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {` + `
	ServerHost		string ` + "`required:" + `"true"` + ` split_words:"true"` + "`" + `
	ServerPort		int ` + "`required:" + `"true"` + ` split_words:"true"` + "`" + `
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		//If you use env file, uncomment this section and the functions:
		//if err := setEnvsFromFile(".env"); err != nil {
		//	log.Panicf("error reading local env file %s", err.Error())
		//	return
		//}

		if err := envconfig.Process("", &Cfg); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return Cfg
}
/*
func setEnvsFromFile(fileName string) error {
	root, err := rootDir()
	if err != nil {
		return err
	}
	file, err := os.Open(fmt.Sprintf("%s/%s", root, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return fmt.Errorf("env file is empty")
	}

	for scanner.Scan() {
		w := strings.Split(scanner.Text(), "=")

		if err = os.Setenv(w[0], w[1]); err != nil {
			return err
		}
	}

	return nil
}

func rootDir() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}
*/
`

const HealthFile = `package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Code    int    ` + "`json:" + `"status"` + "`" + `
	Message string ` + "`json:" + `"message"` + "`" + `
}

func HealthCheck(context echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return context.JSON(http.StatusOK, response)
}
`

const RouterFile = `package router

import (
	"{{.module}}/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server   *echo.Echo
}

func New(server *echo.Echo) *Router {
	return &Router{
		server,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api/microservice") //customize your basePath 
	basePath.GET("/health", handler.HealthCheck)
}
`
