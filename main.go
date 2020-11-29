package main

import(
	"github.com/labstack/echo"
	handlerFunc"Weather_api/Handler"
)

func main(){
	e := echo.New()
	e.POST("api/weather", handlerFunc.Temperature )
	e.Logger.Fatal(e.Start(":8888"))
}