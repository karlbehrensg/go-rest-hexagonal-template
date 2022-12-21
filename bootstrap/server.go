package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	adaptersDB "github.com/karlbehrensg/go-rest-template/auth/adapters/db"
	auth "github.com/karlbehrensg/go-rest-template/auth/app"
	shareDb "github.com/karlbehrensg/go-rest-template/share/db"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hex_go"
)

func StartServer() {
	userDbClient, err := shareDb.NewPostgresClient(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	defer userDbClient.Close()

	userAdapterDb, err := adaptersDB.NewUserAdapter(userDbClient)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	auth.ApplyRoutes(app, &userAdapterDb)

	app.Listen(":3000")
}
