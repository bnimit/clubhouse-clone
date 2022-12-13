package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimit-bhandari/clubhouse-clone/middleware"
	"github.com/nimit-bhandari/clubhouse-clone/routes"
)

func main() {
	app := fiber.New()

	middleware.SetMiddleWare(app)
	routes.SetupApiV1(app)

	port := "8000"
	addr := flag.String("addr", port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))

}
