package main

import (
  "github.com/labstack/echo"

  "net/http"
)

type User struct {
  Status    int    `json:"status"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Age       int    `json:"age"`
  Address struct {
    StateName string `json:"stateName"`
    City string `json:"city"`
    Line string `json:"line"`
    Line2 string `json:"line2"`
    BuildingName string `json:"buildingName"`
  } `json:"address"`
}

func getHome(c echo.Context) error {
  u := new(User)
  u.Status = 200
  u.FirstName = "taro"
  u.LastName = "tanaka"
  u.Age = 24
  u.Address.StateName = "tokyo"
  u.Address.City = "sibuya-ku"
  u.Address.Line = "sibuya 1"
  u.Address.Line2 = "1-1"
  u.Address.BuildingName = "Sibuya Tower"
  return c.JSON(http.StatusOK, u)
}

func getUserName(c echo.Context) error {
  return c.String(http.StatusOK, "Nishikawa Keien")
}

func main(){
  e := echo.New()
  e.GET("/", getHome)
  e.GET("/user", getUserName)
  e.Logger.Fatal(e.Start(":0421"))
}

