package main

import (
	"encoding/base64"
	"net/http"
	"worlder-assessment/handler"
	"worlder-assessment/storage"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	cfg := conf()
	st, err := storage.New(cfg)
	if err != nil {
		panic(err)
	}
	uname:= cfg.GetString("auth.username")
	pass:= cfg.GetString("auth.password")
	e := echo.New()
	e.Use(bindApp(st,uname,pass))
	e.Use(middlewareOne)
	e.GET("/data", handler.RetrieveSensor)
	e.POST("/insertsensors",handler.InsertSensorValue)
	e.Logger.Fatal(e.Start(":8088"))

}
func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		basicAuth := basicAuth(c.Get("uname").(string),c.Get("pass").(string))
		reqAuth := c.Request().Header.Get("Authorization")
		if(basicAuth!=reqAuth){
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic "+base64.StdEncoding.EncodeToString([]byte(auth))
}


func conf() *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.AddConfigPath(".")
	cfg.AutomaticEnv()
	cfg.SetConfigType("yml")

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}
	return cfg
}

func bindApp(db *sqlx.DB, uname string, pass string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			c.Set("uname",uname)
			c.Set("pass",pass)
			return next(c)
		}
	}
}
