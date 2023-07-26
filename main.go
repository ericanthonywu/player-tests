package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "PLayer",
		AppName:       "PLayer Service",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	// cors
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,User-Agent,Content-Length",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// add security header
	app.Use(helmet.New())
	// compress
	app.Use(compress.New())

	// add etag
	app.Use(etag.New())

	// add recover
	app.Use(recover.New())

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
