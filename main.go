package main

import (
	"applyGo/internal/app"
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	port := flag.String("port", "8080", "default port: 8080")
	state := flag.String("state", "dev", "default state: dev")

	flag.Parse()

	config := &app.Configs{}
	err := config.InitWithState(state)
	if err != nil {
		fmt.Printf("err = %s\n", err)
	}

	fmt.Printf("config: %+v\n", config)

	e := echo.New()
	if err := app.InitRoute(e, config); err != nil {
		log.Panicln("new reoutes error:", err)
		return
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
	select {}
}
