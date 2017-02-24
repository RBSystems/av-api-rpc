package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/byuoitav/av-api-rpc/base"
	"github.com/byuoitav/av-api-rpc/helpers"
	"github.com/labstack/echo"
)

//GetRoomByNameAndBuildingHandler .
func GetRoomByNameAndBuildingHandler(context echo.Context) error {
	return context.JSON(http.StatusNotImplemented, "Not implemented.")
}

//SendRoomCommands .
func SendRoomCommands(context echo.Context) error {

	room := context.Param("room")
	building := context.Param("building")

	body := base.RPCRequest{}

	err := context.Bind(&body)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Invalid body.")
	}

	log.Printf("Room: %s Building %s", room, building)

	body.Room = room
	body.Building = building

	report, err := helpers.RunCommands(body)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	b, err := json.Marshal(report)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, string(b))
}
