package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	mysql2 "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/test", handle)
	log.Fatal(e.Start("localhost:8080"))
}

func handle(c echo.Context) error {
	log.Println("request received")
	ctx := c.Request().Context()
	context.AfterFunc(ctx, func() {
		log.Println(context.Cause(ctx))
	})

	err := findData(ctx)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func findData(ctx context.Context) error {
	cfg := mysql2.Config{DBName: "test", User: "root", Passwd: "root", AllowNativePasswords: true, ParseTime: true}
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.WithContext(ctx).Exec("DO SLEEP(5)").Error; err != nil {
		return err
	}
	return nil
}
