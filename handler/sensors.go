package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"worlder-assessment/storage"
)

type InsertRequest struct {
	ID1 string `json:"id1"`
	ID2 string `json:"id2"`
	SensorValue int `json:"sensor_value"`
}

func InsertSensorValue(c echo.Context) error {
	i := c.Get("db")
	v := i.(*sqlx.DB)

	req := new(InsertRequest)
	if err := c.Bind(req);err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	err:=storage.Insert(v,storage.Sensor{
		ID1:         req.ID1,
		ID2:         req.ID2,
		SensorValue: req.SensorValue,
	})
	if(err!=nil){
		return c.String(http.StatusInternalServerError, "")
	}else{
		return c.String(http.StatusOK, "")
	}
}

func RetrieveSensor(c echo.Context) error {
	i := c.Get("db")
	v := i.(*sqlx.DB)

	id1 := c.QueryParam("ID1")
	id2 := c.QueryParam("ID2")
	from := c.QueryParam("start_timestamp")
	to := c.QueryParam("end_timestamp")

	intFrom, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		intFrom = 0
	}
	intTo, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		intTo = 0
	}

	result, err :=storage.Retrieve(v,storage.SensorReq{
		ID1:  id1,
		ID2:  id2,
		From: intFrom,
		To:   intTo,
	})
	if(err!=nil){
		return c.String(http.StatusInternalServerError, "")
	}else{
		return c.JSON(http.StatusOK,result)
	}
}
