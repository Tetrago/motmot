package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/samber/lo"

	"github.com/tetrago/motmot/api/.gen/motmot/public/model"
	. "github.com/tetrago/motmot/api/.gen/motmot/public/table"
)

// Groups godoc
// @Summary Gets groups
// @Description Gets all public groups
// @Tags group
// @Produce json
// @Success 200 {array} GroupModel
// @Failure 500
// @Router /group/all [get]
func Groups(g *gin.Context) {
	var dest []model.Room

	stmt := SELECT(Room.AllColumns).FROM(Room)

	if err := stmt.Query(Database, &dest); err != nil {
		fmt.Printf("[/v1/group/all] Error querying database: %s\n", err.Error())
		g.Status(http.StatusInternalServerError)
		return
	}

	g.JSON(http.StatusOK, lo.Map(dest, MapToGroupModel))
}

func GroupHandler(r *gin.RouterGroup) {
	r.GET("/all", Groups)
}