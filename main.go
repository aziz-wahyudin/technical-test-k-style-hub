package main

import (
	"aziz-wahyudin/technical-test-k-style-hub/config"
	"aziz-wahyudin/technical-test-k-style-hub/factory"
	"aziz-wahyudin/technical-test-k-style-hub/utils/database/mysql"

	"github.com/labstack/echo/v4"

	"fmt"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
